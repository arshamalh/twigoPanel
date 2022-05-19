FROM golang:1.18-alpine3.15 AS builder

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:3.15 AS production
COPY --from=builder /app .
CMD ["./main"]