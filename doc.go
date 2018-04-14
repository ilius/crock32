// Copyright (c) 2014 The btcsuite developers
// Copyright (c) 2018 Saeed Rasooli <saeed.gnu@gmail.com>
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

/*
Package crock32 provides an API for working with modified base32 and Base32Check
encodings.

Modified base32 Encoding
...

Base32Check Encoding Scheme

The Base32Check encoding scheme is primarily used for Bitcoin addresses at the
time of this writing, however it can be used to generically encode arbitrary
byte arrays into human-readable strings along with a version byte that can be
used to differentiate the same payload.  For Bitcoin addresses, the extra
version is used to differentiate the network of otherwise identical public keys
which helps prevent using an address intended for one network on another.
*/
package crock32
