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
                  host: String = "127.0.0.1:9092",
                  key: String = "",
                  topic: String = "qaas-client-topic",
                  groupId: String = "qaas-client-group-id",
                  mode: String = "prod")

object Main {
  val parser = new scopt.OptionParser[Config]("scopt") {
    head("qaas-scala-client", "0.1.0")
    opt[String]('z', "zk") action { (x, c) ⇒ c.copy(zk = x) } text ("Zookeeper address")
    opt[String]('h', "host") action { (x, c) ⇒ c.copy(host = x) } text ("Kafka address")
    opt[String]('k', "key") action { (x, c) ⇒ c.copy(key = x) } text ("Authentication key id")
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
        logger.info("Zk : {}", s"${config.zk}/${config.key}")
        logger.info("host : {}", config.host)
        logger.info("key: {}", config.key)
        logger.info("topic : {}", config.topic)
        logger.info("mode : {}", config.mode)

        val consumerProperties = ConsumerProperties(
          brokerList = config.host,
          zooKeeperHost = s"${config.zk}/${config.key}",
          topic = config.topic,
          groupId = config.groupId,
          decoder = new StringDecoder()).setProperty("client.id", s"${config.key}").setProperty("auto.commit.enable", "true").setProperty("auto.commit.interval.ms", "500")

        val producerProperties = ProducerProperties(
          brokerList = config.host,
          topic = config.topic,
          encoder = new StringEncoder()).setProperty("client.id", s"${config.key}").setProperty("auto.commit.enable", "true")

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
        }

      case None ⇒
        logger.error("Fail to parse configuration")

    }
  }

}
