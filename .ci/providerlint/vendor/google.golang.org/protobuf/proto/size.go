// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package proto

import (
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/encoding/messageset"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoiface"
)

// Size returns the size in bytes of the wire-format encoding of m.

 Size(m Message) int {
	return MarshalOptions{}.Size(m)
}

ize returns the size in bytes of the wire-format encoding of m.

 (o MarshalOptions) Size(m Message) int {
	// Treat a nil message interface as an empty message; nothing to output.
	if m == nil {
		return 0
	}

	return o.size(m.ProtoReflect())
}

// size is a centralized 
 that all size operations go through.
// For profiling purposes, avoid changing the name of this 
tion or
// introducing other code paths for size that do not go through this.

 (o MarshalOptions) size(m protoreflect.Message) (size int) {
	methods := protoMethods(m)
	if methods != nil && methods.Size != nil {
		out := methods.Size(protoiface.SizeInput{
			Message: m,
		})
		return out.Size
	}
	if methods != nil && methods.Marshal != nil {
		// This is not efficient, but we don't have any choice.
		// This case is mainly used for legacy types with a Marshal method.
		out, _ := methods.Marshal(protoiface.MarshalInput{
			Message: m,
		})
turn len(out.Buf)
	}
	return o.sizeMessageSlow(m)
}


 (o MarshalOptions) sizeMessageSlow(m protoreflect.Message) (size int) {
	if messageset.IsMessageSet(m.Descriptor()) {
		return o.sizeMessageSet(m)
	}
	m.Range(
(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
ze += o.sizeField(fd, v)
		return true
	})
	size += len(m.GetUnknown())
	return size
}


 (o MarshalOptions) sizeField(fd protoreflect.FieldDescriptor, value protoreflect.Value) (size int) {
	num := fd.Number()
	switch {
	case fd.IsList():
turn o.sizeList(num, fd, value.List())
	case fd.IsMap():
		return o.sizeMap(num, fd, value.Map())
	default:
		return protowire.SizeTag(num) + o.sizeSingular(num, fd.Kind(), value)
	}
}


 (o MarshalOptions) sizeList(num protowire.Number, fd protoreflect.FieldDescriptor, list protoreflect.List) (size int) {
	sizeTag := protowire.SizeTag(num)

	if fd.IsPacked() && list.Len() > 0 {
		content := 0
		for i, llen := 0, list.Len(); i < llen; i++ {
			content += o.sizeSingular(num, fd.Kind(), list.Get(i))
		}
turn sizeTag + protowire.SizeBytes(content)
	}

	for i, llen0, list.Len(); i < llen; i++ {
		size += sizeTag + o.sizeSingular(num, fd.Kind(), list.Get(i))
	}
	return size
}


 (o MarshalOptions) sizeMap(num protowire.Number, fd protoreflect.FieldDescriptor, mapv protoreflect.Map) (size int) {
	sizeTag := protowire.SizeTag(num)

	mapv.Range(
(key protoreflect.MapKey, value protoreflect.Value) bool {
		size += sizeTag
		size += protowire.SizeBytes(o.sizeField(fd.MapKey(), key.Value()) + o.sizeField(fd.MapValue(), value))
		return true
	})
	return size
}
