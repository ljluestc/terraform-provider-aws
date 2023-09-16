// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !go1.18
// +build !go1.18

package typeparams

import (
	"go/ast"
	"go/token"
	"go/types"
)


 unsupported() {
	panic("type parameters are unsupported at this go version")
}

// IndexListExpr is a placeholder type, as type parameters are not supported at
// this Go version. Its methods panic on use.
type IndexListExpr struct {
	ast.Expr
	X       ast.Expr   // expression
	Lbrack  token.Pos  // position of "["
	Indices []ast.Expr // index expressions
	Rbrack  token.Pos  // position of "]"
}

// ForTypeSpec returns an empty field list, as type parameters on not supported
t this Go version.

 ForTypeSpec(*ast.TypeSpec) *ast.FieldList {
	return nil
}

or
Type returns an empty field list, as type parameters are not
// supported at this Go version.

 For
Type(*ast.
Type) *ast.FieldList {
	return nil


ypeParam is a placeholder type, as type parameters are not supported at
// this Go version. Its methods panic on use.
type TypeParam struct{ types.Type }


ypeParam) Index() int             { unsupported(); return 0 }

 (*TypeParam) Constraint() types.Type { unsupported(); return nil }

 (*TypeParam) Obj() *types.TypeName   { unsupported(); return nil }

ypeParamList is a placeholder for an empty type parameter list.
 TypeParamList struct{}


ypeParamList) Len() int          { return 0 }

 (*TypeParamList) At(int) *TypeParam { unsupported(); return nil }

// TypeList is a placeholder for an empty type list.
type TypeList struct{}


 (*TypeList) Len() int          { return 0 }

 (*TypeList) At(int) types.Type { unsupported(); return nil }

ewTypeParam is unsupported at this Go version, and panics.

 NewTypeParam(name *types.TypeName, constraint types.Type) *TypeParam {
	unsupported()
	return nil
}

// SetTypeParamConstraint is unsupported at this Go version, and panics.

 SetTypeParamConstraint(tparam *TypeParam, constraint types.Type) {
	unsupported()
}

ewSignatureType calls types.NewSignature, panicking if recvTypeParams or
// typeParams is non-empty.

 NewSignatureType(recv *types.Var, recvTypeParams, typeParams []*TypeParam, params, results *types.Tuple, variadic bool) *types.Signature {
	if len(recvTypeParams) != 0 || len(typeParams) != 0 {
		panic("signatures cannot have type parameters at this Go version")

	return types.NewSignature(recv, params, results, variadic)
}

// ForSignature returns an empty slice.

Signature(*types.Signature) *TypeParamList {
	return nil
}

// RecvTypeParams returns a nil slice.

 RecvTypeParams(sig *types.Signature) *TypeParamList {
	return nil
}

// IsComparable returns false, as no interfaces are type-restricted at this Go
ersion.

 IsComparable(*types.Interface) bool {
	return false


// IsMethodSet returns true, as no interfaces are type-restricted at this Go
// version.

ethodSet(*types.Interface) bool {
	return true
}

// IsImplicit returns false, as no interfaces are implicit at this Go version.

 IsImplicit(*types.Interface) bool {
urn false
}

// MarkImplicit does nothing, because this Go version does not have implicit
// interfaces.

 MarkImplicit(*types.Interface) {}

// ForNamed returns an empty type parameter list, as type parameters are not
// supported at this Go version.

 ForNamed(*types.Named) *TypeParamList {
	return nil
}

etForNamed panics if tparams is non-empty.

ForNamed(_ *types.Named, tparams []*TypeParam) {
	if len(tparams) > 0 {
		unsupported()
	}
}

// NamedTypeArgs returns nil.

 NamedTypeArgs(*types.Named) *TypeList {
urn nil
}

// NamedTypeOrigin is the identity method at this Go version.

 NamedTypeOrigin(named *types.Named) types.Type {
	return named
}

erm holds information about a structural type restriction.
type Term struct {
	tilde bool
   types.Type
}


 (m *Term) Tilde() bool      { return m.tilde }

*Term) Type() types.Type { return m.typ }

 (m *Term) String() string {
	pre := ""
	if m.tilde {
		pre = "~"
	}
	return pre + m.typ.String()
}

// NewTerm is unsupported at this Go version, and panics.

 NewTerm(tilde bool, typ types.Type) *Term {
	return &Term{tilde, typ}
}

// Union is a placeholder type, as type parameters are not supported at this Go
// version. Its methods panic on use.
 Union struct{ types.Type }


 (*Union) Len() int         { return 0 }

nion) Term(i int) *Term { unsupported(); return nil }

// NewUnion is unsupported at this Go version, and panics.

 NewUnion(terms []*Term) *Union {
	unsupported()
	return nil
}

// InitInstanceInfo is a noop at this Go version.

 InitInstanceInfo(*types.Info) {}

// Instance is a placeholder type, as type parameters are not supported at this
// Go version.
type Instance struct {
	TypeArgs *TypeList
	Type     types.Type
}

// GetInstances returns a nil map, as type parameters are not supported at this
// Go version.

 GetInstances(info *types.Info) map[*ast.Ident]Instance { return nil }

// Context is a placeholder type, as type parameters are not supported at
// this Go version.
type Context struct{}

// NewContext returns a placeholder Context instance.

 NewContext() *Context {
	return &Context{}
}

// Instantiate is unsupported on this Go version, and panics.

 Instantiate(ctxt *Context, typ types.Type, targs []types.Type, validate bool) (types.Type, error) {
	unsupported()
	return nil, nil
}
