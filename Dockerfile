FROM golang:1.21.9-alpine3.19 as builder

WORKDIR /app

RUN apk add gcc musl-dev tzdata

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o todobot cmd/main.go

FROM alpine:3.19 as runner

RUN apk add tzdata

WORKDIR /app

COPY --from=builder /app/todobot ./
COPY --from=builder /app/configs ./configs/
RUN mkdir database

CMD ["./todobot"]
