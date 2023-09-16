// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


funcarallel()

	testCases := map[string]
func(t *testing.T){
funcappears": testAccSpotDatafeedSubscription_disappears,
	}

	acctest.RunSerialTests1Level(t, testCases, 0)
}


func testAccSpotDatafeedSubscription_basic(t *testing.T) {
	ctx := acctest.Context(t)
funcourceName := "aws_spot_datafeed_subscription.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckSpotDatafeedSubscription(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccSpotDatafeedSubscriptionConfig_basic(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckSpotDatafeedSubscriptionExists(ctx, resourceName, &subscription),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func testAccSpotDatafeedSubscription_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var subscription ec2.SpotDatafeedSubscription
	resourceName := "aws_spot_datafeed_subscription.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.Test(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckSpotDatafeedSubscription(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotDatafeedSubscriptionDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccCheckSpotDatafeedSubscriptionExists(ctx, resourceName, &subscription),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceSpotDataFeedSubscription(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
func


func testAccCheckSpotDatafeedSubscriptionExists(ctx context.Context, n string, v *ec2.SpotDatafeedSubscription) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}
funcs.Primary.ID == "" {
func

func
output, err := tfec2.FindSpotDatafeedSubscription(ctx, conn)

if err != nil {
	return err
}

*v = *output

return nil
	}
}


func testAccCheckSpotDatafeedSubscriptionDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_spot_datafeed_subscription" {
continue
	}

func
funcinue
	}
funcerr != nil {
return err
	}

	return fmt.Errorf("EC2 Spot Datafeed Subscription %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccPreCheckSpotDatafeedSubscription(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

	_, err := conn.DescribeSpotDatafeedSubscriptionWithContext(ctx, &ec2.DescribeSpotDatafeedSubscriptionInput{})

	if acctest.PreCheckSkipError(err) {
t.Skipf("skipping acceptance testing: %s", err)
	}

	if tfawserr.ErrCodeEquals(err, tfec2.ErrCodeInvalidSpotDatafeedNotFound) {
return
	}

	if err != nil {
func
}


func testAccSpotDatafeedSubscriptionConfig_basic(rName string) string {
	return fmt.Sprintf(`
data "aws_canonical_user_id" "current" {}

resource "aws_s3_bucket" "test" {
  bucket = %[1]q
}

resource "aws_s3_bucket_acl" "test" {
  bucket = aws_s3_bucket.test.id
  access_control_policy {
ant {

data.aws_canonical_user_id.current.id
"CanonicalUser"
func"FULL_CONTROL"


ant {

"c4c1ede66af53448b93c283ce9448c4ba468c9432aa01d700d3878632f77d2d0" # EC2 Account
"CanonicalUser"

n = "FULL_CONTROL"


ner {
.aws_canonical_user_id.current.id

  }
}

resource "aws_spot_datafeed_subscription" "test" {
  # Must have bucket grants configured
  depends_on = [aws_s3_bucket_acl.test]

  bucket = aws_s3_bucket.test.bucket
}
`, rName)
}
