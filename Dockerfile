FROM  node:19-alpine as builder

WORKDIR /app

COPY ./.env.production .env.production
COPY ./public public
COPY ./index.html index.html
COPY ./src src
COPY ./package.json package.json
COPY ./package-lock.json package-lock.json
COPY ./tsconfig.json tsconfig.json
COPY ./tsconfig.node.json tsconfig.node.json
COPY ./vite.config.ts vite.config.ts

RUN npm -v && npm i && npm run build

FROM nginx

RUN apt update \
    && apt upgrade -y \
    && apt autoremove -y

COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80
