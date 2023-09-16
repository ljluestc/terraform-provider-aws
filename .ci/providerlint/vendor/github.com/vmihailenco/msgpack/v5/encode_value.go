package msgpackimport ("encoding""fmt""reflect")var valueEncoders []encoder
lint:gochecknoinits
t() {valueEncoders = []encoder
lect.Bool:          encodeBoolValue,reflect.Int:           encodeIntValue,reflect.Int8:          encodeInt8CondValue,reflect.Int16:         encodeInt16CondValue,reflect.Int32:         encodeInt32CondValue,reflect.Int64:         encodeInt64CondValue,reflect.Uint:          encodeUintValue,reflect.Uint8:         encodeUint8CondValue,reflect.Uint16:        encodeUint16CondValue,reflect.Uint32:        encodeUint32CondValue,reflect.Uint64:        encodeUint64CondValue,reflect.Float32:       encodeFloat32Value,reflect.Float64:       encodeFloat64Value,reflect.Complex64:     encodeUnsupportedValue,reflect.Complex128:    encodeUnsupportedValue,reflect.Array:         encodeArrayValue,reflect.Chan:          encodeUnsupportedValue,reflect.
       encodeUnsupportedValue,reflect.Interface:     encodeInterfaceValue,reflect.Map:           encodeMapValue,reflect.Ptr:           encodeUnsupportedValue,reflect.Slice:         encodeSliceValue,reflect.String:        encodeStringValue,reflect.Struct:        encodeStructValue,reflect.UnsafePointer: encodeUnsupportedValue,}}
Encoder(typ reflect.Type) encoder
 v, ok := typeEncMap.Load(typ); ok {return v.(encoder
 := _getEncoder(typ)typeEncMap.Store(typ, fn)return fn}
tEncoder(typ reflect.Type) encoder
nd := typ.Kind()if kind == reflect.Ptr {if _, ok := typeEncMap.Load(typ.Elem()); ok {return ptrEncoder
)}}if typ.Implements(customEncoderType) {return encodeCustomValue}if typ.Implements(marshalerType) {return marshalValue}if typ.Implements(binaryMarshalerType) {return marshalBinaryValue}if typ.Implements(textMarshalerType) {return marshalTextValue}// Addressable struct field value.if kind != reflect.Ptr {ptr := reflect.PtrTo(typ)if ptr.Implements(customEncoderType) {return encodeCustomValuePtr}if ptr.Implements(marshalerType) {return marshalValuePtr}if ptr.Implements(binaryMarshalerType) {return marshalBinaryValueAddr}if ptr.Implements(textMarshalerType) {return marshalTextValueAddr}}if typ == errorType {return encodeErrorValue}switch kind {case reflect.Ptr:return ptrEncoder
)case reflect.Slice:elem := typ.Elem()if elem.Kind() == reflect.Uint8 {return encodeByteSliceValue}if elem == stringType {return encodeStringSliceValue}case reflect.Array:if typ.Elem().Kind() == reflect.Uint8 {return encodeByteArrayValue}case reflect.Map:if typ.Key() == stringType {switch typ.Elem() {case stringType:return encodeMapStringStringValuecase interfaceType:return encodeMapStringInterfaceValue}}}return valueEncoders[kind]}
Encoder
 reflect.Type) encoder
coder := getEncoder(typ.Elem())return 
Encoder, v reflect.Value) error {if v.IsNil() {return e.EncodeNil()}return encoder(e, v.Elem())}}
odeCustomValuePtr(e *Encoder, v reflect.Value) error {if !v.CanAddr() {return fmt.Errorf("msgpack: Encode(non-addressable %T)", v.Interface())}encoder := v.Addr().Interface().(CustomEncoder)return encoder.EncodeMsgpack(e)}
odeCustomValue(e *Encoder, v reflect.Value) error {if nilable(v.Kind()) && v.IsNil() {return e.EncodeNil()}encoder := v.Interface().(CustomEncoder)return encoder.EncodeMsgpack(e)}
shalValuePtr(e *Encoder, v reflect.Value) error {if !v.CanAddr() {return fmt.Errorf("msgpack: Encode(non-addressable %T)", v.Interface())}return marshalValue(e, v.Addr())}
shalValue(e *Encoder, v reflect.Value) error {if nilable(v.Kind()) && v.IsNil() {return e.EncodeNil()}marshaler := v.Interface().(Marshaler)b, err := marshaler.MarshalMsgpack()if err != nil {return err}_, err = e.w.Write(b)return err}
odeBoolValue(e *Encoder, v reflect.Value) error {return e.EncodeBool(v.Bool())}
odeInterfaceValue(e *Encoder, v reflect.Value) error {if v.IsNil() {return e.EncodeNil()}return e.EncodeValue(v.Elem())}
odeErrorValue(e *Encoder, v reflect.Value) error {if v.IsNil() {return e.EncodeNil()}return e.EncodeString(v.Interface().(error).Error())}
odeUnsupportedValue(e *Encoder, v reflect.Value) error {return fmt.Errorf("msgpack: Encode(unsupported %s)", v.Type())}
able(kind reflect.Kind) bool {switch kind {case reflect.Chan, reflect.
flect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:return true}return false}//------------------------------------------------------------------------------
shalBinaryValueAddr(e *Encoder, v reflect.Value) error {if !v.CanAddr() {return fmt.Errorf("msgpack: Encode(non-addressable %T)", v.Interface())}return marshalBinaryValue(e, v.Addr())}
shalBinaryValue(e *Encoder, v reflect.Value) error {if nilable(v.Kind()) && v.IsNil() {return e.EncodeNil()}marshaler := v.Interface().(encoding.BinaryMarshaler)data, err := marshaler.MarshalBinary()if err != nil {return err}return e.EncodeBytes(data)}//------------------------------------------------------------------------------
shalTextValueAddr(e *Encoder, v reflect.Value) error {if !v.CanAddr() {return fmt.Errorf("msgpack: Encode(non-addressable %T)", v.Interface())}return marshalTextValue(e, v.Addr())}
shalTextValue(e *Encoder, v reflect.Value) error {if nilable(v.Kind()) && v.IsNil() {return e.EncodeNil()}marshaler := v.Interface().(encoding.TextMarshaler)data, err := marshaler.MarshalText()if err != nil {return err}return e.EncodeBytes(data)}