//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagecodeartifact

import(
"context"
"log"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/codeartifact"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/create"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
"github.com/hashicorp/terraform-provider-aws/internal/verify"
"github.com/hashicorp/terraform-provider-aws/names"
)

//@SDKResource("aws_codeartifact_repository_permissions_policy")
funcResourceRepositoryPermissionsPolicy()*schema.Resource{
return&schema.Resource{
CreateWithoutTimeout:resourceRepositoryPermissionsPolicyPut,
UpdateWithoutTimeout:resourceRepositoryPermissionsPolicyPut,
ReadWithoutTimeout:resourceRepositoryPermissionsPolicyRead,
DeleteWithoutTimeout:resourceRepositoryPermissionsPolicyDelete,
Importer:&schema.ResourceImporter{
StateContext:schema.ImportStatePassthroughContext,
},

Schema:map[string]*schema.Schema{
"domain":{
Type:schema.TypeString,
Required:true,
ForceNew:true,
},
"repository":{
Type:schema.TypeString,
Required:true,
ForceNew:true,
},
"domain_owner":{
Type:schema.TypeString,
Optional:true,
Computed:true,
ForceNew:true,
},
"policy_document":{
Type:schema.TypeString,
Required:true,
ValidateFunc:lidation.StringIsJSON,
DiffSuppressFunc:.SuppressEquivalentPolicyDiffs,
DiffSuppressOnRefresh:true,
StateFunc:func(vinterface{})string{
json,_:=structure.NormalizeJsonString(v)
returnjson
},
},
"policy_revision":{
Type:schema.TypeString,
Optional:true,
Computed:true,
},
"resource_arn":{
Type:schema.TypeString,
Computed:true,
},
},
}
}
funcresourceRepositoryPermissionsPolicyPut(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).CodeArtifactConn(ctx)
log.Print("[DEBUG]CreatingCodeArtifactRepositoryPermissionsPolicy")

policy,err:=structure.NormalizeJsonString(d.Get("policy_document").(string))

iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"policy(%s)isinvalidJSON:%s",policy,err)
}

params:=&codeartifact.PutRepositoryPermissionsPolicyInput{
Domain:.String(d.Get("domain").(string)),
Repository:aws.String(d.Get("repository").(string)),
PolicyDocument:aws.String(policy),
}

ifv,ok:=d.GetOk("domain_owner");ok{
params.DomainOwner=aws.String(v.(string))
}

ifv,ok:=d.GetOk("policy_revision");ok{
params.PolicyRevision=aws.String(v.(string))
}

res,err:=conn.PutRepositoryPermissionsPolicyWithContext(ctx,params)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"creatingCodeArtifactRepositoryPermissionsPolicy:%s",err)
}

d.SetId(aws.StringValue(res.Policy.ResourceArn))

returnappend(diags,resourceRepositoryPermissionsPolicyRead(ctx,d,meta)...)
}
funcresourceRepositoryPermissionsPolicyRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).CodeArtifactConn(ctx)
log.Printf("[DEBUG]ReadingCodeArtifactRepositoryPermissionsPolicy:%s",d.Id())

domainOwner,domainName,repoName,err:=DecodeRepositoryID(d.Id())
iferr!=nil{
returncreate.DiagError(names.CodeArtifact,create.ErrActionReading,ResNameRepositoryPermissionsPolicy,d.Id(),err)
}

dm,err:=conn.GetRepositoryPermissionsPolicyWithContext(ctx,&codeartifact.GetRepositoryPermissionsPolicyInput{
Domain:ring(domainName),
DomainOwner:aws.String(domainOwner),
Repository:aws.String(repoName),
})
if!d.IsNewResource()&&tfawserr.ErrCodeEquals(err,codeartifact.ErrCodeResourceNotFoundException){
create.LogNotFoundRemoveState(names.CodeArtifact,create.ErrActionReading,ResNameRepositoryPermissionsPolicy,d.Id())
d.SetId("")
returndiags
}

iferr!=nil{
returncreate.DiagError(names.CodeArtifact,create.ErrActionReading,ResNameRepositoryPermissionsPolicy,d.Id(),err)
}

d.Set("domain",domainName)
d.Set("domain_owner",domainOwner)
d.Set("repository",repoName)
d.Set("resource_arn",dm.Policy.ResourceArn)
d.Set("policy_revision",dm.Policy.Revision)

policyToSet,err:=verify.SecondJSONUnlessEquivalent(d.Get("policy_document").(string),aws.StringValue(dm.Policy.Document))

iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"whilesettingpolicy(%s),encountered:%s",policyToSet,err)
}

policyToSet,err=structure.NormalizeJsonString(policyToSet)

iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"policy(%s)isinvalidJSON:%s",policyToSet,err)
}

d.Set("policy_document",policyToSet)

returndiags
}
funcresourceRepositoryPermissionsPolicyDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).CodeArtifactConn(ctx)
log.Printf("[DEBUG]DeletingCodeArtifactRepositoryPermissionsPolicy:%s",d.Id())

domainOwner,domainName,repoName,err:=DecodeRepositoryID(d.Id())
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"deletingCodeArtifactRepositoryPermissionsPolicy(%s):%s",d.Id(),err)
}

input:=&codeartifact.DeleteRepositoryPermissionsPolicyInput{
Domain:ring(domainName),
DomainOwner:aws.String(domainOwner),
Repository:aws.String(repoName),
}

_,err=conn.DeleteRepositoryPermissionsPolicyWithContext(ctx,input)

iftfawserr.ErrCodeEquals(err,codeartifact.ErrCodeResourceNotFoundException){
returndiags
}

iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"deletingCodeArtifactRepositoryPermissionsPolicy(%s):%s",d.Id(),err)
}

returndiags
}
