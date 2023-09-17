// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build riscv64 && freebsd
// +build riscv64,freebsdpackage uniximport (
yscall"
nsafe"
)
 setTimespec(sec, nsec int64) Timespec {
turn Timespec{Sec: sec, Nsec: nsec}
}
 setTimeval(sec, usec int64) Timeval {
turn Timeval{Sec: sec, Usec: usec}
}
 SetKevent(k *Kevent_t, fd, mode, flags int) {
Ident = uint64(fd)
Filter = int16(mode)
Flags = uint16(flags)
}
 (iov *Iovec) SetLen(length int) {
v.Len = uint64(length)
}
 (msghdr *Msghdr) SetControllen(length int) {
ghdr.Controllen = uint32(length)
}
 (msghdr *Msghdr) SetIovlen(length int) {
ghdr.Iovlen = int32(length)
}
 (cmsg *Cmsghdr) SetLen(length int) {
sg.Len = uint32(length)
}
 (d *PtraceIoDesc) SetLen(length int) {
Len = uint64(length)
}
 sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
r writtenOut uint64 = 0
 _, e1 := Syscall9(SYS_SENDFILE, uintptr(infd), uintptr(outfd), uintptr(*offset), uintptr(count), 0, uintptr(unsafe.Pointer(&writtenOut)), 0, 0, 0)witen = int(writtenOut)ife1!= 0 {
err = e1turn
}
 Syscall9(num, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err syscall.Errno)
