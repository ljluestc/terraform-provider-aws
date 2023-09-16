// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
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
	var v ec2.Snapshot
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ebs_snapshot.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
funcource.TestCheckResourceAttr(resourceName, "description", ""),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "storage_tier", "standard"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
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


func TestAccEC2EBSSnapshot_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Snapshot
funcourceName := "aws_ebs_snapshot.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
func
Config: testAccEBSSnapshotConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceEBSSnapshot(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccEC2EBSSnapshot_storageTier(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Snapshot
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ebs_snapshot.test"

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotConfig_storageTier(rName, "archive"),
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "storage_tier", "archive"),
),
	},
	{
ResourceName:ame,
ImportState:
func
},
	})
}


func TestAccEC2EBSSnapshot_outpost(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Snapshot
	outpostDataSourceName := "data.aws_outposts_outpost.test"
	resourceName := "aws_ebs_snapshot.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotConfig_outpost(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
func


func TestAccEC2EBSSnapshot_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Snapshot
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ebs_snapshot.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
func
Config: testAccEBSSnapshotConfig_tags1(rName, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccEBSSnapshotConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
functAccCheckSnapshotExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccEBSSnapshotConfig_tags1(rName, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
func
	})
}


func TestAccEC2EBSSnapshot_withDescription(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Snapshot
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_ebs_snapshot.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotConfig_description(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "description", "test description"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
func
}


func TestAccEC2EBSSnapshot_withKMS(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Snapshot
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcourceName := "aws_ebs_snapshot.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckEBSSnapshotDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotConfig_kms(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSnapshotExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrPair(resourceName, "kms_key_id", kmsKeyResourceName, "arn"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}
func
func testAccCheckSnapshotExists(ctx context.Context, n string, v *ec2.Snapshot) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
func
if rs.Primary.ID == "" {
	return fmt.Errorf("No EBS Snapshot ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindSnapshotByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}

*v = *output

func
func

func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ebs_snapshot" {
continue
	}

	_, err := tfec2.FindSnapshotByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EBS Snapshot %s still exists", rs.Primary.ID)
}

return nil
	}
func
func testAccEBSSnapshotBaseConfig(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
funcailability_zone = data.aws_availability_zones.available.names[0]
  size

  tags = {
me = %[1]q
  }
}
`, rName))
}


func testAccEBSSnapshotConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccEBSSnapshotBaseConfig(rName), `
resource "aws_ebs_snapshot" "test" {
  volume_id = aws_ebs_volume.test.id
}
`)
}


func testAccEBSSnapshotConfig_storageTier(rName, tier string) string {
	return acctest.ConfigCompose(testAccEBSSnapshotBaseConfig(rName), fmt.Sprintf(`
resource "aws_ebs_snapshot" "test" {
  volume_idaws_ebs_volume.test.id
  storage_tier = %[2]q

func %[1]q
  }
}
`, rName, tier))
}


func testAccEBSSnapshotConfig_outpost(rName string) string {
	return acctest.ConfigCompose(testAccEBSSnapshotBaseConfig(rName), fmt.Sprintf(`
data "aws_outposts_outposts" "test" {}

data "aws_outposts_outpost" "test" {
  id = tolist(data.aws_outposts_outposts.test.ids)[0]
}
funcurce "aws_ebs_snapshot" "test" {
  volume_idws_ebs_volume.test.id
  outpost_arn = data.aws_outposts_outpost.test.arn

  tags = {
me = %[1]q
  }
}
`, rName))
func

func testAccEBSSnapshotConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccEBSSnapshotBaseConfig(rName), fmt.Sprintf(`
resource "aws_ebs_snapshot" "test" {
  volume_id = aws_ebs_volume.test.id

  tags = {
1]q = %[2]q
  }
}
`, tagKey1, tagValue1))
}

func testAccEBSSnapshotConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccEBSSnapshotBaseConfig(rName), fmt.Sprintf(`
resource "aws_ebs_snapshot" "test" {
  volume_id = aws_ebs_volume.test.id

  tags = {
1]q = %[2]q
3]q = %[4]q
  }
}
`, tagKey1, tagValue1, tagKey2, tagValue2))
}


func testAccEBSSnapshotConfig_description(rName string) string {
	return acctest.ConfigCompose(testAccEBSSnapshotBaseConfig(rName), fmt.Sprintf(`
resource "aws_ebs_snapshot" "test" {
  volume_idws_ebs_volume.test.id
  description = "test description"

func %[1]q
  }
}
`, rName))
}


func testAccEBSSnapshotConfig_kms(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_kms_key" "test" {
  deletion_window_in_days = 7

  tags = {
func
}

resource "aws_ebs_volume" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  size
  encrypted= true
  kms_key_idms_key.test.arn

  tags = {
me = %[1]q
  }
}

funclume_id = aws_ebs_volume.test.id

  tags = {
me = %[1]q
  }
}
`, rName))
}
func