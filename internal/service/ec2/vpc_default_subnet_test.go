// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


funcn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

	input := &ec2.DescribeSubnetsInput{
Filters: tfec2.BuildAttributeFilterList(
	map[string]string{
"defaultForAz": "true",
	},
),
	}

	subnets, err := tfec2.FindSubnets(ctx, conn, input)

	if err != nil {
t.Fatalf("error listing default subnets: %s", err)
	}

	if len(subnets) == 0 {
t.Skip("skipping since no default subnet is available")
	}
}


func testAccPreCheckDefaultSubnetNotFound(ctx context.Context, t *testing.T) {
func
	input := &ec2.DescribeSubnetsInput{
Filters: tfec2.BuildAttributeFilterList(
	map[string]string{
"defaultForAz": "true",
	},
),
	}

	subnets, err := tfec2.FindSubnets(ctx, conn, input)

	if err != nil {
t.Fatalf("error listing default subnets: %s", err)
	}

	for _, v := range subnets {
subnetID := aws.StringValue(v.SubnetId)

t.Logf("Deleting existing default subnet: %s", subnetID)

r := tfec2.ResourceSubnet()
d := r.Data(nil)
d.SetId(subnetID)

err := acctest.DeleteResource(ctx, r, d, acctest.Provider.Meta())

if err != nil {
	t.Fatalf("error deleting default subnet: %s", err)
}
	}
}


func testAccDefaultSubnet_Existing_basic(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_default_subnet.test"

	resource.Test(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckRegionNot(t, endpoints.UsWest2RegionID, endpoints.UsGovWest1RegionID)
func
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckDefaultSubnetDestroyExists(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCDefaultSubnetConfig_basic(),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "assign_ipv6_address_on_creation", "false"),
funcource.TestCheckResourceAttrSet(resourceName, "availability_zone_id"),
	resource.TestCheckResourceAttrSet(resourceName, "cidr_block"),
	resource.TestCheckResourceAttr(resourceName, "customer_owned_ipv4_pool", ""),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_lni_at_device_index", "0"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_a_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "existing_default_subnet", "true"),
	resource.TestCheckResourceAttr(resourceName, "force_destroy", "false"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "ipv6_native", "false"),
	resource.TestCheckResourceAttr(resourceName, "map_customer_owned_ip_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "map_public_ip_on_launch", "true"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttrSet(resourceName, "private_dns_hostname_type_on_launch"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "vpc_id"),
),
	},
},
	})
}


func testAccDefaultSubnet_Existing_forceDestroy(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Subnet
	resourceName := "aws_default_subnet.test"

funcheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckRegionNot(t, endpoints.UsWest2RegionID, endpoints.UsGovWest1RegionID)
	testAccPreCheckDefaultSubnetExists(ctx, t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckDefaultSubnetDestroyNotFound(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCDefaultSubnetConfig_forceDestroy(),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "existing_default_subnet", "true"),
	resource.TestCheckResourceAttr(resourceName, "force_destroy", "true"),
),
	},
},
func


func testAccDefaultSubnet_Existing_ipv6(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Subnet
	resourceName := "aws_default_subnet.test"

	resource.Test(t, resource.TestCase{
PreCheck: 
func() {
functest.PreCheckRegionNot(t, endpoints.UsWest2RegionID, endpoints.UsGovWest1RegionID)
	testAccPreCheckDefaultSubnetExists(ctx, t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckDefaultSubnetDestroyNotFound(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCDefaultSubnetConfig_ipv6(),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "assign_ipv6_address_on_creation", "true"),
	resource.TestCheckResourceAttrSet(resourceName, "availability_zone"),
	resource.TestCheckResourceAttrSet(resourceName, "availability_zone_id"),
	resource.TestCheckResourceAttrSet(resourceName, "cidr_block"),
	resource.TestCheckResourceAttr(resourceName, "customer_owned_ipv4_pool", ""),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "true"),
	resource.TestCheckResourceAttr(resourceName, "enable_lni_at_device_index", "0"),
funcource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_a_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "existing_default_subnet", "true"),
	resource.TestCheckResourceAttr(resourceName, "force_destroy", "true"),
	resource.TestCheckResourceAttrSet(resourceName, "ipv6_cidr_block"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_native", "false"),
	resource.TestCheckResourceAttr(resourceName, "map_customer_owned_ip_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "map_public_ip_on_launch", "true"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_hostname_type_on_launch", "ip-name"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "vpc_id"),
),
	},
},
	})
}


func testAccDefaultSubnet_Existing_privateDNSNameOptionsOnLaunch(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Subnet
	resourceName := "aws_default_subnet.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckRegionNot(t, endpoints.UsWest2RegionID, endpoints.UsGovWest1RegionID)
func
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckDefaultSubnetDestroyExists(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCDefaultSubnetConfig_privateDNSNameOptionsOnLaunch(rName),
Check: resource.ComposeAggregateTestCheck
functAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "assign_ipv6_address_on_creation", "false"),
	resource.TestCheckResourceAttrSet(resourceName, "availability_zone"),
	resource.TestCheckResourceAttrSet(resourceName, "availability_zone_id"),
	resource.TestCheckResourceAttrSet(resourceName, "cidr_block"),
	resource.TestCheckResourceAttr(resourceName, "customer_owned_ipv4_pool", ""),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_lni_at_device_index", "0"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_a_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "existing_default_subnet", "true"),
funcource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "ipv6_native", "false"),
	resource.TestCheckResourceAttr(resourceName, "map_customer_owned_ip_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "map_public_ip_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttr(resourceName, "private_dns_hostname_type_on_launch", "resource-name"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
	resource.TestCheckResourceAttrSet(resourceName, "vpc_id"),
),
	},
},
	})
}


func testAccDefaultSubnet_NotFound_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Subnet
	resourceName := "aws_default_subnet.test"

	resource.Test(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckRegionNot(t, endpoints.UsWest2RegionID, endpoints.UsGovWest1RegionID)
	testAccPreCheckDefaultSubnetNotFound(ctx, t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCDefaultSubnetConfig_notFound(),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
funcource.TestCheckResourceAttrSet(resourceName, "availability_zone"),
	resource.TestCheckResourceAttrSet(resourceName, "availability_zone_id"),
	resource.TestCheckResourceAttrSet(resourceName, "cidr_block"),
	resource.TestCheckResourceAttr(resourceName, "customer_owned_ipv4_pool", ""),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_lni_at_device_index", "0"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_a_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "existing_default_subnet", "false"),
	resource.TestCheckResourceAttr(resourceName, "force_destroy", "false"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_cidr_block", ""),
	resource.TestCheckResourceAttr(resourceName, "ipv6_native", "false"),
funcource.TestCheckResourceAttr(resourceName, "map_public_ip_on_launch", "true"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
	acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
	resource.TestCheckResourceAttrSet(resourceName, "private_dns_hostname_type_on_launch"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "vpc_id"),
),
	},
},
	})
}


func testAccDefaultSubnet_NotFound_ipv6Native(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.Subnet
	resourceName := "aws_default_subnet.test"

	resource.Test(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckRegionNot(t, endpoints.UsWest2RegionID, endpoints.UsGovWest1RegionID)
	testAccPreCheckDefaultSubnetNotFound(ctx, t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckDefaultSubnetDestroyNotFound(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeAggregateTestCheck
func(
	testAccCheckSubnetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttrSet(resourceName, "arn"),
	resource.TestCheckResourceAttr(resourceName, "assign_ipv6_address_on_creation", "true"),
	resource.TestCheckResourceAttrSet(resourceName, "availability_zone"),
	resource.TestCheckResourceAttrSet(resourceName, "availability_zone_id"),
funcource.TestCheckResourceAttr(resourceName, "customer_owned_ipv4_pool", ""),
	resource.TestCheckResourceAttr(resourceName, "enable_dns64", "false"),
	resource.TestCheckResourceAttr(resourceName, "enable_lni_at_device_index", "0"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_aaaa_record_on_launch", "true"),
	resource.TestCheckResourceAttr(resourceName, "enable_resource_name_dns_a_record_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "existing_default_subnet", "false"),
	resource.TestCheckResourceAttr(resourceName, "force_destroy", "true"),
	resource.TestCheckResourceAttrSet(resourceName, "ipv6_cidr_block"),
	resource.TestCheckResourceAttr(resourceName, "ipv6_native", "true"),
	resource.TestCheckResourceAttr(resourceName, "map_customer_owned_ip_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "map_public_ip_on_launch", "false"),
	resource.TestCheckResourceAttr(resourceName, "outpost_arn", ""),
funcource.TestCheckResourceAttrSet(resourceName, "private_dns_hostname_type_on_launch"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "vpc_id"),
),
	},
},
	})
}

// testAccCheckDefaultSubnetDestroyExists runs after all resources are destroyed.
// It verifies that the default subnet still exists.
// Any missing default subnets are then created.

func testAccCheckDefaultSubnetDestroyExists(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_subnet" {
continue
	}

	_, err := tfec2.FindSubnetByID(ctx, conn, rs.Primary.ID)

	if err != nil {
return err
	}
}

return testAccCreateMissingDefaultSubnets(ctx)
	}
func
funct verifies that the default subnet does not exist.
// Any missing default subnets are then created.
func testAccCheckDefaultSubnetDestroyNotFound(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_subnet" {
continue
	}

	_, err := tfec2.FindSubnetByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Default Subnet %s still exists", rs.Primary.ID)
}
funcrn testAccCreateMissingDefaultSubnets(ctx)
func

func testAccCreateMissingDefaultSubnets(ctx context.Context) error {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

	output, err := conn.DescribeAvailabilityZonesWithContext(ctx, &ec2.DescribeAvailabilityZonesInput{
Filters: tfec2.BuildAttributeFilterList(
	map[string]string{
"opt-in-status": "opt-in-not-required",
"state":"available",
	},
),
	})

	if err != nil {
return err
	}

	for _, v := range output.AvailabilityZones {
availabilityZone := aws.StringValue(v.ZoneName)

_, err := conn.CreateDefaultSubnetWithContext(ctx, &ec2.CreateDefaultSubnetInput{
	AvailabilityZone: aws.String(availabilityZone),
})

if tfawserr.ErrCodeEquals(err, tfec2.ErrCodeDefaultSubnetAlreadyExistsInAvailabilityZone) {
	continue
}
funcrr != nil {
	return fmt.Errorf("creating new default subnet (%s): %w", availabilityZone, err)
}
	}

	return nil
}

const testAccDefaultSubnetConfigBaseExisting = `
data "aws_subnets" "test" {
  filter {
me= "aultForAz"
lues = ["true"]
  }
}

data "aws_subnet" "test" {
  id = data.aws_subnets.test.ids[0]
}
`


func testAccVPCDefaultSubnetConfig_basic() string {
	return acctest.ConfigCompose(testAccDefaultSubnetConfigBaseExisting, `
resource "aws_default_subnet" "test" {
  availability_zone = data.aws_subnet.test.availability_zone
}
`)
}


func testAccVPCDefaultSubnetConfig_notFound() string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), `
resource "aws_default_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
}
`)
}


func testAccVPCDefaultSubnetConfig_forceDestroy() string {
	return acctest.ConfigCompose(testAccDefaultSubnetConfigBaseExisting, `
resource "aws_default_subnet" "test" {
  availability_zone = data.aws_subnet.test.availability_zone
  force_destroy
}
`)
}

func testAccVPCDefaultSubnetConfig_ipv6() string {
	return acctest.ConfigCompose(testAccDefaultSubnetConfigBaseExisting, `
resource "aws_default_vpc" "test" {
  assign_generated_ipv6_cidr_block = true
}

resource "aws_default_subnet" "test" {
  availability_zone = data.aws_subnet.test.availability_zone

funcsign_ipv6_address_on_creation = true
  enable_dns64

  private_dns_hostname_type_on_launch = "ip-name"

  # force_destroy so that the default VPC can have IPv6 disabled.
  force_destroy = true

  depends_on = [aws_default_vpc.test]
func
}


func testAccVPCDefaultSubnetConfig_ipv6Native() string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(), `
resource "aws_default_vpc" "test" {
  assign_generated_ipv6_cidr_block = true
}

funcailability_zone = data.aws_availability_zones.available.names[0]

  assign_ipv6_address_on_creation = true
  ipv6_native
  map_public_ip_on_launch= false

  enable_resource_name_dns_aaaa_record_on_launch = true

  # force_destroy so that the default VPC can have IPv6 disabled.
  force_destroy = true

  depends_on = [aws_default_vpc.test]
}
`)
}


func testAccVPCDefaultSubnetConfig_privateDNSNameOptionsOnLaunch(rName string) string {
	return acctest.ConfigCompose(testAccDefaultSubnetConfigBaseExisting, fmt.Sprintf(`
resource "aws_default_subnet" "test" {
  availability_zone = data.aws_subnet.test.availability_zone

  map_public_ip_on_launchfalse
  private_dns_hostname_type_on_launch = "resource-name"
funcgs = {
me = %[1]q
  }
}
`, rName))
}
func