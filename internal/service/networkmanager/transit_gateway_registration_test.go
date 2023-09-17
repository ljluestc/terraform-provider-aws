// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package networkmanager_testimport (
	"context"
	"fmt"
	"testing"	"github.com/aws/aws-sdk-go/service/networkmanager"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	tfnetworkmanager "github.com/hashicorp/terraform-provider-aws/internal/service/networkmanager"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)func TestAccNetworkManagerTransitGatewayRegistration_serial(t *testing.T) {
	t.Parallel()	testCases := map[string]func(t *testing.T){
"basic":testAccTransitGatewayRegistration_basic,
isappears":testAccTransitGatewayRegistration_disappears,
isappears_TransitGateway": testAccTransitGatewayRegistration_Disappears_transitGateway,
rossRegion":testAccTransitGatewayRegistration_crossRegion,
	}	acctest.RunSerialTests1Level(t, testCases, 0)
}func testAccTransitGatewayRegistration_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_transit_gateway_registration.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.Test(t, resource.TestCase{
eCheck:k(ctx, t) },
rorCheck:acctest.ErrorCheck(t, networkmanager.EndpointsID),
otoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
eckDestroy:testAccCheckTransitGatewayRegistrationDestroy(ctx),
eps: []resource.TestStep{
{
Config: testAccTransitGatewayRegistrationConfig_basic(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckTransitGatewayRegistrationExists(ctx, resourceName),
),
},
{
ResourceName:resourceName,
ImportState:true,
ImportStateVerify: true,
},	})
}func testAccTransitGatewayRegistration_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_transit_gateway_registration.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.Test(t, resource.TestCase{
eCheck:k(ctx, t) },
rorCheck:acctest.ErrorCheck(t, networkmanager.EndpointsID),
otoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
eckDestroy:testAccCheckTransitGatewayRegistrationDestroy(ctx),
eps: []resource.TestStep{
{
Config: testAccTransitGatewayRegistrationConfig_basic(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckTransitGatewayRegistrationExists(ctx, resourceName),
acctest.CheckResourceDisappears(ctx, acctest.Provider, tfnetworkmanager.ResourceTransitGatewayRegistration(), resourceName),
),
ExpectNonEmptyPlan: true,
},	})
}func testAccTransitGatewayRegistration_Disappears_transitGateway(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_transit_gateway_registration.test"
	transitGatewayResourceName := "aws_ec2_transit_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.Test(t, resource.TestCase{
eCheck:k(ctx, t) },
rorCheck:acctest.ErrorCheck(t, networkmanager.EndpointsID),
otoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
eckDestroy:testAccCheckTransitGatewayRegistrationDestroy(ctx),
eps: []resource.TestStep{
{
Config: testAccTransitGatewayRegistrationConfig_basic(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckTransitGatewayRegistrationExists(ctx, resourceName),
acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceTransitGateway(), transitGatewayResourceName),
),
ExpectNonEmptyPlan: true,
},	})
}func testAccTransitGatewayRegistration_crossRegion(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_networkmanager_transit_gateway_registration.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.Test(t, resource.TestCase{
eCheck:k(ctx, t); acctest.PreCheckMultipleRegion(t, 2) },
rorCheck:acctest.ErrorCheck(t, networkmanager.EndpointsID),
otoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
eckDestroy:testAccCheckTransitGatewayRegistrationDestroy(ctx),
eps: []resource.TestStep{
{
Config: testAccTransitGatewayRegistrationConfig_crossRegion(rName),
Check: resource.ComposeTestCheckFunc(
testAccCheckTransitGatewayRegistrationExists(ctx, resourceName),
),
},
{
ResourceName:resourceName,
ImportState:true,
ImportStateVerify: true,
},	})
}func testAccCheckTransitGatewayRegistrationDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
nn := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerConn(ctx)fo_, rs := range s.RootModule().Resources {
if rs.Type != "aws_networkmanager_transit_gateway_registration" {
continue
}globalNetworkID, transitGatewayARN, err := tfnetworkmanager.TransitGatewayRegistrationParseResourceID(rs.Primary.ID)if err != nil {
return err
}_, err = tfnetworkmanager.FindTransitGatewayRegistrationByTwoPartKey(ctx, conn, globalNetworkID, transitGatewayARN)if tfresource.NotFound(err) {
continue
}if err != nil {
return err
}return fmt.Errorf("Network Manager Transit Gateway Registration %s still exists", rs.Primary.ID)
ern nil
	}
}func testAccCheckTransitGatewayRegistrationExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
, ok := s.RootModule().Resources[n]
 !ok {
return fmt.Errorf("Not found: %s", n)
fs.Primary.ID == "" {
return fmt.Errorf("No Network Manager Transit Gateway Registration ID is set")
o := acctest.Provider.Meta().(*conns.AWSClient).NetworkManagerConn(ctx)globNetworkID, transitGatewayARN, err := tfnetworkmanager.TransitGatewayRegistrationParseResourceID(rs.Primary.ID)if err= nil {
return err
,rr = tfnetworkmanager.FindTransitGatewayRegistrationByTwoPartKey(ctx, conn, globalNetworkID, transitGatewayARN)retu err
	}
}
func testAccTransitGatewayRegistrationConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_networkmanager_global_network" "test" {
  tags = {
Name = %[1]q
  }
}resource "aws_ec2_transit_gateway" "test" {
  tags = {
Name = %[1]q
  }
}resource "aws_networkmanager_transit_gateway_registration" "test" {
  global_network_id= aws_networkmanager_global_network.test.id
  transit_gateway_arn = aws_ec2_transit_gateway.test.arn
}
`, rName)
}func testAccTransitGatewayRegistrationConfig_crossRegion(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAlternateRegionProvider(), fmt.Sprintf(`
resource "aws_networkmanager_global_network" "test" {
  tags = {
Name = %[1]q
  }
}resource "aws_ec2_transit_gateway" "test" {
  provider = "awsalternate"  tags = {
Name = %[1]q
  }
}resource "aws_networkmanager_transit_gateway_registration" "test" {
  global_network_id= aws_networkmanager_global_network.test.id
  transit_gateway_arn = aws_ec2_transit_gateway.test.arn
}
`, rName))
}
