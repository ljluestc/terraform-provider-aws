//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

//go:buildsweep
//+buildsweep

packagecodeartifact

import(
"fmt"
"log"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/codeartifact"
"github.com/hashicorp/go-multierror"
"github.com/hashicorp/terraform-plugin-testing/helper/resource"
"github.com/hashicorp/terraform-provider-aws/internal/sweep"
)
funcinit(){
resource.AddTestSweepers("aws_codeartifact_domain",&resource.Sweeper{
Name:"aws_codeartifact_domain",
F:sweepDomains,
})

resource.AddTestSweepers("aws_codeartifact_repository",&resource.Sweeper{
Name:"aws_codeartifact_repository",
F:sweepRepositories,
})
}
funcsweepDomains(regionstring)error{
ctx:=sweep.Context(region)
client,err:=sweep.SharedRegionalSweepClient(ctx,region)
iferr!=nil{
returnfmt.Errorf("errorgettingclient:%s",err)
}
conn:=client.CodeArtifactConn(ctx)
input:=&codeartifact.ListDomainsInput{}
varsweeperErrs*multierror.Error

err=conn.ListDomainsPagesWithContext(ctx,input,func(page*codeartifact.ListDomainsOutput,lastPagebool)bool{
for_,domainPtr:=rangepage.Domains{
ifdomainPtr==nil{
continue
}

domain:=aws.StringValue(domainPtr.Name)
input:=&codeartifact.DeleteDomainInput{
Domain:domainPtr.Name,
}

log.Printf("[INFO]DeletingCodeArtifactDomain:%s",domain)

_,err:=conn.DeleteDomainWithContext(ctx,input)

iferr!=nil{
sweeperErr:=fmt.Errorf("errordeletingCodeArtifactDomain(%s):%w",domain,err)
log.Printf("[ERROR]%s",sweeperErr)
sweeperErrs=multierror.Append(sweeperErrs,sweeperErr)
}
}

return!lastPage
})

ifsweep.SkipSweepError(err){
log.Printf("[WARN]SkippingCodeArtifactDomainsweepfor%s:%s",region,err)
returnnil
}

iferr!=nil{
returnfmt.Errorf("errorlistingCodeArtifactDomains:%w",err)
}

returnsweeperErrs.ErrorOrNil()
}
funcsweepRepositories(regionstring)error{
ctx:=sweep.Context(region)
client,err:=sweep.SharedRegionalSweepClient(ctx,region)
iferr!=nil{
returnfmt.Errorf("errorgettingclient:%w",err)
}
conn:=client.CodeArtifactConn(ctx)
input:=&codeartifact.ListRepositoriesInput{}
varsweeperErrs*multierror.Error

err=conn.ListRepositoriesPagesWithContext(ctx,input,func(page*codeartifact.ListRepositoriesOutput,lastPagebool)bool{
for_,repositoryPtr:=rangepage.Repositories{
ifrepositoryPtr==nil{
continue
}

repository:=aws.StringValue(repositoryPtr.Name)
input:=&codeartifact.DeleteRepositoryInput{
Repository:repositoryPtr.Name,
Domain:toryPtr.DomainName,
DomainOwner:repositoryPtr.DomainOwner,
}

log.Printf("[INFO]DeletingCodeArtifactRepository:%s",repository)

_,err:=conn.DeleteRepositoryWithContext(ctx,input)

iferr!=nil{
sweeperErr:=fmt.Errorf("errordeletingCodeArtifactRepository(%s):%w",repository,err)
log.Printf("[ERROR]%s",sweeperErr)
sweeperErrs=multierror.Append(sweeperErrs,sweeperErr)
}
}

return!lastPage
})

ifsweep.SkipSweepError(err){
log.Printf("[WARN]SkippingCodeArtifactRepositorysweepfor%s:%s",region,err)
returnnil
}

iferr!=nil{
returnfmt.Errorf("errorlistingCodeArtifactRepositories:%w",err)
}

returnsweeperErrs.ErrorOrNil()
}
