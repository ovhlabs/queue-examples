
compile:
	mvn compile

package:
	mvn package

produce:
	java \
		-cp target/kafka-example-jar-with-dependencies.jar \
		-Djava.security.auth.login.config=jaas.conf \
		ovh.queue.Main produce \
		--broker ${K_BROKER} \
		--topic ${K_TOPIC}

consume:
	java \
		-cp target/kafka-example-jar-with-dependencies.jar \
		-Djava.security.auth.login.config=jaas.conf \
		ovh.queue.Main consume \
		--broker ${K_BROKER} \
		--topic ${K_TOPIC} \
		--group ${K_USERNAME}.java
