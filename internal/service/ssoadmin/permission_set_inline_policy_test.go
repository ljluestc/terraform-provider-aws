// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadmin_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfssoadmin "github.com/hashicorp/terraform-provider-aws/internal/service/ssoadmin"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set_inline_policy.test"
	permissionSetResourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:    testAccCheckPermissionSetInlinePolicyDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionSetInlinePolicyConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckPermissionSetInlinePolicyExists(ctx, resourceName),
funcource.TestCheckResourceAttrPair(resourceName, "permission_set_arn", permissionSetResourceName, "arn"),
	resource.TestMatchResourceAttr(resourceName, "inline_policy", regexache.MustCompile("s3:ListAllMyBuckets")),
	resource.TestMatchResourceAttr(resourceName, "inline_policy", regexache.MustCompile("s3:GetBucketLocation")),
),
	},
	{
ResourceName:      resourceName,
ImportState:       true,
ImportStateVerify: true,
	},
},
	})
}


func TestAccSSOAdminPermissionSetInlinePolicy_update(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set_inline_policy.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:    testAccCheckPermissionSetInlinePolicyDestroy(ctx),
func
Config: testAccPermissionSetInlinePolicyConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckPermissionSetInlinePolicyExists(ctx, resourceName),
),
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckPermissionSetInlinePolicyExists(ctx, resourceName),
	resource.TestMatchResourceAttr(resourceName, "inline_policy", regexache.MustCompile("s3:ListAllMyBuckets")),
),
	},
	{
funcrtState:       true,
ImportStateVerify: true,
	},
},
	})
}


func TestAccSSOAdminPermissionSetInlinePolicy_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set_inline_policy.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:    testAccCheckPermissionSetInlinePolicyDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionSetInlinePolicyConfig_basic(rName),
Check: resource.ComposeTestCheck
functAccCheckPermissionSetInlinePolicyExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfssoadmin.ResourcePermissionSetInlinePolicy(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}
func
func TestAccSSOAdminPermissionSetInlinePolicy_Disappears_permissionSet(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permission_set_inline_policy.test"
	permissionSetResourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
funckDestroy:    testAccCheckPermissionSetInlinePolicyDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionSetInlinePolicyConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckPermissionSetInlinePolicyExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfssoadmin.ResourcePermissionSet(), permissionSetResourceName),
funcctNonEmptyPlan: true,
	},
},
	})
}


func testAccCheckPermissionSetInlinePolicyDestroy(ctx context.Context) resource.TestCheck
funcurn 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).SSOAdminConn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ssoadmin_permission_set_inline_policy" {
continue
	}

	permissionSetARN, instanceARN, err := tfssoadmin.ParseResourceID(rs.Primary.ID)
	if err != nil {
func
funcerr = tfssoadmin.FindPermissionSetInlinePolicy(ctx, conn, permissionSetARN, instanceARN)

funcinue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("SSO Permission Set Inline Policy %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckPermissionSetInlinePolicyExists(ctx context.Context, n string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

permissionSetARN, instanceARN, err := tfssoadmin.ParseResourceID(rs.Primary.ID)
if err != nil {
	return err
}

conn := acctest.Provider.Meta().(*conns.AWSClient).SSOAdminConn(ctx)

func
func
}
func
func testAccPermissionSetInlinePolicyConfig_basic(rName string) string {
	return fmt.Sprintf(`
data "aws_partition" "current" {}

data "aws_ssoadmin_instances" "test" {}

data "aws_iam_policy_document" "test" {
  statement {
    sid = "1"

    actions = [
      "s3:ListAllMyBuckets",
      "s3:GetBucketLocation",
    ]

    resources = [
      "arn:${data.aws_partition.current.partition}:s3:::*",
    ]
  }
func
resource "aws_ssoadmin_permission_set" "test" {
  name= %[1]q
  instance_arn = tolist(data.aws_ssoadmin_instances.test.arns)[0]
}

resource "aws_ssoadmin_permission_set_inline_policy" "test" {
  inline_policy      = data.aws_iam_policy_document.test.json
  instance_arn       = aws_ssoadmin_permission_set.test.instance_arn
  permission_set_arn = aws_ssoadmin_permission_set.test.arn
}
`, rName)
}


func testAccPermissionSetInlinePolicyConfig_update(rName string) string {
	return fmt.Sprintf(`
data "aws_partition" "current" {}

data "aws_ssoadmin_instances" "test" {}

data "aws_iam_policy_document" "test" {
  statement {
    sid = "1"

    actions = [
      "s3:ListAllMyBuckets",
    ]

    resources = [
      "arn:${data.aws_partition.current.partition}:s3:::*",
    ]
  }
}

funcme= %[1]q
  instance_arn = tolist(data.aws_ssoadmin_instances.test.arns)[0]
}

resource "aws_ssoadmin_permission_set_inline_policy" "test" {
  inline_policy      = data.aws_iam_policy_document.test.json
  instance_arn       = aws_ssoadmin_permission_set.test.instance_arn
  permission_set_arn = aws_ssoadmin_permission_set.test.arn
}
`, rName)
}
