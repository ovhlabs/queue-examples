# OVH Queue - Kafka examples

[![Build Status](https://travis-ci.org/runabove/queue-examples.svg?branch=master)](https://travis-ci.org/runabove/queue-examples)

Here are Kafka client examples using SASL/SSL in several languages and tested on [OVH Queue](https://www.runabove.com/dbaas-queue.xml) ):

* [Go](go)
* [Python](python)
* [Java](java)
* [NodeJs](nodejs)

Each example can be run to consume messages or produce messages from STDIN
and requires four flags:

    --broker    the Kafka broker adress
    --username  the SASL username
    --password  the SASL password
    --topic     the Kafka topic

***When using OVH Queue the consumer group must be prefixed by the username*** using the
`--consumer-group` flag (ex: --username collector.admin --consumer-group collector.admin.group).

## Go

##### Requirements

* Go >= 1.5

##### Setup

Build binary:

```
cd go
go build -o kafka-client
```

##### Consume

```
kafka-client consume \
    --broker $HOST:9093 \
    --username $SASL_USERNAME --password $SASL_PASSWORD \
    --topic $TOPIC --consumer-group $SASL_USERNAME.group-go
```

##### Produce

```
kafka-client produce \
    --broker $HOST:9093 \
    --username $SASL_USERNAME --password $SASL_PASSWORD \
    --topic $TOPIC
```

## Python

##### Requirements

* Python >= 2.7

##### Setup

Install dependencies:

```
cd python
pip install --upgrade kafka-python==1.3.1
```

##### Consume

```
python client.py consume \
    --broker $HOST:9093 \
    --username $SASL_USERNAME --password $SASL_PASSWORD \
    --topic $TOPIC --consumer-group $SASL_USERNAME.group-python
```

##### Produce

```
python client.py produce \
    --broker $HOST:9093 \
    --username $SASL_USERNAME --password $SASL_PASSWORD \
    --topic $TOPIC
```

## Java

##### Requirements

* Java >= 1.8
* Maven >= 3

##### Setup

The Java client needs a `jaas.conf` file containing SASL login information.

```
KafkaClient {
   org.apache.kafka.common.security.plain.PlainLoginModule required
   username="<YOUR_SASL_USERNAME>"
   password="<YOUR_SASL_PASSWORD>";
};
```

Build the JAR file:

```
cd java
mvn compile package
```

##### Consume

```
java \
	-cp target/kafka-example-jar-with-dependencies.jar \
	-Djava.security.auth.login.config=jaas.conf \
	ovh.queue.Main consume \
	--broker $HOST:9093 \
	--topic $TOPIC --consumer-group $SASL_USERNAME.group-java
```

##### Produce

```
java \
    -cp target/kafka-example-jar-with-dependencies.jar \
    -Djava.security.auth.login.config=jaas.conf \
    ovh.queue.Main produce \
    --broker $HOST:9093 \
    --topic $TOPIC
```

## Node.js

##### Requirements

* Node.js >= 4.x
* NPM >= 2.x

##### Setup

Install dependencies:

```
cd nodejs
npm install
```

##### Consume

```
node client.js consume \
    --broker $HOST:9093 \
    --username $SASL_USERNAME --password $SASL_PASSWORD \
    --topic $TOPIC --consumer-group $SASL_USERNAME.group-node
```

##### Produce

```
node client.js produce \
    --broker $HOST:9093 \
    --username $SASL_USERNAME --password $SASL_PASSWORD \
    --topic $TOPIC
```
