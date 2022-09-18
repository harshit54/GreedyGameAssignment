# Build Stage
# First pull Golang image
FROM golang:1.19.1 AS builder
ENV GOOS linux
ENV CGO_ENABLED _
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./bin/app ./src

FROM debian:latest AS production
COPY --from=builder app .  
ENTRYPOINT ["app"]
EXPOSE 3000