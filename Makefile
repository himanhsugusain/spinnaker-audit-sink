BINARY_NAME=./bin/audit-sink
DOCKER_IMAGE_NAME=spinnaker-audit-sink
DOCKER_TAG=dev

.PHONY: fmt vet build run clean docker-build

all: fmt build

dbuild: docker-build

fmt:
	@echo "Formatting Go code..."
	@go fmt ./...

vet:
	@echo "Running go vet..."
	@go vet ./...

build: vet fmt
	@echo "Building the application..."
	@go build -o $(BINARY_NAME) main.go

run: run
	@echo "Running the application..."
	@./$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@go clean
	@rm -f $(BINARY_NAME)

docker-build:
	@echo "Building Docker image..."
	@docker build -t $(DOCKER_IMAGE_NAME):$(DOCKER_TAG) .
