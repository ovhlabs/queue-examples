
compile:
	mvn compile

package:
	mvn package

produce:
	java \
		-cp target/kafka-example-jar-with-dependencies.jar \
		-Djava.security.auth.login.config=jaas.conf \
		ovh.queue.Main produce \
			--broker ${HOST}:9093 \
			--topic ${TOPIC} \

consume:
	java \
		-cp target/kafka-example-jar-with-dependencies.jar \
		-Djava.security.auth.login.config=jaas.conf \
		ovh.queue.Main consume \
			--broker ${HOST}:9093 \
			--topic ${TOPIC} \
			--consumer-group ${SASL_USERNAME}.java
