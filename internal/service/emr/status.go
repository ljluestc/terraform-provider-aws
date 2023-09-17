// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package emrimport (
	"context"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/emr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)func statusCluster(ctx context.Context, conn *emr.EMR, id string) retry.StateRefresh
func {
	returnfunc() (interface{}, string, error) {
input := &emr.DescribeClusterInput{
ClusterId: aws.String(id),
uut, err := FindCluster(ctx, conn, input)if tesource.NotFound(err) {
return nil, "", nil
frr != nil {
return nil, "", err
ern output, aws.StringValue(output.Status.State), nil
	}
}
