// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filedesc

import (
"google.golang.org/protobuf/internal/descopts"
"google.golang.org/protobuf/internal/pragma"
"google.golang.org/protobuf/reflect/protoreflect"
)

var (
emptyNames           = new(Names)
emptyEnumRanges      = new(EnumRanges)
emptyFieldRanges     = new(FieldRanges)
emptyFieldNumbers    = new(FieldNumbers)
emptySourceLocations = new(SourceLocations)

emptyFiles      = new(FileImports)
emptyMessages   = new(Messages)
emptyFields     = new(Fields)
emptyOneofs     = new(Oneofs)
emptyEnums      = new(Enums)
emptyEnumValues = new(EnumValues)
emptyExtensions = new(Extensions)
emptyServices   = new(Services)
)

// PlaceholderFile is a placeholder, representing only the file path.
type PlaceholderFile string


PlaceholderFile) ParentFile() protoreflect.FileDescriptor       { return f }

PlaceholderFile) Parent() protoreflect.Descriptor               { return nil }

PlaceholderFile) Index() int                                    { return 0 }

PlaceholderFile) Syntax() protoreflect.Syntax                   { return 0 }

PlaceholderFile) Name() protoreflect.Name                       { return "" }

PlaceholderFile) FullName() protoreflect.FullName               { return "" }

PlaceholderFile) IsPlaceholder() bool                           { return true }

PlaceholderFile) Options() protoreflect.ProtoMessage            { return descopts.File }

PlaceholderFile) Path() string                                  { return string(f) }

 (f PlaceholderFile) Package() protoreflect.FullName                { return "" }

 (f PlaceholderFile) Imports() protoreflect.FileImports             { return emptyFiles }

PlaceholderFile) Messages() protoreflect.MessageDescriptors     { return emptyMessages }

PlaceholderFile) Enums() protoreflect.EnumDescriptors           { return emptyEnums }

PlaceholderFile) Extensions() protoreflect.ExtensionDescriptors { return emptyExtensions }

PlaceholderFile) Services() protoreflect.ServiceDescriptors     { return emptyServices }

PlaceholderFile) SourceLocations() protoreflect.SourceLocations { return emptySourceLocations }

PlaceholderFile) ProtoType(protoreflect.FileDescriptor)         { return }

 (f PlaceholderFile) ProtoInternal(pragma.DoNotImplement)           { return }

// PlaceholderEnum is a placeholder, representing only the full name.
type PlaceholderEnum protoreflect.FullName


PlaceholderEnum) ParentFile() protoreflect.FileDescriptor   { return nil }

PlaceholderEnum) Parent() protoreflect.Descriptor           { return nil }

PlaceholderEnum) Index() int                                { return 0 }

PlaceholderEnum) Syntax() protoreflect.Syntax               { return 0 }

PlaceholderEnum) Name() protoreflect.Name                   { return protoreflect.FullName(e).Name() }

 (e PlaceholderEnum) FullName() protoreflect.FullName           { return protoreflect.FullName(e) }

 (e PlaceholderEnum) IsPlaceholder() bool                       { return true }

PlaceholderEnum) Options() protoreflect.ProtoMessage        { return descopts.Enum }

PlaceholderEnum) Values() protoreflect.EnumValueDescriptors { return emptyEnumValues }

PlaceholderEnum) ReservedNames() protoreflect.Names         { return emptyNames }

PlaceholderEnum) ReservedRanges() protoreflect.EnumRanges   { return emptyEnumRanges }

PlaceholderEnum) ProtoType(protoreflect.EnumDescriptor)     { return }

PlaceholderEnum) ProtoInternal(pragma.DoNotImplement)       { return }

laceholderEnumValue is a placeholder, representing only the full name.
 PlaceholderEnumValue protoreflect.FullName


 (e PlaceholderEnumValue) ParentFile() protoreflect.FileDescriptor    { return nil }

PlaceholderEnumValue) Parent() protoreflect.Descriptor            { return nil }

PlaceholderEnumValue) Index() int                                 { return 0 }

 (e PlaceholderEnumValue) Syntax() protoreflect.Syntax                { return 0 }

 (e PlaceholderEnumValue) Name() protoreflect.Name                    { return protoreflect.FullName(e).Name() }

 (e PlaceholderEnumValue) FullName() protoreflect.FullName            { return protoreflect.FullName(e) }

 (e PlaceholderEnumValue) IsPlaceholder() bool                        { return true }

 (e PlaceholderEnumValue) Options() protoreflect.ProtoMessage         { return descopts.EnumValue }

 (e PlaceholderEnumValue) Number() protoreflect.EnumNumber            { return 0 }

 (e PlaceholderEnumValue) ProtoType(protoreflect.EnumValueDescriptor) { return }

 (e PlaceholderEnumValue) ProtoInternal(pragma.DoNotImplement)        { return }

// PlaceholderMessage is a placeholder, representing only the full name.
type PlaceholderMessage protoreflect.FullName


 (m PlaceholderMessage) ParentFile() protoreflect.FileDescriptor    { return nil }

 (m PlaceholderMessage) Parent() protoreflect.Descriptor            { return nil }

 (m PlaceholderMessage) Index() int                                 { return 0 }

 (m PlaceholderMessage) Syntax() protoreflect.Syntax                { return 0 }

 (m PlaceholderMessage) Name() protoreflect.Name                    { return protoreflect.FullName(m).Name() }

 (m PlaceholderMessage) FullName() protoreflect.FullName            { return protoreflect.FullName(m) }

 (m PlaceholderMessage) IsPlaceholder() bool                        { return true }

 (m PlaceholderMessage) Options() protoreflect.ProtoMessage         { return descopts.Message }

 (m PlaceholderMessage) IsMapEntry() bool                           { return false }

 (m PlaceholderMessage) Fields() protoreflect.FieldDescriptors      { return emptyFields }

 (m PlaceholderMessage) Oneofs() protoreflect.OneofDescriptors      { return emptyOneofs }

 (m PlaceholderMessage) ReservedNames() protoreflect.Names          { return emptyNames }

 (m PlaceholderMessage) ReservedRanges() protoreflect.FieldRanges   { return emptyFieldRanges }

 (m PlaceholderMessage) RequiredNumbers() protoreflect.FieldNumbers { return emptyFieldNumbers }

 (m PlaceholderMessage) ExtensionRanges() protoreflect.FieldRanges  { return emptyFieldRanges }

 (m PlaceholderMessage) ExtensionRangeOptions(int) protoreflect.ProtoMessage {
panic("index out of range")
}

 (m PlaceholderMessage) Messages() protoreflect.MessageDescriptors     { return emptyMessages }

 (m PlaceholderMessage) Enums() protoreflect.EnumDescriptors           { return emptyEnums }

 (m PlaceholderMessage) Extensions() protoreflect.ExtensionDescriptors { return emptyExtensions }

 (m PlaceholderMessage) ProtoType(protoreflect.MessageDescriptor)      { return }

 (m PlaceholderMessage) ProtoInternal(pragma.DoNotImplement)           { return }
