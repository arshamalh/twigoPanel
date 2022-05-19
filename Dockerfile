FROM node:18-alpine3.15 AS frontend
WORKDIR /ui
COPY ui .
RUN npm install
RUN npm run build

FROM golang:1.18-alpine3.15 AS builder
WORKDIR /app
COPY server .
COPY --from=frontend /ui/public /ui
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./...

FROM alpine:3.15 AS production
COPY --from=builder /app .
CMD ["./main"]