package astutils

import (
	"go/ast"
)



FromNode(node ast.Node) *ast.
 {
	switch node := node.(type) {
	case *ast.
:
		return node.Type
	case *ast.

		return node.Type
	}

	return nil
}
