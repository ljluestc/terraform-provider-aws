//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagecodeartifact

import(
"context"
"fmt"
"log"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/codeartifact"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

//@SDKDataSource("aws_codeartifact_repository_endpoint")
funcDataSourceRepositoryEndpoint()*schema.Resource{
return&schema.Resource{
ReadWithoutTimeout:dataSourceRepositoryEndpointRead,

Schema:map[string]*schema.Schema{
"domain":{
Type:schema.TypeString,
Required:true,
},
"repository":{
Type:schema.TypeString,
Required:true,
},
"format":{
Type:ema.TypeString,
Required:true,
ValidateFunc:validation.StringInSlice(codeartifact.PackageFormat_Values(),false),
},
"domain_owner":{
Type:ema.TypeString,
Optional:true,
Computed:true,
ValidateFunc:verify.ValidAccountID,
},
"repository_endpoint":{
Type:schema.TypeString,
Computed:true,
},
},
}
}
funcdataSourceRepositoryEndpointRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).CodeArtifactConn(ctx)
domainOwner:=meta.(*conns.AWSClient).AccountID
domain:=d.Get("domain").(string)
repo:=d.Get("repository").(string)
format:=d.Get("format").(string)
params:=&codeartifact.GetRepositoryEndpointInput{
Domain:aws.String(domain),
Repository:aws.String(repo),
Format:aws.String(format),
}

ifv,ok:=d.GetOk("domain_owner");ok{
params.DomainOwner=aws.String(v.(string))
domainOwner=v.(string)
}

log.Printf("[DEBUG]GettingCodeArtifactRepositoryEndpoint")
out,err:=conn.GetRepositoryEndpointWithContext(ctx,params)
iferr!=nil{
returnsdkdiag.AppendErrorf(diags,"gettingCodeArtifactRepositoryEndpoint:%s",err)
}
log.Printf("[DEBUG]CodeArtifactRepositoryEndpoint:%#v",out)

d.SetId(fmt.Sprintf("%s:%s:%s:%s",domainOwner,domain,repo,format))
d.Set("repository_endpoint",out.RepositoryEndpoint)
d.Set("domain_owner",domainOwner)

returndiags
}
