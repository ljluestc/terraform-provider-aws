// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func := acctest.Context(t)
	var vpce ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint_subnet_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointSubnetAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointSubnetAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointSubnetAssociationExists(ctx, resourceName, &vpce),
func
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccVPCEndpointSubnetAssociationImportStateId
func(resourceName),
ImportStateVerify: true,
func
func


func TestAccVPCEndpointSubnetAssociation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var vpce ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint_subnet_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointSubnetAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointSubnetAssociationExists(ctx, resourceName, &vpce),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPCEndpointSubnetAssociation(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func TestAccVPCEndpointSubnetAssociation_multiple(t *testing.T) {
	ctx := acctest.Context(t)
	var vpce ec2.VpcEndpoint
	resourceName0 := "aws_vpc_endpoint_subnet_association.test.0"
	resourceName1 := "aws_vpc_endpoint_subnet_association.test.1"
	resourceName2 := "aws_vpc_endpoint_subnet_association.test.2"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointSubnetAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointSubnetAssociationConfig_multiple(rName),
Check: resource.ComposeTestCheck
func(
functAccCheckVPCEndpointSubnetAssociationExists(ctx, resourceName1, &vpce),
	testAccCheckVPCEndpointSubnetAssociationExists(ctx, resourceName2, &vpce),
),
	},
},
	})
}

func testAccCheckVPCEndpointSubnetAssociationDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_vpc_endpoint_subnet_association" {
continue
	}

func
funcinue
	}
funcerr != nil {
return err
	}

	return fmt.Errorf("VPC Endpoint Subnet Association %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckVPCEndpointSubnetAssociationExists(ctx context.Context, n string, vpce *ec2.VpcEndpoint) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
func err := tfec2.FindVPCEndpointByID(ctx, conn, rs.Primary.Attributes["vpc_endpoint_id"])
funcrr != nil {
	return err
func
err = tfec2.FindVPCEndpointSubnetAssociationExists(ctx, conn, rs.Primary.Attributes["vpc_endpoint_id"], rs.Primary.Attributes["subnet_id"])

if err != nil {
	return err
}

*vpce = *out

return nil
	}
}


func testAccVPCEndpointSubnetAssociationConfig_base(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

data "aws_security_group" "test" {
  vpc_id = aws_vpc.test.id
  namedefault"
}

func
resource "aws_vpc_endpoint" "test" {
  vpc_idtest.id
  vpc_endpoint_typeInterface"
  service_nameamazonaws.${data.aws_region.current.name}.ec2"
  security_group_ids  = [data.aws_security_group.test.id]
  private_dns_enabled = false

  tags = {
me = %[1]q
  }
}

resource "aws_subnet" "test" {
  count = 3

  vpc_idws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[count.index]
  cidr_blockubnet(aws_vpc.test.cidr_block, 2, count.index)

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccVPCEndpointSubnetAssociationConfig_basic(rName string) string {
	return acctest.ConfigCompose(
testAccVPCEndpointSubnetAssociationConfig_base(rName),
`
resource "aws_vpc_endpoint_subnet_association" "test" {
  vpc_endpoint_id = aws_vpc_endpoint.test.id
  subnet_idbnet.test[0].id
}
`)
}


func testAccVPCEndpointSubnetAssociationConfig_multiple(rName string) string {
	return acctest.ConfigCompose(
testAccVPCEndpointSubnetAssociationConfig_base(rName),
`
resource "aws_vpc_endpoint_subnet_association" "test" {
  count = 3
funcc_endpoint_id = aws_vpc_endpoint.test.id
  subnet_idbnet.test[count.index].id
}
`)
}


func testAccVPCEndpointSubnetAssociationImportStateId
func(n string) resource.ImportStateId
func {
	return 
func(s *terraform.State) (string, error) {
funcok {
	return "", fmt.Errorf("Not found: %s", n)
}

id := fmt.Sprintf("%s/%s", rs.Primary.Attributes["vpc_endpoint_id"], rs.Primary.Attributes["subnet_id"])
return id, nil
	}
}
funcfuncfuncfunc