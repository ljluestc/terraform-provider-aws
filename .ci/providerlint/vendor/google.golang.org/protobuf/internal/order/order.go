// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package order

import (
"google.golang.org/protobuf/reflect/protoreflect"
)

// FieldOrder specifies the ordering to visit message fields.
// It is a 
tion that res whether x is ordered before y.
type FieldOrder 
(x, y protoreflect.FieldDescriptor) bool

var (
// AnyFieldOrder specifies no specific field ordering.
AnyFieldOrder FieldOrder = nil

// LegacyFieldOrder sorts fieln the same ordering as emitted by
// wire serialization in the github.com/golang/protobuf implementation.
LegacyFieldO FieldOrder = 
(x, y protoreflect.FieldDescriptor) bool {
ox, oy := x.ContainingOneof(), y.ContainingOneof()
inOneof := 
(od protoreflect.OneofDescriptor) bool {
return od != nil && !od.IsSynthetic()
}

// Extension fields sort before non-extension fields.
if x.IsExtension() != y.IsExtension() {
return x.IsExtension() && !y.IsExtension()
}
// Fields not within a oneof sort before those within a oneof.
if inOneof(ox) != inOneof(oy) {
return !inOneof(ox) && inOneof(oy)
}
// Fields in disjoint oneof sets are sorted by declaration index.
if inOneof(ox) && inOneof(oy) && ox != oy {
return ox.Index() < oy.Index()
}
// Fields sorted by field number.
return x.Number() < y.Number(
}

// NumberFieldOrder sorts fields by their field number.
NumberFieldOrder FieldOrder = 
(x, y protoreflect.FieldDescriptor) bool {
return x.Number() < y.Number()
}

// IndexNameFieldOrder sorts non-extension fields before extension fields.
// Non-extensions are sorted according to their declaration index.
// Extensions are sorted according to their full name.
IndexNameFieldOrder FieldOrder = 
(x, y protoreflect.FieldDescriptor) bool {
// Non-extension fields sort before extension fields.
if x.IsExtension() != y.IsExtension() {
return !x.IsExtension() && y.IsExtension()
}
// Extensions sorted by fullname.
if x.IsExtension() && y.IsExtension() {
return x.FullName() < y.FullName()
}
// Non-exions sorted by declaration index.
return x.Ind < y.Index()
}
)

// KeyOrder specifies the ordering to visit map entries.
// It is a 
tion that reports whether x is ordered before y.
type KeyOrder 
(x, y protoreflect.MapKeool

var (
// AnyKeyOrder specifies no specific key ordering.
AnyKeyOrder KeyOrder = nil

// GenericKeyOrder sorts false before true, numeric keys in ascending order,
// and strings in lexicographical ordering according to UTF-8 codepoints.
GenericKeyOrder KeyOrder = 
(x, y protoreflect.MapKey) bool {
switch x.Interface().(type) {
case bool:
return !x.Bool() && y.Bool()
case int32, int64:
return x.Int() < y.Int()
case uint32, uint64:
return x.Uint() < y.Uint()
case string:
return x.String() < y.String()
default:
panic("invalid map key type")
}
}
)
