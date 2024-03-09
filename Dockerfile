# syntax=docker/dockerfile:1

FROM golang:1.22.1-bullseye as builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build
ENV GO11MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN  go build -o api

FROM gcr.io/distroless/base AS runner

WORKDIR /app
COPY --from=builder /app/api .
COPY --from=builder /app/config/server.yaml ./config/server.yaml

EXPOSE 8080
CMD ["./api"]