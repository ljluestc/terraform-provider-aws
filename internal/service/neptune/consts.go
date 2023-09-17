// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package neptuneimport (
"time"
)const (
propagationTimeout = 2 * time.Minute
)const (
GlobalClusterStatusAvailable = "available"
GlobalClusterStatusCreating  = "creating"
GlobalClusterStatusDeleted   = "deleted"
GlobalClusterStatusDeleting  = "deleting"
GlobalClusterStatusModifying = "modifying"
GlobalClusterStatusUpgrading = "upgrading"
)
