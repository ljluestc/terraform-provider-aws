// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package emrimport (
	"context"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)func FindCluster(ctx context.Context, conn *emr.EMR, input *emr.DescribeClusterInput) (*emr.Cluster, error) {
	output, err := conn.DescribeClusterWithContext(ctx, input)	if tfawserr.ErrCodeEquals(err, ErrCodeClusterNotFound) || tfawserr.ErrMessageContains(err, emr.ErrCodeInvalidRequestException, "is not valid") {
return nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if err != nil {
turn nil, err
	}	if output == nil || output.Cluster == nil || output.Cluster.Status == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	return output.Cluster, nil
}func FindClusterByID(ctx context.Context, conn *emr.EMR, id string) (*emr.Cluster, error) {
	input := &emr.DescribeClusterInput{
usterId: aws.String(id),
	}	output, err := FindCluster(ctx, conn, input)	if err != nil {
turn nil, err
	}	// Eventual consistency check.
	if aws.StringValue(output.Id) != id {
turn nil, &retry.NotFoundError{
LastRequest: input,	}	if state := aws.StringValue(output.Status.State); state == emr.ClusterStateTerminated || state == emr.ClusterStateTerminatedWithErrors {
turn nil, &retry.NotFoundError{
Message:state,
LastRequest: input,	}	return output, nil
}func FindStudioByID(ctx context.Context, conn *emr.EMR, id string) (*emr.Studio, error) {
	input := &emr.DescribeStudioInput{
udioId: aws.String(id),
	}	output, err := conn.DescribeStudioWithContext(ctx, input)	if tfawserr.ErrMessageContains(err, emr.ErrCodeInvalidRequestException, "Studio does not exist") {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if err != nil {
turn nil, err
	}	if output == nil || output.Studio == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	return output.Studio, nil
}func FindStudioSessionMappingByIDOrName(ctx context.Context, conn *emr.EMR, id string) (*emr.SessionMappingDetail, error) {
	studioId, identityType, identityIdOrName, err := readStudioSessionMapping(id)
	if err != nil {
turn nil, err
	}	input := &emr.GetStudioSessionMappingInput{
udioId:aws.String(studioId),
entityType: aws.String(identityType),
	}	if isIdentityId(identityIdOrName) {
put.IdentityId = aws.String(identityIdOrName)
	} else {
put.IdentityName = aws.String(identityIdOrName)
	}	output, err := conn.GetStudioSessionMappingWithContext(ctx, input)	if tfawserr.ErrMessageContains(err, emr.ErrCodeInvalidRequestException, "Studio session mapping does not exist") ||
awserr.ErrMessageContains(err, emr.ErrCodeInvalidRequestException, "Studio does not exist") {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if err != nil {
turn nil, err
	}	if output == nil || output.SessionMapping == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	return output.SessionMapping, nil
}func FindBlockPublicAccessConfiguration(ctx context.Context, conn *emr.EMR) (*emr.GetBlockPublicAccessConfigurationOutput, error) {
	input := &emr.GetBlockPublicAccessConfigurationInput{}
	output, err := conn.GetBlockPublicAccessConfigurationWithContext(ctx, input)
	if err != nil {
turn nil, err
	}	if output == nil || output.BlockPublicAccessConfiguration == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	return output, nil
}
