## QaaS client examples

* [Go](go)
* [NodeJs](nodejs)
* [Python](python)
* [Scala](scala_kafka_0.8.2)

# Basic concepts

Every client uses three constants :

- Application id : The application id given to use QaaS.
- Key   : The key used for the authentication. This key is linked to your application.
- Topic : The topic to which data will be pushed.

Every example uses the key as a client id to authenticate and pushed on a topic prefixed by the application id.
