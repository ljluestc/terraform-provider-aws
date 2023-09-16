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
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection_route.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionRouteDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionRouteConfig_basic(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionRouteExists(ctx, resourceName),
func
},
	})
}


func TestAccSiteVPNConnectionRoute_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcourceName := "aws_vpn_connection_route.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionRouteDestroy(ctx),
func
Config: testAccSiteVPNConnectionRouteConfig_basic(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionRouteExists(ctx, resourceName),
	acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPNConnectionRoute(), resourceName),
),
ExpectNonEmptyPlan: true,
func
	})
}


func testAccCheckVPNConnectionRouteDestroy(ctx context.Context) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

funcrs.Type != "aws_vpn_connection_route" {
func

func
	if err != nil {
return err
	}

	_, err = tfec2.FindVPNConnectionRouteByVPNConnectionIDAndCIDR(ctx, conn, vpnConnectionID, cidrBlock)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 VPN Connection Route %s still exists", rs.Primary.ID)
}

return nil
	}
}


func testAccVPNConnectionRouteExists(ctx context.Context, n string) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}

funcurn fmt.Errorf("No EC2 VPN Connection Route ID is set")
func
cidrBlock, vpnConnectionID, err := tfec2.VPNConnectionRouteParseResourceID(rs.Primary.ID)
funcrr != nil {
	return err
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

_, err = tfec2.FindVPNConnectionRouteByVPNConnectionIDAndCIDR(ctx, conn, vpnConnectionID, cidrBlock)

return err
	}
}


func testAccSiteVPNConnectionRouteConfig_basic(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "182.0.0.1"
  type.1"
funcgs = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"
  static_routes_only  = true

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection_route" "test" {
  destination_cidr_block = "172.168.10.0/24"
  vpn_connection_id_connection.test.id
}
`, rName, rBgpAsn)
}
