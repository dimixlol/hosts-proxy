FROM golang:1.22-alpine AS builder

ARG CONFIG_NAME="config.json"
ARG VERSION="0.0.1"

ENV VERSION=${VERSION}

WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app
COPY ./Makefile /app

RUN apk add make && make download

COPY . /app
RUN make build

FROM scratch
COPY --from=builder /app/bin/app /hosts-proxy
COPY --from=builder /app/config/${CONFIG_NAME} /config.json

EXPOSE 8080

ENTRYPOINT ["/hosts-proxy", "--config-file", "/config.json"]
CMD ["--help"]