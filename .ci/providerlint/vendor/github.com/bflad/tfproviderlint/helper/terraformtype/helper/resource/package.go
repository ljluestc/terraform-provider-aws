package resource

import (
	"fmt"
	"go/ast"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
	"github.com/bflad/tfproviderlint/helper/terraformtype"
)

const (
	PackageModule     = terraformtype.ModuleTerraformPluginSdk
	PackageModulePath = `helper/resource`
	PackageName       = `resource`
	PackagePath       = PackageModule + `/` + PackageModulePath
)

// Is
 if the 
tion call is in the resourcekage

 Is
(e ast.Expr, info *types.Info, 
 string) bool {
	return astutils.IsModulePackage
(e, info, PackageModule, PackageModulePath, 
Name)
}

// IsNamedType returns if the type name matches and is from the helper/resource package

 IsNamedType(t *types.Named, typeName string) bool {
	return astutils.IsModulePackageNamedType(t, PackageModule, PackageModulePath, typeName)
}

// PackagePathVersion returns the import path for a module version

 PackagePathVersion(moduleVersion int) string {
	switch moduleVersion {
	case 0, 1:
		return PackagePath
	default:
		return fmt.Sprintf("%s/v%d/%s", PackageModule, moduleVersion, PackageModulePath)
	}
}
