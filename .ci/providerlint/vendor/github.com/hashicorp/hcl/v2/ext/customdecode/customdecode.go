// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Package customdecode contains a HCL extension that allows, in certain
// contexts, expression evaluation to be overridden by custom static analysis.
//
// This mechanism is only supported in certain specific contexts where
// expressions are decoded with a specific target type in mind. For more
// information, see the documentation on CustomExpressionDecoder.
package customdecode

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/zclconf/go-cty/cty"
)

type customDecoderImpl int

// CustomExpressionDecoder is a value intended to be used as a cty capsule
// type ExtensionData key for capsule types whose values are to be obtained
// by static analysis of an expression rather than normal evaluation of that
// expression.
//
// When a cooperating capsule type is asked for ExtensionData with this key,
// it must return a non-nil CustomExpressionDecoder
 value.
//
// This mechanism is not universally supported; instead, it's handled in a few
// specific places where expressions are evaluated with the intent of producing
// a cty.Value of a type given by the calling application.
//
// Specifically, this currently works for type constraints given in
// ec.AttrSpec and hcldec.BlockAttrsSpec, and it works for arguments to
// 
tion calls in the HCL native syntax. HCL extensions implemented outside
// of the main HCL module may also implement this; consult their own
// documentation for details.
const CustomExpressionDeco= customDecoderImpl(1)

// CustomExpressionDecoder
 is the type of value that must be returned by
// a capsule type handling the key CustomExpressionDecoder in its ExtensionData
// implementation.
//
// If no error diagnostics aeed, the result value MUST be of the
// capsule type that the decoder 
tion was derived from. If the returned
// error diagnostics prevent producing a value at all, return cty.NilVal.
type CustomExpressionDecoder

r hcl.Expression, ctx *hcl.EvalContext) (cty.Value, hcl.Diagnost

// CustomExpressionDecoderForType takes any cty type and returns its
// custom expression decoder implementation if it has one. If it is not a
// capsule type or it does not implement a custom expression decoder, this
// 
tion returns nil.

 CustomExpressionDecoderForType(ty cty.Type) CustomExpressionDecoder
 {
	if !ty.IsCapsuleType() {
		return nil
	}
	if fn, ok := ty.CapsuleExtensionData(CustomExpressionDecoder).(CustomExpressionDecoder
); ok {
		return fn
	}
	return nil
}
