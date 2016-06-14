var cliparse = require("cliparse");
var kafka = require('kafka-node');
var readline = require('readline');
var parsers = cliparse.parsers;

var rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout,
  terminal: false
});

function prodModule(params) {
  var HighLevelProducer = kafka.HighLevelProducer
  var zkUrl = params.options.host+"/"+params.options.key
  var client = new kafka.Client(
    zkUrl,
    params.options.key
  )
  var producer = new HighLevelProducer(client)
  producer.on('ready', function () {
    rl.on('line', function(line){
      payloads = [
          { topic: params.options.topic, messages: line }
      ];
      producer.send(payloads, function (err, data) {
        kv = data[params.options.topic]
        Object.keys(kv).forEach(function(p) {
          console.log("> message sent to partition "+p+" at offset "+kv[p]);
        })
      });
    })
  });
}

function consModule(params) {
  var Consumer = kafka.Consumer
  var zkUrl = params.options.host+"/"+params.options.key
  var client = new kafka.Client(
    zkUrl,
    params.options.key
  )
  consumer = new Consumer(
    client,
    [ { topic: params.options.topic} ],
    {
      groupId: params.options.groupid,
      autoCommit: true,
      autoCommitIntervalMs: 500,
      fetchMaxWaitMs: 100,
      fetchMinBytes: 1,
      fetchMaxBytes: 1024 * 10,
      encoding: 'utf8'
    }
  )
  consumer.on('message', function (message) {
      console.log(message.value);
  });

}
var options = [
  cliparse.option("host", { description: "zookeeper connection string", default: "127.0.0.1:2181"}),
  cliparse.option("key", { description: "key", default: ""}),
  cliparse.option("group-id", { description: "group id", default: "qaas-node-client-group"}),
  cliparse.option("topic", { description: "topic to push to", default: "topic1"})
]


var cliParser = cliparse.cli({
  name: "qaas-client",
  description: "Simple node js producer/consumer",
  commands: [
    cliparse.command(
      "consume",
      {
        description: "consume message on the given topic",
        options: options
      },
      consModule),

    cliparse.command(
      "produce",
      {
        description: "produce message on the given topic from stdin",
        options: options
      },
      prodModule)
  ]
});

cliparse.parse(cliParser);
