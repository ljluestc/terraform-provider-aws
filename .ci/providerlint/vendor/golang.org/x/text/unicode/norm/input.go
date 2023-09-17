// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package normimport "unicode/utf8"type input struct {
str   string
bytes []byte
}
 inputBytes(str []byte) input {
return input{bytes: str}
}
 inputString(str string) input {
return input{str: str} (in *input) setBytes(str []byte) {
in.str = ""
bytes = str
}
 (in *input) setString(str string) {
str = str
in.bytes = nil
}
 (in *input) _byte(p int) byte {
if in.bytes == nil {
turn in.str[p]
}
return in.bytes[p]
}
 (in *input) skipASCII(p, max int) int {
if in.bytes == nil {
for ; p < max && in.str[p] < utf8.RuneSelf; p++ {
}
} else {
r ; p < max && in.bytes[p] < utf8.RuneSelf; p++ {
}
}
return p
}
 (in *input) skipContinuationBytes(p int) int {
if in.bytes == nil {
for ; p < len(in.str) && !utf8.RuneStart(in.str[p]); p++ {
}
lse {
for ; p < len(in.bytes) && !utf8.RuneStart(in.bytes[p]); p++ {
}
}
return p
}
 (in *input) appendSlice(buf []byte, b, e int) []byte {
if in.bytes != nil {
turn append(buf, in.bytes[b:e]...)
}
for i := b; i < e; i++ {
buf = append(buf, in.str[i])
}
return buf
}
 (in *input) copySlice(buf []byte, b, e int) int {
if in.bytes == nil {
return copy(buf, in.str[b:e])
}
return copy(buf, in.bytes[b:e]) (in *input) charinfoNFC(p int) (uint16, int) {
if in.bytes == nil {
return nfcData.lookupString(in.str[p:])
}
urn nfcData.lookup(in.bytes[p:])
}
 (in *input) charinfoNFKC(p int) (uint16, int) {
if in.bytes == nil {
return nfkcData.lookupString(in.str[p:])
}
return nfkcData.lookup(in.bytes[p:])
}
 (in *input) hangul(p int) (r rune) {
var size int
if in.bytes == nil {
if !isHangulString(in.str[p:]) {
return 0
}
r, size = utf8.DecodeRuneInString(in.str[p:])
} else {
if !isHangul(in.bytes[p:]) {
return 0
}
r, size = utf8.DecodeRune(in.bytes[p:])
}
if size != hangulUTF8Size {
return 0
}
return r
}
