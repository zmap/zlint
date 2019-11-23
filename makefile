SHELL := /bin/bash
# Number of linting Go routines to use in integration tests
PARALLELISM := 5
# Additional integration test flags. Example usage:
#   make integration PARALLELISM=99 INT_FLAGS="-serialSummary -forceDownload"
#   make integration INT_FLAGS="-overwriteExpected -config custom.config.json"
INT_FLAGS :=

CMDS = zlint zlint-gtld-update
CMD_PREFIX = ./cmd/
GO_ENV = GO111MODULE="on" GOFLAGS="-mod=vendor"
BUILD = $(GO_ENV) go build
TEST = $(GO_ENV) GORACE=halt_on_error=1 go test -race
INT_TEST = $(GO_ENV) go test -v -tags integration -timeout 20m ./integration/... -parallelism $(PARALLELISM) $(INT_FLAGS)

all: $(CMDS)

zlint:
	$(BUILD) $(CMD_PREFIX)$(@)

zlint-gtld-update:
	$(BUILD) $(CMD_PREFIX)$(@)

clean:
	rm -f $(CMDS)

test:
	$(TEST) ./...

integration:
	$(INT_TEST)

format-check:
	diff <(find . -name '*.go' -not -path './vendor/*' -print | xargs -n1 gofmt -l) <(printf "")

.PHONY: clean zlint zlint-gtld-update test integration format-check
