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
	dataSourceName := "data.aws_ebs_snapshot_ids.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccEBSSnapshotIdsDataSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "ids.#", 0),
func
	},
},
	})
}


func TestAccEC2EBSSnapshotIDsDataSource_sorted(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ebs_snapshot_ids.test"
funcource2Name := "aws_ebs_snapshot.b"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccEBSSnapshotIdsDataSourceConfig_sorted(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(dataSourceName, "ids.#", "2"),
	resource.TestCheckResourceAttrPair(dataSourceName, "ids.0", resource2Name, "id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "ids.1", resource1Name, "id"),
),
func
	})
}


func TestAccEC2EBSSnapshotIDsDataSource_empty(t *testing.T) {
	ctx := acctest.Context(t)
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funcs: []resource.TestStep{
	{
Config: testAccEBSSnapshotIdsDataSourceConfig_empty,
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr("data.aws_ebs_snapshot_ids.empty", "ids.#", "0"),
),
	},
},
	})
}

func testAccEBSSnapshotIdsDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ebs_volume" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
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

data "aws_ebs_snapshot_ids" "test" {
owners = ["self"]

depends_on = [aws_ebs_snapshot.test]
}
`, rName))
}


func testAccEBSSnapshotIdsDataSourceConfig_sorted(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ebs_volume" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
size

count = 2

tags = {
me = %[1]q
func

resource "aws_ebs_snapshot" "a" {
volume_idws_ebs_volume.test[0].id
description = %[1]q

tags = {
me = %[1]q
}
}

resource "aws_ebs_snapshot" "b" {
volume_idws_ebs_volume.test[1].id
description = %[1]q

# We want to ensure that 'aws_ebs_snapshot.a.creation_date' is less than
# 'aws_ebs_snapshot.b.creation_date'/ so that we can ensure that the
# snapshots are being sorted correctly.
depends_on = [aws_ebs_snapshot.a]

tags = {
me = %[1]q
}
}

data "aws_ebs_snapshot_ids" "test" {
owners = ["self"]

filter {
me= "cription"
lues = [%[1]q]
}

depends_on = [aws_ebs_snapshot.a, aws_ebs_snapshot.b]
}
`, rName))
}

const testAccEBSSnapshotIdsDataSourceConfig_empty = `
data "aws_ebs_snapshot_ids" "empty" {
owners = ["000000000000"]
}
`
