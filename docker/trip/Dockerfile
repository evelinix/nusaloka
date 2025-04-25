# Stage 1: Build the Go binary
FROM golang:1.24.2-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o trip-services ./cmd/trip

# Stage 2: Create the final image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/trip-services .

EXPOSE ${TRIP_SERVICE_PORT}

CMD ["./trip-services"]
