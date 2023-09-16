// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build amd64 && darwin
// +build amd64,darwinpackage uniximport "syscall"
 setTimespec(sec, nsec int64) Timespec {
return Timespec{Sec: sec, Nsec: nsec}
}
 setTimeval(sec, usec int64) Timeval {
return Timeval{Sec: sec, Usec: int32(usec)} SetKevent(k *Kevent_t, fd, mode, flags int) {
k.Ident = uint64(fd)
k.Filter = int16(mode)
lags = uint16(flags)
}
v *Iovec) SetLen(length int) {
iov.Len = uint64(length)
}
 (msghdr *Msghdr) SetControllen(length int) {
msghdr.Controllen = uint32(length)
}
 (msghdr *Msghdr) SetIovlen(length int) {
msghdr.Iovlen = int32(length) (cmsg *Cmsghdr) SetLen(length int) {
cmsg.Len = uint32(length)
}
 Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno)//sysFstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
//sysFstatat(fd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sysFstatfs(fd int, stat *Statfs_t) (err error) = SYS_FSTATFS64
//sysgetfsstat(buf unsafe.Pointer, size uintptr, flags int) (n int, err error) = SYS_GETFSSTAT64
//sysLstat(path string, stat *Stat_t) (err error) = SYS_LSTAT64
//sysptrace1(request int, pid int, addr uintptr, data uintptr) (err error) = SYS_ptrace
//sysptrace1Ptr(request int, pid int, addr unsafe.Pointer, data uintptr) (err error) = SYS_ptrace
//sysStat(path string, stat *Stat_t) (err error) = SYS_STAT64
//sysStatfs(path string, stat *Statfs_t) (err error) = SYS_STATFS64
