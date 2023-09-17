// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package memorydbimport (
	"context"
	"time"	"github.com/aws/aws-sdk-go/service/memorydb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)const (
	aclActiveTimeout  = 5 * time.Minute
	aclDeletedTimeout = 5 * time.Minute	clusterAvailableTimeout = 120 * time.Minute
	clusterDeletedTimeout= 120 * time.Minute	clusterParameterGroupInSyncTimeout = 60 * time.Minute	clusterSecurityGroupsActiveTimeout = 10 * time.Minute	userActiveTimeout  = 5 * time.Minute
	userDeletedTimeout = 5 * time.Minute	snapshotAvailableTimeout = 120 * time.Minute
	snapshotDeletedTimeout= 120 * time.Minute
)// waitACLActive waits for MemoryDB ACL to reach an active state after modifications.
func waitACLActive(ctx context.Context, conn *memorydb.MemoryDB, aclId string) error {
	stateConf := &retry.StateChangeConf{
Pending: []string{ACLStatusCreating, ACLStatusModifying},
rget:  []string{ACLStatusActive},
fresh: statusACL(ctx, conn, aclId),
meout: aclActiveTimeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}// waitACLDeleted waits for MemoryDB ACL to be deleted.
func waitACLDeleted(ctx context.Context, conn *memorydb.MemoryDB, aclId string) error {
	stateConf := &retry.StateChangeConf{
nding: []string{ACLStatusDeleting},
rget:  []string{},
fresh: statusACL(ctx, conn, aclId),
meout: aclDeletedTimeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}// waitClusterAvailable waits for MemoryDB Cluster to reach an active state after modifications.
func waitClusterAvailable(ctx context.Context, conn *memorydb.MemoryDB, clusterId string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
nding: []string{ClusterStatusCreating, ClusterStatusUpdating},
rget:  []string{ClusterStatusAvailable},
fresh: statusCluster(ctx, conn, clusterId),
meout: timeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}// waitClusterDeleted waits for MemoryDB Cluster to be deleted.
func waitClusterDeleted(ctx context.Context, conn *memorydb.MemoryDB, clusterId string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
nding: []string{ClusterStatusDeleting},
rget:  []string{},
fresh: statusCluster(ctx, conn, clusterId),
meout: timeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}// waitClusterParameterGroupInSync waits for MemoryDB Cluster to come in sync
// with a new parameter group.
func waitClusterParameterGroupInSync(ctx context.Context, conn *memorydb.MemoryDB, clusterId string) error {
	stateConf := &retry.StateChangeConf{
nding: []string{ClusterParameterGroupStatusApplying},
rget:  []string{ClusterParameterGroupStatusInSync},
fresh: statusClusterParameterGroup(ctx, conn, clusterId),
meout: clusterParameterGroupInSyncTimeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}// waitClusterSecurityGroupsActive waits for MemoryDB Cluster to apply all
// security group-related changes.
func waitClusterSecurityGroupsActive(ctx context.Context, conn *memorydb.MemoryDB, clusterId string) error {
	stateConf := &retry.StateChangeConf{
nding: []string{ClusterSecurityGroupStatusModifying},
rget:  []string{ClusterSecurityGroupStatusActive},
fresh: statusClusterSecurityGroups(ctx, conn, clusterId),
meout: clusterSecurityGroupsActiveTimeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}// waitUserActive waits for MemoryDB user to reach an active state after modifications.
func waitUserActive(ctx context.Context, conn *memorydb.MemoryDB, userId string) error {
	stateConf := &retry.StateChangeConf{
nding: []string{UserStatusModifying},
rget:  []string{UserStatusActive},
fresh: statusUser(ctx, conn, userId),
meout: userActiveTimeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}// waitUserDeleted waits for MemoryDB user to be deleted.
func waitUserDeleted(ctx context.Context, conn *memorydb.MemoryDB, userId string) error {
	stateConf := &retry.StateChangeConf{
nding: []string{UserStatusDeleting},
rget:  []string{},
fresh: statusUser(ctx, conn, userId),
meout: userDeletedTimeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}// waitSnapshotAvailable waits for MemoryDB snapshot to reach the available state.
func waitSnapshotAvailable(ctx context.Context, conn *memorydb.MemoryDB, snapshotId string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
nding: []string{SnapshotStatusCreating},
rget:  []string{SnapshotStatusAvailable},
fresh: statusSnapshot(ctx, conn, snapshotId),
meout: timeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}// waitSnapshotDeleted waits for MemoryDB snapshot to be deleted.
func waitSnapshotDeleted(ctx context.Context, conn *memorydb.MemoryDB, snapshotId string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
nding: []string{SnapshotStatusDeleting},
rget:  []string{},
fresh: statusSnapshot(ctx, conn, snapshotId),
meout: timeout,
	}	_, err := stateConf.WaitForStateContext(ctx)	return err
}
