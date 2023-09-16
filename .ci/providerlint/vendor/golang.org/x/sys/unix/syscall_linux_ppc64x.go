// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build linux && (ppc64 || ppc64le)
// +build linux
// +build ppc64 ppc64lepackage unix//sysEpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sysFadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64
//sysFchown(fd int, uid int, gid int) (err error)
//sysFstat(fd int, stat *Stat_t) (err error)
//sysFstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_NEWFSTATAT
//sysFstatfs(fd int, buf *Statfs_t) (err error)
//sysFtruncate(fd int, length int64) (err error)
//sysnbGetegid() (egid int)
//sysnbGeteuid() (euid int)
//sysnbGetgid() (gid int)
//sysnbGetrlimit(resource int, rlim *Rlimit) (err error) = SYS_UGETRLIMIT
//sysnbGetuid() (uid int)
//sysIoperm(from int, num int, on int) (err error)
//sysIopl(level int) (err error)
//sysLchown(path string, uid int, gid int) (err error)
//sysListen(s int, n int) (err error)
//sysLstat(path string, stat *Stat_t) (err error)
//sysPause() (err error)
//syspread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//syspwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sysRenameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sysSeek(fd int, offset int64, whence int) (off int64, err error) = SYS_LSEEK
//sysSelect(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) = SYS__NEWSELECT
//syssendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
//syssetfsgid(gid int) (prev int, err error)
//syssetfsuid(uid int) (prev int, err error)
//sysShutdown(fd int, how int) (err error)
//sysSplice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int64, err error)
//sysStat(path string, stat *Stat_t) (err error)
//sysStatfs(path string, buf *Statfs_t) (err error)
//sysTruncate(path string, length int64) (err error)
//sysUstat(dev int, ubuf *Ustat_t) (err error)
//sysaccept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error)
//sysbind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysconnect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error)
//sysnbgetgroups(n int, list *_Gid_t) (nn int, err error)
//sysnbsetgroups(n int, list *_Gid_t) (err error)
//sysgetsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error)
//syssetsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error)
//sysnbsocket(domain int, typ int, proto int) (fd int, err error)
//sysnbsocketpair(domain int, typ int, proto int, fd *[2]int32) (err error)
//sysnbgetpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysnbgetsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error)
//sysrecvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error)
//syssendto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error)
//sysrecvmsg(s int, msg *Msghdr, flags int) (n int, err error)
//syssendmsg(s int, msg *Msghdr, flags int) (n int, err error)
//sysmmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error)//sysfutimesat(dirfd int, path string, times *[2]Timeval) (err error)
//sysnbGettimeofday(tv *Timeval) (err error)
//sysnbTime(t *Time_t) (tt Time_t, err error)
//sysUtime(path string, buf *Utimbuf) (err error)
//sysutimes(path string, times *[2]Timeval) (err error)
 setTimespec(sec, nsec int64) Timespec {
return Timespec{Sec: sec, Nsec: nsec}
}
 setTimeval(sec, usec int64) Timeval {
return Timeval{Sec: sec, Usec: usec} (r *PtraceRegs) PC() uint64 { return r.Nip }
 (r *PtraceRegs) SetPC(pc uint64) { r.Nip = pc }
 (iov *Iovec) SetLen(length int) {
iov.Len = uint64(length)
}
 (msghdr *Msghdr) SetControllen(length int) {
msghdr.Controllen = uint64(length) (msghdr *Msghdr) SetIovlen(length int) {
hdr.Iovlen = uint64(length)
}
 (cmsg *Cmsghdr) SetLen(length int) {
cmsg.Len = uint64(length) (rsa *RawSockaddrNFCLLCP) SetServiceNameLen(length int) {
rsa.Service_name_len = uint64(length)
}//syssyncFileRange2(fd int, flags int, off int64, n int64) (err error) = SYS_SYNC_FILE_RANGE2
 SyncFileRange(fd int, off int64, n int64, flags int) error {
// The sync_file_range and sync_file_range2 syscalls differ only in the
// order of their arguments.
return syncFileRange2(fd, flags, off, n)
}//syskexecFileLoad(kernelFd int, initrdFd int, cmdlineLen int, cmdline string, flags int) (err error)
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
