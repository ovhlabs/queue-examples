import argparse
import fileinput
from kafka import SimpleProducer, KafkaClient, KafkaConsumer
import sys

def produce(args):
    client = KafkaClient(args.kafka, client_id=args.key)
    try:
        producer = SimpleProducer(client)
        while True:
            line = sys.stdin.readline()
            result = producer.send_messages(args.topic.encode('UTF-8'), line.rstrip().encode('UTF-8'))
            for r in result:
                print "> message sent to partition %d at offset %d " % (r.partition, r.offset)

    except KeyboardInterrupt:
        producer.close()
        client.close()
        sys.exit(0)

def consume(args):
    client = KafkaClient(args.kafka, client_id=args.key)
    try:
        consumer = KafkaConsumer(args.topic,
            client_id=args.key,
            group_id=args.group,
            bootstrap_servers=[args.kafka],
            enable_auto_commit=True,
            auto_commit_interval_ms=1000)

        for message in consumer:
            print "%s" % message.value

    except KeyboardInterrupt:
        consumer.close()
        client.close()
        sys.exit(0)

parser = argparse.ArgumentParser(description='QAAS producer/consumer')
subparsers = parser.add_subparsers()
c = subparsers.add_parser('consume')
c.add_argument('--kafka', metavar='kafka', help='Kafka broker address')
c.add_argument('--topic', metavar='name', help='Destination topic')
c.add_argument('--key', metavar='key', help='Authentication key')
c.add_argument('--group', metavar='group',  help='Group for consumer', default='python_group_id')
c.set_defaults(func=consume)

p = subparsers.add_parser('produce')
p.add_argument('--kafka', metavar='kafka', help='Kafka broker address')
p.add_argument('--topic', metavar='name', help='Destination topic')
p.add_argument('--key', metavar='key', help='Authentication key')
p.set_defaults(func=produce)

args = parser.parse_args()
args.func(args)
