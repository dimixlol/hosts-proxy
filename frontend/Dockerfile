FROM  node:21-alpine as builder

ARG ENVIRONMENT

ENV ENVIRONMENT=${ENVIRONMENT:-local-backend}
WORKDIR /app

COPY ./package.json package.json
COPY ./package-lock.json package-lock.json

RUN npm -v && npm i

COPY ./tsconfig.json tsconfig.json
COPY ./tsconfig.node.json tsconfig.node.json
COPY ./vite.config.ts vite.config.ts
COPY ./public public
COPY ./index.html index.html
COPY ./.env.$ENVIRONMENT .env.$ENVIRONMENT

COPY ./src src

RUN npm run build

FROM nginx

RUN apt update \
    && apt upgrade -y \
    && apt autoremove -y

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80
