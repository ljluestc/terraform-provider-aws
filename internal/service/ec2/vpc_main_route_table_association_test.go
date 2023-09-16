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
	var rta ec2.RouteTableAssociation
	resourceName := "aws_main_route_table_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckMainRouteTableAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCMainRouteTableAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckMainRouteTableAssociationExists(ctx, resourceName, &rta),
func
	{
Config: testAccVPCMainRouteTableAssociationConfig_updated(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckMainRouteTableAssociationExists(ctx, resourceName, &rta),
),
func
	})
}


func testAccCheckMainRouteTableAssociationDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
func
funcrs.Type != "aws_main_route_table_association" {
continue
func
	_, err := tfec2.FindMainRouteTableAssociationByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("Main route table association %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckMainRouteTableAssociationExists(ctx context.Context, n string, v *ec2.RouteTableAssociation) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
func
funcurn fmt.Errorf("No ID is set")
}
func := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

association, err := tfec2.FindMainRouteTableAssociationByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *association

return nil
	}
}


func testAccMainRouteTableAssociationConfigBaseVPC(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

funcc_idtest.id
  cidr_block = "10.1.1.0/24"

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCMainRouteTableAssociationConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccMainRouteTableAssociationConfigBaseVPC(rName), fmt.Sprintf(`
resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
dr_block = "10.0.0.0/8"
teway_id = aws_internet_gateway.test.id
  }

  tags = {
func
}

resource "aws_main_route_table_association" "test" {
  vpc_id= aws_vpc.test.id
  route_table_id = aws_route_table.test.id
}
`, rName))
}


func testAccVPCMainRouteTableAssociationConfig_updated(rName string) string {
	return acctest.ConfigCompose(testAccMainRouteTableAssociationConfigBaseVPC(rName), fmt.Sprintf(`
# Need to keep the old route table around when we update the
# main_route_table_association, otherwise Terraform will try to destroy the
# route table too early, and will fail because it's still the main one
resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
dr_block = "10.0.0.0/8"
teway_id = aws_internet_gateway.test.id
  }
funcgs = {
me = %[1]q
  }
}

resource "aws_route_table" "test2" {
  vpc_id = aws_vpc.test.id

  route {
dr_block = "10.0.0.0/8"
teway_id = aws_internet_gateway.test.id
  }

  tags = {
me = %[1]q
  }
}

resource "aws_main_route_table_association" "test" {
  vpc_id= aws_vpc.test.id
  route_table_id = aws_route_table.test2.id
}
`, rName))
}
