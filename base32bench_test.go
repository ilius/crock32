// Copyright (c) 2013-2014 The btcsuite developers
// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package crock32_test

import (
	"bytes"
	"testing"

	"github.com/ilius/chunk32/crock32"
)

func BenchmarkBase32Encode(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 5000)
	b.SetBytes(int64(len(data)))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		crock32.Encode(data)
	}
}

func BenchmarkBase32Decode(b *testing.B) {
	b.StopTimer()
	data := bytes.Repeat([]byte{0xff}, 5000)
	encoded := crock32.Encode(data)
	b.SetBytes(int64(len(encoded)))
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		crock32.Decode(encoded)
	}
}
