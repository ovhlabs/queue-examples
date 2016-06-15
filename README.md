# OVH Paas Queue client examples

[![Build Status](https://travis-ci.org/runabove/queue-examples.svg?branch=master)](https://travis-ci.org/runabove/queue-examples)

* [Golang](golang)
* [NodeJs](nodejs)
* [Python](python)
* [Scala](scala_kafka_0.8.2)

## Basic concepts

Each client example can be run in 2 modes: `produce` or `consume`.

The two options `--key` and `--topic` are always mandatory:

- `--key` to configure the key to be authenticated to the OVH Paas Queue.
- `--topic` to declare the authorized topic to produce to.

Depending the client example needs the Kafka or the Zookeeper seed URL is used:

- `--kafka` for the Kafka broker seed URL
- `--zk` for the Zookeeper broker seed URL

In the `consume` mode:

- `--group` is used to configure the consumer group id

***Important to notice***:
- ***The key is used as the Kafka client.id or the Zookeeper root path.***
- ***The topic must be prefixed with the human application id corresponding to the key.***

## Golang

### Requirements

* Go >= 1.5

### Setup

Build Go binary:

    make build-go

### Consume data

    golang/bin/qaas-client-darwin-amd64 consume \
        --kafka $HOST:9092 --key $KEY --topic $PREFIX.$TOPIC --group ${PREFIX}.golang-${GROUP}

### Produce data

    golang/bin/qaas-client-darwin-amd64 produce \
        --kafka $HOST:9092 --key $KEY --topic $PREFIX.$TOPIC

# Node.js

### Requirements

* Node.js
* NPM

### Setup

Install the nodejs dependencies:

    make node-install-deps
    cd nodejs

### Produce data

    node client.js produce \
        --zk $HOST:2181 --key $KEY --topic $PREFIX.$TOPIC

### Consume data

    node client.js consume \
        --zk $HOST:2181 --key $KEY --topic $PREFIX.$TOPIC --group ${PREFIX}.nodejs-${GROUP}

## Python

### Requirements

* Python >= 2.7

### Setup

Install the python dependencies:

    make python-install-deps
    cd python

### Produce

    python client.py produce \
        --kafka $HOST:9092 --key $KEY --topic $PREFIX.$TOPIC

### Consume

    python client.py consume \
        --kafka $HOST:9092 --key $KEY --topic $PREFIX.$TOPIC --group ${PREFIX}.python-${GROUP}

## Scala

This example uses akka-reactive-streams and kafka 0.8.2.1.

### Setup

* Scala
* SBT

### Consume data

    sbt "run produce \
        --kafka $HOST:9092 --zk $HOST:2181 --key $KEY --topic $PREFIX.$TOPIC"

### Produce data

    sbt "run consume \
        --kafka $HOST:9092 --zk $HOST:2181 --key $KEY --topic $PREFIX.$TOPIC --group ${PREFIX}.scala-${GROUP}"


## Docker

Each example can be build using Docker.

    make build-docker