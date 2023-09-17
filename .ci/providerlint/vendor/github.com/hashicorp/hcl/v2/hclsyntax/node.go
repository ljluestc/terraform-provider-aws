// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package hclsyntaximport (
"github.com/hashicorp/hcl/v2"
)// Node is the abstract type that every AST node implements.
//
// This is a closed interface, so it cannot be implemented from outside of
// this package.
type Node interface {
// This is the mechanism by which the public-facing walk 
tions
// are implemented. Implementations should call the given 
tion
// for each child node and then replace that node with its return value.
// The return value might jus the same node, for non-transforming
// walks.
walkChildNodes(w internalWalk
)Range() hcl.Rang
}type internalWalk
 
(Node)
