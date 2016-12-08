
python-install-deps:
	pip install --upgrade kafka-python==1.3.1

produce:
	python client.py produce \
		--broker ${HOST}:9093 \
		--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
		--topic ${TOPIC}

consume:
	python client.py consume \
		--broker ${HOST}:9093 \
		--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
		--topic ${TOPIC} --consumer-group ${SASL_USERNAME}.group-python

##

build-docker:
	docker build -t ovhqueue/py-client-example .

consume-in-docker:
	docker run --rm -ti ovhqueue/py-client-example \
		consume --broker ${HOST}:9093 \
			--username ${SASL_USERNAME} --password ${SASL_PASSWORD} \
			--topic ${TOPIC} --consumer-group ${SASL_USERNAME}.group-dpy
