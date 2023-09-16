// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package connect_testimport (
	"fmt"
	"testing"	"github.com/aws/aws-sdk-go/service/connect"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)functionAssociationDataSource_basic(t *testing.T) {
funcme := sdkacctest.RandStringFromCharSet(8, sdkacctest.CharSetAlpha)
	rName2 := sdkacctest.RandomWithPrefix("resource-test-terraform")
	resourceName := "aws_connect_lambda_
function_association.test"
	datasourceName := "data.aws_connect_lambda_
func
	resource.Test(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, connect.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
funcig: testAccLambda
functionAssociationDataSourceConfig_basic(rName, rName2),
Check: resource.ComposeAggregateTestCheck
func(
	resource.TestCheckResourceAttrPair(datasourceName, "instance_id", resourceName, "instance_id"),
	resource.TestCheckResourceAttrPair(datasourceName, "
function_arn"),
),
func
	})
}
func
functionAssociationDataSourceConfig_base(rName string, rName2 string) string {
	return fmt.Sprintf(`
data "aws_partition" "current" {}resource "aws_lambda_
function" "test" {
filename= "test-fixtures/lambdatest.zip"
funcle = aws_iam_role.test.arn
funcntime = "nodejs14.x"
}resource "aws_iam_role" "test" {
name = %[1]q
funcsume_role_policy = <<EOF
{
"Version": "2012-10-17",
func{
"Action": "sts:AssumeRole",
"Principal": {
"Service": "lambda.${data.aws_partition.current.dns_suffix}"
},
"Effect": "Allow",
"Sid": ""
 }
]
}
EOF
}resource "aws_connect_instance" "test" {
identity_management_type = "CONNECT_MANAGED"
inbound_calls_enabled = true
instance_alias= %[2]q
outbound_calls_enabled= true
}resource "aws_connect_lambda_
function_association" "test" {
instance_id= aws_connect_instance.test.id
function_arn = aws_lambda_
function.test.arn
}
`, rName, rName2)
}
func testAccLambda
functionAssociationDataSourceConfig_basic(rName string, rName2 string) string {
	return fmt.Sprintf(testAccLambda
func "aws_connect_lambda_
func
function_arn = aws_connect_lambda_
function_association.test.
function_arn
func
`)
func
func
func
func
func
func