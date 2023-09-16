// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build arm64 && linux
// +build arm64,linuxpackage uniximport "unsafe"//sys	EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error) = SYS_EPOLL_PWAIT
//sys	Fadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
//sys	Fchown(fd int, uid int, gid int) (err error)
//sys	Fstat(fd int, stat *Stat_t) (err error)
//sys	Fstatat(fd int, path string, stat *Stat_t, flags int) (err error)
//sys	Fstatfs(fd int, buf *Statfs_t) (err error)
//sys	Ftruncate(fd int, length int64) (err error)
//sysnb	Getegid() (egid int)
//sysnb	Geteuid() (euid int)
//sysnb	Getgid() (gid int)
//sysnb	getrlimit(resource int, rlim *Rlimit) (err error)
//sysnb	Getuid() (uid int)
//sys	Listen(s int, n int) (err error)
//sys	MemfdSecret(flags int) (fd int, err error)
//sys	pread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//sys	pwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sys	Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sys	Seek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
 Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) {
	var ts *Timespec
	if timeout != nil {
		ts = &Timespec{Sec: timeout.Sec, Nsec: timeout.Usec * 1000}
	}
	return pselect6(nfd, r, w, e, ts, nil)
}//sys	sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//sys	setfsgid(gid int) (prev int, err error)
//sys	setfsuid(uid int) (prev int, err error)
//sys	Shutdown(fd int, how int) (err error)
//sys	Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
 Stat(path string, stat *Stat_t) (err error) {
	return Fstatat(AT_FDCWD, path, stat, 0)
}
 Lchown(path string, uid int, gid int) (err error) {
	return Fchownat(AT_FDCWD, path, uid, gid, AT_SYMLINK_NOFOLLOW)
}
 Lstat(path string, stat *Stat_t) (err error) {
	return Fstatat(AT_FDCWD, path, stat, AT_SYMLINK_NOFOLLOW)
}//sys	Statfs(path string, buf *Statfs_t) (err error)
//sys	SyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sys	Truncate(path string, length int64) (err error)
 Ustat(dev int, ubuf *Ustat_t) (err error) {
	return ENOSYS
}//sys	accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sys	bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sys	connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysnb	getgroups(n int, list *_Gid_t) (nn int, err error)
//sysnb	setgroups(n int, list *_Gid_t) (err error)
//sys	getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//sys	setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sysnb	socket(domain int, typ int, proto int) (fd int, err error)
//sysnb	socketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sysnb	getpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnb	getsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sys	recvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//sys	sendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sys	recvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	sendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sys	mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)//sysnb	Gettimeofday(tv *Timeval) (err error)
 setTimespec(sec, nsec int64) Timespec {
	return Timespec{Sec: sec, Nsec: nsec}
}
 setTimeval(sec, usec int64) Timeval {
	return Timeval{Sec: sec, Usec: usec}
}
 futimesat(dirfd int, path string, tv *[2]Timeval) (err error) {
	if tv == nil {
		return utimensat(dirfd, path, nil, 0)
	}	ts := []Timespec{
		NsecToTimespec(TimevalToNsec(tv[0])),
		NsecToTimespec(TimevalToNsec(tv[1])),
	}
	return utimensat(dirfd, path, (*[2]Timespec)(unsafe.Pointer(&ts[0])), 0)
}
 Time(t *Time_t) (Time_t, error) {
	var tv Timeval
	err := Gettimeofday(&tv)
	if err != nil {
		return 0, err
	}
	if t != nil {
		*t = Time_t(tv.Sec)
	}
	return Time_t(tv.Sec), nil
}
 Utime(path string, buf *Utimbuf) error {
	tv := []Timeval{
		{Sec: buf.Actime},
		{Sec: buf.Modtime},
	}
	return Utimes(path, tv)
}
 utimes(path string, tv *[2]Timeval) (err error) {
	if tv == nil {
		return utimensat(AT_FDCWD, path, nil, 0)
	}	ts := []Timespec{
		NsecToTimespec(TimevalToNsec(tv[0])),
		NsecToTimespec(TimevalToNsec(tv[1])),
	}
	return utimensat(AT_FDCWD, path, (*[2]Timespec)(unsafe.Pointer(&ts[0])), 0)
}// Getrlimit prefers the prlimit64 system call. See issue 38604. Getrlimit(resource int, rlim *Rlimit) error {
	err := Prlimit(0, resource, nil, rlim)
	if err != ENOSYS {
		return err
	}
	return getrlimit(resource, rlim)
}
 (r *PtraceRegs) PC() uint64 { return r.Pc }
 (r *PtraceRegs) SetPC(pc uint64) { r.Pc = pc }
 (iov *Iovec) SetLen(length int) {
	iov.Len = uint64(length)
}
 (msghdr *Msghdr) SetControllen(length int) {
	msghdr.Controllen = uint64(length)
}
 (msghdr *Msghdr) SetIovlen(length int) {
	msghdr.Iovlen = uint64(length)
}
 (cmsg *Cmsghdr) SetLen(length int) {
	cmsg.Len = uint64(length)
}
 (rsa *RawSockaddrNFCLLCP) SetServiceNameLen(length int) {
	rsa.Service_name_len = uint64(length)
}
 Pause() error {
	_, err := ppoll(nil, 0, nil, nil)
	return err
}//sys	kexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
 KexecFileLoad(kernelFd int, initrdFd int, cmdline string, flags int) error {
	cmdlineLen := len(cmdline)
	if cmdlineLen > 0 {
		// Account for the additional NULL byte added by
		// BytePtrFromString in kexecFileLoad. The kexec_file_load
		// syscall expects a NULL-terminated string.
		cmdlineLen++
	}
	return kexecFileLoad(kernelFd, initrdFd, cmdlineLen, cmdline, flags)
}
