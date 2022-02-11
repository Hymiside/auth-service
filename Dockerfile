FROM golang:latest

WORKDIR /app

COPY . .
RUN go mod vendor

RUN go build cmd/main.go

EXPOSE 5000

CMD ["./main"]
