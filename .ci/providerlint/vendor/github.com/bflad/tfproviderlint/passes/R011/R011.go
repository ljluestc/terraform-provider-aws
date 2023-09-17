packageR011

import(
"go/ast"

"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
"github.com/bflad/tfproviderlint/passes/commentignore"
"github.com/bflad/tfproviderlint/passes/helper/schema/resourceinfo"
"golang.org/x/tools/go/analysis"
)

constDoc=`checkforResourcewithMigrateStateconfigured

TheR011analyzerreportscasesofresourceswhichconfigureMigrateState.
AfterTerraform0.12,resourcesmustconfigurenewstatemigrationsvia
StateUpgraders.ExistingimplementationsofMigrateStatepriortoTerraform
0.12canbeignoredcurrently.`

constanalyzerName="R011"

varAnalyzer=&analysis.Analyzer{
Name:analyzerName,
Doc:Doc,
Requires:[]*analysis.Analyzer{
commentignore.Analyzer,
resourceinfo.Analyzer,
},
Run:run,
}


run(pass*analysis.Pass)(interface{},error){
ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
resourceInfos:=pass.ResultOf[resourceinfo.Analyzer].([]*schema.ResourceInfo)
for_,resourceInfo:=rangeresourceInfos{
ifignorer.ShouldIgnore(analyzerName,resourceInfo.AstCompositeLit){
continue
}

ifresourceInfo.Resource.MigrateState==nil{
continue
}

switcht:=resourceInfo.AstCompositeLit.Type.(type){
default:
pass.Reportf(resourceInfo.AstCompositeLit.Lbrace,"%s:resourceshouldconfigureStateUpgradersinsteadofMigrateState(implementationspriortoTerraform0.12canbeignored)",analyzerName)
case*ast.SelectorExpr:
pass.Reportf(t.Sel.Pos(),"%s:resourceshouldconfigureStateUpgradersinsteadofMigrateState(implementationspriortoTerraform0.12canbeignored)",analyzerName)
}
}

returnnil,nil
}
