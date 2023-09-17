// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package servicecatalogimport (
"context"
"time""github.com/aws/aws-sdk-go/aws"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)// @SDKDataSource("aws_servicecatalog_product")
func DataSourceProduct() *schema.Resource {
return &schema.Resource{
ReadWithoutTimeout: dataSourceProductRead,Timeouts: &schema.ResourceTimeout{
Read: schema.DefaultTimeout(ProductReadTimeout),
},Schema: map[string]*schema.Schema{
"arn": {
Type:a.TypeString,
Computed: true,
},
"accept_language": {
Type:schema.TypeString,
Optional:
Default:,
ValidateFunc: validation.StringInSlice(AcceptLanguage_Values(), false),
},
"created_time": {
Type:a.TypeString,
Computed: true,
},
"description": {
Type:a.TypeString,
Computed: true,
},
"distributor": {
Type:a.TypeString,
Computed: true,
},
"has_default_path": {
Type:a.TypeBool,
Computed: true,
},
"id": {
Type:a.TypeString,
Required: true,
},
"name": {
Type:a.TypeString,
Computed: true,
},
"owner": {
Type:a.TypeString,
Computed: true,
},
"status": {
Type:a.TypeString,
Computed: true,
},
"support_description": {
Type:a.TypeString,
Computed: true,
},
"support_email": {
Type:a.TypeString,
Computed: true,
},
"support_url": {
Type:a.TypeString,
Computed: true,
},
"tags": tftags.TagsSchemaComputed(),
"type": {
Type:a.TypeString,
Computed: true,
},
},
}
}func dataSourceProductRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
var diags diag.Diagnostics
conn := meta.(*conns.AWSClient).ServiceCatalogConn(ctx)output, err := WaitProductReady(ctx, conn, d.Get("accept_language").(string), d.Get("id").(string), d.Timeout(schema.TimeoutRead))if err != nil {
return sdkdiag.AppendErrorf(diags, "describing Service Catalog Product: %s", err)
}if output == nil || output.ProductViewDetail == nil || output.ProductViewDetail.ProductViewSummary == nil {
return sdkdiag.AppendErrorf(diags, "getting Service Catalog Product: empty response")
}pvs := output.ProductViewDetail.ProductViewSummaryd.Set("arn", output.ProductViewDetail.ProductARN)
if output.ProductViewDetail.CreatedTime != nil {
d.Set("created_time", output.ProductViewDetail.CreatedTime.Format(time.RFC3339))
}
d.Set("description", pvs.ShortDescription)
d.Set("distributor", pvs.Distributor)
d.Set("has_default_path", pvs.HasDefaultPath)
d.Set("name", pvs.Name)
d.Set("owner", pvs.Owner)
d.Set("status", output.ProductViewDetail.Status)
d.Set("support_description", pvs.SupportDescription)
d.Set("support_email", pvs.SupportEmail)
d.Set("support_url", pvs.SupportUrl)
d.Set("type", pvs.Type)d.SetId(aws.StringValue(pvs.ProductId))ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfigif err := d.Set("tags", KeyValueTags(ctx, output.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
return sdkdiag.AppendErrorf(diags, "setting tags: %s", err)
}return diags
}
