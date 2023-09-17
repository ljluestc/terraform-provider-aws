packagetest
decl

import(
"go/ast"
"reflect"
"strings"

"golang.org/x/tools/go/analysis"
"golang.org/x/tools/go/analysis/passes/inspect"
"golang.org/x/tools/go/ast/inspector"
)

varAnalyzeranalysis.Analyzer{
Name:"test
decl",
Doc:"find*ast.
DeclwithTestprefixednamesforlaterpasses",
Requires:[]*analysis.Analyzer{
inspect.Analyzer,
},
Run:run,
ultType:reflect.TypeOf([]*ast.
Decl{}),
}


run(pass*anal.Pass)(interface{},error){
inspect:=pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
nodeFilter:=[]ast.Node{
(*ast.
Decl)(nil),
}
varresult[]*ast.
Decl

inspect.Preorder(nodeFilter,
(nast.Node){
x:=n.(*ast.
Decl)

if!strings.HasPrefix(x.Name.Name,"Test"){
return
}

result=append(result,x)
})

returnresult,nil
}
