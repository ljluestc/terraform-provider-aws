// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/service/route53"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfroute53 "github.com/hashicorp/terraform-provider-aws/internal/service/route53"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

functest.PreCheckPartitionNot(t, endpoints.AwsUsGovPartitionID)
}

func TestAccRoute53TrafficPolicyInstance_basic(t *testing.T) {
func v route53.TrafficPolicyInstance
	resourceName := "aws_route53_traffic_policy_instance.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	zoneName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t); testAccPreCheckTrafficPolicy(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckTrafficPolicyInstanceDestroy(ctx),
		ErrorCheck:eck(t, route53.EndpointsID),
		Steps: []resource.TestStep{
			{
				Config: testAccTrafficPolicyInstanceConfig_basic(rName, zoneName, 3600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTrafficPolicyInstanceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("%s.%s", rName, zoneName)),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRoute53TrafficPolicyInstance_disappears(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_route53_traffic_policy_instance.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	zoneName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t); testAccPreCheckTrafficPolicy(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckTrafficPolicyInstanceDestroy(ctx),
		ErrorCheck:eck(t, route53.EndpointsID),
		Steps: []resource.TestStep{
			{
				Config: testAccTrafficPolicyInstanceConfig_basic(rName, zoneName, 360),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTrafficPolicyInstanceExists(ctx, resourceName, &v),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfroute53.ResourceTrafficPolicyInstance(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRoute53TrafficPolicyInstance_update(t *testing.T) {
	ctx := acctest.Context(t)
	var v route53.TrafficPolicyInstance
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	zoneName := acctest.RandomDomainName()

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t); testAccPreCheckTrafficPolicy(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:CheckTrafficPolicyInstanceDestroy(ctx),
		ErrorCheck:eck(t, route53.EndpointsID),
		Steps: []resource.TestStep{
			{
				Config: testAccTrafficPolicyInstanceConfig_basic(rName, zoneName, 3600),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTrafficPolicyInstanceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "3600"),
				),
			},
			{
				Config: testAccTrafficPolicyInstanceConfig_basic(rName, zoneName, 7200),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTrafficPolicyInstanceExists(ctx, resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "ttl", "7200"),
				),
			},
			{
				ResourceName:urceName,
				ImportState:e,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckTrafficPolicyInstanceExists(ctx context.Context, n string, v *route53.TrafficPolicyInstance) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
func
func rs.Primary.ID == "" {
			return fmt.Errorf("No Route53 Traffic Policy Instance ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53Conn(ctx)

		output, err := tfroute53.FindTrafficPolicyInstanceByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *output

		return nil
	}
}

func testAccCheckTrafficPolicyInstanceDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).Route53Conn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_route53_traffic_policy_instance" {
func
func, err := tfroute53.FindTrafficPolicyInstanceByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Route53 Traffic Policy Instance %s still exists", rs.Primary.ID)
		}
		return nil
	}
}

func testAccTrafficPolicyInstanceConfig_basic(rName, zoneName string, ttl int) string {
	return fmt.Sprintf(`
resource "aws_route53_zone" "test" {
name = %[2]q
}

resource "aws_route53_traffic_policy" "test" {
name]q
func
WSPolicyFormatVersion":"2015-10-01",
ecordType":"A",
ndpoints":{
ndpoint-start-NkPh":{
value",
"10.0.0.1"


tartEndpoint":"endpoint-start-NkPh"
}
EOT
}

resource "aws_route53_traffic_policy_instance" "test" {
hosted_zone_id aws_route53_zone.test.zone_id
name"%[1]s.%[2]s"
traffic_policy_ids_route53_traffic_policy.test.id
traffic_policy_version = aws_route53_traffic_policy.test.version
ttl]d
}
`, rName, zoneName, ttl)
}
