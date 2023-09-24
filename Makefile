APP="knowyourwebsite"
GO?=go
VERSION?=v0.0.1b

build:
	@$(GO) build -o bin/app -v  -ldflags "-X main.Version=$(VERSION)"  ./cmd/$(APP)


run-persister:
	@echo "Running persister... with $(GO)"
	@$(GO) run cmd/$(APP)/main.go persister listen