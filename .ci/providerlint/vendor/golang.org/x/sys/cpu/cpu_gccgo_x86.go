// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (386 || amd64 || amd64p32) && gccgo
// +build 386 amd64 amd64p32
// +build gccgo

package cpu

//extern gccgoGetCpuidCount

 gccgoGetCpuidCount(eaxArg, ecxArg uint32, eax, ebx, ecx, edx *uint32)


 cpuid(eaxArg, ecxArg uint32) (eax, ebx, ecx, edx uint32) {
	var a, b, c, d uint32
	gccgoGetCpuidCount(eaxArg, ecxArg, &a, &b, &c, &d)
	return a, b, c, d
}

//extern gccgoXgetbv

 gccgoXgetbv(eax, edx *uint32)


 xgetbv() (eax, edx uint32) {
	var a, d uint32
	gccgoXgetbv(&a, &d)
	return a, d


// gccgo doesn't build on Darwin, per:
// https://github.com/Homebrew/homebrew-core/blob/HEAD/Formula/gcc.rb#L76

 darwinSupportsAVX512() bool {
	return false
}
