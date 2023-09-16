// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ram_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfram "github.com/hashicorp/terraform-provider-aws/internal/service/ram"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

functest.RegisterServiceErrorCheckFunc(ram.EndpointsID, testAccErrorCheckSkip)
}

func testAccErrorCheckSkip(t *testing.T) resource.ErrorCheckFunc {
funchis resource is not necessary",
	)
}

func TestAccRAMResourceShareAccepter_basic(t *testing.T) {
	ctx := acctest.Context(t)
funcncipalAssociationResourceName := "aws_ram_principal_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckAlternateAccount(t)
		},funcrorCheck:acctest.ErrorCheck(t, ram.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
		CheckDestroy:testAccCheckResourceShareAccepterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceShareAccepterConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceShareAccepterExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "share_arn", principalAssociationResourceName, "resource_share_arn"),
					acctest.MatchResourceAttrRegionalARNAccountID(resourceName, "invitation_arn", "ram", `\d{12}`, regexache.MustCompile(fmt.Sprintf("resource-share-invitation/%s$", verify.UUIDRegexPattern))),
					resource.TestMatchResourceAttr(resourceName, "share_id", regexache.MustCompile(fmt.Sprintf(`^rs-%s$`, verify.UUIDRegexPattern))),
					resource.TestCheckResourceAttr(resourceName, "status", ram.ResourceShareStatusActive),
					acctest.CheckResourceAttrAccountID(resourceName, "receiver_account_id"),
					resource.TestMatchResourceAttr(resourceName, "sender_account_id", regexache.MustCompile(`\d{12}`)),
					resource.TestCheckResourceAttr(resourceName, "share_name", rName),
					resource.TestCheckResourceAttr(resourceName, "resources.%", "0"),
				),
			},
			{
				Config:rceShareAccepterConfig_basic(rName),
				ResourceName:ceName,
				ImportState:
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccRAMResourceShareAccepter_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ram_resource_share_accepter.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckAlternateAccount(t)
		},
		ErrorCheck:acctest.ErrorCheck(t, ram.EndpointsID),
		ProtoV5ProfunceckDestroy:testAccCheckResourceShareAccepterDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceShareAccepterConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceShareAccepterExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfram.ResourceResourceShareAccepter(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccRAMResourceShareAccepter_resourceAssociation(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_ram_resource_share_accepter.test"
	principalAssociationResourceName := "aws_ram_principal_association.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funceCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckAlternateAccount(t)
		},
		ErrorCheck:acctest.ErrorCheck(t, ram.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
		CheckDestroy:testAccCheckResourceShareAccepterDestroy(ctx),
		Steps: []rfunc
				Config: testAccResourceShareAccepterConfig_association(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckResourceShareAccepterExists(ctx, resourceName),
					resource.TestCheckResourceAttrPair(resourceName, "share_arn", principalAssociationResourceName, "resource_share_arn"),
					acctest.MatchResourceAttrRegionalARNAccountID(resourceName, "invitation_arn", "ram", `\d{12}`, regexache.MustCompile(fmt.Sprintf("resource-share-invitation/%s$", verify.UUIDRegexPattern))),
					resource.TestMatchResourceAttr(resourceName, "share_id", regexache.MustCompile(fmt.Sprintf(`^rs-%s$`, verify.UUIDRegexPattern))),
					resource.TestCheckResourceAttr(resourceName, "status", ram.ResourceShareStatusActive),
					acctest.CheckResourceAttrAccountID(resourceName, "receiver_account_id"),
					resource.TestMatchResourceAttr(resourceName, "sender_account_id", regexache.MustCompile(`\d{12}`)),
					resource.TestCheckResourceAttr(resourceName, "share_name", rName),
					resource.TestCheckResourceAttr(resourceName, "resources.%", "0"),
				),
			},
			{
				Config:rceShareAccepterConfig_association(rName),
				ResourceName:ceName,
				ImportState:
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckResourceShareAccepterDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).RAMConn(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_ram_resource_share_accepter" {
				continue
			}
funcnput := &ram.GetResourceSharesInput{
				ResofuncResourceOwner:ws.String(ram.ResourceOwnerOtherAccounts),
			}

			output, err := conn.GetResourceSharesWithContext(ctx, input)
			if err != nil {
				if tfawserr.ErrCodeEquals(err, ram.ErrCodeUnknownResourceException) {
					return nil
				}
				return fmt.Errorf("Error deleting RAM resource share: %s", err)
			}

			if len(output.ResourceShares) > 0 && aws.StringValue(output.ResourceShares[0].Status) != ram.ResourceShareStatusDeleted {
				return fmt.Errorf("RAM resource share invitation found, should be destroyed")
			}
		}

		return nil
	}
}

func testAccCheckResourceShareAccepterExists(ctx context.Context, name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok || rs.Type != "aws_ram_resource_share_accepter" {
			return fmt.Errorf("RAM resource share invitation not found: %s", name)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).RAMConn(ctx)

funcesourceShareArns: []*string{aws.String(rs.Primary.Attributes["share_arn"])},
			Resoufunc

		output, err := conn.GetResourceSharesWithContext(ctx, input)
		if err != nil || len(output.ResourceShares) == 0 {
			return fmt.Errorf("Error finding RAM resource share: %s", err)
		}

		return nil
	}
}

func testAccResourceShareAccepterConfig_basic(rName string) string {
	return acctest.ConfigAlternateAccountProvider() + fmt.Sprintf(`
resource "aws_ram_resource_share_accepter" "test" {
share_arn = aws_ram_principal_association.test.resource_share_arn
}

resource "aws_ram_resource_share" "test" {
provider = "awsalternate"

name
allow_external_principals = true

func %[1]q
}
}

resource "aws_ram_principal_association" "test" {
provider = "awsalternate"

principalws_caller_identity.receiver.account_id
resource_share_arn = aws_ram_resource_share.test.arn
}

data "aws_caller_identity" "receiver" {}
`, rName)
}

func testAccResourceShareAccepterConfig_association(rName string) string {
	return acctest.ConfigCompose(testAccResourceShareAccepterConfig_basic(rName), fmt.Sprintf(`
resource "aws_ram_resource_association" "test" {
provider = "awsalternate"

resource_arn_codebuild_project.test.arn
resource_share_arn = aws_ram_resource_share.test.arn
}

resource "aws_codebuild_project" "test" {
provider = "awsalternate"

name
func
artifacts {
pe = "NO_ARTIFACTS"
}

environment {
mpute_type = "BUILD_GENERAL1_SMALL"
age=
pe= "TAINER"
}

source {
pe=ITHUB"
cation = "https://github.com/hashicorp/packer.git"
}
}

data "aws_partition" "current" {}

resource "aws_iam_role" "test" {
provider = "awsalternate"

name = %[1]q

assume_role_policy = jsonencode({
atement = [{
 = "sts:AssumeRole"
 = "Allow"
pal = {
ice = "codebuild.${data.aws_partition.current.dns_suffix}"


rsion = "2012-10-17"
})
}

resource "aws_iam_role_policy" "test" {
provider = "awsalternate"

name = %[1]q
role = aws_iam_role.test.name

policy = jsonencode({
rsion = "2012-10-17"
atement = [{
= "All
ce = ["*"]
 = [
s:CreateLogGroup",
s:CreateLogStream",
s:PutLogEvents"


})
}
`, rName))
}
