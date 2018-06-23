FROM golang:1.10.2-alpine3.7

RUN apk add --no-cache git

COPY . /go/src/hack-health-solution/server

WORKDIR /go/src/hack-health-solution/server

RUN go get -d -v .

RUN go build

CMD go get github.com/pilu/fresh && \
    fresh;

EXPOSE 8080