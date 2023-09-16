// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

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
	resourceName := "aws_ami_launch_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:unchPermissionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMILaunchPermissionConfig_accountID(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMILaunchPermissionExists(ctx, resourceName),
funcource.TestCheckResourceAttr(resourceName, "group", ""),
	resource.TestCheckResourceAttr(resourceName, "organization_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "organizational_unit_arn", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccAMILaunchPermissionImportStateId
func(resourceName),
ImportStateVerify: true,
func
func


func TestAccEC2AMILaunchPermission_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ami_launch_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:unchPermissionDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckAMILaunchPermissionExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceAMILaunchPermission(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func TestAccEC2AMILaunchPermission_Disappears_ami(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ami_launch_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:unchPermissionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMILaunchPermissionConfig_accountID(rName),
Check: resource.ComposeTestCheck
func(
functest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceAMICopy(), "aws_ami_copy.test"),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}

func TestAccEC2AMILaunchPermission_group(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ami_launch_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:unchPermissionDestroy(ctx),
func
Config: testAccAMILaunchPermissionConfig_group(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMILaunchPermissionExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "account_id", ""),
	resource.TestCheckResourceAttr(resourceName, "group", "all"),
funcource.TestCheckResourceAttr(resourceName, "organizational_unit_arn", ""),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccAMILaunchPermissionImportStateId
funcrtStateVerify: true,
	},
},
	})
}


func TestAccEC2AMILaunchPermission_organizationARN(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ami_launch_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
funcrCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:unchPermissionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccAMILaunchPermissionConfig_organizationARN(rName),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "account_id", ""),
	resource.TestCheckResourceAttr(resourceName, "group", ""),
	resource.TestCheckResourceAttrSet(resourceName, "organization_arn"),
	resource.TestCheckResourceAttr(resourceName, "organizational_unit_arn", ""),
),
	},
	{
funcrtState:
ImportStateId
func: testAccAMILaunchPermissionImportStateId
func(resourceName),
ImportStateVerify: true,
	},
},
	})
func

func TestAccEC2AMILaunchPermission_organizationalUnitARN(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ami_launch_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t); acctest.PreCheckOrganizationsAccount(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
funcig: testAccAMILaunchPermissionConfig_organizationalUnitARN(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckAMILaunchPermissionExists(ctx, resourceName),
	resource.TestCheckResourceAttr(resourceName, "account_id", ""),
	resource.TestCheckResourceAttr(resourceName, "group", ""),
	resource.TestCheckResourceAttr(resourceName, "organization_arn", ""),
	resource.TestCheckResourceAttrSet(resourceName, "organizational_unit_arn"),
func
	{
ResourceName:ame,
ImportState:
ImportStateId
func: testAccAMILaunchPermissionImportStateId
func(resourceName),
func
},
	})
}


func testAccAMILaunchPermissionImportStateId
func(resourceName string) resource.ImportStateId
funcurn 
func(s *terraform.State) (string, error) {
rs, ok := s.RootModule().Resources[resourceName]
if !ok {
	return "", fmt.Errorf("Not found: %s", resourceName)
}

imageID := rs.Primary.Attributes["image_id"]

if v := rs.Primary.Attributes["group"]; v != "" {
	return fmt.Sprintf("%s/%s", v, imageID), nil
} else if v := rs.Primary.Attributes["organization_arn"]; v != "" {
funcse if v := rs.Primary.Attributes["organizational_unit_arn"]; v != "" {
funcse {
	return fmt.Sprintf("%s/%s", rs.Primary.Attributes["account_id"], imageID), nil
}
	}
}


func testAccCheckAMILaunchPermissionExists(ctx context.Context, n string) resource.TestCheck
funcurn 
funcok := s.RootModule().Resources[n]
funcurn fmt.Errorf("Not found: %s", n)
}
funcs.Primary.ID == "" {
	return fmt.Errorf("No AMI Launch Permission ID is set")
}

imageID, accountID, group, organizationARN, organizationalUnitARN, err := tfec2.AMILaunchPermissionParseResourceID(rs.Primary.ID)

if err != nil {
	return err
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

_, err = tfec2.FindImageLaunchPermission(ctx, conn, imageID, accountID, group, organizationARN, organizationalUnitARN)

return err
	}
}


func testAccCheckAMILaunchPermissionDestroy(ctx context.Context) resource.TestCheck
func {
func(s *terraform.State) error {
func
for _, rs := range s.RootModule().Resources {
funcinue
	}

	imageID, accountID, group, organizationARN, organizationalUnitARN, err := tfec2.AMILaunchPermissionParseResourceID(rs.Primary.ID)

	if err != nil {
return err
	}

	_, err = tfec2.FindImageLaunchPermission(ctx, conn, imageID, accountID, group, organizationARN, organizationalUnitARN)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("AMI Launch Permission %s still exists", rs.Primary.ID)
}

return nil
	}
}
func
funcurn acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`
data "aws_caller_identity" "current" {}
func "aws_region" "current" {}

resource "aws_ami_copy" "test" {
description
name
source_ami_id_ami.amzn-ami-minimal-hvm-ebs.id
source_ami_region = data.aws_region.current.name
}

resource "aws_ami_launch_permission" "test" {
account_id = data.aws_caller_identity.current.account_id
image_idws_ami_copy.test.id
}
`, rName))
}


func testAccAMILaunchPermissionConfig_group(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`
data "aws_region" "current" {}

resource "aws_ami_copy" "test" {
description
name
source_ami_id_ami.amzn-ami-minimal-hvm-ebs.id
source_ami_region = data.aws_region.current.name
deprecation_time= data.aws_ami.amzn-ami-minimal-hvm-ebs.deprecation_time
}

resource "aws_ami_launch_permission" "test" {
group"all"
image_id = aws_ami_copy.test.id
funcName))
}


func testAccAMILaunchPermissionConfig_organizationARN(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`
data "aws_organizations_organization" "current" {}

data "aws_region" "current" {}

resource "aws_ami_copy" "test" {
description
name
source_ami_id_ami.amzn-ami-minimal-hvm-ebs.id
source_ami_region = data.aws_region.current.name
}

resource "aws_ami_launch_permission" "test" {
organization_arn = data.aws_organizations_organization.current.arn
image_idami_copy.test.id
}
func


func testAccAMILaunchPermissionConfig_organizationalUnitARN(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigLatestAmazonLinuxHVMEBSAMI(), fmt.Sprintf(`
resource "aws_organizations_organization" "test" {}

resource "aws_organizations_organizational_unit" "test" {
name
parent_id = aws_organizations_organization.test.roots[0].id
}

data "aws_region" "current" {}

resource "aws_ami_copy" "test" {
description
name
source_ami_id_ami.amzn-ami-minimal-hvm-ebs.id
source_ami_region = data.aws_region.current.name
}
funcurce "aws_ami_launch_permission" "test" {
organizational_unit_arn = aws_organizations_organizational_unit.test.arn
image_id = aws_ami_copy.test.id
}
`, rName))
}
func