ifeq (/,${HOME})
    GOLANGCI_LINT_CACHE=/tmp/golangci-lint-cache/
else
    GOLANGCI_LINT_CACHE=${HOME}/.cache/golangci-lint
endif
GOLANGCI_LINT ?= GOLANGCI_LINT_CACHE=$(GOLANGCI_LINT_CACHE) go run github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: fmt
fmt:
	$(GOLANGCI_LINT) run --fix

.PHONY: lint
lint:
	$(GOLANGCI_LINT) run

.PHONY: test
test:
	go test -v ./...

test-coverage:
	@go test -cover -outputdir=./ -coverprofile=all.coverprofile ./...

coverage-html: test-coverage
	@go tool cover -html=all.coverprofile

.PHONY: check
check: lint test generate_check
