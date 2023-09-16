//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageivs

import(
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

//@SDKDataSource("aws_ivs_stream_key")
funcDataSourceStreamKey()*schema.Resource{
	return&schema.Resource{
		ReadWithoutTimeout:dataSourceStreamKeyRead,
		Schema:map[string]*schema.Schema{
			"arn":{
				Type:schema.TypeString,
				Computed:true,
			},
			"channel_arn":{
				Type:schema.TypeString,
				Required:true,
			},
			"value":{
				Type:schema.TypeString,
				Computed:true,
			},
			"tags":tftags.TagsSchemaComputed(),
		},
	}
}

const(
	DSNameStreamKey="StreamKeyDataSource"
)
funcdataSourceStreamKeyRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).IVSConn(ctx)

	channelArn:=d.Get("channel_arn").(string)

	out,err:=FindStreamKeyByChannelID(ctx,conn,channelArn)
	iferr!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionReading,DSNameStreamKey,channelArn,err)
	}

	d.SetId(aws.StringValue(out.Arn))

	d.Set("arn",out.Arn)
	d.Set("channel_arn",out.ChannelArn)
	d.Set("value",out.Value)

	ignoreTagsConfig:=meta.(*conns.AWSClient).IgnoreTagsConfig

	//lintignore:AWSR002
	iferr:=d.Set("tags",KeyValueTags(ctx,out.Tags).IgnoreAWS().IgnoreConfig(ignoreTagsConfig).Map());err!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionSetting,DSNameStreamKey,d.Id(),err)
	}

	returnnil
}
