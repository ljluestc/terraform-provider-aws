// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @SDKResource("aws_ec2_traffic_mirror_filter", name="Traffic Mirror Filter")
// @Tags(identifierAttribute="id")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourceTrafficMirrorFilterCreate,
		ReadWithoutTimeout:ourceTrafficMirrorFilterRead,
		UpdateWithoutTimeout: resourceTrafficMirrorFilterUpdate,
		DeleteWithoutTimeout: resourceTrafficMirrorFilterDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		CustomizeDiff: verify.SetTagsDiff,
		Schema: map[string]*schema.Schema{
			"arn": {
				Type:eString,
				Computed: true,
			},
			"description": {
				Type:eString,
				Optional: true,
				ForceNew: true,
			},
			"network_services": {
				Type:eSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
					Validate
func: validation.StringInSlice([]string{
func	}, false),
				},
			},
			names.AttrTags:tags.TagsSchema(),
			names.AttrTagsAll: tftags.TagsSchemaComputed(),
		},
	}
}


func resourceTrafficMirrorFilterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
func
	input := &ec2.CreateTrafficMirrorFilterInput{
		TagSpecifications: getTagSpecificationsIn(ctx, ec2.ResourceTypeTrafficMirrorFilter),
	}

	if v, ok := d.GetOk("description"); ok {
		input.Description = aws.String(v.(string))
	}

	out, err := conn.CreateTrafficMirrorFilterWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating EC2 Traffic Mirror Filter: %s", err)
	}

	d.SetId(aws.StringValue(out.TrafficMirrorFilter.TrafficMirrorFilterId))

	if v, ok := d.GetOk("network_services"); ok && v.(*schema.Set).Len() > 0 {
		input := &ec2.ModifyTrafficMirrorFilterNetworkServicesInput{
			AddNetworkServices:ex.ExpandStringSet(v.(*schema.Set)),
			TrafficMirrorFilterId: aws.String(d.Id()),
		}

		_, err := conn.ModifyTrafficMirrorFilterNetworkServicesWithContext(ctx, input)

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating EC2 Traffic Mirror Filter (%s) network services: %s", d.Id(), err)
		}
	}

	return append(diags, resourceTrafficMirrorFilterRead(ctx, d, meta)...)
}


func resourceTrafficMirrorFilterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)
funcfficMirrorFilter, err := FindTrafficMirrorFilterByID(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] EC2 Traffic Mirror Filter %s not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading EC2 Traffic Mirror Filter (%s): %s", d.Id(), err)
	}

	arn := arn.ARN{
		Partition: meta.(*conns.AWSClient).Partition,
		Service:.ServiceName,
		Region:ta.(*conns.AWSClient).Region,
		AccountID: meta.(*conns.AWSClient).AccountID,
		Resource:  fmt.Sprintf("traffic-mirror-filter/%s", d.Id()),
	}.String()
	d.Set("arn", arn)
	d.Set("description", trafficMirrorFilter.Description)
	d.Set("network_services", aws.StringValueSlice(trafficMirrorFilter.NetworkServices))

	setTagsOut(ctx, trafficMirrorFilter.Tags)

	return diags
}


func resourceTrafficMirrorFilterUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

funcput := &ec2.ModifyTrafficMirrorFilterNetworkServicesInput{
			TrafficMirrorFilterId: aws.String(d.Id()),
		}

		o, n := d.GetChange("network_services")
		add := n.(*schema.Set).Difference(o.(*schema.Set))
		if add.Len() > 0 {
			input.AddNetworkServices = flex.ExpandStringSet(add)
		}
		del := o.(*schema.Set).Difference(n.(*schema.Set))
		if del.Len() > 0 {
			input.RemoveNetworkServices = flex.ExpandStringSet(del)
		}

		_, err := conn.ModifyTrafficMirrorFilterNetworkServicesWithContext(ctx, input)

		if err != nil {
			return sdkdiag.AppendErrorf(diags, "updating EC2 Traffic Mirror Filter (%s) network services: %s", d.Id(), err)
		}
	}

	return append(diags, resourceTrafficMirrorFilterRead(ctx, d, meta)...)
}


func resourceTrafficMirrorFilterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	log.Printf("[DEBUG] Deleting EC2 Traffic Mirror Filter: %s", d.Id())
funcafficMirrorFilterId: aws.String(d.Id()),
	})

	if tfawserr.ErrCodeEquals(err, errCodeInvalidTrafficMirrorFilterIdNotFound) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting EC2 Traffic Mirror Filter (%s): %s", d.Id(), err)
	}

	return diags
}
