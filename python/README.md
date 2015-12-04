# Python QaaS example

## Requirements

* Python version 2.7

## Setup

    make install

## Produce

~~~
 python client.py --key "mysecret" --host 192.168.99.100:9092 --mode prod --group "test" --topic "applicationid.test"
~~~

## Consume

~~~
 python client.py --key "mysecret" --host 192.168.99.100:9092 --mode cons --group "test" --topic "applicationid.test"
~~~
