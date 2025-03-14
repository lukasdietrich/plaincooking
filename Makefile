TARGET = target
BINARY = $(TARGET)/plaincooking

SRC.GO = $(wildcard cmd/**/*.go) $(wildcard internal/**/*.go)

# Generated Files
GEN                 = $(GEN.WIRE) $(GEN.SQLC) $(GEN.API_TYPES) $(GEN.FRONTEND)
GEN.WIRE            = cmd/plaincooking/wire_gen.go
SRC.WIRE            = $(filter-out $(GEN.WIRE),$(SRC.GO))
GEN.SQLC            = internal/database/models
GEN.SQLC.QUERIES    = $(GEN.SQLC)/queries.sql_gen.go
SRC.SQLC.MIGRATIONS = $(wildcard internal/database/migrations/*.sql)
SRC.SQLC.QUERIES    = internal/database/queries.sql
GEN.OPENAPI         = $(TARGET)/openapi.json
GEN.SWAGGER         = $(TARGET)/swagger.json
SRC.SWAGGER         = $(wildcard internal/web/*.go)
GEN.NODE_MODULES    = frontend/node_modules
GEN.API_TYPES       = frontend/src/lib/api/types.gen.ts
GEN.FRONTEND        = frontend/build

# Tools
GO                 = go
MIGRATE            = $(GO) run github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.2
SQLC               = $(GO) run github.com/sqlc-dev/sqlc/cmd/sqlc@v1.27.0
SWAG               = $(GO) run github.com/swaggo/swag/cmd/swag@v1.16.4
WIRE               = $(GO) run github.com/google/wire/cmd/wire@v0.6.0
NPM                = npm --prefix frontend
NPX                = cd frontend; npx
SWAGGER_TO_OPENAPI = $(NPX) swagger2openapi

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
	$(MIGRATE) create \
		-ext sql \
		-dir internal/database/migrations \
		-seq $(name)

$(TARGET):
	mkdir -p $(TARGET)

$(BINARY): $(GEN.FRONTEND) $(GEN.SQLC.QUERIES) $(GEN.WIRE) $(SRC.GO) | $(TARGET)
	$(GO) build \
		-v \
		-o $(BINARY) \
		./cmd/plaincooking

$(GEN.WIRE): $(SRC.WIRE)
	$(WIRE) gen ./cmd/plaincooking

$(GEN.SQLC.QUERIES): $(SRC.SQLC.QUERIES) $(SRC.SQLC.MIGRATIONS)
	$(SQLC) generate

$(GEN.FRONTEND): $(GEN.API_TYPES) $(GEN.NODE_MODULES)
	$(NPM) run build

$(GEN.API_TYPES): $(GEN.OPENAPI) $(GEN.NODE_MODULES)
	$(NPM) run generate:openapi

$(GEN.SWAGGER): $(SRC.SWAGGER) | $(TARGET)
	$(SWAG) init \
		--dir internal/web \
		--parseDependency \
		--generalInfo controller.go \
		--requiredByDefault \
		--outputTypes json \
		--output $(TARGET)

$(GEN.OPENAPI): $(GEN.SWAGGER) $(GEN.NODE_MODULES) | $(TARGET)
	$(SWAGGER_TO_OPENAPI) \
		--outfile $(abspath $(GEN.OPENAPI)) \
		$(abspath $(GEN.SWAGGER))

$(GEN.NODE_MODULES):
	$(NPM) ci
