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
                  key: String = "",
                  topic: String = "qaas-client-topic",
                  groupId: String = "qaas-client-group-id",
                  mode: String = "consume")

object Main {
  val parser = new scopt.OptionParser[Config]("qaas-client") {
    head("qaas-scala-client", "0.1.0")
    opt[String]("zk") action { (x, c) ⇒ c.copy(zk = x) } text ("Zookeeper address")
    opt[String]("kafka") action { (x, c) ⇒ c.copy(kafka = x) } text ("Kafka address")
    opt[String]("key") action { (x, c) ⇒ c.copy(key = x) } text ("Authentication key id")
    opt[String]("topic") action { (x, c) ⇒ c.copy(topic = x) } text ("Topic to read and write from")
    opt[String]("group") action { (x, c) ⇒ c.copy(groupId = x) } text ("Group ip to use")
    cmd("consume").action((_, c) ⇒ c.copy(mode = "consume"))
    cmd("produce").action((_, c) ⇒ c.copy(mode = "produce"))
  }

  val logger = org.slf4j.LoggerFactory.getLogger(Main.getClass)

  def main(args: Array[String]): Unit = {

    parser.parse(args, Config()) match {
      case Some(config) ⇒
        implicit val actorSystem = ActorSystem("ReactiveKafka")
        implicit val materializer = ActorMaterializer()
        logger.info("Zk : {}", s"${config.zk}/${config.key}")
        logger.info("Kafka : {}", config.kafka)
        logger.info("key: {}", config.key)
        logger.info("topic : {}", config.topic)
        logger.info("mode : {}", config.mode)

        val consumerProperties = ConsumerProperties(
          brokerList = config.kafka,
          zooKeeperHost = s"${config.zk}/${config.key}",
          topic = config.topic,
          groupId = config.groupId,
          decoder = new StringDecoder()).setProperty("client.id", s"${config.key}").setProperty("auto.commit.enable", "true").setProperty("auto.commit.interval.ms", "500")

        val producerProperties = ProducerProperties(
          brokerList = config.kafka,
          topic = config.topic,
          encoder = new StringEncoder()).setProperty("client.id", s"${config.key}").setProperty("auto.commit.enable", "true")

        val kafka = new ReactiveKafka()
        config.mode match {
          case "produce" ⇒
            val producer: Subscriber[String] = kafka.publish(producerProperties)
            Source(Stream.from(1)).map(_ ⇒ io.StdIn.readLine("")).to(Sink(producer)).run()
          case "consume" ⇒
            val consumer: Publisher[StringKafkaMessage] = kafka.consume(consumerProperties)
            Source(consumer).runForeach(x ⇒ println(x.message()))
          case _ ⇒
            logger.error("Cannot perform mode {}", config.mode)
        }

      case None ⇒
        logger.error("Fail to parse configuration")

    }
  }

}
