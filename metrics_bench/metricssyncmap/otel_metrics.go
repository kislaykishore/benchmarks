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

// **** DO NOT EDIT - FILE IS AUTO-GENERATED ****

package metricssyncmap

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/kislaykishore/benchmarks/metrics_bench/testutils"
)

var (
	fsOpsCountFsOpBatchForgetAttrSet          = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "BatchForget")))
	fsOpsCountFsOpCreateFileAttrSet           = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "CreateFile")))
	fsOpsCountFsOpCreateLinkAttrSet           = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "CreateLink")))
	fsOpsCountFsOpCreateSymlinkAttrSet        = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "CreateSymlink")))
	fsOpsCountFsOpFallocateAttrSet            = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "Fallocate")))
	fsOpsCountFsOpFlushFileAttrSet            = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "FlushFile")))
	fsOpsCountFsOpForgetInodeAttrSet          = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ForgetInode")))
	fsOpsCountFsOpGetInodeAttributesAttrSet   = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "GetInodeAttributes")))
	fsOpsCountFsOpGetXattrAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "GetXattr")))
	fsOpsCountFsOpListXattrAttrSet            = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ListXattr")))
	fsOpsCountFsOpLookUpInodeAttrSet          = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "LookUpInode")))
	fsOpsCountFsOpMkDirAttrSet                = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "MkDir")))
	fsOpsCountFsOpMkNodeAttrSet               = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "MkNode")))
	fsOpsCountFsOpOpenDirAttrSet              = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "OpenDir")))
	fsOpsCountFsOpOpenFileAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "OpenFile")))
	fsOpsCountFsOpReadDirAttrSet              = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReadDir")))
	fsOpsCountFsOpReadFileAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReadFile")))
	fsOpsCountFsOpReadSymlinkAttrSet          = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReadSymlink")))
	fsOpsCountFsOpReleaseDirHandleAttrSet     = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReleaseDirHandle")))
	fsOpsCountFsOpReleaseFileHandleAttrSet    = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReleaseFileHandle")))
	fsOpsCountFsOpRemoveXattrAttrSet          = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "RemoveXattr")))
	fsOpsCountFsOpRenameAttrSet               = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "Rename")))
	fsOpsCountFsOpRmDirAttrSet                = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "RmDir")))
	fsOpsCountFsOpSetInodeAttributesAttrSet   = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "SetInodeAttributes")))
	fsOpsCountFsOpSetXattrAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "SetXattr")))
	fsOpsCountFsOpStatFSAttrSet               = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "StatFS")))
	fsOpsCountFsOpSyncFSAttrSet               = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "SyncFS")))
	fsOpsCountFsOpSyncFileAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "SyncFile")))
	fsOpsCountFsOpUnlinkAttrSet               = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "Unlink")))
	fsOpsCountFsOpWriteFileAttrSet            = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "WriteFile")))
	fsOpsLatencyFsOpBatchForgetAttrSet        = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "BatchForget")))
	fsOpsLatencyFsOpCreateFileAttrSet         = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "CreateFile")))
	fsOpsLatencyFsOpCreateLinkAttrSet         = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "CreateLink")))
	fsOpsLatencyFsOpCreateSymlinkAttrSet      = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "CreateSymlink")))
	fsOpsLatencyFsOpFallocateAttrSet          = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "Fallocate")))
	fsOpsLatencyFsOpFlushFileAttrSet          = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "FlushFile")))
	fsOpsLatencyFsOpForgetInodeAttrSet        = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ForgetInode")))
	fsOpsLatencyFsOpGetInodeAttributesAttrSet = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "GetInodeAttributes")))
	fsOpsLatencyFsOpGetXattrAttrSet           = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "GetXattr")))
	fsOpsLatencyFsOpListXattrAttrSet          = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ListXattr")))
	fsOpsLatencyFsOpLookUpInodeAttrSet        = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "LookUpInode")))
	fsOpsLatencyFsOpMkDirAttrSet              = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "MkDir")))
	fsOpsLatencyFsOpMkNodeAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "MkNode")))
	fsOpsLatencyFsOpOpenDirAttrSet            = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "OpenDir")))
	fsOpsLatencyFsOpOpenFileAttrSet           = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "OpenFile")))
	fsOpsLatencyFsOpReadDirAttrSet            = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReadDir")))
	fsOpsLatencyFsOpReadFileAttrSet           = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReadFile")))
	fsOpsLatencyFsOpReadSymlinkAttrSet        = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReadSymlink")))
	fsOpsLatencyFsOpReleaseDirHandleAttrSet   = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReleaseDirHandle")))
	fsOpsLatencyFsOpReleaseFileHandleAttrSet  = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "ReleaseFileHandle")))
	fsOpsLatencyFsOpRemoveXattrAttrSet        = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "RemoveXattr")))
	fsOpsLatencyFsOpRenameAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "Rename")))
	fsOpsLatencyFsOpRmDirAttrSet              = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "RmDir")))
	fsOpsLatencyFsOpSetInodeAttributesAttrSet = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "SetInodeAttributes")))
	fsOpsLatencyFsOpSetXattrAttrSet           = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "SetXattr")))
	fsOpsLatencyFsOpStatFSAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "StatFS")))
	fsOpsLatencyFsOpSyncFSAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "SyncFS")))
	fsOpsLatencyFsOpSyncFileAttrSet           = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "SyncFile")))
	fsOpsLatencyFsOpUnlinkAttrSet             = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "Unlink")))
	fsOpsLatencyFsOpWriteFileAttrSet          = metric.WithAttributeSet(attribute.NewSet(attribute.String("fs_op", "WriteFile")))
)
var fsOpsCountAttrSets = map[string]metric.MeasurementOption{"BatchForget": fsOpsCountFsOpBatchForgetAttrSet, "CreateFile": fsOpsCountFsOpCreateFileAttrSet, "CreateLink": fsOpsCountFsOpCreateLinkAttrSet, "CreateSymlink": fsOpsCountFsOpCreateSymlinkAttrSet, "Fallocate": fsOpsCountFsOpFallocateAttrSet, "FlushFile": fsOpsCountFsOpFlushFileAttrSet, "ForgetInode": fsOpsCountFsOpForgetInodeAttrSet, "GetInodeAttributes": fsOpsCountFsOpGetInodeAttributesAttrSet, "GetXattr": fsOpsCountFsOpGetXattrAttrSet, "ListXattr": fsOpsCountFsOpListXattrAttrSet, "LookUpInode": fsOpsCountFsOpLookUpInodeAttrSet, "MkDir": fsOpsCountFsOpMkDirAttrSet, "MkNode": fsOpsCountFsOpMkNodeAttrSet, "OpenDir": fsOpsCountFsOpOpenDirAttrSet, "OpenFile": fsOpsCountFsOpOpenFileAttrSet, "ReadDir": fsOpsCountFsOpReadDirAttrSet, "ReadFile": fsOpsCountFsOpReadFileAttrSet, "ReadSymlink": fsOpsCountFsOpReadSymlinkAttrSet, "ReleaseDirHandle": fsOpsCountFsOpReleaseDirHandleAttrSet, "ReleaseFileHandle": fsOpsCountFsOpReleaseFileHandleAttrSet, "RemoveXattr": fsOpsCountFsOpRemoveXattrAttrSet, "Rename": fsOpsCountFsOpRenameAttrSet, "RmDir": fsOpsCountFsOpRmDirAttrSet, "SetInodeAttributes": fsOpsCountFsOpSetInodeAttributesAttrSet, "SetXattr": fsOpsCountFsOpSetXattrAttrSet, "StatFS": fsOpsCountFsOpStatFSAttrSet, "SyncFS": fsOpsCountFsOpSyncFSAttrSet, "SyncFile": fsOpsCountFsOpSyncFileAttrSet, "Unlink": fsOpsCountFsOpUnlinkAttrSet, "WriteFile": fsOpsCountFsOpWriteFileAttrSet}

type MetricHandle interface {
	FsOpsCount(
		inc int64, fsOp string,
	)
	FsOpsLatency(
		ctx context.Context, duration time.Duration, fsOp string,
	)
}

type otelMetrics struct {
	ch                 chan func()
	fsOpsCountCounters map[string]*atomic.Int64
	fsOpsLatency       metric.Int64Histogram
}

func (o *otelMetrics) FsOpsCount(
	ctx context.Context, inc int64, fsOp string,
) {
	if atomic, ok := o.fsOpsCountCounters[fsOp]; ok {
		atomic.Add(inc)
	}
}

func (o *otelMetrics) FsOpsLatency(
	ctx context.Context, latency time.Duration, fsOp string,
) {
	select {
	case o.ch <- func() {
		switch fsOp {
		case "BatchForget":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpBatchForgetAttrSet)
		case "CreateFile":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpCreateFileAttrSet)
		case "CreateLink":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpCreateLinkAttrSet)
		case "CreateSymlink":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpCreateSymlinkAttrSet)
		case "Fallocate":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpFallocateAttrSet)
		case "FlushFile":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpFlushFileAttrSet)
		case "ForgetInode":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpForgetInodeAttrSet)
		case "GetInodeAttributes":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpGetInodeAttributesAttrSet)
		case "GetXattr":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpGetXattrAttrSet)
		case "ListXattr":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpListXattrAttrSet)
		case "LookUpInode":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpLookUpInodeAttrSet)
		case "MkDir":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpMkDirAttrSet)
		case "MkNode":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpMkNodeAttrSet)
		case "OpenDir":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpOpenDirAttrSet)
		case "OpenFile":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpOpenFileAttrSet)
		case "ReadDir":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpReadDirAttrSet)
		case "ReadFile":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpReadFileAttrSet)
		case "ReadSymlink":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpReadSymlinkAttrSet)
		case "ReleaseDirHandle":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpReleaseDirHandleAttrSet)
		case "ReleaseFileHandle":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpReleaseFileHandleAttrSet)
		case "RemoveXattr":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpRemoveXattrAttrSet)
		case "Rename":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpRenameAttrSet)
		case "RmDir":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpRmDirAttrSet)
		case "SetInodeAttributes":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpSetInodeAttributesAttrSet)
		case "SetXattr":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpSetXattrAttrSet)
		case "StatFS":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpStatFSAttrSet)
		case "SyncFS":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpSyncFSAttrSet)
		case "SyncFile":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpSyncFileAttrSet)
		case "Unlink":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpUnlinkAttrSet)
		case "WriteFile":
			o.fsOpsLatency.Record(ctx, latency.Microseconds(), fsOpsLatencyFsOpWriteFileAttrSet)
		}

	}: // Do nothing
	default: // Unblock writes to channel if it's full.
	}
}

func (o *otelMetrics) Flush() {}

func NewOTelMetrics(ctx context.Context, workers int, bufferSize int, _ func()) (testutils.MetricHandle, error) {
	ch := make(chan func(), bufferSize)
	for range workers {
		go func() {
			for {
				f, ok := <-ch
				if !ok {
					return
				}
				f()
			}
		}()
	}
	meter := otel.Meter("gcsfuse")
	fsOpsCountCounters := map[string]*atomic.Int64{"BatchForget": new(atomic.Int64), "CreateFile": new(atomic.Int64), "CreateLink": new(atomic.Int64), "CreateSymlink": new(atomic.Int64), "Fallocate": new(atomic.Int64), "FlushFile": new(atomic.Int64), "ForgetInode": new(atomic.Int64), "GetInodeAttributes": new(atomic.Int64), "GetXattr": new(atomic.Int64), "ListXattr": new(atomic.Int64), "LookUpInode": new(atomic.Int64), "MkDir": new(atomic.Int64), "MkNode": new(atomic.Int64), "OpenDir": new(atomic.Int64), "OpenFile": new(atomic.Int64), "ReadDir": new(atomic.Int64), "ReadFile": new(atomic.Int64), "ReadSymlink": new(atomic.Int64), "ReleaseDirHandle": new(atomic.Int64), "ReleaseFileHandle": new(atomic.Int64), "RemoveXattr": new(atomic.Int64), "Rename": new(atomic.Int64), "RmDir": new(atomic.Int64), "SetInodeAttributes": new(atomic.Int64), "SetXattr": new(atomic.Int64), "StatFS": new(atomic.Int64), "SyncFS": new(atomic.Int64), "SyncFile": new(atomic.Int64), "Unlink": new(atomic.Int64), "WriteFile": new(atomic.Int64)}

	_, err0 := meter.Int64ObservableCounter("fs/ops_count",
		metric.WithDescription("The cumulative number of ops processed by the file system."),
		metric.WithUnit(""),
		metric.WithInt64Callback(func(_ context.Context, obsrv metric.Int64Observer) error {
			for k, v := range fsOpsCountCounters {
				val := v.Load()
				if val > 0 {
					obsrv.Observe(val, fsOpsCountAttrSets[k])
				}
			}
			return nil
		}))

	fsOpsLatency, err1 := meter.Int64Histogram("fs/ops_latency",
		metric.WithDescription("The cumulative distribution of file system operation latencies"),
		metric.WithUnit("us"),
		metric.WithExplicitBucketBoundaries(1, 2, 3, 4, 5, 6, 8, 10, 13, 16, 20, 25, 30, 40, 50, 65, 80, 100, 130, 160, 200, 250, 300, 400, 500, 650, 800, 1000, 2000, 5000, 10000, 20000, 50000, 100000))

	errs := []error{err0, err1}
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	return &otelMetrics{
		ch: ch, fsOpsCountCounters: fsOpsCountCounters, fsOpsLatency: fsOpsLatency,
	}, nil
}
