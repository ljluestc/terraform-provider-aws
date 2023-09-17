//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagequicksight_test

import(
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/quicksight"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tfquicksight"github.com/hashicorp/terraform-provider-aws/internal/service/quicksight"
	"github.com/hashicorp/terraform-provider-aws/names"
)


funcTestAccQuickSightVPCConnection_basic(t*testing.T){
	ctx:=acctest.Context(t)
	varvpcConnectionquicksight.VPCConnection
	resourceName:="aws_quicksight_vpc_connection.test"
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rId:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t,resource.TestCase{
PreCheck:
func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,quicksight.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckVPCConnectionDestroy(ctx),
Steps:[]resource.TestStep{
	{
Config:testAccVPCConnectionConfig_basic(rId,rName),
Check:resource.ComposeTestCheck
func(
	testAccCheckVPCConnectionExists(ctx,resourceName,&vpcConnection),
	acctest.CheckResourceAttrRegionalARN(resourceName,"arn","quicksight",fmt.Sprintf("vpcConnection/%[1]s",rId)),
	resource.TestCheckResourceAttr(resourceName,"vpc_connection_id",rId),
	resource.TestCheckResourceAttr(resourceName,"name",rName),
	resource.TestCheckResourceAttr(resourceName,"subnet_ids.#","2"),
	resource.TestCheckResourceAttr(resourceName,"security_group_ids.#","1"),
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


funcTestAccQuickSightVPCConnection_disappears(t*testing.T){
	ctx:=acctest.Context(t)
	varvpcConnectionquicksight.VPCConnection
	resourceName:="aws_quicksight_vpc_connection.test"
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rId:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t,resource.TestCase{
PreCheck:
func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,quicksight.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckVPCConnectionDestroy(ctx),
Steps:[]resource.TestStep{
	{
Config:testAccVPCConnectionConfig_basic(rId,rName),
Check:resource.ComposeTestCheck
func(
	testAccCheckVPCConnectionExists(ctx,resourceName,&vpcConnection),
	acctest.CheckFrameworkResourceDisappears(ctx,acctest.Provider,tfquicksight.ResourceVPCConnection,resourceName),
),
ExpectNonEmptyPlan:true,
	},
},
	})
}


funcTestAccQuickSightVPCConnection_tags(t*testing.T){
	ctx:=acctest.Context(t)
	varvpcConnectionquicksight.VPCConnection
	resourceName:="aws_quicksight_vpc_connection.test"
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rId:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t,resource.TestCase{
PreCheck:
func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,quicksight.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckVPCConnectionDestroy(ctx),
Steps:[]resource.TestStep{
	{
Config:testAccVPCConnectionConfig_tags1(rId,rName,"key1","value1"),
Check:resource.ComposeTestCheck
func(
	testAccCheckVPCConnectionExists(ctx,resourceName,&vpcConnection),
	resource.TestCheckResourceAttr(resourceName,"vpc_connection_id",rId),
	resource.TestCheckResourceAttr(resourceName,"name",rName),
	resource.TestCheckResourceAttr(resourceName,"tags.%","1"),
	resource.TestCheckResourceAttr(resourceName,"tags.key1","value1"),
),
	},
	{
ResourceName:resourceName,
ImportState:true,
ImportStateVerify:true,
	},
	{
Config:testAccVPCConnectionConfig_tags2(rId,rName,"key1","value1updated","key2","value2"),
Check:resource.ComposeTestCheck
func(
	testAccCheckVPCConnectionExists(ctx,resourceName,&vpcConnection),
	resource.TestCheckResourceAttr(resourceName,"vpc_connection_id",rId),
	resource.TestCheckResourceAttr(resourceName,"name",rName),
	resource.TestCheckResourceAttr(resourceName,"tags.%","2"),
	resource.TestCheckResourceAttr(resourceName,"tags.key1","value1updated"),
	resource.TestCheckResourceAttr(resourceName,"tags.key2","value2"),
),
	},
	{
Config:testAccVPCConnectionConfig_tags1(rId,rName,"key2","value2"),
Check:resource.ComposeTestCheck
func(
	testAccCheckVPCConnectionExists(ctx,resourceName,&vpcConnection),
	resource.TestCheckResourceAttr(resourceName,"vpc_connection_id",rId),
	resource.TestCheckResourceAttr(resourceName,"name",rName),
	resource.TestCheckResourceAttr(resourceName,"tags.%","1"),
	resource.TestCheckResourceAttr(resourceName,"tags.key2","value2"),
),
	},
},
	})
}


functestAccCheckVPCConnectionExists(ctxcontext.Context,resourceNamestring,vpcConnection*quicksight.VPCConnection)resource.TestCheck
func{
	return
func(s*terraform.State)error{
rs,ok:=s.RootModule().Resources[resourceName]
if!ok{
	returnfmt.Errorf("notfound:%s",resourceName)
}

conn:=acctest.Provider.Meta().(*conns.AWSClient).QuickSightConn(ctx)
output,err:=tfquicksight.FindVPCConnectionByID(ctx,conn,rs.Primary.ID)
iferr!=nil{
	returncreate.Error(names.QuickSight,create.ErrActionCheckingExistence,tfquicksight.ResNameVPCConnection,rs.Primary.ID,err)
}

*vpcConnection=*output

returnnil
	}
}


functestAccCheckVPCConnectionDestroy(ctxcontext.Context)resource.TestCheck
func{
	return
func(s*terraform.State)error{
conn:=acctest.Provider.Meta().(*conns.AWSClient).QuickSightConn(ctx)
for_,rs:=ranges.RootModule().Resources{
	ifrs.Type!="aws_quicksight_vpc_connection"{
continue
	}

	output,err:=tfquicksight.FindVPCConnectionByID(ctx,conn,rs.Primary.ID)
	iferr!=nil{
iftfawserr.ErrCodeEquals(err,quicksight.ErrCodeResourceNotFoundException){
	returnnil
}
returnerr
	}

	ifoutput!=nil&&aws.StringValue(output.Status)==quicksight.VPCConnectionResourceStatusDeleted{
returnnil
	}

	returncreate.Error(names.QuickSight,create.ErrActionCheckingDestroyed,tfquicksight.ResNameVPCConnection,rs.Primary.ID,err)
}

returnnil
	}
}


functestAccBaseVPCConnectionConfig(rNamestring)string{
	returnacctest.ConfigCompose(
acctest.ConfigVPCWithSubnets(rName,2),
`
resource"aws_security_group""test"{
vpc_id=aws_vpc.test.id
}

resource"aws_iam_role""test"{
assume_role_policy=jsonencode({
Version="2012-10-17"
Statement=[
{
Effect="Allow"
Action="sts:AssumeRole"
Principal={
Service="quicksight.amazonaws.com"
}
}
]
})
inline_policy{
name="QuicksightVPCConnectionRolePolicy"
policy=jsonencode({
Version="2012-10-17"
Statement=[
{
Effect="Allow"
Action=[
"ec2:CreateNetworkInterface",
"ec2:ModifyNetworkInterfaceAttribute",
"ec2:DeleteNetworkInterface",
"ec2:DescribeSubnets",
"ec2:DescribeSecurityGroups"
]
Resource=["*"]
}
]
})
}
}
`)
}


functestAccVPCConnectionConfig_basic(rIdstring,rNamestring)string{
	returnacctest.ConfigCompose(
testAccBaseVPCConnectionConfig(rName),
fmt.Sprintf(`
resource"aws_quicksight_vpc_connection""test"{
vpc_connection_id=%[1]q
name=%[2]q
role_arn=aws_iam_role.test.arn
security_group_ids=[
aws_security_group.test.id,
]
subnet_ids=aws_subnet.test[*].id
}
`,rId,rName))
}


functestAccVPCConnectionConfig_tags1(rId,rName,tagKey1,tagValue1string)string{
	returnacctest.ConfigCompose(
testAccBaseVPCConnectionConfig(rName),
fmt.Sprintf(`
resource"aws_quicksight_vpc_connection""test"{
vpc_connection_id=%[1]q
name=%[2]q
role_arn=aws_iam_role.test.arn
security_group_ids=[
aws_security_group.test.id,
]
subnet_ids=aws_subnet.test[*].id

tags={
%[3]q=%[4]q
}
}
`,rId,rName,tagKey1,tagValue1))
}


functestAccVPCConnectionConfig_tags2(rId,rName,tagKey1,tagValue1,tagKey2,tagValue2string)string{
	returnacctest.ConfigCompose(
testAccBaseVPCConnectionConfig(rName),
fmt.Sprintf(`
resource"aws_quicksight_vpc_connection""test"{
vpc_connection_id=%[1]q
name=%[2]q
role_arn=aws_iam_role.test.arn
security_group_ids=[
aws_security_group.test.id,
]
subnet_ids=aws_subnet.test[*].id

tags={
%[3]q=%[4]q
%[5]q=%[6]q
}
}
`,rId,rName,tagKey1,tagValue1,tagKey2,tagValue2))
}
