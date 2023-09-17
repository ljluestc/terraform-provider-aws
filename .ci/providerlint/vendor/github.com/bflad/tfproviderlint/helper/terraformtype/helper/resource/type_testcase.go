packageresource

import(
	"go/ast"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const(
	TestCaseFieldCheckDestroy=`CheckDestroy`
	TestCaseFieldErrorCheck=`ErrorCheck`
	TestCaseFieldIDRefreshName=`IDRefreshName`
	TestCaseFieldIDRefreshIgnore=`IDRefreshIgnore`
	TestCaseFieldIsUnitTest=`IsUnitTest`
	TestCaseFieldPreCheck=`PreCheck`
	TestCaseFieldPreventPostDestroyRefresh=`PreventPostDestroyRefresh`
	TestCaseFieldProviders=`Providers`
	TestCaseFieldProviderFactories=`ProviderFactories`
	TestCaseFieldSteps=`Steps`

	TypeNameTestCase=`TestCase`
)

//testCaseTypeisaninternalrepresentationoftheSDKhelper/resource.TestCasetype
//
//Thisisusedtopreventimportingtherealtypesincetheprojectsupports
//multipleversionsoftheTerraformPluginSDK,whileallowingpassesto
//accessthedatainafamiliarmanner.
typetestCaseTypestruct{}

//TestCaseInforepresentsallgatheredTestCasedataforeasieraccess
typeTestCaseInfostruct{
	AstCompositeLit*ast.CompositeLit
	Fieldsmap[string]*ast.KeyValueExpr
	TestCase*testCaseType
	TypesInfo*types.Info
}

//NewTestCaseInfoinstantiatesaTestCaseInfo

NewTestCaseInfo(cl*ast.CompositeLit,info*types.Info)*TestCaseInfo{
	result:=&TestCaseInfo{
		AstCompositeLit:cl,
		Fields:astutils.CompositeLitFields(cl),
		TestCase:&testCaseType{},
		TypesInfo:info,
	}

	returnresult
}

eclaresFieldreturnstrueifthefieldnameispresentintheAST

(info*TestCaseInfo)DeclaresField(fieldNamestring)bool{
	returninfo.Fields[fieldName]!=nil
}

//IsTypeTestCasereturnsifthetypeisTestCasefromthehelper/schemapackage

IsTypeTestCase(ttypes.Type)bool{
	switcht:=t.(type){
	case*types.Named:
		returnIsNamedType(t,TypeNameTestCase)
	case*types.Pointer:
		returnIsTypeTestCase(t.Elem())
	default:
		returnfalse
	}
}
