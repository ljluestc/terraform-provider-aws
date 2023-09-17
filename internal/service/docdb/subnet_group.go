//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagedocdb

import(
"context"
"fmt"
"log"
"time"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/docdb"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
"github.com/hashicorp/terraform-provider-aws/internal/flex"
tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
"github.com/hashicorp/terraform-provider-aws/internal/verify"
"github.com/hashicorp/terraform-provider-aws/names"
)

//@SDKResource("aws_docdb_subnet_group",name="SubnetGroup")
//@Tags(identifierAttribute="arn")
funcResourceSubnetGroup()*schema.Resource{
return&schema.Resource{
CreateWithoutTimeout:resourceSubnetGroupCreate,
ReadWithoutTimeout:resourceSubnetGroupRead,
UpdateWithoutTimeout:resourceSubnetGroupUpdate,
DeleteWithoutTimeout:resourceSubnetGroupDelete,
Importer:&schema.ResourceImporter{
StateContext:schema.ImportStatePassthroughContext,
},

Schema:map[string]*schema.Schema{
"arn":{
Type:schema.TypeString,
Computed:true,
},

"name":{
Type:hema.TypeString,
Optional:
Computed:
ForceNew:
ConflictsWith:[]string{"name_prefix"},
ValidateFunc:validSubnetGroupName,
},
"name_prefix":{
Type:hema.TypeString,
Optional:
Computed:
ForceNew:
ConflictsWith:[]string{"name"},
ValidateFunc:validSubnetGroupNamePrefix,
},

"description":{
Type:schema.TypeString,
Optional:true,
Default:"ManagedbyTerraform",
},

"subnet_ids":{
Type:schema.TypeSet,
Required:true,
MinItems:1,
Elem:&schema.Schema{Type:schema.TypeString},
Set:.HashString,
},

names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll:tftags.TagsSchemaComputed(),
},

CustomizeDiff:verify.SetTagsDiff,
}
}
funcresourceSubnetGroupCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

subnetIds:=flex.ExpandStringSet(d.Get("subnet_ids").(*schema.Set))

vargroupNamestring
ifv,ok:=d.GetOk("name");ok{
groupName=v.(string)
}elseifv,ok:=d.GetOk("name_prefix");ok{
groupName=id.PrefixedUniqueId(v.(string))
}else{
groupName=id.UniqueId()
}

input:=docdb.CreateDBSubnetGroupInput{
DBSubnetGroupName:String(groupName),
DBSubnetGroupDescription:aws.String(d.Get("description").(string)),
SubnetIds:subnetIds,
Tags:agsIn(ctx),
}

_,err:=conn.CreateDBSubnetGroupWithContext(ctx,&input)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"creatingDocumentDBSubnetGroup:%s",err)
}

d.SetId(groupName)

returnappend(diags,resourceSubnetGroupRead(ctx,d,meta)...)
}
funcresourceSubnetGroupRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

describeOpts:=docdb.DescribeDBSubnetGroupsInput{
DBSubnetGroupName:aws.String(d.Id()),
}

varsubnetGroups[]*docdb.DBSubnetGroup
iferr:=conn.DescribeDBSubnetGroupsPagesWithContext(ctx,&describeOpts,func(resp*docdb.DescribeDBSubnetGroupsOutput,lastPagebool)bool{
subnetGroups=append(subnetGroups,resp.DBSubnetGroups...)
return!lastPage
});err!=nil{
if!d.IsNewResource()&&tfawserr.ErrCodeEquals(err,docdb.ErrCodeDBSubnetGroupNotFoundFault){
log.Printf("[WARN]DocumentDBSubnetGroup(%s)notfound,removingfromstate",d.Id())
d.SetId("")
returndiags
}
returnsdkdiag.AppendErrorf(diags,"readingDocumentDBSubnetGroup(%s)parameters:%s",d.Id(),err)
}

if!d.IsNewResource()&&(len(subnetGroups)!=1||aws.StringValue(subnetGroups[0].DBSubnetGroupName)!=d.Id()){
log.Printf("[WARN]DocumentDBSubnetGroup(%s)notfound,removingfromstate",d.Id())
d.SetId("")
returndiags
}

subnetGroup:=subnetGroups[0]
d.Set("name",subnetGroup.DBSubnetGroupName)
d.Set("description",subnetGroup.DBSubnetGroupDescription)
d.Set("arn",subnetGroup.DBSubnetGroupArn)

subnets:=make([]string,0,len(subnetGroup.Subnets))
for_,s:=rangesubnetGroup.Subnets{
subnets=append(subnets,aws.StringValue(s.SubnetIdentifier))
}
iferr:=d.Set("subnet_ids",subnets);err!=nil{
returnsdkdiag.AppendErrorf(diags,"settingsubnet_ids:%s",err)
}

returndiags
}
funcresourceSubnetGroupUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

ifd.HasChanges("subnet_ids","description"){
_,n:=d.GetChange("subnet_ids")
ifn==nil{
n=new(schema.Set)
}
sIds:=flex.ExpandStringSet(n.(*schema.Set))

_,err:=conn.ModifyDBSubnetGroupWithContext(ctx,&docdb.ModifyDBSubnetGroupInput{
DBSubnetGroupName:String(d.Id()),
DBSubnetGroupDescription:aws.String(d.Get("description").(string)),
SubnetIds:sIds,
})

iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"modifyingDocumentDBSubnetGroup(%s)parameters:%s",d.Id(),err)
}
}

returnappend(diags,resourceSubnetGroupRead(ctx,d,meta)...)
}
funcresourceSubnetGroupDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).DocDBConn(ctx)

delOpts:=docdb.DeleteDBSubnetGroupInput{
DBSubnetGroupName:aws.String(d.Id()),
}

log.Printf("[DEBUG]DeletingDocumentDBSubnetGroup:%s",d.Id())

_,err:=conn.DeleteDBSubnetGroupWithContext(ctx,&delOpts)
iferr!=nil{
iftfawserr.ErrCodeEquals(err,docdb.ErrCodeDBSubnetGroupNotFoundFault){
returndiags
}
returnsdkdiag.AppendErrorf(diags,"deletingDocumentDBSubnetGroup(%s):%s",d.Id(),err)
}

iferr:=WaitForSubnetGroupDeletion(ctx,conn,d.Id());err!=nil{
returnsdkdiag.AppendErrorf(diags,"deletingDocumentDBSubnetGroup(%s):%s",d.Id(),err)
}
returndiags
}
funcWaitForSubnetGroupDeletion(ctxcontext.Context,conn*docdb.DocDB,namestring)error{
params:=&docdb.DescribeDBSubnetGroupsInput{
DBSubnetGroupName:aws.String(name),
}

err:=retry.RetryContext(ctx,10*time.Minute,func()*retry.RetryError{
_,err:=conn.DescribeDBSubnetGroupsWithContext(ctx,params)

iftfawserr.ErrCodeEquals(err,docdb.ErrCodeDBSubnetGroupNotFoundFault){
returnnil
}

iferr!=nil{
returnretry.NonRetryableError(err)
}

returnretry.RetryableError(fmt.Errorf("DocumentDBSubnetGroup(%s)stillexists",name))
})
iftfresource.TimedOut(err){
_,err=conn.DescribeDBSubnetGroupsWithContext(ctx,params)
iftfawserr.ErrCodeEquals(err,docdb.ErrCodeDBSubnetGroupNotFoundFault){
returnnil
}
}
iferr!=nil{
returnfmt.Errorf("deletingDocumentDBsubnetgroup:%s",err)
}
returnnil
}
