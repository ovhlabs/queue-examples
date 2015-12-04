import argparse
from kafka import SimpleProducer, KafkaClient, KafkaConsumer
import sys

parser = argparse.ArgumentParser(description='QAAS producer/consumer')
parser.add_argument('-h', '--host', metavar='host', help='Kafka broker address')
parser.add_argument('-t', '--topic', metavar='name', help='Destination topic')
parser.add_argument('-k', '--key', metavar='key', help='Authentication key')
parser.add_argument('-m', '--mode', metavar='mode',  help='Client mode (prod or cons)')
parser.add_argument('-g', '--group', metavar='group',  help='Group for consumer', default='python_group_id')
args = parser.parse_args()

client = KafkaClient(args.kafka, client_id=args.key)

try:
    if args.mode == "prod":
        producer = SimpleProducer(client)
        for num in range(1, 100):
                message = "message_" + str(num)
                producer.send_messages(args.topic.encode('UTF-8'), message.encode('UTF-8'))
                print "message <" + message + "> produced"
    else:
        consumer = KafkaConsumer(args.topic,
            client_id=client_id,
            group_id=args.group,
            bootstrap_servers=[args.host],
            enable_auto_commit=True,
            auto_commit_interval_ms=5 * 1000,
            auto_offset_reset='largest')

        for message in consumer:
                print("New message %s" % (message.value))
except KeyboardInterrupt:
    sys.exit(0)
