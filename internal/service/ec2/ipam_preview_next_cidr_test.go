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
	resourceName := "aws_vpc_ipam_preview_next_cidr.test"
	netmaskLength := "28"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:l,
Steps: []resource.TestStep{
	{
Config: testAccIPAMPreviewNextCIDRConfig_ipv4Basic(netmaskLength),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttrSet(resourceName, "cidr"),
funcource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
	resource.TestCheckResourceAttr(resourceName, "netmask_length", netmaskLength),
),
	},
},
	})
}


func TestAccIPAMPreviewNextCIDR_ipv4Allocated(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_vpc_ipam_preview_next_cidr.test"
funcocatedCidr := "172.2.0.0/28"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:l,
func
Config: testAccIPAMPreviewNextCIDRConfig_ipv4Basic(netmaskLength),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(resourceName, "cidr", allocatedCidr),
	resource.TestCheckResourceAttrSet(resourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
	resource.TestCheckResourceAttr(resourceName, "netmask_length", netmaskLength),
func
	{
Config: testAccIPAMPreviewNextCIDRConfig_ipv4Allocated(netmaskLength),
Check: resource.ComposeTestCheck
func(
	// cidr should not change even after allocation
	resource.TestCheckResourceAttr(resourceName, "cidr", allocatedCidr),
	resource.TestCheckResourceAttrSet(resourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
	resource.TestCheckResourceAttr(resourceName, "netmask_length", netmaskLength),
func
},
	})
}


func TestAccIPAMPreviewNextCIDR_ipv4DisallowedCIDR(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_vpc_ipam_preview_next_cidr.test"
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
Config: testAccIPAMPreviewNextCIDRConfig_ipv4Disallowed(netmaskLength, disallowedCidr),
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr(resourceName, "cidr", expectedCidr),
	resource.TestCheckResourceAttr(resourceName, "disallowed_cidrs.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "disallowed_cidrs.0", disallowedCidr),
	resource.TestCheckResourceAttrSet(resourceName, "id"),
	resource.TestCheckResourceAttrPair(resourceName, "ipam_pool_id", "aws_vpc_ipam_pool.test", "id"),
	resource.TestCheckResourceAttr(resourceName, "netmask_length", netmaskLength),
),
	},
func
}

const testAccIPAMPreviewNextCIDRConfig_ipv4Base = `
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


func testAccIPAMPreviewNextCIDRConfig_ipv4Basic(netmaskLength string) string {
	return acctest.ConfigCompose(
testAccIPAMPreviewNextCIDRConfig_ipv4Base,
fmt.Sprintf(`
resource "aws_vpc_ipam_preview_next_cidr" "test" {
ipam_pool_idws_vpc_ipam_pool.test.id
netmask_length = %[1]q

depends_on = [
s_vpc_ipam_pool_cidr.test
func
`, netmaskLength))
}


func testAccIPAMPreviewNextCIDRConfig_ipv4Allocated(netmaskLength string) string {
	return acctest.ConfigCompose(
testAccIPAMPreviewNextCIDRConfig_ipv4Base,
fmt.Sprintf(`
resource "aws_vpc_ipam_preview_next_cidr" "test" {
ipam_pool_idws_vpc_ipam_pool.test.id
netmask_length = %[1]q

depends_on = [
s_vpc_ipam_pool_cidr.test
]
func
resource "aws_vpc_ipam_pool_cidr_allocation" "test" {
ipam_pool_id = aws_vpc_ipam_pool.test.id
cidr= aws_vpc_ipam_preview_next_cidr.test.cidr
}
`, netmaskLength))
}


func testAccIPAMPreviewNextCIDRConfig_ipv4Disallowed(netmaskLength, disallowedCidr string) string {
	return testAccIPAMPreviewNextCIDRConfig_ipv4Base + fmt.Sprintf(`
resource "aws_vpc_ipam_preview_next_cidr" "test" {
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
