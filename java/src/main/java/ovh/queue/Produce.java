package ovh.queue;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Properties;
import com.beust.jcommander.Parameters;
import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerRecord;
import ovh.queue.Command;

@Parameters(commandDescription = "Produce messages using TLS/SSL")
public class Produce extends Command {

  public static String NAME = "produce";

  public void run() {
    Properties props = new Properties();
    props.put("bootstrap.servers", this.broker);
    props.put("security.protocol", "SASL_SSL");
    props.put("sasl.mechanism", "PLAIN");
    props.put("acks", "all");
    props.put("retries", 0);
    props.put("batch.size",1 );
    props.put("linger.ms", 1);
    props.put("buffer.memory", 33554432);
    props.put("key.serializer", "org.apache.kafka.common.serialization.StringSerializer");
    props.put("value.serializer", "org.apache.kafka.common.serialization.StringSerializer");

    Producer<String, String> producer = new KafkaProducer<>(props);
    BufferedReader br = null;
    try {
      System.out.println("Ready to produce messages. Write something to stdin...");
      br = new BufferedReader(new InputStreamReader(System.in));
      while (true) {
        String input = br.readLine();
        producer.send(new ProducerRecord<String, String>(this.topic, "", input));
      }
    } catch (IOException e) {
      e.printStackTrace();
    } finally {
      if (br != null) {
        try {
          br.close();
        } catch (IOException e) {
          e.printStackTrace();
        }
      }
    }
  }
}
