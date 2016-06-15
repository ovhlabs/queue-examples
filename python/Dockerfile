FROM alpine:3.3

RUN apk --no-cache add python py-pip

WORKDIR /qaas-client

RUN pip install kafka-python==1.2.1

COPY . /qaas-client

ENTRYPOINT ["python", "client.py"]