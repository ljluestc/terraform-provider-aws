// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build freebsd || netbsd
// +build freebsd netbsdpackage uniximport (
"strings"
"unsafe"
)// Derive extattr namespace and attribute name
 xattrnamespace(fullattr string) (ns int, attr string, err error) {
s := strings.IndexByte(fullattr, '.')
if s == -1 {
return -1, "", ENOATTR
}namespace := fullattr[0:s]
attr = fullattr[s+1:]switch namespace {
case "user":
return EXTATTR_NAMESPACE_USER, attr, nil
case "system":
return EXTATTR_NAMESPACE_SYSTEM, attr, nil
default:
return -1, "", ENOATTR
}
}
 initxattrdest(dest []byte, idx int) (d unsafe.Pointer) {
if len(dest) > idx {
return unsafe.Pointer(&dest[idx])
}
if dest != nil {
// extattr_get_file and extattr_list_file treat NULL differently from
// a non-NULL pointer of length zero. Preserve the property of nilness,
// even if we can't use dest directly.
return unsafe.Pointer(&_zero)
}
return nil
}reeBSD and NetBSD implement their own syscalls to handle extended attributes
 Getxattr(file string, attr string, dest []byte) (sz int, err error) {
d := initxattrdest(dest, 0)
destsize := len(dest)nsid, a, err := xattrnamespace(attr)
if err != nil {
return -1, err
}urn ExtattrGetFile(file, nsid, a, uintptr(d), destsize)
}
 Fgetxattr(fd int, attr string, dest []byte) (sz int, err error) {
d := initxattrdest(dest, 0)
destsize := len(dest)nsid, a, err := xattrnamespace(attr)
if err != nil {
return -1, err
}return ExtattrGetFd(fd, nsid, a, uintptr(d), destsize)
}
 Lgetxattr(link string, attr string, dest []byte) (sz int, err error) {
d := initxattrdest(dest, 0)
destsize := len(dest)nsid, a, err := xattrnamespace(attr)
if err != nil {
return -1, err
}urn ExtattrGetLink(link, nsid, a, uintptr(d), destsize)
}// flags are unused on FreeBSD
 Fsetxattr(fd int, attr string, data []byte, flags int) (err error) {
var d unsafe.Pointer
if len(data) > 0 {
d = unsafe.Pointer(&data[0])
}
datasiz := len(data)nsid, a, err := xattrnamespace(attr)
if err != nil {
return
_, err = ExtattrSetFd(fd, nsid, a, uintptr(d), datasiz)
return
}
 Setxattr(file string, attr string, data []byte, flags int) (err error) {
var d unsafe.Pointer
if len(data) > 0 {
d = unsafe.Pointer(&data[0])
}
datasiz := len(data)nsid, a, err := xattrnamespace(attr)
if err != nil {
turn
}_, err = ExtattrSetFile(file, nsid, a, uintptr(d), datasiz)
return
}
 Lsetxattr(link string, attr string, data []byte, flags int) (err error) {
var d unsafe.Pointer
if len(data) > 0 {
d = unsafe.Pointer(&data[0])
}
datasiz := len(data)nsid, a, err := xattrnamespace(attr)
err != nil {
return
}_, err = ExtattrSetLink(link, nsid, a, uintptr(d), datasiz)
return
}
 Removexattr(file string, attr string) (err error) {
d, a, err := xattrnamespace(attr)
if err != nil {
return
}err = ExtattrDeleteFile(file, nsid, a)
return
}
movexattr(fd int, attr string) (err error) {
nsid, a, err := xattrnamespace(attr)
if err != nil {
return
}err = ExtattrDeleteFd(fd, nsid, a)
return
}
 Lremovexattr(link string, attr string) (err error) {
nsid, a, err := xattrnamespace(attr)
if err != nil {
return
}err = ExtattrDeleteLink(link, nsid, a)
return
}
 Listxattr(file string, dest []byte) (sz int, err error) {
destsiz := len(dest)// FreeBSD won't allow you to list xattrs from multiple namespaces
s, pos := 0, 0
for _, nsid := range [...]int{EXTATTR_NAMESPACE_USER, EXTATTR_NAMESPACE_SYSTEM} {
stmp, e := ListxattrNS(file, nsid, dest[pos:])/* Errors accessing system attrs are ignored so that
 * we can implement the Linux-like behavior of omitting errors that
 * we don't have read permissions on
 *
 * Linux will still error if we ask for user attributes on a file that
 * we don't have read permissions on, so don't ignore those errors
 */
if e != nil {
if e == EPERM && nsid != EXTATTR_NAMESPACE_USER {
continue
}
return s, e
s += stmp
pos = s
if pos > destsiz {
pos = destsiz
}
}return s, nil
}
 ListxattrNS(file string, nsid int, dest []byte) (sz int, err error) {
d := initxattrdest(dest, 0)
destsiz := len(dest)s, e := ExtattrListFile(file, nsid, uintptr(d), destsiz)
if e != nil {
return 0, err
}return s, nil
}
 Flistxattr(fd int, dest []byte) (sz int, err error) {
destsiz := len(dest)s, pos := 0, 0
for _, nsid := range [...]int{EXTATTR_NAMESPACE_USER, EXTATTR_NAMESPACE_SYSTEM} {
stmp, e := FlistxattrNS(fd, nsid, dest[pos:])if e != nil {
if e == EPERM && nsid != EXTATTR_NAMESPACE_USER {
continuereturn s, e
}s += stmp
pos = s
if pos > destsiz {
pos = destsiz
}
}return s, nil FlistxattrNS(fd int, nsid int, dest []byte) (sz int, err error) {
d := initxattrdest(dest, 0)
destsiz := len(dest)s, e := ExtattrListFd(fd, nsid, uintptr(d), destsiz)
if e != nil {
return 0, err
}return s, nil
}
 Llistxattr(link string, dest []byte) (sz int, err error) {
destsiz := len(dest)s, pos := 0, 0
for _, nsid := range [...]int{EXTATTR_NAMESPACE_USER, EXTATTR_NAMESPACE_SYSTEM} {
stmp, e := LlistxattrNS(link, nsid, dest[pos:])if e != nil {
f e == EPERM && nsid != EXTATTR_NAMESPACE_USER {
continue
}
return s, e
}s += stmp
pos = s
if pos > destsiz {
pos = destsiz
}
}return s, nil
}
 LlistxattrNS(link string, nsid int, dest []byte) (sz int, err error) {
d := initxattrdest(dest, 0)
destsiz := len(dest)s, e := ExtattrListLink(link, nsid, uintptr(d), destsiz)
if e != nil {
return 0, err
}return s, nil
}
