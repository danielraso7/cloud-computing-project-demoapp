FROM golang:latest

COPY . /go/src/app

WORKDIR /go/src/app

RUN go mod init main.go

RUN go test -v ./...

RUN go build -o main .

CMD ["./main"]
