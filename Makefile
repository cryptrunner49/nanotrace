# Makefile for NanoTrace

# Variables
BINARY_NAME := nanotrace
BIN_DIR := bin
SRC_DIR := ./cmd

# Create bin directory if not exists
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

# Build the project
build: $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)/main.go
	@echo "✅ Build complete! Binary located at $(BIN_DIR)/$(BINARY_NAME)"

# Run the project
run: build
	@echo "🚀 Running NanoTrace..."
	./$(BIN_DIR)/$(BINARY_NAME)

# Run tests
.PHONY: test

test:
	@echo "🧪 Running tests..."
	go test ./...
	@echo "✅ Tests completed!"

# Clean up build artifacts
.PHONY: clean

clean:
	@echo "🧹 Cleaning up..."
	rm -rf $(BIN_DIR)
	@echo "✅ Cleanup complete!"
