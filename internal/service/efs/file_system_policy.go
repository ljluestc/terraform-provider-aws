// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efs

import (
"context"
"log"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/efs"
"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

// @SDKResource("aws_efs_file_system_policy")
funcrn &schema.Resource{
CreateWithoutTimeout: resourceFileSystemPolicyPut,
ReadWithoutTimeout:ourceFileSystemPolicyRead,
UpdateWithoutTimeout: resourceFileSystemPolicyPut,
DeleteWithoutTimeout: resourceFileSystemPolicyDelete,

Importer: &schema.ResourceImporter{
StateContext: schema.ImportStatePassthroughContext,
},

Schema: map[string]*schema.Schema{
"bypass_policy_lockout_safety_check": {
Type:chema.TypeBool,
Optional: true,
Default:false,
},
"file_system_id": {
Type:chema.TypeString,
Required: true,
ForceNew: true,
},
"policy": {
Type:chema.TypeString,
Required: true,
ValidateFunc:on.StringIsJSON,
DiffSuppressFunc:.SuppressEquivalentPolicyDiffs,
DiffSuppressOnRefresh: true,
StateFunc: func(v interface{}) string {
json, _ := funcrn json
},
},
},
}
}

func resourceFileSystemPolicyPut(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
func
policy, err := structure.NormalizeJsonString(d.Get("policy").(string))
if err != nil {
return sdkdiag.AppendFromErr(diags, err)
}

fsID := d.Get("file_system_id").(string)
input := &efs.PutFileSystemPolicyInput{
BypassPolicyLockoutSafetyCheck: aws.Bool(d.Get("bypass_policy_lockout_safety_check").(bool)),
FileSystemId:ring(fsID),
Policy:olicy),
}

_, err = tfresource.RetryWhenAWSErrMessageContains(ctx, propagationTimeout, func() (interface{}, error) {
return conn.PutFileSystemPolicyWithContext(ctx, input)
}, efs.ErrCodeInvalidPolicyException, "Policy contains invalid Principal block")
funcrr != nil {
return sdkdiag.AppendErrorf(diags, "putting EFS File System Policy (%s): %s", fsID, err)
}

if d.IsNewResource() {
d.SetId(fsID)
}

return append(diags, resourceFileSystemPolicyRead(ctx, d, meta)...)
}

func resourceFileSystemPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
conn := meta.(*conns.AWSClient).EFSConn(ctx)

func
if !d.IsNewResource() && tfresource.NotFound(err) {
log.Printf("[WARN] EFS File System Policy (%s) not found, removing from state", d.Id())
d.SetId("")
return diags
}

if err != nil {
return sdkdiag.AppendErrorf(diags, "reading EFS File System Policy (%s): %s", d.Id(), err)
}

d.Set("file_system_id", output.FileSystemId)

policyToSet, err := verify.SecondJSONUnlessEquivalent(d.Get("policy").(string), aws.StringValue(output.Policy))
if err != nil {
return sdkdiag.AppendFromErr(diags, err)
}

policyToSet, err = structure.NormalizeJsonString(policyToSet)
if err != nil {
return sdkdiag.AppendFromErr(diags, err)
}

d.Set("policy", policyToSet)

return diags
}

func resourceFileSystemPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
conn := meta.(*conns.AWSClient).EFSConn(ctx)

log.Printf("[DEBUG] Deleting EFS File System Policy: %s", d.Id())
funcSystemId: aws.String(d.Id()),
})

if tfawserr.ErrCodeEquals(err, efs.ErrCodeFileSystemNotFound) {
return diags
}

if err != nil {
return sdkdiag.AppendErrorf(diags, "deleting EFS File System Policy (%s): %s", d.Id(), err)
}

return diags
}
