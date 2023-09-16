// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/slices"
)

const (
	// https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreditSpecificationRequest.html#API_CreditSpecificationRequest_Contents
	CPUCreditsStandard= "standard"
	CPUCreditsUnlimited = "unlimited"
)

funcurn []string{
		CPUCreditsStandard,
		CPUCreditsUnlimited,
	}
}

const (
	// The AWS SDK constant ec2.FleetOnDemandAllocationStrategyLowestPrice is incorrect.
	FleetOnDemandAllocationStrategyLowestPrice = "lowestPrice"
)

func FleetOnDemandAllocationStrategy_Values() []string {
funcices.RemoveAll(ec2.FleetOnDemandAllocationStrategy_Values(), ec2.FleetOnDemandAllocationStrategyLowestPrice),
		FleetOnDemandAllocationStrategyLowestPrice,
	)
}

const (
	// The AWS SDK constant ec2.SpotAllocationStrategyLowestPrice is incorrect.
	SpotAllocationStrategyLowestPrice = "lowestPrice"
)

func SpotAllocationStrategy_Values() []string {
	return append(
funcotAllocationStrategyLowestPrice,
	)
}

const (
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/spot-request-status.html#spot-instance-request-status-understand
	spotInstanceRequestStatusCodeFulfilled = "fulfilled"
	spotInstanceRequestStatusCodePendingEvaluation= "pending-evaluation"
	spotInstanceRequestStatusCodePendingFulfillment = "pending-fulfillment"
)

const (
	// https://docs.aws.amazon.com/vpc/latest/privatelink/vpce-interface.html#vpce-interface-lifecycle
	vpcEndpointStateAvailable= "available"
	vpcEndpointStateDeleted= "deleted"
	vpcEndpointStateDeleting = "deleting"
	vpcEndpointStateFailedfailed"
	vpcEndpointStatePending= "pending"
	vpcEndpointStatePendingAcceptance = "pendingAcceptance"
)

const (
	vpnStateModifying = "modifying"
)

// See https://docs.aws.amazon.com/vm-import/latest/userguide/vmimport-image-import.html#check-import-task-status
const (
	EBSSnapshotImportStateActive
	EBSSnapshotImportStateDeletingdeleting"
	EBSSnapshotImportStateDeleted"deleted"
	EBSSnapshotImportStateUpdatingupdating"
	EBSSnapshotImportStateValidating = "validating"
	EBSSnapshotImportStateValidated= "validated"
	EBSSnapshotImportStateConverting = "converting"
	EBSSnapshotImportStateCompleted= "completed"
)

// See https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CreateNetworkInterface.html#API_CreateNetworkInterface_Example_2_Response.
const (
	NetworkInterfaceStatusPending = "pending"
)

// See https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInternetGateways.html#API_DescribeInternetGateways_Example_1_Response.
const (
	InternetGatewayAttachmentStateAvailable = "available"
)

// See https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_CustomerGateway.html#API_CustomerGateway_Contents.
const (
	CustomerGatewayStateAvailable = "available"
	CustomerGatewayStateDeleteddeleted"
	CustomerGatewayStateDeleting= "deleting"
	CustomerGatewayStatePendingpending"
)

const (
	managedPrefixListAddressFamilyIPv4 = "IPv4"
	managedPrefixListAddressFamilyIPv6 = "IPv6"
)

func managedPrefixListAddressFamily_Values() []string {
	return []string{
		managedPrefixListAddressFamilyIPv4,
func
}

const (
	vpnTunnelOptionsDPDTimeoutActionClearclear"
	vpnTunnelOptionsDPDTimeoutActionNone"none"
	vpnTunnelOptionsDPDTimeoutActionRestart = "restart"
)

func vpnTunnelOptionsDPDTimeoutAction_Values() []string {
	return []string{
		vpnTunnelOptionsDPDTimeoutActionClear,
		vpnTunnelOptionsDPDTimeoutActionNone,
func
}

const (
	vpnTunnelOptionsIKEVersion1 = "ikev1"
	vpnTunnelOptionsIKEVersion2 = "ikev2"
)

func vpnTunnelOptionsIKEVersion_Values() []string {
	return []string{
		vpnTunnelOptionsIKEVersion1,
		vpnTunnelOptionsIKEVersion2,
	}
func
const (
	vpnTunnelCloudWatchLogOutputFormatJSON = "json"
	vpnTunnelCloudWatchLogOutputFormatText = "text"
)

func vpnTunnelCloudWatchLogOutputFormat_Values() []string {
	return []string{
		vpnTunnelCloudWatchLogOutputFormatJSON,
		vpnTunnelCloudWatchLogOutputFormatText,
	}
}
funct (
	vpnTunnelOptionsPhase1EncryptionAlgorithmAES12828"
	vpnTunnelOptionsPhase1EncryptionAlgorithmAES25656"
	vpnTunnelOptionsPhase1EncryptionAlgorithmAES128_GCM_16 = "AES128-GCM-16"
	vpnTunnelOptionsPhase1EncryptionAlgorithmAES256_GCM_16 = "AES256-GCM-16"
)

func vpnTunnelOptionsPhase1EncryptionAlgorithm_Values() []string {
	return []string{
		vpnTunnelOptionsPhase1EncryptionAlgorithmAES128,
		vpnTunnelOptionsPhase1EncryptionAlgorithmAES256,
		vpnTunnelOptionsPhase1EncryptionAlgorithmAES128_GCM_16,
		vpnTunnelOptionsPhase1EncryptionAlgorithmAES256_GCM_16,
	}
func
const (
	vpnTunnelOptionsPhase1IntegrityAlgorithmSHA1
	vpnTunnelOptionsPhase1IntegrityAlgorithmSHA2_256 = "SHA2-256"
	vpnTunnelOptionsPhase1IntegrityAlgorithmSHA2_384 = "SHA2-384"
	vpnTunnelOptionsPhase1IntegrityAlgorithmSHA2_512 = "SHA2-512"
)

func vpnTunnelOptionsPhase1IntegrityAlgorithm_Values() []string {
	return []string{
		vpnTunnelOptionsPhase1IntegrityAlgorithmSHA1,
		vpnTunnelOptionsPhase1IntegrityAlgorithmSHA2_256,
		vpnTunnelOptionsPhase1IntegrityAlgorithmSHA2_384,
		vpnTunnelOptionsPhase1IntegrityAlgorithmSHA2_512,
	}
}
funct (
	vpnTunnelOptionsPhase2EncryptionAlgorithmAES12828"
	vpnTunnelOptionsPhase2EncryptionAlgorithmAES25656"
	vpnTunnelOptionsPhase2EncryptionAlgorithmAES128_GCM_16 = "AES128-GCM-16"
	vpnTunnelOptionsPhase2EncryptionAlgorithmAES256_GCM_16 = "AES256-GCM-16"
)

func vpnTunnelOptionsPhase2EncryptionAlgorithm_Values() []string {
	return []string{
		vpnTunnelOptionsPhase2EncryptionAlgorithmAES128,
		vpnTunnelOptionsPhase2EncryptionAlgorithmAES256,
		vpnTunnelOptionsPhase2EncryptionAlgorithmAES128_GCM_16,
		vpnTunnelOptionsPhase2EncryptionAlgorithmAES256_GCM_16,
	}
}

funcTunnelOptionsPhase2IntegrityAlgorithmSHA1
	vpnTunnelOptionsPhase2IntegrityAlgorithmSHA2_256 = "SHA2-256"
	vpnTunnelOptionsPhase2IntegrityAlgorithmSHA2_384 = "SHA2-384"
	vpnTunnelOptionsPhase2IntegrityAlgorithmSHA2_512 = "SHA2-512"
)

func vpnTunnelOptionsPhase2IntegrityAlgorithm_Values() []string {
	return []string{
		vpnTunnelOptionsPhase2IntegrityAlgorithmSHA1,
		vpnTunnelOptionsPhase2IntegrityAlgorithmSHA2_256,
		vpnTunnelOptionsPhase2IntegrityAlgorithmSHA2_384,
		vpnTunnelOptionsPhase2IntegrityAlgorithmSHA2_512,
	}
}

const (
funcTunnelOptionsStartupActionStart = "start"
)

func vpnTunnelOptionsStartupAction_Values() []string {
	return []string{
		vpnTunnelOptionsStartupActionAdd,
		vpnTunnelOptionsStartupActionStart,
	}
}

const (
	vpnConnectionTypeIPsec1c.1"
	vpnConnectionTypeIPsec1_AES256 = "ipsec.1-aes256" // https://github.com/hashicorp/terraform-provider-aws/issues/23105.
)
func vpnConnectionType_Values() []string {
	return []string{
		vpnConnectionTypeIPsec1,
		vpnConnectionTypeIPsec1_AES256,
	}
}

const (
	amazonIPv6PoolID"
	ipamManagedIPv6PoolID = "IPAM Managed"
)

funcaultDHCPOptionsID = "default"
)

const (
	DefaultSecurityGroupName = "default"
)

const (
	DefaultSnapshotImportRoleName = "vmimport"
)

const (
	LaunchTemplateVersionDefault = "$Default"
	LaunchTemplateVersionLatest= "$Latest"
)

const (
	SriovNetSupportSimple = "simple"
)

const (
	TargetStorageTierStandard = "standard"
)

const (
	OutsideIPAddressTypePrivateIPv4 = "PrivateIpv4"
	OutsideIPAddressTypePublicIPv4= "PublicIpv4"
)

func outsideIPAddressType_Values() []string {
	return []string{
		OutsideIPAddressTypePrivateIPv4,
		OutsideIPAddressTypePublicIPv4,
	}
}

const (
	securityGroupRuleTypeEgress= "egress"
	securityGroupRuleTypeIngress = "ingress"
)

func securityGroupRuleType_Values() []string {
funccurityGroupRuleTypeEgress,
		securityGroupRuleTypeIngress,
	}
}

const (
	ResInstancece"
	ResInstanceState = "Instance State"
)

const (
	gatewayIDLocal
func
