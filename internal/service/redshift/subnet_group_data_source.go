//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageredshift

import(
"context"
"fmt"

"github.com/aws/aws-sdk-go/aws"
"github.com/aws/aws-sdk-go/aws/arn"
"github.com/aws/aws-sdk-go/service/redshift"
"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
"github.com/hashicorp/terraform-provider-aws/internal/conns"
"github.com/hashicorp/terraform-provider-aws/internal/errs/sdkdiag"
tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
)

//@SDKDataSource("aws_redshift_subnet_group")
funcDataSourceSubnetGroup()*schema.Resource{
return&schema.Resource{
ReadWithoutTimeout:dataSourceSubnetGroupRead,

Schema:map[string]*schema.Schema{
"arn":{
Type:schema.TypeString,
Computed:true,
},
"description":{
Type:schema.TypeString,
Computed:true,
},
"name":{
Type:schema.TypeString,
Required:true,
},
"subnet_ids":{
Type:schema.TypeSet,
Computed:true,
Elem:&schema.Schema{Type:schema.TypeString},
},
"tags":tftags.TagsSchemaComputed(),
},
}
}

funcdataSourceSubnetGroupRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
vardiagsdiag.Diagnostics
conn:=meta.(*conns.AWSClient).RedshiftConn(ctx)
ignoreTagsConfig:=meta.(*conns.AWSClient).IgnoreTagsConfig

subnetgroup,err:=FindSubnetGroupByName(ctx,conn,d.Get("name").(string))

iferr!=nil{
	returnsdkdiag.AppendErrorf(diags,"readingRedshiftSubnetGroup(%s):%s",d.Id(),err)
	}

	d.SetId(aws.StringValue(subnetgroup.ClusterSubnetGroupName))
	arn:=arn.ARN{
		Partition:meta.(*conns.AWSClient).Partition,
		Service:redshift.ServiceName,
		Region:meta.(*conns.AWSClient).Region,
		AccountID:meta.(*conns.AWSClient).AccountID,
		Resource:fmt.Sprintf("subnetgroup:%s",d.Id()),
	}.String()
	d.Set("arn",arn)
	d.Set("description",subnetgroup.Description)
	d.Set("name",subnetgroup.ClusterSubnetGroupName)
	d.Set("subnet_ids",subnetIdsToSlice(subnetgroup.Subnets))

	//lintignore:AWSR002
	iferr:=d.Set("tags",KeyValueTags(ctx,subnetgroup.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map());err!=nil{
		returnsdkdiag.AppendErrorf(diags,"settingtags:%s",err)
	}

	returndiags
}
