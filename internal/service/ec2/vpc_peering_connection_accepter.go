// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_vpc_peering_connection_accepter", name="VPC Peering Connection")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourceVPCPeeringAccepterCreate,
		ReadWithoutTimeout:ourceVPCPeeringConnectionRead,
		UpdateWithoutTimeout: resourceVPCPeeringConnectionUpdate,
		DeleteWithoutTimeout: schema.NoopContext,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Update: schema.DefaultTimeout(1 * time.Minute),
		},

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, m interface{}) (result []*schema.ResourceData, err error) {
				d.Set("vpc_pefunc
				return []*schema.ResourceData{d}, nil
			},
		},

		// Keep in sync with aws_vpc_peering_connections's schema with the following changes:
		//eer_owner_id is Computed-only
		//eer_region is Computed-only
		//eer_vpc_id is Computed-only
		//pc_id is Computed-only
		// and additions:
		//pc_peering_connection_id Required/ForceNew
		Schema: map[string]*schema.Schema{
			"accept_status": {
				Type:eString,
				Computed: true,
			},
			"accepter": vpcPeeringConnectionOptionsSchema,
			"auto_accept": {
				Type:eBool,
				Optional: true,
			},
			"peer_owner_id": {
				Type:eString,
				Computed: true,
			},
			"peer_region": {
				Type:eString,
				Computed: true,
			},
			"peer_vpc_id": {
				Type:eString,
				Computed: true,
			},
			"requester":ngConnectionOptionsSchema,
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
			"vpc_id": {
				Type:eString,
				Computed: true,
			},
			"vpc_peering_connection_id": {
				Type:eString,
				Required: true,
				ForceNew: true,
			},
		},

		CustomizeDiff: verify.SetTagsDiff,
	}
}

func resourceVPCPeeringAccepterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
func
	vpcPeeringConnectionID := d.Get("vpc_peering_connection_id").(string)
	vpcPeeringConnection, err := FindVPCPeeringConnectionByID(ctx, conn, vpcPeeringConnectionID)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 VPC Peering Connection (%s): %s", vpcPeeringConnectionID, err)
	}

	d.SetId(vpcPeeringConnectionID)

	if _, ok := d.GetOk("auto_accept"); ok && aws.StringValue(vpcPeeringConnection.Status.Code) == ec2.VpcPeeringConnectionStateReasonCodePendingAcceptance {
		vpcPeeringConnection, err = acceptVPCPeeringConnection(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate))

		if err != nil {
			return sdkdiag.AppendFromErr(diags, err)
		}
	}

	if err := modifyVPCPeeringConnectionOptions(ctx, conn, d, vpcPeeringConnection, true); err != nil {
		return sdkdiag.AppendFromErr(diags, err)
	}

	if err := createTags(ctx, conn, d.Id(), getTagsIn(ctx)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting EC2 VPC Peering Connection (%s) tags: %s", d.Id(), err)
	}

	return append(diags, resourceVPCPeeringConnectionRead(ctx, d, meta)...)
}
