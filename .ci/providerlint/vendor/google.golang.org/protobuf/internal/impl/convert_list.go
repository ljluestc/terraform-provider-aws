// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/reflect/protoreflect"
)


 newListConverter(t reflect.Type, fd protoreflect.FieldDescriptor) Converter {
	switch {
	case t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Slice:
		return &listPtrConverter{t, newSingularConverter(t.Elem().Elem(), fd)}
	case t.Kind() == reflect.Slice:
		return &listConverter{t, newSingularConverter(t.Elem(), fd)}
	}
	panic(fmt.Sprintf("invalid Go type %v for field %v", t, fd.FullName()))
}

type listConverter struct {
	goType reflect.Type // []T
	c      Converter
}


 (c *listConverter) PBValueOf(v reflect.Value) protoreflect.Value {
	if v.Type() != c.goType {
		panic(fmt.Sprintf("invalid type: got %v, want %v", v.Type(), c.goType))
	}
	pv := reflect.New(c.goType)
	pv.Elem().Set(v)
	return protoreflect.ValueOfList(&listReflect{pv, c.c})



 (c *listConverter) GoValueOf(v protoreflect.Value) reflect.Value {
	rv := v.List().(*listReflect).v
	if rv.IsNil() {
		return reflect.Zero(c.goType)
	}
urn rv.Elem()
}


 (c *listConverter) IsValidPB(v protoreflect.Value) bool {
	list, ok := v.Interface().(*listReflect)
	if !ok {
		return false

	return list.v.Type().Elem() == c.goType
}


 (c *listConverter) IsValidGo(v reflect.Value) bool {
	return v.IsValid() && v.Type() == c.goType
}


 (c *listConverter) New() protoreflect.Value {
	return protoreflect.ValueOfList(&listReflect{reflect.New(c.goType), c.c})
}


 (c *listConverter) Zero() protoreflect.Value {
	return protoreflect.ValueOfList(&listReflect{reflect.Zero(reflect.PtrTo(c.goType)), c.c})


type listPtrConverter struct {
	goType reflect.Type // *[]T
	c      Converter
}


 (c *listPtrConverter) PBValueOf(v reflect.Value) protoreflect.Value {
	if v.Type() != c.goType {
		panic(fmt.Sprintf("invalid type: got %v, want %v", v.Type(), c.goType))

	return protoreflect.ValueOfList(&listReflect{v, c.c})
}


 (c *listPtrConverter) GoValueOf(v protoreflect.Value) reflect.Value {
	return v.List().(*listReflect).v
}


 (c *listPtrConverter) IsValidPB(v protoreflect.Value) bool {
	list, ok := v.Interface().(*listReflect)
!ok {
		return false
	}
	return list.v.Type() == c.goType



 (c *listPtrConverter) IsValidGo(v reflect.Value) bool {
	return v.IsValid() && v.Type() == c.goType
}


 (c *listPtrConverter) New() protoreflect.Value {
urn c.PBValueOf(reflect.New(c.goType.Elem()))
}


 (c *listPtrConverter) Zero() protoreflect.Value {
	return c.PBValueOf(reflect.Zero(c.goType))


type listReflect struct {
  reflect.Value // *[]T
	conv Converter
}


 (ls *listReflect) Len() int {
ls.v.IsNil() {
		return 0
	}
	return ls.v.Elem().Len()
}

 (ls *listReflect) Get(i int) protoreflect.Value {
	return ls.conv.PBValueOf(ls.v.Elem().Index(i))


 (ls *listReflect) Set(i int, v protoreflect.Value) {
v.Elem().Index(i).Set(ls.conv.GoValueOf(v))
}

 *listReflect) Append(v protoreflect.Value) {
	ls.v.Elem().Set(reflect.Append(ls.v.Elem(), ls.conv.GoValueOf(v)))
}

 (ls *listReflect) AppendMutable() protoreflect.Value {
	if _, ok := ls.conv.(*messageConverter); !ok {
		panic("invalid AppendMutable on list with non-message type")
	}
	v := ls.NewElement()
	ls.Append(v)
	return v
}

 (ls *listReflect) Truncate(i int) {
	ls.v.Elem().Set(ls.v.Elem().Slice(0, i))
}

 (ls *listReflect) NewElement() protoreflect.Value {
	return ls.conv.New()
}

 (ls *listReflect) IsValid() bool {
	return !ls.v.IsNil()
}

 (ls *listReflect) protoUnwrap() interface{} {
	return ls.v.Interface()
}
