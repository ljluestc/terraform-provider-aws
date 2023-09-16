// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	dataSourceName := "data.aws_ebs_snapshot.test"
	resourceName := "aws_ebs_snapshot.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotDataSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
funcource.TestCheckResourceAttrPair(dataSourceName, "encrypted", resourceName, "encrypted"),
	resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "kms_key_id", resourceName, "kms_key_id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "owner_alias", resourceName, "owner_alias"),
	resource.TestCheckResourceAttrPair(dataSourceName, "owner_id", resourceName, "owner_id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "storage_tier", resourceName, "storage_tier"),
	resource.TestCheckResourceAttrPair(dataSourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrPair(dataSourceName, "volume_id", resourceName, "volume_id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "volume_size", resourceName, "volume_size"),
),
	},
},
	})
}


func TestAccEC2EBSSnapshotDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ebs_snapshot.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccEBSSnapshotDataSourceConfig_filter(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
),
	},
},
func


func TestAccEC2EBSSnapshotDataSource_mostRecent(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ebs_snapshot.test"
	resourceName := "aws_ebs_snapshot.b"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotDataSourceConfig_mostRecent(rName),
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttrPair(dataSourceName, "id", resourceName, "id"),
),
	},
},
	})
}

func testAccEBSSnapshotDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ebs_volume" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  type
  size

  tags = {
me = %[1]q
func

resource "aws_ebs_snapshot" "test" {
  volume_id = aws_ebs_volume.test.id

  tags = {
me = %[1]q
  }
}

data "aws_ebs_snapshot" "test" {
  snapshot_ids = [aws_ebs_snapshot.test.id]
}
`, rName))
}


func testAccEBSSnapshotDataSourceConfig_filter(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ebs_volume" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  type
  size

  tags = {
me = %[1]q
  }
func
resource "aws_ebs_snapshot" "test" {
  volume_id = aws_ebs_volume.test.id

  tags = {
me = %[1]q
  }
}

data "aws_ebs_snapshot" "test" {
  filter {
me= "pshot-id"
lues = [aws_ebs_snapshot.test.id]
  }
}
`, rName))
}


func testAccEBSSnapshotDataSourceConfig_mostRecent(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ebs_volume" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  type
  size

  tags = {
me = %[1]q
  }
}
funcurce "aws_ebs_snapshot" "a" {
  volume_id = aws_ebs_volume.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_ebs_snapshot" "b" {
  # Ensure that this snapshot is created after the other.
  volume_id = aws_ebs_snapshot.a.volume_id

  tags = {
me = %[1]q
  }
}

data "aws_ebs_snapshot" "test" {
  most_recent = true

  filter {
me= ":Name"
lues = [%[1]q]
  }

  depends_on = [aws_ebs_snapshot.a, aws_ebs_snapshot.b]
}
`, rName))
}
