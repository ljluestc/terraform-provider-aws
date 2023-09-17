//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageredshiftserverless_test

import(
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/redshiftserverless"
	sdkacctest"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfredshiftserverless"github.com/hashicorp/terraform-provider-aws/internal/service/redshiftserverless"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

funcTestAccRedshiftServerlessEndpointAccess_basic(t*testing.T){
	ctx:=acctest.Context(t)
	resourceName:="aws_redshiftserverless_endpoint_access.test"
	rName:=sdkacctest.RandStringFromCharSet(30,sdkacctest.CharSetAlpha)

	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){acctest.PreCheck(ctx,t)},
		ErrorCheck:acctest.ErrorCheck(t,redshiftserverless.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckEndpointAccessDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccEndpointAccessConfig_basic(rName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckEndpointAccessExists(ctx,resourceName),
					acctest.MatchResourceAttrRegionalARN(resourceName,"arn","redshift-serverless",regexache.MustCompile("managedvpcendpoint/.+$")),
					resource.TestCheckResourceAttr(resourceName,"workgroup_name",rName),
					resource.TestCheckResourceAttr(resourceName,"endpoint_name",rName),
					resource.TestCheckResourceAttrSet(resourceName,"port"),
					resource.TestCheckTypeSetElemAttrPair(resourceName,"subnet_ids.*","aws_subnet.test","id"),
				),
			},
			{
				ResourceName:resourceName,
				ImportState:true,
				ImportStateVerify:true,
			},
			{
				Config:testAccEndpointAccessConfig_updated(rName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckEndpointAccessExists(ctx,resourceName),
					acctest.MatchResourceAttrRegionalARN(resourceName,"arn","redshift-serverless",regexache.MustCompile("managedvpcendpoint/.+$")),
					resource.TestCheckResourceAttr(resourceName,"workgroup_name",rName),
					resource.TestCheckResourceAttr(resourceName,"endpoint_name",rName),
					resource.TestCheckResourceAttrSet(resourceName,"port"),
					resource.TestCheckTypeSetElemAttrPair(resourceName,"subnet_ids.*","aws_subnet.test","id"),
					resource.TestCheckTypeSetElemAttrPair(resourceName,"vpc_security_group_ids.*","aws_security_group.test","id"),
				),
			},
		},
	})
}

funcTestAccRedshiftServerlessEndpointAccess_disappears_workgroup(t*testing.T){
	ctx:=acctest.Context(t)
	resourceName:="aws_redshiftserverless_endpoint_access.test"
	rName:=sdkacctest.RandStringFromCharSet(30,sdkacctest.CharSetAlpha)

	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){acctest.PreCheck(ctx,t)},
		ErrorCheck:acctest.ErrorCheck(t,redshiftserverless.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckEndpointAccessDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccEndpointAccessConfig_basic(rName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckEndpointAccessExists(ctx,resourceName),
					acctest.CheckResourceDisappears(ctx,acctest.Provider,tfredshiftserverless.ResourceWorkgroup(),"aws_redshiftserverless_workgroup.test"),
				),
				ExpectNonEmptyPlan:true,
			},
		},
	})
}

funcTestAccRedshiftServerlessEndpointAccess_disappears(t*testing.T){
	ctx:=acctest.Context(t)
	resourceName:="aws_redshiftserverless_endpoint_access.test"
	rName:=sdkacctest.RandStringFromCharSet(30,sdkacctest.CharSetAlpha)

	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){acctest.PreCheck(ctx,t)},
		ErrorCheck:acctest.ErrorCheck(t,redshiftserverless.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckEndpointAccessDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccEndpointAccessConfig_basic(rName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckEndpointAccessExists(ctx,resourceName),
					acctest.CheckResourceDisappears(ctx,acctest.Provider,tfredshiftserverless.ResourceEndpointAccess(),resourceName),
				),
				ExpectNonEmptyPlan:true,
			},
		},
	})
}

functestAccCheckEndpointAccessDestroy(ctxcontext.Context)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
		conn:=acctest.Provider.Meta().(*conns.AWSClient).RedshiftServerlessConn(ctx)

		for_,rs:=ranges.RootModule().Resources{
			ifrs.Type!="aws_redshiftserverless_endpoint_access"{
				continue
			}
			_,err:=tfredshiftserverless.FindEndpointAccessByName(ctx,conn,rs.Primary.ID)

			iftfresource.NotFound(err){
				continue
			}

			iferr!=nil{
				returnerr
			}

			returnfmt.Errorf("RedshiftServerlessEndpointAccess%sstillexists",rs.Primary.ID)
		}

		returnnil
	}
}

functestAccCheckEndpointAccessExists(ctxcontext.Context,namestring)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
		rs,ok:=s.RootModule().Resources[name]
		if!ok{
			returnfmt.Errorf("notfound:%s",name)
		}

		ifrs.Primary.ID==""{
			returnfmt.Errorf("RedshiftServerlessEndpointAccessIDisnotset")
		}

		conn:=acctest.Provider.Meta().(*conns.AWSClient).RedshiftServerlessConn(ctx)

		_,err:=tfredshiftserverless.FindEndpointAccessByName(ctx,conn,rs.Primary.ID)

		returnerr
	}
}

functestAccEndpointAccessConfig_basic(rNamestring)string{
	returnacctest.ConfigCompose(
		acctest.ConfigAvailableAZsNoOptIn(),
		fmt.Sprintf(`
resource"aws_vpc""test"{
cidr_block="10.0.0.0/16"
}

resource"aws_subnet""test"{
availability_zone=data.aws_availability_zones.available.names[0]
cidr_block=cidrsubnet(aws_vpc.test.cidr_block,8,0)
vpc_id=aws_vpc.test.id
}

resource"aws_redshiftserverless_namespace""test"{
namespace_name=%[1]q
}

resource"aws_redshiftserverless_workgroup""test"{
namespace_name=aws_redshiftserverless_namespace.test.namespace_name
workgroup_name=%[1]q
}

resource"aws_redshiftserverless_endpoint_access""test"{
workgroup_name=aws_redshiftserverless_workgroup.test.workgroup_name
endpoint_name=%[1]q
subnet_ids=[aws_subnet.test.id]
}
`,rName))
}

functestAccEndpointAccessConfig_updated(rNamestring)string{
	returnacctest.ConfigCompose(
		acctest.ConfigAvailableAZsNoOptIn(),
		fmt.Sprintf(`
resource"aws_vpc""test"{
cidr_block="10.0.0.0/16"
}

resource"aws_security_group""test"{
name=%[1]q
vpc_id=aws_vpc.test.id
}

resource"aws_subnet""test"{
availability_zone=data.aws_availability_zones.available.names[0]
cidr_block=cidrsubnet(aws_vpc.test.cidr_block,8,0)
vpc_id=aws_vpc.test.id
}

resource"aws_redshiftserverless_namespace""test"{
namespace_name=%[1]q
}

resource"aws_redshiftserverless_workgroup""test"{
namespace_name=aws_redshiftserverless_namespace.test.namespace_name
workgroup_name=%[1]q
}

resource"aws_redshiftserverless_endpoint_access""test"{
workgroup_name=aws_redshiftserverless_workgroup.test.workgroup_name
endpoint_name=%[1]q
subnet_ids=[aws_subnet.test.id]
vpc_security_group_ids=[aws_security_group.test.id]
}
`,rName))
}
