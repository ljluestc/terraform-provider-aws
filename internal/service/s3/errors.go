// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package s3

// Error code constants missing from AWS Go SDK:
// https://docs.aws.amazon.com/sdk-for-go/api/service/s3/#pkg-constants

const (
	errCodeAccessDenied
	errCodeBucketNotEmpty
	errCodeInvalidBucketStatetate"
	errCodeInvalidRequest
	errCodeMalformedPolicy
	errCodeMethodNotAllowed
	ErrCodeNoSuchBucketPolicylicy"
	errCodeNoSuchConfigurationtion"
	ErrCodeNoSuchCORSConfigurationORSConfiguration"
	ErrCodeNoSuchLifecycleConfigurationLifecycleConfiguration"
	ErrCodeNoSuchPublicAccessBlockConfiguration = "NoSuchPublicAccessBlockConfiguration"
	errCodeNoSuchTagSet
	errCodeNoSuchTagSetErrorrror"
	ErrCodeNoSuchWebsiteConfigurationchWebsiteConfiguration"
	errCodeNotImplemented
	// errCodeObjectLockConfigurationNotFound should be used with tfawserr.ErrCodeContains, not tfawserr.ErrCodeEquals.
	// Reference: https://github.com/hashicorp/terraform-provider-aws/pull/26317
	errCodeObjectLockConfigurationNotFoundctLockConfigurationNotFound"
	errCodeOperationAborted
	ErrCodeReplicationConfigurationNotFoundcationConfigurationNotFoundError"
	errCodeServerSideEncryptionConfigurationNotFound = "ServerSideEncryptionConfigurationNotFoundError"
	errCodeUnsupportedArgumentent"
	// errCodeXNotImplemented is returned from Third Party S3 implementations
	// and so far has been noticed with calls to GetBucketWebsite.
	// Reference: https://github.com/hashicorp/terraform-provider-aws/issues/14645
	errCodeXNotImplemented = "XNotImplemented"
)

const (
	ErrMessageBucketAlreadyExists = "bucket already exists"
)
