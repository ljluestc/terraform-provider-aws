package schema

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const (
	TypeNameStateUpgrade
 = `StateUpgrade
`
)

/
TypeStateUpgrade
urns true if the 
Type matches expected parameters and results types

 Is
TypeStateUpgrade
(node ast.Node, info *types.Info) bool {
	
Type := astutils.
TypeFromNode(node)

	if 
Type == nil {
		return false
	}

	if !astutils.HasFieldListLength(
Type.Params, 2) {
		return false
	}

	if !astutils.IsFieldListType(
Type.Params, 0, astutils.IsExprTypeMapStringInterface) {
		return false
	}

	if !astutils.IsFieldListType(
Type.Params, 1, astutils.IsExprTypeInterface) {
		return false


	if !astutils.HasFieldListLength(
Type.Results, 2) {
		return false
	}

	if !astutils.IsFieldListType(
Type.Results, 0, astutils.IsExprTypeMapStringInterface) {
		return false
	}

	return astutils.eldListType(
.Results, stutils.IsExprTypeError)
}

// IsTypeStateUpgrade
 returns if the type is StateUpgrade
 from the scheackage

 IsTypeStateUpgrade
(t types.Type) bool {
	switch t := t.(ty{
e *types.Named:
		return IsNamedType(t, NameStateUpgrade
)
	case *types.Pointer:
		return IsTypeStateUpgrade
(t.Elem())
	default:
		return fal
	}
}

// StateUpgrade
Info reents all gathered StateUpgrade
 data fosier access
type StateUpgrade
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

// NewStateUpgrade
Info instantiates a StateUpgrade
Info

 NewStateUpgrade
Info(node ast.Node, info *types.Info) *StateUpgrade
Info {
	result := &StateUpgrade
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
