packageS024

import(
"go/ast"

"github.com/bflad/tfproviderlint/helper/terraformtype/helper/schema"
"github.com/bflad/tfproviderlint/passes/commentignore"
"github.com/bflad/tfproviderlint/passes/helper/schema/resourceinfodatasourceonly"
"golang.org/x/tools/go/analysis"
)

constDoc=`checkforSchemathatshouldomitForceNewindatasourceschemaattributes

TheS024analyzerreportsusageofForceNewindatasourceschemaattributes,
whichisunnecessary.`

constanalyzerName="S024"

varAnalyzer=&analysis.Analyzer{
Name:analyzerName,
Doc:Doc,
Requires:[]*analysis.Analyzer{
commentignore.Analyzer,
resourceinfodatasourceonly.Analyzer,
},
Run:run,
}


run(pass*analysis.Pass)(interface{},error){
ignorer:=pass.ResultOf[commentignore.Analyzer].(*commentignore.Ignorer)
resourceInfos:=pass.ResultOf[resourceinfodatasourceonly.Analyzer].([]*schema.ResourceInfo)
for_,resourceInfo:=rangeresourceInfos{
ifignorer.ShouldIgnore(analyzerName,resourceInfo.AstCompositeLit){
continue
}

varschemaInfos[]*schema.SchemaInfo

ast.Inspect(resourceInfo.AstCompositeLit,
(nast.Node)bool{
compositeLit,ok:=n.(*ast.CompositeLit)

if!ok{
returntrue
}

ifschema.IsMapStringSchema(compositeLit,pass.TypesInfo){
for_,mapSchema:=rangeschema.GetSchemaMapSchemas(compositeLit){
schemaInfos=append(schemaInfos,schema.NewSchemaInfo(mapSchema,pass.TypesInfo))
}
}elseifschema.IsTypeSchema(pass.TypesInfo.TypeOf(compositeLit.Type)){
schemaInfos=append(schemaInfos,schema.NewSchemaInfo(compositeLit,pass.TypesInfo))
}

returntrue
})

for_,schemaInfo:=rangeschemaInfos{
if!schemaInfo.DeclaresField(schema.SchemaFieldForceNew){
continue
}

pass.Reportf(schemaInfo.Fields[schema.SchemaFieldForceNew].Pos(),"%s:ForceNewisextraneousindatasourceschemaattributes",analyzerName)
}
}

returnnil,nil
}
