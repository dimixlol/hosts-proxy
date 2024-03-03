APP="hosts-proxy"
GO?=$(shell which go)
VERSION?=v0.0.1b

download:
	@$(GO) mod download

build:
	@$(GO) build -o bin/app -v  -ldflags "-X main.Version=$(VERSION)"  ./cmd/$(APP)

run-persister:
	@echo "Running persister... with $(GO version)"
	@$(GO) run cmd/$(APP)/main.go persister listen

run-proxier:
	@echo "Running proxier... with $(GO version)"
	@$(GO) run cmd/$(APP)/main.go proxier listen