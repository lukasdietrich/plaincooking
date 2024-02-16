TARGET = target
BINARY = $(TARGET)/plaincooking

SRC.GO = $(wildcard cmd/**/*.go) $(wildcard internal/**/*.go)

# Generated Files
GEN                 = $(GEN.WIRE) $(GEN.SQLC)
GEN.WIRE            = cmd/plaincooking/wire_gen.go
SRC.WIRE            = cmd/plaincooking/wire.go
GEN.SQLC            = internal/database/models
GEN.SQLC.QUERIES    = $(GEN.SQLC)/queries.sql_gen.go
SRC.SQLC.MIGRATIONS = $(wildcard internal/database/migrations/*.sql)
SRC.SQLC.QUERIES    = internal/database/queries.sql
GEN.SWAGGER         = $(TARGET)/swagger.json
SRC.SWAGGER         = $(wildcard internal/web/*.go)

# Tools
GO      = go
MIGRATE = $(GO) run github.com/golang-migrate/migrate/v4/cmd/migrate@v4.17.0
SQLC    = $(GO) run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.25.0
SWAG    = $(GO) run github.com/swaggo/swag/cmd/swag@v1.16.3
WIRE    = $(GO) run github.com/google/wire/cmd/wire@v0.6.0

.PHONY: all
all: clean build

.PHONY: clean
clean:
	rm -rf $(TARGET) $(GEN)

.PHONY: build
build: $(BINARY) $(GEN.SWAGGER)

.PHONY: go-vet
go-vet:
	$(GO) vet ./...

.PHONY: migrate-new
migrate-new:
	$(MIGRATE) create -ext sql -dir internal/database/migrations -seq $(name)

$(TARGET):
	mkdir -p $(TARGET)

$(BINARY): $(GEN.SQLC.QUERIES) $(GEN.WIRE) $(SRC.GO) | $(TARGET)
	$(GO) build -v -o $(BINARY) ./cmd/plaincooking

$(GEN.WIRE): $(SRC.WIRE)
	$(WIRE) gen ./cmd/plaincooking

$(GEN.SQLC.QUERIES): $(SRC.SQLC.QUERIES) $(SRC.SQLC.MIGRATIONS)
	$(SQLC) generate

$(GEN.SWAGGER): $(SRC.SWAGGER) | $(TARGET)
	$(SWAG) init --dir internal/web --generalInfo api.go --outputTypes json --output $(TARGET)
