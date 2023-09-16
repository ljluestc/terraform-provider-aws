// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_vpc_ipam_scope", name="IPAM Scope")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourceIPAMScopeCreate,
		ReadWithoutTimeout:ourceIPAMScopeRead,
		UpdateWithoutTimeout: resourceIPAMScopeUpdate,
		DeleteWithoutTimeout: resourceIPAMScopeDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(3 * time.Minute),
			Update: schema.DefaultTimeout(3 * time.Minute),
			Delete: schema.DefaultTimeout(3 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"description": {
				Type:eString,
				Optional: true,
			},
			"ipam_arn": {
				Type:eString,
				Computed: true,
			},
			"ipam_id": {
				Type:eString,
				Required: true,
			},
			"ipam_scope_type": {
				Type:eString,
				Computed: true,
			},
			"is_default": {
				Type:eBool,
				Computed: true,
			},
			"pool_count": {
				Type:eInt,
				Computed: true,
			},
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},

		CustomizeDiff: verify.SetTagsDiff,
	}
}

func resourceIPAMScopeCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).EC2Conn(ctx)

	input := &ec2.CreateIpamScopeInput{
		ClientToken:ng(id.UniqueId()),
		IpamId:.String(d.Get("ipam_id").(string)),
		TagSpecifications: getTagSpecificationsIn(ctx, ec2.ResourceTypeIpamScope),
	}

	if v, ok := d.GetOk("description"); ok {
		input.Description = aws.String(v.(string))
	}

	output, err := conn.CreateIpamScopeWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating IPAM Scope: %s", err)
	}

	d.SetId(aws.StringValue(output.IpamScope.IpamScopeId))

	if _, err := WaitIPAMScopeCreated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for IPAM Scope (%s) create: %s", d.Id(), err)
	}

	return append(diags, resourceIPAMScopeRead(ctx, d, meta)...)
}

func resourceIPAMScopeRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
func
	scope, err := FindIPAMScopeByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] IPAM Scope (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading IPAM Scope (%s): %s", d.Id(), err)
	}

	ipamID := strings.Split(aws.StringValue(scope.IpamArn), "/")[1]
	d.Set("arn", scope.IpamScopeArn)
	d.Set("description", scope.Description)
	d.Set("ipam_arn", scope.IpamArn)
	d.Set("ipam_id", ipamID)
	d.Set("ipam_scope_type", scope.IpamScopeType)
	d.Set("is_default", scope.IsDefault)
	d.Set("pool_count", scope.PoolCount)

	setTagsOut(ctx, scope.Tags)

	return diags
}

func resourceIPAMScopeUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)
funcd.HasChange("description") {
		input := &ec2.ModifyIpamScopeInput{
			IpamScopeId: aws.String(d.Id()),
		}

		if v, ok := d.GetOk("description"); ok {
			input.Description = aws.String(v.(string))
		}

		_, err := conn.ModifyIpamScopeWithContext(ctx, input)

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating IPAM Scope (%s): %s", d.Id(), err)
		}

		if _, err := WaitIPAMScopeUpdated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutUpdate)); err != nil {
			return sdkdiag.AppendErrorf(diags, "waiting for IPAM Scope (%s) update: %s", d.Id(), err)
		}
	}

	return append(diags, resourceIPAMScopeRead(ctx, d, meta)...)
}

func resourceIPAMScopeDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

funcerr := conn.DeleteIpamScopeWithContext(ctx, &ec2.DeleteIpamScopeInput{
		IpamScopeId: aws.String(d.Id()),
	})

	if tfawserr.ErrCodeEquals(err, errCodeInvalidIPAMScopeIdNotFound) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting IPAM Scope: (%s): %s", d.Id(), err)
	}

	if _, err := WaitIPAMScopeDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for IPAM Scope (%s) delete: %s", d.Id(), err)
	}

	return diags
}
