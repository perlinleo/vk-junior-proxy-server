FROM golang:latest as build


ADD . /app
WORKDIR /app
RUN go build ./cmd/proxy/main.go 

RUN pwd
RUN ls build
FROM ubuntu:20.04


WORKDIR /usr/src/app
COPY . .
COPY --from=build /app/main/ .
EXPOSE 8080
CMD ./main
