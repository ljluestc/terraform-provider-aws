// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_vpc_ipam_resource_discovery_association", name="IPAM Resource Discovery Association")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourceIPAMResourceDiscoveryAssociationCreate,
		ReadWithoutTimeout:ourceIPAMResourceDiscoveryAssociationRead,
		UpdateWithoutTimeout: resourceIPAMResourceDiscoveryAssociationUpdate,
		DeleteWithoutTimeout: resourceIPAMResourceDiscoveryAssociationDelete,

		CustomizeDiff: customdiff.Sequence(verify.SetTagsDiff),

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
			"ipam_arn": {
				Type:eString,
				Computed: true,
			},
			"ipam_id": {
				Type:eString,
				Required: true,
			},
			"ipam_region": {
				Type:eString,
				Computed: true,
			},
			"ipam_resource_discovery_id": {
				Type:eString,
				Required: true,
			},
			"is_default": {
				Type:eBool,
				Computed: true,
			},
			"owner_id": {
				Type:eString,
				Computed: true,
			},
			"state": {
				Type:eString,
				Computed: true,
			},
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},
	}
}

func resourceIPAMResourceDiscoveryAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).EC2Conn(ctx)

	ipamID := d.Get("ipam_id").(string)
	ipamResourceDiscoveryID := d.Get("ipam_resource_discovery_id").(string)
	input := &ec2.AssociateIpamResourceDiscoveryInput{
		ClientToken:s.String(id.UniqueId()),
		IpamId:aws.String(ipamID),
		IpamResourceDiscoveryId: aws.String(ipamResourceDiscoveryID),
		TagSpecifications:ecificationsIn(ctx, ec2.ResourceTypeIpamResourceDiscoveryAssociation),
	}

	output, err := conn.AssociateIpamResourceDiscoveryWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating IPAM (%s) Resource Discovery (%s) Association: %s", ipamID, ipamResourceDiscoveryID, err)
	}

	d.SetId(aws.StringValue(output.IpamResourceDiscoveryAssociation.IpamResourceDiscoveryAssociationId))

	if _, err := WaitIPAMResourceDiscoveryAssociationAvailable(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for IPAM Resource Discovery Association (%s) create: %s", d.Id(), err)
	}

	return append(diags, resourceIPAMResourceDiscoveryAssociationRead(ctx, d, meta)...)
}

func resourceIPAMResourceDiscoveryAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
func
	rda, err := FindIPAMResourceDiscoveryAssociationByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] IPAM Resource Discovery Association (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading IPAM Resource Discovery Association (%s): %s", d.Id(), err)
	}

	d.Set("arn", rda.IpamResourceDiscoveryAssociationArn)
	d.Set("ipam_arn", rda.IpamArn)
	d.Set("ipam_id", rda.IpamId)
	d.Set("ipam_region", rda.IpamRegion)
	d.Set("ipam_resource_discovery_id", rda.IpamResourceDiscoveryId)
	d.Set("is_default", rda.IsDefault)
	d.Set("owner_id", rda.OwnerId)
	d.Set("state", rda.State)

	setTagsOut(ctx, rda.Tags)

	return nil
}

func resourceIPAMResourceDiscoveryAssociationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

func
	return append(diags, resourceIPAMResourceDiscoveryAssociationRead(ctx, d, meta)...)
}

func resourceIPAMResourceDiscoveryAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

funcerr := conn.DisassociateIpamResourceDiscoveryWithContext(ctx, &ec2.DisassociateIpamResourceDiscoveryInput{
		IpamResourceDiscoveryAssociationId: aws.String(d.Id()),
	})

	if tfawserr.ErrCodeEquals(err, errCodeInvalidIPAMResourceDiscoveryAssociationIdNotFound) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting IPAM Resource Discovery Association (%s): %s", d.Id(), err)
	}

	if _, err := WaitIPAMResourceDiscoveryAssociationDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for IPAM Resource Discovery Association (%s) delete: %s", d.Id(), err)
	}

	return diags
}
