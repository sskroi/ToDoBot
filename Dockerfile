FROM golang:1.21

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN GOOS=linux go build -o todobot cmd/main.go

CMD ["./todobot"]