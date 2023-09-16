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
	resourceName := "aws_ec2_network_insights_analysis.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsAnalysisDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInsightsAnalysisConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckNetworkInsightsAnalysisExists(ctx, resourceName),
funcource.TestCheckResourceAttr(resourceName, "filter_in_arns.#", "0"),
	resource.TestCheckResourceAttrPair(resourceName, "network_insights_path_id", "aws_ec2_network_insights_path.test", "id"),
	resource.TestCheckResourceAttr(resourceName, "path_found", "true"),
	acctest.CheckResourceAttrRFC3339(resourceName, "start_date"),
	resource.TestCheckResourceAttr(resourceName, "status", "succeeded"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "wait_for_completion", "true"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"wait_for_completion"},
	},
},
	})
}


func TestAccVPCNetworkInsightsAnalysis_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_analysis.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsAnalysisDestroy(ctx),
func
Config: testAccVPCNetworkInsightsAnalysisConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsAnalysisExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceNetworkInsightsAnalysis(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccVPCNetworkInsightsAnalysis_tags(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_analysis.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsAnalysisDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInsightsAnalysisConfig_tags1(rName, "key1", "value1"),
func(
	testAccCheckNetworkInsightsAnalysisExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ourceName,
funcrtStateVerify:
ImportStateVerifyIgnore: []string{"wait_for_completion"},
	},
	{
Config: testAccVPCNetworkInsightsAnalysisConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsAnalysisExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccVPCNetworkInsightsAnalysisConfig_tags1(rName, "key2", "value2"),
func(
	testAccCheckNetworkInsightsAnalysisExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}

func TestAccVPCNetworkInsightsAnalysis_filterInARNs(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_analysis.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsAnalysisDestroy(ctx),
func
Config: testAccVPCNetworkInsightsAnalysisConfig_filterInARNs(rName, "vpc-peering-connection/pcx-fakearn1"),
Check: resource.ComposeTestCheck
func(
	testAccCheckNetworkInsightsAnalysisExists(ctx, resourceName),
	acctest.MatchResourceAttrRegionalARN(resourceName, "filter_in_arns.0", "ec2", regexache.MustCompile(`vpc-peering-connection/pcx-fakearn1$`)),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"wait_for_completion"},
	},
	{
Config: testAccVPCNetworkInsightsAnalysisConfig_filterInARNs(rName, "vpc-peering-connection/pcx-fakearn2"),
func(
	testAccCheckNetworkInsightsAnalysisExists(ctx, resourceName),
	acctest.MatchResourceAttrRegionalARN(resourceName, "filter_in_arns.0", "ec2", regexache.MustCompile(`vpc-peering-connection/pcx-fakearn2$`)),
),
	},
},
	})
}


func TestAccVPCNetworkInsightsAnalysis_waitForCompletion(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ec2_network_insights_analysis.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckNetworkInsightsAnalysisDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCNetworkInsightsAnalysisConfig_waitForCompletion(rName, false),
Check: resource.ComposeTestCheck
functAccCheckNetworkInsightsAnalysisExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "wait_for_completion", "false"),
	resource.TestCheckResourceAttr(resourceName, "status", "running"),
),
	},
	{
Config: testAccVPCNetworkInsightsAnalysisConfig_waitForCompletion(rName, true),
func(
	testAccCheckNetworkInsightsAnalysisExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "wait_for_completion", "true"),
),
	},
},
	})
}
func
func testAccCheckNetworkInsightsAnalysisExists(ctx context.Context, n string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}
funcs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Network Insights Analysis ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

_, err := tfec2.FindNetworkInsightsAnalysisByID(ctx, conn, rs.Primary.ID)

return err
	}
func
func testAccCheckNetworkInsightsAnalysisDestroy(ctx context.Context) resource.TestCheck
func {
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ec2_network_insights_analysis" {
continue
	}

	_, err := tfec2.FindNetworkInsightsAnalysisByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Network Insights Analysis %s still exists", rs.Primary.ID)
func
func
}
func
func testAccVPCNetworkInsightsAnalysisConfig_base(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigVPCWithSubnets(rName, 1), fmt.Sprintf(`
resource "aws_network_interface" "test" {
  count = 2

  subnet_id = aws_subnet.test[0].id

  tags = {
me = %[1]q
  }
}

resource "aws_ec2_network_insights_path" "test" {
  sourcework_interface.test[0].id
  destination = aws_network_interface.test[1].id
  protocol"tcp"

  tags = {
me = %[1]q
  }
}
`, rName))
}


funcurn acctest.ConfigCompose(testAccVPCNetworkInsightsAnalysisConfig_base(rName), `
resource "aws_ec2_network_insights_analysis" "test" {
  network_insights_path_id = aws_ec2_network_insights_path.test.id
}
`)
}


func testAccVPCNetworkInsightsAnalysisConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInsightsAnalysisConfig_base(rName), fmt.Sprintf(`
resource "aws_ec2_network_insights_analysis" "test" {
  network_insights_path_id = aws_ec2_network_insights_path.test.id

  tags = {
1]q = %[2]q
  }
}
`, tagKey1, tagValue1))
}


func testAccVPCNetworkInsightsAnalysisConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(testAccVPCNetworkInsightsAnalysisConfig_base(rName), fmt.Sprintf(`
resource "aws_ec2_network_insights_analysis" "test" {
  network_insights_path_id = aws_ec2_network_insights_path.test.id
funcgs = {
1]q = %[2]q
3]q = %[4]q
  }
}
`, tagKey1, tagValue1, tagKey2, tagValue2))
}


funcurn acctest.ConfigCompose(testAccVPCNetworkInsightsAnalysisConfig_base(rName), fmt.Sprintf(`
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
data "aws_partition" "current" {}

resource "aws_ec2_network_insights_analysis" "test" {
  network_insights_path_id = aws_ec2_network_insights_path.test.id
  filter_in_arns  = ["arn:${data.aws_partition.current.partition}:ec2:${data.aws_region.current.name}:${data.aws_caller_identity.current.id}:%[2]s"]

  tags = {
me = %[1]q
  }
}
func


func testAccVPCNetworkInsightsAnalysisConfig_waitForCompletion(rName string, waitForCompletion bool) string {
	return acctest.ConfigCompose(testAccVPCNetworkInsightsAnalysisConfig_base(rName), fmt.Sprintf(`
resource "aws_ec2_network_insights_analysis" "test" {
  network_insights_path_id = aws_ec2_network_insights_path.test.id
  wait_for_completion

  tags = {
me = %[1]q
  }
}
`, rName, waitForCompletion))
funcfunc