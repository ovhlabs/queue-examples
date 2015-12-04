sbt "run -z 127.0.0.1:2181 -k 192.168.99.100:9092 -a tokenid:tokenkey -t topic1 -m prod"
sbt "run -z 127.0.0.1:2181 -k 192.168.99.100:9092 -a tokenid:tokenkey -t topic1 -m cons"

