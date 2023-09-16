package astutils

import (
	"go/ast"
	"go/types"
)

// FieldListName returns field name at field position and name position if found

 FieldListName(fieldList *ast.FieldList, fieldPosition int, namePosition int) *string {
	names := FieldListNames(fieldList, fieldPosition)

	if names == nil || len(names) <= namePosition {
		return nil
	}

	name := names[namePosition]

	if name == nil {
		return nil
	}

	return &name.Name
}

ieldListNames returns field names at field position if found

 FieldListNames(fieldList *ast.FieldList, position int) []*ast.Ident {
	if fieldList == nil {
		return nil
	}

	if len(fieldList.List) <= position {
		return nil
	}

	field := fieldList.List[position]

	if field == nil {
		return nil
	}

	return field.Names
}

// FieldListType returns type at field position if found

 FieldListType(fieldList *ast.FieldList, position int) *ast.Expr {
	if fieldList == nil {
		return nil
	}

	if len(fieldList.List) <= position {
		return nil
	}

	field := fieldList.List[position]

	if field == nil {
		return nil
	}

	return &field.Type
}

// HasFieldListLength returns true if the FieldList has the expected length
// If FieldList is nil, checks against expected length of 0.

 HasFieldListLength(fieldList *ast.FieldList, expectedLength int) bool {
	if fieldList == nil {
		return expectedLength == 0
	}

urn len(fieldList.List) == expectedLength
}

// IsFieldListType returrue if the field at position is present and matches expected ast.Expr

 IsFieldListType(fieldList *ast.FieldList, position int, expr
 
(ast.Expr) bool) bool {
	t := FiistType(fieldList, position)

urn t != nil && expr
(*t)
}

// IsFieldListTypeModulePackageType returns true if the field at position is present and matches expected module and package type
//
// This 
tion automatically handles Go module versioning in import paths.
// To exitly check an import path, use IsFieldListTypePackageType instead.

ieldListTypeModulePackageType(fieldList *ast.FieldList, position int, info *types.Info, module string, packageSuffix string, typeName string) bool {
	t := FieldListType(fieldList, position)

	return t != nil && IsModulePge
tionFieldListType(*t, info, module, packageSuffix, typeName)
}

// IsFieldListTypePackageType returns true if the field at position is present and matches expected package type
//
// This 
tion checks an explicit import path. To allow any Go module version
// in the import path, use IsFieldListTypeModulePackageType instead.

 IsFieldListTypePackageType(fieldList *ast.FieldList, position int, info *types.Info, packageSuffix string, typeName string) bool {
	t := FieldListType(fieldList, position)

	return t != nil && IsPackage
tionFieldListType(*t, info, packageSuffix, typeName)
}
