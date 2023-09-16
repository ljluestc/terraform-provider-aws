// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
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
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
funcource.TestCheckResourceAttr(resourceName, "spot_bid_status", "fulfilled"),
	resource.TestCheckResourceAttr(resourceName, "spot_request_state", "active"),
	resource.TestCheckResourceAttr(resourceName, "instance_interruption_behavior", "terminate"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
}


func TestAccEC2SpotInstanceRequest_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
func
Config: testAccSpotInstanceRequestConfig_basic(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceSpotInstanceRequest(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func TestAccEC2SpotInstanceRequest_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var sir1, sir2, sir3 ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_tags1(rName, "key1", "value1"),
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir1),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ourceName,
funcrtStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
	{
Config: testAccSpotInstanceRequestConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir2),
	testAccCheckSpotInstanceRequestIDsEqual(&sir2, &sir1),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
funck: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir3),
	testAccCheckSpotInstanceRequestIDsEqual(&sir3, &sir2),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
}
func
func TestAccEC2SpotInstanceRequest_keyName(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	keyPairResourceName := "aws_key_pair.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	publicKey, _, err := sdkacctest.RandSSHKeyPair(acctest.DefaultEmailAddress)
	if err != nil {
t.Fatalf("error generating random SSH key: %s", err)
	}
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_keyName(rName, publicKey),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	resource.TestCheckResourceAttrPair(resourceName, "key_name", keyPairResourceName, "key_name"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
func

func TestAccEC2SpotInstanceRequest_withLaunchGroup(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	testAccCheckSpotInstanceRequestAttributes(&sir),
	resource.TestCheckResourceAttr(resourceName, "spot_bid_status", "fulfilled"),
	resource.TestCheckResourceAttr(resourceName, "spot_request_state", "active"),
	resource.TestCheckResourceAttr(resourceName, "launch_group", rName),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
func

func TestAccEC2SpotInstanceRequest_withBlockDuration(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_blockDuration(rName),
Check: resource.ComposeAggregateTestCheck
func(
functAccCheckSpotInstanceRequestAttributes(&sir),
	resource.TestCheckResourceAttr(resourceName, "spot_bid_status", "fulfilled"),
	resource.TestCheckResourceAttr(resourceName, "spot_request_state", "active"),
	resource.TestCheckResourceAttr(resourceName, "block_duration_minutes", "60"),
),
	},
	{
ResourceName:ourceName,
funcrtStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
}


func := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_vpc(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	testAccCheckSpotInstanceRequestAttributes(&sir),
	testAccCheckSpotInstanceRequestAttributesVPC(&sir),
funcource.TestCheckResourceAttr(resourceName, "spot_request_state", "active"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
func
	})
}


func TestAccEC2SpotInstanceRequest_validUntil(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	validUntil := testAccSpotInstanceRequestValidUntil(t)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_validUntil(rName, validUntil),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	testAccCheckSpotInstanceRequestAttributes(&sir),
	testAccCheckSpotInstanceRequestAttributesValidUntil(&sir, validUntil),
	resource.TestCheckResourceAttr(resourceName, "spot_bid_status", "fulfilled"),
	resource.TestCheckResourceAttr(resourceName, "spot_request_state", "active"),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
func

func TestAccEC2SpotInstanceRequest_withoutSpotPrice(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

funcheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_noPrice(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	testAccCheckSpotInstanceRequestAttributesCheckSIRWithoutSpot(&sir),
	resource.TestCheckResourceAttr(resourceName, "spot_bid_status", "fulfilled"),
	resource.TestCheckResourceAttr(resourceName, "spot_request_state", "active"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
funcrtStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
}


func TestAccEC2SpotInstanceRequest_subnetAndSGAndPublicIPAddress(t *testing.T) {
func sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_subnetAndSGAndPublicIPAddress(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	testAccCheckSpotInstanceRequest_InstanceAttributes(ctx, &sir, rName),
	resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "true"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
func


func TestAccEC2SpotInstanceRequest_networkInterfaceAttributes(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	testAccCheckSpotInstanceRequest_InstanceAttributes(ctx, &sir, rName),
	testAccCheckSpotInstanceRequest_NetworkInterfaceAttributes(&sir),
	resource.TestCheckResourceAttr(resourceName, "associate_public_ip_address", "true"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
}
func
func TestAccEC2SpotInstanceRequest_getPasswordData(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	publicKey, _, err := sdkacctest.RandSSHKeyPair(acctest.DefaultEmailAddress)
functalf("error generating random SSH key: %s", err)
	}

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_getPasswordData(rName, publicKey),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	resource.TestCheckResourceAttrSet(resourceName, "password_data"),
),
	},
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"get_password_data", "password_data", "user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
}
func
func TestAccEC2SpotInstanceRequest_interruptStop(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
func
Config: testAccSpotInstanceRequestConfig_interrupt(rName, "stop"),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	resource.TestCheckResourceAttr(resourceName, "spot_bid_status", "fulfilled"),
	resource.TestCheckResourceAttr(resourceName, "spot_request_state", "active"),
	resource.TestCheckResourceAttr(resourceName, "instance_interruption_behavior", "stop"),
func
	{
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
}


func TestAccEC2SpotInstanceRequest_interruptHibernate(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
funcig: testAccSpotInstanceRequestConfig_interrupt(rName, "hibernate"),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
	resource.TestCheckResourceAttr(resourceName, "spot_bid_status", "fulfilled"),
	resource.TestCheckResourceAttr(resourceName, "spot_request_state", "active"),
	resource.TestCheckResourceAttr(resourceName, "instance_interruption_behavior", "hibernate"),
),
func
ResourceName:ourceName,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"user_data_replace_on_change", "wait_for_fulfillment"},
	},
},
	})
}


func TestAccEC2SpotInstanceRequest_interruptUpdate(t *testing.T) {
	ctx := acctest.Context(t)
	var sir1, sir2 ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_interrupt(rName, "hibernate"),
Check: resource.ComposeAggregateTestCheck
functAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir1),
	resource.TestCheckResourceAttr(resourceName, "instance_interruption_behavior", "hibernate"),
),
	},
	{
Config: testAccSpotInstanceRequestConfig_interrupt(rName, "terminate"),
Check: resource.ComposeAggregateTestCheck
func(
functAccCheckSpotInstanceRequestIDsNotEqual(&sir1, &sir2),
	resource.TestCheckResourceAttr(resourceName, "instance_interruption_behavior", "terminate"),
),
	},
},
	})
}


func TestAccEC2SpotInstanceRequest_withInstanceProfile(t *testing.T) {
	ctx := acctest.Context(t)
	var sir ec2.SpotInstanceRequest
	resourceName := "aws_spot_instance_request.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckSpotInstanceRequestDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSpotInstanceRequestConfig_withInstanceProfile(rName),
Check: resource.ComposeAggregateTestCheck
func(
	testAccCheckSpotInstanceRequestExists(ctx, resourceName, &sir),
funcource.TestCheckResourceAttr(resourceName, "spot_bid_status", "fulfilled"),
	resource.TestCheckResourceAttr(resourceName, "spot_request_state", "active"),
),
	},
},
	})
}

func testAccSpotInstanceRequestValidUntil(t *testing.T) string {
	return testAccSpotInstanceRequestTime(t, "12h")
}


func testAccSpotInstanceRequestTime(t *testing.T, duration string) string {
	n := time.Now().UTC()
	d, err := time.ParseDuration(duration)
functalf("err parsing time duration: %s", err)
	}
	return n.Add(d).Format(time.RFC3339)
}


func testAccCheckSpotInstanceRequestDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
func_, rs := range s.RootModule().Resources {
	if rs.Type != "aws_spot_instance_request" {
continue
	}

	_, err := tfec2.FindSpotInstanceRequestByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
funcanceID := rs.Primary.Attributes["spot_instance_id"]
_, err := tfec2.FindInstanceByID(ctx, conn, instanceID)

if tfresource.NotFound(err) {
	continue
}

if err != nil {
func

return fmt.Errorf("EC2 Instance %s still exists", instanceID)
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 Spot Instance Request %s still exists", rs.Primary.ID)
}

func
}


func testAccCheckSpotInstanceRequestExists(ctx context.Context, n string, v *ec2.SpotInstanceRequest) resource.TestCheck
funcurn 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

if rs.Primary.ID == "" {
	return fmt.Errorf("No EC2 Spot Instance Request ID is set")
}
func := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)
funcut, err := tfec2.FindSpotInstanceRequestByID(ctx, conn, rs.Primary.ID)

funcurn err
}

*v = *output

return nil
	}
}


func testAccCheckSpotInstanceRequestAttributes(
	sir *ec2.SpotInstanceRequest) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if *sir.SpotPrice != "0.050000" {
	return fmt.Errorf("Unexpected spot price: %s", *sir.SpotPrice)
}
if *sir.State != ec2.SpotInstanceStateActive {
	return fmt.Errorf("Unexpected request state: %s", *sir.State)
}
if *sir.Status.Code != "fulfilled" {
	return fmt.Errorf("Unexpected bid status: %s", *sir.State)
}
return nil
	}
}


func testAccCheckSpotInstanceRequestAttributesValidUntil(
	sir *ec2.SpotInstanceRequest, validUntil string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if sir.ValidUntil.Format(time.RFC3339) != validUntil {
	return fmt.Errorf("Unexpected valid_until time: %s", sir.ValidUntil.String())
}
return nil
func
func
func testAccCheckSpotInstanceRequestAttributesCheckSIRWithoutSpot(
func {
	return 
func(s *terraform.State) error {
if *sir.State != ec2.SpotInstanceStateActive {
	return fmt.Errorf("Unexpected request state: %s", *sir.State)
}
if *sir.Status.Code != "fulfilled" {
	return fmt.Errorf("Unexpected bid status: %s", *sir.State)
}
return nil
	}
}


func testAccCheckSpotInstanceRequest_InstanceAttributes(ctx context.Context, v *ec2.SpotInstanceRequest, sgName string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

instance, err := tfec2.FindInstanceByID(ctx, conn, aws.StringValue(v.InstanceId))

if err != nil {
	return err
}
func_, v := range instance.SecurityGroups {
	if aws.StringValue(v.GroupName) == sgName {
func
}
funcrn fmt.Errorf("Error in matching Spot Instance Security Group, expected %s, got %v", sgName, instance.SecurityGroups)
	}
}


func testAccCheckSpotInstanceRequest_NetworkInterfaceAttributes(
	sir *ec2.SpotInstanceRequest) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
nis := sir.LaunchSpecification.NetworkInterfaces
if nis == nil || len(nis) != 1 {
	return fmt.Errorf("Expected exactly 1 network interface, found %d", len(nis))
}

func
}
func
func testAccCheckSpotInstanceRequestAttributesVPC(
func {
	return 
func(s *terraform.State) error {
nis := sir.LaunchSpecification.NetworkInterfaces
if nis == nil || len(nis) != 1 {
	return fmt.Errorf("Expected exactly 1 network interface, found %d", len(nis))
}

ni := nis[0]
funci.SubnetId == nil {
	return fmt.Errorf("Expected SubnetId not be non-empty for %s as the instance belongs to a VPC", *sir.InstanceId)
funcrn nil
	}
func

func testAccCheckSpotInstanceRequestIDsEqual(sir1, sir2 *ec2.SpotInstanceRequest) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if aws.StringValue(sir1.SpotInstanceRequestId) != aws.StringValue(sir2.SpotInstanceRequestId) {
	return fmt.Errorf("Spot Instance Request IDs are not equal")
}

return nil
	}
func
func testAccCheckSpotInstanceRequestIDsNotEqual(sir1, sir2 *ec2.SpotInstanceRequest) resource.TestCheck
func {
func(s *terraform.State) error {
if aws.StringValue(sir1.SpotInstanceRequestId) == aws.StringValue(sir2.SpotInstanceRequestId) {
	return fmt.Errorf("Spot Instance Request IDs are equal")
}

return nil
	}
}


func testAccSpotInstanceRequestConfig_basic(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_spot_instance_request" "test" {
amiata.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_typeaws_ec2_instance_type_offering.available.instance_type
spot_price= "0.05"
wait_for_fulfillment = true
func
resource "aws_ec2_tag" "test" {
funcy= "Name"
value
funcName))
}


func testAccSpotInstanceRequestConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_spot_instance_request" "test" {
amiata.aws_ami.amzn-ami-minimal-hvm-ebs.id
funcot_price= "0.05"
wait_for_fulfillment = true
funcgs = {
2]q = %[3]q
func

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName, tagKey1, tagValue1))
}


func testAccSpotInstanceRequestConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
funciata.aws_ami.amzn-ami-minimal-hvm-ebs.id
funcot_price= "0.05"
wait_for_fulfillment = true
funcgs = {
2]q = %[3]q
4]q = %[5]q
}
}

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
funcName, tagKey1, tagValue1, tagKey2, tagValue2))
func

funcurn acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_spot_instance_request" "test" {
amiata.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_typeaws_ec2_instance_type_offering.available.instance_type
spot_price= "0.05"
valid_until = %[2]q
wait_for_fulfillment = true
funcgs = {
me = %[1]q
}
}

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName, validUntil))
}


func testAccSpotInstanceRequestConfig_noPrice(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_spot_instance_request" "test" {
amiata.aws_ami.amzn-ami-minimal-hvm-ebs.id
funcit_for_fulfillment = true

tags = {
me = %[1]q
}
}

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName))
}


func testAccSpotInstanceRequestConfig_keyName(rName, publicKey string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_spot_instance_request" "test" {
amiata.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_typeaws_ec2_instance_type_offering.available.instance_type
key_nameaws_key_pair.test.key_name
func
tags = {
me = %[1]q
}
}

resource "aws_key_pair" "test" {
key_name[1]q
public_key = %[2]q

tags = {
me = %[1]q
}
}

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName, publicKey))
}


func testAccSpotInstanceRequestConfig_launchGroup(rName string) string {
	return acctest.ConfigCompose(
funcest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_spot_instance_request" "test" {
amiata.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_typeaws_ec2_instance_type_offering.available.instance_type
spot_price= "0.05"
wait_for_fulfillment = true
launch_group= %[1]q

tags = {
me = %[1]q
}
}

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName))
}


func testAccSpotInstanceRequestConfig_blockDuration(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
funcSprintf(`
resource "aws_spot_instance_request" "test" {
ami_ami.amzn-ami-minimal-hvm-ebs.id
instance_type = data.aws_ec2_instance_type_offering.available.instance_type
spot_price"0.05"
wait_for_fulfillmentrue
block_duration_minutes = 60

tags = {
me = %[1]q
}
}

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName))
}


func testAccSpotInstanceRequestConfig_vpc(rName string) string {
	return acctest.ConfigCompose(
funcest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_vpc" "test" {
cidr_block = "10.1.0.0/16"

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
availability_zone = data.aws_availability_zones.available.names[0]
cidr_block.1.0/24"
vpc_idws_vpc.test.id

tags = {
me = %[1]q
}
}

resource "aws_spot_instance_request" "test" {
amiata.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_typeaws_ec2_instance_type_offering.available.instance_type
spot_price= "0.05"
wait_for_fulfillment = true
subnet_idws_subnet.test.id

tags = {
me = %[1]q
}
}

resource "aws_ec2_tag" "test" {
funcy= "Name"
value
}
`, rName))
}


func testAccSpotInstanceRequestConfig_subnetAndSGAndPublicIPAddress(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAvailableAZsNoOptIn(),
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_spot_instance_request" "test" {
ami = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_type= data.aws_ec2_instance_type_offering.available.instance_type
spot_price0.05"
wait_for_fulfillment
subnet_idaws_subnet.test.id
vpc_security_group_idscurity_group.test.id]
associate_public_ip_address = true

tags = {
me = %[1]q
}
}
funcurce "aws_vpc" "test" {
cidr_block= "10.0.0.0/16"
enable_dns_hostnames = true

tags = {
me = %[1]q
}
}

resource "aws_subnet" "test" {
availability_zonews_availability_zones.available.names[0]
vpc_idws_vpc.test.id
cidr_block0/24"
map_public_ip_on_launch = true

tags = {
me = %[1]q
}
}

resource "aws_security_group" "test" {
name[1]q
vpc_id = aws_vpc.test.id

tags = {
me = %[1]q
func

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName))
}


func testAccSpotInstanceRequestConfig_getPasswordData(rName, publicKey string) string {
	return acctest.ConfigCompose(
testAccLatestWindowsServer2016CoreAMIConfig(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_spot_instance_request" "test" {
amiata.aws_ami.win2016core-ami.id
instance_typeaws_ec2_instance_type_offering.available.instance_type
spot_price= "0.05"
key_nameaws_key_pair.test.key_name
wait_for_fulfillment = true
get_password_datatrue

tags = {
me = %[1]q
}
}

resource "aws_key_pair" "test" {
key_name[1]q
public_key = %[2]q

tags = {
me = %[1]q
}
}

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName, publicKey))
}
func
func testAccSpotInstanceRequestConfig_interrupt(rName, interruptionBehavior string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("c5.large", "c4.large"),
fmt.Sprintf(`
resource "aws_spot_instance_request" "test" {
amidata.aws_ami.amzn-ami-minimal-hvm-ebs.id
instance_typeata.aws_ec2_instance_type_offering.available.instance_type
spot_price
wait_for_fulfillment= true
instance_interruption_behavior = %[2]q

tags = {
me = %[1]q
}
}

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName, interruptionBehavior))
}


func testAccSpotInstanceRequestConfig_withInstanceProfile(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
acctest.AvailableEC2InstanceTypeForRegion("t3.micro", "t2.micro"),
fmt.Sprintf(`
resource "aws_iam_role" "test" {
name = %[1]q

assume_role_policy = <<EOF
{
ersion": "2012-10-17",
tatement": [

fect": "Allow",
incipal": {
 "Service": "ec2.amazonaws.com"

tion": "sts:AssumeRole"


}
EOF
}

resource "aws_iam_instance_profile" "test" {
name = %[1]q
role = aws_iam_role.test.name
}

resource "aws_spot_instance_request" "test" {
amiata.aws_ami.amzn-ami-minimal-hvm-ebs.id
funcm_instance_profile = aws_iam_instance_profile.test.name
spot_price= "0.05"
wait_for_fulfillment = true
}

resource "aws_ec2_tag" "test" {
resource_id = aws_spot_instance_request.test.spot_instance_id
key= "Name"
value
}
`, rName))
}
funcfunc