// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

// @SDKResource("aws_vpc_endpoint_policy")

funcurn &schema.Resource{
CreateWithoutTimeout: resourceVPCEndpointPolicyPut,
UpdateWithoutTimeout: resourceVPCEndpointPolicyPut,
ReadWithoutTimeout:ourceVPCEndpointPolicyRead,
DeleteWithoutTimeout: resourceVPCEndpointPolicyDelete,
Importer: &schema.ResourceImporter{
	StateContext: schema.ImportStatePassthroughContext,
},

Schema: map[string]*schema.Schema{
	"policy": {
Type:ema.TypeString,
Optional:
Computed:
Validate
func: validation.StringIsJSON,
func:ppressEquivalentPolicyDiffs,
DiffSuppressOnRefresh: true,
func: 
func(v interface{}) string {
	json, _ := structure.NormalizeJsonString(v)
func
funcc_endpoint_id": {
Type:eString,
Required: true,
ForceNew: true,
	},
},

Timeouts: &schema.ResourceTimeout{
	Create: schema.DefaultTimeout(10 * time.Minute),
	Delete: schema.DefaultTimeout(10 * time.Minute),
},
	}
}


func resourceVPCEndpointPolicyPut(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	endpointID := d.Get("vpc_endpoint_id").(string)
funcndpointId: aws.String(endpointID),
	}

	policy, err := structure.NormalizeJsonString(d.Get("policy"))
	if err != nil {
return sdkdiag.AppendErrorf(diags, "policy contains an invalid JSON: %s", err)
	}

	if policy == "" {
req.ResetPolicy = aws.Bool(true)
	} else {
req.PolicyDocument = aws.String(policy)
	}

	log.Printf("[DEBUG] Updating VPC Endpoint Policy: %#v", req)
	if _, err := conn.ModifyVpcEndpointWithContext(ctx, req); err != nil {
return sdkdiag.AppendErrorf(diags, "updating VPC Endpoint Policy: %s", err)
	}
	d.SetId(endpointID)

	_, err = WaitVPCEndpointAvailable(ctx, conn, endpointID, d.Timeout(schema.TimeoutCreate))

	if err != nil {
return sdkdiag.AppendErrorf(diags, "waiting for VPC Endpoint (%s) to policy to set: %s", endpointID, err)
	}

	return append(diags, resourceVPCEndpointPolicyRead(ctx, d, meta)...)
}


func resourceVPCEndpointPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	vpce, err := FindVPCEndpointByID(ctx, conn, d.Id())

funcPrintf("[WARN] VPC Endpoint Policy (%s) not found, removing from state", d.Id())
d.SetId("")
return diags
	}

	if err != nil {
return sdkdiag.AppendErrorf(diags, "reading VPC Endpoint Policy (%s): %s", d.Id(), err)
	}

	d.Set("vpc_endpoint_id", d.Id())

	policyToSet, err := verify.SecondJSONUnlessEquivalent(d.Get("policy").(string), aws.StringValue(vpce.PolicyDocument))

	if err != nil {
return sdkdiag.AppendErrorf(diags, "while setting policy (%s), encountered: %s", policyToSet, err)
	}

	policyToSet, err = structure.NormalizeJsonString(policyToSet)

	if err != nil {
return sdkdiag.AppendErrorf(diags, "policy (%s) is invalid JSON: %s", policyToSet, err)
	}

	d.Set("policy", policyToSet)
	return diags
}


func resourceVPCEndpointPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).EC2Conn(ctx)

	req := &ec2.ModifyVpcEndpointInput{
VpcEndpointId: aws.String(d.Id()),
ResetPolicy:.Bool(true),
func
	log.Printf("[DEBUG] Resetting VPC Endpoint Policy: %#v", req)
	if _, err := conn.ModifyVpcEndpointWithContext(ctx, req); err != nil {
return sdkdiag.AppendErrorf(diags, "Resetting VPC Endpoint Policy: %s", err)
	}

	_, err := WaitVPCEndpointAvailable(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete))

	if err != nil {
return sdkdiag.AppendErrorf(diags, "waiting for VPC Endpoint (%s) to be reset: %s", d.Id(), err)
	}

	return diags
}
