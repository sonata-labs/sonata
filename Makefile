# Makefile for Sonata

tidy:
	go mod tidy

# Generate code from protobuf files
.PHONY: gen clean
gen: clean
	buf generate
	make tidy

# Clean generated code
clean:
	rm -rf gen