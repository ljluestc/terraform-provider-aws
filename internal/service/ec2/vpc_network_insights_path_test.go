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
	resourceName := "aws_ec2_network_insights_path.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsPathDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInsightsPathConfig_basic(rName, "tcp"),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
funcource.TestCheckResourceAttrPair(resourceName, "destination", "aws_network_interface.test.1", "id"),
	resource.TestCheckResourceAttrPair(resourceName, "destination_arn", "aws_network_interface.test.1", "arn"),
	resource.TestCheckResourceAttr(resourceName, "destination_ip", ""),
	resource.TestCheckResourceAttr(resourceName, "destination_port", "0"),
	resource.TestCheckResourceAttr(resourceName, "protocol", "tcp"),
	resource.TestCheckResourceAttrPair(resourceName, "source", "aws_network_interface.test.0", "id"),
	resource.TestCheckResourceAttrPair(resourceName, "source_arn", "aws_network_interface.test.0", "arn"),
	resource.TestCheckResourceAttr(resourceName, "source_ip", ""),
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


func TestAccVPCNetworkInsightsPath_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_path.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsPathDestroy(ctx),
func
Config: testAccVPCNetworkInsightsPathConfig_basic(rName, "udp"),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceNetworkInsightsPath(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccVPCNetworkInsightsPath_tags(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_path.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsPathDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInsightsPathConfig_tags1(rName, "key1", "value1"),
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccVPCNetworkInsightsPathConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccVPCNetworkInsightsPathConfig_tags1(rName, "key2", "value2"),
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}

func TestAccVPCNetworkInsightsPath_sourceAndDestinationARN(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_path.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsPathDestroy(ctx),
func
Config: testAccVPCNetworkInsightsPathConfig_sourceAndDestinationARN(rName, "tcp"),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	resource.TestCheckResourceAttrPair(resourceName, "destination", "aws_network_interface.test.1", "id"),
	resource.TestCheckResourceAttrPair(resourceName, "destination_arn", "aws_network_interface.test.1", "arn"),
funcource.TestCheckResourceAttrPair(resourceName, "source_arn", "aws_network_interface.test.0", "arn"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
func
}


func TestAccVPCNetworkInsightsPath_sourceIP(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_path.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsPathDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInsightsPathConfig_sourceIP(rName, "1.1.1.1"),
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "source_ip", "1.1.1.1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccVPCNetworkInsightsPathConfig_sourceIP(rName, "8.8.8.8"),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "source_ip", "8.8.8.8"),
func
},
	})
}


func TestAccVPCNetworkInsightsPath_destinationIP(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_path.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsPathDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInsightsPathConfig_destinationIP(rName, "1.1.1.1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "destination_ip", "1.1.1.1"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "destination_ip", "8.8.8.8"),
),
	},
},
	})
func

func TestAccVPCNetworkInsightsPath_destinationPort(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_path.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsPathDestroy(ctx),
func
Config: testAccVPCNetworkInsightsPathConfig_destinationPort(rName, 80),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "destination_port", "80"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccVPCNetworkInsightsPathConfig_destinationPort(rName, 443),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsPathExists(ctx, resourceName),
func
	},
},
	})
}


func testAccCheckNetworkInsightsPathExists(ctx context.Context, n string) resource.TestCheck
funcurn 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Network Insights Path ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

func
return err
	}
}


func testAccCheckNetworkInsightsPathDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
func
funcrs.Type != "aws_ec2_network_insights_path" {
continue
func
	_, err := tfec2.FindNetworkInsightsPathByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Network Insights Path %s still exists", rs.Primary.ID)
}

return nil
	}
}


funcurn acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
funcunt = 2

func
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_network_insights_path" "test" {
  sourcework_interface.test[0].id
  destination = aws_network_interface.test[1].id
  protocol%[2]q
}
`, rName, protocol))
}


func testAccVPCNetworkInsightsPathConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  count = 2

  subnet_id = aws_subnet.test[0].id

  tags = {
me = %[1]q
  }
}
funcurce "aws_ec2_network_insights_path" "test" {
  sourcework_interface.test[0].id
  destination = aws_network_interface.test[1].id
  protocol"tcp"

  tags = {
2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1))
}


func testAccVPCNetworkInsightsPathConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  count = 2

  subnet_id = aws_subnet.test[0].id

  tags = {
func
}

resource "aws_ec2_network_insights_path" "test" {
  sourcework_interface.test[0].id
  destination = aws_network_interface.test[1].id
  protocol"tcp"

  tags = {
2]q = %[3]q
4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2))
}


func testAccVPCNetworkInsightsPathConfig_sourceAndDestinationARN(rName, protocol string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  count = 2

  subnet_id = aws_subnet.test[0].id

  tags = {
func
}

resource "aws_ec2_network_insights_path" "test" {
  sourcework_interface.test[0].arn
  destination = aws_network_interface.test[1].arn
  protocol%[2]q
}
`, rName, protocol))
}


func testAccVPCNetworkInsightsPathConfig_sourceIP(rName, sourceIP string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test[0].id

  tags = {
func
}

resource "aws_ec2_network_insights_path" "test" {
  sourceernet_gateway.test.id
  destination = aws_network_interface.test.id
  protocol"tcp"
  source_ip[2]q

  tags = {
me = %[1]q
  }
}
`, rName, sourceIP))
}


func testAccVPCNetworkInsightsPathConfig_destinationIP(rName, destinationIP string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id
funcgs = {
me = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test[0].id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_network_insights_path" "test" {
  source= aws_network_interface.test.id
  destinationaws_internet_gateway.test.id
  protocol
  destination_ip = %[2]q

  tags = {
me = %[1]q
  }
}
`, rName, destinationIP))
}


func testAccVPCNetworkInsightsPathConfig_destinationPort(rName string, destinationPort int) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  count = 2

func
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_network_insights_path" "test" {
  source  = aws_network_interface.test[0].id
  destinationwork_interface.test[1].id
  protocol= "tcp"
  destination_port = %[2]d

  tags = {
me = %[1]q
  }
}
`, rName, destinationPort))
}
func