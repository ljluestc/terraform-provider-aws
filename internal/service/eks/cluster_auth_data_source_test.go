//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageeks_test

import(
"fmt"
"testing"

"github.com/aws/aws-sdk-go/service/eks"
"github.com/hashicorp/terraform-plugin-testing/helper/resource"
"github.com/hashicorp/terraform-plugin-testing/terraform"
"github.com/hashicorp/terraform-provider-aws/internal/acctest"
tfeks"github.com/hashicorp/terraform-provider-aws/internal/service/eks"
)

funcTestAccEKSClusterAuthDataSource_basic(t*testing.T){
ctx:=acctest.Context(t)
dataSourceResourceName:="data.aws_eks_cluster_auth.test"

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,eks.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
Steps:[]resource.TestStep{
{
Config:testAccClusterAuthDataSourceConfig_basic,
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(dataSourceResourceName,"name","foobar"),
resource.TestCheckResourceAttrSet(dataSourceResourceName,"token"),
testAccCheckClusterAuthToken(dataSourceResourceName),
),
},
},
})
}

functestAccCheckClusterAuthToken(nstring)resource.TestCheckFunc{
returnfunc(s*terraform.State)error{
rs,ok:=s.RootModule().Resources[n]
if!ok{
returnfmt.Errorf("Notfound:%s",n)
}

ifrs.Primary.ID==""{
returnfmt.Errorf("NoresourceIDisset")
}

name:=rs.Primary.Attributes["name"]
tok:=rs.Primary.Attributes["token"]
verifier:=tfeks.NewVerifier(name)
identity,err:=verifier.Verify(tok)
iferr!=nil{
returnfmt.Errorf("Errorverifyingtokenforcluster%q:%v",name,err)
}
ifidentity.ARN==""{
returnfmt.Errorf("UnexpectedblankARNfortokenidentity")
}

returnnil
}
}

consttestAccClusterAuthDataSourceConfig_basic=`
data"aws_eks_cluster_auth""test"{
name="foobar"
}
`
