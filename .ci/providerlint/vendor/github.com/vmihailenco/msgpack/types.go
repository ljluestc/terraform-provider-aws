package msgpackimport ("reflect""sync")var errorType = reflect.TypeOf((*error)(nil)).Elem()var customEncoderType = reflect.TypeOf((*CustomEncoder)(nil)).Elem()var customDecoderType = reflect.TypeOf((*CustomDecoder)(nil)).Elem()var marshalerType = reflect.TypeOf((*Marshaler)(nil)).Elem()var unmarshalerType = reflect.TypeOf((*Unmarshaler)(nil)).Elem()type encoder

coder, reflect.Value) errortype decoder

coder, reflect.Value) errorvar typEncMap = make(map[reflect.Type]encoder
 typDecMap = make(map[reflect.Type]decoder
Register registers encoder and decoder 
s for a value.// This is low level API and in most cases you should prefer implementing// Marshaler/CustomEncoder and Unmarshaler/CustomDecoder interfaces.
ister(value interface{}, enc encoder
c decoder
yp := reflect.TypeOf(value)if enc != nil {typEncMap[typ] = enc}if dec != nil {typDecMap[typ] = dec}}//------------------------------------------------------------------------------var structs = newStructCache(false)var jsonStructs = newStructCache(true)type structCache struct {mu sync.RWMutexm  map[reflect.Type]*fieldsuseJSONTag bool}
StructCache(useJSONTag bool) *structCache {return &structCache{m: make(map[reflect.Type]*fields),useJSONTag: useJSONTag,}}
*structCache) Fields(typ reflect.Type) *fields {m.mu.RLock()fs, ok := m.m[typ]m.mu.RUnlock()if ok {return fs}m.mu.Lock()fs, ok = m.m[typ]if !ok {fs = getFields(typ, m.useJSONTag)m.m[typ] = fs}m.mu.Unlock()return fs}//------------------------------------------------------------------------------type field struct {name      stringindex     []intomitEmpty boolencoder   encoder
der   decoder

*field) value(v reflect.Value) reflect.Value {return fieldByIndex(v, f.index)}
*field) Omit(strct reflect.Value) bool {return f.omitEmpty && isEmptyValue(f.value(strct))}
*field) EncodeValue(e *Encoder, strct reflect.Value) error {return f.encoder(e, f.value(strct))}
*field) DecodeValue(d *Decoder, strct reflect.Value) error {return f.decoder(d, f.value(strct))}//------------------------------------------------------------------------------type fields struct {Table   map[string]*fieldList    []*fieldAsArray boolhasOmitEmpty bool}
Fields(numField int) *fields {return &fields{Table: make(map[string]*field, numField),List:  make([]*field, 0, numField),}}
 *fields) Add(field *field) {fs.Table[field.name] = fieldfs.List = append(fs.List, field)if field.omitEmpty {fs.hasOmitEmpty = true}}
 *fields) OmitEmpty(strct reflect.Value) []*field {if !fs.hasOmitEmpty {return fs.List}fields := make([]*field, 0, len(fs.List))for _, f := range fs.List {if !f.Omit(strct) {fields = append(fields, f)}}return fields}
Fields(typ reflect.Type, useJSONTag bool) *fields {numField := typ.NumField()fs := newFields(numField)var omitEmpty boolfor i := 0; i < numField; i++ {f := typ.Field(i)tag := f.Tag.Get("msgpack")if useJSONTag && tag == "" {tag = f.Tag.Get("json")}name, opt := parseTag(tag)if name == "-" {continue}if f.Name == "_msgpack" {if opt.Contains("asArray") {fs.AsArray = true}if opt.Contains("omitempty") {omitEmpty = true}}if f.PkgPath != "" && !f.Anonymous {continue}field := &field{name:      name,index:     f.Index,omitEmpty: omitEmpty || opt.Contains("omitempty"),encoder:   getEncoder(f.Type),decoder:   getDecoder(f.Type),}if field.name == "" {field.name = f.Name}if f.Anonymous && !opt.Contains("noinline") {inline := opt.Contains("inline")if inline {inlineFields(fs, f.Type, field, useJSONTag)} else {inline = autoinlineFields(fs, f.Type, field, useJSONTag)}if inline {fs.Table[field.name] = fieldcontinue}}fs.Add(field)}return fs}var encodeStructValuePtr uintptrvar decodeStructValuePtr uintptr
t() {encodeStructValuePtr = reflect.ValueOf(encodeStructValue).Pointer()decodeStructValuePtr = reflect.ValueOf(decodeStructValue).Pointer()}
ineFields(fs *fields, typ reflect.Type, f *field, useJSONTag bool) {inlinedFields := getFields(typ, useJSONTag).Listfor _, field := range inlinedFields {if _, ok := fs.Table[field.name]; ok {// Don't inline shadowed fields.continue}field.index = append(f.index, field.index...)fs.Add(field)}}
oinlineFields(fs *fields, typ reflect.Type, f *field, useJSONTag bool) bool {var encoder encoder
decoder decoder
yp.Kind() == reflect.Struct {encoder = f.encoderdecoder = f.decoder} else {for typ.Kind() == reflect.Ptr {typ = typ.Elem()encoder = getEncoder(typ)decoder = getDecoder(typ)}if typ.Kind() != reflect.Struct {return false}}if reflect.ValueOf(encoder).Pointer() != encodeStructValuePtr {return false}if reflect.ValueOf(decoder).Pointer() != decodeStructValuePtr {return false}inlinedFields := getFields(typ, useJSONTag).Listfor _, field := range inlinedFields {if _, ok := fs.Table[field.name]; ok {// Don't auto inline if there are shadowed fields.return false}}for _, field := range inlinedFields {field.index = append(f.index, field.index...)fs.Add(field)}return true}
mptyValue(v reflect.Value) bool {switch v.Kind() {case reflect.Array, reflect.Map, reflect.Slice, reflect.String:return v.Len() == 0case reflect.Bool:return !v.Bool()case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:return v.Int() == 0case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:return v.Uint() == 0case reflect.Float32, reflect.Float64:return v.Float() == 0case reflect.Interface, reflect.Ptr:return v.IsNil()}return false}
ldByIndex(v reflect.Value, index []int) reflect.Value {if len(index) == 1 {return v.Field(index[0])}for i, x := range index {if i > 0 {var ok boolv, ok = indirectNew(v)if !ok {return v}}v = v.Field(x)}return v}
irectNew(v reflect.Value) (reflect.Value, bool) {if v.Kind() == reflect.Ptr {if v.IsNil() {if !v.CanSet() {return v, false}elemType := v.Type().Elem()if elemType.Kind() != reflect.Struct {return v, false}v.Set(reflect.New(elemType))}v = v.Elem()}return v, true}