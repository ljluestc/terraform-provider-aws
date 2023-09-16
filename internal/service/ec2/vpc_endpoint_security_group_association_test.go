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
	var v ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint_security_group_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointSecurityGroupAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointSecurityGroupAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointSecurityGroupAssociationExists(ctx, resourceName, &v),
func
	},
},
	})
}


func TestAccVPCEndpointSecurityGroupAssociation_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.VpcEndpoint
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointSecurityGroupAssociationDestroy(ctx),
func
Config: testAccVPCEndpointSecurityGroupAssociationConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointSecurityGroupAssociationExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPCEndpointSecurityGroupAssociation(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccVPCEndpointSecurityGroupAssociation_multiple(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.VpcEndpoint
	resourceName0 := "aws_vpc_endpoint_security_group_association.test.0"
	resourceName1 := "aws_vpc_endpoint_security_group_association.test.1"
	resourceName2 := "aws_vpc_endpoint_security_group_association.test.2"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointSecurityGroupAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointSecurityGroupAssociationConfig_multiple(rName),
func(
	testAccCheckVPCEndpointSecurityGroupAssociationExists(ctx, resourceName0, &v),
	testAccCheckVPCEndpointSecurityGroupAssociationExists(ctx, resourceName1, &v),
	testAccCheckVPCEndpointSecurityGroupAssociationExists(ctx, resourceName2, &v),
	testAccCheckVPCEndpointSecurityGroupAssociationNumAssociations(&v, 4),
),
	},
},
func


func TestAccVPCEndpointSecurityGroupAssociation_replaceDefaultAssociation(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint_security_group_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointSecurityGroupAssociationDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointSecurityGroupAssociationConfig_replaceDefault(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointSecurityGroupAssociationExists(ctx, resourceName, &v),
func
	},
},
	})
}


func testAccCheckVPCEndpointSecurityGroupAssociationDestroy(ctx context.Context) resource.TestCheck
funcurn 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_vpc_endpoint_security_group_association" {
continue
	}

	err := tfec2.FindVPCEndpointSecurityGroupAssociationExists(ctx, conn, rs.Primary.Attributes["vpc_endpoint_id"], rs.Primary.Attributes["security_group_id"])
functfresource.NotFound(err) {
func

funcrn err
	}

	return fmt.Errorf("VPC Endpoint Security Group Association %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckVPCEndpointSecurityGroupAssociationExists(ctx context.Context, n string, v *ec2.VpcEndpoint) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No VPC Endpoint Security Group Association ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

func
funcurn err
}
func= tfec2.FindVPCEndpointSecurityGroupAssociationExists(ctx, conn, rs.Primary.Attributes["vpc_endpoint_id"], rs.Primary.Attributes["security_group_id"])

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func testAccCheckVPCEndpointSecurityGroupAssociationNumAssociations(v *ec2.VpcEndpoint, n int) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if len := len(v.Groups); len != n {
	return fmt.Errorf("got %d associations; wanted %d", len, n)
}

return nil
	}
}


func testAccVPCEndpointSecurityGroupAssociationConfig_base(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

func %[1]q
func

func
resource "aws_security_group" "test" {
count = 3

vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
}
}
funcurce "aws_vpc_endpoint" "test" {
vpc_idws_vpc.test.id
service_nameazonaws.${data.aws_region.current.name}.ec2"
vpc_endpoint_type = "Interface"

tags = {
me = %[1]q
}
}
`, rName)
}


func testAccVPCEndpointSecurityGroupAssociationConfig_basic(rName string) string {
	return acctest.ConfigCompose(
testAccVPCEndpointSecurityGroupAssociationConfig_base(rName),
`
resource "aws_vpc_endpoint_security_group_association" "test" {
vpc_endpoint_idws_vpc_endpoint.test.id
security_group_id = aws_security_group.test[0].id
}
`)
}


func testAccVPCEndpointSecurityGroupAssociationConfig_multiple(rName string) string {
	return acctest.ConfigCompose(
testAccVPCEndpointSecurityGroupAssociationConfig_base(rName),
`
resource "aws_vpc_endpoint_security_group_association" "test" {
count = length(aws_security_group.test)

vpc_endpoint_idws_vpc_endpoint.test.id
security_group_id = aws_security_group.test[count.index].id
}
func


func testAccVPCEndpointSecurityGroupAssociationConfig_replaceDefault(rName string) string {
	return acctest.ConfigCompose(
testAccVPCEndpointSecurityGroupAssociationConfig_base(rName),
`
resource "aws_vpc_endpoint_security_group_association" "test" {
vpc_endpoint_idws_vpc_endpoint.test.id
security_group_id = aws_security_group.test[0].id

replace_default_association = true
func
}
func