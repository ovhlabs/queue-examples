package ovh.queue;

import com.beust.jcommander.Parameter;

abstract class Command {
  @Parameter(names="--broker", description="Kafka broker address", required=true)
  public String broker;

  @Parameter(names="--topic", description="Topic, prefixed by your namespace (eg. --topic=myns.topic)", required=true)
  public String topic;
}
