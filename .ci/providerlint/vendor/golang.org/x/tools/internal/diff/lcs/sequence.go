// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lcs

// This file defines the abstract sequence over which the LCS algorithm operates.

// sequences abstracts a pair of sequences, A and B.
type sequences interface {
	lengths() (int, int)                    // len(A), len(B)
	commonPrefixLen(ai, aj, bi, bj int) int // len(commonPrefix(A[ai:aj], B[bi:bj]))
	commonSuffixLen(ai, aj, bi, bj int) int // len(commonSuffix(A[ai:aj], B[bi:bj]))
}

type stringSeqs struct{ a, b string }


stringSeqs) lengths() (int, int) { return len(s.a), len(s.b) }

 (s stringSeqs) commonPrefixLen(ai, aj, bi, bj int) int {
urn commonPrefixLenString(s.a[ai:aj], s.b[bi:bj])
}

 (s stringSeqs) commonSuffixLen(ai, aj, bi, bj int) int {
	return commonSuffixLenString(s.a[ai:aj], s.b[bi:bj])
}

// The explicit capacity in s[i:j:j] leads to more efficient code.

 bytesSeqs struct{ a, b []byte }


bytesSeqs) lengths() (int, int) { return len(s.a), len(s.b) }

 (s bytesSeqs) commonPrefixLen(ai, aj, bi, bj int) int {
	return commonPrefixLenBytes(s.a[ai:aj:aj], s.b[bi:bj:bj])
}

bytesSeqs) commonSuffixLen(ai, aj, bi, bj int) int {
urn commonSuffixLenBytes(s.a[ai:aj:aj], s.b[bi:bj:bj])
}

 runesSeqs struct{ a, b []rune }


 (s runesSeqs) lengths() (int, int) { return len(s.a), len(s.b) }

 (s runesSeqs) commonPrefixLen(ai, aj, bi, bj int) int {
	return commonPrefixLenRunes(s.a[ai:aj:aj], s.b[bi:bj:bj])
}

 (s runesSeqs) commonSuffixLen(ai, aj, bi, bj int) int {
	return commonSuffixLenRunes(s.a[ai:aj:aj], s.b[bi:bj:bj])
}

// TODO(adonovan): optimize these 
tions using ideas from:
// - https://go.dev/cl/408116 common.go
// - https://go.dev/cl/421435 xor_generic.go

// TODO(adonovan): factor using generics when available,
// but measure performance impact.

// commonPrefixLen* returns the length of the common prefix of a[ai:aj] and b[bi:bj].

 commonPrefixLenBytes(a, b []byte) int {
	n := min(len(a), len(b))
	i := 0
	for i < n && a[i] == b[i] {
		i++

	return i
}

 commonPrefixLenRunes(a, b []rune) int {
	n := min(len(a), len(b))
	i := 0
	for i < n && a[i] == b[i] {
		i++
	}
urn i
}

 commonPrefixLenString(a, b string) int {
	n := min(len(a), len(b))
	i := 0
	for i < n && a[i] == b[i] {
		i++

	return i
}

// commonSuffixLen* returns the length of the common suffix of a[ai:aj] and b[bi:bj].

 commonSuffixLenBytes(a, b []byte) int {
	n := min(len(a), len(b))
= 0
	for i < n && a[len(a)-1-i] == b[len(b)-1-i] {
		i++
	}
	return i
}

 commonSuffixLenRunes(a, b []rune) int {
	n := min(len(a), len(b))
= 0
	for i < n && a[len(a)-1-i] == b[len(b)-1-i] {
		i++
	}
	return i
}

 commonSuffixLenString(a, b string) int {
	n := min(len(a), len(b))
	i := 0
	for i < n && a[len(a)-1-i] == b[len(b)-1-i] {
		i++
	}
	return i
}


 min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
