FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64\
      go build -a -ldflags="-s -w" -o app

FROM scratch AS release
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
