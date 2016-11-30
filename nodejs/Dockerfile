FROM ubuntu:xenial

RUN apt-get update && apt-get install -y \
    build-essential python curl git openssl \
    zlib1g-dev libssl-dev libsasl2-dev

# Install librdkafka
RUN git clone https://github.com/edenhill/librdkafka && \
    cd librdkafka && \
    git checkout tags/v0.9.2 && \
    ./configure && \
    echo '' >> config.h && \
    echo '#define WITH_SSL 1' >> config.h && \
    echo '#define WITH_SASL 1' >> config.h && \
    make && \
    make install
ENV LD_LIBRARY_PATH /usr/local/lib

RUN curl -sL https://deb.nodesource.com/setup_6.x | bash - && \
    apt-get -y install nodejs

WORKDIR /ovh-kafka
COPY package.json /ovh-kafka/package.json
RUN npm install

COPY client.js /ovh-kafka

ENTRYPOINT ["node", "client.js"]]