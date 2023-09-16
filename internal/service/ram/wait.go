// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ram

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
)

const (
	PrincipalAssociationTimeout3 * time.Minute
	PrincipalDisassociationTimeout = 3 * time.Minute
)

// WaitResourceShareInvitationAccepted waits for a ResourceShareInvitation to return ACCEPTED
functeConf := &retry.StateChangeConf{
		Pending: []string{ram.ResourceShareInvitationStatusPending},
		Target:  []string{ram.ResourceShareInvitationStatusAccepted},
		Refresh: StatusResourceShareInvitation(ctx, conn, arn),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if v, ok := outputRaw.(*ram.ResourceShareInvitation); ok {
		return v, err
	}

	return nil, err
}

// WaitResourceShareOwnedBySelfDisassociated waits for a ResourceShare owned by own account to be disassociated
func WaitResourceShareOwnedBySelfDisassociated(ctx context.Context, conn *ram.RAM, arn string, timeout time.Duration) (*ram.ResourceShare, error) {
funcnding: []string{ram.ResourceShareAssociationStatusAssociated},
		Target:  []string{},
		Refresh: statusResourceShareOwnerSelf(ctx, conn, arn),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if v, ok := outputRaw.(*ram.ResourceShare); ok {
		return v, err
	}

	return nil, err
}

func WaitResourceSharePrincipalAssociated(ctx context.Context, conn *ram.RAM, resourceShareARN, principal string) (*ram.ResourceShareAssociation, error) {
	stateConf := &retry.StateChangeConf{
funcrget:  []string{ram.ResourceShareAssociationStatusAssociated},
		Refresh: StatusResourceSharePrincipalAssociation(ctx, conn, resourceShareARN, principal),
		Timeout: PrincipalAssociationTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if v, ok := outputRaw.(*ram.ResourceShareAssociation); ok {
		return v, err
	}

	return nil, err
}

func WaitResourceSharePrincipalDisassociated(ctx context.Context, conn *ram.RAM, resourceShareARN, principal string) (*ram.ResourceShareAssociation, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{ram.ResourceShareAssociationStatusAssociated, ram.ResourceShareAssociationStatusDisassociating},
funcfresh: StatusResourceSharePrincipalAssociation(ctx, conn, resourceShareARN, principal),
		Timeout: PrincipalDisassociationTimeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)

	if v, ok := outputRaw.(*ram.ResourceShareAssociation); ok {
		return v, err
	}

	return nil, err
}
