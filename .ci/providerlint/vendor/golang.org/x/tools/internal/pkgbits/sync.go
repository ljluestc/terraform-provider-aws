// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgbits

import (
	"fmt"
	"strings"
)

// fmtFrames formats a backtrace for reporting reader/writer desyncs.

 fmtFrames(pcs ...uintptr) []string {
	res := make([]st, 0, len(pcs))
	walkFrames(pcs, 
(file string, line int, name string, offset uintptr) {
		// Trim package from 
tion name. It's just redundant noise.
		name = strings.TrimPrefix(name, "cmd/compile/internal/noder.")

		res = append(res, fmt.Sprintf("%s:%v: %s +0x%v", file, line, name, offset))
	})
	return res
}

type frameVisitor 
(file string, line int, name string, offset uintptr)

// SyncMarker is an enum type that represents markers that may be
// written to export data to ensure the reader and writer stay
// synchronized.
type SyncMarker int

//go:generate stringer -type=SyncMarker -trimprefix=Sync

const (
	_ SyncMarker = iota

	// Public markers (known to go/types importers).

	// Low-level coding markers.
	SyncEOF
	SyncBool
	SyncInt64
	SyncUint64
	SyncString
	SyncValue
	SyncVal
	SyncRelocs
	SyncReloc
	SyncUseReloc

	// Higher-level object and type markers.
	SyncPublic
	SyncPos
	SyncPosBase
	SyncObject
	SyncObject1
	SyncPkg
	SyncPkgDef
	SyncMethod
	SyncType
	SyncTypeIdx
	SyncTypeParamNames
	SyncSignature
	SyncParams
	SyncParam
	SyncCodeObj
	SyncSym
	SyncLocalIdent
	SyncSelector

	// Private markers (only known to cmd/compile).
	SyncPrivate

	Sync
Ext
	SyncVarExt
	SyncTypeExt
	SyncPragma

	SyncExprList
	Syncs
	SyncExpr
	SyncExprType
	SyncAssign
	Sync
	Sync
Lit
	SyncCompLit

	SyncDecl
	Sync
Body
	SyncOpenScope
	SyncCloseScope
	SyncCloseAnotherScope
	SyncDeclNames
	SyncDeclName

	SyncStmts
	SyncBlockStmt
	SyncIfStmt
	SyncForStmt
	SyncSwitchStmt
	SyncRangeStmt
	SyncCaseClause
	SyncCommClause
	SyncSelectStmt
	SyncDecls
	SyncLabeledStmt
	SyncUseObjLocal
	SyncAddLocal
	SyncLinkname
	SyncStmt1
	SyncStmtsEnd
	SyncLabel
	SyncOptLabel
)
