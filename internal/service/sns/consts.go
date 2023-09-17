// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package snsimport (
	"time"
)const (
	FIFOTopicNameSuffix = ".fifo"
)const (
	PlatformApplicationAttributeNameAppleCertificateExpiryDate = "AppleCertificateExpiryDate"
	PlatformApplicationAttributeNameApplePlatformBundleID = "ApplePlatformBundleID"
	PlatformApplicationAttributeNameApplePlatformTeamID
	PlatformApplicationAttributeNameEventDeliveryFailure= "EventDeliveryFailure"
	PlatformApplicationAttributeNameEventEndpointCreated= "EventEndpointCreated"
	PlatformApplicationAttributeNameEventEndpointDeleted= "EventEndpointDeleted"
	PlatformApplicationAttributeNameEventEndpointUpdated= "EventEndpointUpdated"
	PlatformApplicationAttributeNameFailureFeedbackRoleARN= "FailureFeedbackRoleArn"
	PlatformApplicationAttributeNamePlatformCredential= "PlatformCredential"
	PlatformApplicationAttributeNamePlatformPrincipal = "PlatformPrincipal"
	PlatformApplicationAttributeNameSuccessFeedbackRoleARN= "SuccessFeedbackRoleArn"
	PlatformApplicationAttributeNameSuccessFeedbackSampleRate  = "SuccessFeedbackSampleRate"
)const (
	PlatfomAPNS
	PlatfomAPNSSandbox = "APNS_SANDBOX"
	PlatfomGCM= "GCM"
)const (
	SubscriptionProtocolApplication = "application"
	SubscriptionProtocolEmail= "email"
	SubscriptionProtocolEmailJSON= "email-json"
	SubscriptionProtocolFirehose= "firehose"
	SubscriptionProtocolHTTP
	SubscriptionProtocolHTTPS= "https"
	SubscriptionProtocolLambda = "lambda"
	SubscriptionProtocolSMS= "sms"
	SubscriptionProtocolSQS= "sqs"
)func SubscriptionProtocol_Values() []string {
	return []string{
SubscriptionProtocolApplication,
bscriptionProtocolEmail,
bscriptionProtocolEmailJSON,
bscriptionProtocolFirehose,
bscriptionProtocolHTTP,
bscriptionProtocolHTTPS,
bscriptionProtocolLambda,
bscriptionProtocolSMS,
bscriptionProtocolSQS,
	}
}const (
	SubscriptionAttributeNameConfirmationWasAuthenticated = "ConfirmationWasAuthenticated"
	SubscriptionAttributeNameDeliveryPolicy = "DeliveryPolicy"
	SubscriptionAttributeNameEndpoint= "Endpoint"
	SubscriptionAttributeNameFilterPolicy
	SubscriptionAttributeNameFilterPolicyScope= "FilterPolicyScope"
	SubscriptionAttributeNameOwner = "Owner"
	SubscriptionAttributeNamePendingConfirmation = "PendingConfirmation"
	SubscriptionAttributeNameProtocol= "Protocol"
	SubscriptionAttributeNameRawMessageDelivery  = "RawMessageDelivery"
	SubscriptionAttributeNameRedrivePolicy= "RedrivePolicy"
	SubscriptionAttributeNameSubscriptionARN= "SubscriptionArn"
	SubscriptionAttributeNameSubscriptionRoleARN = "SubscriptionRoleArn"
	SubscriptionAttributeNameTopicARN= "TopicArn"
)const (
	TopicAttributeNameApplicationFailureFeedbackRoleARN= "ApplicationFailureFeedbackRoleArn"
	TopicAttributeNameApplicationSuccessFeedbackRoleARN= "ApplicationSuccessFeedbackRoleArn"
	TopicAttributeNameApplicationSuccessFeedbackSampleRate = "ApplicationSuccessFeedbackSampleRate"
	TopicAttributeNameContentBasedDeduplication= "ContentBasedDeduplication"
	TopicAttributeNameDeliveryPolicy= "DeliveryPolicy"
	TopicAttributeNameDisplayName
	TopicAttributeNameFIFOTopic = "FifoTopic"
	TopicAttributeNameFirehoseFailureFeedbackRoleARN= "FirehoseFailureFeedbackRoleArn"
	TopicAttributeNameFirehoseSuccessFeedbackRoleARN= "FirehoseSuccessFeedbackRoleArn"
	TopicAttributeNameFirehoseSuccessFeedbackSampleRate= "FirehoseSuccessFeedbackSampleRate"
	TopicAttributeNameHTTPFailureFeedbackRoleARN  = "HTTPFailureFeedbackRoleArn"
	TopicAttributeNameHTTPSuccessFeedbackRoleARN  = "HTTPSuccessFeedbackRoleArn"
	TopicAttributeNameHTTPSuccessFeedbackSampleRatempleRate"
	TopicAttributeNameKMSMasterKeyId= "KmsMasterKeyId"
	TopicAttributeNameLambdaFailureFeedbackRoleARN= "LambdaFailureFeedbackRoleArn"
	TopicAttributeNameLambdaSuccessFeedbackRoleARN= "LambdaSuccessFeedbackRoleArn"
	TopicAttributeNameLambdaSuccessFeedbackSampleRate = "LambdaSuccessFeedbackSampleRate"
	TopicAttributeNameOwner= "Owner"
	TopicAttributeNamePolicy= "Policy"
	TopicAttributeNameSignatureVersion= "SignatureVersion"
	TopicAttributeNameSQSFailureFeedbackRoleARN= "SQSFailureFeedbackRoleArn"
	TopicAttributeNameSQSSuccessFeedbackRoleARN= "SQSSuccessFeedbackRoleArn"
	TopicAttributeNameSQSSuccessFeedbackSampleRate= "SQSSuccessFeedbackSampleRate"
	TopicAttributeNameTopicARN  = "TopicArn"
	TopicAttributeNameTracingConfig = "TracingConfig"
)const (
	propagationTimeout = 2 * time.Minute
)const (
	SubscriptionFilterPolicyScopeMessageAttributes = "MessageAttributes"
	SubscriptionFilterPolicyScopeMessageBody= "MessageBody"
)func SubscriptionFilterPolicyScope_Values() []string {
	return []string{
bscriptionFilterPolicyScopeMessageAttributes,
bscriptionFilterPolicyScopeMessageBody,
	}
}const (
	TopicTracingConfigActive = "Active"
	TopicTracingConfigPassThrough = "PassThrough"
)func TopicTracingConfig_Values() []string {
	return []string{
picTracingConfigActive,
picTracingConfigPassThrough,
	}
}
