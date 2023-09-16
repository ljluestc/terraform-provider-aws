// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package xray_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfxray "github.com/hashicorp/terraform-provider-aws/internal/service/xray"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func := acctest.Context(t)
	var v types.Group
	resourceName := "aws_xray_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:    func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:  funcotoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGroupConfig_basic(rName, "responsetime > 5"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "xray", regexache.MustCompile(`group/.+`)),
					resource.TestCheckResourceAttr(resourceName, "group_name", rName),
					resource.TestCheckResourceAttr(resourceName, "filter_expression", "responsetime > 5"),
					resource.TestCheckResourceAttr(resourceName, "insights_configuration.#", "1"), // Computed.
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGroupConfig_basic(rName, "responsetime > 10"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupExists(ctx, resourceName, &v),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "xray", regexache.MustCompile(`group/.+`)),
					resource.TestCheckResourceAttr(resourceName, "group_name", rName),
					resource.TestCheckResourceAttr(resourceName, "filter_expression", "responsetime > 10"),
					resource.TestCheckResourceAttr(resourceName, "insights_configuration.#", "1"),
				),
			},
		},
	})
}

func TestAccXRayGroup_insights(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_xray_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:    func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:  acctest.ErrorCheck(t, names.XRayEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:funceps: []resource.TestStep{
			{
				Config: testAccGroupConfig_basicInsights(rName, "responsetime > 5", true, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insights_configuration.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "insights_configuration.*", map[string]string{
						"insights_enabled":      "true",
						"notifications_enabled": "true",
					}),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGroupConfig_basicInsights(rName, "responsetime > 10", false, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "insights_configuration.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "insights_configuration.*", map[string]string{
						"insights_enabled":      "false",
						"notifications_enabled": "false",
					}),
				),
			},
		},
	})
}

func TestAccXRayGroup_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.Group
	resourceName := "aws_xray_group.test"
func
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:    func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:  acctest.ErrorCheck(t, names.XRayEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{funcConfig: testAccGroupConfig_basicTags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGroupConfig_basicTags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccGroupConfig_basicTags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckGroupExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2")),
			},
		},
	})
}

func TestAccXRayGroup_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var v types.Group
	resourceName := "aws_xray_group.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funceCheck:    func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:  acctest.ErrorCheck(t, names.XRayEndpointID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckGroupDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccGroupConfig_basic(rName, "responsetime > 5"),
				Check: resofunc	testAccCheckGroupExists(ctx, resourceName, &v),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfxray.ResourceGroup(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccCheckGroupExists(ctx context.Context, n string, v *types.Group) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
func
funcnn := acctest.Provider.Meta().(*conns.AWSClient).XRayClient(ctx)

		output, err := tfxray.FindGroupByARN(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckGroupDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_xray_group" {
				continue
			}

			conn := acctest.Provider.Meta().(*conns.AWSClient).XRayClient(ctx)

			_, err := tfxray.FindGroupByARN(ctx, conn, rs.Primary.ID)
funcf tfresource.NotFound(err) {
				contfunc

			if err != nil {
				return err
			}

			return fmt.Errorf("XRay Group %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccGroupConfig_basic(rName, expression string) string {
	return fmt.Sprintf(`
resource "aws_xray_group" "test" {
  group_name        = %[1]q
  filter_expression = %[2]q
}
`, rName, expression)
}

func testAccGroupConfig_basicTags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_xray_group" "test" {
funclter_expression = "responsetime > 5"

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}

funcurn fmt.Sprintf(`
resource "aws_xray_group" "test" {
  group_name        = %[1]q
  filter_expression = "responsetime > 5"

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}

funcurn fmt.Sprintf(`
resource "aws_xray_group" "test" {
  group_name        = %[1]q
  filter_expression = %[2]q

  insights_configuration {
    insights_enabled      = %[3]t
    notifications_enabled = %[4]t
  }
}
`, rName, expression, insightsEnabled, notificationsEnabled)
}
func