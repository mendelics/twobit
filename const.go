// Copyright 2015 Andrew E. Bruno. All rights reserved.
// Use of this source code is governed by a BSD style
// license that can be found in the LICENSE file.

package twobit

// SIG -
const SIG = 0x1A412743

const defaultBufSize = 4096

// BASE_N -
const BASE_N = 'N'

// BASE_T -
const BASE_T = 'T'

// BASE_C -
const BASE_C = 'C'

// BASE_A -
const BASE_A = 'A'

// BASE_G -
const BASE_G = 'G'

// BYTES2NT -
var BYTES2NT = []byte{
	BASE_T,
	BASE_C,
	BASE_A,
	BASE_G,
}

// NT2BYTES -
var NT2BYTES = []byte{}
