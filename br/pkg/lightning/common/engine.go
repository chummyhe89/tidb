// Copyright 2023 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package common

import (
	"context"

	"github.com/pingcap/tidb/br/pkg/lightning/log"
)

// Range contains a start key and an end key.
type Range struct {
	Start []byte
	End   []byte // end is always exclusive except import_sstpb.SSTMeta
}

// Engine describes the common interface of local and external engine that
// local backend uses.
type Engine interface {
	// ID is the identifier of an engine.
	ID() string
	// LoadIngestData returns an IngestData that contains the data in [start, end).
	LoadIngestData(ctx context.Context, start, end []byte) (IngestData, error)
	// KVStatistics returns the total kv size and total kv length.
	KVStatistics() (totalKVSize int64, totalKVLength int64)
	// ImportedStatistics returns the imported kv size and imported kv length.
	ImportedStatistics() (importedKVSize int64, importedKVLength int64)
	// SplitRanges splits the range [startKey, endKey) into multiple ranges.
	SplitRanges(startKey, endKey []byte, sizeLimit, keysLimit int64, logger log.Logger) ([]Range, error)
	// TODO(lance6716): add more methods
}
