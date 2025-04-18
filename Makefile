ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Run app and linters 
BINARY_NAME_HTTP=http_app
MAIN_HTTP=cmd/http/main.go
BUILD_DIR=build
COMPLEXITY_THRESHOLD=10
SERVER_FOLDER_HTTP=internal/app/http_app/server

BINARY_NAME_GRPC=grpc_app
MAIN_GRPC=cmd/grpc/main.go

build: lint 
	go build -o $(BUILD_DIR)/$(BINARY_NAME_HTTP) $(MAIN_HTTP)

run: build
	$(BUILD_DIR)/$(BINARY_NAME_HTTP)

clean:
	rm -rf $(BUILD_DIR)
	go clean

build-grpc:
	go build -o $(BUILD_DIR)/$(BINARY_NAME_GRPC) $(MAIN_GRPC)

run-grpc: build-grpc
	$(BUILD_DIR)/$(BINARY_NAME_GRPC)

deps:
	go get -u ./...
	go mod tidy

install-linters:
	@which gocyclo >/dev/null 2>&1 || (echo "Installing gocyclo..." && go install github.com/fzipp/gocyclo/cmd/gocyclo@latest)
	@which gocognit >/dev/null 2>&1 || (echo "Installing gocognit..." && go install github.com/uudashr/gocognit/cmd/gocognit@latest)

lint: install-linters lint-gocyclo lint-gocognit

lint-gocyclo:
	@echo "Run gocyclo (max complexity: $(COMPLEXITY_THRESHOLD))..."
	@result=$$(find . -name '*.go' ! -path './pkg/*' | xargs gocyclo -over $(COMPLEXITY_THRESHOLD)); \
	if [ -n "$$result" ]; then \
		echo "gocyclo check failed:"; \
		echo "$$result"; \
		exit 1; \
	else \
		echo "success"; \
	fi

lint-gocognit:
	@echo "Run gocognit (max complexity: $(COMPLEXITY_THRESHOLD))..."
	@result=$$(find . -name '*.go' ! -path './pkg/*' | xargs gocognit -over $(COMPLEXITY_THRESHOLD)); \
	if [ -n "$$result" ]; then \
		echo "gocognit check failed:"; \
		echo "$$result"; \
		exit 1; \
	else \
		echo "success"; \
	fi

test-unit:
	$(info Running tests...)
	go test ./$(SERVER_FOLDER_HTTP) -coverprofile $(SERVER_FOLDER_HTTP)/coverage.out -v
	go tool cover -html=$(SERVER_FOLDER_HTTP)/coverage.out

test-integration:
	$(info Running tests...)
	go test ./tests/integration/... -v -tags integration

prod-dc-up:	
	docker compose up -d

test-dc-up:
	docker compose -f ./tests/integration/docker-compose.yaml up -d 


# Run migrations
POSTGRES_SETUP_TEST=user=test password=test dbname=test host=localhost port=8001 sslmode=disable
POSTGRES_SETUP_PROD=user=${POSTGRES_USER} password=${POSTGRES_PASSWORD} dbname=library host=localhost port=5432 sslmode=disable


INTERNAL_PKG_PATH=$(CURDIR)/internal
MIGRATION_FOLDER=./migrations

.PHONY: migration-create
migration-create:
	goose -dir "$(MIGRATION_FOLDER)" -s create "$(name)" sql

.PHONY: test-migration-up
test-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" up

.PHONY: test-migration-down
test-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_TEST)" down

.PHONY: prod-migration-up
prod-migration-up:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_PROD)" up

.PHONY: prod-migration-down
prod-migration-down:
	goose -dir "$(MIGRATION_FOLDER)" postgres "$(POSTGRES_SETUP_PROD)" down

.PHONY: generate
generate:
	rm -rf pkg/pb
	mkdir -p pkg/pb
	protoc --proto_path=api/ --go_out=pkg/pb --go-grpc_out=pkg/pb api/api.proto

.PHONY: all build run build-grpc run-grpc clean lint lint-gocyclo lint-gocognit deps prod-dc-up test-dc-up migrations test-unit test-integration
.DEFAULT_GOAL:=run
