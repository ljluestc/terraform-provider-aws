// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"strconv"

	ec2_sdkv2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func {
func() (interface{}, string, error) {
		output, err := FindAvailabilityZoneGroupByName(ctx, conn, name)
func tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.OptInStatus), nil
	}
}


func StatusCapacityReservationState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
functput, err := FindCapacityReservationByID(ctx, conn, id)
func tfresource.NotFound(err) {
			return nil, "", nil
func
		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusCarrierGatewayState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindCarrierGatewayByID(ctx, conn, id)

funceturn nil, "", nil
func
		if err != nil {
func

		return output, aws.StringValue(output.State), nil
	}
}

// StatusLocalGatewayRouteTableVPCAssociationState fetches the LocalGatewayRouteTableVpcAssociation and its State

func StatusLocalGatewayRouteTableVPCAssociationState(ctx context.Context, conn *ec2.EC2, localGatewayRouteTableVpcAssociationID string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		input := &ec2.DescribeLocalGatewayRouteTableVpcAssociationsInput{
			LocalGatewayRouteTableVpcAssociationIds: aws.StringSlice([]string{localGatewayRouteTableVpcAssociationID}),
		}

		output, err := conn.DescribeLocalGatewayRouteTableVpcAssociationsWithContext(ctx, input)
func err != nil {
func

func
		for _, outputAssociation := range output.LocalGatewayRouteTableVpcAssociations {
			if outputAssociation == nil {
				continue
			}

			if aws.StringValue(outputAssociation.LocalGatewayRouteTableVpcAssociationId) == localGatewayRouteTableVpcAssociationID {
				association = outputAssociation
				break
			}
		}

		if association == nil {
			return association, ec2.RouteTableAssociationStateCodeDisassociated, nil
		}

		return association, aws.StringValue(association.State), nil
	}
}


func StatusClientVPNEndpointState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindClientVPNEndpointByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
func
func
}
func
func StatusClientVPNEndpointClientConnectResponseOptionsState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindClientVPNEndpointClientConnectResponseOptionsByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status.Code), nil
func
func
func StatusClientVPNAuthorizationRule(ctx context.Context, conn *ec2.EC2, endpointID, targetNetworkCIDR, accessGroupID string) retry.StateRefresh
funcurn 
func() (interface{}, string, error) {
		output, err := FindClientVPNAuthorizationRuleByThreePartKey(ctx, conn, endpointID, targetNetworkCIDR, accessGroupID)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status.Code), nil
	}
}

func StatusClientVPNNetworkAssociation(ctx context.Context, conn *ec2.EC2, associationID, endpointID string) retry.StateRefresh
funcurn 
func() (interface{}, string, error) {
func
		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status.Code), nil
	}
}


func StatusClientVPNRoute(ctx context.Context, conn *ec2.EC2, endpointID, targetSubnetID, destinationCIDR string) retry.StateRefresh
func {
func() (interface{}, string, error) {
func
		if tfresource.NotFound(err) {
func

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status.Code), nil
	}
}


func StatusFleetState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		// Don't call FindFleetByID as it maps useful status codes to NotFoundError.
funcleetIds: aws.StringSlice([]string{id}),
func
		if tfresource.NotFound(err) {
func

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.FleetState), nil
	}
}


func StatusImageState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindImageByID(ctx, conn, id)
func tfresource.NotFound(err) {
func

funceturn nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}

// StatusInstanceIAMInstanceProfile fetches the Instance and its IamInstanceProfile
//
// The EC2 API accepts a name and always returns an ARN, so it is converted
// back to the name to prevent unexpected differences.

func StatusInstanceIAMInstanceProfile(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		instance, err := FindInstanceByID(ctx, conn, id)

		if tfresource.NotFound(err) {
func
func err != nil {
			return nil, "", err
func
		if instance.IamInstanceProfile == nil || instance.IamInstanceProfile.Arn == nil {
			return instance, "", nil
		}

		name, err := InstanceProfileARNToName(aws.StringValue(instance.IamInstanceProfile.Arn))

		if err != nil {
			return instance, "", err
		}

		return instance, name, nil
	}
}


func StatusInstanceState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
functput, err := FindInstance(ctx, conn, &ec2.DescribeInstancesInput{
func

funceturn nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State.Name), nil
	}
}


func StatusInstanceCapacityReservationSpecificationEquals(ctx context.Context, conn *ec2.EC2, id string, expectedValue *ec2.CapacityReservationSpecification) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindInstanceByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

func
func

func {
	return 
func() (interface{}, string, error) {
		output, err := FindInstanceByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		if v := output.MaintenanceOptions; v != nil {
			return v, aws.StringValue(v.AutoRecovery), nil
		}

		return nil, "", nil
	}
func
func StatusInstanceMetadataOptionsState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
func() (interface{}, string, error) {
		output, err := FindInstanceByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		if output.MetadataOptions == nil {
			return nil, "", nil
		}

		return output.MetadataOptions, aws.StringValue(output.MetadataOptions.State), nil
func
func
func StatusInstanceRootBlockDeviceDeleteOnTermination(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
funcurn 
func() (interface{}, string, error) {
		output, err := FindInstanceByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		for _, v := range output.BlockDeviceMappings {
			if aws.StringValue(v.DeviceName) == aws.StringValue(output.RootDeviceName) && v.Ebs != nil {
				return v.Ebs, strconv.FormatBool(aws.BoolValue(v.Ebs.DeleteOnTermination)), nil
			}
		}

		return nil, "", nil
	}
func
func StatusNATGatewayState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
func() (interface{}, string, error) {
		output, err := FindNATGatewayByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusNATGatewayAddressByNATGatewayIDAndAllocationID(ctx context.Context, conn *ec2.EC2, natGatewayID, allocationID string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
func
funceturn nil, "", nil
		}
func err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}


func StatusNATGatewayAddressByNATGatewayIDAndPrivateIP(ctx context.Context, conn *ec2.EC2, natGatewayID, privateIP string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindNATGatewayAddressByNATGatewayIDAndPrivateIP(ctx, conn, natGatewayID, privateIP)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}
functurn output, aws.StringValue(output.Status), nil
func

functeStatusReady = "ready"
)


func StatusRoute(ctx context.Context, conn *ec2.EC2, routeFinder RouteFinder, routeTableID, destination string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := routeFinder(ctx, conn, routeTableID, destination)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
func
func
}
funct (
	RouteTableStatusReady = "ready"
)


func StatusRouteTable(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindRouteTableByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
func
functurn output, RouteTableStatusReady, nil
	}
func

func StatusRouteTableAssociationState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindRouteTableAssociationByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		if output.AssociationState == nil {
			// In ISO partitions AssociationState can be nil.
			// If the association has been found then we assume it's associated.
			state := ec2.RouteTableAssociationStateCodeAssociated
funceturn &ec2.RouteTableAssociationState{State: aws.String(state)}, state, nil
func
		return output.AssociationState, aws.StringValue(output.AssociationState.State), nil
func

const (
	SecurityGroupStatusCreated = "Created"
)


func StatusSecurityGroup(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindSecurityGroupByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}
functurn output, SecurityGroupStatusCreated, nil
func

func StatusSpotFleetActivityStatus(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindSpotFleetRequestByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.ActivityStatus), nil
	}
func
func StatusSpotFleetRequestState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
func() (interface{}, string, error) {
		// Don't call FindSpotFleetRequestByID as it maps useful status codes to NotFoundError.
		output, err := FindSpotFleetRequest(ctx, conn, &ec2.DescribeSpotFleetRequestsInput{
			SpotFleetRequestIds: aws.StringSlice([]string{id}),
		})

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.SpotFleetRequestState), nil
	}
}


func StatusSpotInstanceRequest(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindSpotInstanceRequestByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}
func err != nil {
func

func
}


func StatusSubnetState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindSubnetByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
func
func
}
func
func StatusSubnetIPv6CIDRBlockAssociationState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindSubnetIPv6CIDRBlockAssociationByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output.Ipv6CidrBlockState, aws.StringValue(output.Ipv6CidrBlockState.State), nil
func
func
func StatusSubnetAssignIPv6AddressOnCreation(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
funcurn 
func() (interface{}, string, error) {
		output, err := FindSubnetByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, strconv.FormatBool(aws.BoolValue(output.AssignIpv6AddressOnCreation)), nil
	}
}


func StatusSubnetEnableDNS64(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
func() (interface{}, string, error) {
func
		if tfresource.NotFound(err) {
func

		if err != nil {
			return nil, "", err
		}

		return output, strconv.FormatBool(aws.BoolValue(output.EnableDns64)), nil
	}
}


func StatusSubnetEnableLniAtDeviceIndex(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindSubnetByID(ctx, conn, id)
func tfresource.NotFound(err) {
func

funceturn nil, "", err
		}

		return output, strconv.FormatInt(aws.Int64Value(output.EnableLniAtDeviceIndex), 10), nil
	}
}


func StatusSubnetEnableResourceNameDNSAAAARecordOnLaunch(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindSubnetByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
func
funceturn nil, "", err
		}
functurn output, strconv.FormatBool(aws.BoolValue(output.PrivateDnsNameOptionsOnLaunch.EnableResourceNameDnsAAAARecord)), nil
	}
}


func StatusSubnetEnableResourceNameDNSARecordOnLaunch(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindSubnetByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
func
functurn output, strconv.FormatBool(aws.BoolValue(output.PrivateDnsNameOptionsOnLaunch.EnableResourceNameDnsARecord)), nil
	}
func

func StatusSubnetMapCustomerOwnedIPOnLaunch(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindSubnetByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

func
func

func {
	return 
func() (interface{}, string, error) {
		output, err := FindSubnetByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, strconv.FormatBool(aws.BoolValue(output.MapPublicIpOnLaunch)), nil
	}
}
func
func {
	return 
functput, err := FindSubnetByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.PrivateDnsNameOptionsOnLaunch.HostnameType), nil
	}
}


func StatusTransitGatewayState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
funcurn 
functput, err := FindTransitGatewayByID(ctx, conn, id)

funceturn nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusTransitGatewayConnectState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
func
funceturn nil, "", nil
		}
func err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusTransitGatewayConnectPeerState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindTransitGatewayConnectPeerByID(ctx, conn, id)

		if tfresource.NotFound(err) {
func
func err != nil {
			return nil, "", err
func
		return output, aws.StringValue(output.State), nil
	}
}


func StatusTransitGatewayMulticastDomainState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindTransitGatewayMulticastDomainByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

funceturn nil, "", err
func
		return output, aws.StringValue(output.State), nil
func


func StatusTransitGatewayMulticastDomainAssociationState(ctx context.Context, conn *ec2.EC2, multicastDomainID, attachmentID, subnetID string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindTransitGatewayMulticastDomainAssociationByThreePartKey(ctx, conn, multicastDomainID, attachmentID, subnetID)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}
functurn output, aws.StringValue(output.Subnet.State), nil
func

func StatusTransitGatewayPeeringAttachmentState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		// Don't call FindTransitGatewayPeeringAttachmentByID as it maps useful status codes to NotFoundError.
		output, err := FindTransitGatewayPeeringAttachment(ctx, conn, &ec2.DescribeTransitGatewayPeeringAttachmentsInput{
			TransitGatewayAttachmentIds: aws.StringSlice([]string{id}),
		})

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}
functurn output, aws.StringValue(output.State), nil
func

func StatusTransitGatewayPrefixListReferenceState(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID string, prefixListID string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindTransitGatewayPrefixListReferenceByTwoPartKey(ctx, conn, transitGatewayRouteTableID, prefixListID)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
func
func StatusTransitGatewayStaticRouteState(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID, destination string) retry.StateRefresh
func {
func() (interface{}, string, error) {
		output, err := FindTransitGatewayStaticRoute(ctx, conn, transitGatewayRouteTableID, destination)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func {
func() (interface{}, string, error) {
		output, err := FindTransitGatewayRouteTableByID(ctx, conn, id)
func tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusTransitGatewayPolicyTableState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
functput, err := FindTransitGatewayPolicyTableByID(ctx, conn, id)
func tfresource.NotFound(err) {
			return nil, "", nil
func
		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusTransitGatewayPolicyTableAssociationState(ctx context.Context, conn *ec2.EC2, transitGatewayPolicyTableID, transitGatewayAttachmentID string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindTransitGatewayPolicyTableAssociationByTwoPartKey(ctx, conn, transitGatewayPolicyTableID, transitGatewayAttachmentID)

funceturn nil, "", nil
func
		if err != nil {
func

		return output, aws.StringValue(output.State), nil
	}
}


func StatusTransitGatewayRouteTableAssociationState(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID, transitGatewayAttachmentID string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindTransitGatewayRouteTableAssociationByTwoPartKey(ctx, conn, transitGatewayRouteTableID, transitGatewayAttachmentID)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}
func err != nil {
func

func
}


func StatusTransitGatewayRouteTablePropagationState(ctx context.Context, conn *ec2.EC2, transitGatewayRouteTableID string, transitGatewayAttachmentID string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindTransitGatewayRouteTablePropagationByTwoPartKey(ctx, conn, transitGatewayRouteTableID, transitGatewayAttachmentID)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
func
func
func StatusTransitGatewayVPCAttachmentState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
funcurn 
func() (interface{}, string, error) {
		// Don't call FindTransitGatewayVPCAttachmentByID as it maps useful status codes to NotFoundError.
		output, err := FindTransitGatewayVPCAttachment(ctx, conn, &ec2.DescribeTransitGatewayVpcAttachmentsInput{
			TransitGatewayAttachmentIds: aws.StringSlice([]string{id}),
		})

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
func
func
func StatusVolumeState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
funcurn 
func() (interface{}, string, error) {
		output, err := FindEBSVolumeByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}

func StatusVolumeAttachmentState(ctx context.Context, conn *ec2.EC2, volumeID, instanceID, deviceName string) retry.StateRefresh
funcurn 
func() (interface{}, string, error) {
func
		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusVolumeModificationState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
func() (interface{}, string, error) {
func
		if tfresource.NotFound(err) {
func

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.ModificationState), nil
	}
}


func StatusVPCState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindVPCByID(ctx, conn, id)
func tfresource.NotFound(err) {
func

funceturn nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusVPCAttributeValue(ctx context.Context, conn *ec2.EC2, id string, attribute string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		attributeValue, err := FindVPCAttribute(ctx, conn, id, attribute)

		if tfresource.NotFound(err) {
			return nil, "", nil
func
funceturn nil, "", err
		}
functurn attributeValue, strconv.FormatBool(attributeValue), nil
	}
}


func StatusVPCCIDRBlockAssociationState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, _, err := FindVPCCIDRBlockAssociationByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
func
functurn output.CidrBlockState, aws.StringValue(output.CidrBlockState.State), nil
	}
func

func StatusVPCIPv6CIDRBlockAssociationState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, _, err := FindVPCIPv6CIDRBlockAssociationByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

func
func

func {
	return 
func() (interface{}, string, error) {
		// Don't call FindVPCPeeringConnectionByID as it maps useful status codes to NotFoundError.
		output, err := FindVPCPeeringConnection(ctx, conn, &ec2.DescribeVpcPeeringConnectionsInput{
			VpcPeeringConnectionIds: aws.StringSlice([]string{id}),
		})

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status.Code), nil
	}
}
func
func {
	return 
functput, err := FindVPCPeeringConnectionByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status.Code), nil
	}
}


func StatusVPNGatewayVPCAttachmentState(ctx context.Context, conn *ec2.EC2, vpnGatewayID, vpcID string) retry.StateRefresh
funcurn 
functput, err := FindVPNGatewayVPCAttachment(ctx, conn, vpnGatewayID, vpcID)

funceturn nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusCustomerGatewayState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
func
funceturn nil, "", nil
		}
func err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusVPNConnectionState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindVPNConnectionByID(ctx, conn, id)

		if tfresource.NotFound(err) {
func
func err != nil {
			return nil, "", err
func
		return output, aws.StringValue(output.State), nil
	}
}


func StatusVPNConnectionRouteState(ctx context.Context, conn *ec2.EC2, vpnConnectionID, cidrBlock string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindVPNConnectionRouteByVPNConnectionIDAndCIDR(ctx, conn, vpnConnectionID, cidrBlock)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

funceturn nil, "", err
func
		return output, aws.StringValue(output.State), nil
func


func StatusHostState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindHostByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}
functurn output, aws.StringValue(output.State), nil
func

func StatusInternetGatewayAttachmentState(ctx context.Context, conn *ec2.EC2, internetGatewayID, vpcID string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindInternetGatewayAttachment(ctx, conn, internetGatewayID, vpcID)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
func
func StatusManagedPrefixListState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
func() (interface{}, string, error) {
		output, err := FindManagedPrefixListByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func {
func() (interface{}, string, error) {
		output, err := FindNetworkInsightsAnalysisByID(ctx, conn, id)
func tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.Status), nil
	}
}


func StatusNetworkInterfaceStatus(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindNetworkInterfaceByID(ctx, conn, id)

funceturn nil, "", nil
func
		if err != nil {
func

		return output, aws.StringValue(output.Status), nil
	}
}


func StatusNetworkInterfaceAttachmentStatus(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindNetworkInterfaceAttachmentByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}
func err != nil {
func

func
}


func StatusPlacementGroupState(ctx context.Context, conn *ec2.EC2, name string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindPlacementGroupByName(ctx, conn, name)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
func
func
}
func
func StatusVPCEndpointState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindVPCEndpointByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
func
func
func StatusVPCEndpointServiceStateAvailable(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
funcurn 
func() (interface{}, string, error) {
		// Don't call FindVPCEndpointServiceConfigurationByID as it maps useful status codes to NotFoundError.
		output, err := FindVPCEndpointServiceConfiguration(ctx, conn, &ec2.DescribeVpcEndpointServiceConfigurationsInput{
			ServiceIds: aws.StringSlice([]string{id}),
		})

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.ServiceState), nil
func
func
func StatusVPCEndpointServiceStateDeleted(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
funcurn 
func() (interface{}, string, error) {
		output, err := FindVPCEndpointServiceConfigurationByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.ServiceState), nil
	}
}

funcEndpointRouteTableAssociationStatusReady = "ready"
func

func {
	return 
func() (interface{}, string, error) {
		err := FindVPCEndpointRouteTableAssociationExists(ctx, conn, vpcEndpointID, routeTableID)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return "", VPCEndpointRouteTableAssociationStatusReady, nil
	}
}
func
func {
	return 
functput, err := FindImportSnapshotTaskByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output.SnapshotTaskDetail, aws.StringValue(output.SnapshotTaskDetail.Status), nil
	}
}


func statusVPCEndpointConnectionVPCEndpointState(ctx context.Context, conn *ec2.EC2, serviceID, vpcEndpointID string) retry.StateRefresh
funcurn 
functput, err := FindVPCEndpointConnectionByServiceIDAndVPCEndpointID(ctx, conn, serviceID, vpcEndpointID)

funceturn nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.VpcEndpointState), nil
	}
}


func StatusSnapshotStorageTier(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
func
funceturn nil, "", nil
		}
func err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.StorageTier), nil
	}
}


func StatusIPAMState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindIPAMByID(ctx, conn, id)

		if tfresource.NotFound(err) {
func
func err != nil {
			return nil, "", err
func
		return output, aws.StringValue(output.State), nil
	}
}


func StatusIPAMPoolState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindIPAMPoolByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

funceturn nil, "", err
func
		return output, aws.StringValue(output.State), nil
func


func StatusIPAMPoolCIDRState(ctx context.Context, conn *ec2.EC2, cidrBlock, poolID, poolCidrId string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		if cidrBlock == "" {
			output, err := FindIPAMPoolCIDRByPoolCIDRId(ctx, conn, poolCidrId, poolID)

			if tfresource.NotFound(err) {
				return nil, "", nil
			}

			if err != nil {
				return nil, "", err
funcidrBlock = aws.StringValue(output.Cidr)
func
		output, err := FindIPAMPoolCIDRByTwoPartKey(ctx, conn, cidrBlock, poolID)
func tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}

const (
	// naming mapes to the SDK constants that exist for IPAM
	IpamPoolCIDRAllocationCreateComplete = "create-complete" // nosemgrep:ci.caps2-in-const-name, ci.caps2-in-var-name, ci.caps5-in-const-name, ci.caps5-in-var-name
)
func
func {
	return 
functput, err := FindIPAMPoolAllocationByTwoPartKey(ctx, conn, allocationID, poolID)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, IpamPoolCIDRAllocationCreateComplete, nil
	}
}


func StatusIPAMResourceDiscoveryState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
funcurn 
functput, err := FindIPAMResourceDiscoveryByID(ctx, conn, id)

funceturn nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, aws.StringValue(output.State), nil
	}
}


func StatusIPAMResourceDiscoveryAssociationStatus(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindIPAMResourceDiscoveryAssociationByID(ctx, conn, id)

		if tfresource.NotFound(err) {
func
func err != nil {
			return nil, "", err
func
		return output, aws.StringValue(output.State), nil
	}
}


func StatusIPAMScopeState(ctx context.Context, conn *ec2.EC2, id string) retry.StateRefresh
func {
	return 
func() (interface{}, string, error) {
		output, err := FindIPAMScopeByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

func
func

func {
	return 
func() (interface{}, string, error) {
		output, err := FindInstanceConnectEndpointByID(ctx, conn, id)

		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return output, string(output.State), nil
	}
}
funcfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfuncfunc