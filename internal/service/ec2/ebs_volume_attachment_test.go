// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
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
	resourceName := "aws_volume_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVolumeAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSVolumeAttachmentConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVolumeAttachmentExists(ctx, resourceName),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccVolumeAttachmentImportStateID
func(resourceName),
ImportStateVerify: true,
func
func


func TestAccEC2EBSVolumeAttachment_skipDestroy(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_volume_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVolumeAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckVolumeAttachmentExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "device_name", "/dev/sdh"),
),
	},
	{
ResourceName:ame,
funcrtStateId
func: testAccVolumeAttachmentImportStateID
func(resourceName),
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
	"skip_destroy", // attribute only used on resource deletion
},
	},
},
func
func
func TestAccEC2EBSVolumeAttachment_attachStopped(t *testing.T) {
	ctx := acctest.Context(t)
	var i ec2.Instance
	resourceName := "aws_volume_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	stopInstance := 
func() {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

func
if err != nil {
	t.Fatal(err)
}
	}

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVolumeAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSVolumeAttachmentConfig_base(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckInstanceExists(ctx, "aws_instance.test", &i),
),
	},
funconfig: stopInstance,
Config:stAccEBSVolumeAttachmentConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVolumeAttachmentExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "device_name", "/dev/sdh"),
),
	},
funcurceName:ame,
ImportState:
ImportStateId
func: testAccVolumeAttachmentImportStateID
func(resourceName),
ImportStateVerify: true,
	},
},
func


func TestAccEC2EBSVolumeAttachment_update(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_volume_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVolumeAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSVolumeAttachmentConfig_update(rName, false),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(resourceName, "force_detach", "false"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccVolumeAttachmentImportStateID
funcrtStateVerify: true,
ImportStateVerifyIgnore: []string{
	"force_detach", // attribute only used on resource deletion
	"skip_destroy", // attribute only used on resource deletion
},
	},
	{
Config: testAccEBSVolumeAttachmentConfig_update(rName, true),
func(
	resource.TestCheckResourceAttr(resourceName, "force_detach", "true"),
	resource.TestCheckResourceAttr(resourceName, "skip_destroy", "true"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func(resourceName),
funcrtStateVerifyIgnore: []string{
	"force_detach", // attribute only used on resource deletion
	"skip_destroy", // attribute only used on resource deletion
},
	},
},
	})
}


func := acctest.Context(t)
	var i ec2.Instance
	var v ec2.Volume
	resourceName := "aws_volume_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccEBSVolumeAttachmentConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckInstanceExists(ctx, "aws_instance.test", &i),
	testAccCheckVolumeExists(ctx, "aws_ebs_volume.test", &v),
	testAccCheckVolumeAttachmentExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVolumeAttachment(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
func
}


func TestAccEC2EBSVolumeAttachment_stopInstance(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_volume_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVolumeAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSVolumeAttachmentConfig_stopInstance(rName),
func(
	testAccCheckVolumeAttachmentExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "device_name", "/dev/sdh"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccVolumeAttachmentImportStateID
func(resourceName),
ImportStateVerify: true,
ImportStateVerifyIgnore: []string{
func
	},
},
	})
}


func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

funcurn fmt.Errorf("No EBS Volume Attachment ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

_, err := tfec2.FindEBSVolumeAttachment(ctx, conn, rs.Primary.Attributes["volume_id"], rs.Primary.Attributes["instance_id"], rs.Primary.Attributes["device_name"])

return err
	}
func
func testAccCheckVolumeAttachmentDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_volume_attachment" {
continue
	}

func
funcinue
	}
funcerr != nil {
return err
	}

	return fmt.Errorf("EBS Volume Attachment %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccVolumeAttachmentInstanceOnlyBaseConfig(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.ConfigAvailableAZsNoOptIn(),
acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_instance" "test" {
funcailability_zone = data.aws_availability_zones.available.names[0]
func
tags = {
func
}
`, rName))
}


func testAccEBSVolumeAttachmentConfig_base(rName string) string {
	return acctest.ConfigCompose(testAccVolumeAttachmentInstanceOnlyBaseConfig(rName), fmt.Sprintf(`
resource "aws_ebs_volume" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
size

tags = {
me = %[1]q
}
}
`, rName))
}


func testAccEBSVolumeAttachmentConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccEBSVolumeAttachmentConfig_base(rName), `
resource "aws_volume_attachment" "test" {
device_name = "/dev/sdh"
volume_idws_ebs_volume.test.id
instance_id = aws_instance.test.id
func
}


func testAccEBSVolumeAttachmentConfig_stopInstance(rName string) string {
	return acctest.ConfigCompose(testAccVolumeAttachmentInstanceOnlyBaseConfig(rName), fmt.Sprintf(`
resource "aws_ebs_volume" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
size

tags = {
me = %[1]q
}
}

resource "aws_volume_attachment" "test" {
device_nameh"
volume_ids_volume.test.id
instance_idance.test.id
func
`, rName))
}


func testAccEBSVolumeAttachmentConfig_skipDestroy(rName string) string {
	return acctest.ConfigCompose(testAccEBSVolumeAttachmentConfig_base(rName), fmt.Sprintf(`
data "aws_ebs_volume" "test" {
filter {
me= "e"
lues = [aws_ebs_volume.test.size]
}

filter {
func = [aws_ebs_volume.test.availability_zone]
}

filter {
me= ":Name"
lues = ["%[1]s"]
}
}

resource "aws_volume_attachment" "test" {
device_name= "/dev/sdh"
funcstance_id= aws_instance.test.id
skip_destroy = true
}
`, rName))
}


func testAccEBSVolumeAttachmentConfig_update(rName string, detach bool) string {
	return acctest.ConfigCompose(testAccEBSVolumeAttachmentConfig_base(rName), fmt.Sprintf(`
resource "aws_volume_attachment" "test" {
device_name= "/dev/sdh"
volume_idaws_ebs_volume.test.id
instance_id= aws_instance.test.id
force_detach = %[1]t
skip_destroy = %[1]t
}
`, detach))
}


func testAccVolumeAttachmentImportStateID
func {
	return 
func(s *terraform.State) (string, error) {
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return "", fmt.Errorf("Not found: %s", resourceName)
}
return fmt.Sprintf("%s:%s:%s", rs.Primary.Attributes["device_name"], rs.Primary.Attributes["volume_id"], rs.Primary.Attributes["instance_id"]), nil
	}
}
funcfuncfuncfuncfunc