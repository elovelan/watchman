// Copyright The Moov Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package csl_us

import (
	"context"
	"os"
	"testing"

	"github.com/moov-io/base/log"

	"github.com/stretchr/testify/require"
)

func TestCSL(t *testing.T) {
	if testing.Short() {
		t.Skip("ignorning network test")
	}

	logger := log.NewNopLogger()
	dir, err := os.MkdirTemp("", "csl")
	require.NoError(t, err)

	file, err := Download(context.Background(), logger, dir)
	require.NoError(t, err)

	cslRecords, err := ReadFile(file["csl.csv"])
	require.NoError(t, err)

	if len(cslRecords.SSIs) == 0 {
		t.Error("parsed zero CSL SSI records")
	}
	if len(cslRecords.ELs) == 0 {
		t.Error("parsed zero CSL EL records")
	}
}
