// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build gccgo && linux && amd64
// +build gccgo,linux,amd64package uniximport "syscall"//extern gettimeofday realGettimeofday(*Timeval, *byte) int32
 gettimeofday(tv *Timeval) (err syscall.Errno) {
	r := realGettimeofday(tv, nil)
	if r < 0 {
		return syscall.GetErrno()
	}
	return 0
}
