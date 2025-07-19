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

package asyncblocking

import (
	"context"
	"errors"
	"sync"
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

type MetricHandle interface {
	FsOpsCount(
		inc int64, fsOp string,
	)
	FsOpsLatency(
		ctx context.Context, duration time.Duration, fsOp string,
	)
	Flush()
}

type otelMetrics struct {
	ch                                     chan func()
	wg                                     sync.WaitGroup
	chFullFn                               func()
	fsOpsCountFsOpBatchForgetAtomic        *atomic.Int64
	fsOpsCountFsOpCreateFileAtomic         *atomic.Int64
	fsOpsCountFsOpCreateLinkAtomic         *atomic.Int64
	fsOpsCountFsOpCreateSymlinkAtomic      *atomic.Int64
	fsOpsCountFsOpFallocateAtomic          *atomic.Int64
	fsOpsCountFsOpFlushFileAtomic          *atomic.Int64
	fsOpsCountFsOpForgetInodeAtomic        *atomic.Int64
	fsOpsCountFsOpGetInodeAttributesAtomic *atomic.Int64
	fsOpsCountFsOpGetXattrAtomic           *atomic.Int64
	fsOpsCountFsOpListXattrAtomic          *atomic.Int64
	fsOpsCountFsOpLookUpInodeAtomic        *atomic.Int64
	fsOpsCountFsOpMkDirAtomic              *atomic.Int64
	fsOpsCountFsOpMkNodeAtomic             *atomic.Int64
	fsOpsCountFsOpOpenDirAtomic            *atomic.Int64
	fsOpsCountFsOpOpenFileAtomic           *atomic.Int64
	fsOpsCountFsOpReadDirAtomic            *atomic.Int64
	fsOpsCountFsOpReadFileAtomic           *atomic.Int64
	fsOpsCountFsOpReadSymlinkAtomic        *atomic.Int64
	fsOpsCountFsOpReleaseDirHandleAtomic   *atomic.Int64
	fsOpsCountFsOpReleaseFileHandleAtomic  *atomic.Int64
	fsOpsCountFsOpRemoveXattrAtomic        *atomic.Int64
	fsOpsCountFsOpRenameAtomic             *atomic.Int64
	fsOpsCountFsOpRmDirAtomic              *atomic.Int64
	fsOpsCountFsOpSetInodeAttributesAtomic *atomic.Int64
	fsOpsCountFsOpSetXattrAtomic           *atomic.Int64
	fsOpsCountFsOpStatFSAtomic             *atomic.Int64
	fsOpsCountFsOpSyncFSAtomic             *atomic.Int64
	fsOpsCountFsOpSyncFileAtomic           *atomic.Int64
	fsOpsCountFsOpUnlinkAtomic             *atomic.Int64
	fsOpsCountFsOpWriteFileAtomic          *atomic.Int64
	fsOpsLatency                           metric.Int64Histogram
}

func (o *otelMetrics) FsOpsCount(
	ctx context.Context, inc int64, fsOp string,
) {
	select {
	case o.ch <- func() {
		switch fsOp {
		case "BatchForget":
			o.fsOpsCountFsOpBatchForgetAtomic.Add(inc)
		case "CreateFile":
			o.fsOpsCountFsOpCreateFileAtomic.Add(inc)
		case "CreateLink":
			o.fsOpsCountFsOpCreateLinkAtomic.Add(inc)
		case "CreateSymlink":
			o.fsOpsCountFsOpCreateSymlinkAtomic.Add(inc)
		case "Fallocate":
			o.fsOpsCountFsOpFallocateAtomic.Add(inc)
		case "FlushFile":
			o.fsOpsCountFsOpFlushFileAtomic.Add(inc)
		case "ForgetInode":
			o.fsOpsCountFsOpForgetInodeAtomic.Add(inc)
		case "GetInodeAttributes":
			o.fsOpsCountFsOpGetInodeAttributesAtomic.Add(inc)
		case "GetXattr":
			o.fsOpsCountFsOpGetXattrAtomic.Add(inc)
		case "ListXattr":
			o.fsOpsCountFsOpListXattrAtomic.Add(inc)
		case "LookUpInode":
			o.fsOpsCountFsOpLookUpInodeAtomic.Add(inc)
		case "MkDir":
			o.fsOpsCountFsOpMkDirAtomic.Add(inc)
		case "MkNode":
			o.fsOpsCountFsOpMkNodeAtomic.Add(inc)
		case "OpenDir":
			o.fsOpsCountFsOpOpenDirAtomic.Add(inc)
		case "OpenFile":
			o.fsOpsCountFsOpOpenFileAtomic.Add(inc)
		case "ReadDir":
			o.fsOpsCountFsOpReadDirAtomic.Add(inc)
		case "ReadFile":
			o.fsOpsCountFsOpReadFileAtomic.Add(inc)
		case "ReadSymlink":
			o.fsOpsCountFsOpReadSymlinkAtomic.Add(inc)
		case "ReleaseDirHandle":
			o.fsOpsCountFsOpReleaseDirHandleAtomic.Add(inc)
		case "ReleaseFileHandle":
			o.fsOpsCountFsOpReleaseFileHandleAtomic.Add(inc)
		case "RemoveXattr":
			o.fsOpsCountFsOpRemoveXattrAtomic.Add(inc)
		case "Rename":
			o.fsOpsCountFsOpRenameAtomic.Add(inc)
		case "RmDir":
			o.fsOpsCountFsOpRmDirAtomic.Add(inc)
		case "SetInodeAttributes":
			o.fsOpsCountFsOpSetInodeAttributesAtomic.Add(inc)
		case "SetXattr":
			o.fsOpsCountFsOpSetXattrAtomic.Add(inc)
		case "StatFS":
			o.fsOpsCountFsOpStatFSAtomic.Add(inc)
		case "SyncFS":
			o.fsOpsCountFsOpSyncFSAtomic.Add(inc)
		case "SyncFile":
			o.fsOpsCountFsOpSyncFileAtomic.Add(inc)
		case "Unlink":
			o.fsOpsCountFsOpUnlinkAtomic.Add(inc)
		case "WriteFile":
			o.fsOpsCountFsOpWriteFileAtomic.Add(inc)
		}

	}: // Do nothing
	default: // Unblock writes to channel if it's full.
		o.chFullFn()
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
		o.chFullFn()
	}
}

func (o *otelMetrics) Flush() {
	close(o.ch)
	o.wg.Wait()
}

func NewOTelMetrics(ctx context.Context, workers int, bufferSize int, chFullFn func()) (testutils.MetricHandle, error) {
	ch := make(chan func(), bufferSize)
	meter := otel.Meter("gcsfuse")
	var fsOpsCountFsOpBatchForgetAtomic,
		fsOpsCountFsOpCreateFileAtomic,
		fsOpsCountFsOpCreateLinkAtomic,
		fsOpsCountFsOpCreateSymlinkAtomic,
		fsOpsCountFsOpFallocateAtomic,
		fsOpsCountFsOpFlushFileAtomic,
		fsOpsCountFsOpForgetInodeAtomic,
		fsOpsCountFsOpGetInodeAttributesAtomic,
		fsOpsCountFsOpGetXattrAtomic,
		fsOpsCountFsOpListXattrAtomic,
		fsOpsCountFsOpLookUpInodeAtomic,
		fsOpsCountFsOpMkDirAtomic,
		fsOpsCountFsOpMkNodeAtomic,
		fsOpsCountFsOpOpenDirAtomic,
		fsOpsCountFsOpOpenFileAtomic,
		fsOpsCountFsOpReadDirAtomic,
		fsOpsCountFsOpReadFileAtomic,
		fsOpsCountFsOpReadSymlinkAtomic,
		fsOpsCountFsOpReleaseDirHandleAtomic,
		fsOpsCountFsOpReleaseFileHandleAtomic,
		fsOpsCountFsOpRemoveXattrAtomic,
		fsOpsCountFsOpRenameAtomic,
		fsOpsCountFsOpRmDirAtomic,
		fsOpsCountFsOpSetInodeAttributesAtomic,
		fsOpsCountFsOpSetXattrAtomic,
		fsOpsCountFsOpStatFSAtomic,
		fsOpsCountFsOpSyncFSAtomic,
		fsOpsCountFsOpSyncFileAtomic,
		fsOpsCountFsOpUnlinkAtomic,
		fsOpsCountFsOpWriteFileAtomic atomic.Int64

	_, err0 := meter.Int64ObservableCounter("fs/ops_count",
		metric.WithDescription("The cumulative number of ops processed by the file system."),
		metric.WithUnit(""),
		metric.WithInt64Callback(func(_ context.Context, obsrv metric.Int64Observer) error {
			obsrv.Observe(fsOpsCountFsOpBatchForgetAtomic.Load(), fsOpsCountFsOpBatchForgetAttrSet)
			obsrv.Observe(fsOpsCountFsOpCreateFileAtomic.Load(), fsOpsCountFsOpCreateFileAttrSet)
			obsrv.Observe(fsOpsCountFsOpCreateLinkAtomic.Load(), fsOpsCountFsOpCreateLinkAttrSet)
			obsrv.Observe(fsOpsCountFsOpCreateSymlinkAtomic.Load(), fsOpsCountFsOpCreateSymlinkAttrSet)
			obsrv.Observe(fsOpsCountFsOpFallocateAtomic.Load(), fsOpsCountFsOpFallocateAttrSet)
			obsrv.Observe(fsOpsCountFsOpFlushFileAtomic.Load(), fsOpsCountFsOpFlushFileAttrSet)
			obsrv.Observe(fsOpsCountFsOpForgetInodeAtomic.Load(), fsOpsCountFsOpForgetInodeAttrSet)
			obsrv.Observe(fsOpsCountFsOpGetInodeAttributesAtomic.Load(), fsOpsCountFsOpGetInodeAttributesAttrSet)
			obsrv.Observe(fsOpsCountFsOpGetXattrAtomic.Load(), fsOpsCountFsOpGetXattrAttrSet)
			obsrv.Observe(fsOpsCountFsOpListXattrAtomic.Load(), fsOpsCountFsOpListXattrAttrSet)
			obsrv.Observe(fsOpsCountFsOpLookUpInodeAtomic.Load(), fsOpsCountFsOpLookUpInodeAttrSet)
			obsrv.Observe(fsOpsCountFsOpMkDirAtomic.Load(), fsOpsCountFsOpMkDirAttrSet)
			obsrv.Observe(fsOpsCountFsOpMkNodeAtomic.Load(), fsOpsCountFsOpMkNodeAttrSet)
			obsrv.Observe(fsOpsCountFsOpOpenDirAtomic.Load(), fsOpsCountFsOpOpenDirAttrSet)
			obsrv.Observe(fsOpsCountFsOpOpenFileAtomic.Load(), fsOpsCountFsOpOpenFileAttrSet)
			obsrv.Observe(fsOpsCountFsOpReadDirAtomic.Load(), fsOpsCountFsOpReadDirAttrSet)
			obsrv.Observe(fsOpsCountFsOpReadFileAtomic.Load(), fsOpsCountFsOpReadFileAttrSet)
			obsrv.Observe(fsOpsCountFsOpReadSymlinkAtomic.Load(), fsOpsCountFsOpReadSymlinkAttrSet)
			obsrv.Observe(fsOpsCountFsOpReleaseDirHandleAtomic.Load(), fsOpsCountFsOpReleaseDirHandleAttrSet)
			obsrv.Observe(fsOpsCountFsOpReleaseFileHandleAtomic.Load(), fsOpsCountFsOpReleaseFileHandleAttrSet)
			obsrv.Observe(fsOpsCountFsOpRemoveXattrAtomic.Load(), fsOpsCountFsOpRemoveXattrAttrSet)
			obsrv.Observe(fsOpsCountFsOpRenameAtomic.Load(), fsOpsCountFsOpRenameAttrSet)
			obsrv.Observe(fsOpsCountFsOpRmDirAtomic.Load(), fsOpsCountFsOpRmDirAttrSet)
			obsrv.Observe(fsOpsCountFsOpSetInodeAttributesAtomic.Load(), fsOpsCountFsOpSetInodeAttributesAttrSet)
			obsrv.Observe(fsOpsCountFsOpSetXattrAtomic.Load(), fsOpsCountFsOpSetXattrAttrSet)
			obsrv.Observe(fsOpsCountFsOpStatFSAtomic.Load(), fsOpsCountFsOpStatFSAttrSet)
			obsrv.Observe(fsOpsCountFsOpSyncFSAtomic.Load(), fsOpsCountFsOpSyncFSAttrSet)
			obsrv.Observe(fsOpsCountFsOpSyncFileAtomic.Load(), fsOpsCountFsOpSyncFileAttrSet)
			obsrv.Observe(fsOpsCountFsOpUnlinkAtomic.Load(), fsOpsCountFsOpUnlinkAttrSet)
			obsrv.Observe(fsOpsCountFsOpWriteFileAtomic.Load(), fsOpsCountFsOpWriteFileAttrSet)
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

	o := &otelMetrics{
		ch:                                     ch,
		chFullFn:                               chFullFn,
		fsOpsCountFsOpBatchForgetAtomic:        &fsOpsCountFsOpBatchForgetAtomic,
		fsOpsCountFsOpCreateFileAtomic:         &fsOpsCountFsOpCreateFileAtomic,
		fsOpsCountFsOpCreateLinkAtomic:         &fsOpsCountFsOpCreateLinkAtomic,
		fsOpsCountFsOpCreateSymlinkAtomic:      &fsOpsCountFsOpCreateSymlinkAtomic,
		fsOpsCountFsOpFallocateAtomic:          &fsOpsCountFsOpFallocateAtomic,
		fsOpsCountFsOpFlushFileAtomic:          &fsOpsCountFsOpFlushFileAtomic,
		fsOpsCountFsOpForgetInodeAtomic:        &fsOpsCountFsOpForgetInodeAtomic,
		fsOpsCountFsOpGetInodeAttributesAtomic: &fsOpsCountFsOpGetInodeAttributesAtomic,
		fsOpsCountFsOpGetXattrAtomic:           &fsOpsCountFsOpGetXattrAtomic,
		fsOpsCountFsOpListXattrAtomic:          &fsOpsCountFsOpListXattrAtomic,
		fsOpsCountFsOpLookUpInodeAtomic:        &fsOpsCountFsOpLookUpInodeAtomic,
		fsOpsCountFsOpMkDirAtomic:              &fsOpsCountFsOpMkDirAtomic,
		fsOpsCountFsOpMkNodeAtomic:             &fsOpsCountFsOpMkNodeAtomic,
		fsOpsCountFsOpOpenDirAtomic:            &fsOpsCountFsOpOpenDirAtomic,
		fsOpsCountFsOpOpenFileAtomic:           &fsOpsCountFsOpOpenFileAtomic,
		fsOpsCountFsOpReadDirAtomic:            &fsOpsCountFsOpReadDirAtomic,
		fsOpsCountFsOpReadFileAtomic:           &fsOpsCountFsOpReadFileAtomic,
		fsOpsCountFsOpReadSymlinkAtomic:        &fsOpsCountFsOpReadSymlinkAtomic,
		fsOpsCountFsOpReleaseDirHandleAtomic:   &fsOpsCountFsOpReleaseDirHandleAtomic,
		fsOpsCountFsOpReleaseFileHandleAtomic:  &fsOpsCountFsOpReleaseFileHandleAtomic,
		fsOpsCountFsOpRemoveXattrAtomic:        &fsOpsCountFsOpRemoveXattrAtomic,
		fsOpsCountFsOpRenameAtomic:             &fsOpsCountFsOpRenameAtomic,
		fsOpsCountFsOpRmDirAtomic:              &fsOpsCountFsOpRmDirAtomic,
		fsOpsCountFsOpSetInodeAttributesAtomic: &fsOpsCountFsOpSetInodeAttributesAtomic,
		fsOpsCountFsOpSetXattrAtomic:           &fsOpsCountFsOpSetXattrAtomic,
		fsOpsCountFsOpStatFSAtomic:             &fsOpsCountFsOpStatFSAtomic,
		fsOpsCountFsOpSyncFSAtomic:             &fsOpsCountFsOpSyncFSAtomic,
		fsOpsCountFsOpSyncFileAtomic:           &fsOpsCountFsOpSyncFileAtomic,
		fsOpsCountFsOpUnlinkAtomic:             &fsOpsCountFsOpUnlinkAtomic,
		fsOpsCountFsOpWriteFileAtomic:          &fsOpsCountFsOpWriteFileAtomic,
		fsOpsLatency:                           fsOpsLatency,
	}
	o.wg.Add(workers)
	for range workers {
		go func() {
			defer o.wg.Done()
			for {
				f, ok := <-ch
				if !ok {
					return
				}
				f()
			}
		}()
	}
	return o, nil
}
