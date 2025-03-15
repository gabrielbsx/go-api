FROM golang:latest

WORKDIR /go/src/app

COPY . .

EXPOSE 3000

RUN go build -o main cmd/main.go

CMD ["./main"]