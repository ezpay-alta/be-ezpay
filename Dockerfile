#STAGE 1
FROM golang:1.17.3-alpine3.15 AS builder
WORKDIR /app
COPY ./ ./
RUN go mod download
RUN go build -o main
RUN go clean --modcache


# STAGE 2
FROM alpine:3.15
WORKDIR /app
COPY --from=builder /app/.env .
COPY --from=builder /app/main  .
EXPOSE 8000
CMD [ "./main" ]
