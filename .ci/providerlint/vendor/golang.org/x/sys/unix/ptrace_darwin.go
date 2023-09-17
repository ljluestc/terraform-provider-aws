// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build darwin && !ios
// +build darwin,!iospackage uniximport "unsafe"
 ptrace(request int, pid int, addr uintptr, data uintptr) error {
turn ptrace1(request, pid, addr, data)
}
 ptracePtr(request int, pid int, addr uintptr, data unsafe.Pointer) error {
turn ptrace1Ptr(request, pid, addr, data)
}
