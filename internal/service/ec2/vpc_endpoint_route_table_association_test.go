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
	resourceName := "aws_vpc_endpoint_route_table_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointRouteTableAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointRouteTableAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointRouteTableAssociationExists(ctx, resourceName),
func
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccVPCEndpointRouteTableAssociationImportStateId
func(resourceName),
ImportStateVerify: true,
func
func


func TestAccVPCEndpointRouteTableAssociation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_vpc_endpoint_route_table_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointRouteTableAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointRouteTableAssociationExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPCEndpointRouteTableAssociation(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func testAccCheckVPCEndpointRouteTableAssociationDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_vpc_endpoint_route_table_association" {
func
func := tfec2.FindVPCEndpointRouteTableAssociationExists(ctx, conn, rs.Primary.Attributes["vpc_endpoint_id"], rs.Primary.Attributes["route_table_id"])

funcinue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("VPC Endpoint Route Table Association %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckVPCEndpointRouteTableAssociationExists(ctx context.Context, n string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No VPC Endpoint Route Table Association ID is set")
func
func
return tfec2.FindVPCEndpointRouteTableAssociationExists(ctx, conn, rs.Primary.Attributes["vpc_endpoint_id"], rs.Primary.Attributes["route_table_id"])
func


func testAccVPCEndpointRouteTableAssociationImportStateId
func(n string) resource.ImportStateId
func {
	return 
func(s *terraform.State) (string, error) {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return "", fmt.Errorf("Not found: %s", n)
}

id := fmt.Sprintf("%s/%s", rs.Primary.Attributes["vpc_endpoint_id"], rs.Primary.Attributes["route_table_id"])
return id, nil
	}
}
func
funcurn fmt.Sprintf(`
funcdr_block = "10.0.0.0/16"

func %[1]q
  }
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
  vpc_idc.test.id
  service_name = "com.amazonaws.${data.aws_region.current.name}.s3"

  tags = {
me = %[1]q
func

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_endpoint_route_table_association" "test" {
  vpc_endpoint_id = aws_vpc_endpoint.test.id
  route_table_id  = aws_route_table.test.id
}
`, rName)
}
