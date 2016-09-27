package ovh.queue;

import com.beust.jcommander.JCommander;
import ovh.queue.Produce;
import ovh.queue.Consume;

public class Main {

  public static void main(String []args) {

    JCommander commander = new JCommander();
    Produce producer = new Produce();
    Consume consumer = new Consume();

    commander.addCommand(Produce.NAME, producer);
    commander.addCommand(Consume.NAME, consumer);
    commander.parse(args);

    String command = commander.getParsedCommand();

    if (command == null) {
      commander.usage();
    } else if (command.equals(Produce.NAME)) {
      producer.run();
    } else if (command.equals(Consume.NAME)) {
      consumer.run();
    }
  }
}
