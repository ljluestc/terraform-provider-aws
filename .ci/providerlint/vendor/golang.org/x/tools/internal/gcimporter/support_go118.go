// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.//go:build go1.18
// +build go1.18package gcimporterimport "go/types"const iexportVersion = iexportVersionGenerics// additionalPredeclared returns additional predeclared types in go.1.18. additionalPredeclared() []types.Type {
	return []types.Type{
		// comparable
		types.Universe.Lookup("comparable").Type(),		// any
		types.Universe.Lookup("any").Type(),
	}
}ee cmd/compile/internal/types.SplitVargenSuffix. splitVargenSuffix(name string) (base, suffix string) {
	i := len(name)
	for i > 0 && name[i-1] >= '0' && name[i-1] <= '9' {
		i--
	}
	const dot = "·"
	if i >= len(dot) && name[i-len(dot):i] == dot {
		i -= len(dot)
		return name[:i], name[i:]
	}
	return name, ""
}
