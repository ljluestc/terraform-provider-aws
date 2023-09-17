// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package sqsconst (
FIFOQueueNameSuffix = ".fifo"
)const (
DefaultQueueDelaySeconds= 0
DefaultQueueKMSDataKeyReusePeriodSeconds  = 300
DefaultQueueMaximumMessageSize= 262_144 // 256 KiB.
DefaultQueueMessageRetentionPeriod
DefaultQueueReceiveMessageWaitTimeSeconds = 0
DefaultQueueVisibilityTimeout= 30
)const (
DeduplicationScopeMessageGroup = "messageGroup"
DeduplicationScopeQueue
)func DeduplicationScope_Values() []string {
return []string{
DeduplicationScopeMessageGroup,
DeduplicationScopeQueue,
}
}const (
FIFOThroughputLimitPerMessageGroupID = "perMessageGroupId"
FIFOThroughputLimitPerQueue = "perQueue"
)func FIFOThroughputLimit_Values() []string {
return []string{
FIFOThroughputLimitPerMessageGroupID,
FIFOThroughputLimitPerQueue,
}
}
