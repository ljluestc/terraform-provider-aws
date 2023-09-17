// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build 386 && freebsd
// +build 386,freebsdpackage uniximport (
"syscall"
"unsafe"
)
 setTimespec(sec, nsec int64) Timespec {
return Timespec{Sec: int32(sec), Nsec: int32(nsec)}
}
 setTimeval(sec, usec int64) Timeval {
return Timeval{Sec: int32(sec), Usec: int32(usec)}
}
 SetKevent(k *Kevent_t, fd, mode, flags int) {
k.Ident = uint32(fd)
k.Filter = int16(mode)
k.Flags = uint16(flags)
}
 (iov *Iovec) SetLen(length int) {
iov.Len = uint32(length)
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
 (d *PtraceIoDesc) SetLen(length int) {
d.Len = uint32(length)
}
 sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
var writtenOut uint64 = 0
_, _, e1 := Syscall9(SYS_SENDFILE, uintptr(infd), uintptr(outfd), uintptr(*offset), uintptr((*offset)>>32), uintptr(count), 0, uintptr(unsafe.Pointer(&writtenOut)), 0, 0)written = int(writtenOut)if e1 != 0 {
err = e1
}
return
}
 Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno)
 PtraceGetFsBase(pid int, fsbase *int64) (err error) {
return ptracePtr(PT_GETFSBASE, pid, unsafe.Pointer(fsbase), 0)
}
