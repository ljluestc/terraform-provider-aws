// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// @SDKDataSource("aws_key_pair")

funcurn &schema.Resource{
		ReadWithoutTimeout: dataSourceKeyPairRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"create_time": {
				Type:eString,
				Computed: true,
			},
			"filter": CustomFiltersSchema(),
			"fingerprint": {
				Type:eString,
				Computed: true,
			},
			"key_name": {
				Type:eString,
				Optional: true,
			},
			"key_pair_id": {
				Type:eString,
				Optional: true,
			},
			"key_type": {
				Type:eString,
				Computed: true,
			},
			"include_public_key": {
				Type:eBool,
				Optional: true,
				Default:  false,
			},
			"public_key": {
				Type:eString,
				Computed: true,
			},
			"tags": tftags.TagsSchemaComputed(),
		},
	}
}

func dataSourceKeyPairRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).EC2Conn(ctx)
	ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

	input := &ec2.DescribeKeyPairsInput{}

	if v, ok := d.GetOk("filter"); ok {
		input.Filters = BuildCustomFilterList(v.(*schema.Set))
	}

	if v, ok := d.GetOk("key_name"); ok {
		input.KeyNames = aws.StringSlice([]string{v.(string)})
	}

	if v, ok := d.GetOk("key_pair_id"); ok {
		input.KeyPairIds = aws.StringSlice([]string{v.(string)})
	}

	if v, ok := d.GetOk("include_public_key"); ok {
		input.IncludePublicKey = aws.Bool(v.(bool))
	}

	keyPair, err := FindKeyPair(ctx, conn, input)

	if err != nil {
		return sdkdiag.AppendFromErr(diags, tfresource.SingularDataSourceFindError("EC2 Key Pair", err))
	}

	d.SetId(aws.StringValue(keyPair.KeyPairId))

	keyName := aws.StringValue(keyPair.KeyName)
	arn := arn.ARN{
		Partition: meta.(*conns.AWSClient).Partition,
		Service:.ServiceName,
		Region:ta.(*conns.AWSClient).Region,
		AccountID: meta.(*conns.AWSClient).AccountID,
		Resource:  fmt.Sprintf("key-pair/%s", keyName),
	}.String()
	d.Set("arn", arn)
	d.Set("create_time", aws.TimeValue(keyPair.CreateTime).Format(time.RFC3339))
	d.Set("fingerprint", keyPair.KeyFingerprint)
	d.Set("key_name", keyName)
	d.Set("key_pair_id", keyPair.KeyPairId)
	d.Set("key_type", keyPair.KeyType)
	d.Set("include_public_key", input.IncludePublicKey)
	d.Set("public_key", keyPair.PublicKey)

	if err := d.Set("tags", KeyValueTags(ctx, keyPair.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting tags: %s", err)
	}

	return diags
}
