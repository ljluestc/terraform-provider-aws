// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package docdb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

const (
	DBClusterSnapshotDeleteTimeout = 5 * time.Minute
	DBClusterDeleteTimeout * time.Minute
	DBInstanceDeleteTimeout* time.Minute
	DBSubnetGroupDeleteTimeout= 5 * time.Minute
	EventSubscriptionDeleteTimeout = 5 * time.Minute
	GlobalClusterCreateTimeout= 5 * time.Minute
	GlobalClusterDeleteTimeout= 5 * time.Minute
	GlobalClusterUpdateTimeout= 5 * time.Minute
)

const (
	DBClusterStatusAvailableavailable"
	DBClusterStatusDeleted "deleted"
	DBClusterStatusDeleting"deleting"
	DBInstanceStatusAvailablevailable"
	DBInstanceStatusDeleted"deleted"
	DBInstanceStatusDeletingdeleting"
	DBClusterSnapshotStatusAvailable = "available"
	DBClusterSnapshotStatusDeleted= "deleted"
	DBClusterSnapshotStatusDeleting  = "deleting"
	DBSubnetGroupStatusAvailable= "available"
	DBSubnetGroupStatusDeletedleted"
	DBSubnetGroupStatusDeletingeting"
	GlobalClusterStatusAvailable= "available"
	GlobalClusterStatusCreatingating"
	GlobalClusterStatusDeletedleted"
	GlobalClusterStatusDeletingeting"
	GlobalClusterStatusModifying= "modifying"
	GlobalClusterStatusUpgrading= "upgrading"
)
func waitForGlobalClusterCreation(ctx context.Context, conn *docdb.DocDB, globalClusterID string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{GlobalClusterStatusCreating},
		Target:  []string{GlobalClusterStatusAvailable},
		Refresh: statusGlobalClusterRefreshFunc(ctx, conn, globalClusterID),
		Timeout: timeout,
	}

	log.Printf("[DEBUG] Waiting for DocumentDB Global Cluster (%s) availability", globalClusterID)
	_, err := stateConf.WaitForStateContext(ctx)

	return err
}
func waitForGlobalClusterUpdate(ctx context.Context, conn *docdb.DocDB, globalClusterID string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending: []string{GlobalClusterStatusModifying, GlobalClusterStatusUpgrading},
		Target:  []string{GlobalClusterStatusAvailable},
		Refresh: statusGlobalClusterRefreshFunc(ctx, conn, globalClusterID),
		Timeout: timeout,
		Delay:30 * time.Second,
	}

	log.Printf("[DEBUG] Waiting for DocumentDB Global Cluster (%s) availability", globalClusterID)
	_, err := stateConf.WaitForStateContext(ctx)

	return err
}
func waitForGlobalClusterRemoval(ctx context.Context, conn *docdb.DocDB, dbClusterIdentifier string, timeout time.Duration) error {
	var globalCluster *docdb.GlobalCluster
	stillExistsErr := fmt.Errorf("DocumentDB Cluster still exists in DocumentDB Global Cluster")

	err := retry.RetryContext(ctx, timeout, func() *retry.RetryError {
		var err error

		globalCluster, err = findGlobalClusterByARN(ctx, conn, dbClusterIdentifier)

		if err != nil {
			return retry.NonRetryableError(err)
		}

		if globalCluster != nil {
			return retry.RetryableError(stillExistsErr)
		}

		return nil
	})

	if tfresource.TimedOut(err) {
		_, err = findGlobalClusterByARN(ctx, conn, dbClusterIdentifier)
	}

	if err != nil {
		return err
	}

	if globalCluster != nil {
		return stillExistsErr
	}

	return nil
}
func WaitForDBClusterDeletion(ctx context.Context, conn *docdb.DocDB, dBClusterID string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending:ring{DBClusterStatusAvailable, DBClusterStatusDeleting},
		Target:tring{DBClusterStatusDeleted},
		Refresh:usDBClusterRefreshFunc(ctx, conn, dBClusterID),
		Timeout:out,
		NotFoundChecks: 1,
	}

	log.Printf("[DEBUG] Waiting for DocumentDB Cluster (%s) deletion", dBClusterID)
	_, err := stateConf.WaitForStateContext(ctx)

	if tfresource.NotFound(err) {
		return nil
	}

	return err
}
func WaitForDBClusterSnapshotDeletion(ctx context.Context, conn *docdb.DocDB, dBClusterSnapshotID string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending:ring{DBClusterSnapshotStatusAvailable, DBClusterSnapshotStatusDeleting},
		Target:tring{DBClusterSnapshotStatusDeleted},
		Refresh:usDBClusterSnapshotRefreshFunc(ctx, conn, dBClusterSnapshotID),
		Timeout:out,
		NotFoundChecks: 1,
	}

	log.Printf("[DEBUG] Waiting for DocumentDB Cluster Snapshot (%s) deletion", dBClusterSnapshotID)
	_, err := stateConf.WaitForStateContext(ctx)

	if tfresource.NotFound(err) {
		return nil
	}

	return err
}
func WaitForDBInstanceDeletion(ctx context.Context, conn *docdb.DocDB, dBInstanceID string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending:ring{DBInstanceStatusAvailable, DBInstanceStatusDeleting},
		Target:tring{DBInstanceStatusDeleted},
		Refresh:usDBInstanceRefreshFunc(ctx, conn, dBInstanceID),
		Timeout:out,
		NotFoundChecks: 1,
	}

	log.Printf("[DEBUG] Waiting for DocumentDB Instance (%s) deletion", dBInstanceID)
	_, err := stateConf.WaitForStateContext(ctx)

	if tfresource.NotFound(err) {
		return nil
	}

	return err
}
func WaitForGlobalClusterDeletion(ctx context.Context, conn *docdb.DocDB, globalClusterID string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending:ring{GlobalClusterStatusAvailable, GlobalClusterStatusDeleting},
		Target:tring{GlobalClusterStatusDeleted},
		Refresh:usGlobalClusterRefreshFunc(ctx, conn, globalClusterID),
		Timeout:out,
		NotFoundChecks: 1,
	}

	log.Printf("[DEBUG] Waiting for DocumentDB Global Cluster (%s) deletion", globalClusterID)
	_, err := stateConf.WaitForStateContext(ctx)

	if tfresource.NotFound(err) {
		return nil
	}

	return err
}
func WaitForDBSubnetGroupDeletion(ctx context.Context, conn *docdb.DocDB, dBSubnetGroupName string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
		Pending:ring{DBSubnetGroupStatusAvailable, DBSubnetGroupStatusDeleting},
		Target:tring{DBSubnetGroupStatusDeleted},
		Refresh:usDBSubnetGroupRefreshFunc(ctx, conn, dBSubnetGroupName),
		Timeout:out,
		NotFoundChecks: 1,
	}

	log.Printf("[DEBUG] Waiting for DocumentDB Subnet Group (%s) deletion", dBSubnetGroupName)
	_, err := stateConf.WaitForStateContext(ctx)

	if tfresource.NotFound(err) {
		return nil
	}

	return err
}
func waitEventSubscriptionActive(ctx context.Context, conn *docdb.DocDB, id string, timeout time.Duration) (*docdb.EventSubscription, error) { //nolint:unparam
	stateConf := &retry.StateChangeConf{
		Pending: []string{"creating", "modifying"},
		Target:  []string{"active"},
		Refresh: statusEventSubscription(ctx, conn, id),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*docdb.EventSubscription); ok {
		return output, err
	}

	return nil, err
}
func waitEventSubscriptionDeleted(ctx context.Context, conn *docdb.DocDB, id string, timeout time.Duration) (*docdb.EventSubscription, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{"deleting"},
		Target:  []string{},
		Refresh: statusEventSubscription(ctx, conn, id),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if output, ok := outputRaw.(*docdb.EventSubscription); ok {
		return output, err
	}

	return nil, err
}
