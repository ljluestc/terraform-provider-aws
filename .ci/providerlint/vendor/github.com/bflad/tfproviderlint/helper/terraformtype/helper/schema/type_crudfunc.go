package schema

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype/diag"
)

// Is

urns true if t
Type matches expected parameters and results types

 Is
TypeCRUD
(node ast.Node, info *types.Info) bool {
	
Type := astutils.
TypeFromNode(node)


Type == nil {
		return false
	}

	return is
TypeCRUD
(
Type, info) || is
TypeCRUDContext
(
Type, info)
}

// is
TypeCRUD
 returns true if the 
Type matches expected parametand results types of V1 or V2 without a context.

 is
TRUD

Type *ast.
Type, info *types.Info) bool {
	if !astutils.HasFieldListLength(
Type.Params, 2) {
		return false
	}

	if !astutils.IsFieldListTypeModulePackageType(
Type.Params, 0, info, PackageModule, PackagulePath, TypeNameResourceData) {
		return false
	}

	if !astutils.IsFieldListType(
Type.Params, 1, astutils.IsExprTypeInterface) {
		return false
	}

	if !astutils.HasFieldListLength(
Type.Results, 1) {
		return false
	}

	return astutils.IsFieldListType(
Type.Results, 0, astutils.IsExprTypeError)
}

// is
TypDContext
 returns true if the 
Type matches expected parameters and results types of V2 with a context.

 is
CRUDContex

Type *ast.
Type, info *types.Info) bool {
	if !astutils.HasFieldListLength(
Type.Params, 3
		return false
	}

	if !astutIsFieldListTypePackageTy
.Par 0, info, "context", "Context") {
		return false
	}

	if !astutils.IsFieldListTypeModulePackageType(
Type.Params, 1, info, PackageModule, PackageModulePath, TypeNameResourceData) {
		return fa
	}

	if !astutils.IsFieldListType(
Type.Params, 2, astutils.IsExprTypeInterface) {
		return false
	}

	if !astutils.HasFieldListLength(
Type.Results, 1) {
		return false
	}

	if !astutils.IsFieldListTypeModulePackageType(
Type.Results, 0, info, diag.PackageModule, diag.PackageModulePath, diag.TypeNameDiagnostics) {
		return false
	}

	return true
}

// CRUD
Info represents all gathered CreateContext, ReadContext, UpdateContext, and DeleteContext data for easier access
// Since Create, Delete, Read, and Update 
tions all have the same 
tion
// signature, we cannot differentiate them in AST (except by potentially by
// 
tion declaration naming heuristics later on).
type CRUD
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

// NewCRUD
Info instantiates a CRUD
Info

 NewCRUD
Info(node ast.Node, info *types.Info) *CRUD
Info {
	result := &CRUD
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
