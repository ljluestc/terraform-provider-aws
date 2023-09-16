// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build (darwin && race) || (linux && race) || (freebsd && race)
// +build darwin,race linux,race freebsd,racepackage uniximport (
	"runtime"
	"unsafe"
)const raceenabled = true
 raceAcquire(addr unsafe.Pointer) {
	runtime.RaceAcquire(addr)
}
 raceReleaseMerge(addr unsafe.Pointer) {
	runtime.RaceReleaseMerge(addr)
}
 raceReadRange(addr unsafe.Pointer, len int) {
	runtime.RaceReadRange(addr, len)
}
 raceWriteRange(addr unsafe.Pointer, len int) {
	runtime.RaceWriteRange(addr, len)
}
