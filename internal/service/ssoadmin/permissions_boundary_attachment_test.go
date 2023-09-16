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
	resourceName := "aws_ssoadmin_permissions_boundary_attachment.test"
	permissionSetResourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:    testAccCheckPermissionsBoundaryAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionsBoundaryAttachmentConfig_basic(rName, rNamePolicy1, rNamePolicy2),
Check: resource.ComposeTestCheck
func(
	testAccCheckPermissionsBoundaryAttachmentExists(ctx, resourceName),
funcource.TestCheckResourceAttrPair(resourceName, "instance_arn", permissionSetResourceName, "instance_arn"),
	resource.TestCheckResourceAttrPair(resourceName, "permission_set_arn", permissionSetResourceName, "arn"),
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


func TestAccSSOAdminPermissionsBoundaryAttachment_forceNew(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permissions_boundary_attachment.test"
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:    testAccCheckPermissionsBoundaryAttachmentDestroy(ctx),
func
Config: testAccPermissionsBoundaryAttachmentConfig_basic(rName, rNamePolicy1, rNamePolicy2),
Check: resource.ComposeTestCheck
func(
	testAccCheckPermissionsBoundaryAttachmentExists(ctx, resourceName),
),
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckPermissionsBoundaryAttachmentExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "permissions_boundary.0.customer_managed_policy_reference.0.name", rNamePolicy2),
	resource.TestCheckResourceAttrPair(resourceName, "instance_arn", permissionSetResourceName, "instance_arn"),
	resource.TestCheckResourceAttrPair(resourceName, "permission_set_arn", permissionSetResourceName, "arn"),
),
func
ResourceName:      resourceName,
ImportState:       true,
ImportStateVerify: true,
	},
},
	})
}


func TestAccSSOAdminPermissionsBoundaryAttachment_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permissions_boundary_attachment.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:    testAccCheckPermissionsBoundaryAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionsBoundaryAttachmentConfig_basic(rName, rNamePolicy1, rNamePolicy2),
Check: resource.ComposeTestCheck
functAccCheckPermissionsBoundaryAttachmentExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfssoadmin.ResourcePermissionsBoundaryAttachment(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}
func
func TestAccSSOAdminPermissionsBoundaryAttachment_Disappears_permissionSet(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ssoadmin_permissions_boundary_attachment.test"
	permissionSetResourceName := "aws_ssoadmin_permission_set.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
funcrCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:    testAccCheckPermissionsBoundaryAttachmentDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccPermissionsBoundaryAttachmentConfig_basic(rName, rNamePolicy1, rNamePolicy2),
Check: resource.ComposeTestCheck
func(
	testAccCheckPermissionsBoundaryAttachmentExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfssoadmin.ResourcePermissionSet(), permissionSetResourceName),
funcctNonEmptyPlan: true,
	},
},
	})
}


func TestAccSSOAdminPermissionsBoundaryAttachment_managedPolicyAndCustomerManagedPolicyRefBothDefined(t *testing.T) {
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy1 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rNamePolicy2 := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckInstances(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ssoadmin.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:    testAccCheckPermissionsBoundaryAttachmentDestroy(ctx),
Steps: []resource.TestStep{
funcig:      testAccPermissionsBoundaryAttachmentConfig_managedPolicyAndCustomerManagedPolicyRefBothDefined(rName, rNamePolicy1, rNamePolicy2),
ExpectError: regexache.MustCompile(".*ValidationException: Only ManagedPolicyArn or CustomerManagedPolicyReference should be given.*"),
	},
},
	})
}


func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).SSOAdminConn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_ssoadmin_permissions_boundary_attachment" {
continue
	}

	permissionSetARN, instanceARN, err := tfssoadmin.PermissionsBoundaryAttachmentParseResourceID(rs.Primary.ID)
	if err != nil {
return err
	}
funcerr = tfssoadmin.FindPermissionsBoundary(ctx, conn, permissionSetARN, instanceARN)
functfresource.NotFound(err) {
continue
func
	if err != nil {
return err
	}

	return fmt.Errorf("SSO Permissions Boundary Attachment %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccCheckPermissionsBoundaryAttachmentExists(ctx context.Context, n string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

permissionSetARN, instanceARN, err := tfssoadmin.PermissionsBoundaryAttachmentParseResourceID(rs.Primary.ID)
if err != nil {
	return err
}

conn := acctest.Provider.Meta().(*conns.AWSClient).SSOAdminConn(ctx)

_, err = tfssoadmin.FindPermissionsBoundary(ctx, conn, permissionSetARN, instanceARN)

func
func

funcurn fmt.Sprintf(`
data "aws_partition" "current" {}

data "aws_ssoadmin_instances" "test" {}

resource "aws_ssoadmin_permission_set" "test" {
  name= %[1]q
  instance_arn = tolist(data.aws_ssoadmin_instances.test.arns)[0]
}

resource "aws_iam_policy" "test1" {
  name        = %[2]q
  path        = "/"
  description = "My test policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
 "ec2:Describe*",
func    Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "aws_iam_policy" "test2" {
  name        = %[3]q
  path        = "/"
  description = "My test policy"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
 "ec2:Describe*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}
`, rName, rNamePolicy1, rNamePolicy2)
}


func testAccPermissionsBoundaryAttachmentConfig_basic(rName, rNamePolicy1, rNamePolicy2 string) string {
	return acctest.ConfigCompose(testAccPermissionsBoundaryAttachmentConfig_base(rName, rNamePolicy1, rNamePolicy2), `
resource "aws_ssoadmin_permissions_boundary_attachment" "test" {
  instance_arn       = aws_ssoadmin_permission_set.test.instance_arn
  permission_set_arn = aws_ssoadmin_permission_set.test.arn

  permissions_boundary {
    customer_managed_policy_reference {
      name = aws_iam_policy.test1.name
      path = "/"
    }
  }
}
`)
}


func testAccPermissionsBoundaryAttachmentConfig_forceNew(rName, rNamePolicy1, rNamePolicy2 string) string {
	return acctest.ConfigCompose(testAccPermissionsBoundaryAttachmentConfig_base(rName, rNamePolicy1, rNamePolicy2), `
resource "aws_ssoadmin_permissions_boundary_attachment" "test" {
  instance_arn       = aws_ssoadmin_permission_set.test.instance_arn
  permission_set_arn = aws_ssoadmin_permission_set.test.arn
funcrmissions_boundary {
    customer_managed_policy_reference {
      name = aws_iam_policy.test2.name
      path = "/"
    }
  }
}
`)
}


func testAccPermissionsBoundaryAttachmentConfig_managedPolicyAndCustomerManagedPolicyRefBothDefined(rName, rNamePolicy1, rNamePolicy2 string) string {
	return acctest.ConfigCompose(testAccPermissionsBoundaryAttachmentConfig_base(rName, rNamePolicy1, rNamePolicy2), `
data "aws_partition" "partition" {}

resource "aws_ssoadmin_permissions_boundary_attachment" "test" {
  instance_arn       = aws_ssoadmin_permission_set.test.instance_arn
func
  permissions_boundary {
    managed_policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/ReadOnlyAccess"
    customer_managed_policy_reference {
      name = aws_iam_policy.test1.name
      path = "/"
    }
  }
}
`)
}
func