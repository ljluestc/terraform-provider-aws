// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package emrimport (
	"context"
	"fmt"
	"time"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)const (
	ClusterCreatedTimeout= 75 * time.Minute
	ClusterCreatedMinTimeout = 10 * time.Second
	ClusterCreatedDelay = 30 * time.Second	ClusterDeletedTimeout= 20 * time.Minute
	ClusterDeletedMinTimeout = 10 * time.Second
	ClusterDeletedDelay = 30 * time.Second
)func waitClusterCreated(ctx context.Context, conn *emr.EMR, id string) (*emr.Cluster, error) {
	stateConf := &retry.StateChangeConf{
Pending:[]string{emr.ClusterStateBootstrapping, emr.ClusterStateStarting},
rget:[]string{emr.ClusterStateRunning, emr.ClusterStateWaiting},
fresh:statusCluster(ctx, conn, id),
meout:ClusterCreatedTimeout,
nTimeout: ClusterCreatedMinTimeout,
lay: ClusterCreatedDelay,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)	if output, ok := outputRaw.(*emr.Cluster); ok {
 stateChangeReason := output.Status.StateChangeReason; stateChangeReason != nil {
tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(stateChangeReason.Code), aws.StringValue(stateChangeReason.Message)))
ern output, err
	}	return nil, err
}func waitClusterDeleted(ctx context.Context, conn *emr.EMR, id string) (*emr.Cluster, error) {
	stateConf := &retry.StateChangeConf{
nding:[]string{emr.ClusterStateTerminating},
rget:[]string{emr.ClusterStateTerminated, emr.ClusterStateTerminatedWithErrors},
fresh:statusCluster(ctx, conn, id),
meout:ClusterDeletedTimeout,
nTimeout: ClusterDeletedMinTimeout,
lay: ClusterDeletedDelay,
	}	outputRaw, err := stateConf.WaitForStateContext(ctx)	if output, ok := outputRaw.(*emr.Cluster); ok {
 stateChangeReason := output.Status.StateChangeReason; stateChangeReason != nil {
tfresource.SetLastError(err, fmt.Errorf("%s: %s", aws.StringValue(stateChangeReason.Code), aws.StringValue(stateChangeReason.Message)))
ern output, err
	}	return nil, err
}
