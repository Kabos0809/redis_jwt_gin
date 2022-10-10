FROM golang:1.19.1-alpine3.16

RUN mkdir -p /go/app && cd /go/app
COPY ./ /go/app
WORKDIR /go/app

RUN apk update && apk add git
RUN go install github.com/cosmtrek/air@v1.29.0 && \
  go install github.com/swaggo/swag/cmd/swag@v1.8.0

RUN go mod download

ENV GO111MODULE=on

CMD ["air", "-c", ".air.toml"]