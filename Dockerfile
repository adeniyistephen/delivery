# Image
FROM golang:alpine as builder

# Set workdir
WORKDIR /app

# Copy over files
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main main.go

# Minimize using busybox
FROM busybox

WORKDIR /app

COPY --from=builder /app/main /usr/bin/

ENTRYPOINT ["main"]