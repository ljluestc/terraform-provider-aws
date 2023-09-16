// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
)


func := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_peering_connection_options.test"
	pcxResourceName := "aws_vpc_peering_connection.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCPeeringConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCPeeringConnectionOptionsConfig_sameRegionSameAccount(rName, true),
Check: resource.ComposeAggregateTestCheck
func(
	// Requester's view:
funcource.TestCheckResourceAttr(resourceName, "requester.0.allow_remote_vpc_dns_resolution", "false"),
	testAccCheckVPCPeeringConnectionOptions(ctx, pcxResourceName,
"requester",
&ec2.VpcPeeringConnectionOptionsDescription{
	AllowDnsResolutionFromRemoteVpc: aws.Bool(false),
},
	),
	// Accepter's view:
	resource.TestCheckResourceAttr(resourceName, "accepter.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "accepter.0.allow_remote_vpc_dns_resolution", "true"),
	testAccCheckVPCPeeringConnectionOptions(ctx, pcxResourceName,
"accepter",
&ec2.VpcPeeringConnectionOptionsDescription{
	AllowDnsResolutionFromRemoteVpc: aws.Bool(true),
},
	),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccVPCPeeringConnectionOptionsConfig_sameRegionSameAccount(rName, false),
Check: resource.ComposeAggregateTestCheck
func(
	// Requester's view:
	resource.TestCheckResourceAttr(
funcuester.#",
"1",
	),
	testAccCheckVPCPeeringConnectionOptions(ctx, pcxResourceName,
"requester",
&ec2.VpcPeeringConnectionOptionsDescription{
	AllowDnsResolutionFromRemoteVpc: aws.Bool(false),
},
	),
	// Accepter's view:
	resource.TestCheckResourceAttr(
resourceName,
"accepter.#",
"1",
	),
	testAccCheckVPCPeeringConnectionOptions(ctx, pcxResourceName,
"accepter",
&ec2.VpcPeeringConnectionOptionsDescription{
	AllowDnsResolutionFromRemoteVpc: aws.Bool(false),
},
	),
),
	},
},
	})
}


func TestAccVPCPeeringConnectionOptions_differentRegionSameAccount(t *testing.T) {
	ctx := acctest.Context(t)
	var providers []*schema.Provider
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcourceNamePeer := "aws_vpc_peering_connection_options.peer"r
	pcxResourceName := "aws_vpc_peering_connection.test"er
	pcxResourceNamePeer := "aws_vpc_peering_connection_accepter.peer" // Accepter

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckMultipleRegion(t, 2)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckVPCPeeringConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCPeeringConnectionOptionsConfig_differentRegionSameAccount(rName, true, true),
Check: resource.ComposeAggregateTestCheck
func(
	// Requester's view:
	resource.TestCheckResourceAttr(resourceName, "requester.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "requester.0.allow_remote_vpc_dns_resolution", "true"),
	testAccCheckVPCPeeringConnectionOptions(ctx, pcxResourceName,
"requester",
funcowDnsResolutionFromRemoteVpc: aws.Bool(true),
},
	),
	// Accepter's view:
	resource.TestCheckResourceAttr(resourceNamePeer, "accepter.#", "1"),
	resource.TestCheckResourceAttr(resourceNamePeer, "accepter.0.allow_remote_vpc_dns_resolution", "true"),
	testAccCheckVPCPeeringConnectionOptionsWithProvider(ctx, pcxResourceNamePeer,
"accepter",
&ec2.VpcPeeringConnectionOptionsDescription{
	AllowDnsResolutionFromRemoteVpc: aws.Bool(true),
},
acctest.RegionProvider
func(acctest.AlternateRegion(), &providers),
	),
),
	},
	{
Config:tAccVPCPeeringConnectionOptionsConfig_differentRegionSameAccount(rName, true, true),
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccVPCPeeringConnectionOptionsConfig_differentRegionSameAccount(rName, false, false),
Check: resource.ComposeAggregateTestCheck
func(
	// Requester's view:
	resource.TestCheckResourceAttr(
resourceName,
"requester.#",
"1",
	),
	testAccCheckVPCPeeringConnectionOptions(ctx, pcxResourceName,
func.VpcPeeringConnectionOptionsDescription{
	AllowDnsResolutionFromRemoteVpc: aws.Bool(false),
},
	),
	// Accepter's view:
	resource.TestCheckResourceAttr(
resourceNamePeer,
"accepter.#",
"1",
	),
	testAccCheckVPCPeeringConnectionOptionsWithProvider(ctx, pcxResourceNamePeer,
"accepter",
&ec2.VpcPeeringConnectionOptionsDescription{
	AllowDnsResolutionFromRemoteVpc: aws.Bool(false),
},
acctest.RegionProvider
func(acctest.AlternateRegion(), &providers),
	),
),
	},
},
	})
}


func := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName := "aws_vpc_peering_connection_options.test"er
	resourceNamePeer := "aws_vpc_peering_connection_options.peer" // Accepter
	pcxResourceName := "aws_vpc_peering_connection.test" // Requester

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
functest.PreCheckAlternateAccount(t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
CheckDestroy:stAccCheckVPCPeeringConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccVPCPeeringConnectionOptionsConfig_sameRegionDifferentAccount(rName),
Check: resource.ComposeAggregateTestCheck
funcRequester's view:
	resource.TestCheckResourceAttr(resourceName, "requester.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "requester.0.allow_remote_vpc_dns_resolution", "true"),
	testAccCheckVPCPeeringConnectionOptions(ctx, pcxResourceName,
"requester",
&ec2.VpcPeeringConnectionOptionsDescription{
	AllowDnsResolutionFromRemoteVpc: aws.Bool(true),
},
	),
	// Accepter's view:
	resource.TestCheckResourceAttr(resourceNamePeer, "accepter.#", "1"),
func
	},
	{
Config:tAccVPCPeeringConnectionOptionsConfig_sameRegionDifferentAccount(rName),
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func testAccCheckVPCPeeringConnectionOptions(ctx context.Context, n, block string, options *ec2.VpcPeeringConnectionOptionsDescription) resource.TestCheck
func {
	return testAccCheckVPCPeeringConnectionOptionsWithProvider(ctx, n, block, options, 
func() *schema.Provider { return acctest.Provider })
}


func testAccCheckVPCPeeringConnectionOptionsWithProvider(ctx context.Context, n, block string, options *ec2.VpcPeeringConnectionOptionsDescription, providerF 
func() *schema.Provider) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
funcurn fmt.Errorf("Not found: %s", n)
func
if rs.Primary.ID == "" {
func

conn := providerF().Meta().(*conns.AWSClient).EC2Conn(ctx)

func
funcurn err
func
o := output.AccepterVpcInfo
func output.RequesterVpcInfo
}

if got, want := aws.BoolValue(o.PeeringOptions.AllowDnsResolutionFromRemoteVpc), aws.BoolValue(options.AllowDnsResolutionFromRemoteVpc); got != want {
	return fmt.Errorf("VPC Peering Connection Options AllowDnsResolutionFromRemoteVpc =%v, want = %v", got, want)
}

return nil
	}
}


func testAccVPCPeeringConnectionOptionsConfig_sameRegionSameAccount(rName string, accepterDnsResolution bool) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "peer" {
  cidr_block  = "10.1.0.0/16"
  enable_dns_hostnames = true

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_peering_connection" "test" {
funcer_vpc_id = aws_vpc.peer.id
  auto_accept = true

  tags = {
me = %[1]q
  }
}

resource "aws_vpc_peering_connection_options" "test" {
  vpc_peering_connection_id = aws_vpc_peering_connection.test.id

  accepter {
low_remote_vpc_dns_resolution = %[2]t
  }
}
`, rName, accepterDnsResolution)
}


func testAccVPCPeeringConnectionOptionsConfig_differentRegionSameAccount(rName string, dnsResolution, dnsResolutionPeer bool) string {
	return acctest.ConfigCompose(acctest.ConfigAlternateRegionProvider(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block  = "10.0.0.0/16"
  enable_dns_hostnames = true

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "peer" {
  provider = "awsalternate"

  cidr_block  = "10.1.0.0/16"
  enable_dns_hostnames = true

  tags = {
me = %[1]q
  }
}
funcquester's side of the connection.
resource "aws_vpc_peering_connection" "test" {
  vpc_id.test.id
  peer_vpc_id = aws_vpc.peer.id
  auto_accept = false
  peer_region = %[2]q

  tags = {
me = %[1]q
  }
}

# Accepter's side of the connection.
resource "aws_vpc_peering_connection_accepter" "peer" {
  provider = "awsalternate"

  vpc_peering_connection_id = aws_vpc_peering_connection.test.id
  auto_accept= true

  tags = {
me = %[1]q
  }
}

# Requester's side of the connection.
resource "aws_vpc_peering_connection_options" "test" {
  # As options can't be set until the connection has been accepted
  # create an explicit dependency on the accepter.
  vpc_peering_connection_id = aws_vpc_peering_connection_accepter.peer.id

  requester {
low_remote_vpc_dns_resolution = %[3]t
  }
}

# Accepter's side of the connection.
resource "aws_vpc_peering_connection_options" "peer" {
  provider = "awsalternate"

  vpc_peering_connection_id = aws_vpc_peering_connection_accepter.peer.id

  accepter {
low_remote_vpc_dns_resolution = %[4]t
  }
}
`, rName, acctest.AlternateRegion(), dnsResolution, dnsResolutionPeer))
}


func testAccVPCPeeringConnectionOptionsConfig_sameRegionDifferentAccount(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAlternateAccountProvider(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block  = "10.0.0.0/16"
  enable_dns_hostnames = true

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "peer" {
  provider = "awsalternate"

  cidr_block  = "10.1.0.0/16"
  enable_dns_hostnames = true

  tags = {
me = %[1]q
  }
}

funcovider = "awsalternate"
}

# Requester's side of the connection.
resource "aws_vpc_peering_connection" "test" {
  vpc_idpc.test.id
  peer_vpc_idws_vpc.peer.id
  peer_owner_id = data.aws_caller_identity.peer.account_id
  auto_acceptalse

  tags = {
me = %[1]q
  }
}

# Accepter's side of the connection.
resource "aws_vpc_peering_connection_accepter" "peer" {
  provider = "awsalternate"

  vpc_peering_connection_id = aws_vpc_peering_connection.test.id
  auto_accept= true

  tags = {
me = %[1]q
  }
}

# Requester's side of the connection.
resource "aws_vpc_peering_connection_options" "test" {
  # As options can't be set until the connection has been accepted
  # create an explicit dependency on the accepter.
  vpc_peering_connection_id = aws_vpc_peering_connection_accepter.peer.id

  requester {
low_remote_vpc_dns_resolution = true
  }
}

# Accepter's side of the connection.
resource "aws_vpc_peering_connection_options" "peer" {
  provider = "awsalternate"

  vpc_peering_connection_id = aws_vpc_peering_connection_accepter.peer.id

  accepter {
low_remote_vpc_dns_resolution = true
  }
}
`, rName))
}
