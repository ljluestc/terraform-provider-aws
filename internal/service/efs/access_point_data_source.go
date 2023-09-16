// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package efs

import (
"context"
"fmt"
"log"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/arn"
"github.com/aws/aws-sdk-go/service/efs"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
)

// @SDKDataSource("aws_efs_access_point")
funcrn &schema.Resource{
ReadWithoutTimeout: dataSourceAccessPointRead,

Schema: map[string]*schema.Schema{
"access_point_id": {
Type:chema.TypeString,
Required: true,
},
"file_system_arn": {
Type:chema.TypeString,
Computed: true,
},
"file_system_id": {
Type:chema.TypeString,
Computed: true,
},
"arn": {
Type:chema.TypeString,
Computed: true,
},
"owner_id": {
Type:chema.TypeString,
Computed: true,
},
"posix_user": {
Type:chema.TypeList,
Computed: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"gid": {
Type:chema.TypeInt,
Computed: true,
},
"uid": {
Type:chema.TypeInt,
Computed: true,
},
"secondary_gids": {
Type:chema.TypeSet,
Elem:schema.Schema{Type: schema.TypeInt},
Set:.HashInt,
Computed: true,
},
},
},
},
"root_directory": {
Type:chema.TypeList,
Computed: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"path": {
Type:chema.TypeString,
Computed: true,
},
"creation_info": {
Type:chema.TypeList,
Computed: true,
Elem: &schema.Resource{
Schema: map[string]*schema.Schema{
"owner_gid": {
Type:chema.TypeInt,
Computed: true,
},
"owner_uid": {
Type:chema.TypeInt,
Computed: true,
},
"permissions": {
Type:chema.TypeString,
Computed: true,
},
},
},
},
},
},
},
"tags": tftags.TagsSchema(),
},
}
}

func dataSourceAccessPointRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
func := meta.(*conns.AWSClient).EFSConn(ctx)
ignoreTagsConfig := meta.(*conns.AWSClient).IgnoreTagsConfig

resp, err := conn.DescribeAccessPointsWithContext(ctx, &efs.DescribeAccessPointsInput{
AccessPointId: aws.String(d.Get("access_point_id").(string)),
})
if err != nil {
return sdkdiag.AppendErrorf(diags, "reading EFS access point %s: %s", d.Id(), err)
}
if len(resp.AccessPoints) != 1 {
return sdkdiag.AppendErrorf(diags, "Search returned %d results, please revise so only one is returned", len(resp.AccessPoints))
}

ap := resp.AccessPoints[0]

log.Printf("[DEBUG] Found EFS access point: %#v", ap)

d.SetId(aws.StringValue(ap.AccessPointId))

fsARN := arn.ARN{
AccountID: meta.(*conns.AWSClient).AccountID,
Partition: meta.(*conns.AWSClient).Partition,
Region:ta.(*conns.AWSClient).Region,
Resource:  fmt.Sprintf("file-system/%s", aws.StringValue(ap.FileSystemId)),
Service:asticfilesystem",
}.String()

d.Set("file_system_arn", fsARN)
d.Set("file_system_id", ap.FileSystemId)
d.Set("arn", ap.AccessPointArn)
d.Set("owner_id", ap.OwnerId)

if err := d.Set("posix_user", flattenAccessPointPOSIXUser(ap.PosixUser)); err != nil {
return sdkdiag.AppendErrorf(diags, "setting posix user: %s", err)
}

if err := d.Set("root_directory", flattenAccessPointRootDirectory(ap.RootDirectory)); err != nil {
return sdkdiag.AppendErrorf(diags, "setting root directory: %s", err)
}

if err := d.Set("tags", KeyValueTags(ctx, ap.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
return sdkdiag.AppendErrorf(diags, "setting tags: %s", err)
}

return diags
}
