//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageds_test

import(
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/directoryservice"
	sdkacctest"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfds"github.com/hashicorp/terraform-provider-aws/internal/service/ds"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

funcTestAccDSSharedDirectory_basic(t*testing.T){
	ctx:=acctest.Context(t)
	varvdirectoryservice.SharedDirectory
	resourceName:="aws_directory_service_shared_directory.test"
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	domainName:=acctest.RandomDomainName()

	resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){
	acctest.PreCheck(ctx,t)
	acctest.PreCheckAlternateAccount(t)
},
ErrorCheck:acctest.ErrorCheck(t,directoryservice.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5FactoriesAlternate(ctx,t),
CheckDestroy:testAccCheckSharedDirectoryDestroy(ctx),
Steps:[]resource.TestStep{
	{
Config:testAccSharedDirectoryConfig_basic(rName,domainName),
Check:resource.ComposeTestCheckFunc(
	testAccCheckSharedDirectoryExists(ctx,resourceName,&v),
	resource.TestCheckResourceAttr(resourceName,"method","HANDSHAKE"),
	resource.TestCheckResourceAttr(resourceName,"notes","test"),
	resource.TestCheckResourceAttrSet(resourceName,"shared_directory_id"),
),
	},
	{
ResourceName:resourceName,
ImportState:true,
ImportStateVerify:true,
	},
},
	})
}

functestAccCheckSharedDirectoryExists(ctxcontext.Context,nstring,v*directoryservice.SharedDirectory)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
rs,ok:=s.RootModule().Resources[n]
if!ok{
	returnfmt.Errorf("Notfound:%s",n)
}

ifrs.Primary.ID==""{
	returnfmt.Errorf("NoDirectoryServiceSharedDirectoryIDisset")
}

conn:=acctest.Provider.Meta().(*conns.AWSClient).DSConn(ctx)

output,err:=tfds.FindSharedDirectory(ctx,conn,rs.Primary.Attributes["directory_id"],rs.Primary.Attributes["shared_directory_id"])

iferr!=nil{
	returnerr
}

*v=*output

returnnil
	}
}

functestAccCheckSharedDirectoryDestroy(ctxcontext.Context)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
conn:=acctest.Provider.Meta().(*conns.AWSClient).DSConn(ctx)

for_,rs:=ranges.RootModule().Resources{
	ifrs.Type!="aws_directory_service_shared_directory"{
continue
	}

	_,err:=tfds.FindSharedDirectory(ctx,conn,rs.Primary.Attributes["directory_id"],rs.Primary.Attributes["shared_directory_id"])

	iftfresource.NotFound(err){
continue
	}

	iferr!=nil{
returnerr
	}

	returnfmt.Errorf("DirectoryServiceSharedDirectory%sstillexists",rs.Primary.ID)
}

returnnil
	}
}

functestAccSharedDirectoryConfig_basic(rName,domainstring)string{
	returnacctest.ConfigCompose(
acctest.ConfigAlternateAccountProvider(),
testAccDirectoryConfig_microsoftStandard(rName,domain),
`
resource"aws_directory_service_shared_directory""test"{
directory_id=aws_directory_service_directory.test.id
notes="test"

target{
id=data.aws_caller_identity.receiver.account_id
}
}

data"aws_caller_identity""receiver"{
provider="awsalternate"
}
`)
}
