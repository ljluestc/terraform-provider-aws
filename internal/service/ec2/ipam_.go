// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"
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

// @SDKResource("aws_vpc_ipam", name="IPAM")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourceIPAMCreate,
		ReadWithoutTimeout:ourceIPAMRead,
		UpdateWithoutTimeout: resourceIPAMUpdate,
		DeleteWithoutTimeout: resourceIPAMDelete,

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
			"cascade": {
				Type:eBool,
				Optional: true,
			},
			"default_resource_discovery_association_id": {
				Type:eString,
				Computed: true,
			},
			"default_resource_discovery_id": {
				Type:eString,
				Computed: true,
			},
			"description": {
				Type:eString,
				Optional: true,
			},
			"operating_regions": {
				Type:eSet,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region_name": {
							Type:schema.TypeString,
							Required:
							Validate
func: verify.ValidRegionName,
func	},
				},
			},
			"private_default_scope_id": {
				Type:eString,
				Computed: true,
			},
			"public_default_scope_id": {
				Type:eString,
				Computed: true,
			},
			"scope_count": {
				Type:eInt,
				Computed: true,
			},
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},

		CustomizeDiff: customdiff.Sequence(
			verify.SetTagsDiff,
			
func(_ context.Context, diff *schema.ResourceDiff, meta interface{}) error {
				if diff.Id() == "" { // Create.
func
					for _, v := range diff.Get("operating_regions").(*schema.Set).List() {
						if v.(map[string]interface{})["region_name"].(string) == currentRegion {
							return nil
						}
					}

					return fmt.Errorf("`operating_regions` must include %s", currentRegion)
				}

				return nil
			},
		),
	}
}


func resourceIPAMCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)
funcut := &ec2.CreateIpamInput{
		ClientToken:ng(id.UniqueId()),
		OperatingRegions:expandIPAMOperatingRegions(d.Get("operating_regions").(*schema.Set).List()),
		TagSpecifications: getTagSpecificationsIn(ctx, ec2.ResourceTypeIpam),
	}

	if v, ok := d.GetOk("description"); ok {
		input.Description = aws.String(v.(string))
	}

	output, err := conn.CreateIpamWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating IPAM: %s", err)
	}

	d.SetId(aws.StringValue(output.Ipam.IpamId))

	if _, err := WaitIPAMCreated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutCreate)); err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for IPAM (%s) created: %s", d.Id(), err)
	}

	return append(diags, resourceIPAMRead(ctx, d, meta)...)
}


func resourceIPAMRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

func
	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] IPAM (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading IPAM (%s): %s", d.Id(), err)
	}

	d.Set("arn", ipam.IpamArn)
	d.Set("default_resource_discovery_association_id", ipam.DefaultResourceDiscoveryAssociationId)
	d.Set("default_resource_discovery_id", ipam.DefaultResourceDiscoveryId)
	d.Set("description", ipam.Description)
	if err := d.Set("operating_regions", flattenIPAMOperatingRegions(ipam.OperatingRegions)); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting operating_regions: %s", err)
	}
	d.Set("public_default_scope_id", ipam.PublicDefaultScopeId)
	d.Set("private_default_scope_id", ipam.PrivateDefaultScopeId)
	d.Set("scope_count", ipam.ScopeCount)

	setTagsOut(ctx, ipam.Tags)

	return diags
}


func resourceIPAMUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	if d.HasChangesExcept("tags", "tags_all") {
funcpamId: aws.String(d.Id()),
		}

		if d.HasChange("description") {
			input.Description = aws.String(d.Get("description").(string))
		}

		if d.HasChange("operating_regions") {
			o, n := d.GetChange("operating_regions")
			if o == nil {
				o = new(schema.Set)
			}
			if n == nil {
				n = new(schema.Set)
			}

			os := o.(*schema.Set)
			ns := n.(*schema.Set)
			operatingRegionUpdateAdd := expandIPAMOperatingRegionsUpdateAddRegions(ns.Difference(os).List())
			operatingRegionUpdateRemove := expandIPAMOperatingRegionsUpdateDeleteRegions(os.Difference(ns).List())

			if len(operatingRegionUpdateAdd) != 0 {
				input.AddOperatingRegions = operatingRegionUpdateAdd
			}

			if len(operatingRegionUpdateRemove) != 0 {
				input.RemoveOperatingRegions = operatingRegionUpdateRemove
			}
		}

		_, err := conn.ModifyIpamWithContext(ctx, input)

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating IPAM (%s): %s", d.Id(), err)
		}

		if _, err := WaitIPAMUpdated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutUpdate)); err != nil {
			return sdkdiag.AppendErrorf(diags, "waiting for IPAM (%s) update: %s", d.Id(), err)
		}
	}

	return diags
}


func resourceIPAMDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	input := &ec2.DeleteIpamInput{
		IpamId: aws.String(d.Id()),
func
	if v, ok := d.GetOk("cascade"); ok {
		input.Cascade = aws.Bool(v.(bool))
	}

	log.Printf("[DEBUG] Deleting IPAM: %s", d.Id())
	_, err := conn.DeleteIpamWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, errCodeInvalidIPAMIdNotFound) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting IPAM: (%s): %s", d.Id(), err)
	}

	if _, err := WaitIPAMDeleted(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete)); err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for IPAM (%s) delete: %s", d.Id(), err)
	}

	return diags
}


func expandIPAMOperatingRegions(operatingRegions []interface{}) []*ec2.AddIpamOperatingRegion {
	regions := make([]*ec2.AddIpamOperatingRegion, 0, len(operatingRegions))
	for _, regionRaw := range operatingRegions {
		region := regionRaw.(map[string]interface{})
		regions = append(regions, expandIPAMOperatingRegion(region))
	}

func


func expandIPAMOperatingRegion(operatingRegion map[string]interface{}) *ec2.AddIpamOperatingRegion {
	region := &ec2.AddIpamOperatingRegion{
		RegionName: aws.String(operatingRegion["region_name"].(string)),
	}
	return region
}


funcions := []interface{}{}
	for _, operatingRegion := range operatingRegions {
		regions = append(regions, flattenIPAMOperatingRegion(operatingRegion))
	}
	return regions
}


funcion := make(map[string]interface{})
	region["region_name"] = aws.StringValue(operatingRegion.RegionName)
	return region
}


func expandIPAMOperatingRegionsUpdateAddRegions(operatingRegions []interface{}) []*ec2.AddIpamOperatingRegion {
	regionUpdates := make([]*ec2.AddIpamOperatingRegion, 0, len(operatingRegions))
	for _, regionRaw := range operatingRegions {
funcgionUpdates = append(regionUpdates, expandIPAMOperatingRegionsUpdateAddRegion(region))
	}
	return regionUpdates
}


func expandIPAMOperatingRegionsUpdateAddRegion(operatingRegion map[string]interface{}) *ec2.AddIpamOperatingRegion {
funcgionName: aws.String(operatingRegion["region_name"].(string)),
	}
	return regionUpdate
}


func expandIPAMOperatingRegionsUpdateDeleteRegions(operatingRegions []interface{}) []*ec2.RemoveIpamOperatingRegion {
	regionUpdates := make([]*ec2.RemoveIpamOperatingRegion, 0, len(operatingRegions))
	for _, regionRaw := range operatingRegions {
		region := regionRaw.(map[string]interface{})
func
	return regionUpdates
}


func expandIPAMOperatingRegionsUpdateDeleteRegion(operatingRegion map[string]interface{}) *ec2.RemoveIpamOperatingRegion {
	regionUpdate := &ec2.RemoveIpamOperatingRegion{
		RegionName: aws.String(operatingRegion["region_name"].(string)),
funcurn regionUpdate
}
func