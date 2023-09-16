// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build arm64 && darwin
// +build arm64,darwinpackage uniximport "syscall"
 setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: sec, Nsec: nsec}
}
 setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: sec, Usec: int32(usec)}
}
 SetKevent(k *Kevent_t, fd, mode, flags int) {
	k.Ident = uint64(fd)
	k.Filter = int16(mode)
	k.Flags = uint16(flags)
}
 (iov *Iovec) SetLen(length int) {
	iov.Len = uint64(length)
}
 (msghdr *Msghdr) SetControllen(length int) {
	msghdr.Controllen = uint32(length)
}
 (msghdr *Msghdr) SetIovlen(length int) {
	msghdr.Iovlen = int32(length)
}
 (cmsg *Cmsghdr) SetLen(length int) {
	cmsg.Len = uint32(length)
}
 Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno) // sic//sys	Fstat(fd int, stat *Stat_t) (err error)
//sys	Fstatat(fd int, path string, stat *Stat_t, flags int) (err error)
//sys	Fstatfs(fd int, stat *Statfs_t) (err error)
//sys	getfsstat(buf unsafe.Pointer, size uintptr, flags int) (n int, err error) = SYS_GETFSSTAT
//sys	Lstat(path string, stat *Stat_t) (err error)
//sys	ptrace1(request int, pid int, addr uintptr, data uintptr) (err error) = SYS_ptrace
//sys	ptrace1Ptr(request int, pid int, addr unsafe.Pointer, data uintptr) (err error) = SYS_ptrace
//sys	Stat(path string, stat *Stat_t) (err error)
//sys	Statfs(path string, stat *Statfs_t) (err error)
