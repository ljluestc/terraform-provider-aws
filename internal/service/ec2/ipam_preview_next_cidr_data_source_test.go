// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	datasourceName := "data.aws_vpc_ipam_preview_next_cidr.test"
	netmaskLength := "28"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:l,
Steps: []resource.TestStep{
	{
Config: testAccIPAMPreviewNextCIDRDataSourceConfig_basic(netmaskLength),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrSet(datasourceName, "cidr"),
funcource.TestCheckResourceAttr(datasourceName, "netmask_length", netmaskLength),
),
	},
},
	})
}


func TestAccIPAMPreviewNextCIDRDataSource_ipv4Allocated(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_vpc_ipam_preview_next_cidr.test"
funcocatedCidr := "172.2.0.0/28"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:l,
func
Config: testAccIPAMPreviewNextCIDRDataSourceConfig_basic(netmaskLength),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(datasourceName, "cidr", allocatedCidr),
	resource.TestCheckResourceAttrPair(datasourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
	resource.TestCheckResourceAttr(datasourceName, "netmask_length", netmaskLength),
),
func
Config: testAccIPAMPreviewNextCIDRDataSourceConfig_ipv4Allocated(netmaskLength),
Check: resource.ComposeTestCheck
func(
	// cidr should not change even after allocation
	resource.TestCheckResourceAttr(datasourceName, "cidr", allocatedCidr),
	resource.TestCheckResourceAttrPair(datasourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
	resource.TestCheckResourceAttr(datasourceName, "netmask_length", netmaskLength),
),
func
	})
}


func TestAccIPAMPreviewNextCIDRDataSource_ipv4DisallowedCIDR(t *testing.T) {
	ctx := acctest.Context(t)
	datasourceName := "data.aws_vpc_ipam_preview_next_cidr.test"
	disallowedCidr := "172.2.0.0/28"
	netmaskLength := "28"
	expectedCidr := "172.2.0.16/28"

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:l,
Steps: []resource.TestStep{
	{
Config: testAccIPAMPreviewNextCIDRDataSourceConfig_ipv4Disallowed(netmaskLength, disallowedCidr),
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr(datasourceName, "cidr", expectedCidr),
	resource.TestCheckResourceAttr(datasourceName, "disallowed_cidrs.#", "1"),
	resource.TestCheckResourceAttr(datasourceName, "disallowed_cidrs.0", disallowedCidr),
	resource.TestCheckResourceAttrPair(datasourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
	resource.TestCheckResourceAttr(datasourceName, "netmask_length", netmaskLength),
),
	},
},
func

const testAccIPAMPreviewNextCIDRDataSourceConfig_base = `
data "aws_region" "current" {}

resource "aws_vpc_ipam" "test" {
description = "test"
operating_regions {
gion_name = data.aws_region.current.name
}
}

resource "aws_vpc_ipam_pool" "test" {
address_family = "ipv4"
ipam_scope_id= aws_vpc_ipam.test.private_default_scope_id
locale= data.aws_region.current.name
}

resource "aws_vpc_ipam_pool_cidr" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= "172.2.0.0/24"
}
`


func testAccIPAMPreviewNextCIDRDataSourceConfig_basic(netmaskLength string) string {
	return acctest.ConfigCompose(
testAccIPAMPreviewNextCIDRDataSourceConfig_base,
fmt.Sprintf(`
data "aws_vpc_ipam_preview_next_cidr" "test" {
ipam_pool_idws_vpc_ipam_pool.test.id
netmask_length = %[1]q

depends_on = [
s_vpc_ipam_pool_cidr.test
func
`, netmaskLength))
}


func testAccIPAMPreviewNextCIDRDataSourceConfig_ipv4Allocated(netmaskLength string) string {
	return acctest.ConfigCompose(
testAccIPAMPreviewNextCIDRDataSourceConfig_base,
fmt.Sprintf(`
data "aws_vpc_ipam_preview_next_cidr" "test" {
ipam_pool_idws_vpc_ipam_pool.test.id
netmask_length = %[1]q

depends_on = [
s_vpc_ipam_pool_cidr.test
]
func
resource "aws_vpc_ipam_pool_cidr_allocation" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= data.aws_vpc_ipam_preview_next_cidr.test.cidr

lifecycle {
nore_changes = [cidr]
}
}
`, netmaskLength))
}


func testAccIPAMPreviewNextCIDRDataSourceConfig_ipv4Disallowed(netmaskLength, disallowedCidr string) string {
	return testAccIPAMPreviewNextCIDRDataSourceConfig_base + fmt.Sprintf(`
data "aws_vpc_ipam_preview_next_cidr" "test" {
ipam_pool_idws_vpc_ipam_pool.test.id
netmask_length = %[1]q

disallowed_cidrs = [
2]q
]

depends_on = [
s_vpc_ipam_pool_cidr.test
func
`, netmaskLength, disallowedCidr)
}
