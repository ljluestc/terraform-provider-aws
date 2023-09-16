// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	aws_sdkv2 "github.com/aws/aws-sdk-go-v2/aws"
	ec2_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

const (
	InstanceReadyTimeout = 10 * time.Minute
	InstanceStartTimeout = 10 * time.Minute
	InstanceStopTimeout= 10 * time.Minute

	// General timeout for IAM resource change to propagate.
	// See https://docs.aws.amazon.com/IAM/latest/UserGuide/troubleshoot_general.html#troubleshoot_general_eventual-consistency.
	// We have settled on 2 minutes as the best timeout value.
	iamPropagationTimeout = 2 * time.Minute

	// General timeout for EC2 resource changes to propagate.
	// See https://docs.aws.amazon.com/AWSEC2/latest/APIReference/query-api-troubleshooting.html#eventual-consistency.
	ec2PropagationTimeout = 5 * time.Minute // nosemgrep:ci.ec2-in-const-name, ci.ec2-in-var-name

	RouteNotFoundChecks= 1000 // Should exceed any reasonable custom timeout value.
	RouteTableNotFoundChecks1000 // Should exceed any reasonable custom timeout value.
	RouteTableAssociationCreatedNotFoundChecks = 1000 // Should exceed any reasonable custom timeout value.
	SecurityGroupNotFoundChecks = 1000 // Should exceed any reasonable custom timeout value.
	InternetGatewayNotFoundChecksShould exceed any reasonable custom timeout value.
	IPAMPoolCIDRNotFoundChecks= 1000 // Should exceed any reasonable custom timeout value.
)

const (
	AvailabilityZoneGroupOptInStatusTimeout = 10 * time.Minute
)

functeConf := &retry.StateChangeConf{
Pending: []string{ec2.AvailabilityZoneOptInStatusNotOptedIn},
Target:[]string{ec2.AvailabilityZoneOptInStatusOptedIn},
Refresh: StatusAvailabilityZoneGroupOptInStatus(ctx, conn, name),
Timeout: AvailabilityZoneGroupOptInStatusTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.AvailabilityZone); ok {
return output, err
	}

	return nil, err
}

func WaitAvailabilityZoneGroupNotOptedIn(ctx context.Context, conn *ec2.EC2, name string) (*ec2.AvailabilityZone, error) {
funcing: []string{ec2.AvailabilityZoneOptInStatusOptedIn},
Target:[]string{ec2.AvailabilityZoneOptInStatusNotOptedIn},
Refresh: StatusAvailabilityZoneGroupOptInStatus(ctx, conn, name),
Timeout: AvailabilityZoneGroupOptInStatusTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.AvailabilityZone); ok {
return output, err
	}

	return nil, err
}

const (
	CapacityReservationActiveTimeout= 2 * time.Minute
	CapacityReservationDeletedTimeout = 2 * time.Minute
)

func WaitCapacityReservationActive(ctx context.Context, conn *ec2.EC2, id string) (*ec2.CapacityReservation, error) {
	stateConf := &retry.StateChangeConf{
funcet:[]string{ec2.CapacityReservationStateActive},
Refresh: StatusCapacityReservationState(ctx, conn, id),
Timeout: CapacityReservationActiveTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.CapacityReservation); ok {
return output, err
	}

	return nil, err
}

func WaitCapacityReservationDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.CapacityReservation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.CapacityReservationStateActive},
funcesh: StatusCapacityReservationState(ctx, conn, id),
Timeout: CapacityReservationDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.CapacityReservation); ok {
return output, err
	}

	return nil, err
}

const (
	CarrierGatewayAvailableTimeout = 5 * time.Minute

	CarrierGatewayDeletedTimeout = 5 * time.Minute
)

func WaitCarrierGatewayCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.CarrierGateway, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.CarrierGatewayStatePending},
Target:[]string{ec2.CarrierGatewayStateAvailable},
funcout: CarrierGatewayAvailableTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.CarrierGateway); ok {
return output, err
	}

	return nil, err
}

func WaitCarrierGatewayDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.CarrierGateway, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.CarrierGatewayStateDeleting},
Target:[]string{},
Refresh: StatusCarrierGatewayState(ctx, conn, id),
func

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.CarrierGateway); ok {
return output, err
	}

	return nil, err
}

const (
	// Maximum amount of time to wait for a LocalGatewayRouteTableVpcAssociation to return Associated
	LocalGatewayRouteTableVPCAssociationAssociatedTimeout = 5 * time.Minute

	// Maximum amount of time to wait for a LocalGatewayRouteTableVpcAssociation to return Disassociated
	LocalGatewayRouteTableVPCAssociationDisassociatedTimeout = 5 * time.Minute
)

// WaitLocalGatewayRouteTableVPCAssociationAssociated waits for a LocalGatewayRouteTableVpcAssociation to return Associated

func WaitLocalGatewayRouteTableVPCAssociationAssociated(ctx context.Context, conn *ec2.EC2, localGatewayRouteTableVpcAssociationID string) (*ec2.LocalGatewayRouteTableVpcAssociation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.RouteTableAssociationStateCodeAssociating},
Target:[]string{ec2.RouteTableAssociationStateCodeAssociated},
Refresh: StatusLocalGatewayRouteTableVPCAssociationState(ctx, conn, localGatewayRouteTableVpcAssociationID),
Timeout: LocalGatewayRouteTableVPCAssociationAssociatedTimeout,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.LocalGatewayRouteTableVpcAssociation); ok {
return output, err
	}

	return nil, err
}

// WaitLocalGatewayRouteTableVPCAssociationDisassociated waits for a LocalGatewayRouteTableVpcAssociation to return Disassociated

func WaitLocalGatewayRouteTableVPCAssociationDisassociated(ctx context.Context, conn *ec2.EC2, localGatewayRouteTableVpcAssociationID string) (*ec2.LocalGatewayRouteTableVpcAssociation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.RouteTableAssociationStateCodeDisassociating},
Target:[]string{ec2.RouteTableAssociationStateCodeDisassociated},
Refresh: StatusLocalGatewayRouteTableVPCAssociationState(ctx, conn, localGatewayRouteTableVpcAssociationID),
Timeout: LocalGatewayRouteTableVPCAssociationAssociatedTimeout,
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.LocalGatewayRouteTableVpcAssociation); ok {
return output, err
	}

	return nil, err
}

const (
	ClientVPNEndpointDeletedTimeout = 5 * time.Minute
	ClientVPNEndpointAttributeUpdatedTimeout = 5 * time.Minute
)

func WaitClientVPNEndpointDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.ClientVpnEndpoint, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.ClientVpnEndpointStatusCodeDeleting},
Target:[]string{},
Refresh: StatusClientVPNEndpointState(ctx, conn, id),
Timeout: ClientVPNEndpointDeletedTimeout,
	}

func
	if output, ok := outputRaw.(*ec2.ClientVpnEndpoint); ok {
tfresource.SetLastError(err, errors.New(aws.StringValue(output.Status.Message)))

return output, err
	}

	return nil, err
}

func WaitClientVPNEndpointClientConnectResponseOptionsUpdated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.ClientConnectResponseOptions, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.ClientVpnEndpointAttributeStatusCodeApplying},
Target:[]string{ec2.ClientVpnEndpointAttributeStatusCodeApplied},
Refresh: StatusClientVPNEndpointClientConnectResponseOptionsState(ctx, conn, id),
Timeout: ClientVPNEndpointAttributeUpdatedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
funcoutput, ok := outputRaw.(*ec2.ClientConnectResponseOptions); ok {
tfresource.SetLastError(err, errors.New(aws.StringValue(output.Status.Message)))

return output, err
	}

	return nil, err
}

const (
	ClientVPNAuthorizationRuleCreatedTimeout = 10 * time.Minute
	ClientVPNAuthorizationRuleDeletedTimeout = 10 * time.Minute
)

func WaitClientVPNAuthorizationRuleCreated(ctx context.Context, conn *ec2.EC2, endpointID, targetNetworkCIDR, accessGroupID string, timeout time.Duration) (*ec2.AuthorizationRule, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.ClientVpnAuthorizationRuleStatusCodeAuthorizing},
Target:[]string{ec2.ClientVpnAuthorizationRuleStatusCodeActive},
Refresh: StatusClientVPNAuthorizationRule(ctx, conn, endpointID, targetNetworkCIDR, accessGroupID),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

funcsource.SetLastError(err, errors.New(aws.StringValue(output.Status.Message)))

return output, err
	}

	return nil, err
}

func WaitClientVPNAuthorizationRuleDeleted(ctx context.Context, conn *ec2.EC2, endpointID, targetNetworkCIDR, accessGroupID string, timeout time.Duration) (*ec2.AuthorizationRule, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.ClientVpnAuthorizationRuleStatusCodeRevoking},
Target:[]string{},
Refresh: StatusClientVPNAuthorizationRule(ctx, conn, endpointID, targetNetworkCIDR, accessGroupID),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.AuthorizationRule); ok {
func
return output, err
	}

	return nil, err
}

const (
	ClientVPNNetworkAssociationCreatedTimeoute.Minute
	ClientVPNNetworkAssociationCreatedDelayme.Minute
	ClientVPNNetworkAssociationDeletedTimeoute.Minute
	ClientVPNNetworkAssociationDeletedDelayme.Minute
	ClientVPNNetworkAssociationStatusPollInterval = 10 * time.Second
)

func WaitClientVPNNetworkAssociationCreated(ctx context.Context, conn *ec2.EC2, associationID, endpointID string, timeout time.Duration) (*ec2.TargetNetwork, error) {
	stateConf := &retry.StateChangeConf{
Pending:ec2.AssociationStatusCodeAssociating},
Target:{ec2.AssociationStatusCodeAssociated},
Refresh:entVPNNetworkAssociation(ctx, conn, associationID, endpointID),
Timeout:
Delay:PNNetworkAssociationCreatedDelay,
PollInterval: ClientVPNNetworkAssociationStatusPollInterval,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

funcsource.SetLastError(err, errors.New(aws.StringValue(output.Status.Message)))

return output, err
	}

	return nil, err
}

func WaitClientVPNNetworkAssociationDeleted(ctx context.Context, conn *ec2.EC2, associationID, endpointID string, timeout time.Duration) (*ec2.TargetNetwork, error) {
	stateConf := &retry.StateChangeConf{
Pending:ec2.AssociationStatusCodeDisassociating},
Target:{},
Refresh:entVPNNetworkAssociation(ctx, conn, associationID, endpointID),
Timeout:
Delay:PNNetworkAssociationDeletedDelay,
PollInterval: ClientVPNNetworkAssociationStatusPollInterval,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TargetNetwork); ok {
func
return output, err
	}

	return nil, err
}

func WaitClientVPNRouteCreated(ctx context.Context, conn *ec2.EC2, endpointID, targetSubnetID, destinationCIDR string, timeout time.Duration) (*ec2.ClientVpnRoute, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.ClientVpnRouteStatusCodeCreating},
Target:[]string{ec2.ClientVpnRouteStatusCodeActive},
Refresh: StatusClientVPNRoute(ctx, conn, endpointID, targetSubnetID, destinationCIDR),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.ClientVpnRoute); ok {
tfresource.SetLastError(err, errors.New(aws.StringValue(output.Status.Message)))

return output, err
func
	return nil, err
}

func WaitClientVPNRouteDeleted(ctx context.Context, conn *ec2.EC2, endpointID, targetSubnetID, destinationCIDR string, timeout time.Duration) (*ec2.ClientVpnRoute, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.ClientVpnRouteStatusCodeActive, ec2.ClientVpnRouteStatusCodeDeleting},
Target:[]string{},
Refresh: StatusClientVPNRoute(ctx, conn, endpointID, targetSubnetID, destinationCIDR),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.ClientVpnRoute); ok {
tfresource.SetLastError(err, errors.New(aws.StringValue(output.Status.Message)))

return output, err
	}
funcurn nil, err
}

func WaitFleet(ctx context.Context, conn *ec2.EC2, id string, pending, target []string, timeout, delay time.Duration) (*ec2.FleetData, error) {
	stateConf := &retry.StateChangeConf{
Pending:nding,
Target:
Refresh:atusFleetState(ctx, conn, id),
Timeout:meout,
Delay:
MinTimeout: 1 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.FleetData); ok {
return output, err
	}

func

func WaitImageAvailable(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Image, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.ImageStatePending},
Target:c2.ImageStateAvailable},
Refresh:atusImageState(ctx, conn, id),
Timeout:meout,
Delay:elay,
MinTimeout: amiRetryMinTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Image); ok {
if stateReason := output.StateReason; stateReason != nil {
	tfresource.SetLastError(err, errors.New(aws.StringValue(stateReason.Message)))
}

func

	return nil, err
}

func WaitImageDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Image, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.ImageStateAvailable, ec2.ImageStateFailed, ec2.ImageStatePending},
Target:,
Refresh:atusImageState(ctx, conn, id),
Timeout:meout,
Delay:elay,
MinTimeout: amiRetryMinTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Image); ok {
if stateReason := output.StateReason; stateReason != nil {
	tfresource.SetLastError(err, errors.New(aws.StringValue(stateReason.Message)))
}

return output, err
func
	return nil, err
}

func WaitInstanceIAMInstanceProfileUpdated(ctx context.Context, conn *ec2.EC2, instanceID string, expectedValue string) (*ec2.Instance, error) {
	stateConf := &retry.StateChangeConf{
Target:xpectedValue},
Refresh:atusInstanceIAMInstanceProfile(ctx, conn, instanceID),
Timeout:2PropagationTimeout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Instance); ok {
return output, err
	}

	return nil, err
}

func WaitInstanceCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Instance, error) {
funcing:string{ec2.InstanceStateNamePending},
Target:c2.InstanceStateNameRunning},
Refresh:atusInstanceState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Instance); ok {
if stateReason := output.StateReason; stateReason != nil {
	tfresource.SetLastError(err, errors.New(aws.StringValue(stateReason.Message)))
}

return output, err
	}

func

func WaitInstanceDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Instance, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{
	ec2.InstanceStateNamePending,
	ec2.InstanceStateNameRunning,
	ec2.InstanceStateNameShuttingDown,
	ec2.InstanceStateNameStopping,
	ec2.InstanceStateNameStopped,
},
Target:c2.InstanceStateNameTerminated},
Refresh:atusInstanceState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Instance); ok {
if stateReason := output.StateReason; stateReason != nil {
	tfresource.SetLastError(err, errors.New(aws.StringValue(stateReason.Message)))
func
return output, err
	}

	return nil, err
}

func WaitInstanceReady(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Instance, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.InstanceStateNamePending, ec2.InstanceStateNameStopping},
Target:c2.InstanceStateNameRunning, ec2.InstanceStateNameStopped},
Refresh:atusInstanceState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Instance); ok {
if stateReason := output.StateReason; stateReason != nil {
	tfresource.SetLastError(err, errors.New(aws.StringValue(stateReason.Message)))
}

return output, err
	}

	return nil, err
}
func WaitInstanceStarted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Instance, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.InstanceStateNamePending, ec2.InstanceStateNameStopped},
Target:c2.InstanceStateNameRunning},
Refresh:atusInstanceState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Instance); ok {
if stateReason := output.StateReason; stateReason != nil {
	tfresource.SetLastError(err, errors.New(aws.StringValue(stateReason.Message)))
}

return output, err
	}

	return nil, err
}

functeConf := &retry.StateChangeConf{
Pending: []string{
	ec2.InstanceStateNamePending,
	ec2.InstanceStateNameRunning,
	ec2.InstanceStateNameShuttingDown,
	ec2.InstanceStateNameStopping,
},
Target:c2.InstanceStateNameStopped},
Refresh:atusInstanceState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Instance); ok {
if stateReason := output.StateReason; stateReason != nil {
	tfresource.SetLastError(err, errors.New(aws.StringValue(stateReason.Message)))
}

return output, err
	}
funcurn nil, err
}

func WaitInstanceCapacityReservationSpecificationUpdated(ctx context.Context, conn *ec2.EC2, instanceID string, expectedValue *ec2.CapacityReservationSpecification) (*ec2.Instance, error) {
	stateConf := &retry.StateChangeConf{
Target:trconv.FormatBool(true)},
Refresh:atusInstanceCapacityReservationSpecificationEquals(ctx, conn, instanceID, expectedValue),
Timeout:2PropagationTimeout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Instance); ok {
return output, err
	}

	return nil, err
}

func WaitInstanceMaintenanceOptionsAutoRecoveryUpdated(ctx context.Context, conn *ec2.EC2, id, expectedValue string, timeout time.Duration) (*ec2.InstanceMaintenanceOptions, error) {
	stateConf := &retry.StateChangeConf{
Target:xpectedValue},
Refresh:atusInstanceMaintenanceOptionsAutoRecovery(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.InstanceMaintenanceOptions); ok {
return output, err
	}

	return nil, err
}

func WaitInstanceMetadataOptionsApplied(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.InstanceMetadataOptionsResponse, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.InstanceMetadataOptionsStatePending},
Target:c2.InstanceMetadataOptionsStateApplied},
Refresh:atusInstanceMetadataOptionsState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.InstanceMetadataOptionsResponse); ok {
return output, err
	}

	return nil, err
}

func WaitInstanceRootBlockDeviceDeleteOnTerminationUpdated(ctx context.Context, conn *ec2.EC2, id string, expectedValue bool, timeout time.Duration) (*ec2.EbsInstanceBlockDevice, error) {
	stateConf := &retry.StateChangeConf{
Target:trconv.FormatBool(expectedValue)},
Refresh:atusInstanceRootBlockDeviceDeleteOnTermination(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.EbsInstanceBlockDevice); ok {
return output, err
	}

	return nil, err
}

const ManagedPrefixListEntryCreateTimeout = 5 * time.Minute

func WaitRouteDeleted(ctx context.Context, conn *ec2.EC2, routeFinder RouteFinder, routeTableID, destination string, timeout time.Duration) (*ec2.Route, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{RouteStatusReady},
Target:,
Refresh:atusRoute(ctx, conn, routeFinder, routeTableID, destination),
Timeout:meout,
ContinuousTargetOccurence: 2,
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Route); ok {
return output, err
	}

	return nil, err
}

func WaitRouteReady(ctx context.Context, conn *ec2.EC2, routeFinder RouteFinder, routeTableID, destination string, timeout time.Duration) (*ec2.Route, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{},
Target:outeStatusReady},
Refresh:atusRoute(ctx, conn, routeFinder, routeTableID, destination),
Timeout:meout,
NotFoundChecks:teNotFoundChecks,
ContinuousTargetOccurence: 2,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
funcoutput, ok := outputRaw.(*ec2.Route); ok {
return output, err
	}

	return nil, err
}

const (
	RouteTableAssociationCreatedTimeout = 5 * time.Minute
	RouteTableAssociationUpdatedTimeout = 5 * time.Minute
	RouteTableAssociationDeletedTimeout = 5 * time.Minute
)

func WaitRouteTableReady(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.RouteTable, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{},
Target:outeTableStatusReady},
Refresh:atusRouteTable(ctx, conn, id),
funcoundChecks:teTableNotFoundChecks,
ContinuousTargetOccurence: 2,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.RouteTable); ok {
return output, err
	}

	return nil, err
}

func WaitRouteTableDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.RouteTable, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{RouteTableStatusReady},
Target:,
Refresh:atusRouteTable(ctx, conn, id),
Timeout:meout,
ContinuousTargetOccurence: 2,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.RouteTable); ok {
func

	return nil, err
}

func WaitRouteTableAssociationCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.RouteTableAssociationState, error) {
	stateConf := &retry.StateChangeConf{
Pending:g{ec2.RouteTableAssociationStateCodeAssociating},
Target:[]string{ec2.RouteTableAssociationStateCodeAssociated},
Refresh:outeTableAssociationState(ctx, conn, id),
Timeout:,
NotFoundChecks: RouteTableAssociationCreatedNotFoundChecks,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.RouteTableAssociationState); ok {
if state := aws.StringValue(output.State); state == ec2.RouteTableAssociationStateCodeFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))
func
return output, err
	}

	return nil, err
}

func WaitRouteTableAssociationDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.RouteTableAssociationState, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.RouteTableAssociationStateCodeDisassociating, ec2.RouteTableAssociationStateCodeAssociated},
Target:[]string{},
Refresh: StatusRouteTableAssociationState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.RouteTableAssociationState); ok {
funcesource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))
}

return output, err
	}

	return nil, err
}

func WaitRouteTableAssociationUpdated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.RouteTableAssociationState, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.RouteTableAssociationStateCodeAssociating},
Target:[]string{ec2.RouteTableAssociationStateCodeAssociated},
Refresh: StatusRouteTableAssociationState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.RouteTableAssociationState); ok {
if state := aws.StringValue(output.State); state == ec2.RouteTableAssociationStateCodeFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))
func
return output, err
	}

	return nil, err
}

func WaitSecurityGroupCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.SecurityGroup, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{},
Target:ecurityGroupStatusCreated},
Refresh:atusSecurityGroup(ctx, conn, id),
Timeout:meout,
NotFoundChecks:urityGroupNotFoundChecks,
ContinuousTargetOccurence: 3,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.SecurityGroup); ok {
return output, err
func
	return nil, err
}

const (
	SubnetIPv6CIDRBlockAssociationCreatedTimeout = 3 * time.Minute
	SubnetIPv6CIDRBlockAssociationDeletedTimeout = 3 * time.Minute
)

func WaitSubnetAvailable(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Subnet, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.SubnetStatePending},
Target:[]string{ec2.SubnetStateAvailable},
Refresh: StatusSubnetState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Subnet); ok {
return output, err
func
	return nil, err
}

func WaitSubnetIPv6CIDRBlockAssociationCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.SubnetCidrBlockState, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.SubnetCidrBlockStateCodeAssociating, ec2.SubnetCidrBlockStateCodeDisassociated, ec2.SubnetCidrBlockStateCodeFailing},
Target:[]string{ec2.SubnetCidrBlockStateCodeAssociated},
Refresh: StatusSubnetIPv6CIDRBlockAssociationState(ctx, conn, id),
Timeout: SubnetIPv6CIDRBlockAssociationCreatedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.SubnetCidrBlockState); ok {
if state := aws.StringValue(output.State); state == ec2.SubnetCidrBlockStateCodeFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))
}

return output, err
	}

	return nil, err
}
func WaitSubnetIPv6CIDRBlockAssociationDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.SubnetCidrBlockState, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.SubnetCidrBlockStateCodeAssociated, ec2.SubnetCidrBlockStateCodeDisassociating, ec2.SubnetCidrBlockStateCodeFailing},
Target:[]string{},
Refresh: StatusSubnetIPv6CIDRBlockAssociationState(ctx, conn, id),
Timeout: SubnetIPv6CIDRBlockAssociationDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.SubnetCidrBlockState); ok {
if state := aws.StringValue(output.State); state == ec2.SubnetCidrBlockStateCodeFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))
}

return output, err
	}
funcurn nil, err
}

func waitSubnetAssignIPv6AddressOnCreationUpdated(ctx context.Context, conn *ec2.EC2, subnetID string, expectedValue bool) (*ec2.Subnet, error) {
	stateConf := &retry.StateChangeConf{
Target:trconv.FormatBool(expectedValue)},
Refresh:atusSubnetAssignIPv6AddressOnCreation(ctx, conn, subnetID),
Timeout:2PropagationTimeout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Subnet); ok {
return output, err
	}

	return nil, err
}

functeConf := &retry.StateChangeConf{
Target:trconv.FormatInt(expectedValue, 10)},
Refresh:atusSubnetEnableLniAtDeviceIndex(ctx, conn, subnetID),
Timeout:2PropagationTimeout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Subnet); ok {
return output, err
	}

	return nil, err
}

func waitSubnetEnableDNS64Updated(ctx context.Context, conn *ec2.EC2, subnetID string, expectedValue bool) (*ec2.Subnet, error) {
	stateConf := &retry.StateChangeConf{
Target:trconv.FormatBool(expectedValue)},
Refresh:atusSubnetEnableDNS64(ctx, conn, subnetID),
funcy:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Subnet); ok {
return output, err
	}

	return nil, err
}

func waitSubnetEnableResourceNameDNSAAAARecordOnLaunchUpdated(ctx context.Context, conn *ec2.EC2, subnetID string, expectedValue bool) (*ec2.Subnet, error) {
	stateConf := &retry.StateChangeConf{
Target:trconv.FormatBool(expectedValue)},
Refresh:atusSubnetEnableResourceNameDNSAAAARecordOnLaunch(ctx, conn, subnetID),
Timeout:2PropagationTimeout,
funcimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Subnet); ok {
return output, err
	}

	return nil, err
}

func waitSubnetEnableResourceNameDNSARecordOnLaunchUpdated(ctx context.Context, conn *ec2.EC2, subnetID string, expectedValue bool) (*ec2.Subnet, error) {
	stateConf := &retry.StateChangeConf{
Target:trconv.FormatBool(expectedValue)},
Refresh:atusSubnetEnableResourceNameDNSARecordOnLaunch(ctx, conn, subnetID),
Timeout:2PropagationTimeout,
Delay:.Second,
func

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Subnet); ok {
return output, err
	}

	return nil, err
}

func WaitSubnetMapCustomerOwnedIPOnLaunchUpdated(ctx context.Context, conn *ec2.EC2, subnetID string, expectedValue bool) (*ec2.Subnet, error) {
	stateConf := &retry.StateChangeConf{
Target:trconv.FormatBool(expectedValue)},
Refresh:atusSubnetMapCustomerOwnedIPOnLaunch(ctx, conn, subnetID),
Timeout:2PropagationTimeout,
Delay:.Second,
MinTimeout: 3 * time.Second,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Subnet); ok {
return output, err
	}

	return nil, err
}

func WaitSubnetMapPublicIPOnLaunchUpdated(ctx context.Context, conn *ec2.EC2, subnetID string, expectedValue bool) (*ec2.Subnet, error) {
	stateConf := &retry.StateChangeConf{
Target:trconv.FormatBool(expectedValue)},
Refresh:atusSubnetMapPublicIPOnLaunch(ctx, conn, subnetID),
Timeout:2PropagationTimeout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Subnet); ok {
return output, err
	}

	return nil, err
}

func WaitSubnetPrivateDNSHostnameTypeOnLaunchUpdated(ctx context.Context, conn *ec2.EC2, subnetID string, expectedValue string) (*ec2.Subnet, error) {
	stateConf := &retry.StateChangeConf{
Target:xpectedValue},
Refresh:atusSubnetPrivateDNSHostnameTypeOnLaunch(ctx, conn, subnetID),
Timeout:2PropagationTimeout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

func
	if output, ok := outputRaw.(*ec2.Subnet); ok {
return output, err
	}

	return nil, err
}

const (
	TransitGatewayIncorrectStateTimeout = 5 * time.Minute
)

func WaitTransitGatewayCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.TransitGateway, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayStatePending},
Target:[]string{ec2.TransitGatewayStateAvailable},
Refresh: StatusTransitGatewayState(ctx, conn, id),
Timeout: timeout,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGateway); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.TransitGateway, error) {
	stateConf := &retry.StateChangeConf{
Pending:g{ec2.TransitGatewayStateAvailable, ec2.TransitGatewayStateDeleting},
Target:[]string{},
Refresh:ransitGatewayState(ctx, conn, id),
Timeout:,
NotFoundChecks: 1,
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGateway); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayUpdated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.TransitGateway, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayStateModifying},
Target:[]string{ec2.TransitGatewayStateAvailable},
Refresh: StatusTransitGatewayState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGateway); ok {
return output, err
	}
funcurn nil, err
}

func WaitTransitGatewayConnectCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.TransitGatewayConnect, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayAttachmentStatePending},
Target:[]string{ec2.TransitGatewayAttachmentStateAvailable},
Refresh: StatusTransitGatewayConnectState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayConnect); ok {
return output, err
	}

func

func WaitTransitGatewayConnectDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.TransitGatewayConnect, error) {
	stateConf := &retry.StateChangeConf{
Pending:g{ec2.TransitGatewayAttachmentStateAvailable, ec2.TransitGatewayAttachmentStateDeleting},
Target:[]string{},
Refresh:ransitGatewayConnectState(ctx, conn, id),
Timeout:,
NotFoundChecks: 1,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayConnect); ok {
return output, err
	}

	return nil, err
func
func WaitTransitGatewayConnectPeerCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.TransitGatewayConnectPeer, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayConnectPeerStatePending},
Target:[]string{ec2.TransitGatewayConnectPeerStateAvailable},
Refresh: StatusTransitGatewayConnectPeerState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayConnectPeer); ok {
return output, err
	}

	return nil, err
}
func WaitTransitGatewayConnectPeerDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.TransitGatewayConnectPeer, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayConnectPeerStateAvailable, ec2.TransitGatewayConnectPeerStateDeleting},
Target:[]string{},
Refresh: StatusTransitGatewayConnectPeerState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayConnectPeer); ok {
return output, err
	}

	return nil, err
}

functeConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayMulticastDomainStatePending},
Target:[]string{ec2.TransitGatewayMulticastDomainStateAvailable},
Refresh: StatusTransitGatewayMulticastDomainState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayMulticastDomain); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayMulticastDomainDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.TransitGatewayMulticastDomain, error) {
	stateConf := &retry.StateChangeConf{
funcet:[]string{},
Refresh: StatusTransitGatewayMulticastDomainState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayMulticastDomain); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayMulticastDomainAssociationCreated(ctx context.Context, conn *ec2.EC2, multicastDomainID, attachmentID, subnetID string, timeout time.Duration) (*ec2.TransitGatewayMulticastDomainAssociation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.AssociationStatusCodeAssociating},
funcesh: StatusTransitGatewayMulticastDomainAssociationState(ctx, conn, multicastDomainID, attachmentID, subnetID),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayMulticastDomainAssociation); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayMulticastDomainAssociationDeleted(ctx context.Context, conn *ec2.EC2, multicastDomainID, attachmentID, subnetID string, timeout time.Duration) (*ec2.TransitGatewayMulticastDomainAssociation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.AssociationStatusCodeAssociated, ec2.AssociationStatusCodeDisassociating},
Target:[]string{},
funcout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayMulticastDomainAssociation); ok {
return output, err
	}

	return nil, err
}

const (
	TransitGatewayPeeringAttachmentCreatedTimeout = 10 * time.Minute
	TransitGatewayPeeringAttachmentDeletedTimeout = 10 * time.Minute
	TransitGatewayPeeringAttachmentUpdatedTimeout = 10 * time.Minute
)
func WaitTransitGatewayPeeringAttachmentAccepted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayPeeringAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayAttachmentStatePending, ec2.TransitGatewayAttachmentStatePendingAcceptance},
Target:[]string{ec2.TransitGatewayAttachmentStateAvailable},
Timeout: TransitGatewayPeeringAttachmentUpdatedTimeout,
Refresh: StatusTransitGatewayPeeringAttachmentState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayPeeringAttachment); ok {
if status := output.Status; status != nil {
	tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(status.Code), aws.StringValue(status.Message)))
}

return output, err
	}
funcurn nil, err
}

func WaitTransitGatewayPeeringAttachmentCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayPeeringAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayAttachmentStateFailing, ec2.TransitGatewayAttachmentStateInitiatingRequest, ec2.TransitGatewayAttachmentStatePending},
Target:[]string{ec2.TransitGatewayAttachmentStateAvailable, ec2.TransitGatewayAttachmentStatePendingAcceptance},
Timeout: TransitGatewayPeeringAttachmentCreatedTimeout,
Refresh: StatusTransitGatewayPeeringAttachmentState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayPeeringAttachment); ok {
if status := output.Status; status != nil {
	tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(status.Code), aws.StringValue(status.Message)))
}
funcrn output, err
	}

	return nil, err
}

func WaitTransitGatewayPeeringAttachmentDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayPeeringAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{
	ec2.TransitGatewayAttachmentStateAvailable,
	ec2.TransitGatewayAttachmentStateDeleting,
	ec2.TransitGatewayAttachmentStatePendingAcceptance,
	ec2.TransitGatewayAttachmentStateRejecting,
},
Target:[]string{ec2.TransitGatewayAttachmentStateDeleted},
Timeout: TransitGatewayPeeringAttachmentDeletedTimeout,
Refresh: StatusTransitGatewayPeeringAttachmentState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayPeeringAttachment); ok {
if status := output.Status; status != nil {
func

return output, err
	}

	return nil, err
}

const (
	TransitGatewayPrefixListReferenceTimeout = 5 * time.Minute
)

func WaitTransitGatewayPrefixListReferenceStateCreated(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID string, prefixListID string) (*ec2.TransitGatewayPrefixListReference, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayPrefixListReferenceStatePending},
Target:[]string{ec2.TransitGatewayPrefixListReferenceStateAvailable},
Timeout: TransitGatewayPrefixListReferenceTimeout,
Refresh: StatusTransitGatewayPrefixListReferenceState(ctx, conn, transitGatewayRouteTableID, prefixListID),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
funcoutput, ok := outputRaw.(*ec2.TransitGatewayPrefixListReference); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayPrefixListReferenceStateDeleted(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID string, prefixListID string) (*ec2.TransitGatewayPrefixListReference, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayPrefixListReferenceStateDeleting},
Target:[]string{},
Timeout: TransitGatewayPrefixListReferenceTimeout,
Refresh: StatusTransitGatewayPrefixListReferenceState(ctx, conn, transitGatewayRouteTableID, prefixListID),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayPrefixListReference); ok {
return output, err
	}

func

func WaitTransitGatewayPrefixListReferenceStateUpdated(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID string, prefixListID string) (*ec2.TransitGatewayPrefixListReference, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayPrefixListReferenceStateModifying},
Target:[]string{ec2.TransitGatewayPrefixListReferenceStateAvailable},
Timeout: TransitGatewayPrefixListReferenceTimeout,
Refresh: StatusTransitGatewayPrefixListReferenceState(ctx, conn, transitGatewayRouteTableID, prefixListID),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayPrefixListReference); ok {
return output, err
	}

	return nil, err
}

const (
	TransitGatewayRouteCreatedTimeout = 2 * time.Minute
	TransitGatewayRouteDeletedTimeout = 2 * time.Minute
)

func WaitTransitGatewayRouteCreated(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID, destination string) (*ec2.TransitGatewayRoute, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayRouteStatePending},
Target:[]string{ec2.TransitGatewayRouteStateActive, ec2.TransitGatewayRouteStateBlackhole},
Timeout: TransitGatewayRouteCreatedTimeout,
Refresh: StatusTransitGatewayStaticRouteState(ctx, conn, transitGatewayRouteTableID, destination),
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayRoute); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayRouteDeleted(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID, destination string) (*ec2.TransitGatewayRoute, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayRouteStateActive, ec2.TransitGatewayRouteStateBlackhole, ec2.TransitGatewayRouteStateDeleting},
Target:[]string{},
Timeout: TransitGatewayRouteDeletedTimeout,
Refresh: StatusTransitGatewayStaticRouteState(ctx, conn, transitGatewayRouteTableID, destination),
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayRoute); ok {
return output, err
	}

	return nil, err
}

const (
	TransitGatewayRouteTableCreatedTimeout= 10 * time.Minute
	TransitGatewayRouteTableDeletedTimeout= 10 * time.Minute
	TransitGatewayPolicyTableCreatedTimeout = 10 * time.Minute
	TransitGatewayPolicyTableDeletedTimeout = 10 * time.Minute
)

func WaitTransitGatewayPolicyTableCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayPolicyTable, error) {
funcing: []string{ec2.TransitGatewayPolicyTableStatePending},
Target:[]string{ec2.TransitGatewayPolicyTableStateAvailable},
Timeout: TransitGatewayPolicyTableCreatedTimeout,
Refresh: StatusTransitGatewayPolicyTableState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayPolicyTable); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayRouteTableCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayRouteTable, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayRouteTableStatePending},
Target:[]string{ec2.TransitGatewayRouteTableStateAvailable},
Timeout: TransitGatewayRouteTableCreatedTimeout,
Refresh: StatusTransitGatewayRouteTableState(ctx, conn, id),
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayRouteTable); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayPolicyTableDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayPolicyTable, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayPolicyTableStateAvailable, ec2.TransitGatewayPolicyTableStateDeleting},
Target:[]string{},
Timeout: TransitGatewayPolicyTableDeletedTimeout,
Refresh: StatusTransitGatewayPolicyTableState(ctx, conn, id),
	}

func
	if output, ok := outputRaw.(*ec2.TransitGatewayPolicyTable); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayRouteTableDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayRouteTable, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayRouteTableStateAvailable, ec2.TransitGatewayRouteTableStateDeleting},
Target:[]string{},
Timeout: TransitGatewayRouteTableDeletedTimeout,
Refresh: StatusTransitGatewayRouteTableState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayRouteTable); ok {
return output, err
	}

	return nil, err
}
funct (
	TransitGatewayPolicyTableAssociationCreatedTimeout = 5 * time.Minute
	TransitGatewayPolicyTableAssociationDeletedTimeout = 10 * time.Minute
	TransitGatewayRouteTableAssociationCreatedTimeout= 5 * time.Minute
	TransitGatewayRouteTableAssociationDeletedTimeout= 10 * time.Minute
)

func WaitTransitGatewayPolicyTableAssociationCreated(ctx context.Context, conn *ec2.EC2, transitGatewayPolicyTableID, transitGatewayAttachmentID string) (*ec2.TransitGatewayPolicyTableAssociation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayAssociationStateAssociating},
Target:[]string{ec2.TransitGatewayAssociationStateAssociated},
Timeout: TransitGatewayPolicyTableAssociationCreatedTimeout,
Refresh: StatusTransitGatewayPolicyTableAssociationState(ctx, conn, transitGatewayPolicyTableID, transitGatewayAttachmentID),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

funcrn output, err
	}

	return nil, err
}

func WaitTransitGatewayPolicyTableAssociationDeleted(ctx context.Context, conn *ec2.EC2, transitGatewayPolicyTableID, transitGatewayAttachmentID string) (*ec2.TransitGatewayPolicyTableAssociation, error) {
	stateConf := &retry.StateChangeConf{
Pending:g{ec2.TransitGatewayAssociationStateAssociated, ec2.TransitGatewayAssociationStateDisassociating},
Target:[]string{},
Timeout:GatewayPolicyTableAssociationDeletedTimeout,
Refresh:ransitGatewayPolicyTableAssociationState(ctx, conn, transitGatewayPolicyTableID, transitGatewayAttachmentID),
NotFoundChecks: 1,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

funcrn output, err
	}

	return nil, err
}

func WaitTransitGatewayRouteTableAssociationCreated(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID, transitGatewayAttachmentID string) (*ec2.TransitGatewayRouteTableAssociation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayAssociationStateAssociating},
Target:[]string{ec2.TransitGatewayAssociationStateAssociated},
Timeout: TransitGatewayRouteTableAssociationCreatedTimeout,
Refresh: StatusTransitGatewayRouteTableAssociationState(ctx, conn, transitGatewayRouteTableID, transitGatewayAttachmentID),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayRouteTableAssociation); ok {
func

	return nil, err
}

func WaitTransitGatewayRouteTableAssociationDeleted(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID, transitGatewayAttachmentID string) (*ec2.TransitGatewayRouteTableAssociation, error) {
	stateConf := &retry.StateChangeConf{
Pending:g{ec2.TransitGatewayAssociationStateAssociated, ec2.TransitGatewayAssociationStateDisassociating},
Target:[]string{},
Timeout:GatewayRouteTableAssociationDeletedTimeout,
Refresh:ransitGatewayRouteTableAssociationState(ctx, conn, transitGatewayRouteTableID, transitGatewayAttachmentID),
NotFoundChecks: 1,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayRouteTableAssociation); ok {
return output, err
	}

	return nil, err
}

const (
funcnsitGatewayRouteTablePropagationDeletedTimeout = 5 * time.Minute
)

func WaitTransitGatewayRouteTablePropagationCreated(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID string, transitGatewayAttachmentID string) (*ec2.TransitGatewayRouteTablePropagation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayPropagationStateEnabling},
Target:[]string{ec2.TransitGatewayPropagationStateEnabled},
Timeout: TransitGatewayRouteTablePropagationCreatedTimeout,
Refresh: StatusTransitGatewayRouteTablePropagationState(ctx, conn, transitGatewayRouteTableID, transitGatewayAttachmentID),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayRouteTablePropagation); ok {
return output, err
	}

func

func WaitTransitGatewayRouteTablePropagationDeleted(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID string, transitGatewayAttachmentID string) (*ec2.TransitGatewayRouteTablePropagation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayPropagationStateDisabling},
Target:[]string{},
Timeout: TransitGatewayRouteTablePropagationDeletedTimeout,
Refresh: StatusTransitGatewayRouteTablePropagationState(ctx, conn, transitGatewayRouteTableID, transitGatewayAttachmentID),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if tfawserr.ErrCodeEquals(err, errCodeInvalidRouteTableIDNotFound) {
return nil, nil
	}

	if output, ok := outputRaw.(*ec2.TransitGatewayRouteTablePropagation); ok {
return output, err
func
	return nil, err
}

const (
	TransitGatewayVPCAttachmentCreatedTimeout = 10 * time.Minute
	TransitGatewayVPCAttachmentDeletedTimeout = 10 * time.Minute
	TransitGatewayVPCAttachmentUpdatedTimeout = 10 * time.Minute
)

func WaitTransitGatewayVPCAttachmentAccepted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayVpcAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayAttachmentStatePending, ec2.TransitGatewayAttachmentStatePendingAcceptance},
Target:[]string{ec2.TransitGatewayAttachmentStateAvailable},
Timeout: TransitGatewayVPCAttachmentUpdatedTimeout,
Refresh: StatusTransitGatewayVPCAttachmentState(ctx, conn, id),
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayVpcAttachment); ok {
return output, err
	}

	return nil, err
}

func WaitTransitGatewayVPCAttachmentCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayVpcAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayAttachmentStateFailing, ec2.TransitGatewayAttachmentStatePending},
Target:[]string{ec2.TransitGatewayAttachmentStateAvailable, ec2.TransitGatewayAttachmentStatePendingAcceptance},
Timeout: TransitGatewayVPCAttachmentCreatedTimeout,
Refresh: StatusTransitGatewayVPCAttachmentState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayVpcAttachment); ok {
return output, err
	}

func

func WaitTransitGatewayVPCAttachmentDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayVpcAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{
	ec2.TransitGatewayAttachmentStateAvailable,
	ec2.TransitGatewayAttachmentStateDeleting,
	ec2.TransitGatewayAttachmentStatePendingAcceptance,
	ec2.TransitGatewayAttachmentStateRejecting,
},
Target:[]string{ec2.TransitGatewayAttachmentStateDeleted},
Timeout: TransitGatewayVPCAttachmentDeletedTimeout,
Refresh: StatusTransitGatewayVPCAttachmentState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

funcrn output, err
	}

	return nil, err
}

func WaitTransitGatewayVPCAttachmentUpdated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.TransitGatewayVpcAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.TransitGatewayAttachmentStateModifying},
Target:[]string{ec2.TransitGatewayAttachmentStateAvailable},
Timeout: TransitGatewayVPCAttachmentUpdatedTimeout,
Refresh: StatusTransitGatewayVPCAttachmentState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.TransitGatewayVpcAttachment); ok {
return output, err
	}

	return nil, err
}

func WaitVolumeCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Volume, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.VolumeStateCreating},
Target:c2.VolumeStateAvailable},
funcout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Volume); ok {
return output, err
	}

	return nil, err
}

func WaitVolumeDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Volume, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.VolumeStateDeleting},
funcesh:atusVolumeState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Volume); ok {
return output, err
	}

	return nil, err
}

func WaitVolumeUpdated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Volume, error) {
	stateConf := &retry.StateChangeConf{
funcet:c2.VolumeStateAvailable, ec2.VolumeStateInUse},
Refresh:atusVolumeState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Volume); ok {
return output, err
	}

	return nil, err
}

func WaitVolumeAttachmentCreated(ctx context.Context, conn *ec2.EC2, volumeID, instanceID, deviceName string, timeout time.Duration) (*ec2.VolumeAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.VolumeAttachmentStateAttaching},
Target:c2.VolumeAttachmentStateAttached},
Refresh:atusVolumeAttachmentState(ctx, conn, volumeID, instanceID, deviceName),
Timeout:meout,
funcimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VolumeAttachment); ok {
return output, err
	}

	return nil, err
}

func WaitVolumeAttachmentDeleted(ctx context.Context, conn *ec2.EC2, volumeID, instanceID, deviceName string, timeout time.Duration) (*ec2.VolumeAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.VolumeAttachmentStateDetaching},
Target:,
Refresh:atusVolumeAttachmentState(ctx, conn, volumeID, instanceID, deviceName),
funcy:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VolumeAttachment); ok {
return output, err
	}

	return nil, err
}

func WaitVolumeModificationComplete(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.VolumeModification, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.VolumeModificationStateModifying},
// The volume is useable once the state is "optimizing", but will not be at full performance.
// Optimization can take hours. e.g. a full 1 TiB drive takes approximately 6 hours to optimize,
// according to https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/monitoring-volume-modifications.html.
funcesh:atusVolumeModificationState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 30 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VolumeModification); ok {
tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))

return output, err
	}

	return nil, err
}

const (
	vpcCreatedTimeout = 10 * time.Minute
func

func WaitVPCCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.Vpc, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.VpcStatePending},
Target:[]string{ec2.VpcStateAvailable},
Refresh: StatusVPCState(ctx, conn, id),
Timeout: vpcCreatedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Vpc); ok {
return output, err
	}

	return nil, err
}

functeConf := &retry.StateChangeConf{
Target:trconv.FormatBool(expectedValue)},
Refresh:atusVPCAttributeValue(ctx, conn, vpcID, attribute),
Timeout:2PropagationTimeout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Vpc); ok {
return output, err
	}

	return nil, err
}

func WaitVPCCIDRBlockAssociationCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.VpcCidrBlockState, error) {
	stateConf := &retry.StateChangeConf{
funcet:c2.VpcCidrBlockStateCodeAssociated},
Refresh:atusVPCCIDRBlockAssociationState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcCidrBlockState); ok {
if state := aws.StringValue(output.State); state == ec2.VpcCidrBlockStateCodeFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))
}

return output, err
	}

	return nil, err
}
func WaitVPCCIDRBlockAssociationDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.VpcCidrBlockState, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.VpcCidrBlockStateCodeAssociated, ec2.VpcCidrBlockStateCodeDisassociating, ec2.VpcCidrBlockStateCodeFailing},
Target:,
Refresh:atusVPCCIDRBlockAssociationState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcCidrBlockState); ok {
if state := aws.StringValue(output.State); state == ec2.VpcCidrBlockStateCodeFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))
}

return output, err
	}

	return nil, err
}

const (
	vpcIPv6CIDRBlockAssociationCreatedTimeout = 10 * time.Minute
	vpcIPv6CIDRBlockAssociationDeletedTimeout = 5 * time.Minute
)

func WaitVPCIPv6CIDRBlockAssociationCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.VpcCidrBlockState, error) {
funcing:string{ec2.VpcCidrBlockStateCodeAssociating, ec2.VpcCidrBlockStateCodeDisassociated, ec2.VpcCidrBlockStateCodeFailing},
Target:c2.VpcCidrBlockStateCodeAssociated},
Refresh:atusVPCIPv6CIDRBlockAssociationState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcCidrBlockState); ok {
if state := aws.StringValue(output.State); state == ec2.VpcCidrBlockStateCodeFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))
}

return output, err
	}
funcurn nil, err
}

func WaitVPCIPv6CIDRBlockAssociationDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.VpcCidrBlockState, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.VpcCidrBlockStateCodeAssociated, ec2.VpcCidrBlockStateCodeDisassociating, ec2.VpcCidrBlockStateCodeFailing},
Target:,
Refresh:atusVPCIPv6CIDRBlockAssociationState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcCidrBlockState); ok {
if state := aws.StringValue(output.State); state == ec2.VpcCidrBlockStateCodeFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))
func
return output, err
	}

	return nil, err
}

func WaitVPCPeeringConnectionActive(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.VpcPeeringConnection, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.VpcPeeringConnectionStateReasonCodeInitiatingRequest, ec2.VpcPeeringConnectionStateReasonCodeProvisioning},
Target:[]string{ec2.VpcPeeringConnectionStateReasonCodeActive, ec2.VpcPeeringConnectionStateReasonCodePendingAcceptance},
Refresh: StatusVPCPeeringConnectionActive(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcPeeringConnection); ok {
tfresource.SetLastError(err, errors.New(aws.StringValue(output.Status.Message)))

return output, err
	}

func

func WaitVPCPeeringConnectionDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.VpcPeeringConnection, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{
	ec2.VpcPeeringConnectionStateReasonCodeActive,
	ec2.VpcPeeringConnectionStateReasonCodeDeleting,
	ec2.VpcPeeringConnectionStateReasonCodePendingAcceptance,
},
Target:[]string{},
Refresh: StatusVPCPeeringConnectionDeleted(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcPeeringConnection); ok {
tfresource.SetLastError(err, errors.New(aws.StringValue(output.Status.Message)))

return output, err
	}

	return nil, err
}

const (
	VPNGatewayDeletedTimeout = 5 * time.Minute

funcGatewayVPCAttachmentDetachedTimeout = 30 * time.Minute
)

func WaitVPNGatewayVPCAttachmentAttached(ctx context.Context, conn *ec2.EC2, vpnGatewayID, vpcID string) (*ec2.VpcAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.AttachmentStatusAttaching},
Target:[]string{ec2.AttachmentStatusAttached},
Refresh: StatusVPNGatewayVPCAttachmentState(ctx, conn, vpnGatewayID, vpcID),
Timeout: VPNGatewayVPCAttachmentAttachedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcAttachment); ok {
return output, err
	}

	return nil, err
}

func WaitVPNGatewayVPCAttachmentDetached(ctx context.Context, conn *ec2.EC2, vpnGatewayID, vpcID string) (*ec2.VpcAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.AttachmentStatusAttached, ec2.AttachmentStatusDetaching},
funcesh: StatusVPNGatewayVPCAttachmentState(ctx, conn, vpnGatewayID, vpcID),
Timeout: VPNGatewayVPCAttachmentDetachedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcAttachment); ok {
return output, err
	}

	return nil, err
}

const (
	customerGatewayCreatedTimeout = 10 * time.Minute
	customerGatewayDeletedTimeout = 5 * time.Minute
)

func WaitCustomerGatewayCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.CustomerGateway, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{CustomerGatewayStatePending},
Target:ustomerGatewayStateAvailable},
Refresh:atusCustomerGatewayState(ctx, conn, id),
funcy:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.CustomerGateway); ok {
return output, err
	}

	return nil, err
}

func WaitCustomerGatewayDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.CustomerGateway, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{CustomerGatewayStateAvailable, CustomerGatewayStateDeleting},
Target:[]string{},
Refresh: StatusCustomerGatewayState(ctx, conn, id),
Timeout: customerGatewayDeletedTimeout,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.CustomerGateway); ok {
return output, err
	}

	return nil, err
}

func WaitNATGatewayCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.NatGateway, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.NatGatewayStatePending},
Target:[]string{ec2.NatGatewayStateAvailable},
Refresh: StatusNATGatewayState(ctx, conn, id),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NatGateway); ok {
if state := aws.StringValue(output.State); state == ec2.NatGatewayStateFailed {
	tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(output.FailureCode), aws.StringValue(output.FailureMessage)))
}

return output, err
	}

	return nil, err
}
func WaitNATGatewayDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.NatGateway, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.NatGatewayStateDeleting},
Target:,
Refresh:atusNATGatewayState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 10 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NatGateway); ok {
if state := aws.StringValue(output.State); state == ec2.NatGatewayStateFailed {
	tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(output.FailureCode), aws.StringValue(output.FailureMessage)))
}

func

	return nil, err
}

func WaitNATGatewayAddressAssigned(ctx context.Context, conn *ec2.EC2, natGatewayID, privateIP string, timeout time.Duration) (*ec2.NatGatewayAddress, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.NatGatewayAddressStatusAssigning},
Target:[]string{ec2.NatGatewayAddressStatusSucceeded},
Refresh: StatusNATGatewayAddressByNATGatewayIDAndPrivateIP(ctx, conn, natGatewayID, privateIP),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NatGatewayAddress); ok {
if status := aws.StringValue(output.Status); status == ec2.NatGatewayAddressStatusFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.FailureMessage)))
}

return output, err
	}
funcurn nil, err
}

func WaitNATGatewayAddressAssociated(ctx context.Context, conn *ec2.EC2, natGatewayID, allocationID string, timeout time.Duration) (*ec2.NatGatewayAddress, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.NatGatewayAddressStatusAssociating},
Target:[]string{ec2.NatGatewayAddressStatusSucceeded},
Refresh: StatusNATGatewayAddressByNATGatewayIDAndAllocationID(ctx, conn, natGatewayID, allocationID),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NatGatewayAddress); ok {
if status := aws.StringValue(output.Status); status == ec2.NatGatewayAddressStatusFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.FailureMessage)))
}

return output, err
func
	return nil, err
}

func WaitNATGatewayAddressDisassociated(ctx context.Context, conn *ec2.EC2, natGatewayID, allocationID string, timeout time.Duration) (*ec2.NatGatewayAddress, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.NatGatewayAddressStatusSucceeded, ec2.NatGatewayAddressStatusDisassociating},
Target:[]string{},
Refresh: StatusNATGatewayAddressByNATGatewayIDAndAllocationID(ctx, conn, natGatewayID, allocationID),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NatGatewayAddress); ok {
if status := aws.StringValue(output.Status); status == ec2.NatGatewayAddressStatusFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.FailureMessage)))
func
return output, err
	}

	return nil, err
}

func WaitNATGatewayAddressUnassigned(ctx context.Context, conn *ec2.EC2, natGatewayID, privateIP string, timeout time.Duration) (*ec2.NatGatewayAddress, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.NatGatewayAddressStatusUnassigning},
Target:[]string{},
Refresh: StatusNATGatewayAddressByNATGatewayIDAndPrivateIP(ctx, conn, natGatewayID, privateIP),
Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NatGatewayAddress); ok {
if status := aws.StringValue(output.Status); status == ec2.NatGatewayAddressStatusFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.FailureMessage)))
}
funcrn output, err
	}

	return nil, err
}

const (
	vpnConnectionCreatedTimeout = 40 * time.Minute
	vpnConnectionDeletedTimeout = 30 * time.Minute
	vpnConnectionUpdatedTimeout = 30 * time.Minute
)

func WaitVPNConnectionCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.VpnConnection, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.VpnStatePending},
Target:c2.VpnStateAvailable},
Refresh:atusVPNConnectionState(ctx, conn, id),
Timeout:nConnectionCreatedTimeout,
Delay:.Second,
MinTimeout: 10 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
funcoutput, ok := outputRaw.(*ec2.VpnConnection); ok {
return output, err
	}

	return nil, err
}

func WaitVPNConnectionDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.VpnConnection, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.VpnStateDeleting},
Target:,
Refresh:atusVPNConnectionState(ctx, conn, id),
Timeout:nConnectionDeletedTimeout,
Delay:.Second,
MinTimeout: 10 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpnConnection); ok {
return output, err
func
	return nil, err
}

func WaitVPNConnectionUpdated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.VpnConnection, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{vpnStateModifying},
Target:c2.VpnStateAvailable},
Refresh:atusVPNConnectionState(ctx, conn, id),
Timeout:nConnectionUpdatedTimeout,
Delay:.Second,
MinTimeout: 10 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpnConnection); ok {
return output, err
	}

	return nil, err
func
const (
	vpnConnectionRouteCreatedTimeout = 15 * time.Second
	vpnConnectionRouteDeletedTimeout = 15 * time.Second
)

func WaitVPNConnectionRouteCreated(ctx context.Context, conn *ec2.EC2, vpnConnectionID, cidrBlock string) (*ec2.VpnStaticRoute, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.VpnStatePending},
Target:[]string{ec2.VpnStateAvailable},
Refresh: StatusVPNConnectionRouteState(ctx, conn, vpnConnectionID, cidrBlock),
Timeout: vpnConnectionRouteCreatedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpnStaticRoute); ok {
return output, err
	}

	return nil, err
func
func WaitVPNConnectionRouteDeleted(ctx context.Context, conn *ec2.EC2, vpnConnectionID, cidrBlock string) (*ec2.VpnStaticRoute, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.VpnStatePending, ec2.VpnStateAvailable, ec2.VpnStateDeleting},
Target:[]string{},
Refresh: StatusVPNConnectionRouteState(ctx, conn, vpnConnectionID, cidrBlock),
Timeout: vpnConnectionRouteDeletedTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpnStaticRoute); ok {
return output, err
	}

	return nil, err
}

const (
	HostCreatedTimeout = 10 * time.Minute
	HostUpdatedTimeout = 10 * time.Minute
	HostDeletedTimeout = 20 * time.Minute
)

func WaitHostCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.Host, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.AllocationStatePending},
funcout: HostCreatedTimeout,
Refresh: StatusHostState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Host); ok {
return output, err
	}

	return nil, err
}

func WaitHostUpdated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.Host, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.AllocationStatePending},
Target:[]string{ec2.AllocationStateAvailable},
Timeout: HostUpdatedTimeout,
Refresh: StatusHostState(ctx, conn, id),
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Host); ok {
return output, err
	}

	return nil, err
}

func WaitHostDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.Host, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.AllocationStateAvailable},
Target:[]string{},
Timeout: HostDeletedTimeout,
Refresh: StatusHostState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
funcoutput, ok := outputRaw.(*ec2.Host); ok {
return output, err
	}

	return nil, err
}

const (
	dhcpOptionSetDeletedTimeout = 3 * time.Minute
)

func WaitInternetGatewayAttached(ctx context.Context, conn *ec2.EC2, internetGatewayID, vpcID string, timeout time.Duration) (*ec2.InternetGatewayAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending:g{ec2.AttachmentStatusAttaching},
Target:[]string{InternetGatewayAttachmentStateAvailable},
Timeout:,
NotFoundChecks: InternetGatewayNotFoundChecks,
Refresh:nternetGatewayAttachmentState(ctx, conn, internetGatewayID, vpcID),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.InternetGatewayAttachment); ok {
return output, err
func
	return nil, err
}

func WaitInternetGatewayDetached(ctx context.Context, conn *ec2.EC2, internetGatewayID, vpcID string, timeout time.Duration) (*ec2.InternetGatewayAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{InternetGatewayAttachmentStateAvailable, ec2.AttachmentStatusDetaching},
Target:[]string{},
Timeout: timeout,
Refresh: StatusInternetGatewayAttachmentState(ctx, conn, internetGatewayID, vpcID),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.InternetGatewayAttachment); ok {
return output, err
	}
funcurn nil, err
}

const (
	ManagedPrefixListTimeout = 15 * time.Minute
)

func WaitManagedPrefixListCreated(ctx context.Context, conn *ec2.EC2, id string) (*ec2.ManagedPrefixList, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.PrefixListStateCreateInProgress},
Target:[]string{ec2.PrefixListStateCreateComplete},
Timeout: ManagedPrefixListTimeout,
Refresh: StatusManagedPrefixListState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.ManagedPrefixList); ok {
if state := aws.StringValue(output.State); state == ec2.PrefixListStateCreateFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StateMessage)))
}

return output, err
func
	return nil, err
}

func WaitManagedPrefixListModified(ctx context.Context, conn *ec2.EC2, id string) (*ec2.ManagedPrefixList, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.PrefixListStateModifyInProgress},
Target:[]string{ec2.PrefixListStateModifyComplete},
Timeout: ManagedPrefixListTimeout,
Refresh: StatusManagedPrefixListState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.ManagedPrefixList); ok {
if state := aws.StringValue(output.State); state == ec2.PrefixListStateModifyFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StateMessage)))
func
return output, err
	}

	return nil, err
}

func WaitManagedPrefixListDeleted(ctx context.Context, conn *ec2.EC2, id string) (*ec2.ManagedPrefixList, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.PrefixListStateDeleteInProgress},
Target:[]string{},
Timeout: ManagedPrefixListTimeout,
Refresh: StatusManagedPrefixListState(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

functate := aws.StringValue(output.State); state == ec2.PrefixListStateDeleteFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StateMessage)))
}

return output, err
	}

	return nil, err
}

func WaitNetworkInsightsAnalysisCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.NetworkInsightsAnalysis, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.AnalysisStatusRunning},
Target:c2.AnalysisStatusSucceeded},
Timeout:meout,
Refresh:atusNetworkInsightsAnalysis(ctx, conn, id),
Delay:.Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
funcoutput, ok := outputRaw.(*ec2.NetworkInsightsAnalysis); ok {
tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))

return output, err
	}

	return nil, err
}

const (
	networkInterfaceAttachedTimeout = 5 * time.Minute
	NetworkInterfaceDetachedTimeout = 10 * time.Minute
)

func WaitNetworkInterfaceAttached(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.NetworkInterfaceAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.AttachmentStatusAttaching},
Target:[]string{ec2.AttachmentStatusAttached},
funcesh: StatusNetworkInterfaceAttachmentStatus(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NetworkInterfaceAttachment); ok {
return output, err
	}

	return nil, err
}

func WaitNetworkInterfaceAvailableAfterUse(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.NetworkInterface, error) {
	// Hyperplane attached ENI.
	// Wait for it to be moved into a removable state.
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.NetworkInterfaceStatusInUse},
Target:c2.NetworkInterfaceStatusAvailable},
Timeout:meout,
Refresh:atusNetworkInterfaceStatus(ctx, conn, id),
Delay:.Second,
funcandle EC2 ENI eventual consistency. It can take up to 3 minutes.
ContinuousTargetOccurence: 18,
NotFoundChecks:
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NetworkInterface); ok {
return output, err
	}

	return nil, err
}

func WaitNetworkInterfaceCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.NetworkInterface, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{NetworkInterfaceStatusPending},
Target:[]string{ec2.NetworkInterfaceStatusAvailable},
Timeout: timeout,
Refresh: StatusNetworkInterfaceStatus(ctx, conn, id),
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NetworkInterface); ok {
return output, err
	}

	return nil, err
}

func WaitNetworkInterfaceDetached(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.NetworkInterfaceAttachment, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.AttachmentStatusAttached, ec2.AttachmentStatusDetaching},
Target:[]string{ec2.AttachmentStatusDetached},
Timeout: timeout,
Refresh: StatusNetworkInterfaceAttachmentStatus(ctx, conn, id),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.NetworkInterfaceAttachment); ok {
return output, err
func
	return nil, err
}

const (
	PlacementGroupCreatedTimeout = 5 * time.Minute
	PlacementGroupDeletedTimeout = 5 * time.Minute
)

func WaitPlacementGroupCreated(ctx context.Context, conn *ec2.EC2, name string) (*ec2.PlacementGroup, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.PlacementGroupStatePending},
Target:[]string{ec2.PlacementGroupStateAvailable},
Timeout: PlacementGroupCreatedTimeout,
Refresh: StatusPlacementGroupState(ctx, conn, name),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.PlacementGroup); ok {
return output, err
func
	return nil, err
}

func WaitPlacementGroupDeleted(ctx context.Context, conn *ec2.EC2, name string) (*ec2.PlacementGroup, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.PlacementGroupStateDeleting},
Target:[]string{},
Timeout: PlacementGroupDeletedTimeout,
Refresh: StatusPlacementGroupState(ctx, conn, name),
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.PlacementGroup); ok {
return output, err
	}

	return nil, err
}

func WaitSpotFleetRequestCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.SpotFleetRequestConfig, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.BatchStateSubmitted},
Target:c2.BatchStateActive},
Refresh:atusSpotFleetRequestState(ctx, conn, id),
funcimeout: 10 * time.Second,
Delay:.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.SpotFleetRequestConfig); ok {
return output, err
	}

	return nil, err
}

func WaitSpotFleetRequestFulfilled(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.SpotFleetRequestConfig, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.ActivityStatusPendingFulfillment},
Target:c2.ActivityStatusFulfilled},
funcout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.SpotFleetRequestConfig); ok {
if activityStatus := aws.StringValue(output.ActivityStatus); activityStatus == ec2.ActivityStatusError {
	var errs *multierror.Error

	input := &ec2.DescribeSpotFleetRequestHistoryInput{
SpotFleetRequestId: aws.String(id),
StartTime: aws.Time(time.UnixMilli(0)),
	}

	if output, err := FindSpotFleetRequestHistoryRecords(ctx, conn, input); err == nil {
for _, v := range output {
	if eventType := aws.StringValue(v.EventType); eventType == ec2.EventTypeError || eventType == ec2.EventTypeInformation {
errs = multierror.Append(errs, errors.New(v.String()))
	}
}
	}

func

return output, err
	}

	return nil, err
}

func WaitSpotFleetRequestUpdated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.SpotFleetRequestConfig, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.BatchStateModifying},
Target:c2.BatchStateActive},
Refresh:atusSpotFleetRequestState(ctx, conn, id),
Timeout:meout,
MinTimeout: 10 * time.Second,
Delay:.Second,
	}
funcputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.SpotFleetRequestConfig); ok {
return output, err
	}

	return nil, err
}

func WaitSpotInstanceRequestFulfilled(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.SpotInstanceRequest, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{spotInstanceRequestStatusCodePendingEvaluation, spotInstanceRequestStatusCodePendingFulfillment},
Target:potInstanceRequestStatusCodeFulfilled},
Refresh:atusSpotInstanceRequest(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.SpotInstanceRequest); ok {
funcFault := fmt.Errorf("%s: %s", aws.StringValue(fault.Code), aws.StringValue(fault.Message))
	tfresource.SetLastError(err, fmt.Errorf("%s %w", aws.StringValue(output.Status.Message), errFault))
} else {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.Status.Message)))
}

return output, err
	}

	return nil, err
}

func WaitVPCEndpointAccepted(ctx context.Context, conn *ec2.EC2, vpcEndpointID string, timeout time.Duration) (*ec2.VpcEndpoint, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{vpcEndpointStatePendingAcceptance},
Target:pcEndpointStateAvailable},
Timeout:meout,
funcy:Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcEndpoint); ok {
if state, lastError := aws.StringValue(output.State), output.LastError; state == vpcEndpointStateFailed && lastError != nil {
	tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(lastError.Code), aws.StringValue(lastError.Message)))
}

return output, err
	}

	return nil, err
}

functeConf := &retry.StateChangeConf{
Pending:string{vpcEndpointStatePending},
Target:pcEndpointStateAvailable, vpcEndpointStatePendingAcceptance},
Timeout:meout,
Refresh:atusVPCEndpointState(ctx, conn, vpcEndpointID),
Delay:Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcEndpoint); ok {
if state, lastError := aws.StringValue(output.State), output.LastError; state == vpcEndpointStateFailed && lastError != nil {
	tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(lastError.Code), aws.StringValue(lastError.Message)))
}

return output, err
	}

func

func WaitVPCEndpointDeleted(ctx context.Context, conn *ec2.EC2, vpcEndpointID string, timeout time.Duration) (*ec2.VpcEndpoint, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{vpcEndpointStateDeleting},
Target:,
Refresh:atusVPCEndpointState(ctx, conn, vpcEndpointID),
Timeout:meout,
Delay:Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcEndpoint); ok {
return output, err
	}

	return nil, err
}

func WaitVPCEndpointServiceAvailable(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.ServiceConfiguration, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.ServiceStatePending},
Target:c2.ServiceStateAvailable},
Refresh:atusVPCEndpointServiceStateAvailable(ctx, conn, id),
Timeout:meout,
Delay:Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.ServiceConfiguration); ok {
return output, err
	}

	return nil, err
func
func WaitVPCEndpointServiceDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.ServiceConfiguration, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.ServiceStateAvailable, ec2.ServiceStateDeleting},
Target:,
Timeout:meout,
Refresh:atusVPCEndpointServiceStateDeleted(ctx, conn, id),
Delay:Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.ServiceConfiguration); ok {
return output, err
	}

	return nil, err
}
func WaitVPCEndpointRouteTableAssociationDeleted(ctx context.Context, conn *ec2.EC2, vpcEndpointID, routeTableID string) error {
	stateConf := &retry.StateChangeConf{
Pending:string{VPCEndpointRouteTableAssociationStatusReady},
Target:,
Refresh:atusVPCEndpointRouteTableAssociation(ctx, conn, vpcEndpointID, routeTableID),
Timeout:2PropagationTimeout,
ContinuousTargetOccurence: 2,
	}

	_, err := stateConf.WaitForStateContext(ctx)

	return err
}

func WaitVPCEndpointRouteTableAssociationReady(ctx context.Context, conn *ec2.EC2, vpcEndpointID, routeTableID string) error {
	stateConf := &retry.StateChangeConf{
Pending:string{},
Target:PCEndpointRouteTableAssociationStatusReady},
Refresh:atusVPCEndpointRouteTableAssociation(ctx, conn, vpcEndpointID, routeTableID),
Timeout:2PropagationTimeout,
ContinuousTargetOccurence: 2,
	}

	_, err := stateConf.WaitForStateContext(ctx)

	return err
func
func WaitEBSSnapshotImportComplete(ctx context.Context, conn *ec2.EC2, importTaskID string, timeout time.Duration) (*ec2.SnapshotTaskDetail, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{
	EBSSnapshotImportStateActive,
	EBSSnapshotImportStateUpdating,
	EBSSnapshotImportStateValidating,
	EBSSnapshotImportStateValidated,
	EBSSnapshotImportStateConverting,
},
Target:[]string{EBSSnapshotImportStateCompleted},
Refresh: StatusEBSSnapshotImport(ctx, conn, importTaskID),
Timeout: timeout,
Delay:* time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.SnapshotTaskDetail); ok {
tfresource.SetLastError(err, errors.New(aws.StringValue(output.StatusMessage)))

return output, err
	}
funcurn nil, err
}

func waitVPCEndpointConnectionAccepted(ctx context.Context, conn *ec2.EC2, serviceID, vpcEndpointID string, timeout time.Duration) (*ec2.VpcEndpointConnection, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{vpcEndpointStatePendingAcceptance, vpcEndpointStatePending},
Target:pcEndpointStateAvailable},
Refresh:atusVPCEndpointConnectionVPCEndpointState(ctx, conn, serviceID, vpcEndpointID),
Timeout:meout,
Delay:Second,
MinTimeout: 5 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.VpcEndpointConnection); ok {
return output, err
	}

	return nil, err
}

const (
func

func waitEBSSnapshotTierArchive(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.SnapshotTierStatus, error) { //nolint:unparam
	stateConf := &retry.StateChangeConf{
Pending: []string{TargetStorageTierStandard},
Target:[]string{ec2.TargetStorageTierArchive},
Refresh: StatusSnapshotStorageTier(ctx, conn, id),
Timeout: timeout,
Delay:* time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.SnapshotTierStatus); ok {
tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(output.LastTieringOperationStatus), aws.StringValue(output.LastTieringOperationStatusDetail)))

return output, err
	}

func

func WaitIPAMCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Ipam, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamStateCreateInProgress},
Target:[]string{ec2.IpamStateCreateComplete},
Refresh: StatusIPAMState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Ipam); ok {
return output, err
	}

	return nil, err
}
func WaitIPAMDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Ipam, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamStateCreateComplete, ec2.IpamStateModifyComplete, ec2.IpamStateDeleteInProgress},
Target:[]string{},
Refresh: StatusIPAMState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Ipam); ok {
return output, err
	}

	return nil, err
}

func WaitIPAMUpdated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Ipam, error) {
funcing: []string{ec2.IpamStateModifyInProgress},
Target:[]string{ec2.IpamStateModifyComplete},
Refresh: StatusIPAMState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Ipam); ok {
return output, err
	}

	return nil, err
func
func WaitIPAMPoolCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.IpamPool, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamPoolStateCreateInProgress},
Target:[]string{ec2.IpamPoolStateCreateComplete},
Refresh: StatusIPAMPoolState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamPool); ok {
if state := aws.StringValue(output.State); state == ec2.IpamPoolStateCreateFailed {
func

return output, err
	}

	return nil, err
}

func WaitIPAMPoolDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.IpamPool, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamPoolStateDeleteInProgress},
Target:[]string{},
Refresh: StatusIPAMPoolState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamPool); ok {
if state := aws.StringValue(output.State); state == ec2.IpamPoolStateDeleteFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StateMessage)))
}

return output, err
	}
funcurn nil, err
}

func WaitIPAMPoolUpdated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.IpamPool, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamPoolStateModifyInProgress},
Target:[]string{ec2.IpamPoolStateModifyComplete},
Refresh: StatusIPAMPoolState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamPool); ok {
if state := aws.StringValue(output.State); state == ec2.IpamPoolStateModifyFailed {
	tfresource.SetLastError(err, errors.New(aws.StringValue(output.StateMessage)))
}

return output, err
	}

	return nil, err
func
func WaitIPAMPoolCIDRIdCreated(ctx context.Context, conn *ec2.EC2, poolCidrId, poolID, cidrBlock string, timeout time.Duration) (*ec2.IpamPoolCidr, error) {
	stateConf := &retry.StateChangeConf{
Pending:g{ec2.IpamPoolCidrStatePendingProvision},
Target:[]string{ec2.IpamPoolCidrStateProvisioned},
Refresh:PAMPoolCIDRState(ctx, conn, cidrBlock, poolID, poolCidrId),
Timeout:,
Delay: 5 * time.Second,
NotFoundChecks: IPAMPoolCIDRNotFoundChecks,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamPoolCidr); ok {
if state, failureReason := aws.StringValue(output.State), output.FailureReason; state == ec2.IpamPoolCidrStateFailedProvision && failureReason != nil {
	tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(failureReason.Code), aws.StringValue(failureReason.Message)))
}

return output, err
	}
funcurn nil, err
}

func WaitIPAMPoolCIDRDeleted(ctx context.Context, conn *ec2.EC2, cidrBlock, poolID, poolCidrId string, timeout time.Duration) (*ec2.IpamPoolCidr, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamPoolCidrStatePendingDeprovision, ec2.IpamPoolCidrStateProvisioned},
Target:[]string{},
Refresh: StatusIPAMPoolCIDRState(ctx, conn, cidrBlock, poolID, poolCidrId),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamPoolCidr); ok {
if state, failureReason := aws.StringValue(output.State), output.FailureReason; state == ec2.IpamPoolCidrStateFailedDeprovision && failureReason != nil {
	tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(failureReason.Code), aws.StringValue(failureReason.Message)))
}
funcrn output, err
	}

	return nil, err
}

func WaitIPAMPoolCIDRAllocationCreated(ctx context.Context, conn *ec2.EC2, allocationID, poolID string, timeout time.Duration) (*ec2.IpamPoolAllocation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{},
Target:[]string{IpamPoolCIDRAllocationCreateComplete},
Refresh: StatusIPAMPoolCIDRAllocationState(ctx, conn, allocationID, poolID),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamPoolAllocation); ok {
func

	return nil, err
}

func WaitIPAMResourceDiscoveryAvailable(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Ipam, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamResourceDiscoveryStateCreateInProgress},
Target:[]string{ec2.IpamResourceDiscoveryStateCreateComplete},
Refresh: StatusIPAMResourceDiscoveryState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Ipam); ok {
return output, err
func
	return nil, err
}

func WaitIPAMResourceDiscoveryAssociationAvailable(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.IpamResourceDiscoveryAssociation, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamResourceDiscoveryAssociationStateAssociateInProgress},
Target:[]string{ec2.IpamResourceDiscoveryAssociationStateAssociateComplete},
Refresh: StatusIPAMResourceDiscoveryAssociationStatus(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamResourceDiscoveryAssociation); ok {
return output, err
	}

	return nil, err
}

functeConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamResourceDiscoveryAssociationStateAssociateComplete, ec2.IpamResourceDiscoveryAssociationStateDisassociateInProgress},
Target:[]string{},
Refresh: StatusIPAMResourceDiscoveryAssociationStatus(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamResourceDiscoveryAssociation); ok {
return output, err
	}

	return nil, err
}

func WaitIPAMResourceDiscoveryDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.IpamResourceDiscovery, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamResourceDiscoveryStateCreateComplete, ec2.IpamResourceDiscoveryStateModifyComplete, ec2.IpamResourceDiscoveryStateDeleteInProgress},
Target:[]string{},
Refresh: StatusIPAMResourceDiscoveryState(ctx, conn, id),
funcy: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamResourceDiscovery); ok {
return output, err
	}

	return nil, err
}

func WaitIPAMResourceDiscoveryUpdated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.IpamResourceDiscovery, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamResourceDiscoveryStateModifyInProgress},
Target:[]string{ec2.IpamResourceDiscoveryStateModifyComplete},
Refresh: StatusIPAMResourceDiscoveryState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
funcoutput, ok := outputRaw.(*ec2.IpamResourceDiscovery); ok {
return output, err
	}

	return nil, err
}

func WaitIPAMScopeCreated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.IpamScope, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamScopeStateCreateInProgress},
Target:[]string{ec2.IpamScopeStateCreateComplete},
Refresh: StatusIPAMScopeState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamScope); ok {
return output, err
	}

	return nil, err
func
func WaitIPAMScopeDeleted(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.IpamScope, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamScopeStateCreateComplete, ec2.IpamScopeStateModifyComplete, ec2.IpamScopeStateDeleteInProgress},
Target:[]string{},
Refresh: StatusIPAMScopeState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamScope); ok {
return output, err
	}

	return nil, err
}

func WaitIPAMScopeUpdated(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.IpamScope, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{ec2.IpamScopeStateModifyInProgress},
funcesh: StatusIPAMScopeState(ctx, conn, id),
Timeout: timeout,
Delay: time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.IpamScope); ok {
return output, err
	}

	return nil, err
}

func WaitInstanceStoppedWithContext(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Instance, error) {
	stateConf := &retry.StateChangeConf{
Pending: []string{
	ec2.InstanceStateNamePending,
func.InstanceStateNameShuttingDown,
	ec2.InstanceStateNameStopping,
},
Target:c2.InstanceStateNameStopped},
Refresh:atusInstanceState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Instance); ok {
if stateReason := output.StateReason; stateReason != nil {
	tfresource.SetLastError(err, errors.New(aws.StringValue(stateReason.Message)))
}

return output, err
func
	return nil, err
}

func WaitInstanceStartedWithContext(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Instance, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.InstanceStateNamePending, ec2.InstanceStateNameStopped},
Target:c2.InstanceStateNameRunning},
Refresh:atusInstanceState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*ec2.Instance); ok {
if stateReason := output.StateReason; stateReason != nil {
func

return output, err
	}

	return nil, err
}

func WaitInstanceReadyWithContext(ctx context.Context, conn *ec2.EC2, id string, timeout time.Duration) (*ec2.Instance, error) {
	stateConf := &retry.StateChangeConf{
Pending:string{ec2.InstanceStateNamePending, ec2.InstanceStateNameStopping},
Target:c2.InstanceStateNameRunning, ec2.InstanceStateNameStopped},
Refresh:atusInstanceState(ctx, conn, id),
Timeout:meout,
Delay:.Second,
MinTimeout: 3 * time.Second,
	}

func
	if output, ok := outputRaw.(*ec2.Instance); ok {
if stateReason := output.StateReason; stateReason != nil {
	tfresource.SetLastError(err, errors.New(aws.StringValue(stateReason.Message)))
}

return output, err
	}

	return nil, err
}

func WaitInstanceConnectEndpointCreated(ctx context.Context, conn *ec2_sdkv2.Client, id string, timeout time.Duration) (*types.Ec2InstanceConnectEndpoint, error) {
	stateConf := &retry.StateChangeConf{
Pending: enum.Slice(types.Ec2InstanceConnectEndpointStateCreateInProgress),
Target:enum.Slice(types.Ec2InstanceConnectEndpointStateCreateComplete),
Refresh: StatusInstanceConnectEndpointState(ctx, conn, id),
Timeout: timeout,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*types.Ec2InstanceConnectEndpoint); ok {
tfresource.SetLastError(err, errors.New(aws_sdkv2.ToString(output.StateMessage)))

return output, err
	}

	return nil, err
}

func WaitInstanceConnectEndpointDeleted(ctx context.Context, conn *ec2_sdkv2.Client, id string, timeout time.Duration) (*types.Ec2InstanceConnectEndpoint, error) {
	stateConf := &retry.StateChangeConf{
Pending: enum.Slice(types.Ec2InstanceConnectEndpointStateDeleteInProgress),
Target:[]string{},
Refresh: StatusInstanceConnectEndpointState(ctx, conn, id),
Timeout: timeout,
func
	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*types.Ec2InstanceConnectEndpoint); ok {
tfresource.SetLastError(err, errors.New(aws_sdkv2.ToString(output.StateMessage)))

return output, err
	}

	return nil, err
}
funcfuncfuncfuncfuncfuncfunc