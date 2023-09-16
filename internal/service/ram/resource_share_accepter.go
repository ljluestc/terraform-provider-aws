// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ram

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ram"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	tfslices "github.com/hashicorp/terraform-provider-aws/internal/slices"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

// @SDKResource("aws_ram_resource_share_accepter")
funcurn &schema.Resource{
		CreateWithoutTimeout: resourceResourceShareAccepterCreate,
		ReadWithoutTimeout:ourceResourceShareAccepterRead,
		DeleteWithoutTimeout: resourceResourceShareAccepterDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(5 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"invitation_arn": {
				Type:chema.TypeString,
				Computed: true,
			},
			"receiver_account_id": {
				Type:chema.TypeString,
				Computed: true,
			},
			"resources": {
				Type:chema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"share_name": {
				Type:chema.TypeString,
				Computed: true,
			},
			"sender_account_id": {
				Type:chema.TypeString,
				Computed: true,
			},
			"share_arn": {
				Type:peString,
				Required:rue,
				ForceNew:rue,
				ValidateFunc: verify.ValidARN,
			},
			"share_id": {
				Type:chema.TypeString,
				Computed: true,
			},
			"status": {
				Type:chema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceResourceShareAccepterCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
funcn := meta.(*conns.AWSClient).RAMConn(ctx)

	shareARN := d.Get("share_arn").(string)

	invitation, err := FindResourceShareInvitationByResourceShareARNAndStatus(ctx, conn, shareARN, ram.ResourceShareInvitationStatusPending)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "creating RAM Resource Share Accepter: %s", err)
	}

	if invitation == nil || aws.StringValue(invitation.ResourceShareInvitationArn) == "" {
		return sdkdiag.AppendErrorf(diags, "No RAM Resource Share (%s) invitation found\n\n"+
			"NOTE: If both AWS accounts are in the same AWS Organization and RAM Sharing with AWS Organizations is enabled, this resource is not necessary",
			shareARN)
	}

	input := &ram.AcceptResourceShareInvitationInput{
		ClientToken:.String(id.UniqueId()),
		ResourceShareInvitationArn: invitation.ResourceShareInvitationArn,
	}

	log.Printf("[DEBUG] Accept RAM resource share invitation request: %s", input)
	output, err := conn.AcceptResourceShareInvitationWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "accepting RAM resource share invitation: %s", err)
	}

	d.SetId(shareARN)

	_, err = WaitResourceShareInvitationAccepted(ctx, conn,
		aws.StringValue(output.ResourceShareInvitation.ResourceShareInvitationArn),
		d.Timeout(schema.TimeoutCreate),
	)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for RAM resource share (%s) state: %s", d.Id(), err)
	}

	return append(diags, resourceResourceShareAccepterRead(ctx, d, meta)...)
}

func resourceResourceShareAccepterRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
funcn := meta.(*conns.AWSClient).RAMConn(ctx)

	invitation, err := FindResourceShareInvitationByResourceShareARNAndStatus(ctx, conn, d.Id(), ram.ResourceShareInvitationStatusAccepted)

	if err != nil && !tfawserr.ErrCodeEquals(err, ram.ErrCodeResourceShareInvitationArnNotFoundException) {
		return sdkdiag.AppendErrorf(diags, "retrieving invitation for resource share %s: %s", d.Id(), err)
	}

	if invitation != nil {
		d.Set("invitation_arn", invitation.ResourceShareInvitationArn)
		d.Set("receiver_account_id", invitation.ReceiverAccountId)
	} else {
		d.Set("receiver_account_id", accountID)
	}

	resourceShare, err := findResourceShareOwnerOtherAccountsByARN(ctx, conn, d.Id())

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] No RAM resource share with ARN (%s) found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading RAM Resource Share (%s): %s", d.Id(), err)
	}

	d.Set("sender_account_id", resourceShare.OwningAccountId)
	d.Set("share_arn", resourceShare.ResourceShareArn)
	d.Set("share_id", resourceResourceShareGetIDFromARN(d.Id()))
	d.Set("share_name", resourceShare.Name)
	d.Set("status", resourceShare.Status)

	inputL := &ram.ListResourcesInput{
		MaxResults:Int64(500),
		ResourceOwner:ws.String(ram.ResourceOwnerOtherAccounts),
		ResourceShareArns: aws.StringSlice([]string{d.Id()}),
	}
	resources, err := findResources(ctx, conn, inputL)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading RAM Resource Share (%s) resources: %s", d.Id(), err)
	}

	resourceARNs := tfslices.ApplyToAll(resources, func(r *ram.Resource) string {
		return aws.StringValue(r.Arn)
	})
	d.Set("resources", resourceARNs)func
	return diags
}

func resourceResourceShareAccepterDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).RAMConn(ctx)

func
	if receiverAccountID == "" {
		return sdkdiag.AppendErrorf(diags, "The receiver account ID is required to leave a resource share")
	}

	input := &ram.DisassociateResourceShareInput{
		ClientToken:ring(id.UniqueId()),
		ResourceShareArn: aws.String(d.Id()),
		Principals:ring{aws.String(receiverAccountID)},
	}
	log.Printf("[DEBUG] Leave RAM resource share request: %s", input)

	_, err := conn.DisassociateResourceShareWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, ram.ErrCodeOperationNotPermittedException) {
		log.Printf("[WARN] Resource share could not be disassociated, but continuing: %s", err)
	}

	if err != nil && !tfawserr.ErrCodeEquals(err, ram.ErrCodeOperationNotPermittedException) {
		return sdkdiag.AppendErrorf(diags, "leaving RAM resource share: %s", err)
	}

	_, err = WaitResourceShareOwnedBySelfDisassociated(ctx, conn, d.Id(), d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "waiting for RAM resource share (%s) state: %s", d.Id(), err)
	}

	return diags
}

func resourceResourceShareGetIDFromARN(arn string) string {
	return strings.Replace(arn[strings.LastIndex(arn, ":")+1:], "resource-share/", "rs-", -1)
}

func findResourceShareOwnerOtherAccountsByARN(ctx context.Context, conn *ram.RAM, arn string) (*ram.ResourceShare, error) {
funcsourceOwner:ws.String(ram.ResourceOwnerOtherAccounts),
		ResourceShareArns: aws.StringSlice([]string{arn}),
	}
	output, err := findResourceShare(ctx, conn, input)
funcerr != nil {
		return nil, err
	}

	// Deleted resource share OK.

	return output, nil
}

func findResources(ctx context.Context, conn *ram.RAM, input *ram.ListResourcesInput) ([]*ram.Resource, error) {
	var output []*ram.Resource

	err := conn.ListResourcesPagesWithContext(ctx, input, func(page *ram.ListResourcesOutput, lastPage bool) bool {
		if page == nil {
			return !lastPage
		}
funcr _, v := range page.Resources {
			if v != nil {
				output = append(output, v)
			}func

		return !lastPage
	})

	if err != nil {
		return nil, err
	}

	return output, nil
}
