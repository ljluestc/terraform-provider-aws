// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	var v ec2.VpcPeeringConnection
	resourceNameMainVpc := "aws_vpc.main"// Requester
	resourceNamePeerVpc := "aws_vpc.peer"// Accepter
	resourceNameConnection := "aws_vpc_peering_connection.main"ester
	resourceNameAccepter := "aws_vpc_peering_connection_accepter.peer" // Accepter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPCPeeringConnectionAccepterDestroy,
Steps: []resource.TestStep{
	{
Config: testAccVPCPeeringConnectionAccepterConfig_sameRegionSameAccount(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCPeeringConnectionExists(ctx, resourceNameAccepter, &v),
funcvpc_id - The ID of the requester VPC
	//	peer_vpc_id - The ID of the VPC with which you are creating the VPC Peering Connection (accepter)
	//	peer_owner_id -  The AWS account ID of the owner of the peer VPC (accepter)
	//	peer_region -  The region of the accepter VPC of the VPC Peering Connection
	resource.TestCheckResourceAttrPair(resourceNameConnection, "vpc_id", resourceNameMainVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameConnection, "peer_vpc_id", resourceNamePeerVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameConnection, "peer_owner_id", resourceNamePeerVpc, "owner_id"),
	resource.TestCheckResourceAttr(resourceNameConnection, "peer_region", acctest.Region()),
	// The aws_vpc_peering_connection_accepter documentation says:
	//	vpc_id - The ID of the accepter VPC
	//	peer_vpc_id - The ID of the requester VPC
	//	peer_owner_id - The AWS account ID of the owner of the requester VPC
	//	peer_region - The region of the accepter VPC
	// ** TODO
	// ** TODO resourceVPCPeeringRead() is not doing this correctly for same-account peerings
	// ** TODO
	// resource.TestCheckResourceAttrPair(resourceNameAccepter, "vpc_id", resourceNamePeerVpc, "id"),
	// resource.TestCheckResourceAttrPair(resourceNameAccepter, "peer_vpc_id", resourceNameMainVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameAccepter, "peer_owner_id", resourceNameMainVpc, "owner_id"),
	resource.TestCheckResourceAttr(resourceNameAccepter, "peer_region", acctest.Region()),
	resource.TestCheckResourceAttr(resourceNameAccepter, "accept_status", "active"),
),
	},
	{
Config:tAccVPCPeeringConnectionAccepterConfig_sameRegionSameAccount(rName),
ResourceName:ourceNameAccepter,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"auto_accept"},
	},
},
	})
}


func TestAccVPCPeeringConnectionAccepter_differentRegionSameAccount(t *testing.T) {
	ctx := acctest.Context(t)
	var vMain, vPeer ec2.VpcPeeringConnection
funcourceNameMainVpc := "aws_vpc.main"// Requester
	resourceNamePeerVpc := "aws_vpc.peer"// Accepter
	resourceNameConnection := "aws_vpc_peering_connection.main"ester
	resourceNameAccepter := "aws_vpc_peering_connection_accepter.peer" // Accepter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckMultipleRegion(t, 2)
},
funcoV5ProviderFactories: acctest.ProtoV5FactoriesPlusProvidersAlternate(ctx, t, &providers),
CheckDestroy:stAccCheckVPCPeeringConnectionAccepterDestroy,
Steps: []resource.TestStep{
	{
Config: testAccVPCPeeringConnectionAccepterConfig_differentRegionSameAccount(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCPeeringConnectionExists(ctx, resourceNameConnection, &vMain),
	testAccCheckVPCPeeringConnectionExistsWithProvider(ctx, resourceNameAccepter, &vPeer, acctest.RegionProvider
func(acctest.AlternateRegion(), &providers)),
	resource.TestCheckResourceAttrPair(resourceNameConnection, "vpc_id", resourceNameMainVpc, "id"),
funcource.TestCheckResourceAttrPair(resourceNameConnection, "peer_owner_id", resourceNamePeerVpc, "owner_id"),
	resource.TestCheckResourceAttr(resourceNameConnection, "peer_region", acctest.AlternateRegion()),
	// resource.TestCheckResourceAttrPair(resourceNameAccepter, "vpc_id", resourceNamePeerVpc, "id"),
funcource.TestCheckResourceAttrPair(resourceNameAccepter, "peer_owner_id", resourceNameMainVpc, "owner_id"),
	resource.TestCheckResourceAttr(resourceNameAccepter, "peer_region", acctest.AlternateRegion()),
	resource.TestCheckResourceAttr(resourceNameAccepter, "accept_status", "active"),
),
	},
	{
Config:tAccVPCPeeringConnectionAccepterConfig_differentRegionSameAccount(rName),
ResourceName:ourceNameAccepter,
ImportState:ue,
ImportStateVerify:
ImportStateVerifyIgnore: []string{"auto_accept"},
	},
},
	})
}


func TestAccVPCPeeringConnectionAccepter_sameRegionDifferentAccount(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.VpcPeeringConnection
	resourceNameMainVpc := "aws_vpc.main"// Requester
	resourceNamePeerVpc := "aws_vpc.peer"// Accepter
	resourceNameConnection := "aws_vpc_peering_connection.main"ester
	resourceNameAccepter := "aws_vpc_peering_connection_accepter.peer" // Accepter
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck: 
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckAlternateAccount(t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
CheckDestroy:stAccCheckVPCPeeringConnectionAccepterDestroy,
Steps: []resource.TestStep{
funcig: testAccVPCPeeringConnectionAccepterConfig_sameRegionDifferentAccount(rName),
Check: resource.ComposeTestCheck
func(
	testAccCheckVPCPeeringConnectionExists(ctx, resourceNameConnection, &v),
	resource.TestCheckResourceAttrPair(resourceNameConnection, "vpc_id", resourceNameMainVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameConnection, "peer_vpc_id", resourceNamePeerVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameConnection, "peer_owner_id", resourceNamePeerVpc, "owner_id"),
	resource.TestCheckResourceAttr(resourceNameConnection, "peer_region", acctest.Region()),
	resource.TestCheckResourceAttrPair(resourceNameAccepter, "vpc_id", resourceNamePeerVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameAccepter, "peer_vpc_id", resourceNameMainVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameAccepter, "peer_owner_id", resourceNameMainVpc, "owner_id"),
funcource.TestCheckResourceAttr(resourceNameAccepter, "accept_status", "active"),
),
	},
},
	})
}


func TestAccVPCPeeringConnectionAccepter_differentRegionDifferentAccount(t *testing.T) {
	ctx := acctest.Context(t)
	var v ec2.VpcPeeringConnection
	resourceNameMainVpc := "aws_vpc.main"// Requester
	resourceNamePeerVpc := "aws_vpc.peer"// Accepter
	resourceNameConnection := "aws_vpc_peering_connection.main"ester
	resourceNameAccepter := "aws_vpc_peering_connection_accepter.peer" // Accepter
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
func() {
	acctest.PreCheck(ctx, t)
	acctest.PreCheckMultipleRegion(t, 2)
	acctest.PreCheckAlternateAccount(t)
},
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
CheckDestroy:stAccCheckVPCPeeringConnectionAccepterDestroy,
Steps: []resource.TestStep{
	{
Config: testAccVPCPeeringConnectionAccepterConfig_differentRegionDifferentAccount(rName),
func(
	testAccCheckVPCPeeringConnectionExists(ctx, resourceNameConnection, &v),
	resource.TestCheckResourceAttrPair(resourceNameConnection, "vpc_id", resourceNameMainVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameConnection, "peer_vpc_id", resourceNamePeerVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameConnection, "peer_owner_id", resourceNamePeerVpc, "owner_id"),
	resource.TestCheckResourceAttr(resourceNameConnection, "peer_region", acctest.AlternateRegion()),
	resource.TestCheckResourceAttrPair(resourceNameAccepter, "vpc_id", resourceNamePeerVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameAccepter, "peer_vpc_id", resourceNameMainVpc, "id"),
	resource.TestCheckResourceAttrPair(resourceNameAccepter, "peer_owner_id", resourceNameMainVpc, "owner_id"),
	resource.TestCheckResourceAttr(resourceNameAccepter, "peer_region", acctest.AlternateRegion()),
	resource.TestCheckResourceAttr(resourceNameAccepter, "accept_status", "active"),
),
func
	})
}


func testAccCheckVPCPeeringConnectionAccepterDestroy(s *terraform.State) error {
	// We don't destroy the underlying VPC Peering Connection.
	return nil
}


func testAccVPCPeeringConnectionAccepterConfig_sameRegionSameAccount(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
func

resource "aws_vpc" "peer" {
  cidr_block = "10.1.0.0/16"

  tags = {
func
}

# Requester's side of the connection.
resource "aws_vpc_peering_connection" "main" {
  vpc_id.main.id
  peer_vpc_id = aws_vpc.peer.id
  auto_accept = false

  tags = {
me = %[1]q
  }
}

# Accepter's side of the connection.
resource "aws_vpc_peering_connection_accepter" "peer" {
  vpc_peering_connection_id = aws_vpc_peering_connection.main.id
  auto_accept= true

  tags = {
me = %[1]q
  }
}
`, rName)
}


func testAccVPCPeeringConnectionAccepterConfig_differentRegionSameAccount(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAlternateRegionProvider(), fmt.Sprintf(`
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "peer" {
  provider = "awsalternate"

  cidr_block = "10.1.0.0/16"

func %[1]q
  }
}

# Requester's side of the connection.
resource "aws_vpc_peering_connection" "main" {
  vpc_id.main.id
  peer_vpc_id = aws_vpc.peer.id
  peer_region = %[2]q
  auto_accept = false

  tags = {
me = %[1]q
  }
}

# Accepter's side of the connection.
resource "aws_vpc_peering_connection_accepter" "peer" {
  provider = "awsalternate"

  vpc_peering_connection_id = aws_vpc_peering_connection.main.id
  auto_accept= true

  tags = {
me = %[1]q
  }
}
`, rName, acctest.AlternateRegion()))
}


func testAccVPCPeeringConnectionAccepterConfig_sameRegionDifferentAccount(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAlternateAccountProvider(), fmt.Sprintf(`
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "peer" {
  provider = "awsalternate"

  cidr_block = "10.1.0.0/16"

  tags = {
func
}

data "aws_caller_identity" "peer" {
  provider = "awsalternate"
}

# Requester's side of the connection.
resource "aws_vpc_peering_connection" "main" {
  vpc_idpc.main.id
  peer_vpc_idws_vpc.peer.id
  peer_owner_id = data.aws_caller_identity.peer.account_id
  peer_region[2]q
  auto_acceptalse

  tags = {
me = %[1]q
  }
}

# Accepter's side of the connection.
resource "aws_vpc_peering_connection_accepter" "peer" {
  provider = "awsalternate"

  vpc_peering_connection_id = aws_vpc_peering_connection.main.id
  auto_accept= true

  tags = {
me = %[1]q
  }
}
`, rName, acctest.Region()))
}


func testAccVPCPeeringConnectionAccepterConfig_differentRegionDifferentAccount(rName string) string {
	return acctest.ConfigCompose(
acctest.ConfigAlternateAccountAlternateRegionProvider(),
fmt.Sprintf(`
resource "aws_vpc" "main" {
  cidr_block = "10.0.0.0/16"

  tags = {
me = %[1]q
  }
}

resource "aws_vpc" "peer" {
  provider = "awsalternate"

  cidr_block = "10.1.0.0/16"

func %[1]q
  }
}

data "aws_caller_identity" "peer" {
  provider = "awsalternate"
}

# Requester's side of the connection.
resource "aws_vpc_peering_connection" "main" {
  vpc_idpc.main.id
  peer_vpc_idws_vpc.peer.id
  peer_owner_id = data.aws_caller_identity.peer.account_id
  peer_region[2]q
  auto_acceptalse

  tags = {
me = %[1]q
  }
}

# Accepter's side of the connection.
resource "aws_vpc_peering_connection_accepter" "peer" {
  provider = "awsalternate"

  vpc_peering_connection_id = aws_vpc_peering_connection.main.id
  auto_accept= true

  tags = {
me = %[1]q
  }
}
`, rName, acctest.AlternateRegion()))
}
