# syntax=docker/dockerfile:1

FROM node:lts-alpine as frontend_builder
WORKDIR /app
COPY frontend/ .
RUN npm install
RUN npm run build

FROM golang:1.16-alpine as backend_builder
COPY . .
COPY --from=frontend_builder /app/build ./frontend/
RUN go mod download
RUN go build -o /app

FROM alpine:latest
COPY --from=backend_builder /app /app
EXPOSE 8080
CMD /app -addr 0.0.0.0 -port 8080
