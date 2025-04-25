# Stage 1: Build the Go binary
FROM golang:1.24.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o account-services ./cmd/account

# Stage 2: Create the final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/account-services .

EXPOSE ${ACCOUNT_SERVICE_PORT}

CMD ["./account-services"]
