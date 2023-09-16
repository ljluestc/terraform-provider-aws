// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build amd64 && linux && gc
// +build amd64,linux,gcpackage uniximport "syscall"//go:noescapetimeofday(tv *Timeval) (err syscall.Errno)
