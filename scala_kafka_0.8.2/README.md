# Scala QaaS example with Kafka 0.8.2

This example uses akka-reactive-streams and kafka 0.8.2.1.

## Consume data

    sbt "run -z $HOST:2181 -k $HOST:9092 -i $KEY_ID -s $KEY_SECRET -t applicationid.topic1 -m cons"

## Produce data

    sbt "run -z $HOST:2181 -k $HOST:9092 -i $KEY_ID -s $KEY_SECRET -t applicationid.topic1 -m prod"
