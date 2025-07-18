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

package paramchannel

import (
	"context"
	"testing"

	"github.com/kislaykishore/benchmarks/metrics_bench/testutils"
)

func BenchmarkFsOpsCountAsync(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, testutils.BufferSize, func() {
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

func BenchmarkFsOpsCountAsyncDiscardMetrics(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, 1, func() {})
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
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, testutils.BufferSize, func() {
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

func BenchmarkFsOpsCountAsyncMultipleOps(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, testutils.BufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := testutils.FsOps[i%len(testutils.FsOps)]
			metrics.FsOpsCount(1, op)
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsCountAsyncMultipleOpsFlush(b *testing.B) {
	// We use a no-op meter provider to avoid any overhead from metric exporters.
	ctx := context.Background()
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, testutils.BufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := testutils.FsOps[i%len(testutils.FsOps)]
			metrics.FsOpsCount(1, op)
		}
	})
	metrics.Flush()
}

func BenchmarkFsOpsLatencyAsync(b *testing.B) {
	ctx := context.Background()
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, testutils.BufferSize, func() {
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

func BenchmarkFsOpsLatencyAsyncDiscardMetrics(b *testing.B) {
	ctx := context.Background()
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, 1, func() {})
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
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, testutils.BufferSize, func() {
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

func BenchmarkFsOpsLatencyAsyncMultipleOps(b *testing.B) {
	ctx := context.Background()
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, testutils.BufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := testutils.FsOps[i%len(testutils.FsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
	b.StopTimer()
	metrics.Flush()
}

func BenchmarkFsOpsLatencyAsyncMultipleOpsFlush(b *testing.B) {
	ctx := context.Background()
	shFn := testutils.SetupOTelMetricExporters(ctx)
	b.Cleanup(func() {
		shFn(ctx)
	})
	// The otelMetrics struct uses a channel and workers for some operations, but
	// FsOpsCount uses atomics directly.
	metrics, err := NewOTelMetrics(ctx, 3, testutils.BufferSize, func() {
		b.FailNow()
	})
	if err != nil {
		b.Fatalf("NewOTelMetrics() error = %v", err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for i := 0; pb.Next(); i++ {
			op := testutils.FsOps[i%len(testutils.FsOps)]
			metrics.FsOpsLatency(ctx, 100, op)
		}
	})
	metrics.Flush()
}
