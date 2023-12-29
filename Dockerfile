FROM golang:1.21

WORKDIR /go/src/app

COPY . .

RUN go build -o main *.go

CMD ["./main"]