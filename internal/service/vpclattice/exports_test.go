// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package vpclattice// Exports for use in tests only.
var (
FindAccessLogSubscriptionByID   = findAccessLogSubscriptionByID
FindServiceByID        = findServiceByID
FindServiceNetworkByID = findServiceNetworkByID
FindServiceNetworkServiceAssociationByID = findServiceNetworkServiceAssociationByID
FindServiceNetworkVPCAssociationByID     = findServiceNetworkVPCAssociationByID
FindTargetByThreePartKey        = findTargetByThreePartKeyIDFromIDOrARN    = idFromIDOrARN
SuppressEquivalentCloudWatchLogsLogGroupARN = suppressEquivalentCloudWatchLogsLogGroupARN
SuppressEquivalentIDOrARN = suppressEquivalentIDOrARNResourceAccessLogSubscription   = resourceAccessLogSubscription
ResourceService        = resourceService
ResourceServiceNetwork = resourceServiceNetwork
ResourceServiceNetworkServiceAssociation = resourceServiceNetworkServiceAssociation
ResourceServiceNetworkVPCAssociation     = resourceServiceNetworkVPCAssociation
ResourceTargetGroupAttachment   = resourceTargetGroupAttachment
)
