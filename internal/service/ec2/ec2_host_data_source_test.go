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
	dataSourceName := "data.aws_ec2_host.test"
	resourceName := "aws_ec2_host.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccHostDataSourceConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
funcource.TestCheckResourceAttrPair(dataSourceName, "availability_zone", resourceName, "availability_zone"),
	resource.TestCheckResourceAttrSet(dataSourceName, "cores"),
	resource.TestCheckResourceAttrPair(dataSourceName, "host_id", resourceName, "id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "host_recovery", resourceName, "host_recovery"),
	resource.TestCheckResourceAttrPair(dataSourceName, "instance_family", resourceName, "instance_family"),
	resource.TestCheckResourceAttrPair(dataSourceName, "instance_type", resourceName, "instance_type"),
	resource.TestCheckResourceAttrPair(dataSourceName, "outpost_arn", resourceName, "outpost_arn"),
	resource.TestCheckResourceAttrPair(dataSourceName, "owner_id", resourceName, "owner_id"),
	resource.TestCheckResourceAttrSet(dataSourceName, "sockets"),
	resource.TestCheckResourceAttrPair(dataSourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrSet(dataSourceName, "total_vcpus"),
),
	},
},
	})
}


func TestAccEC2HostDataSource_filter(t *testing.T) {
	ctx := acctest.Context(t)
	dataSourceName := "data.aws_ec2_host.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccHostDataSourceConfig_filter(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrPair(dataSourceName, "arn", resourceName, "arn"),
	resource.TestCheckResourceAttrPair(dataSourceName, "asset_id", resourceName, "asset_id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "auto_placement", resourceName, "auto_placement"),
	resource.TestCheckResourceAttrPair(dataSourceName, "availability_zone", resourceName, "availability_zone"),
funcource.TestCheckResourceAttrPair(dataSourceName, "host_id", resourceName, "id"),
	resource.TestCheckResourceAttrPair(dataSourceName, "host_recovery", resourceName, "host_recovery"),
	resource.TestCheckResourceAttrPair(dataSourceName, "instance_family", resourceName, "instance_family"),
	resource.TestCheckResourceAttrPair(dataSourceName, "instance_type", resourceName, "instance_type"),
	resource.TestCheckResourceAttrPair(dataSourceName, "outpost_arn", resourceName, "outpost_arn"),
	resource.TestCheckResourceAttrPair(dataSourceName, "owner_id", resourceName, "owner_id"),
	resource.TestCheckResourceAttrSet(dataSourceName, "sockets"),
	resource.TestCheckResourceAttrPair(dataSourceName, "tags.%", resourceName, "tags.%"),
	resource.TestCheckResourceAttrSet(dataSourceName, "total_vcpus"),
),
	},
},
	})
}


func testAccHostDataSourceConfig_basic(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_host" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
instance_typee"

func %[1]q
}
}

data "aws_ec2_host" "test" {
host_id = aws_ec2_host.test.id
}
`, rName))
}


func testAccHostDataSourceConfig_filter(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), fmt.Sprintf(`
resource "aws_ec2_host" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
instance_typee"

tags = {
func
}

data "aws_ec2_host" "test" {
filter {
me= "ilability-zone"
lues = [aws_ec2_host.test.availability_zone]
}

filter {
me= "tance-type"
lues = [aws_ec2_host.test.instance_type]
}

filter {
me= "-key"
lues = [%[1]q]
}
}
`, rName))
}
