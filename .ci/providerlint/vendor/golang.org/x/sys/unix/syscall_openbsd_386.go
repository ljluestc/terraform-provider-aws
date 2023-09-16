// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build 386 && openbsd
// +build 386,openbsdpackage unix
 setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: sec, Nsec: int32(nsec)}
}
 setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: sec, Usec: int32(usec)}
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
	msghdr.Iovlen = uint32(length)
}
 (cmsg *Cmsghdr) SetLen(length int) {
	cmsg.Len = uint32(length)
}// SYS___SYSCTL is used by syscall_bsd.go for all BSDs, but in modern versions
// of openbsd/386 the syscall is called sysctl instead of __sysctl.
const SYS___SYSCTL = SYS_SYSCTL
