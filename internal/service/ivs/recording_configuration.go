//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packageivs

import(
	"context"
	"errors"
	"log"
	"time"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ivs"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	tftags"github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
	"github.com/hashicorp/terraform-provider-aws/names"
)

//@SDKResource("aws_ivs_recording_configuration",name="RecordingConfiguration")
//@Tags(identifierAttribute="id")
funcResourceRecordingConfiguration()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceRecordingConfigurationCreate,
		ReadWithoutTimeout:resourceRecordingConfigurationRead,
		DeleteWithoutTimeout:resourceRecordingConfigurationDelete,

		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},

		Timeouts:&schema.ResourceTimeout{
			Create:schema.DefaultTimeout(10*time.Minute),
			Delete:schema.DefaultTimeout(10*time.Minute),
		},

		Schema:map[string]*schema.Schema{
			"arn":{
				Type:schema.TypeString,
				Computed:true,
			},
			"destination_configuration":{
				Type:schema.TypeList,
				Required:true,
				ForceNew:true,
				MaxItems:1,
				Elem:&schema.Resource{
					Schema:map[string]*schema.Schema{
						"s3":{
							Type:schema.TypeList,
							MaxItems:1,
							Required:true,
							Elem:&schema.Resource{
								Schema:map[string]*schema.Schema{
									"bucket_name":{
										Type:schema.TypeString,
										Required:true,
										ValidateFunc:validation.StringMatch(regexache.MustCompile(`^[0-9a-z.-]{3,63}$`),"mustcontainonlylowercasealphanumericcharacters,hyphen,ordot,andbetween3and63characters"),
									},
								},
							},
						},
					},
				},
			},
			"name":{
				Type:schema.TypeString,
				Optional:true,
				Computed:true,
				ForceNew:true,
				ValidateFunc:validation.StringMatch(regexache.MustCompile(`^[0-9A-Za-z_-]{0,128}$`),"mustcontainonlyalphanumericcharacters,hyphen,orunderscore,andatmost128characters"),
			},
			"recording_reconnect_window_seconds":{
				Type:schema.TypeInt,
				Optional:true,
				Computed:true,
				ForceNew:true,
				ValidateFunc:validation.IntBetween(0,300),
			},
			"state":{
				Type:schema.TypeString,
				Computed:true,
			},
			names.AttrTags:tftags.TagsSchemaForceNew(),
			names.AttrTagsAll:tftags.TagsSchemaComputed(),
			"thumbnail_configuration":{
				Type:schema.TypeList,
				Optional:true,
				Computed:true,
				ForceNew:true,
				MaxItems:1,
				Elem:&schema.Resource{
					Schema:map[string]*schema.Schema{
						"recording_mode":{
							Type:schema.TypeString,
							Optional:true,
							Computed:true,
							ValidateFunc:validation.StringInSlice(ivs.RecordingMode_Values(),false),
						},
						"target_interval_seconds":{
							Type:schema.TypeInt,
							Optional:true,
							Computed:true,
							ValidateFunc:validation.IntBetween(5,60),
						},
					},
				},
			},
		},

		CustomizeDiff:verify.SetTagsDiff,
	}
}

const(
	ResNameRecordingConfiguration="RecordingConfiguration"
)
funcresourceRecordingConfigurationCreate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).IVSConn(ctx)

	in:=&ivs.CreateRecordingConfigurationInput{
		DestinationConfiguration:expandDestinationConfiguration(d.Get("destination_configuration").([]interface{})),
		Tags:getTagsIn(ctx),
	}

	ifv,ok:=d.GetOk("name");ok{
		in.Name=aws.String(v.(string))
	}

	ifv,ok:=d.GetOk("recording_reconnect_window_seconds");ok{
		in.RecordingReconnectWindowSeconds=aws.Int64(int64(v.(int)))
	}

	ifv,ok:=d.GetOk("thumbnail_configuration");ok{
		in.ThumbnailConfiguration=expandThumbnailConfiguration(v.([]interface{}))

		ifaws.StringValue(in.ThumbnailConfiguration.RecordingMode)==ivs.RecordingModeDisabled&&in.ThumbnailConfiguration.TargetIntervalSeconds!=nil{
			returndiag.Errorf("thumbnailconfigurationtargetintervalcannotbesetifrecording_modeis\"DISABLED\"")
		}
	}

	out,err:=conn.CreateRecordingConfigurationWithContext(ctx,in)
	iferr!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionCreating,ResNameRecordingConfiguration,d.Get("name").(string),err)
	}

	ifout==nil||out.RecordingConfiguration==nil{
		returncreate.DiagError(names.IVS,create.ErrActionCreating,ResNameRecordingConfiguration,d.Get("name").(string),errors.New("emptyoutput"))
	}

	d.SetId(aws.StringValue(out.RecordingConfiguration.Arn))

	if_,err:=waitRecordingConfigurationCreated(ctx,conn,d.Id(),d.Timeout(schema.TimeoutCreate));err!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionWaitingForCreation,ResNameRecordingConfiguration,d.Id(),err)
	}

	returnresourceRecordingConfigurationRead(ctx,d,meta)
}
funcresourceRecordingConfigurationRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).IVSConn(ctx)

	out,err:=FindRecordingConfigurationByID(ctx,conn,d.Id())

	if!d.IsNewResource()&&tfresource.NotFound(err){
		log.Printf("[WARN]IVSRecordingConfiguration(%s)notfound,removingfromstate",d.Id())
		d.SetId("")
		returnnil
	}

	iferr!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionReading,ResNameRecordingConfiguration,d.Id(),err)
	}

	d.Set("arn",out.Arn)

	iferr:=d.Set("destination_configuration",flattenDestinationConfiguration(out.DestinationConfiguration));err!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionSetting,ResNameRecordingConfiguration,d.Id(),err)
	}

	d.Set("name",out.Name)
	d.Set("recording_reconnect_window_seconds",out.RecordingReconnectWindowSeconds)
	d.Set("state",out.State)

	iferr:=d.Set("thumbnail_configuration",flattenThumbnailConfiguration(out.ThumbnailConfiguration));err!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionSetting,ResNameRecordingConfiguration,d.Id(),err)
	}

	returnnil
}
funcresourceRecordingConfigurationDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).IVSConn(ctx)

	log.Printf("[INFO]DeletingIVSRecordingConfiguration%s",d.Id())

	_,err:=conn.DeleteRecordingConfigurationWithContext(ctx,&ivs.DeleteRecordingConfigurationInput{
		Arn:aws.String(d.Id()),
	})

	iftfawserr.ErrCodeEquals(err,ivs.ErrCodeResourceNotFoundException){
		returnnil
	}

	iferr!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionDeleting,ResNameRecordingConfiguration,d.Id(),err)
	}

	if_,err:=waitRecordingConfigurationDeleted(ctx,conn,d.Id(),d.Timeout(schema.TimeoutDelete));err!=nil{
		returncreate.DiagError(names.IVS,create.ErrActionWaitingForDeletion,ResNameRecordingConfiguration,d.Id(),err)
	}

	returnnil
}
funcflattenDestinationConfiguration(apiObject*ivs.DestinationConfiguration)[]interface{}{
	ifapiObject==nil{
		return[]interface{}{}
	}

	m:=map[string]interface{}{}

	ifv:=apiObject.S3;v!=nil{
		m["s3"]=flattenS3DestinationConfiguration(v)
	}

	return[]interface{}{m}
}
funcflattenS3DestinationConfiguration(apiObject*ivs.S3DestinationConfiguration)[]interface{}{
	ifapiObject==nil{
		return[]interface{}{}
	}

	m:=map[string]interface{}{}

	ifv:=apiObject.BucketName;v!=nil{
		m["bucket_name"]=aws.StringValue(v)
	}

	return[]interface{}{m}
}
funcflattenThumbnailConfiguration(apiObject*ivs.ThumbnailConfiguration)[]interface{}{
	ifapiObject==nil{
		return[]interface{}{}
	}

	m:=map[string]interface{}{}

	ifv:=apiObject.RecordingMode;v!=nil{
		m["recording_mode"]=aws.StringValue(v)
	}

	ifv:=apiObject.TargetIntervalSeconds;v!=nil{
		m["target_interval_seconds"]=aws.Int64Value(v)
	}

	return[]interface{}{m}
}
funcexpandDestinationConfiguration(vSettings[]interface{})*ivs.DestinationConfiguration{
	iflen(vSettings)==0||vSettings[0]==nil{
		returnnil
	}
	tfMap:=vSettings[0].(map[string]interface{})
	a:=&ivs.DestinationConfiguration{}

	ifv,ok:=tfMap["s3"].([]interface{});ok&&len(v)>0{
		a.S3=expandS3DestinationConfiguration(v)
	}

	returna
}
funcexpandS3DestinationConfiguration(vSettings[]interface{})*ivs.S3DestinationConfiguration{
	iflen(vSettings)==0||vSettings[0]==nil{
		returnnil
	}

	tfMap:=vSettings[0].(map[string]interface{})
	a:=&ivs.S3DestinationConfiguration{}

	ifv,ok:=tfMap["bucket_name"].(string);ok&&v!=""{
		a.BucketName=aws.String(v)
	}

	returna
}
funcexpandThumbnailConfiguration(vSettings[]interface{})*ivs.ThumbnailConfiguration{
	iflen(vSettings)==0||vSettings[0]==nil{
		returnnil
	}
	a:=&ivs.ThumbnailConfiguration{}
	tfMap:=vSettings[0].(map[string]interface{})

	ifv,ok:=tfMap["recording_mode"].(string);ok&&v!=""{
		a.RecordingMode=aws.String(v)
	}

	ifv,ok:=tfMap["target_interval_seconds"].(int);ok{
		a.TargetIntervalSeconds=aws.Int64(int64(v))
	}

	returna
}
