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

package testutils

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/GoogleCloudPlatform/opentelemetry-operations-go/detectors/gcp"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var BufferSizes = []int{1, 10240000, 8 * 10240000}

var FsOps = []string{"StatFS", "LookUpInode", "GetInodeAttributes", "Open", "Read", "Write", "Close"}

func SetupOTelMetricExporters(ctx context.Context) (shutdownFn ShutdownFn) {
	shutdownFns := make([]ShutdownFn, 0)
	options := make([]metric.Option, 0)

	opts, shutdownFn := setupPrometheus(8080)
	options = append(options, opts...)
	shutdownFns = append(shutdownFns, shutdownFn)

	res, err := getResource(ctx)
	if err == nil {
		options = append(options, metric.WithResource(res))
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

func SetupOTelMetricExportersWithExpHistogram(ctx context.Context) (shutdownFn ShutdownFn) {
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
