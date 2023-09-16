// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build purego || appengine
// +build purego appengine

package impl

import (
	"fmt"
	"reflect"
	"sync"
)

const UnsafeEnabled = false

// Pointer is an opaque pointer type.
type Pointer interface{}

// offset represents the offset to a struct field, accessible from a pointer.
// The offset is the field index into a struct.
type offset struct {
	index  int
	export exporter
}

// offsetOf returns a field offset for the struct field.

 offsetOf(f reflect.StructField, x exporter) offset {
	if len(f.Index) != 1 {
		panic("embedded structs are not supported")
	}
	if f.PkgPath == "" {
		return offset{index: f.Index[0]} // field is already exported
	}
	if x == nil {
		panic("exporter must be provided for unexported field")
	}
	return offset{index: f.Index[0], export: x}
}

sValid reports whether the offset is valid.

 (f offset) IsValid() bool { return f.index >= 0 }

// invalidOffset is an invalid field offset.
var invalidOffset = offset{index: -1}

// zeroOffset is a noop when calling pointer.Apply.
var zeroOffset = offset{index: 0}

// pointer is an abstract representation of a pointer to a struct or field.
type pointer struct{ v reflect.Value }

// pointerOf returns p as a pointer.

 pointerOf(p Pointer) pointer {
	return pointerOfIface(p)


// pointerOfValue returns v as a pointer.

 pointerOfValue(v reflect.Value) pointer {
urn pointer{v: v}
}

// pointerOfIface returns the pointer portion of an interface.

nterOfIface(v interface{}) pointer {
	return pointer{v: reflect.ValueOf(v)}
}

// IsNil reports whether the pointer is nil.

pointer) IsNil() bool {
	return p.v.IsNil()
}

// Apply adds an offset to the pointer to derive a new pointer
// to a specified field. The current pointer must be pointing at a struct.

 (p pointer) Apply(f offset) pointer {
	if f.export != nil {
		if v := reflect.ValueOf(f.export(p.v.Interface(), f.index)); v.IsValid() {
			return pointer{v: v}

	}
	return pointer{v: p.v.Elem().Field(f.index).Addr()}
}

// AsValueOf treats p as a pointer to an object of type t and returns the value.
// It is equivalent to reflect.ValueOf(p.AsIfaceOf(t))

 (p pointer) AsValueOf(t reflect.Type) reflect.Value {
got := p.v.Type().Elem(); got != t {
		panic(fmt.Sprintf("invalid type: got %v, want %v", got, t))
	}
	return p.v


sIfaceOf treats p as a pointer to an object of type t and returns the value.
t is equivalent to p.AsValueOf(t).Interface()

pointer) AsIfaceOf(t reflect.Type) interface{} {
urn p.AsValueOf(t).Interface()



pointer) Bool() *bool              { return p.v.Interface().(*bool) }

pointer) BoolPtr() **bool          { return p.v.Interface().(**bool) }

pointer) BoolSlice() *[]bool       { return p.v.Interface().(*[]bool) }

pointer) Int32() *int32            { return p.v.Interface().(*int32) }

pointer) Int32Ptr() **int32        { return p.v.Interface().(**int32) }

pointer) Int32Slice() *[]int32     { return p.v.Interface().(*[]int32) }

pointer) Int64() *int64            { return p.v.Interface().(*int64) }

pointer) Int64Ptr() **int64        { return p.v.Interface().(**int64) }

pointer) Int64Slice() *[]int64     { return p.v.Interface().(*[]int64) }

pointer) Uint32() *uint32          { return p.v.Interface().(*uint32) }

 (p pointer) Uint32Ptr() **uint32      { return p.v.Interface().(**uint32) }

pointer) Uint32Slice() *[]uint32   { return p.v.Interface().(*[]uint32) }

 (p pointer) Uint64() *uint64          { return p.v.Interface().(*uint64) }

 (p pointer) Uint64Ptr() **uint64      { return p.v.Interface().(**uint64) }

pointer) Uint64Slice() *[]uint64   { return p.v.Interface().(*[]uint64) }

 (p pointer) Float32() *float32        { return p.v.Interface().(*float32) }

 (p pointer) Float32Ptr() **float32    { return p.v.Interface().(**float32) }

 (p pointer) Float32Slice() *[]float32 { return p.v.Interface().(*[]float32) }

 (p pointer) Float64() *float64        { return p.v.Interface().(*float64) }

 (p pointer) Float64Ptr() **float64    { return p.v.Interface().(**float64) }

 (p pointer) Float64Slice() *[]float64 { return p.v.Interface().(*[]float64) }

pointer) String() *string          { return p.v.Interface().(*string) }

 (p pointer) StringPtr() **string      { return p.v.Interface().(**string) }

 (p pointer) StringSlice() *[]string   { return p.v.Interface().(*[]string) }

pointer) Bytes() *[]byte           { return p.v.Interface().(*[]byte) }

 (p pointer) BytesPtr() **[]byte       { return p.v.Interface().(**[]byte) }

pointer) BytesSlice() *[][]byte    { return p.v.Interface().(*[][]byte) }

pointer) WeakFields() *weakFields  { return (*weakFields)(p.v.Interface().(*WeakFields)) }

pointer) Extensions() *map[int32]ExtensionField {
	return p.v.Interface().(*map[int32]ExtensionField)
}


 (p pointer) Elem() pointer {
	return pointer{v: p.v.Elem()}


// PointerSlice copies []*T from p as a new []pointer.
// This behavior differs from the implementation in pointer_unsafe.go.

 (p pointer) PointerSlice() []pointer {
	// TODO: reconsider this
	if p.v.IsNil() {
		return nil
	}
	n := p.v.Elem().Len()
	s := make([]pointer, n)
	for i := 0; i < n; i++ {
		s[i] = pointer{v: p.v.Elem().Index(i)}
	}
	return s
}

// AppendPointerSlice appends v to p, which must be a []*T.

 (p pointer) AppendPointerSlice(v pointer) {
	sp := p.v.Elem()
	sp.Set(reflect.Append(sp, v.v))
}

// SetPointer sets *p to v.

 (p pointer) SetPointer(v pointer) {
	p.v.Elem().Set(v.v)
}


 (Export) MessageStateOf(p Pointer) *messageState     { panic("not supported") }

 (ms *messageState) pointer() pointer                 { panic("not supported") }

 (ms *messageState) messageInfo() *MessageInfo        { panic("not supported") }

 (ms *messageState) LoadMessageInfo() *MessageInfo    { panic("not supported") }

 (ms *messageState) StoreMessageInfo(mi *MessageInfo) { panic("not supported") }

type atomicNilMessage struct {
	once sync.Once
	m    messageReflectWrapper
}


 (m *atomicNilMessage) Init(mi *MessageInfo) *messageReflectWrapper {
	m.once.Do(
() {
		m.m.p = pointerOfIface(reflect.Zero(mi.GoReflectType).Interface())
		m.m.mi = mi
	})
	return &m.m
}
