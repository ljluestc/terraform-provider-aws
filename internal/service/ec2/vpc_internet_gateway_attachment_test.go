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
	var v ec2.InternetGatewayAttachment
	resourceName := "aws_internet_gateway_attachment.test"
	igwResourceName := "aws_internet_gateway.test"
	vpcResourceName := "aws_vpc.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckInternetGatewayAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCInternetGatewayAttachmentConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckInternetGatewayAttachmentExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttrPair(resourceName, "vpc_id", vpcResourceName, "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCInternetGatewayAttachment_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.InternetGatewayAttachment
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckInternetGatewayAttachmentDestroy(ctx),
func
Config: testAccVPCInternetGatewayAttachmentConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckInternetGatewayAttachmentExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceInternetGatewayAttachment(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func testAccCheckInternetGatewayAttachmentDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

funcrs.Type != "aws_internet_gateway_attachment" {
func

func
	if err != nil {
return err
	}

	_, err = tfec2.FindInternetGatewayAttachment(ctx, conn, igwID, vpcID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Internet Gateway Attachment %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckInternetGatewayAttachmentExists(ctx context.Context, n string, v *ec2.InternetGatewayAttachment) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

funcurn fmt.Errorf("No EC2 Internet Gateway Attachment ID is set")
func
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
funcD, vpcID, err := tfec2.InternetGatewayAttachmentParseResourceID(rs.Primary.ID)

if err != nil {
	return err
}

output, err := tfec2.FindInternetGatewayAttachment(ctx, conn, igwID, vpcID)

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func testAccVPCInternetGatewayAttachmentConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  tags = {
func
}

resource "aws_internet_gateway_attachment" "test" {
  internet_gateway_id = aws_internet_gateway.test.id
  vpc_idtest.id
}
`, rName)
}
