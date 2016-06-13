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

    golang/bin/qaas-client-darwin-amd64 consume --key $KEY --topic $TOPIC --host $HOST

## Produce data

    golang/bin/qaas-client-darwin-amd64 produce --key $KEY --topic $TOPIC --host $HOST

# Node.js

## Requirements

* Node.js
* NPM

## Setup

Install Node dependencies:

  make build-node

## Consume data

    node client.js cons -k $HOST:9092 -a $KEY -t applicationid.topic1

## Produce data

    node client.js prod -k $HOST:9092 -a $KEY -t applicationid.topic1

# Python

## Requirements

* Python version 2.7

## Setup

    make build-python

## Produce

~~~
 python client.py --key "mysecret" --host 192.168.99.100:9092 --mode prod --group "test" --topic "applicationid.test"
~~~

## Consume

~~~
 python client.py --key "mysecret" --host 192.168.99.100:9092 --mode cons --group "test" --topic "applicationid.test"
~~~

# Scala

This example uses akka-reactive-streams and kafka 0.8.2.1.

## Consume data

    sbt "run -z $HOST:2181 -k $HOST:9092 -i $KEY_ID -s $KEY_SECRET -t applicationid.topic1 -m cons"

## Produce data

    sbt "run -z $HOST:2181 -k $HOST:9092 -i $KEY_ID -s $KEY_SECRET -t applicationid.topic1 -m prod"
