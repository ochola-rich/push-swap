# Makefile

# Variables
APP_PUSH_SWAP=push-swap
APP_CHECKER=checker
BUILD_DIR=build
CMD_PUSH_SWAP=cmd/$(APP_PUSH_SWAP)/main.go
CMD_CHECKER=cmd/$(APP_CHECKER)/main.go

# .PHONY explicitly declares these targets as abstract commands, not files.
# This prevents conflicts with the 'build' directory.
.PHONY: all build run clean

# Default target (runs when you just type 'make')
all: build

# 1. Build the binary locally into the build/ folder
build:
	@echo "Building $(APP_PUSH_SWAP)..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_PUSH_SWAP) $(CMD_PUSH_SWAP)
	@echo "Building $(APP_CHECKER)..."
	go build -o $(BUILD_DIR)/$(APP_CHECKER) $(CMD_CHECKER)
	@echo "Build complete! Try running: ./$(BUILD_DIR)/$(APP_CHECKER) '3 1 0 6 2 9' or ./$(BUILD_DIR)/$(APP_CHECKER) '3 1 0 6 2 9'"

# 2. Run the local binary
run: build
	@echo "Starting application..."
	./$(BUILD_DIR)/$(APP_PUSH_SWAP) '3 1 0 6 2 9' | ./$(BUILD_DIR)/$(APP_CHECKER) '3 1 0 6 2 9'

# 3. Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BUILD_DIR)
	@echo "Clean complete."

