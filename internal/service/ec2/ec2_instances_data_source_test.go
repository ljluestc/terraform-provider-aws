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
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccInstancesDataSourceConfig_ids(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_instances.test", "ids.#", "2"),
funcource.TestCheckResourceAttr("data.aws_instances.test", "private_ips.#", "2"),
	// Public IP values are flakey for new EC2 instances due to eventual consistency
	resource.TestCheckResourceAttrSet("data.aws_instances.test", "public_ips.#"),
),
	},
},
	})
}


func TestAccEC2InstancesDataSource_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccInstancesDataSourceConfig_tags(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_instances.test", "ids.#", "2"),
),
	},
},
func


func TestAccEC2InstancesDataSource_instanceStateNames(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccInstancesDataSourceConfig_instanceStateNames(rName),
Check: resource.ComposeTestCheck
funcource.TestCheckResourceAttr("data.aws_instances.test", "ids.#", "2"),
),
	},
},
	})
}

func TestAccEC2InstancesDataSource_empty(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
func
Config: testAccInstancesDataSourceConfig_empty(rName),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr("data.aws_instances.test", "ids.#", "0"),
	resource.TestCheckResourceAttr("data.aws_instances.test", "ipv6_addresses.#", "0"),
funcource.TestCheckResourceAttr("data.aws_instances.test", "public_ips.#", "0"),
),
	},
},
	})
}

func TestAccEC2InstancesDataSource_timeout(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccInstancesDataSourceConfig_timeout(rName),
func(
	resource.TestCheckResourceAttr("data.aws_instances.test", "ids.#", "2"),
	resource.TestCheckResourceAttrSet("data.aws_instances.test", "ipv6_addresses.#"),
	resource.TestCheckResourceAttr("data.aws_instances.test", "private_ips.#", "2"),
	resource.TestCheckResourceAttrSet("data.aws_instances.test", "public_ips.#"),
),
func
	})
}


func testAccInstancesDataSourceConfig_ids(rName string) string {
	return acctest.ConfigCompose(
funcest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
acctest.ConfigVPCWithSubnetsIPv6(rName, 1),
fmt.Sprintf(`
resource "aws_instance" "test" {
  count
  ami = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
  instance_types_ec2_instance_type_offering.available.instance_type
  subnet_id = aws_subnet.test[0].id
  ipv6_address_count = 1

  tags = {
me = %[1]q
func

data "aws_instances" "test" {
  filter {
me= "tance-id"
lues = aws_instance.test[*].id
  }
}
`, rName))
}


func testAccInstancesDataSourceConfig_tags(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_instance" "test" {
  count= 2
  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type

  tags = {
me
condTag = "%[1]s-2"
  }
}

funcstance_tags = {
metance.test[0].tags["Name"]
condTag = aws_instance.test[0].tags["SecondTag"]
  }

  depends_on = [aws_instance.test]
}
`, rName))
}


func testAccInstancesDataSourceConfig_instanceStateNames(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_instance" "test" {
  count= 2
  ami  = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type

  tags = {
me = %[1]q
  }
}

data "aws_instances" "test" {
  instance_tags = {
func

  instance_state_names = ["pending", "running"]
  depends_on  = [aws_instance.test]
}
`, rName))
}


func testAccInstancesDataSourceConfig_empty(rName string) string {
	return fmt.Sprintf(`
data "aws_instances" "test" {
  instance_tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccInstancesDataSourceConfig_timeout(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_instance" "test" {
  count= 2
funcstance_type = data.aws_ec2_instance_type_offering.available.instance_type

  tags = {
me = %[1]q
  }
}

data "aws_instances" "test" {
  filter {
me= "tance-id"
lues = aws_instance.test[*].id
func
  timeouts {
ad = "60m"
  }
}
`, rName))
}
