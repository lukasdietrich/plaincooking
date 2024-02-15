TARGET = target
BINARY = $(TARGET)/plaincooking

# Generated Files
GEN      = $(GEN.WIRE)
GEN.WIRE = cmd/plaincooking/wire_gen.go
SRC.WIRE = cmd/plaincooking/wire.go

# Tools
GO   = go
SQLC = $(GO) run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.25.0
WIRE = $(GO) run github.com/google/wire/cmd/wire@v0.6.0

.PHONY: all
all: clean build

.PHONY: clean
clean:
	rm -rf $(TARGET) $(GEN)

.PHONY: build
build: $(BINARY)

$(TARGET):
	mkdir -p $(TARGET)

$(BINARY): $(GEN.WIRE) | $(TARGET)
	$(GO) build -v -o $(BINARY) ./cmd/plaincooking

$(GEN.WIRE): $(SRC.WIRE)
	$(WIRE) gen ./cmd/plaincooking
