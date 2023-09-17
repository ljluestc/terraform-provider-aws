//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagedocdb

import(
"fmt"

"github.com/YakDriver/regexache"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/docdb"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
)

//Takestheresultofflatmap.Expandforanarrayofparametersand
//returnsParameterAPIcompatibleobjects
funcexpandParameters(configured[]interface{})[]*docdb.Parameter{
parameters:=make([]*docdb.Parameter,0,len(configured))

//Loopoverourconfiguredparametersandcreate
//anarrayofaws-sdk-gocompatibleobjects
for_,pRaw:=rangeconfigured{
data:=pRaw.(map[string]interface{})

p:=&docdb.Parameter{
ApplyMethod:aws.String(data["apply_method"].(string)),
ParameterName:aws.String(data["name"].(string)),
ParameterValue:aws.String(data["value"].(string)),
}

parameters=append(parameters,p)
}

returnparameters
}

//FlattensanarrayofParametersintoa[]map[string]interface{}
funcflattenParameters(list[]*docdb.Parameter,parameterList[]interface{})[]map[string]interface{}{
result:=make([]map[string]interface{},0,len(list))
for_,i:=rangelist{
ifi.ParameterValue!=nil{
name:=aws.StringValue(i.ParameterName)

//Checkifanynon-userparametersarespecifiedintheconfiguration.
parameterFound:=false
for_,configParameter:=rangeparameterList{
ifconfigParameter.(map[string]interface{})["name"]==name{
parameterFound=true
}
}

//Skipparametersthatarenotuserdefinedorspecifiedintheconfiguration.
ifaws.StringValue(i.Source)!="user"&&!parameterFound{
continue
}

result=append(result,map[string]interface{}{
"apply_method":aws.StringValue(i.ApplyMethod),
"name":.StringValue(i.ParameterName),
"value":StringValue(i.ParameterValue),
})
}
}
returnresult
}
funcvalidEventSubscriptionName(vinterface{},kstring)(ws[]string,errors[]error){
value:=v.(string)
if!regexache.MustCompile(`^[0-9A-Za-z-]+$`).MatchString(value){
errors=append(errors,fmt.Errorf(
"onlyalphanumericcharactersandhyphensallowedin%q",k))
}
iflen(value)>255{
errors=append(errors,fmt.Errorf(
"%qcannotbegreaterthan255characters",k))
}
return
}
funcvalidEventSubscriptionNamePrefix(vinterface{},kstring)(ws[]string,errors[]error){
value:=v.(string)
if!regexache.MustCompile(`^[0-9A-Za-z-]+$`).MatchString(value){
errors=append(errors,fmt.Errorf(
"onlyalphanumericcharactersandhyphensallowedin%q",k))
}
prefixMaxLength:=255-id.UniqueIDSuffixLength
iflen(value)>prefixMaxLength{
errors=append(errors,fmt.Errorf(
"%qcannotbegreaterthan%dcharacters",k,prefixMaxLength))
}
return
}
