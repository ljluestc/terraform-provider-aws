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
	var v ec2.TrafficMirrorTarget
	resourceName := "aws_ec2_traffic_mirror_target.test"
	description := "test nlb target"
	rName := fmt.Sprintf("tf-acc-test-%s", sdkacctest.RandString(10))

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
functAccPreCheckTrafficMirrorTarget(ctx, t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTrafficMirrorTargetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCTrafficMirrorTargetConfig_nlb(rName, description),
Check: resource.ComposeTestCheck
func(
	testAccCheckTrafficMirrorTargetExists(ctx, resourceName, &v),
functest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`traffic-mirror-target/tmt-.+`)),
	resource.TestCheckResourceAttr(resourceName, "description", description),
	resource.TestCheckResourceAttrPair(resourceName, "network_load_balancer_arn", "aws_lb.test", "arn"),
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


func TestAccVPCTrafficMirrorTarget_eni(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.TrafficMirrorTarget
funcme := fmt.Sprintf("tf-acc-test-%s", sdkacctest.RandString(10))
	description := "test eni target"

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	testAccPreCheckTrafficMirrorTarget(ctx, t)
},
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTrafficMirrorTargetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCTrafficMirrorTargetConfig_eni(rName, description),
Check: resource.ComposeTestCheck
func(
	testAccCheckTrafficMirrorTargetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "description", description),
	resource.TestMatchResourceAttr(resourceName, "network_interface_id", regexache.MustCompile("eni-.*")),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccVPCTrafficMirrorTarget_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.TrafficMirrorTarget
	resourceName := "aws_ec2_traffic_mirror_target.test"
	description := "test nlb target"
	rName := fmt.Sprintf("tf-acc-test-%s", sdkacctest.RandString(10))
funcource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	testAccPreCheckTrafficMirrorTarget(ctx, t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTrafficMirrorTargetDestroy(ctx),
func
Config: testAccVPCTrafficMirrorTargetConfig_tags1(rName, description, "key1", "value1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckTrafficMirrorTargetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccVPCTrafficMirrorTargetConfig_tags2(rName, description, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckTrafficMirrorTargetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccVPCTrafficMirrorTargetConfig_tags1(rName, description, "key2", "value2"),
func(
	testAccCheckTrafficMirrorTargetExists(ctx, resourceName, &v),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}

func TestAccVPCTrafficMirrorTarget_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.TrafficMirrorTarget
	resourceName := "aws_ec2_traffic_mirror_target.test"
	description := "test nlb target"
	rName := fmt.Sprintf("tf-acc-test-%s", sdkacctest.RandString(10))

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
func
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckTrafficMirrorTargetDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCTrafficMirrorTargetConfig_nlb(rName, description),
Check: resource.ComposeTestCheck
func(
functest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTrafficMirrorTarget(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}


func TestAccVPCTrafficMirrorTarget_gwlb(t *testing.T) {
	ctx := acctest.Context(t)
funcme := fmt.Sprintf("tf-acc-test-%s", sdkacctest.RandString(10))
	description := "test gwlb endpoint target"

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	testAccPreCheckTrafficMirrorTarget(ctx, t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccVPCTrafficMirrorTargetConfig_gwlb(rName, description),
Check: resource.ComposeTestCheck
func(
	resource.TestCheckResourceAttr(resourceName, "description", description),
	resource.TestCheckResourceAttrPair(resourceName, "gateway_load_balancer_endpoint_id", "aws_vpc_endpoint.test", "id"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func testAccCheckTrafficMirrorTargetDestroy(ctx context.Context) resource.TestCheck
funcurn 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_traffic_mirror_target" {
continue
	}

	_, err := tfec2.FindTrafficMirrorTargetByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

funcrn err
func
	return fmt.Errorf("EC2 Traffic Mirror Target %s still exists", rs.Primary.ID)
func
return nil
	}
}


func testAccCheckTrafficMirrorTargetExists(ctx context.Context, n string, v *ec2.TrafficMirrorTarget) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Traffic Mirror Target ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindTrafficMirrorTargetByID(ctx, conn, rs.Primary.ID)

if err != nil {
	return err
}
func *output
funcrn nil
	}
func

func testAccVPCTrafficMirrorTargetConfig_nlb(rName, description string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 2), fmt.Sprintf(`
resource "aws_lb" "test" {
name= %[1]q
internal= true
load_balancer_type = "network"
subnetsws_subnet.test[*].id

enable_deletion_protection = false
}

resource "aws_ec2_traffic_mirror_target" "test" {
description= %[2]q
network_load_balancer_arn = aws_lb.test.arn
}
`, rName, description))
}


func testAccVPCTrafficMirrorTargetConfig_eni(rName, description string) string {
	return acctest.ConfigCompose(
acctest.ConfigVPCWithSubnets(rName, 1),
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
funcurce "aws_instance" "test" {
ami= data.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_type = "t2.micro"
subnet_idet.test[0].id

tags = {
me = %[1]q
}
}

resource "aws_ec2_traffic_mirror_target" "test" {
description = %[2]q
network_interface_id = aws_instance.test.primary_network_interface_id

tags = {
me = %[1]q
}
}
`, rName, description))
func

func testAccVPCTrafficMirrorTargetConfig_tags1(rName, description, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 2), fmt.Sprintf(`
resource "aws_lb" "test" {
name= %[1]q
internal= true
load_balancer_type = "network"
subnetsws_subnet.test[*].id

enable_deletion_protection = false
}

resource "aws_ec2_traffic_mirror_target" "test" {
description= %[2]q
network_load_balancer_arn = aws_lb.test.arn

tags = {
3]q = %[4]q
}
}
`, rName, description, tagKey1, tagValue1))
}


func testAccVPCTrafficMirrorTargetConfig_tags2(rName, description, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 2), fmt.Sprintf(`
funcme= %[1]q
internal= true
load_balancer_type = "network"
subnetsws_subnet.test[*].id

enable_deletion_protection = false
}

resource "aws_ec2_traffic_mirror_target" "test" {
description= %[2]q
network_load_balancer_arn = aws_lb.test.arn

tags = {
3]q = %[4]q
5]q = %[6]q
}
}
`, rName, description, tagKey1, tagValue1, tagKey2, tagValue2))
}


func testAccVPCTrafficMirrorTargetConfig_gwlb(rName, description string) string {
	return acctest.ConfigCompose(
funcSprintf(`
resource "aws_ec2_traffic_mirror_target" "test" {
description
gateway_load_balancer_endpoint_id = aws_vpc_endpoint.test.id

tags = {
me = %[1]q
}
}
`, rName, description))
}


func testAccPreCheckTrafficMirrorTarget(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

	_, err := conn.DescribeTrafficMirrorTargetsWithContext(ctx, &ec2.DescribeTrafficMirrorTargetsInput{})

	if acctest.PreCheckSkipError(err) {
t.Skip("skipping traffic mirror target acceptance test: ", err)
	}

	if err != nil {
t.Fatal("Unexpected PreCheck error: ", err)
func
func