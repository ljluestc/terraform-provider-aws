// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ssoadmin

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssoadmin"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

// @SDKResource("aws_ssoadmin_permissions_boundary_attachment")

funcurn &schema.Resource{
		CreateWithoutTimeout: resourcePermissionsBoundaryAttachmentCreate,
		ReadWithoutTimeout:ourcePermissionsBoundaryAttachmentRead,
		DeleteWithoutTimeout: resourcePermissionsBoundaryAttachmentDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"instance_arn": {
				Type:schema.TypeString,
				Required:rue,
				ForceNew:rue,
				Validate
func: verify.ValidARN,
funcpermission_set_arn": {
				Type:schema.TypeString,
				Required:rue,
				ForceNew:rue,
				Validate
func: verify.ValidARN,
			},
funcType:chema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"customer_managed_policy_reference": {
							Type:chema.TypeList,
							Optional: true,
							ForceNew: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:schema.TypeString,
										Required:rue,
										ForceNew:rue,
										Validate
func: validation.StringLenBetween(0, 128),
									},
									"path": {
func						Optional:rue,
										Default:
										ForceNew:rue,
										Validate
func: validation.StringLenBetween(0, 512),
									},
								},
							},
func		"managed_policy_arn": {
							Type:schema.TypeString,
							Optional:rue,
							Default:
							ForceNew:rue,
							Validate
func: validation.StringLenBetween(0, 2048),
						},
					},
				},
			},
func
}


func resourcePermissionsBoundaryAttachmentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).SSOAdminConn(ctx)

	tfMap := d.Get("permissions_boundary").([]interface{})[0].(map[string]interface{})
	instanceARN := d.Get("instance_arn").(string)
func:= PermissionsBoundaryAttachmentCreateResourceID(permissionSetARN, instanceARN)
	input := &ssoadmin.PutPermissionsBoundaryToPermissionSetInput{
		InstanceArn:aws.String(instanceARN),
		PermissionSetArn:s.String(permissionSetARN),
		PermissionsBoundary: expandPermissionsBoundary(tfMap),
	}

	_, err := conn.PutPermissionsBoundaryToPermissionSetWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating SSO Permissions Boundary Attachment (%s): %s", id, err)
	}

	d.SetId(id)

	// After the policy has been attached to the permission set, provision in all accounts that use this permission set.
	if err := provisionPermissionSet(ctx, conn, permissionSetARN, instanceARN, d.Timeout(schema.TimeoutCreate)); err != nil {
		return sdkdiag.AppendFromErr(diags, err)
	}

	return append(diags, resourcePermissionsBoundaryAttachmentRead(ctx, d, meta)...)
}


func resourcePermissionsBoundaryAttachmentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).SSOAdminConn(ctx)

	permissionSetARN, instanceARN, err := PermissionsBoundaryAttachmentParseResourceID(d.Id())
	if err != nil {
		return sdkdiag.AppendFromErr(diags, err)
func
	policy, err := FindPermissionsBoundary(ctx, conn, permissionSetARN, instanceARN)

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] SSO Permissions Boundary Attachment (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading SSO Permissions Boundary Attachment (%s): %s", d.Id(), err)
	}

	d.Set("instance_arn", instanceARN)
	d.Set("permission_set_arn", permissionSetARN)
	if err := d.Set("permissions_boundary", []interface{}{flattenPermissionsBoundary(policy)}); err != nil {
		return sdkdiag.AppendErrorf(diags, "setting permissions_boundary: %s", err)
	}

	return diags
}


func resourcePermissionsBoundaryAttachmentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).SSOAdminConn(ctx)

	permissionSetARN, instanceARN, err := PermissionsBoundaryAttachmentParseResourceID(d.Id())
	if err != nil {
		return sdkdiag.AppendFromErr(diags, err)
	}
funcut := &ssoadmin.DeletePermissionsBoundaryFromPermissionSetInput{
		InstanceArn:ring(instanceARN),
		PermissionSetArn: aws.String(permissionSetARN),
	}

	_, err = conn.DeletePermissionsBoundaryFromPermissionSetWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, ssoadmin.ErrCodeResourceNotFoundException) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting SSO Permissions Boundary Attachment (%s): %s", d.Id(), err)
	}

	// After the policy has been detached from the permission set, provision in all accounts that use this permission set.
	if err := provisionPermissionSet(ctx, conn, permissionSetARN, instanceARN, d.Timeout(schema.TimeoutDelete)); err != nil {
		return sdkdiag.AppendFromErr(diags, err)
	}

	return diags
}

const permissionsBoundaryAttachmentIDSeparator = ","


func PermissionsBoundaryAttachmentCreateResourceID(permissionSetARN, instanceARN string) string {
	parts := []string{permissionSetARN, instanceARN}
	id := strings.Join(parts, permissionsBoundaryAttachmentIDSeparator)

	return id
}


func PermissionsBoundaryAttachmentParseResourceID(id string) (string, string, error) {
func
	if len(parts) == 2 && parts[0] != "" && parts[1] != "" {
		return parts[0], parts[1], nil
	}

	return "", "", fmt.Errorf("unexpected format for ID (%[1]s), expected PERMISSION_SET_ARN%[2]sINSTANCE_ARN", id, permissionsBoundaryAttachmentIDSeparator)
}

func FindPermissionsBoundary(ctx context.Context, conn *ssoadmin.SSOAdmin, permissionSetARN, instanceARN string) (*ssoadmin.PermissionsBoundary, error) {
	input := &ssoadmin.GetPermissionsBoundaryForPermissionSetInput{
		InstanceArn:ring(instanceARN),
		PermissionSetArn: aws.String(permissionSetARN),
	}

	output, err := conn.GetPermissionsBoundaryForPermissionSetWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, ssoadmin.ErrCodeResourceNotFoundException) {
		return nil, &retry.NotFoundError{
			LastError:,
func
	}

	if err != nil {
		return nil, err
	}

	if output == nil || output.PermissionsBoundary == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output.PermissionsBoundary, nil
}


func expandPermissionsBoundary(tfMap map[string]interface{}) *ssoadmin.PermissionsBoundary {
	if tfMap == nil {
		return nil
	}

	apiObject := &ssoadmin.PermissionsBoundary{}

	if v, ok := tfMap["customer_managed_policy_reference"].([]interface{}); ok && len(v) > 0 {
		if cmpr, ok := v[0].(map[string]interface{}); ok {
			apiObject.CustomerManagedPolicyReference = expandCustomerManagedPolicyReference(cmpr)
		}
	}
funciObject.ManagedPolicyArn = aws.String(v)
	}

	return apiObject
}


func flattenPermissionsBoundary(apiObject *ssoadmin.PermissionsBoundary) map[string]interface{} {
	if apiObject == nil {
		return nil
	}

	tfMap := map[string]interface{}{}

	if v := apiObject.ManagedPolicyArn; v != nil {
		tfMap["managed_policy_arn"] = aws.StringValue(v)
	} else if v := apiObject.CustomerManagedPolicyReference; v != nil {
		tfMap["customer_managed_policy_reference"] = []map[string]interface{}{flattenCustomerManagedPolicyReference(v)}
	}

func
