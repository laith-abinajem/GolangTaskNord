#################### RUN ####################
	
# Run the API server
.PHONY: run
run:
	@ $(MAKE) build-quick
	@ $(shell source exports.sh)
	@ ./bin/server

.PHONY: seed
seed:
	@ go run cmd/seed/main.go


#################### Build Executable ####################
# Build amd64	for alpine
.PHONY: build
build:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w -s' -o ./bin/server cmd/api/*.go

# Build depending on the OS
.PHONY: build-quick
build-quick:
	@go build  -o ./bin/server cmd/api/*.go


#################### Docker Compose ####################

# Stack
.PHONY: up
up:
	@docker compose up -d --build --force-recreate --remove-orphans

.PHONY: down
down:
	@docker compose down 

.PHONY: top
top:
	@docker stats
