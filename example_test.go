// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package crock32_test

import (
	"fmt"

	"github.com/ilius/crock32"
)

// This example demonstrates how to decode modified base32 encoded data.
func ExampleDecode() {
	// Decode example modified base32 encoded data.
	encoded := "1A6AWVM41J62X31"
	decoded, err := crock32.Decode(encoded)
	if err != nil {
		panic(err)
	}

	// Show the decoded data.
	fmt.Println("Decoded Data:", string(decoded))

	// Output:
	// Decoded Data: Test data
}

// This example demonstrates how to encode data using the modified base32
// encoding scheme.
func ExampleEncode() {
	// Encode example data with the modified base32 encoding scheme.
	data := []byte("Test data")
	encoded := crock32.Encode(data)

	// Show the encoded data.
	fmt.Println("Encoded Data:", encoded)

	// Output:
	// Encoded Data: 1A6AWVM41J62X31
}

// This example demonstrates how to decode base32Check encoded data.
// func ExampleCheckDecode() {
// 	// Decode an example base32Check encoded data.
// 	encoded := "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa"
// 	decoded, version, err := crock32.CheckDecode(encoded)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// Show the decoded data.
// 	fmt.Printf("Decoded data: %x\n", decoded)
// 	fmt.Println("Version Byte:", version)

// 	// Output:
// 	// Decoded data: 62e907b15cbf27d5425399ebf6f0fb50ebb88f18
// 	// Version Byte: 0
// }

// This example demonstrates how to encode data using the base32Check encoding
// scheme.
// func ExampleCheckEncode() {
// 	// Encode example data with the base32Check encoding scheme.
// 	data := []byte("Test data")
// 	encoded := crock32.CheckEncode(data, 0)

// 	// Show the encoded data.
// 	fmt.Println("Encoded Data:", encoded)

// 	// Output:
// 	// Encoded Data: 182iP79GRURMp7oMHDU
// }
