# Default target
default: build

# Build the application
build:
	@echo "Building the application..."
	go build -v -o ./bin/loanengine ./.

# Run the application locally
run:
	@echo "Running the application..."
	./bin/loanengine

# Build and then run the application
restart: build run

# Build for production (Linux AMD64)
prod:
	@echo "Building for production (GOOS=linux, GOARCH=amd64)..."
	GOOS=linux GOARCH=amd64 go build -v -o ./bin/loanengine ./main.go

# Push the built binary to the production server
push: prod
	@echo "Pushing binary to production server..."
	scp bin/loanengine amamrthaloan:/tmp

# Release process: build, push, and deploy
release: push
	@echo "Releasing application..."
	ssh amamrthaloan 'bash /tmp/build_loanengine.sh'

# Clean up build artifacts
clean:
	@echo "Cleaning up build artifacts..."
	rm -f ./bin/loanengine

# Help message
help:
	@echo "Makefile commands:"
	@echo "  default    - Build the application"
	@echo "  build      - Build the application"
	@echo "  run        - Run the application"
	@echo "  restart    - Build and run the application"
	@echo "  prod       - Build the application for production (Linux AMD64)"
	@echo "  push       - Push the binary to the production server"
	@echo "  release    - Build, push, and deploy the application"
	@echo "  clean      - Remove build artifacts"
	@echo "  help       - Display this help message"
