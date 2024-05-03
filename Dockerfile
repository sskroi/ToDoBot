FROM golang:1.21.9-alpine3.19 as builder

WORKDIR /app

RUN apk add gcc musl-dev

COPY go.mod go.sum .

RUN go mod download

COPY . .

RUN go build -o todobot cmd/main.go

FROM alpine:3.19 as runner

WORKDIR /app

COPY --from=builder /app/todobot .

CMD ["./todobot"]
