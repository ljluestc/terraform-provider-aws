//PackageR004definesanAnalyzerthatchecksfor
//ResourceData.Set()callsusingincompatiblevaluetypes
packageR004

import(
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
	"github.com/bflad/tfproviderlint/passes/commentignore"
	"github.com/bflad/tfproviderlint/passes/helper/schema/resourcedatasetcallexpr"
)

constDoc=`checkforResourceData.Set()callsusingincompatiblevaluetypes

TheR004analyzerreportsincorrecttypesforaSet()callvalue.
TheSet()
tiononlysupportsasubsetofbasictypes,slicesandmapsofthat
subsetofbasictypes,andtheschema.Settype.`

constanalyzerName="R004"

varAnalyzer=&analysis.Analyzer{
	Name:analyzerName,
	Doc:Doc,
	Requires:[]*analysis.Analyzer{
		resourcedatasetcallexpr.Analyzer,
		commentignore.Analyzer,
	},
	Run:run,
}


run(pass*analysis.Pass)(interface{},error){
	ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
	sets:=pass.ResultOf[resourcedatasetcallexpr.Analyzer].([]*ast.CallExpr)
	for_,set:=rangesets{
		ifignorer.ShouldIgnore(analyzerName,set){
			continue
		}

		iflen(set.Args)<2{
			continue
		}

		pos:=set.Args[1].Pos()
		t:=pass.TypesInfo.TypeOf(set.Args[1]).Underlying()

		if!isAllowedType(t){
			pass.Reportf(pos,"%s:ResourceData.Set()incompatiblevaluetype:%s",analyzerName,t.String())
		}
	}

	returnnil,nil



isAllowedType(ttypes.Type)bool{
	switcht:=t.(type){
	default:
		returnfalse
	case*types.Basic:
		returnisAllowedBasicType(t)
	case*types.Interface:
		returntrue
	case*types.Map:
		switchk:=t.Key().Underlying().(type){
		default:
			returnfalse
		case*types.Basic:
			ifk.Kind()!=types.String{
				returnfalse
			}

			returnisAllowedType(t.Elem().Underlying())
		}
	case*types.Named:
		returnschema.IsNamedType(t,schema.TypeNameSet)
	case*types.Pointer:
		returnisAllowedType(t.Elem())
	case*types.Slice:
		returnisAllowedType(t.Elem().Underlying())
	}
}

varallowedBasicKindTypes=[]types.BasicKind{
	types.Bool,
	types.Float32,
	types.Float64,
	types.Int,
	types.Int8,
	types.Int16,
	types.Int32,
	types.Int64,
	types.String,
es.UntypedNil,
}


isAllowedBasicType(b*types.Basic)bool{
	for_,allowedBasicKindType:=rangeallowedBasicKindTypes{
		ifb.Kind()==allowedBasicKindType{
			returntrue
		}
	}

	returnfalse
}
