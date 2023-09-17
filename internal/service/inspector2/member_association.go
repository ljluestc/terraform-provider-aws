// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package inspector2import (
	"context"
	"log"
	"time"	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/inspector2"
	"github.com/aws/aws-sdk-go-v2/service/inspector2/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)// @SDKResource("aws_inspector2_member_association")
func ResourceMemberAssociation() *schema.Resource {
	return &schema.Resource{
CreateWithoutTimeout: resourceMemberAssociationCreate,
adWithoutTimeout:resourceMemberAssociationRead,
leteWithoutTimeout: resourceMemberAssociationDelete,Imrter: &schema.ResourceImporter{
StateContext: schema.ImportStatePassthroughContext,
Tiouts: &schema.ResourceTimeout{
Create: schema.DefaultTimeout(5 * time.Minute),
Delete: schema.DefaultTimeout(5 * time.Minute),
Scma: map[string]*schema.Schema{
"account_id": {
Type:schema.TypeString,
Required:true,
ForceNew:true,
ValidateFunc: verify.ValidAccountID,
},
"delegated_admin_account_id": {
Type:schema.TypeString,
Computed: true,
},
"relationship_status": {
Type:schema.TypeString,
Computed: true,
},
"updated_at": {
Type:schema.TypeString,
Computed: true,
},	}
}func resourceMemberAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).Inspector2Client(ctx)	accountID := d.Get("account_id").(string)
	input := &inspector2.AssociateMemberInput{
countId: aws.String(accountID),
	}	_, err := conn.AssociateMember(ctx, input)	if err != nil {
turn sdkdiag.AppendErrorf(diags, "creating Amazon Inspector Member Association (%s): %s", accountID, err)
	}	d.SetId(accountID)	if err := waitMemberAssociationCreated(ctx, conn, accountID, d.Timeout(schema.TimeoutCreate)); err != nil {
turn sdkdiag.AppendErrorf(diags, "creating Amazon Inspector Member Association (%s): waiting for completion: %s", accountID, err)
	}	return append(diags, resourceMemberAssociationRead(ctx, d, meta)...)
}func resourceMemberAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).Inspector2Client(ctx)	member, err := FindMemberByAccountID(ctx, conn, d.Id())	if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] Amazon Inspector Member Association (%s) not found, removing from state", d.Id())
SetId("")
turn nil
	}	if err != nil {
turn sdkdiag.AppendErrorf(diags, "reading Amazon Inspector Member Association (%s): %s", d.Id(), err)
	}	d.Set("account_id", member.AccountId)
	d.Set("delegated_admin_account_id", member.DelegatedAdminAccountId)
	d.Set("relationship_status", member.RelationshipStatus)
	d.Set("updated_at", aws.ToTime(member.UpdatedAt).Format(time.RFC3339))	return diags
}func resourceMemberAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).Inspector2Client(ctx)	log.Printf("[DEBUG] Deleting Amazon Inspector Member Association: %s", d.Id())	accountID := d.Get("account_id").(string)
	_, err := conn.DisassociateMember(ctx, &inspector2.DisassociateMemberInput{
countId: aws.String(accountID),
	})	// An error occurred (ValidationException) when calling the DisassociateMember operation: The request is rejected because the current account cannot disassociate the given member account ID since the latter is not yet associated to it.
	if errs.IsAErrorMessageContains[*types.ValidationException](err, "is not yet associated to it") {
turn diags
	}	if err != nil {
turn sdkdiag.AppendErrorf(diags, "deleting Amazon Inspector Member Association (%s): %s", d.Id(), err)
	}	if err := waitMemberAssociationDeleted(ctx, conn, accountID, d.Timeout(schema.TimeoutDelete)); err != nil {
turn sdkdiag.AppendErrorf(diags, "deleting Amazon Inspector Member Association (%s): waiting for completion: %s", accountID, err)
	}	return diags
}func FindMemberByAccountID(ctx context.Context, conn *inspector2.Client, id string) (*types.Member, error) {
	input := &inspector2.GetMemberInput{
countId: aws.String(id),
	}	output, err := conn.GetMember(ctx, input)	if errs.IsA[*types.AccessDeniedException](err) || errs.IsA[*types.ResourceNotFoundException](err) {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if err != nil {
turn nil, err
	}	if output == nil || output.Member == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	if status := output.Member.RelationshipStatus; status == types.RelationshipStatusRemoved {
turn nil, &retry.NotFoundError{
Message:string(status),
LastRequest: input,	}	return output.Member, nil
}func waitMemberAssociationCreated(ctx context.Context, conn *inspector2.Client, id string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
nding: enum.Slice(types.RelationshipStatusCreated),
rget:  enum.Slice(types.RelationshipStatusEnabled),
fresh: statusMemberAssociation(ctx, conn, id),
meout: timeout,
	}	_, err := stateConf.WaitForStateContext(ctx)
	return err
}func waitMemberAssociationDeleted(ctx context.Context, conn *inspector2.Client, id string, timeout time.Duration) error {
	stateConf := &retry.StateChangeConf{
nding: enum.Slice(types.RelationshipStatusCreated, types.RelationshipStatusEnabled),
rget:  []string{},
fresh: statusMemberAssociation(ctx, conn, id),
meout: timeout,
	}	_, err := stateConf.WaitForStateContext(ctx)
	return err
}func statusMemberAssociation(ctx context.Context, conn *inspector2.Client, id string) retry.StateRefreshFunc {
	return func() (any, string, error) {
mber, err := FindMemberByAccountID(ctx, conn, id)
 tfresource.NotFound(err) {
return nil, "", nil err != nil {
return nil, "", err
ern member, string(member.RelationshipStatus), nil
	}
}
