# syntax=docker/dockerfile:1

FROM node:lts-alpine as frontend_builder
WORKDIR /app
COPY frontend/ .
RUN npm install
RUN npm run build

FROM golang:1-bullseye as backend_builder
WORKDIR /app
COPY . .
COPY --from=frontend_builder /app/build ./frontend/build/
RUN go mod download
RUN go build -o ./main

FROM ubuntu:focal
VOLUME /db
EXPOSE 8080
WORKDIR /db
COPY --from=backend_builder /app/main /main
CMD /main -addr 0.0.0.0 -port 8080
