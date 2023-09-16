//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagevpclattice_test

import(
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/vpclattice"
	sdkacctest"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tfvpclattice"github.com/hashicorp/terraform-provider-aws/internal/service/vpclattice"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

funcTestAccVPCLatticeServiceNetworkVPCAssociation_basic(t*testing.T){
	ctx:=acctest.Context(t)

	varservicenetworkvpcascvpclattice.GetServiceNetworkVpcAssociationOutput
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName:="aws_vpclattice_service_network_vpc_association.test"

	resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){
	acctest.PreCheck(ctx,t)
	acctest.PreCheckPartitionHasService(t,names.VPCLatticeEndpointID)
	testAccPreCheck(ctx,t)
},
ErrorCheck:acctest.ErrorCheck(t,names.VPCLatticeEndpointID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckServiceNetworkVPCAssociationDestroy(ctx),
Steps:[]resource.TestStep{
	{
Config:testAccServiceNetworkVPCAssociationConfig_basic(rName),
Check:resource.ComposeTestCheckFunc(
	testAccCheckServiceNetworkVPCAssociationExists(ctx,resourceName,&servicenetworkvpcasc),
	acctest.MatchResourceAttrRegionalARN(resourceName,"arn","vpc-lattice",regexache.MustCompile("servicenetworkvpcassociation/.+$")),
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

funcTestAccVPCLatticeServiceNetworkVPCAssociation_arn(t*testing.T){
	ctx:=acctest.Context(t)

	varservicenetworkvpcascvpclattice.GetServiceNetworkVpcAssociationOutput
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName:="aws_vpclattice_service_network_vpc_association.test"

	resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){
	acctest.PreCheck(ctx,t)
	acctest.PreCheckPartitionHasService(t,names.VPCLatticeEndpointID)
	testAccPreCheck(ctx,t)
},
ErrorCheck:acctest.ErrorCheck(t,names.VPCLatticeEndpointID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckServiceNetworkVPCAssociationDestroy(ctx),
Steps:[]resource.TestStep{
	{
Config:testAccServiceNetworkVPCAssociationConfig_arn(rName),
Check:resource.ComposeTestCheckFunc(
	testAccCheckServiceNetworkVPCAssociationExists(ctx,resourceName,&servicenetworkvpcasc),
	acctest.MatchResourceAttrRegionalARN(resourceName,"arn","vpc-lattice",regexache.MustCompile("servicenetworkvpcassociation/.+$")),
	resource.TestCheckResourceAttrSet(resourceName,"service_network_identifier"),
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

funcTestAccVPCLatticeServiceNetworkVPCAssociation_disappears(t*testing.T){
	ctx:=acctest.Context(t)

	varservicenetworkvpcascvpclattice.GetServiceNetworkVpcAssociationOutput
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName:="aws_vpclattice_service_network_vpc_association.test"

	resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){
	acctest.PreCheck(ctx,t)
	acctest.PreCheckPartitionHasService(t,names.VPCLatticeEndpointID)
	testAccPreCheck(ctx,t)
},
ErrorCheck:acctest.ErrorCheck(t,names.VPCLatticeEndpointID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckServiceNetworkVPCAssociationDestroy(ctx),
Steps:[]resource.TestStep{
	{
Config:testAccServiceNetworkVPCAssociationConfig_basic(rName),
Check:resource.ComposeTestCheckFunc(
	testAccCheckServiceNetworkVPCAssociationExists(ctx,resourceName,&servicenetworkvpcasc),
	acctest.CheckResourceDisappears(ctx,acctest.Provider,tfvpclattice.ResourceServiceNetworkVPCAssociation(),resourceName),
),
ExpectNonEmptyPlan:true,
	},
},
	})
}

funcTestAccVPCLatticeServiceNetworkVPCAssociation_full(t*testing.T){
	ctx:=acctest.Context(t)

	varservicenetworkvpcascvpclattice.GetServiceNetworkVpcAssociationOutput
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName:="aws_vpclattice_service_network_vpc_association.test"

	resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){
	acctest.PreCheck(ctx,t)
	acctest.PreCheckPartitionHasService(t,names.VPCLatticeEndpointID)
	testAccPreCheck(ctx,t)
},
ErrorCheck:acctest.ErrorCheck(t,names.VPCLatticeEndpointID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckServiceNetworkVPCAssociationDestroy(ctx),
Steps:[]resource.TestStep{
	{
Config:testAccServiceNetworkVPCAssociationConfig_full(rName),
Check:resource.ComposeTestCheckFunc(
	testAccCheckServiceNetworkVPCAssociationExists(ctx,resourceName,&servicenetworkvpcasc),
	acctest.MatchResourceAttrRegionalARN(resourceName,"arn","vpc-lattice",regexache.MustCompile("servicenetworkvpcassociation/.+$")),
	resource.TestCheckResourceAttrSet(resourceName,"service_network_identifier"),
	resource.TestCheckResourceAttrSet(resourceName,"vpc_identifier"),
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

funcTestAccVPCLatticeServiceNetworkVPCAssociation_tags(t*testing.T){
	ctx:=acctest.Context(t)
	varservicenetworkvpcasc1,servicenetworkvpcasc2,service3vpclattice.GetServiceNetworkVpcAssociationOutput
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName:="aws_vpclattice_service_network_vpc_association.test"

	resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){
	acctest.PreCheck(ctx,t)
	acctest.PreCheckPartitionHasService(t,names.VPCLatticeEndpointID)
	testAccPreCheck(ctx,t)
},
ErrorCheck:acctest.ErrorCheck(t,names.VPCLatticeEndpointID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckServiceNetworkVPCAssociationDestroy(ctx),
Steps:[]resource.TestStep{
	{
Config:testAccServiceNetworkVPCAssociationConfig_tags1(rName,"key1","value1"),
Check:resource.ComposeTestCheckFunc(
	testAccCheckServiceNetworkVPCAssociationExists(ctx,resourceName,&servicenetworkvpcasc1),
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
Config:testAccServiceNetworkVPCAssociationConfig_tags2(rName,"key1","value1updated","key2","value2"),
Check:resource.ComposeTestCheckFunc(
	testAccCheckServiceNetworkVPCAssociationExists(ctx,resourceName,&servicenetworkvpcasc2),
	resource.TestCheckResourceAttr(resourceName,"tags.%","2"),
	resource.TestCheckResourceAttr(resourceName,"tags.key1","value1updated"),
	resource.TestCheckResourceAttr(resourceName,"tags.key2","value2"),
),
	},
	{
Config:testAccServiceNetworkVPCAssociationConfig_tags1(rName,"key2","value2"),
Check:resource.ComposeTestCheckFunc(
	testAccCheckServiceNetworkVPCAssociationExists(ctx,resourceName,&service3),
	resource.TestCheckResourceAttr(resourceName,"tags.%","1"),
	resource.TestCheckResourceAttr(resourceName,"tags.key2","value2"),
),
	},
},
	})
}

functestAccCheckServiceNetworkVPCAssociationDestroy(ctxcontext.Context)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
conn:=acctest.Provider.Meta().(*conns.AWSClient).VPCLatticeClient(ctx)

for_,rs:=ranges.RootModule().Resources{
	ifrs.Type!="aws_vpclattice_service_network_vpc_association"{
continue
	}

	_,err:=tfvpclattice.FindServiceNetworkVPCAssociationByID(ctx,conn,rs.Primary.ID)

	iftfresource.NotFound(err){
continue
	}

	iferr!=nil{
returnerr
	}

	returnfmt.Errorf("VPCLatticeServiceNetworkVPCAssociation%sstillexists",rs.Primary.ID)
}

returnnil
	}
}

functestAccCheckServiceNetworkVPCAssociationExists(ctxcontext.Context,namestring,service*vpclattice.GetServiceNetworkVpcAssociationOutput)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
rs,ok:=s.RootModule().Resources[name]
if!ok{
	returncreate.Error(names.VPCLattice,create.ErrActionCheckingExistence,tfvpclattice.ResNameService,name,errors.New("notfound"))
}

ifrs.Primary.ID==""{
	returncreate.Error(names.VPCLattice,create.ErrActionCheckingExistence,tfvpclattice.ResNameService,name,errors.New("notset"))
}

conn:=acctest.Provider.Meta().(*conns.AWSClient).VPCLatticeClient(ctx)
resp,err:=tfvpclattice.FindServiceNetworkVPCAssociationByID(ctx,conn,rs.Primary.ID)

iferr!=nil{
	returnerr
}

*service=*resp

returnnil
	}
}

functestAccServiceNetworkVPCAssociationConfig_base(rNamestring)string{
	returnacctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName,0),fmt.Sprintf(`
resource"aws_vpclattice_service_network""test"{
name=%[1]q
}
`,rName))
}

functestAccServiceNetworkVPCAssociationConfig_basic(rNamestring)string{
	returnacctest.ConfigCompose(testAccServiceNetworkVPCAssociationConfig_base(rName),`
resource"aws_vpclattice_service_network_vpc_association""test"{
vpc_identifier=aws_vpc.test.id
service_network_identifier=aws_vpclattice_service_network.test.id
}
`)
}

functestAccServiceNetworkVPCAssociationConfig_arn(rNamestring)string{
	returnacctest.ConfigCompose(testAccServiceNetworkVPCAssociationConfig_base(rName),`
resource"aws_vpclattice_service_network_vpc_association""test"{
vpc_identifier=aws_vpc.test.id
service_network_identifier=aws_vpclattice_service_network.test.arn
}
`)
}

functestAccServiceNetworkVPCAssociationConfig_full(rNamestring)string{
	returnacctest.ConfigCompose(testAccServiceNetworkVPCAssociationConfig_base(rName),fmt.Sprintf(`
resource"aws_security_group""test"{
name=%[1]q
vpc_id=aws_vpc.test.id

tags={
Name=%[1]q
}
}

resource"aws_vpclattice_service_network_vpc_association""test"{
vpc_identifier=aws_vpc.test.id
security_group_ids=[aws_security_group.test.id]
service_network_identifier=aws_vpclattice_service_network.test.id
}
`,rName))
}

functestAccServiceNetworkVPCAssociationConfig_tags1(rName,tagKey1,tagValue1string)string{
	returnacctest.ConfigCompose(testAccServiceNetworkVPCAssociationConfig_base(rName),fmt.Sprintf(`
resource"aws_vpclattice_service_network_vpc_association""test"{
vpc_identifier=aws_vpc.test.id
service_network_identifier=aws_vpclattice_service_network.test.id

tags={
%[1]q=%[2]q
}
}
`,tagKey1,tagValue1))
}

functestAccServiceNetworkVPCAssociationConfig_tags2(rName,tagKey1,tagValue1,tagKey2,tagValue2string)string{
	returnacctest.ConfigCompose(testAccServiceNetworkVPCAssociationConfig_base(rName),fmt.Sprintf(`
resource"aws_vpclattice_service_network_vpc_association""test"{
vpc_identifier=aws_vpc.test.id
service_network_identifier=aws_vpclattice_service_network.test.id

tags={
%[1]q=%[2]q
%[3]q=%[4]q
}
}
`,tagKey1,tagValue1,tagKey2,tagValue2))
}
