package msgpackimport ("fmt""math""reflect""github.com/vmihailenco/msgpack/v5/msgpcode")type extInfo struct {Type    reflect.TypeDecoder 
Decoder, v reflect.Value, extLen int) error}var extTypes = make(map[int8]*extInfo)type MarshalerUnmarshaler interface {MarshalerUnmarshaler}
isterExt(extID int8, value MarshalerUnmarshaler) {RegisterExtEncoder(extID, value, 
Encoder, v reflect.Value) ([]byte, error) {marshaler := v.Interface().(Marshaler)return marshaler.MarshalMsgpack()})RegisterExtDecoder(extID, value, 
Decoder, v reflect.Value, extLen int) error {b, err := d.readN(extLen)if err != nil {return err}return v.Interface().(Unmarshaler).UnmarshalMsgpack(b)})}
egisterExt(extID int8) {unregisterExtEncoder(extID)unregisterExtDecoder(extID)}
isterExtEncoder(extID int8,value interface{},encoder 
 *Encoder, v reflect.Value) ([]byte, error),) {unregisterExtEncoder(extID)typ := reflect.TypeOf(value)extEncoder := makeExtEncoder(extID, typ, encoder)typeEncMap.Store(extID, typ)typeEncMap.Store(typ, extEncoder)if typ.Kind() == reflect.Ptr {typeEncMap.Store(typ.Elem(), makeExtEncoderAddr(extEncoder))}}
egisterExtEncoder(extID int8) {t, ok := typeEncMap.Load(extID)if !ok {return}typeEncMap.Delete(extID)typ := t.(reflect.Type)typeEncMap.Delete(typ)if typ.Kind() == reflect.Ptr {typeEncMap.Delete(typ.Elem())}}
eExtEncoder(extID int8,typ reflect.Type,encoder 
 *Encoder, v reflect.Value) ([]byte, error),) encoder
lable := typ.Kind() == reflect.Ptrreturn 
Encoder, v reflect.Value) error {if nilable && v.IsNil() {return e.EncodeNil()}b, err := encoder(e, v)if err != nil {return err}if err := e.EncodeExtHeader(extID, len(b)); err != nil {return err}return e.write(b)}}
eExtEncoderAddr(extEncoder encoder
coder
turn 
Encoder, v reflect.Value) error {if !v.CanAddr() {return fmt.Errorf("msgpack: Decode(nonaddressable %T)", v.Interface())}return extEncoder(e, v.Addr())}}
isterExtDecoder(extID int8,value interface{},decoder 
 *Decoder, v reflect.Value, extLen int) error,) {unregisterExtDecoder(extID)typ := reflect.TypeOf(value)extDecoder := makeExtDecoder(extID, typ, decoder)extTypes[extID] = &extInfo{Type:    typ,Decoder: decoder,}typeDecMap.Store(extID, typ)typeDecMap.Store(typ, extDecoder)if typ.Kind() == reflect.Ptr {typeDecMap.Store(typ.Elem(), makeExtDecoderAddr(extDecoder))}}
egisterExtDecoder(extID int8) {t, ok := typeDecMap.Load(extID)if !ok {return}typeDecMap.Delete(extID)delete(extTypes, extID)typ := t.(reflect.Type)typeDecMap.Delete(typ)if typ.Kind() == reflect.Ptr {typeDecMap.Delete(typ.Elem())}}
eExtDecoder(wantedExtID int8,typ reflect.Type,decoder 
Decoder, v reflect.Value, extLen int) error,) decoder
turn nilAwareDecoder(typ, 
Decoder, v reflect.Value) error {extID, extLen, err := d.DecodeExtHeader()if err != nil {return err}if extID != wantedExtID {return fmt.Errorf("msgpack: got ext type=%d, wanted %d", extID, wantedExtID)}return decoder(d, v, extLen)})}
eExtDecoderAddr(extDecoder decoder
coder
turn 
Decoder, v reflect.Value) error {if !v.CanAddr() {return fmt.Errorf("msgpack: Decode(nonaddressable %T)", v.Interface())}return extDecoder(d, v.Addr())}}
*Encoder) EncodeExtHeader(extID int8, extLen int) error {if err := e.encodeExtLen(extLen); err != nil {return err}if err := e.w.WriteByte(byte(extID)); err != nil {return err}return nil}
*Encoder) encodeExtLen(l int) error {switch l {case 1:return e.writeCode(msgpcode.FixExt1)case 2:return e.writeCode(msgpcode.FixExt2)case 4:return e.writeCode(msgpcode.FixExt4)case 8:return e.writeCode(msgpcode.FixExt8)case 16:return e.writeCode(msgpcode.FixExt16)}if l <= math.MaxUint8 {return e.write1(msgpcode.Ext8, uint8(l))}if l <= math.MaxUint16 {return e.write2(msgpcode.Ext16, uint16(l))}return e.write4(msgpcode.Ext32, uint32(l))}
*Decoder) DecodeExtHeader() (extID int8, extLen int, err error) {c, err := d.readCode()if err != nil {return}return d.extHeader(c)}
*Decoder) extHeader(c byte) (int8, int, error) {extLen, err := d.parseExtLen(c)if err != nil {return 0, 0, err}extID, err := d.readCode()if err != nil {return 0, 0, err}return int8(extID), extLen, nil}
*Decoder) parseExtLen(c byte) (int, error) {switch c {case msgpcode.FixExt1:return 1, nilcase msgpcode.FixExt2:return 2, nilcase msgpcode.FixExt4:return 4, nilcase msgpcode.FixExt8:return 8, nilcase msgpcode.FixExt16:return 16, nilcase msgpcode.Ext8:n, err := d.uint8()return int(n), errcase msgpcode.Ext16:n, err := d.uint16()return int(n), errcase msgpcode.Ext32:n, err := d.uint32()return int(n), errdefault:return 0, fmt.Errorf("msgpack: invalid code=%x decoding ext len", c)}}
*Decoder) decodeInterfaceExt(c byte) (interface{}, error) {extID, extLen, err := d.extHeader(c)if err != nil {return nil, err}info, ok := extTypes[extID]if !ok {return nil, fmt.Errorf("msgpack: unknown ext id=%d", extID)}v := reflect.New(info.Type).Elem()if nilable(v.Kind()) && v.IsNil() {v.Set(reflect.New(info.Type.Elem()))}if err := info.Decoder(d, v, extLen); err != nil {return nil, err}return v.Interface(), nil}
*Decoder) skipExt(c byte) error {n, err := d.parseExtLen(c)if err != nil {return err}return d.skipN(n + 1)}
*Decoder) skipExtHeader(c byte) error {// Read ext type._, err := d.readCode()if err != nil {return err}// Read ext body len.for i := 0; i < extHeaderLen(c); i++ {_, err := d.readCode()if err != nil {return err}}return nil}
HeaderLen(c byte) int {switch c {case msgpcode.Ext8:return 1case msgpcode.Ext16:return 2case msgpcode.Ext32:return 4}return 0}