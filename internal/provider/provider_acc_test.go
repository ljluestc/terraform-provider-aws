//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageprovider_test

import(
"context"
"fmt"
"reflect"
"strings"
"testing"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/endpoints"
"github.com/hashicorp/terraform-plugin-go/tfprotov5"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
sdkacctest"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
"github.com/hashicorp/terraform-plugin-testing/helper/resource"
"github.com/hashicorp/terraform-plugin-testing/terraform"
"github.com/hashicorp/terraform-provider-aws/internal/acctest"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/provider"
"github.com/hashicorp/terraform-provider-aws/names"
)

funcTestAccProvider_DefaultTags_emptyBlock(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_defaultTagsEmptyConfigurationBlock(),
Check:resource.ComposeTestCheckFunc(
testAccCheckProviderDefaultTags_Tags(ctx,t,&provider,map[string]string{}),
),
},
},
})
}

funcTestAccProvider_DefaultTagsTags_none(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{//nosemgrep:ci.test-config-funcs-correct-form
Config:acctest.ConfigDefaultTags_Tags0(),
Check:resource.ComposeTestCheckFunc(
testAccCheckProviderDefaultTags_Tags(ctx,t,&provider,map[string]string{}),
),
},
},
})
}

funcTestAccProvider_DefaultTagsTags_one(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{//nosemgrep:ci.test-config-funcs-correct-form
Config:acctest.ConfigDefaultTags_Tags1("test","value"),
Check:resource.ComposeTestCheckFunc(
testAccCheckProviderDefaultTags_Tags(ctx,t,&provider,map[string]string{"test":"value"}),
),
},
},
})
}

funcTestAccProvider_DefaultTagsTags_multiple(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{//nosemgrep:ci.test-config-funcs-correct-form
Config:acctest.ConfigDefaultTags_Tags2("test1","value1","test2","value2"),
Check:resource.ComposeTestCheckFunc(
testAccCheckProviderDefaultTags_Tags(ctx,t,&provider,map[string]string{
"test1":"value1",
"test2":"value2",
}),
),
},
},
})
}

funcTestAccProvider_DefaultAndIgnoreTags_emptyBlocks(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_defaultAndIgnoreTagsEmptyConfigurationBlock(),
Check:resource.ComposeTestCheckFunc(
testAccCheckProviderDefaultTags_Tags(ctx,t,&provider,map[string]string{}),
testAccCheckIgnoreTagsKeys(ctx,t,&provider,[]string{}),
testAccCheckIgnoreTagsKeyPrefixes(ctx,t,&provider,[]string{}),
),
},
},
})
}

funcTestAccProvider_endpoints(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider
varendpointsstrings.Builder

//Initializeeachendpointconfigurationwithmatchingnameandvalue
for_,serviceKey:=rangenames.ProviderPackages(){
endpoints.WriteString(fmt.Sprintf("%s=\"http://%s\"\n",serviceKey,serviceKey))
}

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_endpoints(endpoints.String()),
Check:resource.ComposeTestCheckFunc(
testAccCheckEndpoints(ctx,&provider),
),
},
},
})
}

funcTestAccProvider_fipsEndpoint(t*testing.T){
ctx:=acctest.Context(t)
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
resourceName:="aws_s3_bucket.test"

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_fipsEndpoint(fmt.Sprintf("https://s3-fips.%s.%s",acctest.Region(),acctest.PartitionDNSSuffix()),rName),
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(resourceName,"bucket",rName),
),
},
},
})
}

typeunusualEndpointstruct{
fieldNamestring
thingstring
urlstring
}

funcTestAccProvider_unusualEndpoints(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider
unusual1:=unusualEndpoint{"es","elasticsearch","http://notarealendpoint"}
unusual2:=unusualEndpoint{"databasemigration","dms","http://alsonotarealendpoint"}
unusual3:=unusualEndpoint{"lexmodelbuildingservice","lexmodels","http://kingofspain"}

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_unusualEndpoints(unusual1,unusual2,unusual3),
Check:resource.ComposeTestCheckFunc(
testAccCheckUnusualEndpoints(ctx,&provider,unusual1),
testAccCheckUnusualEndpoints(ctx,&provider,unusual2),
testAccCheckUnusualEndpoints(ctx,&provider,unusual3),
),
},
},
})
}

funcTestAccProvider_IgnoreTags_emptyBlock(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_ignoreTagsEmptyConfigurationBlock(),
Check:resource.ComposeTestCheckFunc(
testAccCheckIgnoreTagsKeys(ctx,t,&provider,[]string{}),
testAccCheckIgnoreTagsKeyPrefixes(ctx,t,&provider,[]string{}),
),
},
},
})
}

funcTestAccProvider_IgnoreTagsKeyPrefixes_none(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_ignoreTagsKeyPrefixes0(),
Check:resource.ComposeTestCheckFunc(
testAccCheckIgnoreTagsKeyPrefixes(ctx,t,&provider,[]string{}),
),
},
},
})
}

funcTestAccProvider_IgnoreTagsKeyPrefixes_one(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_ignoreTagsKeyPrefixes3("test"),
Check:resource.ComposeTestCheckFunc(
testAccCheckIgnoreTagsKeyPrefixes(ctx,t,&provider,[]string{"test"}),
),
},
},
})
}

funcTestAccProvider_IgnoreTagsKeyPrefixes_multiple(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_ignoreTagsKeyPrefixes2("test1","test2"),
Check:resource.ComposeTestCheckFunc(
testAccCheckIgnoreTagsKeyPrefixes(ctx,t,&provider,[]string{"test1","test2"}),
),
},
},
})
}

funcTestAccProvider_IgnoreTagsKeys_none(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_ignoreTagsKeys0(),
Check:resource.ComposeTestCheckFunc(
testAccCheckIgnoreTagsKeys(ctx,t,&provider,[]string{}),
),
},
},
})
}

funcTestAccProvider_IgnoreTagsKeys_one(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_ignoreTagsKeys1("test"),
Check:resource.ComposeTestCheckFunc(
testAccCheckIgnoreTagsKeys(ctx,t,&provider,[]string{"test"}),
),
},
},
})
}

funcTestAccProvider_IgnoreTagsKeys_multiple(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_ignoreTagsKeys2("test1","test2"),
Check:resource.ComposeTestCheckFunc(
testAccCheckIgnoreTagsKeys(ctx,t,&provider,[]string{"test1","test2"}),
),
},
},
})
}

funcTestAccProvider_Region_c2s(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_region(endpoints.UsIsoEast1RegionID),
Check:resource.ComposeTestCheckFunc(
testAccCheckDNSSuffix(ctx,t,&provider,"c2s.ic.gov"),
testAccCheckPartition(ctx,t,&provider,endpoints.AwsIsoPartitionID),
testAccCheckReverseDNSPrefix(ctx,t,&provider,"gov.ic.c2s"),
),
PlanOnly:true,
},
},
})
}

funcTestAccProvider_Region_china(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_region(endpoints.CnNorthwest1RegionID),
Check:resource.ComposeTestCheckFunc(
testAccCheckDNSSuffix(ctx,t,&provider,"amazonaws.com.cn"),
testAccCheckPartition(ctx,t,&provider,endpoints.AwsCnPartitionID),
testAccCheckReverseDNSPrefix(ctx,t,&provider,"cn.com.amazonaws"),
),
PlanOnly:true,
},
},
})
}

funcTestAccProvider_Region_commercial(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_region(endpoints.UsWest2RegionID),
Check:resource.ComposeTestCheckFunc(
testAccCheckDNSSuffix(ctx,t,&provider,"amazonaws.com"),
testAccCheckPartition(ctx,t,&provider,endpoints.AwsPartitionID),
testAccCheckReverseDNSPrefix(ctx,t,&provider,"com.amazonaws"),
),
PlanOnly:true,
},
},
})
}

funcTestAccProvider_Region_govCloud(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_region(endpoints.UsGovWest1RegionID),
Check:resource.ComposeTestCheckFunc(
testAccCheckDNSSuffix(ctx,t,&provider,"amazonaws.com"),
testAccCheckPartition(ctx,t,&provider,endpoints.AwsUsGovPartitionID),
testAccCheckReverseDNSPrefix(ctx,t,&provider,"com.amazonaws"),
),
PlanOnly:true,
},
},
})
}

funcTestAccProvider_Region_sc2s(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_region(endpoints.UsIsobEast1RegionID),
Check:resource.ComposeTestCheckFunc(
testAccCheckDNSSuffix(ctx,t,&provider,"sc2s.sgov.gov"),
testAccCheckPartition(ctx,t,&provider,endpoints.AwsIsoBPartitionID),
testAccCheckReverseDNSPrefix(ctx,t,&provider,"gov.sgov.sc2s"),
),
PlanOnly:true,
},
},
})
}

funcTestAccProvider_Region_stsRegion(t*testing.T){
ctx:=acctest.Context(t)
varprovider*schema.Provider

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:testAccProtoV5ProviderFactoriesInternal(ctx,t,&provider),
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_stsRegion(endpoints.UsEast1RegionID,endpoints.UsWest2RegionID),
Check:resource.ComposeTestCheckFunc(
testAccCheckRegion(ctx,t,&provider,endpoints.UsEast1RegionID),
testAccCheckSTSRegion(ctx,t,&provider,endpoints.UsWest2RegionID),
),
PlanOnly:true,
},
},
})
}

funcTestAccProvider_AssumeRole_empty(t*testing.T){
ctx:=acctest.Context(t)
resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccProviderConfig_assumeRoleEmpty,
Check:resource.ComposeTestCheckFunc(
acctest.CheckCallerIdentityAccountID("data.aws_caller_identity.current"),
),
},
},
})
}

functestAccProtoV5ProviderFactoriesInternal(ctxcontext.Context,t*testing.T,v**schema.Provider)map[string]func()(tfprotov5.ProviderServer,error){
providerServerFactory,p,err:=provider.ProtoV5ProviderServerFactory(ctx)

iferr!=nil{
t.Fatal(err)
}

providerServer:=providerServerFactory()
*v=p

returnmap[string]func()(tfprotov5.ProviderServer,error){
acctest.ProviderName:func()(tfprotov5.ProviderServer,error){//nolint:unparam
returnproviderServer,nil
},
}
}

functestAccCheckPartition(ctxcontext.Context,t*testing.T,p**schema.Provider,expectedPartitionstring)resource.TestCheckFunc{//nolint:unparam
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}

providerPartition:=(*p).Meta().(*conns.AWSClient).Partition

ifproviderPartition!=expectedPartition{
returnfmt.Errorf("expectedDNSSuffix(%s),got:%s",expectedPartition,providerPartition)
}

returnnil
}
}

functestAccCheckDNSSuffix(ctxcontext.Context,t*testing.T,p**schema.Provider,expectedDnsSuffixstring)resource.TestCheckFunc{//nolint:unparam
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}

providerDnsSuffix:=(*p).Meta().(*conns.AWSClient).DNSSuffix

ifproviderDnsSuffix!=expectedDnsSuffix{
returnfmt.Errorf("expectedDNSSuffix(%s),got:%s",expectedDnsSuffix,providerDnsSuffix)
}

returnnil
}
}

functestAccCheckRegion(ctxcontext.Context,t*testing.T,p**schema.Provider,expectedRegionstring)resource.TestCheckFunc{//nolint:unparam
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}

ifgot:=(*p).Meta().(*conns.AWSClient).Region;got!=expectedRegion{
returnfmt.Errorf("expectedRegion(%s),got:%s",expectedRegion,got)
}

returnnil
}
}

functestAccCheckSTSRegion(ctxcontext.Context,t*testing.T,p**schema.Provider,expectedRegionstring)resource.TestCheckFunc{//nolint:unparam
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}

stsRegion:=aws.StringValue((*p).Meta().(*conns.AWSClient).STSConn(ctx).Config.Region)

ifstsRegion!=expectedRegion{
returnfmt.Errorf("expectedSTSRegion(%s),got:%s",expectedRegion,stsRegion)
}

returnnil
}
}

functestAccCheckReverseDNSPrefix(ctxcontext.Context,t*testing.T,p**schema.Provider,expectedReverseDnsPrefixstring)resource.TestCheckFunc{//nolint:unparam
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}
providerReverseDnsPrefix:=(*p).Meta().(*conns.AWSClient).ReverseDNSPrefix

ifproviderReverseDnsPrefix!=expectedReverseDnsPrefix{
returnfmt.Errorf("expectedDNSSuffix(%s),got:%s",expectedReverseDnsPrefix,providerReverseDnsPrefix)
}

returnnil
}
}

functestAccCheckIgnoreTagsKeyPrefixes(ctxcontext.Context,t*testing.T,p**schema.Provider,expectedKeyPrefixes[]string)resource.TestCheckFunc{//nolint:unparam
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}

providerClient:=(*p).Meta().(*conns.AWSClient)
ignoreTagsConfig:=providerClient.IgnoreTagsConfig

ifignoreTagsConfig==nil||ignoreTagsConfig.KeyPrefixes==nil{
iflen(expectedKeyPrefixes)!=0{
returnfmt.Errorf("expectedkey_prefixes(%d)length,got:0",len(expectedKeyPrefixes))
}

returnnil
}

actualKeyPrefixes:=ignoreTagsConfig.KeyPrefixes.Keys()

iflen(actualKeyPrefixes)!=len(expectedKeyPrefixes){
returnfmt.Errorf("expectedkey_prefixes(%d)length,got:%d",len(expectedKeyPrefixes),len(actualKeyPrefixes))
}

for_,expectedElement:=rangeexpectedKeyPrefixes{
varfoundbool

for_,actualElement:=rangeactualKeyPrefixes{
ifactualElement==expectedElement{
found=true
break
}
}

if!found{
returnfmt.Errorf("expectedkey_prefixeselement,butwasmissing:%s",expectedElement)
}
}

for_,actualElement:=rangeactualKeyPrefixes{
varfoundbool

for_,expectedElement:=rangeexpectedKeyPrefixes{
ifactualElement==expectedElement{
found=true
break
}
}

if!found{
returnfmt.Errorf("unexpectedkey_prefixeselement:%s",actualElement)
}
}

returnnil
}
}

functestAccCheckIgnoreTagsKeys(ctxcontext.Context,t*testing.T,p**schema.Provider,expectedKeys[]string)resource.TestCheckFunc{//nolint:unparam
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}

providerClient:=(*p).Meta().(*conns.AWSClient)
ignoreTagsConfig:=providerClient.IgnoreTagsConfig

ifignoreTagsConfig==nil||ignoreTagsConfig.Keys==nil{
iflen(expectedKeys)!=0{
returnfmt.Errorf("expectedkeys(%d)length,got:0",len(expectedKeys))
}

returnnil
}

actualKeys:=ignoreTagsConfig.Keys.Keys()

iflen(actualKeys)!=len(expectedKeys){
returnfmt.Errorf("expectedkeys(%d)length,got:%d",len(expectedKeys),len(actualKeys))
}

for_,expectedElement:=rangeexpectedKeys{
varfoundbool

for_,actualElement:=rangeactualKeys{
ifactualElement==expectedElement{
found=true
break
}
}

if!found{
returnfmt.Errorf("expectedkeyselement,butwasmissing:%s",expectedElement)
}
}

for_,actualElement:=rangeactualKeys{
varfoundbool

for_,expectedElement:=rangeexpectedKeys{
ifactualElement==expectedElement{
found=true
break
}
}

if!found{
returnfmt.Errorf("unexpectedkeyselement:%s",actualElement)
}
}

returnnil
}
}

functestAccCheckProviderDefaultTags_Tags(ctxcontext.Context,t*testing.T,p**schema.Provider,expectedTagsmap[string]string)resource.TestCheckFunc{//nolint:unparam
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}

providerClient:=(*p).Meta().(*conns.AWSClient)
defaultTagsConfig:=providerClient.DefaultTagsConfig

ifdefaultTagsConfig==nil||len(defaultTagsConfig.Tags)==0{
iflen(expectedTags)!=0{
returnfmt.Errorf("expectedkeys(%d)length,got:0",len(expectedTags))
}

returnnil
}

actualTags:=defaultTagsConfig.Tags

iflen(actualTags)!=len(expectedTags){
returnfmt.Errorf("expectedtags(%d)length,got:%d",len(expectedTags),len(actualTags))
}

for_,expectedElement:=rangeexpectedTags{
varfoundbool

for_,actualElement:=rangeactualTags{
ifaws.StringValue(actualElement.Value)==expectedElement{
found=true
break
}
}

if!found{
returnfmt.Errorf("expectedtagselement,butwasmissing:%s",expectedElement)
}
}

for_,actualElement:=rangeactualTags{
varfoundbool

for_,expectedElement:=rangeexpectedTags{
ifaws.StringValue(actualElement.Value)==expectedElement{
found=true
break
}
}

if!found{
returnfmt.Errorf("unexpectedtagselement:%s",actualElement)
}
}

returnnil
}
}

functestAccCheckEndpoints(_context.Context,p**schema.Provider)resource.TestCheckFunc{
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}

providerClient:=(*p).Meta().(*conns.AWSClient)

for_,serviceKey:=rangenames.Aliases(){
methodName:=serviceConn(serviceKey)
method:=reflect.ValueOf(providerClient).MethodByName(methodName)
if!method.IsValid(){
continue
}
ifmethod.Kind()!=reflect.Func{
returnfmt.Errorf("value%qisnotafunction",methodName)
}
if!funcHasConnFuncSignature(method){
returnfmt.Errorf("function%qdoesnotmatchexpectedsignature",methodName)
}

result:=method.Call([]reflect.Value{
reflect.ValueOf(context.Background()),
})
ifl:=len(result);l!=1{
returnfmt.Errorf("expected1result,got%d",l)
}
providerClientField:=result[0]

if!providerClientField.IsValid(){
returnfmt.Errorf("unabletomatchconns.AWSClientstructfieldnameforendpointname:%s",serviceKey)
}

if!reflect.Indirect(providerClientField).FieldByName("Config").IsValid(){
continue//currentlyunknownhowtodothischeckforv2clients
}

actualEndpoint:=reflect.Indirect(reflect.Indirect(providerClientField).FieldByName("Config").FieldByName("Endpoint")).String()
expectedEndpoint:=fmt.Sprintf("http://%s",serviceKey)

ifactualEndpoint!=expectedEndpoint{
returnfmt.Errorf("expectedendpoint(%s)value(%s),got:%s",serviceKey,expectedEndpoint,actualEndpoint)
}
}

returnnil
}
}

functestAccCheckUnusualEndpoints(_context.Context,p**schema.Provider,unusualunusualEndpoint)resource.TestCheckFunc{
returnfunc(s*terraform.State)error{
ifp==nil||*p==nil||(*p).Meta()==nil||(*p).Meta().(*conns.AWSClient)==nil{
returnfmt.Errorf("providernotinitialized")
}

providerClient:=(*p).Meta().(*conns.AWSClient)

methodName:=serviceConn(unusual.thing)
method:=reflect.ValueOf(providerClient).MethodByName(methodName)
ifmethod.Kind()!=reflect.Func{
returnfmt.Errorf("value%qisnotafunction",methodName)
}
if!funcHasConnFuncSignature(method){
returnfmt.Errorf("function%qdoesnotmatchexpectedsignature",methodName)
}

result:=method.Call([]reflect.Value{
reflect.ValueOf(context.Background()),
})
ifl:=len(result);l!=1{
returnfmt.Errorf("expected1result,got%d",l)
}
providerClientField:=result[0]

if!providerClientField.IsValid(){
returnfmt.Errorf("unabletomatchconns.AWSClientstructfieldnameforendpointname:%s",unusual.thing)
}

actualEndpoint:=reflect.Indirect(reflect.Indirect(providerClientField).FieldByName("Config").FieldByName("Endpoint")).String()
expectedEndpoint:=unusual.url

ifactualEndpoint!=expectedEndpoint{
returnfmt.Errorf("expectedendpoint(%s)value(%s),got:%s",unusual.thing,expectedEndpoint,actualEndpoint)
}

returnnil
}
}

funcfuncHasConnFuncSignature(methodreflect.Value)bool{
typ:=method.Type()
iftyp.NumIn()!=1{
returnfalse
}

fn:=func(ctxcontext.Context){}
ftyp:=reflect.TypeOf(fn)

returntyp.In(0)==ftyp.In(0)
}

funcserviceConn(keystring)string{
serviceUpper:=""
varerrerror
ifserviceUpper,err=names.ProviderNameUpper(key);err!=nil{
return""
}

returnfmt.Sprintf("%sConn",serviceUpper)
}

consttestAccProviderConfig_assumeRoleEmpty=`
provider"aws"{
assume_role{
}
}

data"aws_caller_identity""current"{}
`//lintignore:AT004

consttestAccProviderConfig_base=`
data"aws_region""provider_test"{}

#Requiredtoinitializetheprovider.
data"aws_service""provider_test"{
region=data.aws_region.provider_test.name
service_id="s3"
}
`

functestAccProviderConfig_endpoints(endpointsstring)string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,fmt.Sprintf(`
provider"aws"{
skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true

endpoints{
%[1]s
}
}
`,endpoints))
}

functestAccProviderConfig_fipsEndpoint(endpoint,rNamestring)string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,fmt.Sprintf(`
provider"aws"{
endpoints{
s3=%[1]q
}
}

resource"aws_s3_bucket""test"{
bucket=%[2]q
force_destroy=true
}
`,endpoint,rName))
}

functestAccProviderConfig_unusualEndpoints(unusual1,unusual2,unusual3unusualEndpoint)string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,fmt.Sprintf(`
provider"aws"{
skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true

endpoints{
%[1]s=%[2]q
%[3]s=%[4]q
%[5]s=%[6]q
}
}
`,unusual1.fieldName,unusual1.url,unusual2.fieldName,unusual2.url,unusual3.fieldName,unusual3.url))
}

functestAccProviderConfig_ignoreTagsKeys0()string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,`
provider"aws"{
skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`)
}

functestAccProviderConfig_ignoreTagsKeys1(tag1string)string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,fmt.Sprintf(`
provider"aws"{
ignore_tags{
keys=[%[1]q]
}

skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`,tag1))
}

functestAccProviderConfig_ignoreTagsKeys2(tag1,tag2string)string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,fmt.Sprintf(`
provider"aws"{
ignore_tags{
keys=[%[1]q,%[2]q]
}

skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`,tag1,tag2))
}

functestAccProviderConfig_ignoreTagsKeyPrefixes0()string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,`
provider"aws"{
skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`)
}

functestAccProviderConfig_ignoreTagsKeyPrefixes3(tagPrefix1string)string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,fmt.Sprintf(`
provider"aws"{
ignore_tags{
key_prefixes=[%[1]q]
}

skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`,tagPrefix1))
}

functestAccProviderConfig_ignoreTagsKeyPrefixes2(tagPrefix1,tagPrefix2string)string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,fmt.Sprintf(`
provider"aws"{
ignore_tags{
key_prefixes=[%[1]q,%[2]q]
}

skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`,tagPrefix1,tagPrefix2))
}

functestAccProviderConfig_defaultTagsEmptyConfigurationBlock()string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,`
provider"aws"{
default_tags{}

skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`)
}

functestAccProviderConfig_defaultAndIgnoreTagsEmptyConfigurationBlock()string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,`
provider"aws"{
default_tags{}
ignore_tags{}

skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`)
}

functestAccProviderConfig_ignoreTagsEmptyConfigurationBlock()string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,`
provider"aws"{
ignore_tags{}

skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`)
}

functestAccProviderConfig_region(regionstring)string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,fmt.Sprintf(`
provider"aws"{
region=%[1]q
skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`,region))
}

functestAccProviderConfig_stsRegion(region,stsRegionstring)string{
//lintignore:AT004
returnacctest.ConfigCompose(testAccProviderConfig_base,fmt.Sprintf(`
provider"aws"{
region=%[1]q
sts_region=%[2]q
skip_credentials_validation=true
skip_metadata_api_check=true
skip_requesting_account_id=true
}
`,region,stsRegion))
}
