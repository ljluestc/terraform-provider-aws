packageresource

import(
	"go/ast"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const(
	TestStepFieldCheck=`Check`
	TestStepFieldConfig=`Config`
	TestStepFieldDestroy=`Destroy`
	TestStepFieldExpectError=`ExpectError`
	TestStepFieldExpectNonEmptyPlan=`ExpectNonEmptyPlan`
	TestStepFieldImportState=`ImportState`
	TestStepFieldImportStateId=`ImportStateId`
	TestStepFieldImportStateId
=`ImportStateId
`
	TestStepFieldImportStateIdPrefix=`ImportStateIdPrefix`
	TestStepFieldImportStateCheck=`ImportStateCheck`
	TestStepFieldImportStateVerify=`ImportStateVerify`
	TestStepFieldImportStateVerifyIgnore=`ImportStateVerifyIgnore`
	TestStepFieldPlanOnly=`PlanOnly`
	TestStepFieldPreConfig=`PreConfig`
	TestStepFieldPreventDiskCleanup=`PreventDiskCleanup`
	TestStepFieldPrevostDestroyRefresh=`PrevostDestroyRefresh`
	TestStepFieldResourceName=`ResourceName`
	TestStepFieldSkip
=`Skip
`
	TestStepFieldTaint=`Taint`

	TypeNameTestStep=`TestStep`
)

//testStepTypeisaninternalrepresentationoftheSDKhelper/resource.TestSteptype
//
//Thisisusedtopreventimportingtherealtypesincetheprojectsupports
//multipleversionsoftheTerraformPluginSDK,whileallowingpassesto
//accessthedatainafamiliarmanner.
typetestStepTypestruct{}

//TestStepInforepresentsallgatheredTestStepdataforeasieraccess
typeTestStepInfostruct{
	AstCompositeLit*ast.CompositeLit
	Fieldsmap[string]*ast.KeyValueExpr
	TestStep*testStepType
esInfo*types.Info
}

//NewTestStepInfoinstantiatesaTestStepInfo

NewTestStepInfo(cl*ast.CompositeLit,info*types.Info)*TestStepInfo{
	result:=&TestStepInfo{
		AstCompositeLit:cl,
		Fields:astutils.CompositeLitFields(cl),
		TestStep:&testStepType{},
		TypesInfo:info,
	}

	returnresult
}

//DeclaresFieldreturnstrueifthefieldnameispresentintheAST

(info*TestStepInfo)DeclaresField(fieldNamestring)bool{
	returninfo.Fields[fieldName]!=nil
}

//IsTypeTestStepreturnsifthetypeisTestStepfromthehelper/schemapackage

IsTypeTestStep(ttypes.Type)bool{
	switcht:=t.(type){
	case*types.Named:
		returnIsNamedType(t,TypeNameTestStep)
	case*types.Pointer:
		returnIsTypeTestStep(t.Elem())
	default:
		returnfalse
	}
}
