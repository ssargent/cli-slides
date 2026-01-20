# Justfile for CLI Slides project

# Build the binary and place it in the bin directory
build:
    mkdir -p bin
    go build -o bin/slides ./cmd/slides

# Clean build artifacts
clean:
    rm -rf bin/

# Run tests
test:
    go test ./...

# Format code
fmt:
    go fmt ./...

# Run the app with test slides
run-test:
    just build
    ./bin/slides show test_slides.md
