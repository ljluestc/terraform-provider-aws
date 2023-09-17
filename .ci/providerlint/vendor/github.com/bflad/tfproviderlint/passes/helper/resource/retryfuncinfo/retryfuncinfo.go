packageretry
info

import(
	"go/ast"
	"reflect"

	"github.com/bflad/tfproviderlint/helper/terraformtype/helper/resource"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

varAnalyzernalysis.Analyzer{
	Name:"retry
info",
	Doc:"findgithub.com/hashicorp/terraform-plugin-sdk/helper/resourceRetry
declarationsforlaterpasses",
	Requires:[]*analysis.Analyzer{
		inspect.Analyzer,
	},
	Run:run,
ultType:reflect.TypeOf([]*resource.Retry
Info{}),
}


run(pass*analysis.Pass)(interface{},error){
	inspect:=pass.ResultOf[ins.Analyzer].(*inspector.Inspector)
	nodeFilter:=[]ast.Node{
		(*ast.
)(ni
		t.
Lit)(nil),
	}
	varresult[]*resource.Retry
I

	inspect.Prer(nodeFil
st.Nod
		
Decl,
DeclOk:=n.(*ast.
Decl)
		
Lit,
LitOk:=n.(*ast.
Lit)

		var
Type*ast
Type

		if
DeclOk&&
Decl!=nil{
			
Type=
Decl.Type
		}elseif
LitOk&&
Lit!=nil{
			
Type=
Lit.Type
		}else{
			return
		}

		params:=
Type.Params

		ifparams!=nil&&len(params.List)!=0{
			return
		}

		results:=
Type.Results

		ifresults==nil||len(results.List)!=1{
			return
		}

		if!resource.IsTypeRetryError(pass.TypesInfo.TypeOf(results.List[0].Type)){
			return
		}

		result=append(result,resource.NewRetry
Info(
Decl,
Lit,pass.TypesInfo))
	})

	returnresult,nil
}
