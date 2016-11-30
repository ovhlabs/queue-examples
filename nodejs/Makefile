
node-install-deps:
	npm install

produce:
	node client.js produce \
		--broker ${HOST}:9093 \
		--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
		--topic ${TOPIC}

consume:
	node client.js consume \
		--broker ${HOST}:9093 \
		--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
		--topic ${TOPIC} --consumer-group ${SASL_USERNAME}.group-node

##

build-docker:
	docker build -t ovhqueue/node-client-example .

consume-in-docker:
	docker run --rm -ti ovhqueue/node-client-example \
		consume --broker ${HOST}:9093 \
			--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
			--topic ${TOPIC} --consumer-group ${SASL_USERNAME}.group-dnode