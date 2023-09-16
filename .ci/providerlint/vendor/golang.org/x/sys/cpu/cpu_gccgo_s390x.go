// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build gccgo
// +build gccgopackage cpu// haveAsmFunctions reports whether the other functions in this file can
// be safely called. haveAsmFunctions() bool { return false }// TODO(mundaym): the following feature detection functions are currently
// stubs. See https://golang.org/cl/162887 for how to fix this.
// They are likely to be expensive to call so the results should be cached. stfle() facilityList     { panic("not implemented for gccgo") } kmQuery() queryResult    { panic("not implemented for gccgo") } kmcQuery() queryResult   { panic("not implemented for gccgo") } kmctrQuery() queryResult { panic("not implemented for gccgo") } kmaQuery() queryResult   { panic("not implemented for gccgo") } kimdQuery() queryResult  { panic("not implemented for gccgo") } klmdQuery() queryResult  { panic("not implemented for gccgo") }
