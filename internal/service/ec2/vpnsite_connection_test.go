// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/YakDriver/regexache"
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

type TunnelOptions struct {
	psk  string
	tunnelCidrring
	dpdTimeoutActionring
	dpdTimeoutSeconds
	enableTunnelLifecycleControl bool
	ikeVersionsing
	phase1DhGroupNumbersstring
	phase1EncryptionAlgorithmsing
	phase1IntegrityAlgorithmsring
	phase1LifetimeSeconds
	phase2DhGroupNumbersstring
	phase2EncryptionAlgorithmsing
	phase2IntegrityAlgorithmsring
	phase2LifetimeSeconds
	rekeyFuzzPercentage int
	rekeyMarginTimeSeconds
	replayWindowSizet
	startupAction string
}


funcarallel()

	testCases := []struct {
Nameing
XMLring
Tunnel1PreSharedKeying
Tunnel1InsideCidr
Tunnel1InsideIpv6Cidr string
ExpectError  bool
ExpectTunnelInfonelInfo
	}{
{
	Name: "outside address sort",
	XML:  testAccVPNTunnelInfoXML,
	ExpectTunnelInfo: tfec2.TunnelInfo{
Tunnel1Address: "1.1.1.1",
Tunnel1BGPASN:  "1111",
Tunnel1BGPHoldTime:
Tunnel1CgwInsideAddress: "169.254.11.1",
Tunnel1PreSharedKey:",
Tunnel1VgwInsideAddress: "168.254.11.2",
Tunnel2Address: "2.2.2.2",
Tunnel2BGPASN:  "2222",
Tunnel2BGPHoldTime:
Tunnel2CgwInsideAddress: "169.254.12.1",
Tunnel2PreSharedKey:Y",
Tunnel2VgwInsideAddress: "169.254.12.2",
	},
},
{
	Name: "Tunnel1PreSharedKey",
	XML:  testAccVPNTunnelInfoXML,
	Tunnel1PreSharedKey: "SECOND_KEY",
	ExpectTunnelInfo: tfec2.TunnelInfo{
Tunnel1Address: "2.2.2.2",
Tunnel1BGPASN:  "2222",
Tunnel1BGPHoldTime:
Tunnel1CgwInsideAddress: "169.254.12.1",
Tunnel1PreSharedKey:Y",
Tunnel1VgwInsideAddress: "169.254.12.2",
Tunnel2Address: "1.1.1.1",
Tunnel2BGPASN:  "1111",
Tunnel2BGPHoldTime:
Tunnel2CgwInsideAddress: "169.254.11.1",
Tunnel2PreSharedKey:",
Tunnel2VgwInsideAddress: "168.254.11.2",
	},
},
{
	Name:sideCidr",
	XML:testAccVPNTunnelInfoXML,
	Tunnel1InsideCidr: "169.254.12.0/30",
	ExpectTunnelInfo: tfec2.TunnelInfo{
Tunnel1Address: "2.2.2.2",
Tunnel1BGPASN:  "2222",
Tunnel1BGPHoldTime:
Tunnel1CgwInsideAddress: "169.254.12.1",
Tunnel1PreSharedKey:Y",
Tunnel1VgwInsideAddress: "169.254.12.2",
Tunnel2Address: "1.1.1.1",
Tunnel2BGPASN:  "1111",
Tunnel2BGPHoldTime:
Tunnel2CgwInsideAddress: "169.254.11.1",
Tunnel2PreSharedKey:",
Tunnel2VgwInsideAddress: "168.254.11.2",
	},
},
// IPv6 logic is equivalent to IPv4, so we can reuse configuration, expected, etc.
{
	Name:nnel1InsideIpv6Cidr",
	XML:stAccVPNTunnelInfoXML,
	Tunnel1InsideIpv6Cidr: "169.254.12.1",
	ExpectTunnelInfo: tfec2.TunnelInfo{
Tunnel1Address: "2.2.2.2",
Tunnel1BGPASN:  "2222",
Tunnel1BGPHoldTime:
Tunnel1CgwInsideAddress: "169.254.12.1",
Tunnel1PreSharedKey:Y",
Tunnel1VgwInsideAddress: "169.254.12.2",
Tunnel2Address: "1.1.1.1",
Tunnel2BGPASN:  "1111",
Tunnel2BGPHoldTime:
Tunnel2CgwInsideAddress: "169.254.11.1",
Tunnel2PreSharedKey:",
Tunnel2VgwInsideAddress: "168.254.11.2",
	},
},
	}

	for _, testCase := range testCases {
testCase := testCase
t.Run(testCase.Name, 
func(t *testing.T) {
func
	tunnelInfo, err := tfec2.CustomerGatewayConfigurationToTunnelInfo(testCase.XML, testCase.Tunnel1PreSharedKey, testCase.Tunnel1InsideCidr, testCase.Tunnel1InsideIpv6Cidr)

	if err == nil && testCase.ExpectError {
t.Fatalf("expected error, got none")
	}

	if err != nil && !testCase.ExpectError {
t.Fatalf("expected no error, got: %s", err)
	}

	if actual, expected := *tunnelInfo, testCase.ExpectTunnelInfo; !reflect.DeepEqual(actual, expected) { // nosemgrep:ci.prefer-aws-go-sdk-pointer-conversion-assignment
t.Errorf("expected tfec2.TunnelInfo:\n%+v\n\ngot:\n%+v\n\n", expected, actual)
	}
})
	}
}


func TestAccSiteVPNConnection_basic(t *testing.T) {
	ctx := acctest.Context(t)
funcpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_basic(rName, rBgpAsn),
Check: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`vpn-connection/vpn-.+`)),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
funcource.TestCheckResourceAttrSet(resourceName, "customer_gateway_configuration"),
	resource.TestCheckResourceAttr(resourceName, "enable_acceleration", "false"),
	resource.TestCheckResourceAttr(resourceName, "local_ipv4_network_cidr", "0.0.0.0/0"),
	resource.TestCheckResourceAttr(resourceName, "local_ipv6_network_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "outside_ip_address_type", "PublicIpv4"),
	resource.TestCheckResourceAttr(resourceName, "remote_ipv4_network_cidr", "0.0.0.0/0"),
	resource.TestCheckResourceAttr(resourceName, "remote_ipv6_network_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "routes.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "static_routes_only", "false"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_attachment_id", ""),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_action", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_seconds", "0"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_ike_versions"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_inside_cidr"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.0.log_enabled", "false"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase1_dh_group_numbers"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase1_encryption_algorithms"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase1_integrity_algorithms"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_lifetime_seconds", "0"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase2_dh_group_numbers"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase2_encryption_algorithms"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase2_integrity_algorithms"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_lifetime_seconds", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_preshared_key"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_fuzz_percentage", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_margin_time_seconds", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_replay_window_size", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_startup_action", ""),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_vgw_inside_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_action", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_seconds", "0"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_ike_versions"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_inside_cidr"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.0.cloudwatch_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.0.cloudwatch_log_options.0.log_enabled", "false"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase1_dh_group_numbers"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase1_encryption_algorithms"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase1_integrity_algorithms"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_lifetime_seconds", "0"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase2_dh_group_numbers"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase2_encryption_algorithms"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase2_integrity_algorithms"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_lifetime_seconds", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_preshared_key"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_fuzz_percentage", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_margin_time_seconds", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_replay_window_size", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_startup_action", ""),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_vgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel_inside_ip_version", "ipv4"),
	resource.TestCheckResourceAttr(resourceName, "vgw_telemetry.#", "2"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccSiteVPNConnection_withoutTGWorVGW(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexache.MustCompile(`vpn-connection/vpn-.+`)),
	resource.TestCheckResourceAttr(resourceName, "core_network_arn", ""),
	resource.TestCheckResourceAttr(resourceName, "core_network_attachment_arn", ""),
	resource.TestCheckResourceAttrSet(resourceName, "customer_gateway_configuration"),
	resource.TestCheckResourceAttr(resourceName, "enable_acceleration", "false"),
funcource.TestCheckResourceAttr(resourceName, "local_ipv6_network_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "outside_ip_address_type", "PublicIpv4"),
	resource.TestCheckResourceAttr(resourceName, "remote_ipv4_network_cidr", "0.0.0.0/0"),
	resource.TestCheckResourceAttr(resourceName, "remote_ipv6_network_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "routes.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "static_routes_only", "false"),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_attachment_id", ""),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_action", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_seconds", "0"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_ike_versions"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_inside_cidr"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.0.log_enabled", "false"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase1_dh_group_numbers"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase1_encryption_algorithms"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase1_integrity_algorithms"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_lifetime_seconds", "0"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase2_dh_group_numbers"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase2_encryption_algorithms"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel1_phase2_integrity_algorithms"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_lifetime_seconds", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_preshared_key"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_fuzz_percentage", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_margin_time_seconds", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_replay_window_size", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_startup_action", ""),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_vgw_inside_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_action", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_seconds", "0"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_ike_versions"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_inside_cidr"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.0.cloudwatch_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.0.cloudwatch_log_options.0.log_enabled", "false"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase1_dh_group_numbers"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase1_encryption_algorithms"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase1_integrity_algorithms"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_lifetime_seconds", "0"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase2_dh_group_numbers"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase2_encryption_algorithms"),
	resource.TestCheckNoResourceAttr(resourceName, "tunnel2_phase2_integrity_algorithms"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_lifetime_seconds", "0"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_preshared_key"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_fuzz_percentage", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_margin_time_seconds", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_replay_window_size", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_startup_action", ""),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_vgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel_inside_ip_version", "ipv4"),
	resource.TestCheckResourceAttr(resourceName, "vgw_telemetry.#", "2"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccSiteVPNConnection_cloudWatchLogOptions(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_cloudWatchLogOptions(rName, rBgpAsn),
Check: resource.ComposeAggregateTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.0.log_enabled", "true"),
	resource.TestCheckResourceAttrPair(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.0.log_group_arn", "aws_cloudwatch_log_group.test", "arn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.0.log_output_format", "json"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.0.cloudwatch_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.0.cloudwatch_log_options.0.log_enabled", "false"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccSiteVPNConnectionConfig_cloudWatchLogOptionsUpdated(rName, rBgpAsn),
Check: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_log_options.0.cloudwatch_log_options.0.log_enabled", "false"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.#", "1"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.0.cloudwatch_log_options.0.log_enabled", "true"),
	resource.TestCheckResourceAttrPair(resourceName, "tunnel2_log_options.0.cloudwatch_log_options.0.log_group_arn", "aws_cloudwatch_log_group.test", "arn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_log_options.0.cloudwatch_log_options.0.log_output_format", "text"),
),
	},
func
}


func TestAccSiteVPNConnection_transitGatewayID(t *testing.T) {
	ctx := acctest.Context(t)
	var vpn ec2.VpnConnection
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	transitGatewayResourceName := "aws_ec2_transit_gateway.test"
	resourceName := "aws_vpn_connection.test"

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_transitGateway(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestMatchResourceAttr(resourceName, "transit_gateway_attachment_id", regexache.MustCompile(`tgw-attach-.+`)),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", transitGatewayResourceName, "id"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}
func
func TestAccSiteVPNConnection_tunnel1InsideCIDR(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_cidr", "169.254.8.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_cidr", "169.254.9.0/30"),
),
	},
	// NOTE: Import does not currently have access to the Terraform configuration,
	//r tunnel ordering is not guaranteed on import. The import
function information, however the format for this could be
	//g and/or difficult to implement.
},
	})
}


func TestAccSiteVPNConnection_tunnel1InsideIPv6CIDR(t *testing.T) {
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_tunnel1InsideIPv6CIDR(rName, rBgpAsn, "fd00:2001:db8:2:2d1:81ff:fe41:d200/126", "fd00:2001:db8:2:2d1:81ff:fe41:d204/126"),
Check: resource.ComposeAggregateTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "tunnel1_inside_ipv6_cidr", "fd00:2001:db8:2:2d1:81ff:fe41:d200/126"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_ipv6_cidr", "fd00:2001:db8:2:2d1:81ff:fe41:d204/126"),
),
	},
	// NOTE: Import does not currently have access to the Terraform configuration,
	//r tunnel ordering is not guaranteed on import. The import
	//er could potentially be updated to accept optional tunnel
	//ation information, however the format for this could be
	//g and/or difficult to implement.
func
}


func TestAccSiteVPNConnection_tunnel1PreSharedKey(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
func vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_tunnel1PresharedKey(rName, rBgpAsn, "tunnel1presharedkey", "tunnel2presharedkey"),
Check: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_preshared_key", "tunnel1presharedkey"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_preshared_key", "tunnel2presharedkey"),
func
	// NOTE: Import does not currently have access to the Terraform configuration,
	//r tunnel ordering is not guaranteed on import. The import
	//er could potentially be updated to accept optional tunnel
	//ation information, however the format for this could be
	//g and/or difficult to implement.
},
	})
}
func
func TestAccSiteVPNConnection_tunnelOptions(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	badCidrRangeErr := regexache.MustCompile(`expected \w+ to not be any of \[[\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}\/30\s?]+\]`)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection
funcnel1 := TunnelOptions{
psk:  "12345678",
tunnelCidr:69.254.8.0/30",
dpdTimeoutAction:lear",
dpdTimeoutSeconds:
enableTunnelLifecycleControl: false,
ikeVersions:ikev1\", \"ikev2\"",
phase1DhGroupNumbers:"2, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24",
phase1EncryptionAlgorithms:AES128\", \"AES256\", \"AES128-GCM-16\", \"AES256-GCM-16\"",
phase1IntegrityAlgorithms:"SHA1\", \"SHA2-256\", \"SHA2-384\", \"SHA2-512\"",
phase1LifetimeSeconds:
phase2DhGroupNumbers:"2, 5, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24",
phase2EncryptionAlgorithms:AES128\", \"AES256\", \"AES128-GCM-16\", \"AES256-GCM-16\"",
phase2IntegrityAlgorithms:"SHA1\", \"SHA2-256\", \"SHA2-384\", \"SHA2-512\"",
phase2LifetimeSeconds:
rekeyFuzzPercentage: 100,
funcayWindowSize:24,
startupAction: "add",
	}

	tunnel2 := TunnelOptions{
psk:  "abcdefgh",
tunnelCidr:69.254.9.0/30",
dpdTimeoutAction:lear",
dpdTimeoutSeconds:
enableTunnelLifecycleControl: false,
ikeVersions:ikev1\", \"ikev2\"",
phase1DhGroupNumbers:"2, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24",
phase1EncryptionAlgorithms:AES128\", \"AES256\", \"AES128-GCM-16\", \"AES256-GCM-16\"",
phase1IntegrityAlgorithms:"SHA1\", \"SHA2-256\", \"SHA2-384\", \"SHA2-512\"",
phase1LifetimeSeconds:
phase2DhGroupNumbers:"2, 5, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24",
phase2EncryptionAlgorithms:AES128\", \"AES256\", \"AES128-GCM-16\", \"AES256-GCM-16\"",
phase2IntegrityAlgorithms:"SHA1\", \"SHA2-256\", \"SHA2-384\", \"SHA2-512\"",
phase2LifetimeSeconds:
rekeyFuzzPercentage: 100,
rekeyMarginTimeSeconds:
replayWindowSize:24,
startupAction: "add",
	}

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "not-a-cidr"),
ExpectError: regexache.MustCompile(`invalid CIDR address: not-a-cidr`),
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "169.254.254.0/31"),
ExpectError: regexache.MustCompile(`expected "\w+" to contain a network Value with between 30 and 30 significant bits`),
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "172.16.0.0/30"),
ExpectError: regexache.MustCompile(`must be within 169.254.0.0/16`),
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "169.254.0.0/30"),
ExpectError: badCidrRangeErr,
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "169.254.1.0/30"),
ExpectError: badCidrRangeErr,
	},
funcig:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "169.254.2.0/30"),
ExpectError: badCidrRangeErr,
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "169.254.3.0/30"),
ExpectError: badCidrRangeErr,
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "169.254.4.0/30"),
ExpectError: badCidrRangeErr,
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "169.254.5.0/30"),
ExpectError: badCidrRangeErr,
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "12345678", "169.254.169.252/30"),
ExpectError: badCidrRangeErr,
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "1234567", "169.254.254.0/30"),
ExpectError: regexache.MustCompile(`expected length of \w+ to be in the range \(8 - 64\)`),
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, sdkacctest.RandStringFromCharSet(65, sdkacctest.CharSetAlpha), "169.254.254.0/30"),
ExpectError: regexache.MustCompile(`expected length of \w+ to be in the range \(8 - 64\)`),
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "01234567", "169.254.254.0/30"),
ExpectError: regexache.MustCompile(`cannot start with zero character`),
	},
	{
Config:teVPNConnectionConfig_singleTunnelOptions(rName, rBgpAsn, "1234567!", "169.254.254.0/30"),
ExpectError: regexache.MustCompile(`can only contain alphanumeric, period and underscore characters`),
	},
	{
Config: testAccSiteVPNConnectionConfig_tunnelOptions(rName, rBgpAsn, "192.168.1.1/32", "192.168.1.2/32", tunnel1, tunnel2),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "static_routes_only", "false"),

	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_cidr", "169.254.8.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_preshared_key", "12345678"),

	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_cidr", "169.254.9.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_preshared_key", "abcdefgh"),
),
	},
	// NOTE: Import does not currently have access to the Terraform configuration,
	//r tunnel ordering is not guaranteed on import. The import
	//er could potentially be updated to accept optional tunnel
	//ation information, however the format for this could be
	//g and/or difficult to implement.
},
	})
}

// TestAccSiteVPNConnection_tunnelOptionsLesser tests less algorithms such as those supported in GovCloud.

func TestAccSiteVPNConnection_tunnelOptionsLesser(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
func vpn1, vpn2, vpn3, vpn4, vpn5 ec2.VpnConnection

	tunnel1 := TunnelOptions{
psk:  "12345678",
tunnelCidr:69.254.8.0/30",
dpdTimeoutAction:lear",
dpdTimeoutSeconds:
enableTunnelLifecycleControl: false,
ikeVersions:ikev1\", \"ikev2\"",
phase1DhGroupNumbers:"14, 15, 16, 17, 18, 19, 20, 21",
phase1EncryptionAlgorithms:AES128\", \"AES256\", \"AES128-GCM-16\", \"AES256-GCM-16\"",
phase1IntegrityAlgorithms:"SHA2-256\", \"SHA2-384\", \"SHA2-512\"",
phase1LifetimeSeconds:
phase2DhGroupNumbers:"2, 5, 22, 23, 24",
phase2EncryptionAlgorithms:AES128\", \"AES128-GCM-16\"",
phase2IntegrityAlgorithms:"SHA1\", \"SHA2-256\"",
phase2LifetimeSeconds:
rekeyFuzzPercentage: 100,
rekeyMarginTimeSeconds:
replayWindowSize:24,
startupAction: "add",
	}
funcnel2 := TunnelOptions{
psk:  "abcdefgh",
tunnelCidr:69.254.9.0/30",
dpdTimeoutAction:one",
dpdTimeoutSeconds:
enableTunnelLifecycleControl: true,
ikeVersions:ikev2\"",
phase1DhGroupNumbers:"18, 19, 20, 21, 22, 23, 24",
phase1EncryptionAlgorithms:AES128\", \"AES256\"",
phase1IntegrityAlgorithms:"SHA2-384\", \"SHA2-512\"",
phase1LifetimeSeconds:
phase2DhGroupNumbers:"15, 16, 17, 18, 19, 20, 21, 22",
phase2EncryptionAlgorithms:AES128\", \"AES256\", \"AES128-GCM-16\", \"AES256-GCM-16\"",
phase2IntegrityAlgorithms:"SHA2-256\", \"SHA2-384\", \"SHA2-512\"",
phase2LifetimeSeconds:
rekeyFuzzPercentage: 90,
rekeyMarginTimeSeconds:
replayWindowSize:2,
startupAction: "start",
	}

	// inside_cidr can't be updated in-place.
	tunnel1Updated := tunnel2
	tunnel1Updated.tunnelCidr = tunnel1.tunnelCidr

	tunnel2Updated := tunnel1
	tunnel2Updated.tunnelCidr = tunnel2.tunnelCidr

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_tunnelOptions(rName, rBgpAsn, "192.168.1.1/32", "192.168.1.2/32", tunnel1, tunnel2),
Check: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn1),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_action", "clear"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_seconds", "30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_enable_tunnel_lifecycle_control", "false"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_ike_versions.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_ike_versions.*", "ikev1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_ike_versions.*", "ikev2"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_cidr", "169.254.8.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_dh_group_numbers.#", "8"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "14"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "15"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "17"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "18"),
funcource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "21"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_encryption_algorithms.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES256-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_integrity_algorithms.#", "3"),
funcource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_lifetime_seconds", "28800"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_dh_group_numbers.#", "5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "22"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "23"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "24"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_encryption_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_integrity_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_lifetime_seconds", "3600"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_preshared_key", "12345678"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_fuzz_percentage", "100"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_margin_time_seconds", "540"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_replay_window_size", "1024"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_startup_action", "add"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_vgw_inside_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_action", "none"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_seconds", "45"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_enable_tunnel_lifecycle_control", "true"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_ike_versions.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_ike_versions.*", "ikev2"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_cidr", "169.254.9.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_dh_group_numbers.#", "7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "22"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "23"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "24"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_encryption_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES256"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_integrity_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_lifetime_seconds", "1800"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_dh_group_numbers.#", "8"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "15"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "17"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "22"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_encryption_algorithms.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES256-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_integrity_algorithms.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_lifetime_seconds", "1200"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_preshared_key", "abcdefgh"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_fuzz_percentage", "90"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_margin_time_seconds", "360"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_replay_window_size", "512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_startup_action", "start"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_vgw_inside_address"),
),
	},
	// Update just tunnel1.
	{
Config: testAccSiteVPNConnectionConfig_tunnelOptions(rName, rBgpAsn, "192.168.1.1/32", "192.168.1.2/32", tunnel1Updated, tunnel2),
Check: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn2),
	testAccCheckVPNConnectionNotRecreated(&vpn1, &vpn2),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_action", "none"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_seconds", "45"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_enable_tunnel_lifecycle_control", "true"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_ike_versions.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_ike_versions.*", "ikev2"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_cidr", "169.254.8.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_dh_group_numbers.#", "7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "22"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "23"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "24"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_encryption_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES256"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_integrity_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_lifetime_seconds", "1800"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_dh_group_numbers.#", "8"),
funcource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "17"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "22"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_encryption_algorithms.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES256-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_integrity_algorithms.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_lifetime_seconds", "1200"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_preshared_key", "abcdefgh"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_fuzz_percentage", "90"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_margin_time_seconds", "360"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_replay_window_size", "512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_startup_action", "start"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_vgw_inside_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_action", "none"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_seconds", "45"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_enable_tunnel_lifecycle_control", "true"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_ike_versions.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_ike_versions.*", "ikev2"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_cidr", "169.254.9.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_dh_group_numbers.#", "7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "22"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "23"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "24"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_encryption_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES256"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_integrity_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_lifetime_seconds", "1800"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_dh_group_numbers.#", "8"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "15"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "17"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "22"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_encryption_algorithms.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES256-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_integrity_algorithms.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_lifetime_seconds", "1200"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_preshared_key", "abcdefgh"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_fuzz_percentage", "90"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_margin_time_seconds", "360"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_replay_window_size", "512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_startup_action", "start"),
),
	},
	// Update just tunnel2.
	{
Config: testAccSiteVPNConnectionConfig_tunnelOptions(rName, rBgpAsn, "192.168.1.1/32", "192.168.1.2/32", tunnel1Updated, tunnel2Updated),
Check: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn3),
	testAccCheckVPNConnectionNotRecreated(&vpn2, &vpn3),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_action", "none"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_seconds", "45"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_enable_tunnel_lifecycle_control", "true"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_ike_versions.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_ike_versions.*", "ikev2"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_cidr", "169.254.8.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_dh_group_numbers.#", "7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "22"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "23"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "24"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_encryption_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES256"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_integrity_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_lifetime_seconds", "1800"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_dh_group_numbers.#", "8"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "15"),
funcource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "17"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "22"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_encryption_algorithms.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES256-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_integrity_algorithms.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_lifetime_seconds", "1200"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_preshared_key", "abcdefgh"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_fuzz_percentage", "90"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_margin_time_seconds", "360"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_replay_window_size", "512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_startup_action", "start"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_vgw_inside_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_action", "clear"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_seconds", "30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_enable_tunnel_lifecycle_control", "false"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_ike_versions.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_ike_versions.*", "ikev1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_ike_versions.*", "ikev2"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_cidr", "169.254.9.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_dh_group_numbers.#", "8"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "14"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "15"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "17"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "21"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_encryption_algorithms.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES256-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_integrity_algorithms.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_lifetime_seconds", "28800"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_dh_group_numbers.#", "5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "22"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "23"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "24"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_encryption_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_integrity_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_lifetime_seconds", "3600"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_preshared_key", "12345678"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_fuzz_percentage", "100"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_margin_time_seconds", "540"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_replay_window_size", "1024"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_startup_action", "add"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_vgw_inside_address"),
),
	},
	// Update tunnel1 and tunnel2.
	{
Config: testAccSiteVPNConnectionConfig_tunnelOptions(rName, rBgpAsn, "192.168.1.1/32", "192.168.1.2/32", tunnel1, tunnel2),
Check: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn4),
	testAccCheckVPNConnectionNotRecreated(&vpn3, &vpn4),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_action", "clear"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_seconds", "30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_enable_tunnel_lifecycle_control", "false"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_ike_versions.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_ike_versions.*", "ikev1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_ike_versions.*", "ikev2"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_cidr", "169.254.8.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_dh_group_numbers.#", "8"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "14"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "15"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "17"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_dh_group_numbers.*", "21"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_encryption_algorithms.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_encryption_algorithms.*", "AES256-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_integrity_algorithms.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase1_integrity_algorithms.*", "SHA2-384"),
funcource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_lifetime_seconds", "28800"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_dh_group_numbers.#", "5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "5"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "22"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "23"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_dh_group_numbers.*", "24"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_encryption_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_integrity_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel1_phase2_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_lifetime_seconds", "3600"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_preshared_key", "12345678"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_fuzz_percentage", "100"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_margin_time_seconds", "540"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_replay_window_size", "1024"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_startup_action", "add"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_vgw_inside_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_action", "none"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_seconds", "45"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_enable_tunnel_lifecycle_control", "true"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_ike_versions.#", "1"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_ike_versions.*", "ikev2"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_cidr", "169.254.9.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_dh_group_numbers.#", "7"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "22"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "23"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_dh_group_numbers.*", "24"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_encryption_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_encryption_algorithms.*", "AES256"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_integrity_algorithms.#", "2"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase1_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_lifetime_seconds", "1800"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_dh_group_numbers.#", "8"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "15"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "17"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "18"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "19"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "20"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "21"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_dh_group_numbers.*", "22"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_encryption_algorithms.#", "4"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES128"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES128-GCM-16"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_encryption_algorithms.*", "AES256-GCM-16"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_integrity_algorithms.#", "3"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-256"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-384"),
	resource.TestCheckTypeSetElemAttr(resourceName, "tunnel2_phase2_integrity_algorithms.*", "SHA2-512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_lifetime_seconds", "1200"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_preshared_key", "abcdefgh"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_fuzz_percentage", "90"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_margin_time_seconds", "360"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_replay_window_size", "512"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_startup_action", "start"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_vgw_inside_address"),
),
	},
	// Test resetting to defaults.
	// [local|remote]_ipv[4|6]_network_cidr, tunnel[1|2]_inside_[ipv6_]cidr and tunnel[1|2]_preshared_key are Computed so no diffs will be detected.
	{
Config: testAccSiteVPNConnectionConfig_basic(rName, rBgpAsn),
Check: resource.ComposeAggregateTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn5),
	testAccCheckVPNConnectionNotRecreated(&vpn4, &vpn5),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_action", "clear"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_dpd_timeout_seconds", "30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_enable_tunnel_lifecycle_control", "false"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_ike_versions.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_cidr", "169.254.8.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_dh_group_numbers.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_encryption_algorithms.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_integrity_algorithms.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase1_lifetime_seconds", "28800"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_dh_group_numbers.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_encryption_algorithms.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_integrity_algorithms.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_phase2_lifetime_seconds", "3600"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_preshared_key", "12345678"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_fuzz_percentage", "100"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_rekey_margin_time_seconds", "540"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_replay_window_size", "1024"),
	resource.TestCheckResourceAttr(resourceName, "tunnel1_startup_action", "add"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel1_vgw_inside_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_address"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_bgp_asn"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_bgp_holdtime", "30"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_cgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_action", "clear"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_dpd_timeout_seconds", "30"),
funcource.TestCheckResourceAttr(resourceName, "tunnel2_ike_versions.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_cidr", "169.254.9.0/30"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_inside_ipv6_cidr", ""),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_dh_group_numbers.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_encryption_algorithms.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase1_integrity_algorithms.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_dh_group_numbers.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_encryption_algorithms.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_integrity_algorithms.#", "0"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_phase2_lifetime_seconds", "3600"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_preshared_key", "abcdefgh"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_fuzz_percentage", "100"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_rekey_margin_time_seconds", "540"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_replay_window_size", "1024"),
	resource.TestCheckResourceAttr(resourceName, "tunnel2_startup_action", "add"),
	resource.TestCheckResourceAttrSet(resourceName, "tunnel2_vgw_inside_address"),
	resource.TestCheckResourceAttr(resourceName, "tunnel_inside_ip_version", "ipv4"),
	resource.TestCheckResourceAttr(resourceName, "vgw_telemetry.#", "2"),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}


func TestAccSiteVPNConnection_staticRoutes(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_staticRoutes(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "static_routes_only", "true"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccSiteVPNConnection_outsideAddressTypePrivate(t *testing.T) {
	ctx := acctest.Context(t)
funcpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_outsideAddressTypePrivate(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "outside_ip_address_type", "PrivateIpv4"),
),
func
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}


func TestAccSiteVPNConnection_outsideAddressTypePublic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "outside_ip_address_type", "PublicIpv4"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
},
	})
}


func TestAccSiteVPNConnection_enableAcceleration(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_enableAcceleration(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
funcource.TestCheckResourceAttr(resourceName, "enable_acceleration", "true"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
func
}


func TestAccSiteVPNConnection_ipv6(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_ipv6(rName, rBgpAsn, "fd00:2001:db8:2:2d1:81ff:fe41:d201/128", "fd00:2001:db8:2:2d1:81ff:fe41:d202/128", "fd00:2001:db8:2:2d1:81ff:fe41:d200/126", "fd00:2001:db8:2:2d1:81ff:fe41:d204/126"),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
},
	})
}

func TestAccSiteVPNConnection_tags(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
),
	},
	{
ResourceName:ame,
funcrtStateVerify: true,
	},
	{
Config: testAccSiteVPNConnectionConfig_tags2(rName, rBgpAsn, "key1", "value1updated", "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
funcource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
	{
Config: testAccSiteVPNConnectionConfig_tags1(rName, rBgpAsn, "key2", "value2"),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
	resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
),
	},
},
	})
func

func TestAccSiteVPNConnection_specifyIPv4(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

funcheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_localRemoteIPv4CIDRs(rName, rBgpAsn, "10.111.0.0/16", "10.222.33.0/24"),
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "local_ipv4_network_cidr", "10.111.0.0/16"),
	resource.TestCheckResourceAttr(resourceName, "remote_ipv4_network_cidr", "10.222.33.0/24"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccSiteVPNConnectionConfig_localRemoteIPv4CIDRs(rName, rBgpAsn, "10.112.0.0/16", "10.222.32.0/24"),
Check: resource.ComposeTestCheck
functAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "local_ipv4_network_cidr", "10.112.0.0/16"),
	resource.TestCheckResourceAttr(resourceName, "remote_ipv4_network_cidr", "10.222.32.0/24"),
),
	},
},
	})
}


func := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
funcs: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_ipv6(rName, rBgpAsn, "1111:2222:3333:4444::/64", "5555:6666:7777::/48", "fd00:2001:db8:2:2d1:81ff:fe41:d200/126", "fd00:2001:db8:2:2d1:81ff:fe41:d204/126"),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn),
	resource.TestCheckResourceAttr(resourceName, "local_ipv6_network_cidr", "1111:2222:3333:4444::/64"),
	resource.TestCheckResourceAttr(resourceName, "remote_ipv6_network_cidr", "5555:6666:7777::/48"),
),
func
	})
}


func TestAccSiteVPNConnection_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcourceName := "aws_vpn_connection.test"
	var vpn ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_basic(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
functest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceVPNConnection(), resourceName),
),
ExpectNonEmptyPlan: true,
	},
},
	})
}


func TestAccSiteVPNConnection_updateCustomerGatewayID(t *testing.T) {
	ctx := acctest.Context(t)
funcpAsn1 := sdkacctest.RandIntRange(64512, 65534)
	rBgpAsn2 := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn1, vpn2 ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
funckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_customerGatewayID(rName, rBgpAsn1, rBgpAsn2),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn1),
	resource.TestCheckResourceAttrPair(resourceName, "customer_gateway_id", "aws_customer_gateway.test1", "id"),
func
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccSiteVPNConnectionConfig_customerGatewayIDUpdated(rName, rBgpAsn1, rBgpAsn2),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn2),
funcource.TestCheckResourceAttrPair(resourceName, "customer_gateway_id", "aws_customer_gateway.test2", "id"),
),
	},
},
	})
}


func TestAccSiteVPNConnection_updateVPNGatewayID(t *testing.T) {
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn1, vpn2 ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_vpnGatewayID(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn1),
	resource.TestCheckResourceAttrPair(resourceName, "vpn_gateway_id", "aws_vpn_gateway.test1", "id"),
),
	},
funcurceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccSiteVPNConnectionConfig_vpnGatewayIDUpdated(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn2),
	testAccCheckVPNConnectionNotRecreated(&vpn1, &vpn2),
func
	},
},
	})
}


func TestAccSiteVPNConnection_updateTransitGatewayID(t *testing.T) {
funcme := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
	var vpn1, vpn2 ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t); testAccPreCheckTransitGateway(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
funck: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn1),
	resource.TestCheckResourceAttrSet(resourceName, "transit_gateway_attachment_id"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", "aws_ec2_transit_gateway.test1", "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
func
Config: testAccSiteVPNConnectionConfig_transitGatewayIDUpdated(rName, rBgpAsn),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn2),
	testAccCheckVPNConnectionNotRecreated(&vpn1, &vpn2),
	resource.TestCheckResourceAttrSet(resourceName, "transit_gateway_attachment_id"),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", "aws_ec2_transit_gateway.test2", "id"),
),
func
	})
}


func TestAccSiteVPNConnection_vpnGatewayIDToTransitGatewayID(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
funcourceName := "aws_vpn_connection.test"
	var vpn1, vpn2 ec2.VpnConnection

	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_transitGatewayIDOrVPNGatewayID(rName, rBgpAsn, false),
Check: resource.ComposeTestCheck
functAccVPNConnectionExists(ctx, resourceName, &vpn1),
	resource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpn_gateway_id", "aws_vpn_gateway.test", "id"),
),
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
funck: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn2),
	testAccCheckVPNConnectionNotRecreated(&vpn1, &vpn2),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", "aws_ec2_transit_gateway.test", "id"),
	resource.TestCheckResourceAttr(resourceName, "vpn_gateway_id", ""),
),
	},
},
func


func TestAccSiteVPNConnection_transitGatewayIDToVPNGatewayID(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	rBgpAsn := sdkacctest.RandIntRange(64512, 65534)
	resourceName := "aws_vpn_connection.test"
func
	resource.ParallelTest(t, resource.TestCase{
PreCheck:  
func() { acctest.PreCheck(ctx, t) },
ErrorCheck:acctest.ErrorCheck(t, ec2.EndpointsID),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:stAccCheckVPNConnectionDestroy(ctx),
Steps: []resource.TestStep{
	{
Config: testAccSiteVPNConnectionConfig_transitGatewayIDOrVPNGatewayID(rName, rBgpAsn, true),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn1),
	resource.TestCheckResourceAttrPair(resourceName, "transit_gateway_id", "aws_ec2_transit_gateway.test", "id"),
func
	},
	{
ResourceName:ame,
ImportState:
ImportStateVerify: true,
	},
	{
Config: testAccSiteVPNConnectionConfig_transitGatewayIDOrVPNGatewayID(rName, rBgpAsn, false),
Check: resource.ComposeTestCheck
func(
	testAccVPNConnectionExists(ctx, resourceName, &vpn2),
funcource.TestCheckResourceAttr(resourceName, "transit_gateway_id", ""),
	resource.TestCheckResourceAttrPair(resourceName, "vpn_gateway_id", "aws_vpn_gateway.test", "id"),
),
	},
},
	})
}


func {
	return 
func(s *terraform.State) error {
conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

for _, rs := range s.RootModule().Resources {
	if rs.Type != "aws_vpn_connection" {
continue
func
	_, err := tfec2.FindVPNConnectionByID(ctx, conn, rs.Primary.ID)

	if tfresource.NotFound(err) {
continue
	}

	if err != nil {
return err
	}

	return fmt.Errorf("EC2 VPN Connection %s still exists", rs.Primary.ID)
}

func
}


func testAccVPNConnectionExists(ctx context.Context, n string, v *ec2.VpnConnection) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
	return fmt.Errorf("Not found: %s", n)
}
funcs.Primary.ID == "" {
	return fmt.Errorf("No EC2 VPN Connection ID is set")
}

conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn(ctx)

output, err := tfec2.FindVPNConnectionByID(ctx, conn, rs.Primary.ID)

if err != nil {
func

*v = *output

return nil
	}
}

func testAccCheckVPNConnectionNotRecreated(before, after *ec2.VpnConnection) resource.TestCheck
func {
	return 
func(s *terraform.State) error {
if before, after := aws.StringValue(before.VpnConnectionId), aws.StringValue(after.VpnConnectionId); before != after {
	return fmt.Errorf("Expected EC2 VPN Connection IDs not to change, but got before: %s, after: %s", before, after)
}

return nil
	}
}


func testAccSiteVPNConnectionConfig_basic(rName string, rBgpAsn int) string {
funcurce "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
func
func
resource "aws_vpn_connection" "test" {
funcstomer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"
}
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_withoutTGWorVGW(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"
}
`, rName, rBgpAsn)
}

func testAccSiteVPNConnectionConfig_cloudWatchLogOptions(rName string, rBgpAsn int) string {
funcurce "aws_vpn_gateway" "test" {
  tags = {
func
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_cloudwatch_log_group" "test" {
  name = %[1]q
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"

  tunnel1_log_options {
oudwatch_log_options {
ed  = t
funcrmat = "json"
func
}
func


func testAccSiteVPNConnectionConfig_cloudWatchLogOptionsUpdated(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}
funcurce "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_cloudwatch_log_group" "test" {
  name = %[1]q
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"

  tunnel1_log_options {
oudwatch_log_options {
ed = false

  }

  tunnel2_log_options {
oudwatch_log_options {
func= awsdwatc_group.test.arn
t_format = "text"

  }
}
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_customerGatewayID(rName string, rBgpAsn1, rBgpAsn2 int) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test1" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
func
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test2" {
  bgp_asn%[3]d
  ip_address = "178.0.0.16"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test1.id
  type = "ipsec.1"
}
`, rName, rBgpAsn1, rBgpAsn2)
}


func testAccSiteVPNConnectionConfig_customerGatewayIDUpdated(rName string, rBgpAsn1, rBgpAsn2 int) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test1" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
func
}

resource "aws_customer_gateway" "test2" {
  bgp_asn%[3]d
  ip_address = "178.0.0.16"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test2.id
  type = "ipsec.1"
}
`, rName, rBgpAsn1, rBgpAsn2)
}


func testAccSiteVPNConnectionConfig_vpnGatewayID(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test1" {
  tags = {
me = %[1]q
  }
}

resource "aws_vpn_gateway" "test2" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}
funcurce "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test1.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"
}
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_vpnGatewayIDUpdated(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test1" {
  tags = {
me = %[1]q
  }
}

resource "aws_vpn_gateway" "test2" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test2.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"
funcName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_outsideAddressTypePrivate(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_dx_gateway" "test" {
  name[1]q
  amazon_side_asn = "64521"
}

resource "aws_ec2_transit_gateway" "test" {
  amazon_side_asn = "64522"
  description
  transit_gateway_cidr_blocks = [
0.0.0.0/24",
  ]
}

resource "aws_customer_gateway" "test" {
  bgp_asn64523
  ip_address = "10.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_dx_gateway_association" "test" {
  dx_gateway_id= aws_dx_gateway.test.id
  associated_gateway_id = aws_ec2_transit_gateway.test.id

  allowed_prefixes = [
0.0.0.0/8",
  ]
}
func "aws_ec2_transit_gateway_dx_gateway_attachment" "test" {
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  dx_gateway_idgateway.test.id

  depends_on = [
s_dx_gateway_association.test
  ]
}

resource "aws_vpn_connection" "test" {
  customer_gateway_idtomer_gateway.test.id
  outside_ip_address_type  = "PrivateIpv4"
  transit_gateway_id2_transit_gateway.test.id
  transport_transit_gateway_attachment_id = data.aws_ec2_transit_gateway_dx_gateway_attachment.test.id
  type1"

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_outsideAddressTypePublic(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
func_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_idomer_gateway.test.id
  outside_ip_address_type = "PublicIpv4"
  type"
  vpn_gateway_id = aws_vpn_gateway.test.id

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_staticRoutes(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
func
  tags = {
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
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_enableAcceleration(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  description = %[1]q
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_id = aws_customer_gateway.test.id
  transit_gateway_id  = aws_ec2_transit_gateway.test.id
  type = "ipsec.1"
  static_routes_only  = false
  enable_acceleration = true

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_ipv6(rName string, rBgpAsn int, localIpv6NetworkCidr string, remoteIpv6NetworkCidr string, tunnel1InsideIpv6Cidr string, tunnel2InsideIpv6Cidr string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  description = %[1]q
}

resource "aws_customer_gateway" "test" {
func_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_id = aws_customer_gateway.test.id
  transit_gateway_id  = aws_ec2_transit_gateway.test.id
  type = "ipsec.1"
  static_routes_only  = false
  enable_acceleration = false

  local_ipv6_network_cidr  = %[3]q
  remote_ipv6_network_cidr = %[4]q
  tunnel_inside_ip_version = "ipv6"

  tunnel1_inside_ipv6_cidr = %[5]q
  tunnel2_inside_ipv6_cidr = %[6]q

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn, localIpv6NetworkCidr, remoteIpv6NetworkCidr, tunnel1InsideIpv6Cidr, tunnel2InsideIpv6Cidr)
}


func testAccSiteVPNConnectionConfig_singleTunnelOptions(rName string, rBgpAsn int, psk string, tunnelCidr string) string {
	return fmt.Sprintf(`
funcgs = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"
  static_routes_only  = false

  tunnel1_inside_cidr[3]q
  tunnel1_preshared_key = %[4]q

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn, tunnelCidr, psk)
}


funcurn fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  description = %[1]q
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_id = aws_customer_gateway.test.id
  transit_gateway_id  = aws_ec2_transit_gateway.test.id
  type = aws_customer_gateway.test.type

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_transitGatewayID(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test1" {
func

resource "aws_ec2_transit_gateway" "test2" {
  description = "%[1]s-2"
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_id = aws_customer_gateway.test.id
  transit_gateway_id  = aws_ec2_transit_gateway.test1.id
  type = "ipsec.1"

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_transitGatewayIDUpdated(rName string, rBgpAsn int) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test1" {
  description = "%[1]s-1"
}

resource "aws_ec2_transit_gateway" "test2" {
  description = "%[1]s-2"
}
funcurce "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_id = aws_customer_gateway.test.id
  transit_gateway_id  = aws_ec2_transit_gateway.test2.id
  type = "ipsec.1"

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn)
}


func testAccSiteVPNConnectionConfig_tunnel1InsideCIDR(rName string, rBgpAsn int, tunnel1InsideCidr string, tunnel2InsideCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"
funcgs = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_id = aws_customer_gateway.test.id
  tunnel1_inside_cidr = %[3]q
  tunnel2_inside_cidr = %[4]q
  type = "ipsec.1"
  vpn_gateway_id_gateway.test.id

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn, tunnel1InsideCidr, tunnel2InsideCidr)
}


func testAccSiteVPNConnectionConfig_tunnel1InsideIPv6CIDR(rName string, rBgpAsn int, tunnel1InsideIpv6Cidr string, tunnel2InsideIpv6Cidr string) string {
	return fmt.Sprintf(`
resource "aws_ec2_transit_gateway" "test" {
  description = %[1]q
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
func
  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_idtomer_gateway.test.id
  transit_gateway_id2_transit_gateway.test.id
  tunnel_inside_ip_version = "ipv6"
  tunnel1_inside_ipv6_cidr = %[3]q
  tunnel2_inside_ipv6_cidr = %[4]q
  type1"

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn, tunnel1InsideIpv6Cidr, tunnel2InsideIpv6Cidr)
}


func testAccSiteVPNConnectionConfig_tunnel1PresharedKey(rName string, rBgpAsn int, tunnel1PresharedKey string, tunnel2PresharedKey string) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
func
  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_idws_customer_gateway.test.id
  tunnel1_preshared_key = %[3]q
  tunnel2_preshared_key = %[4]q
  typeipsec.1"
  vpn_gateway_idpn_gateway.test.id

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn, tunnel1PresharedKey, tunnel2PresharedKey)
}


func testAccSiteVPNConnectionConfig_tunnelOptions(
	rName string,
	rBgpAsn int,
	localIpv4NetworkCidr string,
	remoteIpv4NetworkCidr string,
	tunnel1 TunnelOptions,
	tunnel2 TunnelOptions,
) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
func

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"

  local_ipv4_network_cidr  = %[3]q
  remote_ipv4_network_cidr = %[4]q

  tunnel1_inside_cidr
  tunnel1_preshared_key%[6]q
  tunnel1_dpd_timeout_action
  tunnel1_dpd_timeout_seconds%[8]d
  tunnel1_enable_tunnel_lifecycle_control = %[9]t
  tunnel1_ike_versions
  tunnel1_phase1_dh_group_numbers= [%[11]s]
  tunnel1_phase1_encryption_algorithms[%[12]s]
  tunnel1_phase1_integrity_algorithms
  tunnel1_phase1_lifetime_seconds= %[14]d
  tunnel1_phase2_dh_group_numbers= [%[15]s]
  tunnel1_phase2_encryption_algorithms[%[16]s]
  tunnel1_phase2_integrity_algorithms
funcnnel1_rekey_fuzz_percentage  = %[19]d
  tunnel1_rekey_margin_time_seconds
  tunnel1_replay_window_size
  tunnel1_startup_action[22]q

  tunnel2_inside_cidr
  tunnel2_preshared_key%[24]q
  tunnel2_dpd_timeout_action
  tunnel2_dpd_timeout_seconds%[26]d
  tunnel2_enable_tunnel_lifecycle_control = %[27]t
  tunnel2_ike_versions
  tunnel2_phase1_dh_group_numbers= [%[29]s]
  tunnel2_phase1_encryption_algorithms[%[30]s]
  tunnel2_phase1_integrity_algorithms
  tunnel2_phase1_lifetime_seconds= %[32]d
  tunnel2_phase2_dh_group_numbers= [%[33]s]
  tunnel2_phase2_encryption_algorithms[%[34]s]
  tunnel2_phase2_integrity_algorithms
  tunnel2_phase2_lifetime_seconds= %[36]d
  tunnel2_rekey_fuzz_percentage  = %[37]d
  tunnel2_rekey_margin_time_seconds
  tunnel2_replay_window_size
  tunnel2_startup_action[40]q

  tags = {
me = %[1]q
  }
}
`,
rName,
rBgpAsn,
localIpv4NetworkCidr,
funcel1.tunnelCidr,
tunnel1.psk,
tunnel1.dpdTimeoutAction,
tunnel1.dpdTimeoutSeconds,
tunnel1.enableTunnelLifecycleControl,
tunnel1.ikeVersions,
tunnel1.phase1DhGroupNumbers,
tunnel1.phase1EncryptionAlgorithms,
tunnel1.phase1IntegrityAlgorithms,
tunnel1.phase1LifetimeSeconds,
tunnel1.phase2DhGroupNumbers,
tunnel1.phase2EncryptionAlgorithms,
tunnel1.phase2IntegrityAlgorithms,
tunnel1.phase2LifetimeSeconds,
tunnel1.rekeyFuzzPercentage,
tunnel1.rekeyMarginTimeSeconds,
tunnel1.replayWindowSize,
tunnel1.startupAction,
tunnel2.tunnelCidr,
tunnel2.psk,
tunnel2.dpdTimeoutAction,
tunnel2.dpdTimeoutSeconds,
tunnel2.enableTunnelLifecycleControl,
tunnel2.ikeVersions,
tunnel2.phase1DhGroupNumbers,
tunnel2.phase1EncryptionAlgorithms,
tunnel2.phase1IntegrityAlgorithms,
tunnel2.phase1LifetimeSeconds,
tunnel2.phase2DhGroupNumbers,
tunnel2.phase2EncryptionAlgorithms,
tunnel2.phase2IntegrityAlgorithms,
tunnel2.phase2LifetimeSeconds,
tunnel2.rekeyFuzzPercentage,
funcel2.replayWindowSize,
tunnel2.startupAction)
}


func testAccSiteVPNConnectionConfig_tags1(rName string, rBgpAsn int, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"
  static_routes_only  = true

  tags = {
3]q = %[4]q
  }
}
`, rName, rBgpAsn, tagKey1, tagValue1)
}


func testAccSiteVPNConnectionConfig_tags2(rName string, rBgpAsn int, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"
  static_routes_only  = true

  tags = {
3]q = %[4]q
5]q = %[6]q
  }
}
`, rName, rBgpAsn, tagKey1, tagValue1, tagKey2, tagValue2)
}


func testAccSiteVPNConnectionConfig_localRemoteIPv4CIDRs(rName string, rBgpAsn int, localIpv4Cidr string, remoteIpv4Cidr string) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
  ip_address = "178.0.0.1"
  type.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  vpn_gateway_id_gateway.test.id
  customer_gateway_id = aws_customer_gateway.test.id
  type = "ipsec.1"
  static_routes_only  = false

  local_ipv4_network_cidr  = %[3]q
  remote_ipv4_network_cidr = %[4]q

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn, localIpv4Cidr, remoteIpv4Cidr)
}


func testAccSiteVPNConnectionConfig_transitGatewayIDOrVPNGatewayID(rName string, rBgpAsn int, useTransitGateway bool) string {
	return fmt.Sprintf(`
resource "aws_vpn_gateway" "test" {
  tags = {
me = %[1]q
  }
}

resource "aws_ec2_transit_gateway" "test" {
  description = %[1]q
}

resource "aws_customer_gateway" "test" {
  bgp_asn%[2]d
funcpe.1"

  tags = {
me = %[1]q
  }
}

resource "aws_vpn_connection" "test" {
  customer_gateway_id = aws_customer_gateway.test.id
  transit_gateway_id  = %[3]t ? aws_ec2_transit_gateway.test.id : null
  vpn_gateway_id null : aws_vpn_gateway.test.id
  type = "ipsec.1"

  tags = {
me = %[1]q
  }
}
`, rName, rBgpAsn, useTransitGateway)
}

// Test our VPN tunnel config XML parsing
const testAccVPNTunnelInfoXML = `
<vpn_connection id="vpn-abc123">
  <ipsec_tunnel>
ustomer_gateway>
utside_address>
ress>22.22.22.22</ip_address>
outside_address>
nside_address>
ress>169.254.12.1</ip_address>
k_mask>255.255.255.252</network_mask>
k_cidr>30</network_cidr>
funcomer_gateway>
pn_gateway>
utside_address>
ress>2.2.2.2</ip_address>
outside_address>
nside_address>
ress>169.254.12.2</ip_address>
k_mask>255.255.255.252</network_mask>
k_cidr>30</network_cidr>
inside_address>

22</asn>
ime>32</hold_time>

vpn_gateway>
ke>
ed_key>SECOND_KEY</pre_shared_key>
ike>
  </ipsec_tunnel>
  <ipsec_tunnel>
ustomer_gateway>
utside_address>
ress>11.11.11.11</ip_address>
outside_address>
nside_address>
ress>169.254.11.1</ip_address>
k_mask>255.255.255.252</network_mask>
k_cidr>30</network_cidr>
inside_address>
customer_gateway>
pn_gateway>
utside_address>
ress>1.1.1.1</ip_address>
funce_address>
ress>168.254.11.2</ip_address>
k_mask>255.255.255.252</network_mask>
k_cidr>30</network_cidr>
inside_address>

11</asn>
ime>31</hold_time>

vpn_gateway>
ke>
ed_key>FIRST_KEY</pre_shared_key>
ike>
  </ipsec_tunnel>
</vpn_connection>
`
func