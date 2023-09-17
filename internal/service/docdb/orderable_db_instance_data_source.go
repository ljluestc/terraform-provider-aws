//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagedocdb

import(
"context"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/docdb"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
)

//@SDKDataSource("aws_docdb_orderable_db_instance")
funcDataSourceOrderableDBInstance()*schema.Resource{
return&schema.Resource{
ReadWithoutTimeout:dataSourceOrderableDBInstanceRead,
Schema:map[string]*schema.Schema{
"availability_zones":{
Type:schema.TypeList,
Computed:true,
Elem:&schema.Schema{Type:schema.TypeString},
},

"engine":{
Type:schema.TypeString,
Optional:true,
Default:"docdb",
},

"engine_version":{
Type:schema.TypeString,
Optional:true,
Computed:true,
},

"instance_class":{
Type:hema.TypeString,
Optional:
Computed:
ConflictsWith:[]string{"preferred_instance_classes"},
},

"license_model":{
Type:schema.TypeString,
Optional:true,
Default:"na",
},

"preferred_instance_classes":{
Type:hema.TypeList,
Optional:
Elem:chema.Schema{Type:schema.TypeString},
ConflictsWith:[]string{"instance_class"},
},

"vpc":{
Type:schema.TypeBool,
Optional:true,
Computed:true,
},
},
}
}
funcdataSourceOrderableDBInstanceRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

input:=&docdb.DescribeOrderableDBInstanceOptionsInput{}

ifv,ok:=d.GetOk("instance_class");ok{
input.DBInstanceClass=aws.String(v.(string))
}

ifv,ok:=d.GetOk("engine");ok{
input.Engine=aws.String(v.(string))
}

ifv,ok:=d.GetOk("engine_version");ok{
input.EngineVersion=aws.String(v.(string))
}

ifv,ok:=d.GetOk("license_model");ok{
input.LicenseModel=aws.String(v.(string))
}

ifv,ok:=d.GetOk("vpc");ok{
input.Vpc=aws.Bool(v.(bool))
}

varinstanceClassResults[]*docdb.OrderableDBInstanceOption

err:=conn.DescribeOrderableDBInstanceOptionsPagesWithContext(ctx,input,func(resp*docdb.DescribeOrderableDBInstanceOptionsOutput,lastPagebool)bool{
for_,instanceOption:=rangeresp.OrderableDBInstanceOptions{
ifinstanceOption==nil{
continue
}

instanceClassResults=append(instanceClassResults,instanceOption)
}
return!lastPage
})

iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"readingDocumentDBorderableDBinstanceoptions:%s",err)
}

iflen(instanceClassResults)==0{
returnsdkdiag.AppendErrorf(diags,"noDocumentDBOrderableDBInstanceoptionsfoundmatchingcriteria;trydifferentsearch")
}

//preferredclasses
varfound*docdb.OrderableDBInstanceOption
ifl:=d.Get("preferred_instance_classes").([]interface{});len(l)>0{
for_,elem:=rangel{
preferredInstanceClass,ok:=elem.(string)

if!ok{
continue
}

for_,instanceClassResult:=rangeinstanceClassResults{
ifpreferredInstanceClass==aws.StringValue(instanceClassResult.DBInstanceClass){
found=instanceClassResult
break
}
}

iffound!=nil{
break
}
}
}

iffound==nil&&len(instanceClassResults)>1{
returnsdkdiag.AppendErrorf(diags,"multipleDocumentDBDBInstanceClasses(%v)matchthecriteria;tryadifferentsearch",instanceClassResults)
}

iffound==nil&&len(instanceClassResults)==1{
found=instanceClassResults[0]
}

iffound==nil{
returnsdkdiag.AppendErrorf(diags,"noDocumentDBDBInstanceClassesmatchthecriteria;tryadifferentsearch")
}

d.SetId(aws.StringValue(found.DBInstanceClass))

d.Set("instance_class",found.DBInstanceClass)

varavailabilityZones[]string
for_,az:=rangefound.AvailabilityZones{
availabilityZones=append(availabilityZones,aws.StringValue(az.Name))
}
d.Set("availability_zones",availabilityZones)

d.Set("engine",found.Engine)
d.Set("engine_version",found.EngineVersion)
d.Set("license_model",found.LicenseModel)
d.Set("vpc",found.Vpc)

returndiags
}
