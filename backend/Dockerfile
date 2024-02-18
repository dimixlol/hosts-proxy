FROM golang:1.20-alpine AS builder

ARG CONFIG_NAME="config.json"
ARG VERSION="0.0.1"
ARG WORKER

ENV VERSION=${VERSION}

COPY . /app
WORKDIR /app

RUN apk add make && make build

FROM scratch
COPY --from=builder /app/bin/app /kyw-backend
COPY --from=builder /app/config/${CONFIG_NAME} /config.json

EXPOSE 8080

ENTRYPOINT ["/kyw-backend", "--config-file", "/config.json"]
CMD ["--help"]