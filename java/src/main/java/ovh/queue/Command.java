package ovh.queue;

import com.beust.jcommander.Parameter;

abstract class Command {
  @Parameter(names="--broker", description="Kafka broker address", required=true)
  public String broker;

  @Parameter(names="--topic", description="Topic", required=true)
  public String topic;
}
