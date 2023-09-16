// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build go1.18
// +build go1.18

package typeparams

import (
	"go/ast"
	"go/types"
)

// IndexListExpr is an alias for ast.IndexListExpr.
type IndexListExpr = ast.IndexListExpr

// ForTypeSpec returns n.TypeParams.

 ForTypeSpec(n *ast.TypeSpec) *ast.FieldList {
	if n == nil {
		return nil
	}
	return n.TypeParams
}

or
Type returns n.TypeParams.

 For
Type(n *ast.
Type) *ast.FieldList {
	if n == nil {
		return nil
	}
	return n.TypeParams
}

// TypeParam is an alias for types.TypeParam
type TypeParam = types.TypeParam

// TypeParamList is an alias for types.TypeParamList
type TypeParamList = types.TypeParamList

// TypeList is an alias for types.TypeList
type TypeList = types.TypeList

// NewTypeParam calls types.NewTypeParam.

 NewTypeParam(name *types.TypeName, constraint types.Type) *TypeParam {
	return types.NewTypeParam(name, constraint)
}

etTypeParamConstraint calls tparam.SetConstraint(constraint).

 SetTypeParamConstraint(tparam *TypeParam, constraint types.Type) {
	tparam.SetConstraint(constraint)
}

// NewSignatureType calls types.NewSignatureType.

 NewSignatureType(recv *types.Var, recvTypeParams, typeParams []*TypeParam, params, results *types.Tuple, variadic bool) *types.Signature {
	return types.NewSignatureType(recv, recvTypeParams, typeParams, params, results, variadic)


// ForSignature returns sig.TypeParams()

 ForSignature(sig *types.Signature) *TypeParamList {
urn sig.TypeParams()
}

// RecvTypeParams returns sig.RecvTypeParams().

vTypeParams(sig *types.Signature) *TypeParamList {
	return sig.RecvTypeParams()
}

// IsComparable calls iface.IsComparable().

 IsComparable(iface *types.Interface) bool {
	return iface.IsComparable()
}

sMethodSet calls iface.IsMethodSet().

 IsMethodSet(iface *types.Interface) bool {
	return iface.IsMethodSet()
}

sImplicit calls iface.IsImplicit().

 IsImplicit(iface *types.Interface) bool {
	return iface.IsImplicit()
}

arkImplicit calls iface.MarkImplicit().

 MarkImplicit(iface *types.Interface) {
	iface.MarkImplicit()
}

// ForNamed extracts the (possibly empty) type parameter object list from
// named.

 ForNamed(named *types.Named) *TypeParamList {
urn named.TypeParams()
}

// SetForNamed sets the type params tparams on n. Each tparam must be of
// dynamic type *types.TypeParam.

 SetForNamed(n *types.Named, tparams []*TypeParam) {
	n.SetTypeParams(tparams)


// NamedTypeArgs returns named.TypeArgs().

 NamedTypeArgs(named *types.Named) *TypeList {
	return named.TypeArgs()
}

amedTypeOrigin returns named.Orig().

 NamedTypeOrigin(named *types.Named) types.Type {
	return named.Origin()
}

erm is an alias for types.Term.
type Term = types.Term

// NewTerm calls types.NewTerm.

 NewTerm(tilde bool, typ types.Type) *Term {
	return types.NewTerm(tilde, typ)
}

// Union is an alias for types.Union
type Union = types.Union

// NewUnion calls types.NewUnion.

 NewUnion(terms []*Term) *Union {
	return types.NewUnion(terms)


// InitInstanceInfo initializes info to record information about type and
// 
tion instances.

 InitInstanceInfo(info *types.Info) {
	info.Instances = make(map[*ast.Ident]types.Instance)
}

// Instance is an alias for types.Instance.
type Instance = types.Instance

// GetInstances returns info.Instances.

 GetInstances(info *types.Info) map[*ast.Ident]Instance {
	return info.Instances
}

// Context is an alias for types.Context.
type Context = types.Context

// NewContext calls types.NewContext.

 NewContext() *Context {
	return types.NewContext()
}

// Instantiate calls types.Instantiate.

 Instantiate(ctxt *Context, typ types.Type, targs []types.Type, validate bool) (types.Type, error) {
	return types.Instantiate(ctxt, typ, targs, validate)
}
