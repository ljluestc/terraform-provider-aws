// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package connect

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/connect"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

// @SDKResource("aws_connect_lambda_
function_association")

func ResourceLambda
functionAssociation() *schema.Resource {
	return &schema.Resource{
		CreateWithoutTimeout: resourceLambda
functionAssociationCreate,
		ReadWithoutTimeout:   resourceLambda
functionAssociationRead,
		DeleteWithoutTimeout: resourceLambda
functionAssociationDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"
function_arn": {
				Type:schema.TypeString,
				Required:     true,
				ForceNew:     true,
				Validate
func: verify.ValidARN,
			},
			"instance_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
		},
	}
}


func resourceLambda
functionAssociationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ConnectConn(ctx)

	instanceId := d.Get("instance_id").(string)
	
functionArn := d.Get("
function_arn").(string)

	input := &connect.AssociateLambda
functionInput{
		InstanceId:  aws.String(instanceId),
		
functionArn: aws.String(
functionArn),
	}

	_, err := conn.AssociateLambda
functionWithContext(ctx, input)
	if err != nil {
		return diag.Errorf("creating Connect Lambda 
function Association (%s,%s): %s", instanceId, 
functionArn, err)
	}

	d.SetId(Lambda
functionAssociationCreateResourceID(instanceId, 
functionArn))

	return resourceLambda
functionAssociationRead(ctx, d, meta)
}


func resourceLambda
functionAssociationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ConnectConn(ctx)

	instanceID, 
functionArn, err := Lambda
functionAssociationParseResourceID(d.Id())

	if err != nil {
		return diag.FromErr(err)
	}

	lfaArn, err := FindLambda
functionAssociationByARNWithContext(ctx, conn, instanceID, 
functionArn)

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] Connect Lambda 
function Association (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	if err != nil {
		return diag.Errorf("finding Connect Lambda 
function Association by 
function ARN (%s): %s", 
functionArn, err)
	}

	d.Set("
function_arn", lfaArn)
	d.Set("instance_id", instanceID)

	return nil
}


func resourceLambda
functionAssociationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	conn := meta.(*conns.AWSClient).ConnectConn(ctx)

	instanceID, 
functionArn, err := Lambda
functionAssociationParseResourceID(d.Id())
	if err != nil {
		return diag.FromErr(err)
	}

	input := &connect.DisassociateLambda
functionInput{
		InstanceId:  aws.String(instanceID),
		
functionArn: aws.String(
functionArn),
	}

	_, err = conn.DisassociateLambda
functionWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, connect.ErrCodeResourceNotFoundException) {
		return nil
	}

	if err != nil {
		return diag.Errorf("deleting Connect Lambda 
function Association (%s): %s", d.Id(), err)
	}

	return nil
}
