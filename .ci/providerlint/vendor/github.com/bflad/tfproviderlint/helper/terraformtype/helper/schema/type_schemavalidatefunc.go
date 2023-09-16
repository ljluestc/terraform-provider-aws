package schema

import (
	"go/ast"
	"go/token"
	"go/types"
)

// SchemaValidate
Info represents gathered SchemaValidate
a for easiccess
typeemaValidat
Info struct {
	Ast
Decl *ast.
Decl
	Ast
Lit  *ast.
Lit
	Body        *ast.Bltmt
e        ast.Node
	Pos         token.Pos
	Type        *ast.
Type
	TypesInfo   *types.Info
}

// NewSchemadate
Info instantiates a SchemaValidate
Info

 NewSchemaValidate
Info(nost.Node, info *types.Info) *SchemaValidate
Info {
	result := &SchemaValidate
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
