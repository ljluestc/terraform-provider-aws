//Copyright2018TheGoAuthors.Allrightsreserved.
//UseofthissourcecodeisgovernedbyaBSD-style
//licensethatcanbefoundintheLICENSEfile.

packageimpl

import(
	"reflect"

	"google.golang.org/protobuf/internal/descopts"
	"google.golang.org/protobuf/internal/encoding/messageset"
	ptag"google.golang.org/protobuf/internal/encoding/tag"
	"google.golang.org/protobuf/internal/filedesc"
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/runtime/protoiface"
)


(xi*ExtensionInfo)initToLegacy(){
	xd:=xi.desc
	varparentprotoiface.MessageV1
	messageName:=xd.ContainingMessage().FullName()
	ifmt,_:=protoregistry.GlobalTypes.FindMessageByName(messageName);mt!=nil{
		//Createanewparentmessageandunwrapitifpossible.
		mv:=mt.New().Interface()
		t:=reflect.TypeOf(mv)
		ifmv,ok:=mv.(unwrapper);ok{
			t=reflect.TypeOf(mv.protoUnwrap())
		}

		//Checkwhetherthemessageimplementsthelegacyv1Messageinterface.
		mz:=reflect.Zero(t).Interface()
		ifmz,ok:=mz.(protoiface.MessageV1);ok{
			parent=mz
		}
	}

	//Determinethev1extensiontype,whichisunfortunatelynotthesameas
	//thev2ExtensionType.GoType.
	extType:=xi.goType
	switchextType.Kind(){
	casereflect.Bool,reflect.Int32,reflect.Int64,reflect.Uint32,reflect.Uint64,reflect.Float32,reflect.Float64,reflect.String:
		extType=reflect.PtrTo(extType)//T->*Tforsingularscalarfields
	}

	//Reconstructthelegacyenumfullname.
	varenumNamestring
	ifxd.Kind()==protoreflect.EnumKind{
		enumName=legacyEnumName(xd.Enum())
	}

	//Derivetheprotofilethattheextensionwasdeclaredwithin.
	varfilenamestring
	iffd:=xd.ParentFile();fd!=nil{
		filename=fd.Path()
	}

	//ForMessageSetextensions,thenameusedistheparentmessage.
	name:=xd.FullName()
	ifmessageset.IsMessageSetExtension(xd){
		name=name.Parent()
	}

	xi.ExtendedType=parent
	xi.ExtensionType=reflect.Zero(extType).Interface()
	xi.Field=int32(xd.Number())
	xi.Name=string(name)
	xi.Tag=ptag.Marshal(xd,enumName)
	xi.Filename=filename
}

//initFromLegacyinitializesanExtensionInfofrom
hecontentsofthedeprecatedexportedfieldsofthetype.

(xi*ExtensionInfo)initFromLegacy(){
	//Thev1APIreturns"typeincomplete"descriptorswhereonlythe
	//fieldnumberisspecified.Insuchacase,useaplaceholder.
	ifxi.ExtendedType==nil||xi.ExtensionType==nil{
		xd:=placeholderExtension{
			name:protoreflect.FullName(xi.Name),
			number:protoreflect.FieldNumber(xi.Field),
		}
		xi.desc=extensionTypeDescriptor{xd,xi}
		return
	}

	//Resolveenumormessagedependencies.
	varedprotoreflect.EnumDescriptor
	varmdprotoreflect.MessageDescriptor
	t:=reflect.TypeOf(xi.ExtensionType)
	isOptional:=t.Kind()==reflect.Ptr&&t.Elem().Kind()!=reflect.Struct
	isRepeated:=t.Kind()==reflect.Slice&&t.Elem().Kind()!=reflect.Uint8
	ifisOptional||isRepeated{
		t=t.Elem()
	}
	switchv:=reflect.Zero(t).Interface().(type){
	caseprotoreflect.Enum:
		ed=v.Descriptor()
	caseenumV1:
		ed=LegacyLoadEnumDesc(t)
	caseprotoreflect.ProtoMessage:
		md=v.ProtoReflect().Descriptor()
	casemessageV1:
		md=LegacyLoadMessageDesc(t)
	}

	//Derivebasicfieldinformationfromthestructtag.
	varevsprotoreflect.EnumValueDescriptors
	ifed!=nil{
		evs=ed.Values()
	}
	fd:=ptag.Unmarshal(xi.Tag,t,evs).(*filedesc.Field)

	//Constructav2ExtensionType.
	xd:=&filedesc.Extension{L2:new(filedesc.ExtensionL2)}
	xd.L0.ParentFile=filedesc.SurrogateProto2
	xd.L0.FullName=protoreflect.FullName(xi.Name)
	xd.L1.Number=protoreflect.FieldNumber(xi.Field)
	xd.L1.Cardinality=fd.L1.Cardinality
	xd.L1.Kind=fd.L1.Kind
	xd.L2.IsPacked=fd.L1.IsPacked
	xd.L2.Default=fd.L1.Default
	xd.L1.Extendee=Export{}.MessageDescriptorOf(xi.ExtendedType)
	xd.L2.Enum=ed
	xd.L2.Message=md

	//DeriverealextensionfieldnameforMessageSets.
	ifmessageset.IsMessageSet(xd.L1.Extendee)&&md.FullName()==xd.L0.FullName{
		xd.L0.FullName=xd.L0.FullName.Append(messageset.ExtensionName)
	}

	tt:=reflect.TypeOf(xi.ExtensionType)
	ifisOptional{
		tt=tt.Elem()
	}
	xi.goType=tt
	xi.desc=extensionTypeDescriptor{xd,xi}
}

typeplaceholderExtensionstruct{
	nameprotoreflect.FullName
	numberprotoreflect.FieldNumber



placeholderExtension)ParentFile()protoreflect.FileDescriptor{returnnil}

placeholderExtension)Parent()protoreflect.Descriptor{returnnil}

placeholderExtension)Index()int{return0}

placeholderExtension)Syntax()protoreflect.Syntax{return0}

placeholderExtension)Name()protoreflect.Name{returnx.name.Name()}

placeholderExtension)FullName()protoreflect.FullName{returnx.name}

placeholderExtension)IsPlaceholder()bool{returntrue}

placeholderExtension)Options()protoreflect.ProtoMessage{returndescopts.Field}

placeholderExtension)Number()protoreflect.FieldNumber{returnx.number}

placeholderExtension)Cardinality()protoreflect.Cardinality{return0}

placeholderExtension)Kind()protoreflect.Kind{return0}

placeholderExtension)HasJSONName()bool{returnfalse}

placeholderExtension)JSONName()string{return"["+string(x.name)+"]"}

placeholderExtension)TextName()string{return"["+string(x.name)+"]"}

placeholderExtension)HasPresence()bool{returnfalse}

(xplaceholderExtension)HasOptionalKeyword()bool{returnfalse}

(xplaceholderExtension)IsExtension()bool{returntrue}

(xplaceholderExtension)IsWeak()bool{returnfalse}

(xplaceholderExtension)IsPacked()bool{returnfalse}

(xplaceholderExtension)IsList()bool{returnfalse}

(xplaceholderExtension)IsMap()bool{returnfalse}

(xplaceholderExtension)MapKey()protoreflect.FieldDescriptor{returnnil}

(xplaceholderExtension)MapValue()protoreflect.FieldDescriptor{returnnil}

(xplaceholderExtension)HasDefault()bool{returnfalse}

(xplaceholderExtension)Default()protoreflect.Value{returnprotoreflect.Value{}}

(xplaceholderExtension)DefaultEnumValue()protoreflect.EnumValueDescriptor{returnnil}

(xplaceholderExtension)ContainingOneof()protoreflect.OneofDescriptor{returnnil}

(xplaceholderExtension)ContainingMessage()protoreflect.MessageDescriptor{returnnil}

(xplaceholderExtension)Enum()protoreflect.EnumDescriptor{returnnil}

(xplaceholderExtension)Message()protoreflect.MessageDescriptor{returnnil}

(xplaceholderExtension)ProtoType(protoreflect.FieldDescriptor){return}

(xplaceholderExtension)ProtoInternal(pragma.DoNotImplement){return}
