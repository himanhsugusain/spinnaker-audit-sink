# syntax=docker/dockerfile:1
FROM  --platform=$BUILDPLATFORM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH\
      go build -a -ldflags="-s -w" -o app

FROM scratch AS release
COPY --from=builder /app/app /app
ENTRYPOINT ["/app"]
