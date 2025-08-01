// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maintest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/kislaykishore/benchmarks/metrics_bench/asyncblocking"
	"github.com/kislaykishore/benchmarks/metrics_bench/exphistogram"
	"github.com/kislaykishore/benchmarks/metrics_bench/metricssync"
	"github.com/kislaykishore/benchmarks/metrics_bench/metricssyncmap"
	"github.com/kislaykishore/benchmarks/metrics_bench/oldoptimizedimplementation"
	"github.com/kislaykishore/benchmarks/metrics_bench/oldunoptimizedimplementation"
	"github.com/kislaykishore/benchmarks/metrics_bench/paramchannel"
	"github.com/kislaykishore/benchmarks/metrics_bench/reducedbuckets"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/exemplar"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const (
	bufferSize = 8 * 10240000
)

var fsOps = []string{"StatFS", "LookUpInode", "GetInodeAttributes", "Open", "Read", "Write", "Close"}
var exemplarFilterOff = os.Getenv("EXEMPLAR_FILTER_OFF")

func BenchmarkFsOpsCountAsync(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(1, "StatFS")
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsCountAsyncFlush(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(1, "StatFS")
		}
	})
	metrics.Flush()
}

func BenchmarkFsOpsCountAsyncDiscardMetrics(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, 1, func() {})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(1, "StatFS")
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsCountAsyncMultipleOps(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsCount(1, op)
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsCountAsyncMultipleOpsDiscardMetrics(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, 1, func() {})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsCount(1, op)
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsCountMapSync(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := metricssyncmap.NewOTelMetrics(ctx, 3, 1024000)
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(1, "StatFS")
		}
	})
}

func BenchmarkFsOpsCountMapSyncMultipleOps(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := metricssyncmap.NewOTelMetrics(ctx, 3, 1024000)
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsCount(1, op)
		}
	})
}

func BenchmarkFsOpsCountParamsChAsync(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(1, "StatFS")
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsCountParamsChAsyncDiscardMetrics(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, 1, func() {})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(1, "StatFS")
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsCountParamsChAsyncFlush(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(1, "StatFS")
		}
	})
	metrics.Flush()
}

func BenchmarkFsOpsCountAsyncParamsChMultipleOps(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsCount(1, op)
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsCountAsyncParamsChMultipleOpsFlush(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsCount(1, op)
		}
	})
	metrics.Flush()
}

func BenchmarkFsOpsCountAsyncMultipleOpsFlush(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsCount(1, op)
		}
	})
	metrics.Flush()
}

func BenchmarkFsOpsCountSync(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := metricssync.NewOTelMetrics(ctx, 3, 1)
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(1, "StatFS")
		}
	})

}

func BenchmarkFsOpsCountSyncMultipleOps(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := metricssync.NewOTelMetrics(ctx, 3, 1024000)
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsCount(1, op)
		}
	})
}

func BenchmarkFsOpsCountSyncMultipleOpsOldOptimizedImpl(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := oldoptimizedimplementation.NewOTelMetrics()
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsCount(ctx, 1, op)
		}
	})
}

func BenchmarkFsOpsCountSyncMultipleOpsOldUnoptimizedImpl(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := oldunoptimizedimplementation.NewOTelMetrics()
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsCount(ctx, 1, op)
		}
	})
}

func BenchmarkFsOpsCountSyncOldOptimizedImpl(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := oldoptimizedimplementation.NewOTelMetrics()
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(ctx, 1, "StatFS")
		}
	})
}

func BenchmarkFsOpsCountSyncOldUnoptimizedImpl(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := oldunoptimizedimplementation.NewOTelMetrics()
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsCount(ctx, 1, "StatFS")
		}
	})
}

func BenchmarkFsOpsLatencySync(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := metricssync.NewOTelMetrics(ctx, 3, 1)
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
}

func BenchmarkFsOpsLatencySyncExponentialHistogram(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExportersWithExpHistogram(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := exphistogram.NewOTelMetrics(ctx, 3, 1)
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
}

func BenchmarkFsOpsLatencyAsync(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsLatencyAsyncMultipleOps(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsLatencyAsyncMultipleOpsDiscardMetrics(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, 1, func() {})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsLatencySyncMultipleOps(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := metricssync.NewOTelMetrics(ctx, 3, 1024000)
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
}

func BenchmarkFsOpsLatencyAsyncDiscardMetrics(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, 1, func() {})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsLatencyAsyncFlush(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
	metrics.Flush()
}

func BenchmarkFsOpsLatencyAsyncMultipleOpsFlush(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := asyncblocking.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
	metrics.Flush()
}

func BenchmarkFsOpsLatencyParamsChAsync(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsLatencyParamsChAsyncDiscardMetrics(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, 1, func() {})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsLatencyParamsChAsyncFlush(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
	metrics.Flush()
}

func BenchmarkFsOpsLatencyParamsChAsyncMultipleOps(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsLatencyParamsChAsyncMultipleOpsFlush(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := paramchannel.NewOTelMetrics(ctx, 3, bufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
	metrics.Flush()
}

func BenchmarkFsOpsLatencyReducedBucketsSync(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := metricssync.NewOTelMetrics(ctx, 3, bufferSize)
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
}

func BenchmarkFsOpsLatencyReducedBucketsSyncMultipleOps(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := reducedbuckets.NewOTelMetrics(ctx, 3, 1024000)
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
}

func BenchmarkFsOpsLatencySyncMultipleOpsOldOptimizedImpl(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := oldoptimizedimplementation.NewOTelMetrics()
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
}

func BenchmarkFsOpsLatencySyncOldOptimizedImpl(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := oldoptimizedimplementation.NewOTelMetrics()
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
}

func BenchmarkFsOpsLatencySyncMultipleOpsOldUnoptimizedImpl(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := oldunoptimizedimplementation.NewOTelMetrics()
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := fsOps[i%len(fsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
}

func BenchmarkFsOpsLatencySyncOldUnoptimizedImpl(b *testing.B) {
	ctx := context.Background()
	shFn := setupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := oldunoptimizedimplementation.NewOTelMetrics()
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			metrics.FsOpsLatency(ctx, 100, "StatFS")
		}
	})
}

func setupOTelMetricExporters(ctx context.Context) (shutdownFn ShutdownFn) {
	shutdownFns := make([]ShutdownFn, 0)
	options := make([]metric.Option, 0)

	opts, shutdownFn := setupPrometheus(8080)
	options = append(options, opts...)
	shutdownFns = append(shutdownFns, shutdownFn)

	res, err := getResource(ctx)
	if err == nil {
		options = append(options, metric.WithResource(res))
	}
	if exemplarFilterOff != "" {
		options = append(options, metric.WithExemplarFilter(exemplar.AlwaysOffFilter))
	}

	meterProvider := metric.NewMeterProvider(options...)
	shutdownFns = append(shutdownFns, meterProvider.Shutdown)

	otel.SetMeterProvider(meterProvider)

	return JoinShutdownFunc(shutdownFns...)
}

// setExponentialAggregation is an OTel View that drops the metrics that don't match the allowed prefixes.
func setExponentialAggregation(i metric.Instrument) (metric.Stream, bool) {
	s := metric.Stream{Name: i.Name, Description: i.Description, Unit: i.Unit}
	i.Name = "fs/ops_latency"
	s.Aggregation = metric.AggregationBase2ExponentialHistogram{MaxSize: 34, MaxScale: 0}
	return s, true
}

func setupOTelMetricExportersWithExpHistogram(ctx context.Context) (shutdownFn ShutdownFn) {
	shutdownFns := make([]ShutdownFn, 0)
	options := make([]metric.Option, 0)

	opts, shutdownFn := setupPrometheus(8080)
	options = append(options, opts...)
	shutdownFns = append(shutdownFns, shutdownFn)

	res, err := getResource(ctx)
	if err == nil {
		options = append(options, metric.WithResource(res))
	}
	options = append(options, metric.WithView(setExponentialAggregation))
	if exemplarFilterOff != "" {
		options = append(options, metric.WithExemplarFilter(exemplar.AlwaysOffFilter))
	}

	meterProvider := metric.NewMeterProvider(options...)

	shutdownFns = append(shutdownFns, meterProvider.Shutdown)

	otel.SetMeterProvider(meterProvider)

	return JoinShutdownFunc(shutdownFns...)
}

func setupPrometheus(port int64) ([]metric.Option, ShutdownFn) {
	if port <= 0 {
		return nil, nil
	}
	exporter, err := prometheus.New(prometheus.WithoutUnits(), prometheus.WithoutCounterSuffixes(), prometheus.WithoutScopeInfo(), prometheus.WithoutTargetInfo())
	if err != nil {
		return nil, nil
	}
	shutdownCh := make(chan context.Context)
	done := make(chan interface{})
	go serveMetrics(port, shutdownCh, done)
	return []metric.Option{metric.WithReader(exporter)}, func(ctx context.Context) error {
		shutdownCh <- ctx
		close(shutdownCh)
		<-done
		close(done)
		return nil
	}
}

func serveMetrics(port int64, shutdownCh <-chan context.Context, done chan<- interface{}) {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	prometheusServer := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := prometheusServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Failed to start Prometheus server: %v\n", err)
		}
	}()
	go func() {
		ctx := <-shutdownCh
		defer func() { done <- true }()
		if err := prometheusServer.Shutdown(ctx); err != nil {
			return
		}
	}()
}

func getResource(ctx context.Context) (*resource.Resource, error) {
	return resource.New(ctx,
		// Use the GCP resource detector to detect information about the GCP platform
		resource.WithDetectors(gcp.NewDetector()),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceName("gcsfuse"),
			semconv.ServiceVersion("2.1"),
		),
	)
}

type ShutdownFn func(ctx context.Context) error

// JoinShutdownFunc combines the provided shutdown functions into a single function.
func JoinShutdownFunc(shutdownFns ...ShutdownFn) ShutdownFn {
	return func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFns {
			if fn == nil {
				continue
			}
			err = errors.Join(err, fn(ctx))
		}
		return err
	}
}
