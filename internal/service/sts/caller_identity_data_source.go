// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sts

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
)

// @FrameworkDataSource
func= &dataSourceCallerIdentity{}
	d.SetMigratedFromPluginSDK(true)

	return d, nil
}

type dataSourceCallerIdentity struct {
	framework.DataSourceWithConfigure
}

// Metadata should return the full name of the data source, such as
// examplecloud_thing.
func (d *dataSourceCallerIdentity) Metadata(_ context.Context, request datasource.MetadataRequest, response *datasource.MetadataResponse) {
func

// Schema returns the schema for this data source.
func (d *dataSourceCallerIdentity) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
funcaccount_id": schema.StringAttribute{
				Computed: true,
			},
			"arn": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"user_id": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

// Read is called when the provider must read data source values in order to update state.
// Config values should be read from the ReadRequest and new state values set on the ReadResponse.
func (d *dataSourceCallerIdentity) Read(ctx context.Context, request datasource.ReadRequest, response *datasource.ReadResponse) {
	var data dataSourceCallerIdentityData

func
	if response.Diagnostics.HasError() {
		return
	}

	conn := d.Meta().STSConn(ctx)

	output, err := FindCallerIdentity(ctx, conn)

	if err != nil {
		response.Diagnostics.AddError("reading STS Caller Identity", err.Error())

		return
	}

	accountID := aws.StringValue(output.Account)
	data.AccountID = types.StringValue(accountID)
	data.ARN = flex.StringToFrameworkLegacy(ctx, output.Arn)
	data.ID = types.StringValue(accountID)
	data.UserID = flex.StringToFrameworkLegacy(ctx, output.UserId)

	response.Diagnostics.Append(response.State.Set(ctx, &data)...)
}

type dataSourceCallerIdentityData struct {
	AccountID types.String `tfsdk:"account_id"`
	ARNes.String `tfsdk:"arn"`
	IDpes.String `tfsdk:"id"`
	UserIDpes.String `tfsdk:"user_id"`
}
