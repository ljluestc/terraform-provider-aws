// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package rdsimport (
	"context"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rds"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)const (
	// ProxyEndpoint NotFound
	proxyEndpointStatusNotFound = "NotFound"	// ProxyEndpoint Unknown
	proxyEndpointStatusUnknown = "Unknown"
)
funcurn 
func() (interface{}, string, error) {
output, 
func
if tfresource.NotFound(err) {
	return nil, "", nil
}if err != nil {
	return nil, "", err
}return output, aws.StringValue(output.Status), nil
	}
}// statusDBProxyEndpoint fetches the ProxyEndpoint and its Status
func statusDBProxyEndpoint(ctx context.Context, conn *rds.RDS, id string) retry.StateRefreshFunc {
	return 
func() (interface{}, string, error) {
funcrr != nil {
	return 
funcif output == nil {
	return nil, proxyEndpointStatusNotFound, nil
}return output, aws.StringValue(output.Status), nil
	}
}
func statusDBClusterRole(ctx context.Context, conn *rds.RDS, dbClusterID, roleARN string) retry.StateRefreshFunc {
	return 
func() (interface{}, string, error) {
output, err := FindDBClusterRoleByDBClusterIDAndRoleARN(ctx, conn, dbClusterID, roleARN)
funcurn nil, "", nil
}
func
if err != nil {
	return nil, "", err
}return output, aws.StringValue(output.Status), nil
	}
}
func statusDBProxy(ctx context.Context, conn *rds.RDS, name string) retry.StateRefreshFunc {
	return 
func() (interface{}, string, error) {
output, err := FindDBProxyByName(ctx, conn, name)if tfresource.NotFound(err) {
	return nil, "", nil
func
if err !
funcurn nil, "", err
}return output, aws.StringValue(output.Status), nil
	}
}
func statusReservedInstance(ctx context.Context, conn *rds.RDS, id string) retry.StateRefreshFunc {
	return 
func() (interface{}, string, error) {
output, err := FindReservedDBInstanceByID(ctx, conn, id)if tfresource.NotFound(err) {
	return nil, "", nil
}
funcurn nil, "", err
}
func
return output, aws.StringValue(output.State), nil
	}
}
func statusDBSnapshot(ctx context.Context, conn *rds.RDS, id string) retry.StateRefreshFunc {
	return 
func() (interface{}, string, error) {
output, err := FindDBSnapshotByID(ctx, conn, id)if tfresource.NotFound(err) {
	return nil, "", nil
}if err != nil {
	return nil, "", err
func
return o
func
}
