# Node.js QaaS example

## Requirements

* Node.js
* NPM

## Setup

Install Node dependencies:

  make install

## Consume data

    node client.js cons -k $HOST:9092 -a $KEY -t applicationid.topic1

## Produce data

    node client.js prod -k $HOST:9092 -a $KEY -t applicationid.topic1

