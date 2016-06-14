## QaaS client examples

[![Build Status](https://travis-ci.org/runabove/queue-examples.svg?branch=master)](https://travis-ci.org/runabove/queue-examples)

* [Go](golang)
* [NodeJs](nodejs)
* [Python](python)
* [Scala](scala_kafka_0.8.2)

# Basic concepts

Every client uses two constants :

- Key   : The key used for the authentication. This key is linked to your application.
- Topic : The topic to which data will be pushed. This topic should be prefixed
    by your application id

Every example uses the key as a client id to authenticate and pushed on a topic prefixed by the application id.

# Golang

## Requirements

* Go >= 1.5

## Setup

Build Go binary:

    make build-go

## Consume data

    golang/bin/qaas-client-darwin-amd64 consume --host $HOST:9092 --key $KEY --topic $PREFIX.$TOPIC --group ${PREFIX}.golang-${GROUP}

## Produce data

    golang/bin/qaas-client-darwin-amd64 produce --host $HOST:9092 --key $KEY --topic $PREFIX.$TOPIC

# Node.js

## Requirements

* Node.js
* NPM

## Setup

Install Node dependencies:

  make build-node

## Produce data

    node client.js produce --host $HOST:2181 --key $KEY --topic $PREFIX.$TOPIC

## Consume data

    node client.js consume --host $HOST:2181 --key $KEY --topic $PREFIX.$TOPIC --group ${PREFIX}.nodejs-${GROUP}

# Python

## Requirements

* Python version 2.7

## Setup

    make build-python

## Produce

~~~
 python client.py produce --host $HOST:9092 --key $KEY --topic $PREFIX.$TOPIC
~~~

## Consume

~~~
 python client.py consume --host $HOST:9092 --key $KEY --topic $PREFIX.$TOPIC --group ${PREFIX}.python-${GROUP}
 ~~~

# Scala

This example uses akka-reactive-streams and kafka 0.8.2.1.

## Consume data

    sbt "run -z $HOST:2181 -k $HOST:9092 -i $KEY_ID -s $KEY_SECRET -t applicationid.topic1 -m cons"

## Produce data

    sbt "run -z $HOST:2181 -k $HOST:9092 -i $KEY_ID -s $KEY_SECRET -t applicationid.topic1 -m prod"
