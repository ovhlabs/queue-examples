
build:
	docker build -t kafkacat .

produce:
	docker run --rm -ti kafkacat \
		-b ${HOST}:9093 \
		-X ssl.ca.location=/etc/ssl/certs \
		-X security.protocol=SASL_SSL \
		-X api.version.request=true \
		-X sasl.mechanisms=PLAIN \
		-X sasl.username=${SASL_USERNAME} \
		-X sasl.password=${SASL_PASSWORD} \
		-t ${TOPIC} -P

consume:
	docker run --rm -ti kafkacat \
		-b ${HOST}:9093 \
		-X ssl.ca.location=/etc/ssl/certs \
		-X security.protocol=SASL_SSL \
		-X api.version.request=true \
		-X sasl.mechanisms=PLAIN \
		-X sasl.username=${SASL_USERNAME} \
		-X sasl.password=${SASL_PASSWORD} \
		-t ${TOPIC} -G ${SASL_USERNAME}.group-kafkacat -C
