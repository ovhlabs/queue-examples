import argparse
import fileinput
import sys
import ssl
from kafka import KafkaProducer, KafkaClient, KafkaConsumer
from kafka.errors import KafkaError

def produce(args):
    try:
        producer = KafkaProducer(
                bootstrap_servers = [args.broker],
                security_protocol="SASL_SSL",
                ssl_context=ssl.SSLContext(ssl.PROTOCOL_TLSv1),
                sasl_mechanism="PLAIN",
                sasl_plain_username=args.username,
                sasl_plain_password=args.password)

        print "Ready to produce messages. Write something to stdin..."
        while True:
            value = sys.stdin.readline().rstrip().encode('UTF-8')
            future = producer.send(args.topic.encode('UTF-8'), value)
            try:
                r = future.get(timeout=10)
                print "> Message '%s' sent to partition %d at offset %d" % (value, r.partition, r.offset)
            except KafkaError:
                print "Fail to produce message to kafka"
                pass

    except KeyboardInterrupt:
        producer.close()
        sys.exit(0)

def consume(args):
    try:
        if args.consumer_group == None:
            args.consumer_group = args.username + ".py"

        consumer = KafkaConsumer(args.topic,
                bootstrap_servers=[args.broker],
                security_protocol="SASL_SSL",
                ssl_context=ssl.SSLContext(ssl.PROTOCOL_TLSv1),
                sasl_mechanism="PLAIN",
                sasl_plain_username=args.username,
                sasl_plain_password=args.password,
                enable_auto_commit=True,
                auto_commit_interval_ms=1000,
                group_id=args.consumer_group)

        print "Ready to consume messages..."
        for message in consumer:
            print message.value

    except KeyboardInterrupt:
        consumer.close()
        sys.exit(0)

parser = argparse.ArgumentParser(description='Python Kafka client example')
subparsers = parser.add_subparsers()

c = subparsers.add_parser('consume')
c.add_argument('--broker', metavar='broker', help='Kafka broker address')
c.add_argument('--topic', metavar='topic', help='Topic')
c.add_argument('--username', metavar='username', help='SASL username')
c.add_argument('--password', metavar='password', help='SASL password')
c.add_argument('--consumer-group', metavar='consumer_group', help='Consumer group')
c.set_defaults(func=consume)

p = subparsers.add_parser('produce')
p.add_argument('--broker', metavar='broker', help='Kafka broker address')
p.add_argument('--topic', metavar='topic', help='Topic')
p.add_argument('--username', metavar='username', help='SASL username')
p.add_argument('--password', metavar='password', help='SASL password')
p.add_argument('--consumer-group', metavar='consumer_group', help='Consumer group')
p.set_defaults(func=produce)

args = parser.parse_args()
args.func(args)
