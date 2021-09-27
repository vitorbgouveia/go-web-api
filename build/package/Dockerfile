FROM golang

WORKDIR /go/src/web-api

ADD . /go/src/web-api

RUN go mod tidy
RUN gofmt -w ./**/*

EXPOSE 5000

CMD ["sh", "-c", "go run cmd/web-api/main.go"]
