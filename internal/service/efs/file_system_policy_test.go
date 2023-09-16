// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efs_test

import (
"context"
"fmt"
"testing"

"github.com/aws/aws-sdk-go/service/efs"
sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
"github.com/hashicorp/terraform-plugin-testing/helper/resource"
"github.com/hashicorp/terraform-plugin-testing/terraform"
"github.com/hashicorp/terraform-provider-aws/internal/acctest"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
tfefs "github.com/hashicorp/terraform-provider-aws/internal/service/efs"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func:= acctest.Context(t)
var desc efs.DescribeFileSystemPolicyOutput
resourceName := "aws_efs_file_system_policy.test"
rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

resource.ParallelTest(t, resource.TestCase{
PreCheck:func() { acctest.PreCheck(ctx, t) },
ErrorCheck:funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckFileSystemPolicyDestroy(ctx),
Steps: []resource.TestStep{
{
Config: testAccFileSystemPolicyConfig_basic(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckFileSystemPolicyExists(ctx, resourceName, &desc),
resource.TestCheckResourceAttrSet(resourceName, "policy"),
),
},
{
ResourceName:,
ImportState:true,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"bypass_policy_lockout_safety_check"},
},
{
Config: testAccFileSystemPolicyConfig_updated(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckFileSystemPolicyExists(ctx, resourceName, &desc),
resource.TestCheckResourceAttrSet(resourceName, "policy"),
),
},
},
})
}

func TestAccEFSFileSystemPolicy_disappears(t *testing.T) {
ctx := acctest.Context(t)
funcurceName := "aws_efs_file_system_policy.test"
rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

resource.ParallelTest(t, resource.TestCase{
PreCheck:func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, efs.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestrofuncs: []resource.TestStep{
{
Config: testAccFileSystemPolicyConfig_basic(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckFileSystemPolicyExists(ctx, resourceName, &desc),
acctest.CheckResourceDisappears(ctx, acctest.Provider, tfefs.ResourceFileSystemPolicy(), resourceName),
),
ExpectNonEmptyPlan: true,
},
},
})
}

func TestAccEFSFileSystemPolicy_policyBypass(t *testing.T) {
ctx := acctest.Context(t)
var desc efs.DescribeFileSystemPolicyOutput
resourceName := "aws_efs_file_system_policy.test"
func
resource.ParallelTest(t, resource.TestCase{
PreCheck:func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, efs.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckFileSystemPolicyDestroy(ctx),
Steps: []resource.TestStep{
{funcig: testAccFileSystemPolicyConfig_basic(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckFileSystemPolicyExists(ctx, resourceName, &desc),
resource.TestCheckResourceAttr(resourceName, "bypass_policy_lockout_safety_check", "false"),
),
},
{
ResourceName:,
ImportState:true,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"bypass_policy_lockout_safety_check"},
},
{
Config: testAccFileSystemPolicyConfig_bypass(rName, true),
Check: resource.ComposeTestCheckFunc(
testAccCheckFileSystemPolicyExists(ctx, resourceName, &desc),
resource.TestCheckResourceAttr(resourceName, "bypass_policy_lockout_safety_check", "true"),
),
},
},
})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/21968
func TestAccEFSFileSystemPolicy_equivalentPolicies(t *testing.T) {
ctx := acctest.Context(t)
var desc efs.DescribeFileSystemPolicyOutput
resourceName := "aws_efs_file_system_policy.test"
rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, efs.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckFileSystemPolicyDestroy(ctx),
Steps: []resource.TestStep{
{
Config: testAccFileSystemPolicyConfig_firstEquivalent(rName),
Check: resofuncAccCheckFileSystemPolicyExists(ctx, resourceName, &desc),
resource.TestCheckResourceAttrSet(resourceName, "policy"),
),
},
{
Config:tAccFileSystemPolicyConfig_secondEquivalent(rName),
PlanOnly: true,
},
},
})
}

// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/19245
func TestAccEFSFileSystemPolicy_equivalentPoliciesIAMPolicyDoc(t *testing.T) {
ctx := acctest.Context(t)
var desc efs.DescribeFileSystemPolicyOutput
resourceName := "aws_efs_file_system_policy.test"
rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

resource.ParallelTest(t, resource.TestCase{
PreCheck:func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:testAccCheckFileSystemPolicyDestroy(ctx),
Steps: []resource.TestStep{
{
Config: testAccFileSystemPolicyConfig_equivalentIAMDoc(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckFileSystemPolicyExists(ctx, resourceName, &desc),
resource.Tefunc
},
{
Config:tAccFileSystemPolicyConfig_equivalentIAMDoc(rName),
PlanOnly: true,
},
},
})
}

func testAccCheckFileSystemPolicyDestroy(ctx context.Context) resource.TestCheckFunc {
return func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EFSConn(ctx)

for _, rs := range s.RootModule().Resources {
if rs.Type != "aws_efs_file_system_policy" {
continue
}

_, err := tfefs.FindFileSystemPolicyByID(ctx, conn, rs.Primary.ID)
funcfresource.NotFound(err) {
continufunc

if err != nil {
return err
}

return fmt.Errorf("EFS File System Policy %s still exists", rs.Primary.ID)
}

return nil
}
}

func testAccCheckFileSystemPolicyExists(ctx context.Context, n string, v *efs.DescribeFileSystemPolicyOutput) resource.TestCheckFunc {
return func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
return fmt.Errorf("No EFS File System Policy ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EFSConn(ctx)
funcut, err := tfefs.FindFileSystemPolicyByID(ctx, conn, rs.Primary.ID)
funcrr != nil {
return err
}

*v = *output

return nil
}
}

func testAccFileSystemPolicyConfig_basic(rName string) string {
return fmt.Sprintf(`
resource "aws_efs_file_system" "test" {
creation_token = %[1]q
}

resource "aws_efs_file_system_policy" "test" {
file_system_id = aws_efs_file_system.test.id

policy = <<POLICY
{
ersion": "2012-10-17",
d": "ExamplePolicy01",
tatement": [
functatement01",
llow",
 {
 "AWS": "*"

"${aws_efs_file_system.test.arn}",

 "elasticfilesystem:ClientMount",
 "elasticfilesystem:ClientWrite"

 {
 "Bool": {
aws:SecureTransport": "true"
 }



}
POLICY
}
`, rName)
}

func testAccFileSystemPolicyConfig_updated(rName string) string {
return fmt.Sprintf(`
resource "aws_efs_file_system" "test" {
creation_token = %[1]q
}

resource "aws_efs_file_system_policy" "test" {
file_system_id = aws_efs_file_system.test.id

policy = <<POLICY
{
ersion": "2012-10-17",
d": "ExamplePolicy01",
tatement": [

func",
 {
 "AWS": "*"

"${aws_efs_file_system.test.arn}",
lasticfilesystem:ClientMount",
 {
 "Bool": {
aws:SecureTransport": "true"
 }



}
POLICY
}
`, rName)
}

func testAccFileSystemPolicyConfig_bypass(rName string, bypass bool) string {
return fmt.Sprintf(`
resource "aws_efs_file_system" "test" {
creation_token = %[1]q
}

resource "aws_efs_file_system_policy" "test" {
file_system_id = aws_efs_file_system.test.id

bypass_policy_lockout_safety_check = %[2]t

policy = <<POLICY
{
ersion": "2012-10-17",
d": "ExamplePolicy01",
tatement": [
functatement01",
llow",
 {
 "AWS": "*"

"${aws_efs_file_system.test.arn}",

 "elasticfilesystem:ClientMount",
 "elasticfilesystem:ClientWrite"

 {
 "Bool": {
aws:SecureTransport": "true"
 }



}
POLICY
}
`, rName, bypass)
}

func testAccFileSystemPolicyConfig_firstEquivalent(rName string) string {
return fmt.Sprintf(`
resource "aws_efs_file_system" "test" {
creation_token = %[1]q
}

resource "aws_efs_file_system_policy" "test" {
file_system_id = aws_efs_file_system.test.id

policy = jsonencode({
rsion = "2012-10-17"
= "ePolicy01"
atement = [{
"ExleStatement01"
 = "Allow"
pal = {
= "*"
func aws_efs_file_system.test.arn
 = [
sticfilesystem:ClientMount",
sticfilesystem:ClientWrite",

ion = {
 = {
ureTransport" = "true"



})
}
`, rName)
}

func testAccFileSystemPolicyConfig_secondEquivalent(rName string) string {
return fmt.Sprintf(`
resource "aws_efs_file_system" "test" {
creation_token = %[1]q
}

resource "aws_efs_file_system_policy" "test" {
file_system_id = aws_efs_file_system.test.id

policy = jsonencode({
rsion = "2012-10-17"
= "ePolicy01"
atement = [{
"ExleStatement01"
 = "Allow"
pal = {
= "*"

func
sticfilesystem:ClientWrite",
sticfilesystem:ClientMount",

ion = {
 = {
ureTransport" = ["true"]



})
}
`, rName)
}

func testAccFileSystemPolicyConfig_equivalentIAMDoc(rName string) string {
return fmt.Sprintf(`
resource "aws_efs_file_system" "test" {
creation_token = %[1]q
}

resource "aws_efs_file_system_policy" "test" {
file_system_id = aws_efs_file_system.test.id
policys_iam_policy_document.test.json
}

data "aws_iam_policy_document" "test" {
version = "2012-10-17"

statement {
d = "Allow mount and write"

tions = [
icfilesystem:ClientWrite",
funclesystem:ClientMount",


sources = [aws_efs_file_system.test.arn]

incipals {
= "A
fiers = ["*"]

}

statement {
d = ce in-transit encryption for all clients"
fect = ny"
tions= []
sources = [aws_efs_file_system.test.arn]

incipals {
= "A
fiers = ["*"]


ndition {
= "B"
le = "aws:SecureTransport"
= ["fa"]

}
}
`, rName)
}
