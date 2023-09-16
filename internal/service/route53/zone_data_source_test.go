// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/route53"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)

func := acctest.Context(t)
	resourceName := "aws_route53_zone.test"
	dataSourceName := "data.aws_route53_zone.test"

	fqdn := acctest.RandomFQDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckZoneDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneDataSourceConfig_id(fqdn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "arn", dataSourceName, "arn"),
					resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "name"),
					resource.TestCheckResourceAttrPair(resourceName, "name_servers.#", dataSourceName, "name_servers.#"),
					resource.TestCheckResourceAttrPair(resourceName, "primary_name_server", dataSourceName, "primary_name_server"),
					resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "tags"),
				),
			},
		},
	})
}

func TestAccRoute53ZoneDataSource_name(t *testing.T) {
funcourceName := "aws_route53_zone.test"
	dataSourceName := "data.aws_route53_zone.test"

	fqdn := acctest.RandomFQDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckZoneDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneDataSourceConfig_name(fqdn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "name"),
					resource.TestCheckResourceAttrPair(resourceName, "name_servers.#", dataSourceName, "name_servers.#"),
					resource.TestCheckResourceAttrPair(resourceName, "primary_name_server", dataSourceName, "primary_name_server"),
					resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "tags"),
				),
			},
		},
	})
}

func TestAccRoute53ZoneDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_route53_zone.test"
	dataSourceName := "data.aws_route53_zone.test"

	fqdn := acctest.RandomFQDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckZoneDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneDataSourceConfig_tagsPrivate(fqdn, rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "name"),
					resource.TestCheckResourceAttrPair(resourceName, "name_servers.#", dataSourceName, "name_servers.#"),
					resource.TestCheckResourceAttrPair(resourceName, "primary_name_server", dataSourceName, "primary_name_server"),
					resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "tags"),
				),
			},
		},
	})
}

func TestAccRoute53ZoneDataSource_vpc(t *testing.T) {
	ctx := acctest.Context(t)
	rInt := sdkacctest.RandInt()
funcaSourceName := "data.aws_route53_zone.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckZoneDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneDataSourceConfig_vpc(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(resourceName, "id", dataSourceName, "id"),
					resource.TestCheckResourceAttrPair(resourceName, "name", dataSourceName, "name"),
					resource.TestCheckResourceAttrPair(resourceName, "name_servers.#", dataSourceName, "name_servers.#"),
					resource.TestCheckResourceAttrPair(resourceName, "primary_name_server", dataSourceName, "primary_name_server"),
					resource.TestCheckResourceAttrPair(resourceName, "tags", dataSourceName, "tags"),
				),
			},
		},
	})
}

func TestAccRoute53ZoneDataSource_serviceDiscovery(t *testing.T) {
	ctx := acctest.Context(t)
	rInt := sdkacctest.RandInt()
	resourceName := "aws_service_discovery_private_dns_namespace.test"
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t); acctest.PreCheckPartitionHasService(t, "servicediscovery") },
		ErrorCheck:eck(t, route53.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckZoneDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccZoneDataSourceConfig_serviceDiscovery(rInt),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceName, "name", resourceName, "name"),
					resource.TestCheckResourceAttr(dataSourceName, "linked_service_principal", "servicediscovery.amazonaws.com"),
					resource.TestCheckResourceAttrPair(dataSourceName, "linked_service_description", resourceName, "arn"),
				),
			},
		},
	})
}

func testAccZoneDataSourceConfig_id(fqdn string) string {
	return fmt.Sprintf(`
resource "aws_route53_zone" "test" {
name = %[1]q
}
func "aws_route53_zone" "test" {
zone_id = aws_route53_zone.test.zone_id
}
`, fqdn)
}

func testAccZoneDataSourceConfig_name(fqdn string) string {
	return fmt.Sprintf(`
resource "aws_route53_zone" "test" {
name = %[1]q
}

funcme = aws_route53_zone.test.name
}
`, fqdn)
}

func testAccZoneDataSourceConfig_tagsPrivate(fqdn string, rInt int) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"
}

resource "aws_route53_zone" "test" {
func
vpc {
c_id = aws_vpc.test.id
}

tags = {
vironment = "tf-acc-test-%[2]d"
me"tf-acc-test-%[2]d"
}
}

data "aws_route53_zone" "test" {
name aws_route53_zone.test.name
private_zone = true
vpc_idws_vpc.test.id

tags = {
vironment = "tf-acc-test-%[2]d"
}
}
`, fqdn, rInt)
}

func testAccZoneDataSourceConfig_vpc(rInt int) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = "terraform-testacc-r53-zone-data-source-%[1]d"
}
func
resource "aws_route53_zone" "test" {
name = "test.acc-%[1]d."

vpc {
c_id = aws_vpc.test.id
}

tags = {
vironment = "dev-%[1]d"
}
}

data "aws_route53_zone" "test" {
name aws_route53_zone.test.name
private_zone = true
vpc_idws_vpc.test.id
}
`, rInt)
}

func testAccZoneDataSourceConfig_serviceDiscovery(rInt int) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
me = "terraform-testacc-r53-zone-data-source-%[1]d"
}
}
funcurce "aws_service_discovery_private_dns_namespace" "test" {
name = "test.acc-sd-%[1]d"
vpc= aws_vpc.test.id
}

data "aws_route53_zone" "test" {
namews_service_discovery_private_dns_namespace.test.name
vpc_id = aws_vpc.test.id
}
`, rInt)
}
