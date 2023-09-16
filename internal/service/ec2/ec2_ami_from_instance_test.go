// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)


func := acctest.Context(t)
	var image ec2.Image
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ami_from_instance.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMIFromInstanceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &image),
funcource.TestCheckResourceAttr(resourceName, "description", "Testing Terraform aws_ami_from_instance resource"),
	resource.TestCheckResourceAttr(resourceName, "usage_operation", "RunInstances"),
	resource.TestCheckResourceAttr(resourceName, "platform_details", "Linux/UNIX"),
	resource.TestCheckResourceAttr(resourceName, "image_type", "machine"),
	resource.TestCheckResourceAttr(resourceName, "hypervisor", "xen"),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
),
	},
},
	})
}


func TestAccEC2AMIFromInstance_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var image ec2.Image
funcourceName := "aws_ami_from_instance.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
func
Config: testAccAMIFromInstanceConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &image),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
func
Config: testAccAMIFromInstanceConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &image),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
func
Config: testAccAMIFromInstanceConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &image),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
func


func TestAccEC2AMIFromInstance_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var image ec2.Image
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ami_from_instance.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMIFromInstanceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
functest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceAMIFromInstance(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}

func testAccAMIFromInstanceBaseConfig(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_instance" "test" {
ami= data.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_type = data.aws_ec2_instance_type_offering.available.instance_type

tags = {
me = %[1]q
func
`, rName))
}


func testAccAMIFromInstanceConfig_basic(rName string) string {
	return acctest.ConfigCompose(
testAccAMIFromInstanceBaseConfig(rName),
fmt.Sprintf(`
resource "aws_ami_from_instance" "test" {
name= %[1]q
descriptioning Terraform aws_ami_from_instance resource"
source_instance_id = aws_instance.test.id
}
`, rName))
}

func testAccAMIFromInstanceConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(
testAccAMIFromInstanceBaseConfig(rName),
fmt.Sprintf(`
resource "aws_ami_from_instance" "test" {
name= %[1]q
descriptioning Terraform aws_ami_from_instance resource"
source_instance_id = aws_instance.test.id

tags = {
2]q = %[3]q
}
}
func


func testAccAMIFromInstanceConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(
testAccAMIFromInstanceBaseConfig(rName),
fmt.Sprintf(`
resource "aws_ami_from_instance" "test" {
name= %[1]q
descriptioning Terraform aws_ami_from_instance resource"
source_instance_id = aws_instance.test.id

tags = {
2]q = %[3]q
4]q = %[5]q
}
}
func
