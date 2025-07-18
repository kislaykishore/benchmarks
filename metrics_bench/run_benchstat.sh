
ALL_TEST_DIRS := $(shell find . -type f -name 'otel_metrics_test.go' -exec dirname {} \; | sort)

export bench=asyncblocking && go test ${bench} -run '^$' -bench '^Benchmark'  -benchtime .1s -count 6 -cpu 2 -benchmem  | tee ${bench}.txt
export bench=exphistogram && go test ${bench} -run '^$' -bench '^Benchmark'  -benchtime .1s -count 6 -cpu 2 -benchmem  | tee ${bench}.txt
export bench=metricssync && go test ${bench} -run '^$' -bench '^Benchmark'  -benchtime .1s -count 6 -cpu 2 -benchmem  | tee ${bench}.txt
export bench=metricssyncmap && go test ${bench} -run '^$' -bench '^Benchmark'  -benchtime .1s -count 6 -cpu 2 -benchmem  | tee ${bench}.txt
export bench=oldoptimizedimplementation && go test ${bench} -run '^$' -bench '^Benchmark'  -benchtime .1s -count 6 -cpu 2 -benchmem  | tee ${bench}.txt
export bench=oldunoptimizedimplementation && go test ${bench} -run '^$' -bench '^Benchmark'  -benchtime .1s -count 6 -cpu 2 -benchmem  | tee ${bench}.txt
export bench=paramchannel && go test ${bench} -run '^$' -bench '^Benchmark'  -benchtime .1s -count 6 -cpu 2 -benchmem  | tee ${bench}.txt
export bench=reducedbuckets && go test ${bench} -run '^$' -bench '^Benchmark'  -benchtime .1s -count 6 -cpu 2 -benchmem  | tee ${bench}.txt