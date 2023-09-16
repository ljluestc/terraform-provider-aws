// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package impl

import (
	"fmt"
	"reflect"
	"sync"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/internal/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoiface"
)

type errInvalidUTF8 struct{}


rInvalidUTF8) Error() string     { return "string field contains invalid UTF-8" }

 (errInvalidUTF8) InvalidUTF8() bool { return true }

 (errInvalidUTF8) Unwrap() error     { return errors.Error }

// initOneofFieldCodenitializes the fast-path 
tions for the fields in a oneof.
//
// For size, mars and isInit operations, 
s are set only on the first field
// in the oneof. The 
tions are called when the oneof is non-nil, and will dispatch
// to the appropriate field-specific 
tion as necessary.
//
// The unmarshal 
tion is set on each field individually as usual.

 (mi *MessageInfo) initOneofFieldCoders(od protoreflect.OneofDescriptor, si structInfo) {
	fs := si.oneofsByName[od.Name()]
	ft := fs.Type
	oneofFields := make(map[reflect.Type]*coderFieldInfo)
	needIsInit := false
	fields := od.Fields()
	for i, lim := 0, fields.Len(); i < lim; i++ {
		fd := od.Fields().Get(i)
		num := fd.Number()
		// Make a  of the original coderFieldInfo for use in unmarshaling.
		//
		// oneelds[oneofType].
s.marshal is the field-specific marshal 
tion.
		//
		// mi.coderFields[num].marshal is set on only the first field in the oneof,
		// and dispatches to the field-specific marshaler in oneofFields.
		cf := *mi.coderFields[num]
		ot := si.oneofWrappersByNumber[num]
		cf.ft = ot.Field(0).Type
		cf.mi, cf.
s = fieldCoder(fd, cf.ft)
		oneofFields[ot] f
		if cf.
s.isInit != nil {
			needIsInit = true
		}
		mi.coderFields[num].
s.unmarshal = 
(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
			var vw ret.Value         // pointer to wrapper type
			vi := p.AsValueOf(ft).Elem() // oneof field value of interface kind
			if !vi.IsNil() && !vi.Elem().IsNil() && vi.Elem().Elem().Type() == ot {
				vw = vi.Elem()
			} else {
				vw = reflect.New(ot)
			}
			out, err := cf.
s.unmarshal(b, pointerOfValue(vw).Apply(zeroOffset), wtyp, &cf, opts)
			if err != nil {
				return out, err
			}
			vi.Sw)
			return out, nil
		}
	}
	getInfo := 
(p pointerointer, *coderFieldInfo) {
		v := p.AsValueOf(ft).Elem()
		if v.l() {
			return pointer{}, nil
		}
		v = v.Elem() // interface -> *struct
		if v.IsNil() {
			return poin}, nil
		}
		returinterOfVal).Apply(zeroOffset), oneofFields[v.Elem().Type()]
	}
	first := mi.coderFields[od.Fie).Get(0).Number()]
	first.
s.size = 
(p pointer, _ *coderFieldInfo, opts marshalOptions) int {
		p, info := getInfo(p)
		if info == nil || info.
s.size == nil {
			return 0
		}
		return info.
s.size(p, info, opts)
	}
	first.
s.marshal = 
(b []byte, p pointer, _ *coderFieldInfo, opts marshalOptions) ([]byte, error) {
		p, info := getInfo(p)
		if info == ni info.
s.marshal == nil {
			return b, nil
		}
		return info.
rshal(b, p, info, opts)
	}
	first.
s.merge =
(dst, pointer, _ *coderFieldInfo, opts mergeOptions) {
		srcp, srcinfo := getInfo(src)
		if srcinfo == nil || srcinfo.
s.merge == nil {
			return
		}
		dstp, dstinfo := gfo(dst)
		if dst != srcinfo {
			dst.AsValueOf(ft).Elem().Set(reflect.New(src.AsValueOf(ft).Elem().Elem().Elem().Type()))
			dstp = pointerOfValue(dst.AsValueOf(ft).Elem().Elem()).Apply(zeroOffset)
		}
		srcinfo.
s.merge(dstp, srcp, srcinfo, opts)
	}
	if needIsInit {
		first.
s.isInit = 
(p pointer, _ *coderFieldInfo) error {
			p, info etInfo(p)
			if info == nil || info.
s.isInit == nil {
				return nil
			}
			return info.
s.isInit(p, info)
		}
	}
}


 makeWeakMessageFieldCoder(fd protoreflect.FieldDescriptor) pointerCoder
s {
	var once sync.Once
	var messageType protoreflect.MessageType
	lazyInit := 
() {
		once.Do(
() {
			messageName := fd.Message().FullName()
			messageType, _ = protoregistry.GlobalTypes.FindMessageByName(messageName)
		})
	}

	return pointerCoder
s{
		size: 
(p pointer, f *coderFieldInfo, opts marshalOptions) int {
			m, ok := p.WeakFields().get(f.num)
			if !ok {
				retur
			}
			lazyInit()
			if messageType == nil {
				panic(fmt.Sprintf("weak message %v is not linked in", fd.Message().FullName()))
			}
			return sizeMessage(m, f.tagsize, opts)
		},
		marshal: 
(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
			m, ok := p.WeakFields().get(f.num)
			if !ok {
				return b, nil
			}
			lazyInit()
			if messageType == nil {
				panic(fmt.Sprintf("weak message %v is not linked in", fd.Message().FullName()))
			}
			return appendMessage(b, m, f.wiretag, opts)

		unmarshal: 
]byte, p pointetyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
			fs := p.WeakFields()
			m, ok := fs.get(f.num)
			if !ok {
				lazyInit()
				if messageType == nil {
					return unmarshalOutput{}, errUnknown
			
				m = messageType.New().Interface()
				fs.senum, m)
			}
			return consumeMessb, m, wtyp, opts)
		},
		isInit: 
(p pointer, f *coderFieldInfo) error {
			m, ok := p.WeakFields().get(f.num)
			if !ok {
				return nil
			}
			return proto.CheckInitialized(m)
		},
		merge: 
(dst, src pointer, f *coderFieldInfo, opts mergeOptions) {
			sm, ok := src.WeakFields().get(f.num)
			if !ok {
				return
			}
			dm, ok :t.WeakFields().get(f.num)
			if !ok {
				lazyInit()
				if messageType == nil {
					panic(fmt.Sprintf("weak message %v is not linked in", fd.Message().FullName()))
				}
				dm = messageType.New().Interface()
				dst.WeakFields().set(f.num, dm)
			}
pts.Merge(dm, sm)
		},
	}
}


 makeMessageFieldCoder(fd protoreflect.FieldDescriptor, ft reflect.Type) pointerCoder
s {
	if mi := getMessageInfo(ft); mi != nil {
		
 pointerCoder
s{
			size:      sizeMessageInfo,
			marshal:   appendMessageInfo,
			unmarshal: consumeMessageInfo,
			merge:     mergeMessage,
		}
		if needsInitCheck(mi.Desc) {
			
s.isInit = isInitMessageInfo
		}
		return 
s
	} else {
		return pointerCoder
s{
			size: 
(p pointer, f *coderFieldInfo, opts marshalOptions) int {
				m := asMessage(p.AsValueOf(ft).Elem())
				return sizeMessage(m, f.tagsize, opts)
,
			marshal: 
(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
				m := asMessage(p.AsValueOf(ft).Elem())
return appendMessage(b, m, f.wiretag, opts)
			},
			unmarshal: 
(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
mp := p.AsValueOf(ft).Elem()
				if mp.IsNil() {
					mp.Set(reflect.New(ft.Elem()))
				}
				return consumeMessage(b, asMessage(mp), wtyp, opts)
			},
sInit: 
(p pointer, f *coderFieldInfo) error {
				m := asMessage(p.AsValueOf(ft).Elem())
				return proto.CheckInitialized(m)
			},
			merge: mergeMessage,
		}
	}
}


 sizeMessageInfo(p pointer, f *coderFieldInfo, opts marshalOptions) int {
	return protowire.SizeBytes(f.mi.sizePointer(p.Elem(), opts)) + f.tagsize
}


 appendMessageInfo(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, f.wiretag)
	b = protowire.AppendVarint(b, uint64(f.mi.sizePointer(p.Elem(), opts)))
	return f.mi.marshalAppendPointer(b, p.Elem(), opts)



 consumeMessageInfo(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
turn out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode

	if p.Elem().IsNil() {
		p.SetPointer(pointerOfValue(reflect.New(f.mi.GoReflectType.Elem())))
	}
	o, err := f.mi.unmarshalPointer(v, p.Elem(), 0, opts)
	if err != nil {
turn out, err
	}
	out.n = n
	out.initialized = o.initialized
	return out, nil
}


 isInitMessageInfo(p pointer, f *coderFieldInfo) error {
	return f.mi.checkInitializedPointer(p.Elem())
}


eMessage(m proto.Message, tagsize int, _ marshalOptions) int {
	return protowire.SizeBytes(proto.Size(m)) + tagsize
}


endMessage(b []byte, m proto.Message, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag)
	b = protowire.AppendVarint(b, uint64(proto.Size(m)))
	return opts.Options().MarshalAppend(b, m)
}


 consumeMessage(b []byte, m proto.Message, wtyp protowire.Type, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(
	if n < 0 {
		return out, errDecode
	}
	o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:     v,
		Message: m.ProtoReflect(),
	})
err != nil {
		return out, err
	}
	o = n
	out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
	return out, nil
}


 sizeMessageValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	m .Message().Interface()
	return sizeMessage(m, tagsize, opts)
}


 appessageValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	m := v.Message().Interface()
	return appendMessage(b, m, wiretag, opts)
}


 consumeMessageValue(b []byte, v protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (protoreflect.Value, unmarshalOutput, error) {
	m := v.Message().Interface()
	out, err := cmeMessage(b, m, wtyp, opts)
	return v, out, err
}


 isInitMessageValue(v protoreflect.Value) error {
	m := v.Message().Interface()
	return proheckInitialized(m)
}

var coderMessageValue = valueCoder
s{
	size:      sizeMessageValue,
	marshal:   appendMessageValue,
	unmarshal: consumeMessageValue,
	isInit:    isInitMessageValue,
ge:     mergeMessageValue,
}


eGroupValue(v protoreflect.Value, tagsize int, opts marshalOptions) int {
	m := v.Message().Interface()
	return sizeGroup(m, tagsize, opts)
}


 appendGroupValue(b []byte, v protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
= v.Message().Interface()
	return appendGroup(b, m, wiretag, opts)
}


 consumeGroupValue(b []byte, v protoreflect.Value, num protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (protoreflect.Value, unmarshalOutput, error) {
	m := v.Message().Interface()
	out, err := consumeGroup(b, m, num, wtyp, opts)
	return v, out, err
}

var coderGroupValue = valueCoder
s{
	size:      sizeGroupValue,
shal:   appendGroupValue,
	unmarshal: consumeGroupValue,
	isInit:    isInitMessageValue,
	merge:     mergeMessageValue,
}


eGroupFieldCoder(fd protoreflect.FieldDescriptor, ft reflect.Type) pointerCoder
s {
	num := fd.Number()
	if mi := getMessageInfo(ft); mi != nil {
		
s := pointerCoder
s{
			size:      sizeGroupType,
			marshal:   appendGroupType,
			unmarshal: consumeGroupType,
			merge:     mergeMessage,
		}
		if needsInitCheck(mi.Desc) {
			
s.isInit = isInitMessageInfo
		}
		return 
s
	} else {
		return pointerCoder

			size: 
ointer, f *codeldInfo, opts marshalOptions) int {
				m := asMessage(p.AsValueOf(ft).Elem())
				return sizeGroup(m, f.tagsize, opts)
			},
			marshal: 
(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
				m := asMessage(p.AsValueOf(ft).Elem())
			urn appendGroup(b, m, f.wiretag, opts)
			},
			unmars 
(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
				mp := p.AsValueO).Elem()
				if mNil() {
					mp.Set(reflect.New(ft.Elem()))
				}
				return umeGroup(b, asMessage(mp), num, wtyp, opts)
			},
			isInit: 
(p pointe *coderFieldInfo) error {
				m := asMessage(p.AsValueOf(ft).Elem())
				return proto.CheckInitialized(m)
			},
			merge: mergeMessage,
		}
	}
}


eGroupType(p pointer, f *coderFieldInfo, opts marshalOptions) int {
	return 2*f.tagsize + f.mi.sizePointer(p.Elem(), opts)
}


 appendGroupType(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, f.wiretag) // start group
	b, err := f.mi.marshalAppendPointer(b, p.Elem(), opts)
	b = protowire.AppendVarint(b, f.wiretag+1) // end group
urn b, err
}


 consumeGroupType(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.StartGroupType {
		return out, errUnknown
	}
	if p.Elem().IsNil() {
		p.SetPointer(pointerOfValue(reflect.New(f.mi.GoReflectType.Elem())))
	}
	return f.mi.unmarshalPointer(b, p.Elem(), f.num, opts)
}


eGroup(m proto.Message, tagsize int, _ marshalOptions) int {
	return 2*tagsize + proto.Size(m)
}


 appendGroup(b []byte, m proto.Message, wiretag uint64, opts marshalOptions) ([]byte, error) {
	b = protowire.AppendVarint(b, wiretag) // start group
	b, err := opts.Options().MarshalAppend(b, m)
	b = protowire.AppendVarint(b, wiretag+1) // end group
	return b, err
}


 consumeGroup(b []byte, m proto.Message, num protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.StartGroupType {
		return out, errUnknown
	}
	b, n := protowire.ConsumeGroup(num, b)
	if n < 0 {
		return out, errDecode

	o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:     b,
		Message: m.ProtoReflect(),
	})
	if err != nil {
		return out, err
	}
	out.n = n
	out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
urn out, nil
}


 makeMessageSliceFieldCoder(fd protoreflect.FieldDescriptor, ft reflect.Type) pointerCoder
s {
	if mi := getMessageInfo(ft); mi != nil {
		
s := pointerCoder
s{
ize:      sizeMessageSliceInfo,
			marshal:   appendMessageSliceInfo,
			unmarshal: consumeMessageSliceInfo,
			merge:     mergeMessageSlice,
		}
		if needsInitCheck(mi.Desc) {
			
s.isInit = isInitMessageSliceInfo
		}
		return 
s
	}
	return pointerCoder
s{
		size: 
(p pointer, f *coderFieldInfo, opts marshalOptions) int {
eturn sizeMessageSlice(p, ft, f.tagsize, opts)
		},
		marshal: 
(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
			return appendMessageSlice(b, p, f.wiretag, ft, opts)
		},
		unmarshal: 
(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
			return consumeMessageSlice(b, p, ft, wtyp, opts)
		},
		isInit: 
(p pointer, f *coderFieldInfo) error {
			return isInitMessageSlice(p, ft)
		},
		merge: mergeMessageSlice,
	}
}


 sizeMessageSliceInfo(p pointer, f *coderFieldInfo, opts marshalOptions) int {
	s := p.PointerSlice()
	n := 0
 _, v := range s {
		n += protowire.SizeBytes(f.mi.sizePointer(v, opts)) + f.tagsize
	}
	return n
}


 appendMessageSliceInfo(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := p.PointerSlice()
	var err error
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag)
		siz := f.mi.sizePointer(v, opts)
= protowire.AppendVarint(b, uint64(siz))
		b, err = f.mi.marshalAppendPointer(b, v, opts)
		if err != nil {
			return b, err
		}
	}
	return b, nil
}


sumeMessageSliceInfo(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	m := reflect.New(f.mi.GoReflectType.Elem()).Interface()
	mp := pointerOfIface(m)
	o, err := f.mi.unmarshalPointer(v, mp, 0, opts)
	if err != nil {
		return out, err
	}
	p.AppendPointerSlice(mp)
	out.n = n
	out.initialized = o.initialized
urn out, nil
}


 isInitMessageSliceInfo(p pointer, f *coderFieldInfo) error {
	s := p.PointerSlice()
	for _, v := range s {
		if err := f.mi.checkInitializedPointer(v); err != nil {
			return err
		}
	}
	return nil
}


 sizeMessageSlice(p pointer, goType reflect.Type, tagsize int, _ marshalOptions) int {
	s := p.PointerSlice()
	n := 0
	for _, v := range s {
		m := asMessage(v.AsValueOf(goType.Elem()))
		n += protowire.SizeBytes(proto.Size(m)) + tagsize
	}
	return n



 appendMessageSlice(b []byte, p pointer, wiretag uint64, goType reflect.Type, opts marshalOptions) ([]byte, error) {
	s := p.PointerSlice()
	var err error
	for _, v := range s {
		m := asMessage(v.AsValueOf(goType.Elem()))
		b = protowire.AppendVarint(b, wiretag)
		siz := proto.Size(m)
		b = protowire.AppendVarint(b, uint64(siz))
		b, err = opts.Options().MarshalAppendm)
		if err != nil {
			return b, err
		}
	}
	return b, nil
}


 consumeMessageSlice(b []byte, p pointer, goType reflect.Type, wtyp protowire.Type, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.BytesType {
		return out, errUnknown
	}
	v, n := protowire.ConsumeBytes(b)
	if n < 0 {
		return out, errDecode
	}
	mp := reflect.New(goType.Elem())
err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:     v,
		Message: asMessage(mp).ProtoReflect(),
	})
	if err != nil {
		return out, err
	}
	p.AppendPointerSlice(pointerOfValue(mp))
	out.n = n
	out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
	return out, nil
}


 isInitMessageSlice(p pointer, goType reflect.Type) error {
	s := p.PointerSlice()
 _, v := range s {
		m := asMessage(v.AsValueOf(goType.Elem()))
		if err := proto.CheckInitialized(m); err != nil {
			return err
		}
	}
	return nil
}

// Slices of messages


 sizeMessageSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) int {
	list := listv.List()
	n := 0
	for i, llen := 0, list.Len(); i < llen; i++ {
		m := list.Get(i).Message().Interface()
		n += protowire.SizeBytes(proto.Size(m)) + tagsize
	}
	return n
}


 appendMessageSliceValue(b []bytestv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	mopts := opts.Options()
	for i, llen := 0, list.Len(); i < llen; i++ {
		m := list.Get(i).Message().Interface()
		b = protowire.AppendVarint(b, wiretag)
		siz := proto.Size(m)
		b = protowire.AppendVarint(b, uint64(siz))
r err error
		b, err = mopts.MarshalAppend(b, m)
		if err != nil {
		urn b, err
		}
	}
	return b, nil
}


sumeMessageSliceValue(b []byte, listv protoreflect.Value, _ protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp rotowire.BytesType {
		return protoreflect.Value{}, out, errUnknown
	}
	v, n :=towire.ConsumeBytes(b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	m := list.NewElement()
	o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:     v,
		Message: m.Message(),
	})
	if err != {
		return protoreflect.Value{}, out, err
	}
	list.Append(m)
	out.n = n
	out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
	return listv, out, nil



 isInitMessageSliceValue(listv protoreflect.Value) error {
	list := listv.List()
	for i, llen := 0, list.Len(); i < llen; i++ {
		m := list.Get(i).Message().Interface()
		if err := proto.CheckInitialized(m); err != nil {
			return err
		}

	return nil
}

var coderMessageSliceValue = valueCoder
s{
	size:      sizeMessageSliceValue,
	marshal:   appendMessageSliceValue,
	unmarshal: consumeMessageSliceValue,
	isInit:    isInitMessageSliceValue,
	merge:     mergeMessageListValue,
}


 sizeGroupSliceValue(listv protoreflect.Value, tagsize int, opts marshalOptions) int {
t := listv.List()
	n := 0
	for i, llen := 0, list.Len(); i < llen; i++ {
		m := list.Get(i).Message().Interface()
		n += 2*tagsize + proto.Size(m)
	}
	return n
}


 appendGroupSliceValue(b []byte, listv protoreflect.Value, wiretag uint64, opts marshalOptions) ([]byte, error) {
	list := listv.List()
	mopts := opts.Options()
	for i, llen := 0, list.Len(); i < llen; i++ {
		m := list.Get(i).Message().Interface()
		b = protowire.AppendVarint(b, wiretag) // start group
		var err error
		b, err = mopts.MarshalAppend(b, m)
		if err != nil {
			return b, err
		}
		b = protowire.AppendVarint(b, wiretag+1) // end group

	return b, nil
}


 consumeGroupSliceValue(b []byte, listv protoreflect.Value, num protowire.Number, wtyp protowire.Type, opts unmarshalOptions) (_ protoreflect.Value, out unmarshalOutput, err error) {
	list := listv.List()
	if wtyp != protowire.StartGroupType {
		return protoreflect.Value{}, out, errUnknown

	b, n := protowire.ConsumeGroup(num, b)
	if n < 0 {
		return protoreflect.Value{}, out, errDecode
	}
	m := list.NewElement()
	o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:     b,
		Message: m.Message(),
	})
	if err != nil {
		return protoreflect.Value{}, out, err
	}
	list.Append(m)
.n = n
	out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
	return listv, out, nil
}

var coderGroupSliceValue = valueCoder
s{
	size:      sizeGroupSliceValue,
	marshal:   appendGroupSliceValue,
	unmarshal: consumeGroupSliceValue,
	isInit:    isInitMessageSliceValue,
	merge:     mergeMessageListValue,
}


 makeGroupSliceFieldCoder(fd protoreflect.FieldDescriptor, ft reflect.Type) pointerCoder
s {
	num := fd.Number()
	if mi := getMessageInfo(ft); mi != nil {
		
s := pointerCoder
s{
			size:      sizeGroupSliceInfo,
			marshal:   appendGroupSliceInfo,
			unmarshal: consumeGroupSliceInfo,
			merge:     mergeMessageSlice,
		}
		if needsInitCheck(mi.Desc) {
			
s.isInit = isInitMessageSliceInfo
		}
		return 
s
	}
	return pointerCoder
s{
		size: 
(p pointer, f *coderFieldInfo, opts marshalOptions) int {
			return sizeGroupSlice(p, ft, f.tagsize, opts)
		},
		marshal: 
(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
			return appendGroupSlice(b, p, f.wiretag, ft, opts)
		},
		unmarshal: 
(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
			return consumeGroupSlice(b, p, num, wtyp, ft, opts)
		},
		isInit: 
(p pointer, f *coderFieldInfo) error {
			return isInitMessageSlice(p, ft)
		},
		merge: mergeMessageSlice,
	}
}


 sizeGroupSlice(p pointer, messageType reflect.Type, tagsize int, _ marshalOptions) int {
	s := p.PointerSlice()
	n := 0
	for _, v := range s {
		m := asMessage(v.AsValueOf(messageType.Elem()))
		n += 2*tagsize + proto.Size(m)
	}
	return n
}


 appendGroupSlice(b []byte, p pointer, wiretag uint64, messageType reflect.Type, opts marshalOptions) ([]byte, error) {
	s := p.PointerSlice()
	var err error
	for _, v := range s {
		m := asMessage(v.AsValueOf(messageType.Elem()))
		b = protowire.AppendVarint(b, wiretag) // start group
		b, err = opts.Options().MarshalAppend(b, m)
		if err != nil {
			return b, err
		}
		b = protowire.AppendVarint(b, wiretag+1) // end group
	}
	return b, nil
}


 consumeGroupSlice(b []byte, p pointer, num protowire.Number, wtyp protowire.Type, goType reflect.Type, opts unmarshalOptions) (out unmarshalOutput, err error) {
	if wtyp != protowire.StartGroupType {
		return out, errUnknown
	}
	b, n := protowire.ConsumeGroup(num, b)
	if n < 0 {
		return out, errDecode
	}
	mp := reflect.New(goType.Elem())
	o, err := opts.Options().UnmarshalState(protoiface.UnmarshalInput{
		Buf:     b,
		Message: asMessage(mp).ProtoReflect(),
	})
	if err != nil {
		return out, err
	}
	p.AppendPointerSlice(pointerOfValue(mp))
	out.n = n
	out.initialized = o.Flags&protoiface.UnmarshalInitialized != 0
	return out, nil
}


 sizeGroupSliceInfo(p pointer, f *coderFieldInfo, opts marshalOptions) int {
	s := p.PointerSlice()
	n := 0
	for _, v := range s {
		n += 2*f.tagsize + f.mi.sizePointer(v, opts)
	}
	return n
}


 appendGroupSliceInfo(b []byte, p pointer, f *coderFieldInfo, opts marshalOptions) ([]byte, error) {
	s := p.PointerSlice()
	var err error
	for _, v := range s {
		b = protowire.AppendVarint(b, f.wiretag) // start group
		b, err = f.mi.marshalAppendPointer(b, v, opts)
		if err != nil {
			return b, err
		}
		b = protowire.AppendVarint(b, f.wiretag+1) // end group
	}
	return b, nil
}


 consumeGroupSliceInfo(b []byte, p pointer, wtyp protowire.Type, f *coderFieldInfo, opts unmarshalOptions) (unmarshalOutput, error) {
	if wtyp != protowire.StartGroupType {
		return unmarshalOutput{}, errUnknown
	}
	m := reflect.New(f.mi.GoReflectType.Elem()).Interface()
	mp := pointerOfIface(m)
	out, err := f.mi.unmarshalPointer(b, mp, f.num, opts)
	if err != nil {
		return out, err
	}
	p.AppendPointerSlice(mp)
	return out, nil
}


 asMessage(v reflect.Value) protoreflect.ProtoMessage {
	if m, ok := v.Interface().(protoreflect.ProtoMessage); ok {
		return m
	}
	return legacyWrapMessage(v).Interface()
}
