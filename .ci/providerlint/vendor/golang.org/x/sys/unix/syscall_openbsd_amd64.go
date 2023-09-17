// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build amd64 && openbsd
// +build amd64,openbsdpackage unix
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
ghdr.Iovlen = uint32(length)
}
 (cmsg *Cmsghdr) SetLen(length int) {
sg.Len = uint32(length)
}// SYS___SYSCTL is used by syscall_bsd.go for all BSDs, but in modern versions
// of openbsd/amd64 the syscall is called sysctl instead of __sysctl.
const SYS___SYSCTL = SYS_SYSCTL
