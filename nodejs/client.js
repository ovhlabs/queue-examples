var cliparse = require("cliparse");
var readline = require('readline');
var kafka = require('node-rdkafka');

var stdin = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false
});

function produce(params) {

  var producer = new kafka.Producer({
    'metadata.broker.list': params.options.broker,
    'security.protocol': 'SASL_SSL',
    'ssl.ca.location': '/etc/ssl/certs',
    'api.version.request': 'true',
    'sasl.mechanisms': 'PLAIN',
    'sasl.username': params.options.username,
    'sasl.password': params.options.password,
    'dr_cb': true // Specifies that we want a delivery-report event to be generated
  });

  producer.on('event.log', function(log) {
    console.log('log', log);
  });

  producer.on('error', function(err) {
    console.error('Error on producer', err);
  });

  producer.on('delivery-report', function(report) {
    console.log('Producer delivery-report', report);
  });

  producer.connect();

  producer.on('ready', function() {
    console.log('Ready to produce messages. Write something to stdin...')
    try {

      var topic = producer.Topic(params.options.topic, {
       // Make the Kafka broker acknowledges each message (optional)
       'request.required.acks': 1
      });

      stdin.on('line', function(line) {
        // if partition is set to -1, the default partitioner is used
        var partition = -1;
        var value = new Buffer(line);
        producer.produce(topic, partition, value);
      });

    } catch (err) {
      console.error('Fail to produce message', err);
    }
  });
}

function consume(params) {
  if (!params.options["consumer-group"]) {
    params.options["consumer-group"] = params.options.username + ".node";
  }

  var consumer = new kafka.KafkaConsumer({
    'metadata.broker.list': params.options.broker,
    'security.protocol': 'SASL_SSL',
    'ssl.ca.location': '/etc/ssl/certs',
    'api.version.request': true,
    'debug': 'protocol,security,broker',
    'sasl.mechanisms': 'PLAIN',
    'sasl.username': params.options.username,
    'sasl.password': params.options.password,
    'group.id': params.options["consumer-group"]
  });

  consumer.connect();

  consumer.on('ready', function() {
    console.log('Ready to consume messages...');

    consumer.consume(params.options.topic, function(err, message) {
      console.log(message.value.toString());
    });

  });
}

var options = [
  cliparse.option("broker", { description: "Kafka broker address"}),
  cliparse.option("topic", { description: "Topic"}),
  cliparse.option("username", { description: "SASL username"}),
  cliparse.option("password", { description: "SASL password"}),
  cliparse.option("consumer-group", { description: "Consumer group"}),
];

cliparse.parse(
  cliparse.cli({
    name: "kafka-client",
    description: "Node.js Kafka client to produce/consume using SASL/SSL",
    commands: [
      cliparse.command("produce", { description: "Produce messages", options: options }, produce),
      cliparse.command("consume", { description: "Consume messages", options: options }, consume)
    ]
  })
);
