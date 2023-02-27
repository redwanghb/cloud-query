# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

# Install tools
.PHONY: install-tools
install-tools:
	@echo Installing tools from tools/tool.go
	@cat tools/tool.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml

.PHONY: lint-fix
lint-fix:
	golangci-lint run --fix --config ../../.golangci.yml

.PHONY: gen-docs
gen-docs:
	rm -rf ./docs/tables/*
	go run main.go doc ./docs/tables
	sed 's_(\(.*\))_(https://github.com/cloudquery/cloudquery/blob/main/plugins/source/homebrew/docs/tables/\1)_' docs/tables/README.md > ../../../website/pages/docs/plugins/sources/homebrew/tables.md

# All gen targets
.PHONY: gen
gen: gen-docs