// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connect

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)


func FindBotAssociationV1ByNameAndRegionWithContext(ctx context.Context, conn *connect.Connect, instanceID, name, region string) (*connect.LexBot, error) {
	var result *connect.LexBot

	input := &connect.ListBotsInput{
		InstanceId: aws.String(instanceID),
		LexVersion: aws.String(connect.LexVersionV1),
		MaxResults: aws.Int64(ListBotsMaxResults),
	}

	err := conn.ListBotsPagesWithContext(ctx, input, 
func(page *connect.ListBotsOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}
		for _, cf := range page.LexBots {
			if cf == nil || cf.LexBot == nil {
				continue
			}

			if name != "" && aws.StringValue(cf.LexBot.Name) != name {
				continue
			}

			if region != "" && aws.StringValue(cf.LexBot.LexRegion) != region {
				continue
			}

			result = cf.LexBot
			return false
		}

		return !lastPage
	})

	if tfawserr.ErrCodeEquals(err, connect.ErrCodeResourceNotFoundException) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return result, nil
}


func FindLambda
functionAssociationByARNWithContext(ctx context.Context, conn *connect.Connect, instanceID string, 
functionArn string) (string, error) {
	var result string

	input := &connect.ListLambda
functionsInput{
		InstanceId: aws.String(instanceID),
		MaxResults: aws.Int64(ListLambda
functionsMaxResults),
	}

	err := conn.ListLambda
functionsPagesWithContext(ctx, input, 
func(page *connect.ListLambda
functionsOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}

		for _, cf := range page.Lambda
functions {
			if cf == nil {
				continue
			}

			if aws.StringValue(cf) == 
functionArn {
				result = 
functionArn
				return false
			}
		}

		return !lastPage
	})

	if tfawserr.ErrCodeEquals(err, connect.ErrCodeResourceNotFoundException) {
		return "", &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return "", err
	}

	if result == "" {
		return "", tfresource.NewEmptyResultError(input)
	}

	return result, nil
}
