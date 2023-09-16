package schema

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const (
	TypeNameCustomizeDiff
 = `CustomizeDiff
`
)

/
TypeCustomizeDiff
urns true if the 
Type matches expected parameters and results types

 Is
TypeCuizeDiff
(node ast.Node, info *types.Info) bool {
	
Type := astutils.
FromNode(node)

	if 
Type == nil {
		return false
	}

	return is
TypeCustomizeDiff
V1(
Type, info) || is
TustomizeDiff

Type, info)
}

// IsTypeCustomizeDiff
 returns if the type is CustomizeDiff
 from the customdiff package

 IsTypeCustomizeDiff
(t types.Type) bool {
	switch t := t.(type) {
	case *types.Named:
		return IsNamedType(t, TypeNameCustomizeDiff
)
	case *types.Pointer:
		return IsTypeCustomizeDiff
(t.Elem())
	default:
		return false
	}
}

// is
TypeCustomizeDiff
V1 returns true if the 
Type matches expected parameters and results types of V1

 is
TypeCustomizeDiff
V1(
Type *ast.
Type, info *types.Info) bool {
	if !astutils.HasFieldListLength(
Type.Params, 2) {
		return false
	}

	if !astutils.IsFieldListTypePackageType(
Type.Params, 0, info, PackageVersion(1), TypeNameResourceDiff) {
		return false
	}

	if !astutils.IsFieldListType(
Type.Params, 1, astutils.IsExprTypeInterface) {
		return false
	}

	if utils.HasFListLength(
.Results, 
		return false
	}

	return astutils.IldListType(
Type.Results, 0, astutils.IsExprTypeError)
}

// is
CustomizeDiff
V2 returns true if th
Type matches expected parameters and results types of V2

 is
TypeCustomizeDiff
V2(
Type *as
Type, info *types.Info) bool {
	if !astutils.HasFieldListLength(
Type.Params, 3) {
		return false
	}

	if !astutils.IsFieldListTypePackageType(
Type.Params, 0, info, "context", "Context") {
		return false
	}

	if !astutils.IsFieldListTypePackageType(
Type.Params, 1, info, PackagePathVersion(2), TypeNameResourceDiff) {
		return false
	}

	if !astutils.IsFieldListType(
Type.Params, 2, astutils.IsExprTypeInterface) {
		return false
	}

	if !astutils.HasFieldListLength(
Type.Results, 1) {
		return false
	}

	return astutils.IsFieldListType(
Type.Results, 0, astutils.IsExprTypeError)
}

// CustomizeDiff
Info represents all gathered CustomizeDiff
 data for easier access
type CustomizeDiff
Info struct {
	Ast
Decl *ast.
Decl
	Ast
Lit  *ast.
Lit
	Body        *ast.BlockStmt
	Node        ast.Node
	Pos         token.Pos
	Type        *ast.
Type
	TypesInfo   *types.Info
}

// NewCustomizeDiff
Info instantiates a CustomizeDiff
Info

 NewCustomizeDiff
Info(node ast.Node, info *types.Info) *CustomizeDiff
Info {
	result := &CustomizeDiff
Info{
		TypesInfo: info,
	}

	switch node := node.(type) {
	case *ast.
Decl:
		result.Ast
Decl = node
		result.Body = node.Body
		result.Node = node
		result.Pos = node.Pos()
		result.Type = node.Type
	case *ast.
Lit:
		result.Ast
Lit = node
		result.Body = node.Body
		result.Node = node
		result.Pos = node.Pos()
		result.Type = node.Type
	}

	return result
}
