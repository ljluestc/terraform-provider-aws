//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageservicecatalog_test

import(
"context"
"fmt"
"testing"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/servicecatalog"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
sdkacctest"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
"github.com/hashicorp/terraform-plugin-testing/helper/resource"
"github.com/hashicorp/terraform-plugin-testing/terraform"
"github.com/hashicorp/terraform-provider-aws/internal/acctest"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
tfservicecatalog"github.com/hashicorp/terraform-provider-aws/internal/service/servicecatalog"
)

//addsweepertodeleteknowntestservicecattagoptions

funcTestAccServiceCatalogTagOption_basic(t*testing.T){
ctx:=acctest.Context(t)
resourceName:="aws_servicecatalog_tag_option.test"
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,servicecatalog.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckTagOptionDestroy(ctx),
Steps:[]resource.TestStep{
{
Config:testAccTagOptionConfig_basic(rName,"värde","active=true"),
Check:resource.ComposeTestCheckFunc(
testAccCheckTagOptionExists(ctx,resourceName),
resource.TestCheckResourceAttr(resourceName,"active","true"),
resource.TestCheckResourceAttr(resourceName,"key",rName),
resource.TestCheckResourceAttrSet(resourceName,"owner"),
resource.TestCheckResourceAttr(resourceName,"value","värde"),
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

funcTestAccServiceCatalogTagOption_disappears(t*testing.T){
ctx:=acctest.Context(t)
resourceName:="aws_servicecatalog_tag_option.test"
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,servicecatalog.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckTagOptionDestroy(ctx),
Steps:[]resource.TestStep{
{
Config:testAccTagOptionConfig_basic(rName,"värde","active=true"),
Check:resource.ComposeTestCheckFunc(
testAccCheckTagOptionExists(ctx,resourceName),
acctest.CheckResourceDisappears(ctx,acctest.Provider,tfservicecatalog.ResourceTagOption(),resourceName),
),
ExpectNonEmptyPlan:true,
},
},
})
}

funcTestAccServiceCatalogTagOption_update(t*testing.T){
ctx:=acctest.Context(t)
resourceName:="aws_servicecatalog_tag_option.test"
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
rName2:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

//UpdateTagOption()isveryparticularaboutwhatitreceives.Onlyfieldsthatchangeshould
//beincludedoritwillthrowservicecatalog.ErrCodeDuplicateResourceException,"alreadyexists"

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,servicecatalog.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckTagOptionDestroy(ctx),
Steps:[]resource.TestStep{
{
Config:testAccTagOptionConfig_basic(rName,"värdeett",""),
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(resourceName,"active","true"),
resource.TestCheckResourceAttr(resourceName,"key",rName),
resource.TestCheckResourceAttrSet(resourceName,"owner"),
resource.TestCheckResourceAttr(resourceName,"value","värdeett"),
),
},
{
Config:testAccTagOptionConfig_basic(rName,"värdetvå","active=true"),
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(resourceName,"active","true"),
resource.TestCheckResourceAttr(resourceName,"key",rName),
resource.TestCheckResourceAttrSet(resourceName,"owner"),
resource.TestCheckResourceAttr(resourceName,"value","värdetvå"),
),
},
{
Config:testAccTagOptionConfig_basic(rName,"värdetvå","active=false"),
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(resourceName,"active","false"),
resource.TestCheckResourceAttr(resourceName,"key",rName),//cannotbeupdatedinplace
resource.TestCheckResourceAttrSet(resourceName,"owner"),
resource.TestCheckResourceAttr(resourceName,"value","värdetvå"),
),
},
{
Config:testAccTagOptionConfig_basic(rName,"värdetvå","active=true"),
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(resourceName,"active","true"),
resource.TestCheckResourceAttr(resourceName,"key",rName),//cannotbeupdatedinplace
resource.TestCheckResourceAttrSet(resourceName,"owner"),
resource.TestCheckResourceAttr(resourceName,"value","värdetvå"),
),
},
{
Config:testAccTagOptionConfig_basic(rName2,"värdeett","active=true"),
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(resourceName,"active","true"),
resource.TestCheckResourceAttr(resourceName,"key",rName2),
resource.TestCheckResourceAttrSet(resourceName,"owner"),
resource.TestCheckResourceAttr(resourceName,"value","värdeett"),
),
},
},
})
}

funcTestAccServiceCatalogTagOption_notActive(t*testing.T){
ctx:=acctest.Context(t)
resourceName:="aws_servicecatalog_tag_option.test"
rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,servicecatalog.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckTagOptionDestroy(ctx),
Steps:[]resource.TestStep{
{
Config:testAccTagOptionConfig_basic(rName,"värdeett","active=false"),
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(resourceName,"active","false"),
resource.TestCheckResourceAttr(resourceName,"key",rName),
resource.TestCheckResourceAttrSet(resourceName,"owner"),
resource.TestCheckResourceAttr(resourceName,"value","värdeett"),
),
},
},
})
}

functestAccCheckTagOptionDestroy(ctxcontext.Context)resource.TestCheckFunc{
returnfunc(s*terraform.State)error{
conn:=acctest.Provider.Meta().(*conns.AWSClient).ServiceCatalogConn(ctx)

for_,rs:=ranges.RootModule().Resources{
ifrs.Type!="aws_servicecatalog_tag_option"{
continue
}

input:=&servicecatalog.DescribeTagOptionInput{
Id:aws.String(rs.Primary.ID),
}

output,err:=conn.DescribeTagOptionWithContext(ctx,input)

iftfawserr.ErrCodeEquals(err,servicecatalog.ErrCodeResourceNotFoundException){
continue
}

iferr!=nil{
returnfmt.Errorf("errorgettingServiceCatalogTagOption(%s):%w",rs.Primary.ID,err)
}

ifoutput!=nil{
returnfmt.Errorf("ServiceCatalogTagOption(%s)stillexists",rs.Primary.ID)
}
}

returnnil
}
}

functestAccCheckTagOptionExists(ctxcontext.Context,resourceNamestring)resource.TestCheckFunc{
returnfunc(s*terraform.State)error{
rs,ok:=s.RootModule().Resources[resourceName]

if!ok{
returnfmt.Errorf("resourcenotfound:%s",resourceName)
}

conn:=acctest.Provider.Meta().(*conns.AWSClient).ServiceCatalogConn(ctx)

input:=&servicecatalog.DescribeTagOptionInput{
Id:aws.String(rs.Primary.ID),
}

_,err:=conn.DescribeTagOptionWithContext(ctx,input)

iferr!=nil{
returnfmt.Errorf("errordescribingServiceCatalogTagOption(%s):%w",rs.Primary.ID,err)
}

returnnil
}
}

functestAccTagOptionConfig_basic(key,value,activestring)string{
returnfmt.Sprintf(`
resource"aws_servicecatalog_tag_option""test"{
key=%[1]q
value=%[2]q
%[3]s
}
`,key,value,active)
}
