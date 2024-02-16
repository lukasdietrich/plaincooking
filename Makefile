TARGET = target
BINARY = $(TARGET)/plaincooking

# Generated Files
GEN                 = $(GEN.WIRE) $(GEN.SQLC)
GEN.WIRE            = cmd/plaincooking/wire_gen.go
SRC.WIRE            = cmd/plaincooking/wire.go
GEN.SQLC            = internal/database/models
GEN.SQLC.QUERIES    = $(GEN.SQLC)/queries.sql_gen.go
SRC.SQLC.MIGRATIONS = $(wildcard internal/database/migrations/*.sql)

# Tools
GO      = go
MIGRATE = $(GO) run github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0
SQLC    = $(GO) run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.25.0
WIRE    = $(GO) run github.com/google/wire/cmd/wire@v0.6.0

.PHONY: all
all: clean build

.PHONY: clean
clean:
	rm -rf $(TARGET) $(GEN)

.PHONY: build
build: $(BINARY)

.PHONY: migrate-new
migrate-new:
	$(MIGRATE) create -ext sql -dir internal/database/migrations -seq $(name)

$(TARGET):
	mkdir -p $(TARGET)

$(BINARY): $(GEN.SQLC.QUERIES) $(GEN.WIRE) | $(TARGET)
	$(GO) build -v -o $(BINARY) ./cmd/plaincooking

$(GEN.WIRE): $(SRC.WIRE)
	$(WIRE) gen ./cmd/plaincooking

$(GEN.SQLC.QUERIES): $(SRC.SQLC.QUERIES) $(SRC.SQLC.MIGRATIONS)
	$(SQLC) generate
