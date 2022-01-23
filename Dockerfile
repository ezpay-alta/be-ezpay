FROM golang:1.17.3-alpine3.15 as builder
# Define build env #
ENV GOOS linux
ENV CGO_ENABLED 0

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o app

FROM alpine:3.15 as production
RUN apk add --no-cache ca-certificates
COPY --from=builder app .
CMD ./app
