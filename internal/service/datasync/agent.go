// Copyright (c) HashiCorp, Inc.// SPDX-License-Identifier: MPL-2.0package datasyncimport (	"context"	"fmt"	"log"	"net"	"net/http"	"time"	"github.com/aws/aws-sdk-go/aws"	"github.com/aws/aws-sdk-go/service/datasync"	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"	"github.com/hashicorp/terraform-provider-aws/internal/conns"	"github.com/hashicorp/terraform-provider-aws/internal/errs"	"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"	"github.com/hashicorp/terraform-provider-aws/internal/flex"	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"	"github.com/hashicorp/terraform-provider-aws/internal/verify"	"github.com/hashicorp/terraform-provider-aws/names")// @SDKResource("aws_datasync_agent", name="Agent")// @Tags(identifierAttribute="id")func ResourceAgent() *schema.Resource {	return &schema.Resource{CreateWithoutTimeout: resourceAgentCreate,ReadWithoutTimeout:resourceAgentRead,UpdateWithoutTimeout: resourceAgentUpdate,DeleteWithoutTimeout: resourceAgentDelete,Importer: &schema.ResourceImporter{StateContext: schema.ImportStatePassthroughContext,},Timeouts: &schema.ResourceTimeout{Create: schema.DefaultTimeout(10 * time.Minute),},Schema: map[string]*schema.Schema{"arn": {Type:schema.TypeString,Computed: true,},"activation_key": {Type: schema.TypeString,Optional: true,Computed: true,ForceNew: true,ExactlyOneOf:  []string{"activation_key", "ip_address"},ConflictsWith: []string{"private_link_endpoint"},},"ip_address": {Type:schema.TypeString,Optional:true,Computed:true,ForceNew:true,ExactlyOneOf: []string{"activation_key", "ip_address"},},"private_link_endpoint": {
Type: schema.TypeString,
Optional: true,
Computed: true,
ForceNew: true,
ConflictsWith: []string{"activation_key"},
},
"name": {
Type:schema.TypeString,
Optional: true,
},
"security_group_arns": {
Type:schema.TypeSet,
Optional: true,
ForceNew: true,
Elem:&schema.Schema{Type: schema.TypeString},
},
"subnet_arns": {
Type:schema.TypeSet,
Optional: true,
ForceNew: true,
Elem:&schema.Schema{Type: schema.TypeString},
},
names.AttrTags:tftags.TagsSchema(),
names.AttrTagsAll: tftags.TagsSchemaComputed(),
"vpc_endpoint_id": {
Type:schema.TypeString,
Optional: true,
ForceNew: true,
},
CuomizeDiff: verify.SetTagsDiff,
	}
}func resourceAgentCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	activationKey := d.Get("activation_key").(string)
	agentIpAddress := d.Get("ip_address").(string)	// Perform one time fetch of activation key from gateway IP address.
	if activationKey == "" {
ient := &http.Client{
CheckRedirect: func(req *http.Request, via []*http.Request) error {
return http.ErrUseLastResponse
},
Timeout: time.Second * 10,gion := meta.(*conns.AWSClient).RegionvarequestURL string
 v, ok := d.GetOk("private_link_endpoint"); ok {
requestURL = fmt.Sprintf("http://%s/?gatewayType=SYNC&activationRegion=%s&endpointType=PRIVATE_LINK&privateLinkEndpoint=%s", agentIpAddress, region, v.(string))
else {
requestURL = fmt.Sprintf("http://%s/?gatewayType=SYNC&activationRegion=%s", agentIpAddress, region)
eest, err := http.NewRequest("GET", requestURL, nil)
 err != nil {
return sdkdiag.AppendErrorf(diags, "creating HTTP request: %s", err)
aresponse *http.Response
r = retry.RetryContext(ctx, d.Timeout(schema.TimeoutCreate), func() *retry.RetryError {
response, err = client.Do(request)if errs.IsA[net.Error](err) {
return retry.RetryableError(fmt.Errorf("making HTTP request: %w", err))
}if err != nil {
return retry.NonRetryableError(fmt.Errorf("making HTTP request: %w", err))
}if response == nil {
return retry.NonRetryableError(fmt.Errorf("no response for activation key request"))
}log.Printf("[DEBUG] Received HTTP response: %#v", response)
if expected := http.StatusFound; expected != response.StatusCode {
return retry.NonRetryableError(fmt.Errorf("expected HTTP status code %d, received: %d", expected, response.StatusCode))
}redirectURL, err := response.Location()
if err != nil {
return retry.NonRetryableError(fmt.Errorf("extracting HTTP Location header: %w", err))
}if errorType := redirectURL.Query().Get("errorType"); errorType == "PRIVATE_LINK_ENDPOINT_UNREACHABLE" {
errMessage := fmt.Errorf("during activation: %s", errorType)
return retry.RetryableError(errMessage)
}activationKey = redirectURL.Query().Get("activationKey")return nil
iffresource.TimedOut(err) {
return sdkdiag.AppendErrorf(diags, "timeout retrieving activation key from IP Address (%s): %s", agentIpAddress, err)
frr != nil {
return sdkdiag.AppendErrorf(diags, "retrieving activation key from IP Address (%s): %s", agentIpAddress, err)
fctivationKey == "" {
return sdkdiag.AppendErrorf(diags, "empty activationKey received from IP Address: %s", agentIpAddress)	}	input := &datasync.CreateAgentInput{
tivationKey: aws.String(activationKey),
gs: getTagsIn(ctx),
	}	if v, ok := d.GetOk("name"); ok {
put.AgentName = aws.String(v.(string))
	}	if v, ok := d.GetOk("security_group_arns"); ok {
put.SecurityGroupArns = flex.ExpandStringSet(v.(*schema.Set))
	}	if v, ok := d.GetOk("subnet_arns"); ok {
put.SubnetArns = flex.ExpandStringSet(v.(*schema.Set))
	}	if v, ok := d.GetOk("vpc_endpoint_id"); ok {
put.VpcEndpointId = aws.String(v.(string))
	}	output, err := conn.CreateAgentWithContext(ctx, input)	if err != nil {
turn sdkdiag.AppendErrorf(diags, "creating DataSync Agent: %s", err)
	}	d.SetId(aws.StringValue(output.AgentArn))	_, err = tfresource.RetryWhenNotFound(ctx, d.Timeout(schema.TimeoutCreate), func() (interface{}, error) {
turn FindAgentByARN(ctx, conn, d.Id())
	})	if err != nil {
turn sdkdiag.AppendErrorf(diags, "waiting for DataSync Agent (%s) create: %s", d.Id(), err)
	}	return append(diags, resourceAgentRead(ctx, d, meta)...)
}func resourceAgentRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	output, err := FindAgentByARN(ctx, conn, d.Id())	if !d.IsNewResource() && tfresource.NotFound(err) {
g.Printf("[WARN] DataSync Agent (%s) not found, removing from state", d.Id())
SetId("")
turn diags
	}	if err != nil {
turn sdkdiag.AppendErrorf(diags, "reading DataSync Agent (%s): %s", d.Id(), err)
	}	d.Set("arn", output.AgentArn)
	d.Set("name", output.Name)
	if plc := output.PrivateLinkConfig; plc != nil {
Set("private_link_endpoint", plc.PrivateLinkEndpoint)
Set("security_group_arns", flex.FlattenStringList(plc.SecurityGroupArns))
Set("subnet_arns", flex.FlattenStringList(plc.SubnetArns))
Set("vpc_endpoint_id", plc.VpcEndpointId)
	} else {
Set("private_link_endpoint", "")
Set("security_group_arns", nil)
Set("subnet_arns", nil)
Set("vpc_endpoint_id", "")
	}	return diags
}func resourceAgentUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	if d.HasChange("name") {
put := &datasync.UpdateAgentInput{
AgentArn: aws.String(d.Id()),
Name:aws.String(d.Get("name").(string)),
,rr := conn.UpdateAgentWithContext(ctx, input)if e != nil {
return sdkdiag.AppendErrorf(diags, "updating DataSync Agent (%s): %s", d.Id(), err)	}	return append(diags, resourceAgentRead(ctx, d, meta)...)
}func resourceAgentDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	conn := meta.(*conns.AWSClient).DataSyncConn(ctx)	log.Printf("[DEBUG] Deleting DataSync Agent: %s", d.Id())
	_, err := conn.DeleteAgentWithContext(ctx, &datasync.DeleteAgentInput{
entArn: aws.String(d.Id()),
	})	if tfawserr.ErrMessageContains(err, datasync.ErrCodeInvalidRequestException, "does not exist") {
turn diags
	}	if err != nil {
turn sdkdiag.AppendErrorf(diags, "deleting DataSync Agent (%s): %s", d.Id(), err)
	}	return diags
}func FindAgentByARN(ctx context.Context, conn *datasync.DataSync, arn string) (*datasync.DescribeAgentOutput, error) {
	input := &datasync.DescribeAgentInput{
entArn: aws.String(arn),
	}	output, err := conn.DescribeAgentWithContext(ctx, input)	if tfawserr.ErrMessageContains(err, datasync.ErrCodeInvalidRequestException, "does not exist") {
turn nil, &retry.NotFoundError{
LastError:err,
LastRequest: input,	}	if err != nil {
turn nil, err
	}	if output == nil {
turn nil, tfresource.NewEmptyResultError(input)
	}	return output, nil
}
