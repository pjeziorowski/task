# Development
FROM node:20-slim AS dev

WORKDIR /client

COPY package.json /client/package.json
COPY package-lock.json /client/package-lock.json

RUN npm install

ARG SERVER_URL
RUN echo "SERVER_URL=$SERVER_URL" > .env

COPY . .
