// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.// Solaris system calls.
// This file is compiled as ordinary Go code,
// but it is also input to mksyscall,
// which parses the //sys lines and generates system call stubs.
// Note that sometimes we use a lowercase //sys name and wrap
// it in our own nicer implementation, either here or in
// syscall_solaris.go or syscall_unix.go.package uniximport (
mt"
s"
untime"
ync"
yscall"
nsafe"
)// Implemented in runtime/syscall_solaris.go.
type syscallFunc uintptr
 rawSysvicall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno) sysvicall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)// SockaddrDatalink implements the Sockaddr interface for AF_LINK type sockets.
type SockaddrDatalink struct {
mily uint16
dex  uint16
peuint8
enuint8
enuint8
enuint8
ta[244]int8
wRawSockaddrDatalink
}
 direntIno(buf []byte) (uint64, bool) {
turn readInt(buf, unsafe.Offsetof(Dirent{}.Ino), unsafe.Sizeof(Dirent{}.Ino))
}
 direntReclen(buf []byte) (uint64, bool) {
turn readInt(buf, unsafe.Offsetof(Dirent{}.Reclen), unsafe.Sizeof(Dirent{}.Reclen))
}
 direntNamlen(buf []byte) (uint64, bool) {
clen, ok := direntReclen(buf)
 !ok {
return 0, falseturn reclen - uint64(unsafe.Offsetof(Dirent{}.Name)), true
}//sysnbpe(p *[2]_C_int) (n int, err error)
 Pipe(p []int) (err error) {
 len(p) != 2 {
return EINVALr pp [2]_C_int
 err := pipe(&pp)
 n != 0 {
return err err == nil {
p[0] = int(pp[0])
p[1] = int(pp[1])turn nil
}//sysnbpe2(p *[2]_C_int, flags int) (err error)
 Pipe2(p []int, flags int) error {
 len(p) != 2 {
return EINVALr pp [2]_C_int
r := pipe2(&pp, flags)
 err == nil {
p[0] = int(pp[0])
p[1] = int(pp[1])turn err
}
 (sa *SockaddrInet4) sockaddr() (unsafe.Pointer, _Socklen, error) {
 sa.Port < 0 || sa.Port > 0xFFFF {
return nil, 0, EINVAL.raw.Family = AF_INET
:= (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
0] = byte(sa.Port >> 8)
1] = byte(sa.Port)
.raw.Addr = sa.Addr
turn unsafe.Pointer(&sa.raw), SizeofSockaddrInet4, nil
}
 (sa *SockaddrInet6) sockaddr() (unsafe.Pointer, _Socklen, error) {
 sa.Port < 0 || sa.Port > 0xFFFF {
return nil, 0, EINVAL.raw.Family = AF_INET6
:= (*[2]byte)(unsafe.Pointer(&sa.raw.Port))
0] = byte(sa.Port >> 8)
1] = byte(sa.Port)
.raw.Scope_id = sa.ZoneId
.raw.Addr = sa.Addr
turn unsafe.Pointer(&sa.raw), SizeofSockaddrInet6, nil
}
 (sa *SockaddrUnix) sockaddr() (unsafe.Pointer, _Socklen, error) {
me := sa.Name
:= len(name)
 n >= len(sa.raw.Path) {
return nil, 0, EINVAL.raw.Family = AF_UNIX
r i := 0; i < n; i++ {
sa.raw.Path[i] = int8(name[i]) length is family (uint16), name, NUL.
 := _Socklen(2)
 n > 0 {
sl += _Socklen(n) + 1 sa.raw.Path[0] == '@' {
sa.raw.Path[0] = 0
// Don't count trailing NUL for abstract address.
sl--
trn unsafe.Pointer(&sa.raw), sl, nil
}//systsockname(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) = libsocket.getsockname
 Getsockname(fd int) (sa Sockaddr, err error) {
r rsa RawSockaddrAny
r len _Socklen = SizeofSockaddrAny
 err = getsockname(fd, &rsa, &len); err != nil {
returnturn anyToSockaddr(fd, &rsa)
}// GetsockoptString returns the string value of the socket option opt for the
// socket associated with fd at the given socket level. GetsockoptString(fd, level, opt int) (string, error) {
f := make([]byte, 256)
llen := _Socklen(len(buf))
r := getsockopt(fd, level, opt, unsafe.Pointer(&buf[0]), &vallen)
 err != nil {
return "", errturn string(buf[:vallen-1]), nil
}const ImplementsGetwd = true//systcwd(buf []byte) (n int, err error)
 Getwd() (wd string, err error) {
r buf [PathMax]byte
 Getcwd will return an error if it failed for any reason.
 err = Getcwd(buf[0:])
 err != nil {
return "", err:= clen(buf[:])
 n < 1 {
return "", EINVALturn string(buf[:n]), nil
}/*
 * Wrapped
 *///sysnbtgroups(ngid int, gid *_Gid_t) (n int, err error)
//sysnbtgroups(ngid int, gid *_Gid_t) (err error)
 Getgroups() (gids []int, err error) {
 err := getgroups(0, nil)
 Check for error and sanity check group count. Newer versions of
 Solaris allow up to 1024 (NGROUPS_MAX).
 n < 0 || n > 1024 {
if err != nil {
return nil, err
}
return nil, EINVAL
else if n == 0 {
return nil, nil
: make([]_Gid_t, n)
 err = getgroups(n, &a[0])
 n == -1 {
return nil, errds = make([]int, n)
r i, v := range a[0:n] {
gids[i] = int(v)turn
}
 Setgroups(gids []int) (err error) {
 len(gids) == 0 {
return setgroups(0, nil)
: make([]_Gid_t, len(gids))
r i, v := range gids {
a[i] = _Gid_t(v)turn setgroups(len(a), &a[0])
}// ReadDirent reads directory entries from fd and writes them into buf. ReadDirent(fd int, buf []byte) (n int, err error) {
 Final argument is (basep *uintptr) and the syscall doesn't take nil.
 TODO(rsc): Can we use a single global basep for all calls?
turn Getdents(fd, buf, new(uintptr))
}// Wait status is 7 bits at bottom, either 0 (exited),
// 0x7F (stopped), or a signal number that caused an exit.
// The 0x80 bit is whether there was a core dump.
// An extra number (exit code, signal causing a stop)
// is in the high bits.type WaitStatus uint32const (
sk  = 0x7F
re  = 0x80
ift = 8eied  = 0
opped = 0x7F
)
 (w WaitStatus) Exited() bool { return w&mask == exited }
 (w WaitStatus) ExitStatus() int {
 w&mask != exited {
return -1turn int(w >> shift)
}
 (w WaitStatus) Signaled() bool { return w&mask != stopped && w&mask != 0 }
 (w WaitStatus) Signal() syscall.Signal {
g := syscall.Signal(w & mask)
 sig == stopped || sig == 0 {
return -1turn sig
}
 (w WaitStatus) CoreDump() bool { return w.Signaled() && w&core != 0 }
 (w WaitStatus) Stopped() bool { return w&mask == stopped && syscall.Signal(w>>shift) != SIGSTOP }
 (w WaitStatus) Continued() bool { return w&mask == stopped && syscall.Signal(w>>shift) == SIGSTOP }
 (w WaitStatus) StopSignal() syscall.Signal {
 !w.Stopped() {
return -1turn syscall.Signal(w>>shift) & 0xFF
}
 (w WaitStatus) TrapCause() int { return -1 }//sysit4(pid int32, statusp *_C_int, options int, rusage *Rusage) (wpid int32, err error)
 Wait4(pid int, wstatus *WaitStatus, options int, rusage *Rusage) (int, error) {
r status _C_int
id, err := wait4(int32(pid), &status, options, rusage)
id := int(rpid)
 wpid == -1 {
return wpid, err wstatus != nil {
*wstatus = WaitStatus(status)turn wpid, nil
}//systhostname(buf []byte) (n int, err error)
 Gethostname() (name string, err error) {
r buf [MaxHostNameLen]byte
 err := gethostname(buf[:])
 n != 0 {
return "", err= clen(buf[:])
 n < 1 {
return "", EFAULTturn string(buf[:n]), nil
}//sysimes(path string, times *[2]Timeval) (err error)
 Utimes(path string, tv []Timeval) (err error) {
 tv == nil {
return utimes(path, nil) len(tv) != 2 {
return EINVALturn utimes(path, (*[2]Timeval)(unsafe.Pointer(&tv[0])))
}//sysimensat(fd int, path string, times *[2]Timespec, flag int) (err error)
 UtimesNano(path string, ts []Timespec) error {
 ts == nil {
return utimensat(AT_FDCWD, path, nil, 0) len(ts) != 2 {
return EINVALturn utimensat(AT_FDCWD, path, (*[2]Timespec)(unsafe.Pointer(&ts[0])), 0)
}
 UtimesNanoAt(dirfd int, path string, ts []Timespec, flags int) error {
 ts == nil {
return utimensat(dirfd, path, nil, flags) len(ts) != 2 {
return EINVALturn utimensat(dirfd, path, (*[2]Timespec)(unsafe.Pointer(&ts[0])), flags)
}//sysntl(fd int, cmd int, arg int) (val int, err error)// FcntlInt performs a fcntl syscall on fd with the provided command and argument. FcntlInt(fd uintptr, cmd, arg int) (int, error) {
lptr, _, errno := sysvicall6(uintptr(unsafe.Pointer(&procfcntl)), 3, uintptr(fd), uintptr(cmd), uintptr(arg), 0, 0, 0)
r err error
 errno != 0 {
err = errnoturn int(valptr), err
}// FcntlFlock performs a fcntl syscall for the F_GETLK, F_SETLK or F_SETLKW command. FcntlFlock(fd uintptr, cmd int, lk *Flock_t) error {
 _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procfcntl)), 3, uintptr(fd), uintptr(cmd), uintptr(unsafe.Pointer(lk)), 0, 0, 0)
 e1 != 0 {
return e1turn nil
}//systimesat(fildes int, path *byte, times *[2]Timeval) (err error)
 Futimesat(dirfd int, path string, tv []Timeval) error {
thp, err := BytePtrFromString(path)
 err != nil {
return err tv == nil {
return futimesat(dirfd, pathp, nil) len(tv) != 2 {
return EINVALturn futimesat(dirfd, pathp, (*[2]Timeval)(unsafe.Pointer(&tv[0])))
}// Solaris doesn't have an futimes function because it allows NULL to be
// specified as the path for futimesat. However, Go doesn't like
// NULL-style string interfaces, so this simple wrapper is provided. Futimes(fd int, tv []Timeval) error {
 tv == nil {
return futimesat(fd, nil, nil) len(tv) != 2 {
return EINVALturn futimesat(fd, nil, (*[2]Timeval)(unsafe.Pointer(&tv[0])))
}
 anyToSockaddr(fd int, rsa *RawSockaddrAny) (Sockaddr, error) {
itch rsa.Addr.Family {
se AF_UNIX:
pp := (*RawSockaddrUnix)(unsafe.Pointer(rsa))
sa := new(SockaddrUnix)
// Assume path ends at NUL.
// This is not technically the Solaris semantics for
// abstract Unix domain sockets -- they are supposed
// to be uninterpreted fixed-size binary blobs -- but
// everyone uses this convention.
n := 0
for n < len(pp.Path) && pp.Path[n] != 0 {
n++
}
sa.Name = string(unsafe.Slice((*byte)(unsafe.Pointer(&pp.Path[0])), n))
return sa, nilse AF_INET:
pp := (*RawSockaddrInet4)(unsafe.Pointer(rsa))
sa := new(SockaddrInet4)
p := (*[2]byte)(unsafe.Pointer(&pp.Port))
sa.Port = int(p[0])<<8 + int(p[1])
sa.Addr = pp.Addr
return sa, nilse AF_INET6:
pp := (*RawSockaddrInet6)(unsafe.Pointer(rsa))
sa := new(SockaddrInet6)
p := (*[2]byte)(unsafe.Pointer(&pp.Port))
sa.Port = int(p[0])<<8 + int(p[1])
sa.ZoneId = pp.Scope_id
sa.Addr = pp.Addr
return sa, nilturn nil, EAFNOSUPPORT
}//syscept(s int, rsa *RawSockaddrAny, addrlen *_Socklen) (fd int, err error) = libsocket.accept
 Accept(fd int) (nfd int, sa Sockaddr, err error) {
r rsa RawSockaddrAny
r len _Socklen = SizeofSockaddrAny
d, err = accept(fd, &rsa, &len)
 nfd == -1 {
return, err = anyToSockaddr(fd, &rsa)
 err != nil {
Close(nfd)
nfd = 0turn
}//syscvmsg(s int, msg *Msghdr, flags int) (n int, err error) = libsocket.__xnet_recvmsg
 recvmsgRaw(fd int, iov []Iovec, oob []byte, flags int, rsa *RawSockaddrAny) (n, oobn int, recvflags int, err error) {
r msg Msghdr
g.Name = (*byte)(unsafe.Pointer(rsa))
g.Namelen = uint32(SizeofSockaddrAny)
r dummy byte
 len(oob) > 0 {
// receive at least one normal byte
if emptyIovecs(iov) {
var iova [1]Iovec
iova[0].Base = &dummy
iova[0].SetLen(1)
iov = iova[:]
}
msg.Accrightslen = int32(len(oob)) len(iov) > 0 {
msg.Iov = &iov[0]
msg.SetIovlen(len(iov)) n, err = recvmsg(fd, &msg, flags); n == -1 {
returnbn = int(msg.Accrightslen)
turn
}//sysndmsg(s int, msg *Msghdr, flags int) (n int, err error) = libsocket.__xnet_sendmsg
 sendmsgN(fd int, iov []Iovec, oob []byte, ptr unsafe.Pointer, salen _Socklen, flags int) (n int, err error) {
r msg Msghdr
g.Name = (*byte)(unsafe.Pointer(ptr))
g.Namelen = uint32(salen)
r dummy byte
r empty bool
 len(oob) > 0 {
// send at least one normal byte
empty = emptyIovecs(iov)
if empty {
var iova [1]Iovec
iova[0].Base = &dummy
iova[0].SetLen(1)
iov = iova[:]
}
msg.Accrightslen = int32(len(oob)) len(iov) > 0 {
msg.Iov = &iov[0]
msg.SetIovlen(len(iov)) n, err = sendmsg(fd, &msg, flags); err != nil {
return 0, err len(oob) > 0 && empty {
n = 0turn n, nil
}//sysct(path *byte) (err error)
 Acct(path string) (err error) {
 len(path) == 0 {
// Assume caller wants to disable accounting.
return acct(nil)
tp, err := BytePtrFromString(path)
 err != nil {
return errturn acct(pathp)
}//sysmakedev(version int, major uint, minor uint) (val uint64)
 Mkdev(major, minor uint32) uint64 {
turn __makedev(NEWDEV, uint(major), uint(minor))
}//sysmajor(version int, dev uint64) (val uint)
 Major(dev uint64) uint32 {
turn uint32(__major(NEWDEV, dev))
}//sysminor(version int, dev uint64) (val uint)
 Minor(dev uint64) uint32 {
turn uint32(__minor(NEWDEV, dev))
}/*
 * Expose the ioctl function
 *///sysctlRet(fd int, req int, arg uintptr) (ret int, err error) = libc.ioctl
//sysctlPtrRet(fd int, req int, arg unsafe.Pointer) (ret int, err error) = libc.ioctl
 ioctl(fd int, req int, arg uintptr) (err error) {
 err = ioctlRet(fd, req, arg)
turn err
}
 ioctlPtr(fd int, req int, arg unsafe.Pointer) (err error) {
 err = ioctlPtrRet(fd, req, arg)
turn err
}
 IoctlSetTermio(fd int, req int, value *Termio) error {
turn ioctlPtr(fd, req, unsafe.Pointer(value))
}
 IoctlGetTermio(fd int, req int) (*Termio, error) {
r value Termio
r := ioctlPtr(fd, req, unsafe.Pointer(&value))
turn &value, err
}//sysll(fds *PollFd, nfds int, timeout int) (n int, err error)
 Poll(fds []PollFd, timeout int) (n int, err error) {
 len(fds) == 0 {
return poll(nil, 0, timeout)turn poll(&fds[0], len(fds), timeout)
}
 Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) {
 raceenabled {
raceReleaseMerge(unsafe.Pointer(&ioSync))turn sendfile(outfd, infd, offset, count)
}/*
 * Exposed directly
 */
//syscess(path string, mode uint32) (err error)
//sysjtime(delta *Timeval, olddelta *Timeval) (err error)
//sysdir(path string) (err error)
//sysmod(path string, mode uint32) (err error)
//sysown(path string, uid int, gid int) (err error)
//sysroot(path string) (err error)
//sysockGettime(clockid int32, time *Timespec) (err error)
//sysose(fd int) (err error)
//syseat(path string, mode uint32) (fd int, err error)
//sysp(fd int) (nfd int, err error)
//sysp2(oldfd int, newfd int) (err error)
//sysit(code int)
//sysccessat(dirfd int, path string, mode uint32, flags int) (err error)
//syshdir(fd int) (err error)
//syshmod(fd int, mode uint32) (err error)
//syshmodat(dirfd int, path string, mode uint32, flags int) (err error)
//syshown(fd int, uid int, gid int) (err error)
//syshownat(dirfd int, path string, uid int, gid int, flags int) (err error)
//sysatasync(fd int) (err error)
//sysock(fd int, how int) (err error)
//sysathconf(fd int, name int) (val int, err error)
//systat(fd int, stat *Stat_t) (err error)
//systatat(fd int, path string, stat *Stat_t, flags int) (err error)
//systatvfs(fd int, vfsstat *Statvfs_t) (err error)
//systdents(fd int, buf []byte, basep *uintptr) (n int, err error)
//sysnbtgid() (gid int)
//sysnbtpid() (pid int)
//sysnbtpgid(pid int) (pgid int, err error)
//sysnbtpgrp() (pgid int, err error)
//systeuid() (euid int)
//systegid() (egid int)
//systppid() (ppid int)
//systpriority(which int, who int) (n int, err error)
//sysnbtrlimit(which int, lim *Rlimit) (err error)
//sysnbtrusage(who int, rusage *Rusage) (err error)
//sysnbtsid(pid int) (sid int, err error)
//sysnbttimeofday(tv *Timeval) (err error)
//sysnbtuid() (uid int)
//sysll(pid int, signum syscall.Signal) (err error)
//syshown(path string, uid int, gid int) (err error)
//sysnk(path string, link string) (err error)
//syssten(s int, backlog int) (err error) = libsocket.__xnet_llisten
//systat(path string, stat *Stat_t) (err error)
//sysdvise(b []byte, advice int) (err error)
//sysdir(path string, mode uint32) (err error)
//sysdirat(dirfd int, path string, mode uint32) (err error)
//sysfifo(path string, mode uint32) (err error)
//sysfifoat(dirfd int, path string, mode uint32) (err error)
//sysnod(path string, mode uint32, dev int) (err error)
//sysnodat(dirfd int, path string, mode uint32, dev int) (err error)
//sysock(b []byte) (err error)
//sysockall(flags int) (err error)
//sysrotect(b []byte, prot int) (err error)
//sysync(b []byte, flags int) (err error)
//sysnlock(b []byte) (err error)
//sysnlockall() (err error)
//sysnosleep(time *Timespec, leftover *Timespec) (err error)
//sysen(path string, mode int, perm uint32) (fd int, err error)
//sysenat(dirfd int, path string, flags int, mode uint32) (fd int, err error)
//systhconf(path string, name int) (val int, err error)
//sysuse() (err error)
//sysead(fd int, p []byte, offset int64) (n int, err error)
//sysrite(fd int, p []byte, offset int64) (n int, err error)
//sysad(fd int, p []byte) (n int, err error)
//sysadlink(path string, buf []byte) (n int, err error)
//sysname(from string, to string) (err error)
//sysnameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)
//sysdir(path string) (err error)
//sysek(fd int, offset int64, whence int) (newoffset int64, err error) = lseek
//syslect(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error)
//sysnbtegid(egid int) (err error)
//sysnbteuid(euid int) (err error)
//sysnbtgid(gid int) (err error)
//systhostname(p []byte) (err error)
//sysnbtpgid(pid int, pgid int) (err error)
//systpriority(which int, who int, prio int) (err error)
//sysnbtregid(rgid int, egid int) (err error)
//sysnbtreuid(ruid int, euid int) (err error)
//sysnbtsid() (pid int, err error)
//sysnbtuid(uid int) (err error)
//sysutdown(s int, how int) (err error) = libsocket.shutdown
//sysat(path string, stat *Stat_t) (err error)
//sysatvfs(path string, vfsstat *Statvfs_t) (err error)
//sysmlink(path string, link string) (err error)
//sysnc() (err error)
//syssconf(which int) (n int64, err error)
//sysnbmes(tms *Tms) (ticks uintptr, err error)
//sysuncate(path string, length int64) (err error)
//sysync(fd int) (err error)
//sysruncate(fd int, length int64) (err error)
//sysask(mask int) (oldmask int)
//sysnbame(buf *Utsname) (err error)
//sysmount(target string, flags int) (err error) = libc.umount
//syslink(path string) (err error)
//syslinkat(dirfd int, path string, flags int) (err error)
//systat(dev int, ubuf *Ustat_t) (err error)
//sysime(path string, buf *Utimbuf) (err error)
//sysnd(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) = libsocket.__xnet_bind
//sysnnect(s int, addr unsafe.Pointer, addrlen _Socklen) (err error) = libsocket.__xnet_connect
//sysap(addr uintptr, length uintptr, prot int, flag int, fd int, pos int64) (ret uintptr, err error)
//sysnmap(addr uintptr, length uintptr) (err error)
//sysndfile(outfd int, infd int, offset *int64, count int) (written int, err error) = libsendfile.sendfile
//sysndto(s int, buf []byte, flags int, to unsafe.Pointer, addrlen _Socklen) (err error) = libsocket.__xnet_sendto
//syscket(domain int, typ int, proto int) (fd int, err error) = libsocket.__xnet_socket
//sysnbcketpair(domain int, typ int, proto int, fd *[2]int32) (err error) = libsocket.__xnet_socketpair
//sysite(fd int, p []byte) (n int, err error)
//systsockopt(s int, level int, name int, val unsafe.Pointer, vallen *_Socklen) (err error) = libsocket.__xnet_getsockopt
//sysnbtpeername(fd int, rsa *RawSockaddrAny, addrlen *_Socklen) (err error) = libsocket.getpeername
//systsockopt(s int, level int, name int, val unsafe.Pointer, vallen uintptr) (err error) = libsocket.setsockopt
//syscvfrom(fd int, p []byte, flags int, from *RawSockaddrAny, fromlen *_Socklen) (n int, err error) = libsocket.recvfrom
 readlen(fd int, buf *byte, nbuf int) (n int, err error) {
, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procread)), 3, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf), 0, 0, 0)
= int(r0)
 e1 != 0 {
err = e1turn
}
 writelen(fd int, buf *byte, nbuf int) (n int, err error) {
, _, e1 := sysvicall6(uintptr(unsafe.Pointer(&procwrite)), 3, uintptr(fd), uintptr(unsafe.Pointer(buf)), uintptr(nbuf), 0, 0, 0)
= int(r0)
 e1 != 0 {
err = e1turn
}// Event Portstype fileObjCookie struct {
bj*fileObj
okie interface{}
}// EventPort provides a safe abstraction on top of Solaris/illumos Event Ports.
type EventPort struct {
rt  int
sync.Mutex
smap[uintptr]*fileObjCookie
ths map[string]*fileObjCookie
 The user cookie presents an interesting challenge from a memory management perspective.
 There are two paths by which we can discover that it is no longer in use:
 1. The user calls port_dissociate before any events fire
 2. An event fires and we return it to the user
 The tricky situation is if the event has fired in the kernel but
 the user hasn't requested/received it yet.
 If the user wants to port_dissociate before the event has been processed,
 we should handle things gracefully. To do so, we need to keep an extra
 reference to the cookie around until the event is processed
 thus the otherwise seemingly extraneous "cookies" map
 The key of this map is a pointer to the corresponding fCookie
okies map[*fileObjCookie]struct{}
}// PortEvent is an abstraction of the port_event C struct.
// Compare Source against PORT_SOURCE_FILE or PORT_SOURCE_FD
// to see if Path or Fd was the event source. The other will be
// uninitialized.
type PortEvent struct {
okie interface{}
ents int32
uintptr
thstring
urce uint16
bj*fileObj
}// NewEventPort creates a new EventPort including the
// underlying call to port_create(3c). NewEventPort() (*EventPort, error) {
rt, err := port_create()
 err != nil {
return nil, err:= &EventPort{
port:port,
fds:make(map[uintptr]*fileObjCookie),
paths:make(map[string]*fileObjCookie),
cookies: make(map[*fileObjCookie]struct{}),turn e, nil
}//sysrt_create() (n int, err error)
//sysrt_associate(port int, source int, object uintptr, events int, user *byte) (n int, err error)
//sysrt_dissociate(port int, source int, object uintptr) (n int, err error)
//sysrt_get(port int, pe *portEvent, timeout *Timespec) (n int, err error)
//sysrt_getn(port int, pe *portEvent, max uint32, nget *uint32, timeout *Timespec) (n int, err error)// Close closes the event port. (e *EventPort) Close() error {
mu.Lock()
fer e.mu.Unlock()
r := Close(e.port)
 err != nil {
return errfds = nil
paths = nil
cookies = nil
turn nil
}// PathIsWatched checks to see if path is associated with this EventPort. (e *EventPort) PathIsWatched(path string) bool {
mu.Lock()
fer e.mu.Unlock()
 found := e.paths[path]
turn found
}// FdIsWatched checks to see if fd is associated with this EventPort. (e *EventPort) FdIsWatched(fd uintptr) bool {
mu.Lock()
fer e.mu.Unlock()
 found := e.fds[fd]
turn found
}// AssociatePath wraps port_associate(3c) for a filesystem path including
// creating the necessary file_obj from the provided stat information. (e *EventPort) AssociatePath(path string, stat os.FileInfo, events int, cookie interface{}) error {
mu.Lock()
fer e.mu.Unlock()
 _, found := e.paths[path]; found {
return fmt.Errorf("%v is already associated with this Event Port", path)ookie, err := createFileObjCookie(path, stat, cookie)
 err != nil {
return err err = port_associate(e.port, PORT_SOURCE_FILE, uintptr(unsafe.Pointer(fCookie.fobj)), events, (*byte)(unsafe.Pointer(fCookie)))
 err != nil {
return errpaths[path] = fCookie
cookies[fCookie] = struct{}{}
turn nil
}// DissociatePath wraps port_dissociate(3c) for a filesystem path. (e *EventPort) DissociatePath(path string) error {
mu.Lock()
fer e.mu.Unlock()
 ok := e.paths[path]
 !ok {
return fmt.Errorf("%v is not associated with this Event Port", path) err := port_dissociate(e.port, PORT_SOURCE_FILE, uintptr(unsafe.Pointer(f.fobj)))
 If the path is no longer associated with this event port (ENOENT)
 we should delete it from our map. We can still return ENOENT to the caller.
 But we need to save the cookie
 err != nil && err != ENOENT {
return err err == nil {
// dissociate was successful, safe to delete the cookie
fCookie := e.paths[path]
delete(e.cookies, fCookie)lete(e.paths, path)
turn err
}// AssociateFd wraps calls to port_associate(3c) on file descriptors. (e *EventPort) AssociateFd(fd uintptr, events int, cookie interface{}) error {
mu.Lock()
fer e.mu.Unlock()
 _, found := e.fds[fd]; found {
return fmt.Errorf("%v is already associated with this Event Port", fd)ookie, err := createFileObjCookie("", nil, cookie)
 err != nil {
return err err = port_associate(e.port, PORT_SOURCE_FD, fd, events, (*byte)(unsafe.Pointer(fCookie)))
 err != nil {
return errfds[fd] = fCookie
cookies[fCookie] = struct{}{}
turn nil
}// DissociateFd wraps calls to port_dissociate(3c) on file descriptors. (e *EventPort) DissociateFd(fd uintptr) error {
mu.Lock()
fer e.mu.Unlock()
 ok := e.fds[fd]
 !ok {
return fmt.Errorf("%v is not associated with this Event Port", fd) err := port_dissociate(e.port, PORT_SOURCE_FD, fd)
 err != nil && err != ENOENT {
return err err == nil {
// dissociate was successful, safe to delete the cookie
fCookie := e.fds[fd]
delete(e.cookies, fCookie)lete(e.fds, fd)
turn err
}
 createFileObjCookie(name string, stat os.FileInfo, cookie interface{}) (*fileObjCookie, error) {
ookie := new(fileObjCookie)
ookie.cookie = cookie
 name != "" && stat != nil {
fCookie.fobj = new(fileObj)
bs, err := ByteSliceFromString(name)
if err != nil {
return nil, err
}
fCookie.fobj.Name = (*int8)(unsafe.Pointer(&bs[0]))
s := stat.Sys().(*syscall.Stat_t)
fCookie.fobj.Atim.Sec = s.Atim.Sec
fCookie.fobj.Atim.Nsec = s.Atim.Nsec
fCookie.fobj.Mtim.Sec = s.Mtim.Sec
fCookie.fobj.Mtim.Nsec = s.Mtim.Nsec
fCookie.fobj.Ctim.Sec = s.Ctim.Sec
fCookie.fobj.Ctim.Nsec = s.Ctim.Nsecturn fCookie, nil
}// GetOne wraps port_get(3c) and returns a single PortEvent. (e *EventPort) GetOne(t *Timespec) (*PortEvent, error) {
 := new(portEvent)
 err := port_get(e.port, pe, t)
 err != nil {
return nil, err:= new(PortEvent)
mu.Lock()
fer e.mu.Unlock()
r = e.peIntToExt(pe, p)
 err != nil {
return nil, errturn p, nil
}// peIntToExt converts a cgo portEvent struct into the friendlier PortEvent
// NOTE: Always call this function while holding the e.mu mutex (e *EventPort) peIntToExt(peInt *portEvent, peExt *PortEvent) error {
 e.cookies == nil {
return fmt.Errorf("this EventPort is already closed")Ext.Events = peInt.Events
Ext.Source = peInt.Source
ookie := (*fileObjCookie)(unsafe.Pointer(peInt.User))
 found := e.cookies[fCookie]i found {
panic("unexpected event port address; may be due to kernel bug; see https://go.dev/issue/54254")Ext.Cookie = fCookie.cookie
lete(e.cookies, fCookie)sich peInt.Source {
se PORT_SOURCE_FD:
peExt.Fd = uintptr(peInt.Object)
// Only remove the fds entry if it exists and this cookie matches
if fobj, ok := e.fds[peExt.Fd]; ok {
if fobj == fCookie {
delete(e.fds, peExt.Fd)
}
}
se PORT_SOURCE_FILE:
peExt.fobj = fCookie.fobj
peExt.Path = BytePtrToString((*byte)(unsafe.Pointer(peExt.fobj.Name)))
// Only remove the paths entry if it exists and this cookie matches
if fobj, ok := e.paths[peExt.Path]; ok {
if fobj == fCookie {
delete(e.paths, peExt.Path)
}
}turn nil
}// Pending wraps port_getn(3c) and returns how many events are pending. (e *EventPort) Pending() (int, error) {
r n uint32 = 0
 err := port_getn(e.port, nil, 0, &n, nil)
turn int(n), err
}// Get wraps port_getn(3c) and fills a slice of PortEvent.
// It will block until either min events have been received
// or the timeout has been exceeded. It will return how many
// events were actually received along with any error information. (e *EventPort) Get(s []PortEvent, min int, timeout *Timespec) (int, error) {
 min == 0 {
return 0, fmt.Errorf("need to request at least one event or use Pending() instead") len(s) < min {
return 0, fmt.Errorf("len(s) (%d) is less than min events requested (%d)", len(s), min)t := uint32(min)
x := uint32(len(s))
r err error
 := make([]portEvent, max)
 err = port_getn(e.port, &ps[0], max, &got, timeout)
 got will be trustworthy with ETIME, but not any other error.
 err != nil && err != ETIME {
return 0, errmu.Lock()
fer e.mu.Unlock()
lid := 0
r i := 0; i < int(got); i++ {
err2 := e.peIntToExt(&ps[i], &s[i])
if err2 != nil {
if valid == 0 && err == nil {
// If err2 is the only error and there are no valid events
// to return, return it to the caller.
err = err2
}
break
}
valid = i + 1turn valid, err
}//systmsg(fd int, clptr *strbuf, dataptr *strbuf, flags int) (err error)
 Putmsg(fd int, cl []byte, data []byte, flags int) (err error) {
r clp, datap *strbuf
 len(cl) > 0 {
clp = &strbuf{
Len: int32(len(cl)),
Buf: (*int8)(unsafe.Pointer(&cl[0])),
} len(data) > 0 {
datap = &strbuf{
Len: int32(len(data)),
Buf: (*int8)(unsafe.Pointer(&data[0])),
}turn putmsg(fd, clp, datap, flags)
}//systmsg(fd int, clptr *strbuf, dataptr *strbuf, flags *int) (err error)
 Getmsg(fd int, cl []byte, data []byte) (retCl []byte, retData []byte, flags int, err error) {
r clp, datap *strbuf
 len(cl) > 0 {
clp = &strbuf{
Maxlen: int32(len(cl)),
Buf:(*int8)(unsafe.Pointer(&cl[0])),
} len(data) > 0 {
datap = &strbuf{
Maxlen: int32(len(data)),
Buf:(*int8)(unsafe.Pointer(&data[0])),
}
 rr = getmsg(fd, clp, datap, &flags); err != nil {
return nil, nil, 0, err
 en(cl) > 0 {
retCl = cl[:clp.Len] len(data) > 0 {
retData = data[:datap.Len]turn retCl, retData, flags, nil
}
 IoctlSetIntRetInt(fd int, req int, arg int) (int, error) {
turn ioctlRet(fd, req, uintptr(arg))
}
 IoctlSetString(fd int, req int, val string) error {
 := make([]byte, len(val)+1)
py(bs[:len(bs)-1], val)
r := ioctlPtr(fd, req, unsafe.Pointer(&bs[0]))
ntime.KeepAlive(&bs[0])
turn err
}// Lifreq Helpers
 (l *Lifreq) SetName(name string) error {
 len(name) >= len(l.Name) {
return fmt.Errorf("name cannot be more than %d characters", len(l.Name)-1)r i := range name {
l.Name[i] = int8(name[i])turn nil
}
 (l *Lifreq) SetLifruInt(d int) {
*int)(unsafe.Pointer(&l.Lifru[0])) = d
}
 (l *Lifreq) GetLifruInt() int {
turn *(*int)(unsafe.Pointer(&l.Lifru[0]))
}
 (l *Lifreq) SetLifruUint(d uint) {
*uint)(unsafe.Pointer(&l.Lifru[0])) = d
}
 (l *Lifreq) GetLifruUint() uint {
turn *(*uint)(unsafe.Pointer(&l.Lifru[0]))
}
 IoctlLifreq(fd int, req int, l *Lifreq) error {
turn ioctlPtr(fd, req, unsafe.Pointer(l))
}// Strioctl Helpers
 (s *Strioctl) SetInt(i int) {
Len = int32(unsafe.Sizeof(i))
Dp = (*int8)(unsafe.Pointer(&i))
}
 IoctlSetStrioctlRetInt(fd int, req int, s *Strioctl) (int, error) {
turn ioctlPtrRet(fd, req, unsafe.Pointer(s))
}
