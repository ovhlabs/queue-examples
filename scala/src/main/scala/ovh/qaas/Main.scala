package ovh.qaas

import scopt._
import akka.actor.ActorSystem
import akka.stream.ActorMaterializer
import akka.stream.scaladsl.{ Sink, Source }
import kafka.serializer.{ StringDecoder, StringEncoder }
import org.reactivestreams.{ Publisher, Subscriber }

import com.softwaremill.react.kafka._
import com.softwaremill.react.kafka.KafkaMessages._

case class Config(zk: String = "127.0.0.1:2181",
                  kafka: String = "127.0.0.1:9092",
                  auth: String = "",
                  topic: String = "qaas-client-topic",
                  groupId: String = "qaas-client-group-id",
                  mode: String = "prod")

object Main {
  val parser = new scopt.OptionParser[Config]("scopt") {
    head("qaas-scala-client", "0.1.0")
    opt[String]('z', "zk") action { (x, c) ⇒ c.copy(zk = x) } text ("Zookeeper address")
    opt[String]('k', "kafka") action { (x, c) ⇒ c.copy(kafka = x) } text ("Kafka address")
    opt[String]('a', "auth") action { (x, c) ⇒ c.copy(auth = x) } text ("Authentication (tokenid-tokenkey)")
    opt[String]('t', "topic") action { (x, c) ⇒ c.copy(topic = x) } text ("Topic to read and write from")
    opt[String]('g', "group-id") action { (x, c) ⇒ c.copy(groupId = x) } text ("Group ip to use")
    opt[String]('m', "mode") action { (x, c) ⇒ c.copy(mode = x) } text ("mode")
  }

  val logger = org.slf4j.LoggerFactory.getLogger(Main.getClass)
  def main(args: Array[String]): Unit = {

    parser.parse(args, Config()) match {
      case Some(config) ⇒
        implicit val actorSystem = ActorSystem("ReactiveKafka")
        implicit val materializer = ActorMaterializer()
        logger.info("Zk : {}", s"${config.zk}/${config.auth}")
        logger.info("kafka : {}", config.kafka)
        logger.info("auth : {}", config.auth)
        logger.info("topic : {}", config.topic)
        logger.info("mode : {}", config.mode)

        val consumerProperties = ConsumerProperties(
          brokerList = config.kafka,
          zooKeeperHost = s"${config.zk}/${config.auth}",
          topic = config.topic,
          groupId = config.groupId,
          decoder = new StringDecoder()).setProperty("client.id", config.auth).setProperty("auto.commit.enable", "true").setProperty("auto.commit.interval.ms", "500")

        val producerProperties = ProducerProperties(
          brokerList = config.kafka,
          topic = config.topic,
          encoder = new StringEncoder()).setProperty("client.id", config.auth).setProperty("auto.commit.enable", "true")

        val kafka = new ReactiveKafka()
        config.mode match {
          case "prod" ⇒
            val producer: Subscriber[String] = kafka.publish(producerProperties)
            Source(Stream.from(1)).map(_.toString()).to(Sink(producer)).run()
          case "cons" ⇒
            val consumer: Publisher[StringKafkaMessage] = kafka.consume(consumerProperties)
            Source(consumer).runForeach(println)
          case _ ⇒
            logger.error("Cannot perform mode {}", config.mode)
          //TODO correctly exit
        }

      case None ⇒
        logger.error("Fail to parse configuration")
      //TODO correctly exit

    }
  }

}
