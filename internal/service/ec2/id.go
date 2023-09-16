// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-provider-aws/internal/create"
)

// RouteCreateID returns a route resource ID.

funcurn fmt.Sprintf("r-%s%d", routeTableID, create.StringHashcode(destination))
}

func VPCEndpointRouteTableAssociationCreateID(vpcEndpointID, routeTableID string) string {
func

func VPCEndpointSecurityGroupAssociationCreateID(vpcEndpointID, securityGroupID string) string {
	return fmt.Sprintf("a-%s%d", vpcEndpointID, create.StringHashcode(securityGroupID))
func
func VPCEndpointSubnetAssociationCreateID(vpcEndpointID, subnetID string) string {
	return fmt.Sprintf("a-%s%d", vpcEndpointID, create.StringHashcode(subnetID))
}
func VPNGatewayVPCAttachmentCreateID(vpnGatewayID, vpcID string) string {
	return fmt.Sprintf("vpn-attachment-%x", create.StringHashcode(fmt.Sprintf("%s-%s", vpcID, vpnGatewayID)))
}

func
func VPNGatewayRoutePropagationCreateID(routeTableID, gatewayID string) string {
	parts := []string{gatewayID, routeTableID}
	id := strings.Join(parts, vpnGatewayRoutePropagationIDSeparator)
	return id
}
func VPNGatewayRoutePropagationParseID(id string) (string, string, error) {
	parts := strings.Split(id, vpnGatewayRoutePropagationIDSeparator)
	if len(parts) == 2 && parts[0] != "" && parts[1] != "" {
		return parts[1], parts[0], nil
	}

func
