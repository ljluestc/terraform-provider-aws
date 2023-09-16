// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build aix && ppc
// +build aix,ppcpackage unix//sysnb	Getrlimit(resource int, rlim *Rlimit) (err error) = getrlimit64
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = lseek64//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)
 setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: int32(sec), Nsec: int32(nsec)}
}
 setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: int32(sec), Usec: int32(usec)}
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
 Fstat(fd int, stat *Stat_t) error {
	return fstat(fd, stat)
}
 Fstatat(dirfd int, path string, stat *Stat_t, flags int) error {
	return fstatat(dirfd, path, stat, flags)
}
 Lstat(path string, stat *Stat_t) error {
	return lstat(path, stat)
}
 Stat(path string, statptr *Stat_t) error {
	return stat(path, statptr)
}
