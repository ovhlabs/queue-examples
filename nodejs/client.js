var cliparse = require("cliparse");
var kafka = require('kafka-node');
var parsers = cliparse.parsers;

function prodModule(params) {
  console.log("prod")
  console.log(params)
  var HighLevelProducer = kafka.HighLevelProducer
  var zkUrl = params.options.zk+"/"+params.options.auth
  var client = new kafka.Client(
    zkUrl,
    params.options.auth
  )
  var producer = new HighLevelProducer(client)
  payloads = [
      { topic: params.options.topic, messages: 'hi' }
  ];
  producer.on('ready', function () {
      console.log("Producer ready")
      producer.send(payloads, function (err, data) {
          console.log(data);
      });
  });
}

function consModule(params) {
  console.log("cons")
  console.log(params)
  var Consumer = kafka.Consumer
  var zkUrl = params.options.zk+"/"+params.options.auth
  var client = new kafka.Client(
    zkUrl,
    params.options.auth
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
      // fromOffset: false,
      encoding: 'utf8'
    }
  )
  consumer.on('message', function (message) {
      console.log(message);
  });

}
var options = [
  cliparse.option("zk", { aliases: ["z"], description: "zookeeper connection string", default: "127.0.0.1:2181"}),
  cliparse.option("key", { aliases: ["k"], description: "key", default: ""}),
  cliparse.option("group-id", { aliases: ["g"], description: "group id", default: "qaas-node-client-group"}),
  cliparse.option("topic", { aliases: ["t"], description: "topic to push to", default: "topic1"})
]


var cliParser = cliparse.cli({
  name: "qaas-client",
  description: "Simple node js producer/consumer",
  commands: [
    cliparse.command(
      "cons",
      {
        description: "consume message on the given topic",
        options: options
      },
      consModule),

    cliparse.command(
      "prod",
      {
        description: "produce message on the given topic",
        options: options
      },
      prodModule)
  ]
});

cliparse.parse(cliParser);
