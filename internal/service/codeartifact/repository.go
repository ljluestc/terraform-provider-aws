//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagecodeartifact

import(
"context"
"fmt"
"log"
"strings"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/arn"
"github.com/aws/aws-sdk-go/service/codeartifact"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/create"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
"github.com/hashicorp/terraform-provider-aws/internal/verify"
"github.com/hashicorp/terraform-provider-aws/names"
)

//@SDKResource("aws_codeartifact_repository",name="Repository")
//@Tags(identifierAttribute="arn")
funcResourceRepository()*schema.Resource{
return&schema.Resource{
CreateWithoutTimeout:resourceRepositoryCreate,
ReadWithoutTimeout:resourceRepositoryRead,
UpdateWithoutTimeout:resourceRepositoryUpdate,
DeleteWithoutTimeout:resourceRepositoryDelete,
Importer:&schema.ResourceImporter{
StateContext:schema.ImportStatePassthroughContext,
},

Schema:map[string]*schema.Schema{
"arn":{
Type:schema.TypeString,
Computed:true,
},
"repository":{
Type:schema.TypeString,
Required:true,
ForceNew:true,
},
"domain":{
Type:schema.TypeString,
Required:true,
ForceNew:true,
},
"domain_owner":{
Type:ema.TypeString,
Optional:true,
ForceNew:true,
Computed:true,
ValidateFunc:verify.ValidAccountID,
},
"description":{
Type:schema.TypeString,
Optional:true,
},
"upstream":{
Type:schema.TypeList,
MinItems:1,
Optional:true,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
"repository_name":{
Type:schema.TypeString,
Required:true,
},
},
},
},
"external_connections":{
Type:schema.TypeList,
Optional:true,
MaxItems:1,
Elem:&schema.Resource{
Schema:map[string]*schema.Schema{
"external_connection_name":{
Type:schema.TypeString,
Required:true,
},
"package_format":{
Type:schema.TypeString,
Computed:true,
},
"status":{
Type:schema.TypeString,
Computed:true,
},
},
},
},
"administrator_account":{
Type:schema.TypeString,
Computed:true,
},
names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll:tftags.TagsSchemaComputed(),
},

CustomizeDiff:verify.SetTagsDiff,
}
}
funcresourceRepositoryCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).CodeArtifactConn(ctx)

input:=&codeartifact.CreateRepositoryInput{
Repository:aws.String(d.Get("repository").(string)),
Domain:aws.String(d.Get("domain").(string)),
Tags:gsIn(ctx),
}

ifv,ok:=d.GetOk("description");ok{
input.Description=aws.String(v.(string))
}

ifv,ok:=d.GetOk("domain_owner");ok{
input.DomainOwner=aws.String(v.(string))
}

ifv,ok:=d.GetOk("upstream");ok{
input.Upstreams=expandUpstreams(v.([]interface{}))
}

res,err:=conn.CreateRepositoryWithContext(ctx,input)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"creatingCodeArtifactRepository:%s",err)
}

repo:=res.Repository
d.SetId(aws.StringValue(repo.Arn))

ifv,ok:=d.GetOk("external_connections");ok{
externalConnection:=v.([]interface{})[0].(map[string]interface{})
input:=&codeartifact.AssociateExternalConnectionInput{
Domain:repo.DomainName,
Repository:o.Name,
DomainOwner:.DomainOwner,
ExternalConnection:aws.String(externalConnection["external_connection_name"].(string)),
}

_,err:=conn.AssociateExternalConnectionWithContext(ctx,input)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"associatingexternalconnectiontoCodeArtifactrepository:%s",err)
}
}

returnappend(diags,resourceRepositoryRead(ctx,d,meta)...)
}
funcresourceRepositoryRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).CodeArtifactConn(ctx)

owner,domain,repo,err:=DecodeRepositoryID(d.Id())
iferr!=nil{
returncreate.DiagError(names.CodeArtifact,create.ErrActionReading,ResNameRepository,d.Id(),err)
}
sm,err:=conn.DescribeRepositoryWithContext(ctx,&codeartifact.DescribeRepositoryInput{
Repository:aws.String(repo),
Domain:ring(domain),
DomainOwner:aws.String(owner),
})
if!d.IsNewResource()&&tfawserr.ErrCodeEquals(err,codeartifact.ErrCodeResourceNotFoundException){
create.LogNotFoundRemoveState(names.CodeArtifact,create.ErrActionReading,ResNameRepository,d.Id())
d.SetId("")
returndiags
}

iferr!=nil{
returncreate.DiagError(names.CodeArtifact,create.ErrActionReading,ResNameRepository,d.Id(),err)
}

arn:=aws.StringValue(sm.Repository.Arn)
d.Set("repository",sm.Repository.Name)
d.Set("arn",arn)
d.Set("domain_owner",sm.Repository.DomainOwner)
d.Set("domain",sm.Repository.DomainName)
d.Set("administrator_account",sm.Repository.AdministratorAccount)
d.Set("description",sm.Repository.Description)

ifsm.Repository.Upstreams!=nil{
iferr:=d.Set("upstream",flattenUpstreams(sm.Repository.Upstreams));err!=nil{
returnsdkdiag.AppendErrorf(diags,"[WARN]Errorsettingupstream:%s",err)
}
}

ifsm.Repository.ExternalConnections!=nil{
iferr:=d.Set("external_connections",flattenExternalConnections(sm.Repository.ExternalConnections));err!=nil{
returnsdkdiag.AppendErrorf(diags,"[WARN]Errorsettingexternal_connections:%s",err)
}
}

returndiags
}
funcresourceRepositoryUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).CodeArtifactConn(ctx)
log.Print("[DEBUG]UpdatingCodeArtifactRepository")

needsUpdate:=false
params:=&codeartifact.UpdateRepositoryInput{
Repository:aws.String(d.Get("repository").(string)),
Domain:ring(d.Get("domain").(string)),
DomainOwner:aws.String(d.Get("domain_owner").(string)),
}

ifd.HasChange("description"){
ifv,ok:=d.GetOk("description");ok{
params.Description=aws.String(v.(string))
needsUpdate=true
}
}

ifd.HasChange("upstream"){
ifv,ok:=d.GetOk("upstream");ok{
params.Upstreams=expandUpstreams(v.([]interface{}))
needsUpdate=true
}
}

ifneedsUpdate{
_,err:=conn.UpdateRepositoryWithContext(ctx,params)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"updatingCodeArtifactRepository:%s",err)
}
}

ifd.HasChange("external_connections"){
ifv,ok:=d.GetOk("external_connections");ok{
externalConnection:=v.([]interface{})[0].(map[string]interface{})
input:=&codeartifact.AssociateExternalConnectionInput{
Repository:.String(d.Get("repository").(string)),
Domain:aws.String(d.Get("domain").(string)),
DomainOwner:String(d.Get("domain_owner").(string)),
ExternalConnection:aws.String(externalConnection["external_connection_name"].(string)),
}

_,err:=conn.AssociateExternalConnectionWithContext(ctx,input)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"associatingexternalconnectiontoCodeArtifactrepository:%s",err)
}
}else{
oldConn,_:=d.GetChange("external_connections")
externalConnection:=oldConn.([]interface{})[0].(map[string]interface{})
input:=&codeartifact.DisassociateExternalConnectionInput{
Repository:.String(d.Get("repository").(string)),
Domain:aws.String(d.Get("domain").(string)),
DomainOwner:String(d.Get("domain_owner").(string)),
ExternalConnection:aws.String(externalConnection["external_connection_name"].(string)),
}

_,err:=conn.DisassociateExternalConnectionWithContext(ctx,input)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"disassociatingexternalconnectiontoCodeArtifactrepository:%s",err)
}
}
}

returnappend(diags,resourceRepositoryRead(ctx,d,meta)...)
}
funcresourceRepositoryDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).CodeArtifactConn(ctx)
log.Printf("[DEBUG]DeletingCodeArtifactRepository:%s",d.Id())

owner,domain,repo,err:=DecodeRepositoryID(d.Id())
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"deletingCodeArtifactRepository(%s):%s",d.Id(),err)
}
input:=&codeartifact.DeleteRepositoryInput{
Repository:aws.String(repo),
Domain:ring(domain),
DomainOwner:aws.String(owner),
}

_,err=conn.DeleteRepositoryWithContext(ctx,input)

iftfawserr.ErrCodeEquals(err,codeartifact.ErrCodeResourceNotFoundException){
returndiags
}

iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"deletingCodeArtifactRepository(%s):%s",d.Id(),err)
}

returndiags
}
funcexpandUpstreams(l[]interface{})[]*codeartifact.UpstreamRepository{
upstreams:=[]*codeartifact.UpstreamRepository{}

for_,mRaw:=rangel{
m:=mRaw.(map[string]interface{})
upstream:=&codeartifact.UpstreamRepository{
RepositoryName:aws.String(m["repository_name"].(string)),
}

upstreams=append(upstreams,upstream)
}

returnupstreams
}
funcflattenUpstreams(upstreams[]*codeartifact.UpstreamRepositoryInfo)[]interface{}{
iflen(upstreams)==0{
returnnil
}

varls[]interface{}

for_,upstream:=rangeupstreams{
m:=map[string]interface{}{
"repository_name":aws.StringValue(upstream.RepositoryName),
}

ls=append(ls,m)
}

returnls
}
funcflattenExternalConnections(connections[]*codeartifact.RepositoryExternalConnectionInfo)[]interface{}{
iflen(connections)==0{
returnnil
}

varls[]interface{}

for_,connection:=rangeconnections{
m:=map[string]interface{}{
"external_connection_name":aws.StringValue(connection.ExternalConnectionName),
"package_format":ws.StringValue(connection.PackageFormat),
"status":ringValue(connection.Status),
}

ls=append(ls,m)
}

returnls
}
funcDecodeRepositoryID(idstring)(string,string,string,error){
repoArn,err:=arn.Parse(id)
iferr!=nil{
return"","","",err
}

idParts:=strings.Split(strings.TrimPrefix(repoArn.Resource,"repository/"),"/")
iflen(idParts)!=2{
return"","","",fmt.Errorf("expectedresourcepartofarninformatDomainName/RepositoryName,received:%s",repoArn.Resource)
}
returnrepoArn.AccountID,idParts[0],idParts[1],nil
}
