// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build (darwin || dragonfly || freebsd || (linux && !ppc64 && !ppc64le) || netbsd || openbsd || solaris) && gc
// +build darwin dragonfly freebsd linux,!ppc64,!ppc64le netbsd openbsd solaris
// +build gcpackage uniximport "syscall"
 Syscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) Syscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno) RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno) RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)
