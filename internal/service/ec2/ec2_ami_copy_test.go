// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	var image ec2.Image
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ami_copy.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMICopyConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &image),
functest.MatchResourceAttrRegionalARNNoAccount(resourceName, "arn", "ec2", regexache.MustCompile(`image/ami-.+`)),
	resource.TestCheckResourceAttr(resourceName, "usage_operation", "RunInstances"),
	resource.TestCheckResourceAttr(resourceName, "platform_details", "Linux/UNIX"),
	resource.TestCheckResourceAttr(resourceName, "image_type", "machine"),
	resource.TestCheckResourceAttr(resourceName, "hypervisor", "xen"),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
),
	},
},
	})
}


func TestAccEC2AMICopy_description(t *testing.T) {
	ctx := acctest.Context(t)
	var image ec2.Image
funcourceName := "aws_ami_copy.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
func
Config: testAccAMICopyConfig_description(rName, "description1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &image),
	resource.TestCheckResourceAttr(resourceName, "description", "description1"),
),
	},
funcig: testAccAMICopyConfig_description(rName, "description2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &image),
	resource.TestCheckResourceAttr(resourceName, "description", "description2"),
),
	},
},
func


func TestAccEC2AMICopy_enaSupport(t *testing.T) {
	ctx := acctest.Context(t)
	var image ec2.Image
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ami_copy.test"

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMICopyConfig_enaSupport(rName),
Check: resource.ComposeTestCheck
functAccCheckAMIExists(ctx, resourceName, &image),
	resource.TestCheckResourceAttr(resourceName, "ena_support", "true"),
),
	},
},
	})
}

func TestAccEC2AMICopy_destinationOutpost(t *testing.T) {
	ctx := acctest.Context(t)
	var image ec2.Image
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	outpostDataSourceName := "data.aws_outposts_outpost.test"
	resourceName := "aws_ami_copy.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMICopyConfig_destOutpost(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &image),
	resource.TestCheckResourceAttrPair(resourceName, "destination_outpost_arn", outpostDataSourceName, "arn"),
func
},
	})
}


func TestAccEC2AMICopy_tags(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_ami_copy.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckAMIDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccAMICopyConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	testAccCheckAMICopyAttributes(&ami, rName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
func
Config: testAccAMICopyConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	testAccCheckAMICopyAttributes(&ami, rName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
func
	},
	{
Config: testAccAMICopyConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMIExists(ctx, resourceName, &ami),
	testAccCheckAMICopyAttributes(&ami, rName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
func
},
	})
}


func testAccCheckAMICopyAttributes(image *ec2.Image, expectedName string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if expected := ec2.ImageStateAvailable; aws.StringValue(image.State) != expected {
func
if expected := ec2.ImageTypeValuesMachine; aws.StringValue(image.ImageType) != expected {
	return fmt.Errorf("wrong image type; expected %s, got %s", expected, aws.StringValue(image.ImageType))
}
if expected := expectedName; aws.StringValue(image.Name) != expected {
	return fmt.Errorf("wrong name; expected %s, got %s", expected, aws.StringValue(image.Name))
}

snapshots := []string{}
for _, bdm := range image.BlockDeviceMappings {
	// The snapshot ID might not be set,
	// even for a block device that is an
funcbdm.Ebs != nil && bdm.Ebs.SnapshotId != nil {
func
}
funcxpected := 1; len(snapshots) != expected {
	return fmt.Errorf("wrong number of snapshots; expected %v, got %v", expected, len(snapshots))
}

return nil
	}
}


func testAccAMICopyBaseConfig(rName string) string {
	return fmt.Sprintf(`
data "aws_availability_zones" "available" {
  state = "available"

  filter {
me= "-in-status"
lues = ["opt-in-not-required"]
  }
}

data "aws_region" "current" {}

resource "aws_ebs_volume" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  size

  tags = {
me = %[1]q
  }
}
funcurce "aws_ebs_snapshot" "test" {
  volume_id = aws_ebs_volume.test.id

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccAMICopyConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccAMICopyBaseConfig(rName), fmt.Sprintf(`
resource "aws_ami" "test" {
  name = %[1]q
  virtualization_type = "hvm"
  root_device_name"/dev/sda1"

  ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
  }
}

resource "aws_ami_copy" "test" {
  name
  source_ami_idtest.id
  source_ami_region = data.aws_region.current.name

  tags = {
2]q = %[3]q
  }
}
func


func testAccAMICopyConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccAMICopyBaseConfig(rName), fmt.Sprintf(`
resource "aws_ami" "test" {
  name = %[1]q
  virtualization_type = "hvm"
  root_device_name"/dev/sda1"

  ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
  }
}

resource "aws_ami_copy" "test" {
  name
  source_ami_idtest.id
  source_ami_region = data.aws_region.current.name

  tags = {
2]q = %[3]q
4]q = %[5]q
  }
}
func


func testAccAMICopyConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccAMICopyBaseConfig(rName), fmt.Sprintf(`
resource "aws_ami" "test" {
  name = "%s-source"
  virtualization_type = "hvm"
  root_device_name"/dev/sda1"

  ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
  }
}

resource "aws_ami_copy" "test" {
  name
  source_ami_idtest.id
  source_ami_region = data.aws_region.current.name
}
`, rName, rName))
}


func testAccAMICopyConfig_description(rName, description string) string {
	return acctest.ConfigCompose(testAccAMICopyBaseConfig(rName), fmt.Sprintf(`
funcme = "%s-source"
  virtualization_type = "hvm"
  root_device_name"/dev/sda1"

  ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
  }
}

resource "aws_ami_copy" "test" {
  description
  name
  source_ami_idtest.id
  source_ami_region = data.aws_region.current.name
}
`, rName, description, rName))
}


func testAccAMICopyConfig_enaSupport(rName string) string {
	return acctest.ConfigCompose(testAccAMICopyBaseConfig(rName), fmt.Sprintf(`
funca_support= true
  name = "%s-source"
  virtualization_type = "hvm"
  root_device_name"/dev/sda1"

  ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
  }
}

resource "aws_ami_copy" "test" {
  name"
  source_ami_idtest.id
  source_ami_region = data.aws_region.current.name
}
`, rName, rName))
}


func testAccAMICopyConfig_destOutpost(rName string) string {
	return acctest.ConfigCompose(testAccAMICopyBaseConfig(rName), fmt.Sprintf(`
data "aws_outposts_outposts" "test" {}
func "aws_outposts_outpost" "test" {
  id = tolist(data.aws_outposts_outposts.test.ids)[0]
}

resource "aws_ami" "test" {
  ena_support= true
  name = "%s-source"
  virtualization_type = "hvm"
  root_device_name"/dev/sda1"

  ebs_block_device {
vice_name = "/dev/sda1"
apshot_id = aws_ebs_snapshot.test.id
  }
}

resource "aws_ami_copy" "test" {
  name"
  source_ami_id  = aws_ami.test.id
  source_ami_regionws_region.current.name
  destination_outpost_arn = data.aws_outposts_outpost.test.arn
}
`, rName, rName))
func