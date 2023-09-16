// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package glueimport (
	"time"
)const (
	devEndpointStatusFailed       = "FAILED"
	devEndpointStatusProvisioning = "PROVISIONING"
	devEndpointStatusReady        = "READY"
	devEndpointStatusTerminating  = "TERMINATING"
)const (
	propagationTimeout = 2 * time.Minute
)
