packageresource

import(
	"go/ast"
	"go/types"

	"github.com/bflad/tfproviderlint/helper/astutils"
)

const(
	RetryErrorFieldErr=`Err`
	RetryErrorFieldRetryable=`Retryable`

	TypeNameRetryError=`RetryError`
)

//retryErrorTypeisaninternalrepresentationoftheSDKhelper/resource.RetryErrortype
//
//Thisisusedtopreventimportingtherealtypesincetheprojectsupports
//multipleversionsoftheTerraformPluginSDK,whileallowingpassesto
//accessthedatainafamiliarmanner.
typeretryErrorTypestruct{
	Errerror
	Retryablebool
}

//RetryErrorInforepresentsallgatheredRetryErrordataforeasieraccess
typeRetryErrorInfostruct{
	AstCompositeLit*ast.CompositeLit
	Fieldsmap[string]*ast.KeyValueExpr
	RetryError*retryErrorType
	TypesInfo*types.Info
}

//NewRetryErrorInfoinstantiatesaRetryErrorInfo

NewRetryErrorInfo(cl*ast.CompositeLit,info*types.Info)*RetryErrorInfo{
	result:=&RetryErrorInfo{
		AstCompositeLit:cl,
		Fields:astutils.CompositeLitFields(cl),
		RetryError:&retryErrorType{},
		TypesInfo:info,
	}

	returnresult
}

eclaresFieldreturnstrueifthefieldnameispresentintheAST

(info*RetryErrorInfo)DeclaresField(fieldNamestring)bool{
	returninfo.Fields[fieldName]!=nil
}

//IsTypeRetryErrorreturnsifthetypeisRetryErrorfromthehelper/resourcepackage

IsTypeRetryError(ttypes.Type)bool{
	switcht:=t.(type){
	case*types.Named:
		returnIsNamedType(t,TypeNameRetryError)
	case*types.Pointer:
		returnIsTypeRetryError(t.Elem())
	default:
		returnfalse
	}
}
