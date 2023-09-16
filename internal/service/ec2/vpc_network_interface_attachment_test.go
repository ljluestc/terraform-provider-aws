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
	var conf ec2.NetworkInterface
	resourceName := "aws_network_interface_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckENIDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInterfaceAttachmentConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckENIExists(ctx, "aws_network_interface.test", &conf),
funcource.TestCheckResourceAttr(resourceName, "device_index", "1"),
	resource.TestCheckResourceAttrSet(resourceName, "instance_id"),
	resource.TestCheckResourceAttrSet(resourceName, "network_interface_id"),
	resource.TestCheckResourceAttrSet(resourceName, "status"),
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


func testAccVPCNetworkInterfaceAttachmentConfig_basic(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
funcest.ConfigAvailableAZsNoOptIn(),
fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "172.16.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
vpc_idws_vpc.test.id
cidr_block16.10.0/24"
availability_zone = data.aws_availability_zones.available.names[0]

tags = {
me = %[1]q
}
}

resource "aws_security_group" "test" {
vpc_id = aws_vpc.test.id
name[1]q

egress {
om_port= 0
_port
otocol = p"
dr_blocks = ["10.0.0.0/16"]
}
}

resource "aws_network_interface" "test" {
subnet_idbnet.test.id
private_ips.10.100"]
security_groups = [aws_security_group.test.id]

tags = {
me = %[1]q
}
}

resource "aws_instance" "test" {
ami= data.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_type = data.aws_ec2_instance_type_offering.available.instance_type
subnet_idet.test.id

tags = {
me = %[1]q
}
}

resource "aws_network_interface_attachment" "test" {
device_index= 1
instance_id = aws_instance.test.id
network_interface_id = aws_network_interface.test.id
}
`, rName))
}
