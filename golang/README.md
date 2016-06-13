# Go QaaS example

## Requirements

* Go >= 1.5

## Setup

Build Go binary:

    make build

## Consume data

    ./bin/qaas-client consume -t applicationid.topic1 -a $KEY -h $HOST:9092

## Produce data

    ./bin/qaas-client produce -t applicationid.topic1 -a $KEY -h $HOST:9092
