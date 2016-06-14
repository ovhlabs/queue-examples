FROM alpine:3.3

RUN apk --no-cache add python py-pip
#&& \
    #pip install virtualenv

WORKDIR /qaas-client
#COPY requirements.txt /qaas-client/requirements.txt
#RUN virtualenv /env && /env/bin/pip install -r /qaas-client/requirements.txt
RUN pip install kafka-python==1.2.1

COPY . /qaas-client

ENTRYPOINT ["python", "client.py"]