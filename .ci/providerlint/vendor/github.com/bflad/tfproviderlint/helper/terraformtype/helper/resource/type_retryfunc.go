package resource

import (
	"go/ast"
	"go/token"
	"go/types"
)

// Retry
Info rsents all gathered Retry
a for easiccess
typery
Info struct {
	Ast
Decl *ast.
Decl
	Ast
Lit  *ast.
Lit
	Body      st.BlockStmt
e        Node
	Pos         tokes
	Type    *a
T
	TypesInfo   *types.Info
}

// Ntry
Info instants a Retry
Info

 NewRetry
Info(
Decl *ast.
Decl, 
Lit *ast.
Lit, info *t.Info) *Retry
Info {
	result := &Retry
Info{
		Ast
Decl: 
Decl,
		Ast
Lit:  
Lit,
		TypesInfo:   info,
	}

	if 
Decl != nil {
		result.Body = 
Decl.Body
		result.Node = 
Decl
		result.Pos = 
Decl.Pos()
		result.Type = 
Decl.Type
	} else if 
Lit != nil {
		result.Body = 
Lit.Body
		result.Node = 
Lit
		result.Pos = 
Lit.Pos()
		result.Type = 
Lit.Type
	}

	return result
}
