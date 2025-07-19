// Copyright 2024 Google LLC
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

package oldoptimizedimplementation

import (
	"context"
	"errors"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/kislaykishore/benchmarks/metrics_bench/testutils"
)

var (
	fsOpsMeter = otel.Meter("fs_op")

	// fsOpKey specifies the FS operation like LookupInode, ReadFile etc.
	fsOpKey = attribute.Key("fs_op")

	fsOpsOptionCache sync.Map
)

func loadOrStoreAttrOption[K comparable](mp *sync.Map, key K, attrSetGenFunc func() attribute.Set) metric.MeasurementOption {
	attrSet, ok := mp.Load(key)
	if ok {
		return attrSet.(metric.MeasurementOption)
	}
	v, _ := mp.LoadOrStore(key, metric.WithAttributeSet(attrSetGenFunc()))
	return v.(metric.MeasurementOption)
}

func fsOpsAttrOption(fsOps string) metric.MeasurementOption {
	return loadOrStoreAttrOption(&fsOpsOptionCache, fsOps,
		func() attribute.Set {
			return attribute.NewSet(fsOpKey.String(fsOps))
		})
}

// otelMetrics maintains the list of all metrics computed in GCSFuse.
type otelMetrics struct {
	fsOpsCount   metric.Int64Counter
	fsOpsLatency metric.Float64Histogram
}

func (o *otelMetrics) Flush() {}

func (o *otelMetrics) FsOpsCount(ctx context.Context, inc int64, fsOp string) {
	o.fsOpsCount.Add(ctx, inc, fsOpsAttrOption(fsOp))
}

func (o *otelMetrics) FsOpsLatency(ctx context.Context, latency time.Duration, fsOp string) {
	o.fsOpsLatency.Record(ctx, float64(latency.Microseconds()), fsOpsAttrOption(fsOp))
}

func NewOTelMetrics(ctx context.Context, _ int, _ int, _ func()) (testutils.MetricHandle, error) {
	fsOpsCount, err1 := fsOpsMeter.Int64Counter("fs/ops_count", metric.WithDescription("The cumulative number of ops processed by the file system."))
	fsOpsLatency, err2 := fsOpsMeter.Float64Histogram("fs/ops_latency", metric.WithDescription("The cumulative distribution of file system operation latencies"), metric.WithUnit("us"),
		metric.WithExplicitBucketBoundaries(1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000))

	if err := errors.Join(err1, err2); err != nil {
		return nil, err
	}

	return &otelMetrics{
		fsOpsCount:   fsOpsCount,
		fsOpsLatency: fsOpsLatency,
	}, nil
}
