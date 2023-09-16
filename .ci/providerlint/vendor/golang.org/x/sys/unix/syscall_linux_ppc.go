// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build linux && ppc
// +build linux,ppcpackage uniximport (
"syscall"
"unsafe"
)//sysEpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sysFchown(fd int, uid int, gid int) (err error)
//sysFstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
//sysFstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sysFtruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64
//sysnbGetegid() (egid int)
//sysnbGeteuid() (euid int)
//sysnbGetgid() (gid int)
//sysnbGetuid() (uid int)
//sysIoperm(from int, num int, on int) (err error)
//sysIopl(level int) (err error)
//sysLchown(path string, uid int, gid int) (err error)
//sysListen(s int, n int) (err error)
//sysLstat(path string, stat *Stat_t) (err error) = SYS_LSTAT64
//sysPause() (err error)
//syspread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//syspwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sysRenameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sysSelect(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) = SYS__NEWSELECT
//syssendfile(outfd int, infd int, offset *int64, count int) (written int, err error) = SYS_SENDFILE64
//syssetfsgid(gid int) (prev int, err error)
//syssetfsuid(uid int) (prev int, err error)
//sysShutdown(fd int, how int) (err error)
//sysSplice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error)
//sysStat(path string, stat *Stat_t) (err error) = SYS_STAT64
//sysTruncate(path string, length int64) (err error) = SYS_TRUNCATE64
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
//syssendmsg(s int, msg *Msghdr, flags int) (n int, err error)//sysfutimesat(dirfd int, path string, times *[2]Timeval) (err error)
//sysnbGettimeofday(tv *Timeval) (err error)
//sysnbTime(t *Time_t) (tt Time_t, err error)
//sysUtime(path string, buf *Utimbuf) (err error)
//sysutimes(path string, times *[2]Timeval) (err error)
 Fadvise(fd int, offset int64, length int64, advice int) (err error) {
_, _, e1 := Syscall6(SYS_FADVISE64_64, uintptr(fd), uintptr(advice), uintptr(offset>>32), uintptr(offset), uintptr(length>>32), uintptr(length))
if e1 != 0 {
err = errnoErr(e1)
}
return
}
 seek(fd int, offset int64, whence int) (int64, syscall.Errno) {
var newoffset int64
offsetLow := uint32(offset & 0xffffffff)
offsetHigh := uint32((offset >> 32) & 0xffffffff)
_, _, err := Syscall6(SYS__LLSEEK, uintptr(fd), uintptr(offsetHigh), uintptr(offsetLow), uintptr(unsafe.Pointer(&newoffset)), uintptr(whence), 0)
return newoffset, err Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
newoffset, errno := seek(fd, offset, whence)
if errno != 0 {
return 0, errno
}
urn newoffset, nil
}
 Fstatfs(fd int, buf *Statfs_t) (err error) {
_, _, e := Syscall(SYS_FSTATFS64, uintptr(fd), unsafe.Sizeof(*buf), uintptr(unsafe.Pointer(buf)))
if e != 0 {
err = ereturn
}
 Statfs(path string, buf *Statfs_t) (err error) {
pathp, err := BytePtrFromString(path)
if err != nil {
return err
}
_, _, e := Syscall(SYS_STATFS64, uintptr(unsafe.Pointer(pathp)), unsafe.Sizeof(*buf), uintptr(unsafe.Pointer(buf)))
if e != 0 {
err = e
}
urn
}//sysmmap2(addr uintptr, length uintptr, prot int, flags int, fd int, pageOffset uintptr) (xaddr uintptr, err error)
 mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error) {
page := uintptr(offset / 4096)
offset != int64(page)*4096 {
return 0, EINVAL
}
return mmap2(addr, length, prot, flags, fd, page) setTimespec(sec, nsec int64) Timespec {
return Timespec{Sec: int32(sec), Nsec: int32(nsec)}
}
 setTimeval(sec, usec int64) Timeval {
return Timeval{Sec: int32(sec), Usec: int32(usec)}
}type rlimit32 struct {
Cur uint32
 uint32
}//sysnbgetrlimit(resource int, rlim *rlimit32) (err error) = SYS_UGETRLIMITconst rlimInf32 = ^uint32(0)
const rlimInf64 = ^uint64(0)
 Getrlimit(resource int, rlim *Rlimit) (err error) {
err = Prlimit(0, resource, nil, rlim)
if err != ENOSYS {
return err
}rl := rlimit32{}
err = getrlimit(resource, &rl)
if err != nil {
return
}if rl.Cur == rlimInf32 {
rlim.Cur = rlimInf64
} else {
rlim.Cur = uint64(rl.Cur)
}if rl.Max == rlimInf32 {
im.Max = rlimInf64
} else {
im.Max = uint64(rl.Max)
}
return
}
 (r *PtraceRegs) PC() uint32 { return r.Nip }
 (r *PtraceRegs) SetPC(pc uint32) { r.Nip = pc }
v *Iovec) SetLen(length int) {
iov.Len = uint32(length)
}
 (msghdr *Msghdr) SetControllen(length int) {
msghdr.Controllen = uint32(length)
}
ghdr *Msghdr) SetIovlen(length int) {
msghdr.Iovlen = uint32(length)
}
 (cmsg *Cmsghdr) SetLen(length int) {
cmsg.Len = uint32(length)
}
 (rsa *RawSockaddrNFCLLCP) SetServiceNameLen(length int) {
rsa.Service_name_len = uint32(length)
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
