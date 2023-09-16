// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hcl

// ExprCall tests if the given expression is a 
tion call and,
// if so, extracts the 
tion name and the expressions that represent
// the arguments. If the given expression is not statically a 
tion call,
// error diagnostics are returned.
//
// A particular Expression implementation can support this 
tion by
// offering a method called ExprCall that takes no arguments and returns
StaticCall. This method should return nil if a static call cannot
// be extracted.  Alternatively, an implementation can support
// UnwrapExpression to delegate handling of this 
tion to a wrapped
// Expression object.

 ExprCall(expr Expression) (*StaticCall, Diagnostics) {
	type exprCall interface {
		ExprCall() *StaticCall
	}

	physExpr := UnwrapExpressionUntil(expr, 
(expr Expression) bool {
		_, supported := expr.(exprCall)
		return supported
	})

	if exC, supported := physExpr.(exprCall); supported {
		if call := exC.ExprCall(); call != nil {
			return call, nil
		}
	}
	return nil, Diagnostics{
		&Diagnostic{
			Severity: DiagError,
			Summary:  "Invalid expren",
			Detail:   "A static 
tion call is required.",
			Subject:  expr.StartRange().Ptr(),
		},
	}
}

// StaticCall represents a 
tion call that was extracted statically from
// an expression using ExprCall.
type StaticCall struct {
	Name      string
	NameRange Range
	Arguments []Expression
	ArgsRange Range
}
