# Generate mocks for mock/unit testing
.PHONY: gen-mocks
gen-mocks:
	go install github.com/golang/mock/mockgen@v1.7.0-rc.1
	go run codegen/main.go
	go generate ./client/services/...

# Test unit
.PHONY: test
test:
	go test -race -timeout 3m ./...

.PHONY: gen-docs
gen-docs:
	rm -rf ../../../website/tables/datadog
	go run main.go doc ../../../website/tables/datadog
	sed 's_(\(.*\))_(../../../../../website/tables/datadog/\1)_' ../../../website/tables/datadog/README.md > ./docs/tables/README.md
	sed -i.bak -e 's_(\(.*\).md)_(tables/\1)_' ../../../website/tables/datadog/README.md
	mv ../../../website/tables/datadog/README.md ../../../website/pages/docs/plugins/sources/datadog/tables.md
	sed -i.bak -e 's_(\(.*\).md)_(\1)_' ../../../website/tables/datadog/*.md
	rm -rf ../../../website/tables/datadog/*.bak

.PHONY: lint
lint:
	golangci-lint run --config ../../.golangci.yml 

# All gen targets
.PHONY: gen
gen: gen-mocks gen-docs