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

package reducedbuckets

import (
	"context"
	"fmt"
	"testing"

	"github.com/kislaykishore/benchmarks/metrics_bench/testutils"
)

func BenchmarkFsOpsCount(b *testing.B) {
	workers := 3
	for _, bufferSize := range testutils.BufferSizes {
		for _, discard := range []bool{true, false} {
			for _, multipleObs := range []bool{true, false} {
				b.Run(fmt.Sprintf("workers=%v/bufferSize=%d/discard=%v/multipleObs=%v", workers, bufferSize, discard, multipleObs), func(b *testing.B) {
					// We use a no-op meter provider to avoid any overhead from metric exporters.
					ctx := context.Background()
					shFn := testutils.SetupOTelMetricExporters(ctx)
					b.Cleanup(func() {
						shFn(ctx)
					})
					// The otelMetrics struct uses a channel and workers for some operations, but
					// FsOpsCount uses atomics directly.
					metrics, err := NewOTelMetrics(ctx, workers, bufferSize)
					if err != nil {
						b.Fatalf("NewOTelMetrics() error = %v", err)
					}

					b.ResetTimer()
					b.RunParallel(func(pb *testing.PB) {
						for i := 0; pb.Next(); i++ {
							var op string
							if multipleObs {
								op = testutils.FsOps[i%len(testutils.FsOps)]
							} else {
								op = "StatFS"
							}
							metrics.FsOpsCount(1, op)
						}
					})
					b.StopTimer()
				})
			}
		}
	}
}

func BenchmarkFsOpsLatency(b *testing.B) {
	workers := 3
	for _, bufferSize := range testutils.BufferSizes {
		for _, discard := range []bool{true, false} {
			for _, multipleObs := range []bool{true, false} {
				b.Run(fmt.Sprintf("workers=%v/bufferSize=%d/discard=%v/multipleObs=%v", workers, bufferSize, discard, multipleObs), func(b *testing.B) {
					// We use a no-op meter provider to avoid any overhead from metric exporters.
					ctx := context.Background()
					shFn := testutils.SetupOTelMetricExporters(ctx)
					b.Cleanup(func() {
						shFn(ctx)
					})
					// The otelMetrics struct uses a channel and workers for some operations, but
					// FsOpsCount uses atomics directly.
					metrics, err := NewOTelMetrics(ctx, workers, bufferSize)
					if err != nil {
						b.Fatalf("NewOTelMetrics() error = %v", err)
					}

					b.ResetTimer()
					b.RunParallel(func(pb *testing.PB) {
						for i := 0; pb.Next(); i++ {
							var op string
							if multipleObs {
								op = testutils.FsOps[i%len(testutils.FsOps)]
							} else {
								op = "StatFS"
							}
							metrics.FsOpsLatency(ctx, 100, op)
						}
					})
					b.StopTimer()
				})
			}
		}
	}
}
