FROM alpine:3.3

RUN apk --no-cache add nodejs python-dev build-base

WORKDIR /qaas-client
COPY package.json /qaas-client/package.json
RUN npm install

COPY client.js /qaas-client

ENTRYPOINT ["node", "client.js"]