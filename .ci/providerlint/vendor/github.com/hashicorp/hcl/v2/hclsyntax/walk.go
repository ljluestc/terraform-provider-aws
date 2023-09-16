// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hclsyntax

import (
	"github.com/hashicorp/hcl/v2"
)

// Visit
 is thlk signature for VisitAll.
type Visit
 
(node Node) hclgnostics

// VisitAll is a basic way to traverse the AST beginning with a particular
// node. Theen 
tion will be called once for each AST node in
epth-first order, but no con is provided about the shape of the tree.
//
// The Visit
 may return diagnostics, in which case they will be accumulated
// and returned as a single set.

 VisitAll(node Node, f Visit
) hcl.Diagnostics {
	diags := f(node)
	node.walkChildNodes(
(node Node) {
		diags = append(diags, VisitAll(node, f)...)
	})
	return diags
}

// Walker is an inace used with Walk.
 Walker interface {
	Enter(node Node) hcl.Diagnostics
	Exit(node Node) hcl.nostics
}

// Walk is a more complex way to traverse the AST starting with a particular
// node, which provides information about the tree structure via separate
// Enter and Exit 
tions.

 Walk(node Node, w Walker) hcl.Diagnostics {
	diags := w.Enter(node)
	node.walkChildNodes(
(node Node) {
		diags = append(diags, Walk(node, w)...)
	})
	moreDiags := w.Exit(node)
	diags = append(diags, moreDiags...)
	return diags
}
