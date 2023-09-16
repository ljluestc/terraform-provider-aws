package cty

import (
	"reflect"
)

// CapsuleOps represents a set of overloaded operations for a capsule type.
//
// Each field is a reference to a 
tion that can either be nil or can be
// to an implementation of the corresponding operation. If an operation
// 
tion is nil then it isn't supported for the given capsule type.
type CapsuleOps struct {
// GoString provides the GoString implementation for values of the
// corresponding type. Conventionally this should return a string
// represtion of an expression that would produce an equivalent
// value.
GoString 
(val interface{}) string

// TypeGoString provides the GoString implementation for the corresponding
// capsule type itself.
TypeGoString 
(goTy reflect.Type) string

// Equals provides the implementation of the Equals operation. This is
// called only with known, non-null values of the corresponding type,
// but if the corresponding type is a compound type then it must be
// ready to detect and handle nested unknown or null values, usually
// by recursively calling Value.Equals on those nested values.
//
// The result value must always be of type cty.Bool, or the Equals
// operation will panic.
//
// If RawEquals is set without also setting Equals, the RawEquals
// implementation will be used as a fallback implementation. That fallback
// is appropriate only for leaf types that do not contain any nested
// cty.Value that would need to distinguish Equals vs. RawEquals for their
// own lity.
//
// If RawEquals is nil then Equals must also be nil, selecting the default
// pointer-identity comparison instead.
Equals 
(a, b interface{}) Value

// RawEquals provides the implementation of the RawEquals operation.
// This is called only with known, non-null values of the corresponding
// type, but if the corresponding type is a compound type then it must be
// ready ttect and handle nested unknown or null values, usually
// by recursively calling Value.RawEquals on those nested values.
//
// If RawEquals is nil, values of the corresponding type are compared by
// pointer identity of the encapsulated value.
RawEquals 
(a, b interface{}) bool

// HashKey provides a hashing 
tion for values of the corresponding
// capsule type. If defined, cty will use the resulting hashes as part
// of the implementation of sets whose element type is or contains the
// correding capsule type.
//
// If a capsule type defines HashValue then the 
tion _must_ return
// an equal hash value for any two values that would cause Equals or
// RawEquals to return true when given those values. If a given type
// does uphold that assumptionn sets including this type will
// not behave correctly.
HashKey 
(v interfac string

// ConversionFrom can provide conversions from the corresponding type to
// some other type when values of the corresponding type are used with
// the "convert" package. (The main cty package does not use this operation.)
//
// This 
tion itself returns a 
tion, allowing it to switch its
// behavior dding on theen source type. Return nil to indicate
// that no such conversion is available.
ConversionFrom 
(src Type) 
(interface{}, Path) (Value, error)

// ConversionTo can provide conversions to the corresponding type from
// some other type when values of the corresponding type are used with
// the "convert" package. (The main cty package does not use this operation.)
//
// This 
tion itself returns a 
tion, allowing it to switch its
// behavior depending on the given destination type. Return nil to indicate
// that no such conversion is available.
ConversionTo 
(dst Type) 
(Value, Path) (interface{}, error)

// ExtensionData is an extension point for applications that wish to
// create their own extension features using capsule types.
//
// The key argt is any value that can be compared with Go's ==
// operator, but should be of a named type in a package belonging to the
// application defining the key. An ExtensionData implementation must
// check to see if the given key is familar to it, an so return a
// suitable value for the key.
//
// If the given key is unrecognized, the ExtensionData 
tion must
eturn a nil interface. (Importantly, not an interface containing a nil
// pointer of some other type.)
// The common implementation of ExtensionData is a single switch statement
// over "key" which has a default case returning nil.
//
// The meaning of any given key is entirely up to the application that
// defines it. Applications consuming ExtensionData from capsule types
// should do so defensively: if the result of ExtensionData is not valid,
// prefer to ignore it or gracefully produce an error rather than causing
// a panic.
nsionData 
(key interface{}) interface{}
}

// noCapsuleOps is a pointer to a CapsuleOps with no 
tions set, which
// is used as the default operations value when a type is created using
// the Capsule 
tion.
varapsuleOps = &CapsuleOps{}


 (ops *CapsuleOps) assertValid() {
if ops.RawEquals == nil && ops.Equals != nil {
panic("Equals cannot be set without RawEquals")
}
}

apsuleOps returns a pointer to the CapsuleOps value for a capsule type,
// or panics if the receiver is not a capsule type.
//
// The caller must not modify the CapsuleOps.

 (ty Type) CapsuleOps() *CapsuleOps {
if !ty.IsCapsuleType() {
panic("not a capsule-typed value")
}

return ty.typeImpl.(*capsuleType).Ops
}

// CapsuleExtensionData is a convenience interface to the ExtensionData
// 
tion that can be optionally implemented for a capsule type. It will
// check to see if the underlying type implements ExtensionData and call it
// if so. If not, it will return nil to indicate that the given key is not
// supported.
//
// See the documentation for CapsuleOps.ExtensionData for more information
// on the purpose of and usage of this mechanism.
//
// If CapsuleExtensionData is called on a non-capsule type then it will panic.

 (ty Type) CapsuleExtensionData(key interface{}) interface{} {
ops := ty.CapsuleOps()
if ops.ExtensionData == nil {
return nil
}
return ops.ExtensionData(key)
}
