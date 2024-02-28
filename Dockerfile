# ------------------------------------------------ Builder
FROM golang:alpine AS builder

RUN apk add git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o /entrypoint

# -------------------------------------------------- Runtime
FROM apline:latest AS runtime

LABEL maintainer="alireza feizi <alirezafeyze44@gmail.com>"

WORKDIR /app

COPY --from=builder /entrypoint .

ENTRYPOINT ["./entrypoint"]
