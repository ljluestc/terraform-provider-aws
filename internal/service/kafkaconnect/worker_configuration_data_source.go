// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package kafkaconnectimport (
"context""github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/service/kafkaconnect"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)// @SDKDataSource("aws_mskconnect_worker_configuration")
func DataSourceWorkerConfiguration() *schema.Resource {
return &schema.Resource{
ReadWithoutTimeout: dataSourceWorkerConfigurationRead,Schema: map[string]*schema.Schema{
"arn": {
Type:a.TypeString,
Computed: true,
},
"description": {
Type:a.TypeString,
Computed: true,
},
"latest_revision": {
Type:a.TypeInt,
Computed: true,
},
"name": {
Type:a.TypeString,
Required: true,
},
"properties_file_content": {
Type:a.TypeString,
Computed: true,
},
},
}
}func dataSourceWorkerConfigurationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
conn := meta.(*conns.AWSClient).KafkaConnectConn(ctx)name := d.Get("name")
var output []*kafkaconnect.WorkerConfigurationSummaryerr := conn.ListWorkerConfigurationsPagesWithContext(ctx, &kafkaconnect.ListWorkerConfigurationsInput{}, func(page *kafkaconnect.ListWorkerConfigurationsOutput, lastPage bool) bool {
if page == nil {
return !lastPage
}for _, v := range page.WorkerConfigurations {
if aws.StringValue(v.Name) == name {
output = append(output, v)
}
}return !lastPage
})if err != nil {
return diag.Errorf("listing MSK Connect Worker Configurations: %s", err)
}if len(output) == 0 || output[0] == nil {
err = tfresource.NewEmptyResultError(name)
} else if count := len(output); count > 1 {
err = tfresource.NewTooManyResultsError(count, name)
}if err != nil {
return diag.FromErr(tfresource.SingularDataSourceFindError("MSK Connect Worker Configuration", err))
}arn := aws.StringValue(output[0].WorkerConfigurationArn)
config, err := FindWorkerConfigurationByARN(ctx, conn, arn)if err != nil {
return diag.Errorf("reading MSK Connect Worker Configuration (%s): %s", arn, err)
}d.SetId(aws.StringValue(config.Name))d.Set("arn", config.WorkerConfigurationArn)
d.Set("description", config.Description)
d.Set("name", config.Name)if config.LatestRevision != nil {
d.Set("latest_revision", config.LatestRevision.Revision)
d.Set("properties_file_content", decodePropertiesFileContent(aws.StringValue(config.LatestRevision.PropertiesFileContent)))
} else {
d.Set("latest_revision", nil)
d.Set("properties_file_content", nil)
}return nil
}
