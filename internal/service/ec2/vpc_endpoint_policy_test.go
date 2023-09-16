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
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)


func := acctest.Context(t)
	var endpoint ec2.VpcEndpoint

	resourceName := "aws_vpc_endpoint_policy.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointPolicyConfig_basic(rName, policy1),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCEndpointPolicyConfig_basic(rName, policy2),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
),
func
	})
}


func TestAccVPCEndpointPolicy_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint_policy.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccVPCEndpointPolicyConfig_basic(rName, policy1),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPCEndpointPolicy(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
func
}


func TestAccVPCEndpointPolicy_disappears_endpoint(t *testing.T) {
	ctx := acctest.Context(t)
	var endpoint ec2.VpcEndpoint
	resourceName := "aws_vpc_endpoint_policy.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCEndpointDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCEndpointPolicyConfig_basic(rName, policy1),
Check: resource.ComposeTestCheck
functAccCheckVPCEndpointExists(ctx, resourceName, &endpoint),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPCEndpoint(), "aws_vpc_endpoint.test"),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}
funct policy1 = `
{
"Version": "2012-10-17",
"Statement": [

eadOnly",
l": "*",
 [
db:DescribeTable",
db:ListTables"

 "Allow",
": "*"

]
}
`

const policy2 = `
{
"Version": "2012-10-17",
"Statement": [

llowAll",
 "Allow",
l": {
"*"

 "*",
": "*"

]
}
`


func testAccVPCEndpointPolicyConfig_basic(rName, policy string) string {
	return fmt.Sprintf(`
data "aws_vpc_endpoint_service" "test" {
service = "dynamodb"
}

resource "aws_vpc" "test" {
cidr_block = "10.0.0.0/16"

tags = {
func
}

resource "aws_vpc_endpoint" "test" {
service_name = data.aws_vpc_endpoint_service.test.service_name
vpc_idc.test.id

tags = {
me = %[1]q
}
}

resource "aws_vpc_endpoint_policy" "test" {
vpc_endpoint_id = aws_vpc_endpoint.test.id
policy = <<POLICY
%[2]s
POLICY
}
`, rName, policy)
}
