GO = go
TEST_DIRS := $(shell find . -type f -name 'otel_metrics_test.go' -exec dirname {} \; | sort)

.PHONY: bench
bench: $(TEST_DIRS:%=bench/%)
bench/%: DIR=$*
bench/%:
	@echo "$(GO) test $(DIR) -run=xxxxxMatchNothingxxxxx -bench=. -benchtime .1s -count 6 -cpu 2 -benchmem  | tee $(DIR).txt" \
		&& $(GO) test $(DIR) -run=xxxxxMatchNothingxxxxx -bench=. -benchtime .1s -count 6 -cpu 2 -benchmem  | tee $(DIR).txt

.PHONY: benchstat
benchstat:
	benchstat -table .name oldunoptimizedimplementation.txt metricssync.txt oldoptimizedimplementation.txt paramchannel.txt reducedbuckets.txt asyncblocking.txt exphistogram.txt metricssyncmap.txt > results.txt
