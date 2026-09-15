FROM golang:1.25 AS build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /messenger ./cmd/messenger
EXPOSE 8080 50051
ENTRYPOINT ["/messenger"]