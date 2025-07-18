GO = go
TEST_DIRS := $(shell find . -type f -name 'otel_metrics_test.go' -exec dirname {} \; | sort)

.PHONY: bench
bench: $(TEST_DIRS:%=bench/%)
bench/%: DIR=$*
bench/%:
	@echo "$(GO) test $(DIR) -run=xxxxxMatchNothingxxxxx -bench=. -benchtime .1s -count 6 -cpu 2 -benchmem  | tee $(DIR).txt" \
		&& $(GO) test $(DIR) -run=xxxxxMatchNothingxxxxx -bench=. -benchtime .1s -count 6 -cpu 2 -benchmem  | tee $(DIR).txt
