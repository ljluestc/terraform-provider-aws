//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagecodeartifact_test

import(
"context"
"fmt"
"testing"

"github.com/YakDriver/regexache"
"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/codeartifact"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
sdkacctest"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
"github.com/hashicorp/terraform-plugin-testing/helper/resource"
"github.com/hashicorp/terraform-plugin-testing/terraform"
"github.com/hashicorp/terraform-provider-aws/internal/acctest"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
tfcodeartifact"github.com/hashicorp/terraform-provider-aws/internal/service/codeartifact"
)
functestAccRepositoryPermissionsPolicy_basic(t*testing.T){
ctx:=acctest.Context(t)
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
resourceName:="aws_codeartifact_repository_permissions_policy.test"

resource.Test(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t);acctest.PreCheckPartitionHasService(t,codeartifact.EndpointsID)},
ErrorCheck:acctest.ErrorCheck(t,codeartifact.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckRepositoryPermissionsDestroy(ctx),
Steps:[]resource.TestStep{
{
Config:testAccRepositoryPermissionsPolicyConfig_basic(rName),
Check:resource.ComposeTestCheckFunc(
testAccCheckRepositoryPermissionsExists(ctx,resourceName),
resource.TestCheckResourceAttrPair(resourceName,"resource_arn","aws_codeartifact_repository.test","arn"),
resource.TestCheckResourceAttr(resourceName,"domain",rName),
resource.TestMatchResourceAttr(resourceName,"policy_document",regexache.MustCompile("codeartifact:CreateRepository")),
resource.TestCheckResourceAttrPair(resourceName,"domain_owner","aws_codeartifact_domain.test","owner"),
),
},
{
ResourceName:ceName,
ImportState:
ImportStateVerify:true,
},
{
Config:testAccRepositoryPermissionsPolicyConfig_updated(rName),
Check:resource.ComposeTestCheckFunc(
testAccCheckRepositoryPermissionsExists(ctx,resourceName),
resource.TestCheckResourceAttrPair(resourceName,"resource_arn","aws_codeartifact_repository.test","arn"),
resource.TestCheckResourceAttr(resourceName,"domain",rName),
resource.TestMatchResourceAttr(resourceName,"policy_document",regexache.MustCompile("codeartifact:CreateRepository")),
resource.TestMatchResourceAttr(resourceName,"policy_document",regexache.MustCompile("codeartifact:ListRepositoriesInDomain")),
resource.TestCheckResourceAttrPair(resourceName,"domain_owner","aws_codeartifact_domain.test","owner"),
),
},
},
})
}
functestAccRepositoryPermissionsPolicy_ignoreEquivalent(t*testing.T){
ctx:=acctest.Context(t)
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
resourceName:="aws_codeartifact_repository_permissions_policy.test"

resource.Test(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t);acctest.PreCheckPartitionHasService(t,codeartifact.EndpointsID)},
ErrorCheck:acctest.ErrorCheck(t,codeartifact.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckRepositoryPermissionsDestroy(ctx),
Steps:[]resource.TestStep{
{
Config:testAccRepositoryPermissionsPolicyConfig_order(rName),
Check:resource.ComposeTestCheckFunc(
testAccCheckRepositoryPermissionsExists(ctx,resourceName),
resource.TestCheckResourceAttrPair(resourceName,"resource_arn","aws_codeartifact_repository.test","arn"),
resource.TestCheckResourceAttr(resourceName,"domain",rName),
resource.TestMatchResourceAttr(resourceName,"policy_document",regexache.MustCompile("codeartifact:CreateRepository")),
resource.TestCheckResourceAttrPair(resourceName,"domain_owner","aws_codeartifact_domain.test","owner"),
),
},
{
Config:testAccRepositoryPermissionsPolicyConfig_newOrder(rName),
PlanOnly:true,
},
},
})
}
functestAccRepositoryPermissionsPolicy_owner(t*testing.T){
ctx:=acctest.Context(t)
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
resourceName:="aws_codeartifact_repository_permissions_policy.test"

resource.Test(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t);acctest.PreCheckPartitionHasService(t,codeartifact.EndpointsID)},
ErrorCheck:acctest.ErrorCheck(t,codeartifact.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckRepositoryPermissionsDestroy(ctx),
Steps:[]resource.TestStep{
{
Config:testAccRepositoryPermissionsPolicyConfig_owner(rName),
Check:resource.ComposeTestCheckFunc(
testAccCheckRepositoryPermissionsExists(ctx,resourceName),
resource.TestCheckResourceAttrPair(resourceName,"resource_arn","aws_codeartifact_repository.test","arn"),
resource.TestCheckResourceAttr(resourceName,"domain",rName),
resource.TestMatchResourceAttr(resourceName,"policy_document",regexache.MustCompile("codeartifact:CreateRepository")),
resource.TestCheckResourceAttrPair(resourceName,"domain_owner","aws_codeartifact_domain.test","owner"),
),
},
{
ResourceName:ceName,
ImportState:
ImportStateVerify:true,
},
},
})
}
functestAccRepositoryPermissionsPolicy_disappears(t*testing.T){
ctx:=acctest.Context(t)
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
resourceName:="aws_codeartifact_repository_permissions_policy.test"

resource.Test(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t);acctest.PreCheckPartitionHasService(t,codeartifact.EndpointsID)},
ErrorCheck:acctest.ErrorCheck(t,codeartifact.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckRepositoryPermissionsDestroy(ctx),
Steps:[]resource.TestStep{
{
Config:testAccRepositoryPermissionsPolicyConfig_basic(rName),
Check:resource.ComposeTestCheckFunc(
testAccCheckRepositoryPermissionsExists(ctx,resourceName),
acctest.CheckResourceDisappears(ctx,acctest.Provider,tfcodeartifact.ResourceRepositoryPermissionsPolicy(),resourceName),
),
ExpectNonEmptyPlan:true,
},
},
})
}
functestAccRepositoryPermissionsPolicy_Disappears_domain(t*testing.T){
ctx:=acctest.Context(t)
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
resourceName:="aws_codeartifact_repository_permissions_policy.test"

resource.Test(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t);acctest.PreCheckPartitionHasService(t,codeartifact.EndpointsID)},
ErrorCheck:acctest.ErrorCheck(t,codeartifact.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckRepositoryPermissionsDestroy(ctx),
Steps:[]resource.TestStep{
{
Config:testAccRepositoryPermissionsPolicyConfig_basic(rName),
Check:resource.ComposeTestCheckFunc(
testAccCheckRepositoryPermissionsExists(ctx,resourceName),
acctest.CheckResourceDisappears(ctx,acctest.Provider,tfcodeartifact.ResourceRepositoryPermissionsPolicy(),resourceName),
),
ExpectNonEmptyPlan:true,
},
},
})
}
functestAccCheckRepositoryPermissionsExists(ctxcontext.Context,nstring)resource.TestCheckFunc{
returnfunc(s*terraform.State)error{
rs,ok:=s.RootModule().Resources[n]
if!ok{
returnfmt.Errorf("Notfound:%s",n)
}

ifrs.Primary.ID==""{
returnfmt.Errorf("noCodeArtifactdomainset")
}

conn:=acctest.Provider.Meta().(*conns.AWSClient).CodeArtifactConn(ctx)

domainOwner,domainName,repoName,err:=tfcodeartifact.DecodeRepositoryID(rs.Primary.ID)
iferr!=nil{
returnerr
}

_,err=conn.GetRepositoryPermissionsPolicyWithContext(ctx,&codeartifact.GetRepositoryPermissionsPolicyInput{
Domain:ring(domainName),
DomainOwner:aws.String(domainOwner),
Repository:aws.String(repoName),
})

returnerr
}
}
functestAccCheckRepositoryPermissionsDestroy(ctxcontext.Context)resource.TestCheckFunc{
returnfunc(s*terraform.State)error{
for_,rs:=ranges.RootModule().Resources{
ifrs.Type!="aws_codeartifact_repository_permissions_policy"{
continue
}

conn:=acctest.Provider.Meta().(*conns.AWSClient).CodeArtifactConn(ctx)

domainOwner,domainName,repoName,err:=tfcodeartifact.DecodeRepositoryID(rs.Primary.ID)
iferr!=nil{
returnerr
}

resp,err:=conn.GetRepositoryPermissionsPolicyWithContext(ctx,&codeartifact.GetRepositoryPermissionsPolicyInput{
Domain:ring(domainName),
DomainOwner:aws.String(domainOwner),
Repository:aws.String(repoName),
})

iferr==nil{
ifaws.StringValue(resp.Policy.ResourceArn)==rs.Primary.ID{
returnfmt.Errorf("CodeArtifactDomain%sstillexists",rs.Primary.ID)
}
}

iftfawserr.ErrCodeEquals(err,codeartifact.ErrCodeResourceNotFoundException){
returnnil
}

returnerr
}

returnnil
}
}
functestAccRepositoryPermissionsPolicyConfig_basic(rNamestring)string{
returnfmt.Sprintf(`
resource"aws_kms_key""test"{
description=%[1]q
deletion_window_in_days=7
}

resource"aws_codeartifact_domain""test"{
domain[1]q
encryption_key=aws_kms_key.test.arn
}

resource"aws_codeartifact_repository""test"{
repository=%[1]q
domain=aws_codeartifact_domain.test.domain
}

resource"aws_codeartifact_repository_permissions_policy""test"{
domainaws_codeartifact_domain.test.domain
repositorycodeartifact_repository.test.repository
policy_document=<<EOF
{
"Version":"2012-10-17",
"Statement":[

odeartifact:CreateRepository",
llow",
"*",
"${aws_codeartifact_domain.test.arn}"

]
}
EOF
}
`,rName)
}
functestAccRepositoryPermissionsPolicyConfig_owner(rNamestring)string{
returnfmt.Sprintf(`
resource"aws_kms_key""test"{
description=%[1]q
deletion_window_in_days=7
}

resource"aws_codeartifact_domain""test"{
domain[1]q
encryption_key=aws_kms_key.test.arn
}

resource"aws_codeartifact_repository""test"{
repository=%[1]q
domain=aws_codeartifact_domain.test.domain
}

resource"aws_codeartifact_repository_permissions_policy""test"{
domainaws_codeartifact_domain.test.domain
domain_owner=aws_codeartifact_domain.test.owner
repositorycodeartifact_repository.test.repository
policy_document=<<EOF
{
"Version":"2012-10-17",
"Statement":[

odeartifact:CreateRepository",
llow",
"*",
"${aws_codeartifact_domain.test.arn}"

]
}
EOF
}
`,rName)
}
functestAccRepositoryPermissionsPolicyConfig_updated(rNamestring)string{
returnfmt.Sprintf(`
resource"aws_kms_key""test"{
description=%[1]q
deletion_window_in_days=7
}

resource"aws_codeartifact_domain""test"{
domain[1]q
encryption_key=aws_kms_key.test.arn
}

resource"aws_codeartifact_repository""test"{
repository=%[1]q
domain=aws_codeartifact_domain.test.domain
}

resource"aws_codeartifact_repository_permissions_policy""test"{
domainaws_codeartifact_domain.test.domain
repositorycodeartifact_repository.test.repository
policy_document=<<EOF
{
"Version":"2012-10-17",
"Statement":[


"codeartifact:CreateRepository",
"codeartifact:ListRepositoriesInDomain"
],
llow",
"*",
"${aws_codeartifact_domain.test.arn}"

]
}
EOF
}
`,rName)
}
functestAccRepositoryPermissionsPolicyConfig_order(rNamestring)string{
returnfmt.Sprintf(`
resource"aws_kms_key""test"{
description=%[1]q
deletion_window_in_days=7
}

resource"aws_codeartifact_domain""test"{
domain[1]q
encryption_key=aws_kms_key.test.arn
}

resource"aws_codeartifact_repository""test"{
repository=%[1]q
domain=aws_codeartifact_domain.test.domain
}

resource"aws_codeartifact_repository_permissions_policy""test"{
domain=aws_codeartifact_domain.test.domain
repository=aws_codeartifact_repository.test.repository
policy_document=jsonencode({
Version="2012-10-17"
Statement=[{
=[
eartifact:CreateRepository",
eartifact:ListRepositoriesInDomain",

="Allow"
pal="*"
ce=aws_codeartifact_domain.test.arn
}]
})
}
`,rName)
}
functestAccRepositoryPermissionsPolicyConfig_newOrder(rNamestring)string{
returnfmt.Sprintf(`
resource"aws_kms_key""test"{
description=%[1]q
deletion_window_in_days=7
}

resource"aws_codeartifact_domain""test"{
domain[1]q
encryption_key=aws_kms_key.test.arn
}

resource"aws_codeartifact_repository""test"{
repository=%[1]q
domain=aws_codeartifact_domain.test.domain
}

resource"aws_codeartifact_repository_permissions_policy""test"{
domain=aws_codeartifact_domain.test.domain
repository=aws_codeartifact_repository.test.repository
policy_document=jsonencode({
Version="2012-10-17"
Statement=[{
=[
eartifact:ListRepositoriesInDomain",
eartifact:CreateRepository",

="Allow"
pal="*"
ce=aws_codeartifact_domain.test.arn
}]
})
}
`,rName)
}
