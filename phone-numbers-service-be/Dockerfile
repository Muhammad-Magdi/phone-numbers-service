FROM golang:1.17.2-alpine3.14
RUN apk add build-base

EXPOSE 8080

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . ./

RUN go build -o /phone-numbers-service

CMD [ "/phone-numbers-service" ]