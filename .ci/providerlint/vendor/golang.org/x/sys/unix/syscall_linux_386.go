// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build 386 && linux
// +build 386,linuxpackage uniximport (
"unsafe"
)
 setTimespec(sec, nsec int64) Timespec {
return Timespec{Sec: int32(sec), Nsec: int32(nsec)}
}
 setTimeval(sec, usec int64) Timeval {
return Timeval{Sec: int32(sec), Usec: int32(usec)}
}// 64-bit file system and 32-bit uid calls
// (386 default is 32-bit file system and 16-bit uid).
//sysEpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)
//sysFadvise(fd int, offset int64, length int64, advice int) (err error) = SYS_FADVISE64_64
//sysFchown(fd int, uid int, gid int) (err error) = SYS_FCHOWN32
//sysFstat(fd int, stat *Stat_t) (err error) = SYS_FSTAT64
//sysFstatat(dirfd int, path string, stat *Stat_t, flags int) (err error) = SYS_FSTATAT64
//sysFtruncate(fd int, length int64) (err error) = SYS_FTRUNCATE64
//sysnbGetegid() (egid int) = SYS_GETEGID32
//sysnbGeteuid() (euid int) = SYS_GETEUID32
//sysnbGetgid() (gid int) = SYS_GETGID32
//sysnbGetuid() (uid int) = SYS_GETUID32
//sysIoperm(from int, num int, on int) (err error)
//sysIopl(level int) (err error)
//sysLchown(path string, uid int, gid int) (err error) = SYS_LCHOWN32
//sysLstat(path string, stat *Stat_t) (err error) = SYS_LSTAT64
//syspread(fd int, p []byte, offset int64) (n int, err error) = SYS_PREAD64
//syspwrite(fd int, p []byte, offset int64) (n int, err error) = SYS_PWRITE64
//sysRenameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//syssendfile(outfd int, infd int, offset *int64, count int) (written int, err error) = SYS_SENDFILE64
//syssetfsgid(gid int) (prev int, err error) = SYS_SETFSGID32
//syssetfsuid(uid int) (prev int, err error) = SYS_SETFSUID32
//sysSplice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error)
//sysStat(path string, stat *Stat_t) (err error) = SYS_STAT64
//sysSyncFileRange(fd int, off int64, n int64, flags int) (err error)
//sysTruncate(path string, length int64) (err error) = SYS_TRUNCATE64
//sysUstat(dev int, ubuf *Ustat_t) (err error)
//sysnbgetgroups(n int, list *_Gid_t) (nn int, err error) = SYS_GETGROUPS32
//sysnbsetgroups(n int, list *_Gid_t) (err error) = SYS_SETGROUPS32
//sysSelect(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error) = SYS__NEWSELECT//sysmmap2(addr uintptr, length uintptr, prot int, flags int, fd int, pageOffset uintptr) (xaddr uintptr, err error)
sPause() (err error)
 mmap(addr uintptr, length uintptr, prot int, flags int, fd int, offset int64) (xaddr uintptr, err error) {
page := uintptr(offset / 4096)
if offset != int64(page)*4096 {
return 0, EINVAL
}
return mmap2(addr, length, prot, flags, fd, page)
}type rlimit32 struct {
Cur uint32
Max uint32
}//sysnbgetrlimit(resource int, rlim *rlimit32) (err error) = SYS_GETRLIMITt rlimInf32 = ^uint32(0)
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
rlim.Max = rlimInf64
} else {
rlim.Max = uint64(rl.Max)return
}
 Seek(fd int, offset int64, whence int) (newoffset int64, err error) {
newoffset, errno := seek(fd, offset, whence)
if errno != 0 {
return 0, errno
}
return newoffset, nil
}//sysfutimesat(dirfd int, path string, times *[2]Timeval) (err error)
//sysnbGettimeofday(tv *Timeval) (err error)
//sysnbTime(t *Time_t) (tt Time_t, err error)
//sysUtime(path string, buf *Utimbuf) (err error)
//sysutimes(path string, times *[2]Timeval) (err error)// On x86 Linux, all the socket calls go through an extra indirection,
// I think because the 5-register system call interface can't handle
// the 6-argument calls like sendto and recvfrom. Instead the
// arguments to the underlying system call are the number below
// and a pointer to an array of uintptr. We hide the pointer in the
// socketcall assembly to avoid allocation on every system call.const (
// see linux/net.h
_SOCKET      = 1
_BIND        = 2
_CONNECT     = 3
_LISTEN      = 4
_ACCEPT      = 5
_GETSOCKNAME = 6
_GETPEERNAME = 7
_SOCKETPAIR  = 8
_SEND        = 9
_RECV        = 10
_SENDTO      = 11
_RECVFROM    = 12
_SHUTDOWN    = 13
_SETSOCKOPT  = 14
_GETSOCKOPT  = 15
_SENDMSG     = 16
_RECVMSG     = 17
CEPT4     = 18
_RECVMMSG    = 19
_SENDMMSG    = 20
)
 accept4(s int, rsa *RawSockaddrAny, addrlen *_Socklen, flags int) (fd int, err error) {
fd, e := socketcall(_ACCEPT4, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), uintptr(flags), 0, 0)
e != 0 {
err = e
}
return
}
 getsockname(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
e := rawsocketcall(_GETSOCKNAME, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
if e != 0 {
err = e
}
return
}
peername(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) {
_, e := rawsocketcall(_GETPEERNAME, uintptr(s), uintptr(unsafe.Pointer(rsa)), uintptr(unsafe.Pointer(addrlen)), 0, 0, 0)
if e != 0 {
err = e
}
return
}
 socketpair(domain int, typ int, flags int, fd *[2]int32) (err error) {
_, e := rawsocketcall(_SOCKETPAIR, uintptr(domain), uintptr(typ), uintptr(flags), uintptr(unsafe.Pointer(fd)), 0, 0)
if e != 0 {
err = e
}
return
}
 bind(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
_, e := socketcall(_BIND, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
if e != 0 {
err = e
}
return connect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) {
_, e := socketcall(_CONNECT, uintptr(s), uintptr(addr), uintptr(addrlen), 0, 0, 0)
if e != 0 {
err = e
}
urn
}
 socket(domain int, typ int, proto int) (fd int, err error) {
fd, e := rawsocketcall(_SOCKET, uintptr(domain), uintptr(typ), uintptr(proto), 0, 0, 0)
if e != 0 {
err = ereturn
}
 getsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) {
_, e := socketcall(_GETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), uintptr(unsafe.Pointer(vallen)), 0)
if e != 0 {
r = e
}
return
}
 setsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) {
_, e := socketcall(_SETSOCKOPT, uintptr(s), uintptr(level), uintptr(name), uintptr(val), vallen, 0)
if e != 0 {
err = e
}
return recvfrom(s int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) {
var base uintptr
if len(p) > 0 {
base = uintptr(unsafe.Pointer(&p[0]))
}
n, e := socketcall(_RECVFROM, uintptr(s), base, uintptr(len(p)), uintptr(flags), uintptr(unsafe.Pointer(from)), uintptr(unsafe.Pointer(fromlen)))
if e != 0 {
err = e
}
urn
}
 sendto(s int, p []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) {
var base uintptr
if len(p) > 0 {
base = uintptr(unsafe.Pointer(&p[0]))_, e := socketcall(_SENDTO, uintptr(s), base, uintptr(len(p)), uintptr(flags), uintptr(to), uintptr(addrlen))
if e != 0 {
err = e
}
return
}
 recvmsg(s int, msg *Msghdr, flags int) (n int, err error) {
n, e := socketcall(_RECVMSG, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
if e != 0 {
err = e
}
return
}
 sendmsg(s int, msg *Msghdr, flags int) (n int, err error) {
n, e := socketcall(_SENDMSG, uintptr(s), uintptr(unsafe.Pointer(msg)), uintptr(flags), 0, 0, 0)
if e != 0 {
err = e
}
return Listen(s int, n int) (err error) {
_, e := socketcall(_LISTEN, uintptr(s), uintptr(n), 0, 0, 0, 0)
if e != 0 {
err = e
}
urn
}
 Shutdown(s, how int) (err error) {
_, e := socketcall(_SHUTDOWN, uintptr(s), uintptr(how), 0, 0, 0, 0)
if e != 0 {
err = e
}
return
}
 Fstatfs(fd int, buf *Statfs_t) (err error) {
_, e := Syscall(SYS_FSTATFS64, uintptr(fd), unsafe.Sizeof(*buf), uintptr(unsafe.Pointer(buf)))
if e != 0 {
r = e
}
return
}
 Statfs(path string, buf *Statfs_t) (err error) {
pathp, err := BytePtrFromString(path)
err != nil {
return err
}
_, _, e := Syscall(SYS_STATFS64, uintptr(unsafe.Pointer(pathp)), unsafe.Sizeof(*buf), uintptr(unsafe.Pointer(buf)))
e != 0 {
err = e
}
return (r *PtraceRegs) PC() uint64 { return uint64(uint32(r.Eip)) }
 (r *PtraceRegs) SetPC(pc uint64) { r.Eip = int32(pc) }
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
}
 (rsa *RawSockaddrNFCLLCP) SetServiceNameLen(length int) {
rsa.Service_name_len = uint32(length)
}
