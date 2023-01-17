FROM golang:latest

COPY . /go/src/app

WORKDIR /go/src/app

RUN go mod init http_request

RUN go test -v ./...

RUN go build -o http_request .

CMD ["./http_request"]
