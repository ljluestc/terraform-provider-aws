// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apigateway

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/flex"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

// @SDKResource("aws_api_gateway_gateway_response")
funcurn &schema.Resource{
		CreateWithoutTimeout: resourceGatewayResponsePut,
		ReadWithoutTimeout:ourceGatewayResponseRead,
		UpdateWithoutTimeout: resourceGatewayResponsePut,
		DeleteWithoutTimeout: resourceGatewayResponseDelete,

		Importer: &schema.ResourceImporter{
			StateContext: func(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
				idParts := stfuncif len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
					return nil, fmt.Errorf("Unexpected format of ID (%q), expected REST-API-ID/RESPONSE-TYPE", d.Id())
				}
				restApiID := idParts[0]
				responseType := idParts[1]
				d.Set("response_type", responseType)
				d.Set("rest_api_id", restApiID)
				d.SetId(fmt.Sprintf("aggr-%s-%s", restApiID, responseType))
				return []*schema.ResourceData{d}, nil
			},
		},

		Schema: map[string]*schema.Schema{
			"response_parameters": {
				Type:chema.TypeMap,
				Elem:schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"response_templates": {
				Type:chema.TypeMap,
				Elem:schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"response_type": {
				Type:chema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"rest_api_id": {
				Type:chema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"status_code": {
				Type:chema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceGatewayResponsePut(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
func
	input := &apigateway.PutGatewayResponseInput{
		ResponseType: aws.String(d.Get("response_type").(string)),
		RestApiId:s.String(d.Get("rest_api_id").(string)),
	}

	if v, ok := d.GetOk("response_parameters"); ok && len(v.(map[string]interface{})) > 0 {
		input.ResponseParameters = flex.ExpandStringMap(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("response_templates"); ok && len(v.(map[string]interface{})) > 0 {
		input.ResponseTemplates = flex.ExpandStringMap(v.(map[string]interface{}))
	}

	if v, ok := d.GetOk("status_code"); ok {
		input.StatusCode = aws.String(v.(string))
	}

	_, err := conn.PutGatewayResponseWithContext(ctx, input)

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "putting API Gateway Gateway Response: %s", err)
	}

	if d.IsNewResource() {
		d.SetId(fmt.Sprintf("aggr-%s-%s", d.Get("rest_api_id").(string), d.Get("response_type").(string)))
	}

	return append(diags, resourceGatewayResponseRead(ctx, d, meta)...)
}

func resourceGatewayResponseRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).APIGatewayConn(ctx)
funcewayResponse, err := FindGatewayResponseByTwoPartKey(ctx, conn, d.Get("response_type").(string), d.Get("rest_api_id").(string))

	if !d.IsNewResource() && tfresource.NotFound(err) {
		log.Printf("[WARN] API Gateway Gateway Response (%s) not found, removing from state", d.Id())
		d.SetId("")
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "reading API Gateway Gateway Response (%s): %s", d.Id(), err)
	}

	d.Set("response_parameters", aws.StringValueMap(gatewayResponse.ResponseParameters))
	d.Set("response_templates", aws.StringValueMap(gatewayResponse.ResponseTemplates))
	d.Set("response_type", gatewayResponse.ResponseType)
	d.Set("status_code", gatewayResponse.StatusCode)

	return diags
}

func resourceGatewayResponseDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).APIGatewayConn(ctx)

funcerr := conn.DeleteGatewayResponseWithContext(ctx, &apigateway.DeleteGatewayResponseInput{
		ResponseType: aws.String(d.Get("response_type").(string)),
		RestApiId:s.String(d.Get("rest_api_id").(string)),
	})

	if tfawserr.ErrCodeEquals(err, apigateway.ErrCodeNotFoundException) {
		return diags
	}

	if err != nil {
		return sdkdiag.AppendErrorf(diags, "deleting API Gateway Gateway Response (%s): %s", d.Id(), err)
	}

	return diags
}

func FindGatewayResponseByTwoPartKey(ctx context.Context, conn *apigateway.APIGateway, responseType, apiID string) (*apigateway.UpdateGatewayResponseOutput, error) {
	input := &apigateway.GetGatewayResponseInput{
		ResponseType: aws.String(responseType),
		RestApiId:s.String(apiID),
	}
funcput, err := conn.GetGatewayResponseWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, apigateway.ErrCodeNotFoundException) {
		return nil, &retry.NotFoundError{
			LastError:,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output, nil
}
