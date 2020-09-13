// Copyright (c) 2013-2017 The btcsuite developers
// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package crock32_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/ilius/crock32"
)

var stringTests = []struct {
	in  string
	out string
}{
	{"", ""},
	{" ", "10"},
	{"-", "1D"},
	{"0", "1G"},
	{"1", "1H"},
	{"-1", "B9H"},
	{"11", "C9H"},
	{"abc", "62RK3"},
	{"1234598760", "64S36D1N74W3EDHG"},
	{"abcdefghijklmnopqrstuvwxyz", "31C9HP8SB6CXM6JTKBDHPPWVVGE5S76X3NESVQGYBT"},
	{"00000000000000000000000000000000000000000000000000000000000000", "C1G60R30C1G60R30C1G60R30C1G60R30C1G60R30C1G60R30C1G60R30C1G60R30C1G60R30C1G60R30C1G60R30C1G60R30C1G"},
}

var invalidStringTests = []struct {
	in  string
	out string
}{
	{"3mUJr0", ""},
	{"!@#$%^&*()-_=+~`", ""},
}

var hexTests = []struct {
	in  string
	out string
}{
	{"61", "31"},
	{"626262", "64rk2"},
	{"636363", "66rv3"},
	{"73696d706c792061206c6f6e6720737472696e67", "edmptw3cf4g6283cdxq6e83kehs6jvk7"},
	{"00eb15231dfceb60925886b67d065299925915aeb172c06647", "03nha8rxzknp14jrgtv7t1jjk695j5dep5sc0sj7"},
	{"516b6fcd0f", "a5npzk8f"},
	{"bf4f89001e670274dd", "2zmz2803skg4x6x"},
	{"572e4794", "1bjwhwm"},
	{"ecac89cad93923c02321", "xjp8kjps74hw08s1"},
	{"10c8511e", "8cgm8y"},
	{"00000000000000000000", "0000000000"},
}

func TestBase32(t *testing.T) {
	// Encode tests
	for _, test := range stringTests {
		tmp := []byte(test.in)
		if res := crock32.Encode(tmp); res != test.out {
			t.Errorf(
				"Encode test in=%#v failed: got: %s want: %s",
				test.in,
				res,
				test.out,
			)
			continue
		}
	}

	// Decode tests
	for x, test := range hexTests {
		b, err := hex.DecodeString(test.in)
		if err != nil {
			t.Errorf("hex.DecodeString failed failed #%d: got: %s", x, test.in)
			continue
		}
		res, err := crock32.Decode(test.out)
		if err != nil {
			t.Logf("%v ---> %v\n", test.in, crock32.Encode(b))
			t.Errorf("Decode returned error: %v", err)
			continue
		}
		if !bytes.Equal(res, b) {
			t.Logf("%v ---> %v\n", test.in, crock32.Encode(b))
			t.Errorf(
				"Decode test out=%#v failed: got: %q want: %q",
				test.out,
				hex.EncodeToString(res),
				test.in,
			)
			continue
		}
	}

	// Decode with invalid input
	for _, test := range invalidStringTests {
		res, _ := crock32.Decode(test.in)
		// TODO: check err == "invalid character %#v"
		if string(res) != test.out {
			t.Errorf(
				"Decode invalidString test in=%v failed: got: %q want: %q",
				test.in,
				res,
				test.out,
			)
			continue
		}
	}
}
