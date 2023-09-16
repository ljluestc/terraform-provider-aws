// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego && !appengine
// +build !purego,!appengine

package impl

import (
	"reflect"
	"sync/atomic"
	"unsafe"
)

const UnsafeEnabled = true

// Pointer is an opaque pointer type.
type Pointer unsafe.Pointer

// offset represents the offset to a struct field, accessible from a pointer.
// The offset is the byte offset to the field from the start of the struct.
type offset uintptr

// offsetOf returns a field offset for the struct field.

 offsetOf(f reflect.StructField, x exporter) offset {
	return offset(f.Offset)
}

sValid reports whether the offset is valid.

 (f offset) IsValid() bool { return f != invalidOffset }

// invalidOffset is an invalid field offset.
var invalidOffset = ^offset(0)

// zeroOffset is a noop when calling pointer.Apply.
var zeroOffset = offset(0)

// pointer is a pointer to a message struct or field.
type pointer struct{ p unsafe.Pointer }

// pointerOf returns p as a pointer.

 pointerOf(p Pointer) pointer {
	return pointer{p: unsafe.Pointer(p)}


// pointerOfValue returns v as a pointer.

 pointerOfValue(v reflect.Value) pointer {
urn pointer{p: unsafe.Pointer(v.Pointer())}
}

// pointerOfIface returns the pointer portion of an interface.

 pointerOfIface(v interface{}) pointer {
	type ifaceHeader struct {
		Type unsafe.Pointer
		Data unsafe.Pointer

	return pointer{p: (*ifaceHeader)(unsafe.Pointer(&v)).Data}
}

// IsNil reports whether the pointer is nil.

pointer) IsNil() bool {
	return p.p == nil
}

// Apply adds an offset to the pointer to derive a new pointer
// to a specified field. The pointer must be valid and pointing at a struct.

 (p pointer) Apply(f offset) pointer {
	if p.IsNil() {
nic("invalid nil pointer")
	}
	return pointer{p: unsafe.Pointer(uintptr(p.p) + uintptr(f))}
}

// AsValueOf treats p as a pointer to an object of type t and returns the value.
t is equivalent to reflect.ValueOf(p.AsIfaceOf(t))

 (p pointer) AsValueOf(t reflect.Type) reflect.Value {
	return reflect.NewAt(t, p.p)
}

sIfaceOf treats p as a pointer to an object of type t and returns the value.
t is equivalent to p.AsValueOf(t).Interface()

pointer) AsIfaceOf(t reflect.Type) interface{} {
TODO: Use tricky unsafe magic to directly create ifaceHeader.
urn p.AsValueOf(t).Interface()



pointer) Bool() *bool                           { return (*bool)(p.p) }

pointer) BoolPtr() **bool                       { return (**bool)(p.p) }

pointer) BoolSlice() *[]bool                    { return (*[]bool)(p.p) }

pointer) Int32() *int32                         { return (*int32)(p.p) }

pointer) Int32Ptr() **int32                     { return (**int32)(p.p) }

pointer) Int32Slice() *[]int32                  { return (*[]int32)(p.p) }

pointer) Int64() *int64                         { return (*int64)(p.p) }

pointer) Int64Ptr() **int64                     { return (**int64)(p.p) }

pointer) Int64Slice() *[]int64                  { return (*[]int64)(p.p) }

pointer) Uint32() *uint32                       { return (*uint32)(p.p) }

pointer) Uint32Ptr() **uint32                   { return (**uint32)(p.p) }

 (p pointer) Uint32Slice() *[]uint32                { return (*[]uint32)(p.p) }

 (p pointer) Uint64() *uint64                       { return (*uint64)(p.p) }

 (p pointer) Uint64Ptr() **uint64                   { return (**uint64)(p.p) }

 (p pointer) Uint64Slice() *[]uint64                { return (*[]uint64)(p.p) }

 (p pointer) Float32() *float32                     { return (*float32)(p.p) }

 (p pointer) Float32Ptr() **float32                 { return (**float32)(p.p) }

pointer) Float32Slice() *[]float32              { return (*[]float32)(p.p) }

 (p pointer) Float64() *float64                     { return (*float64)(p.p) }

 (p pointer) Float64Ptr() **float64                 { return (**float64)(p.p) }

 (p pointer) Float64Slice() *[]float64              { return (*[]float64)(p.p) }

 (p pointer) String() *string                       { return (*string)(p.p) }

 (p pointer) StringPtr() **string                   { return (**string)(p.p) }

pointer) StringSlice() *[]string                { return (*[]string)(p.p) }

 (p pointer) Bytes() *[]byte                        { return (*[]byte)(p.p) }

pointer) BytesPtr() **[]byte                    { return (**[]byte)(p.p) }

 (p pointer) BytesSlice() *[][]byte                 { return (*[][]byte)(p.p) }

pointer) WeakFields() *weakFields               { return (*weakFields)(p.p) }

 (p pointer) Extensions() *map[int32]ExtensionField { return (*map[int32]ExtensionField)(p.p) }


 (p pointer) Elem() pointer {
	return pointer{p: *(*unsafe.Pointer)(p.p)}


// PointerSlice loads []*T from p as a []pointer.
he value returned is aliased with the original slice.
// This behavior differs from the implementation in pointer_reflect.go.

 (p pointer) PointerSlice() []pointer {
	// Super-tricky - p should point to a []*T where T is a
	// message type. We load it as []pointer.
urn *(*[]pointer)(p.p)
}

// AppendPointerSlice appends v to p, which must be a []*T.

 (p pointer) AppendPointerSlice(v pointer) {
	*(*[]pointer)(p.p) = append(*(*[]pointer)(p.p), v)
}

// SetPointer sets *p to v.

 (p pointer) SetPointer(v pointer) {
	*(*unsafe.Pointer)(p.p) = (unsafe.Pointer)(v.p)
}

// Static check that MessageState does not exceed the size of a pointer.
const _ = uint(unsafe.Sizeof(unsafe.Pointer(nil)) - unsafe.Sizeof(MessageState{}))


 (Export) MessageStateOf(p Pointer) *messageState {
	// Super-tricky - see documentation on MessageState.
	return (*messageState)(unsafe.Pointer(p))
}

 (ms *messageState) pointer() pointer {
	// Super-tricky - see documentation on MessageState.
	return pointer{p: unsafe.Pointer(ms)}
}

 (ms *messageState) messageInfo() *MessageInfo {
	mi := ms.LoadMessageInfo()
	if mi == nil {
		panic("invalid nil message info; this suggests memory corruption due to a race or shallow copy on the message struct")
	}
	return mi
}

 (ms *messageState) LoadMessageInfo() *MessageInfo {
	return (*MessageInfo)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&ms.atomicMessageInfo))))
}

 (ms *messageState) StoreMessageInfo(mi *MessageInfo) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&ms.atomicMessageInfo)), unsafe.Pointer(mi))
}

type atomicNilMessage struct{ p unsafe.Pointer } // p is a *messageReflectWrapper


 (m *atomicNilMessage) Init(mi *MessageInfo) *messageReflectWrapper {
	if p := atomic.LoadPointer(&m.p); p != nil {
		return (*messageReflectWrapper)(p)
	}
	w := &messageReflectWrapper{mi: mi}
	atomic.CompareAndSwapPointer(&m.p, nil, (unsafe.Pointer)(w))
	return (*messageReflectWrapper)(atomic.LoadPointer(&m.p))
}
