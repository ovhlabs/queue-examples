package ovh.queue;

import java.util.Arrays;
import java.util.Properties;
import com.beust.jcommander.Parameters;
import com.beust.jcommander.Parameter;
import org.apache.kafka.clients.consumer.KafkaConsumer;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import ovh.queue.Command;

@Parameters(commandDescription = "Consume messages using TLS/SSL")
public class Consume extends Command {

  public static String NAME = "consume";

  @Parameter(names="--consumer-group", description = "Consumer group", required=true)
  public String group;

  public void run() {
    Properties props = new Properties();
    props.put("bootstrap.servers", this.broker);
    props.put("security.protocol", "SASL_SSL");
    props.put("sasl.mechanism", "PLAIN");
    props.put("enable.auto.commit", "true");
    props.put("auto.commit.interval.ms", "1000");
    props.put("session.timeout.ms", "30000");
    props.put("key.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
    props.put("value.deserializer", "org.apache.kafka.common.serialization.StringDeserializer");
    props.put("group.id", this.group);

    KafkaConsumer<String, String> consumer = new KafkaConsumer<>(props);

    try {
      consumer.subscribe(Arrays.asList(this.topic));

      System.out.println("Ready to consume messages...");
      while (true) {
        ConsumerRecords<String, String> records = consumer.poll(100);
        for (ConsumerRecord<String, String> record : records)
          System.out.printf("offset = %d, key = %s, value = %s\n", record.offset(), record.key(), record.value());
      }
    } finally {
      consumer.close();
    }
  }
}
