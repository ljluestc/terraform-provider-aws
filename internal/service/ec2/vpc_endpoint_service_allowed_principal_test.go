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
	rName := sdkacctest.RandomWithPrefix("tfacctest")

	resourceName := "aws_vpc_endpoint_service_allowed_principal.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceAllowedPrincipalDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointServiceAllowedPrincipalConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckVPCEndpointServiceAllowedPrincipalExists(ctx, resourceName),
funcource.TestCheckResourceAttrPair(resourceName, "vpc_endpoint_service_id", "aws_vpc_endpoint_service.test", "id"),
	resource.TestCheckResourceAttrPair(resourceName, "principal_arn", "data.aws_iam_session_context.current", "issuer_arn"),
),
	},
},
	})
}


func TestAccVPCEndpointServiceAllowedPrincipal_multiple(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix("tfacctest")
funcourceName := "aws_vpc_endpoint_service_allowed_principal.test"
	serviceResourceName := "aws_vpc_endpoint_service.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceAllowedPrincipalDestroy(ctx),
func
Config: testAccVPCEndpointServiceAllowedPrincipalConfig_Multiple(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckVPCEndpointServiceAllowedPrincipalExists(ctx, resourceName),
	resource.TestMatchResourceAttr(resourceName, "id", regexache.MustCompile(`^vpce-svc-perm-\w{17}$`)),
	resource.TestCheckResourceAttrPair(resourceName, "vpc_endpoint_service_id", "aws_vpc_endpoint_service.test", "id"),
	resource.TestCheckResourceAttr(serviceResourceName, "allowed_principals.#", "1"),
func
	},
},
	})
}


func TestAccVPCEndpointServiceAllowedPrincipal_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix("tfacctest")

	resourceName := "aws_vpc_endpoint_service_allowed_principal.test"
	tagResourceName := "aws_ec2_tag.test"
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointServiceAllowedPrincipalDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointServiceAllowedPrincipalConfig_tag(rName),
func(
	testAccCheckVPCEndpointServiceAllowedPrincipalExists(ctx, resourceName),
	resource.TestCheckResourceAttrPair(tagResourceName, "resource_id", resourceName, "id"),
	resource.TestCheckResourceAttr(tagResourceName, "key", "Name"),
	resource.TestCheckResourceAttr(tagResourceName, "value", rName),
),
	},
},
func


func TestAccVPCEndpointServiceAllowedPrincipal_migrateID(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix("tfacctest")

	resourceName := "aws_vpc_endpoint_service_allowed_principal.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funckDestroy: testAccCheckVPCEndpointServiceAllowedPrincipalDestroy(ctx),
Steps: []resource.TestStep{
	{
ExternalProviders: map[string]resource.ExternalProvider{
	"aws": {
Source:shicorp/aws",
VersionConstraint: "4.63.0",
	},
funcig: testAccVPCEndpointServiceAllowedPrincipalConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckVPCEndpointServiceAllowedPrincipalExists(ctx, resourceName),
),
	},
	{
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Config:stAccVPCEndpointServiceAllowedPrincipalConfig_basic(rName),
PlanOnly:true,
	},
},
	})
func
// Verify that the resource returns an ID usable for creating an `aws_ec2_tag`

func TestAccVPCEndpointServiceAllowedPrincipal_migrateAndTag(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix("tfacctest")

	resourceName := "aws_vpc_endpoint_service_allowed_principal.test"
	tagResourceName := "aws_ec2_tag.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:test.ErrorCheck(t, ec2.EndpointsID),
CheckDestroy: testAccCheckVPCEndpointServiceAllowedPrincipalDestroy(ctx),
func
ExternalProviders: map[string]resource.ExternalProvider{
	"aws": {
Source:shicorp/aws",
VersionConstraint: "4.63.0",
	},
},
Config: testAccVPCEndpointServiceAllowedPrincipalConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
functAccCheckVPCEndpointServiceAllowedPrincipalExists(ctx, resourceName),
),
	},
	{
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Config:stAccVPCEndpointServiceAllowedPrincipalConfig_tag(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckVPCEndpointServiceAllowedPrincipalExists(ctx, resourceName),
	resource.TestMatchResourceAttr(resourceName, "id", regexache.MustCompile(`^vpce-svc-perm-\w{17}$`)),
	resource.TestCheckResourceAttrPair(tagResourceName, "resource_id", resourceName, "id"),
	resource.TestCheckResourceAttr(tagResourceName, "key", "Name"),
	resource.TestCheckResourceAttr(tagResourceName, "value", rName),
func
},
	})
}


func testAccCheckVPCEndpointServiceAllowedPrincipalDestroy(ctx context.Context) resource.TestCheck
func {
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_vpc_endpoint_service_allowed_principal" {
continue
	}

	_, err := tfec2.FindVPCEndpointServicePermission(ctx, conn, rs.Primary.Attributes["vpc_endpoint_service_id"], rs.Primary.Attributes["principal_arn"])

	if tfresource.NotFound(err) {
continue
	}
funcerr != nil {
func

func

return nil
	}
}


func testAccCheckVPCEndpointServiceAllowedPrincipalExists(ctx context.Context, n string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 VPC Endpoint Service Allowed Principal ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

_, err := tfec2.FindVPCEndpointServicePermission(ctx, conn, rs.Primary.Attributes["vpc_endpoint_service_id"], rs.Primary.Attributes["principal_arn"])

return err
	}
func
func testAccVPCEndpointServiceAllowedPrincipalConfig_basic(rName string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseNetworkLoadBalancer(rName, 1), fmt.Sprintf(`
func
data "aws_iam_session_context" "current" {
arn = data.aws_caller_identity.current.arn
}

resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn

tags = {
me = %[1]q
}
}

resource "aws_vpc_endpoint_service_allowed_principal" "test" {
vpc_endpoint_service_id = aws_vpc_endpoint_service.test.id

principal_arn = data.aws_iam_session_context.current.issuer_arn
}
func


func testAccVPCEndpointServiceAllowedPrincipalConfig_Multiple(rName string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceConfig_baseNetworkLoadBalancer(rName, 1), fmt.Sprintf(`
data "aws_caller_identity" "current" {}
data "aws_partition" "current" {}

data "aws_iam_session_context" "current" {
arn = data.aws_caller_identity.current.arn
}

resource "aws_vpc_endpoint_service" "test" {
acceptance_required
network_load_balancer_arns = aws_lb.test[*].arn
allowed_principals= ["arn:${data.aws_partition.current.partition}:iam::123456789012:root"]

tags = {
me = %[1]q
}

lifecycle {
nore_changes = [
rincipals

}
func
resource "aws_vpc_endpoint_service_allowed_principal" "test" {
vpc_endpoint_service_id = aws_vpc_endpoint_service.test.id

principal_arn = data.aws_iam_session_context.current.issuer_arn
}
`, rName))
}


func testAccVPCEndpointServiceAllowedPrincipalConfig_tag(rName string) string {
	return acctest.ConfigCompose(testAccVPCEndpointServiceAllowedPrincipalConfig_basic(rName), fmt.Sprintf(`
resource "aws_ec2_tag" "test" {
resource_id = aws_vpc_endpoint_service_allowed_principal.test.id

keyName"
value = %[1]q
}
`, rName))
}
func