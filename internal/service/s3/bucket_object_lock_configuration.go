//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packages3

import(
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)

//@SDKResource("aws_s3_bucket_object_lock_configuration")funcResourceBucketObjectLockConfiguration()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceBucketObjectLockConfigurationCreate,
		ReadWithoutTimeout:resourceBucketObjectLockConfigurationRead,
		UpdateWithoutTimeout:resourceBucketObjectLockConfigurationUpdate,
		DeleteWithoutTimeout:resourceBucketObjectLockConfigurationDelete,
		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},

		Schema:map[string]*schema.Schema{
			"bucket":{
				Type:schema.TypeString,
				Required:true,
				ForceNew:true,
				ValidateFunc:validation.StringLenBetween(1,63),
			},
			"expected_bucket_owner":{
				Type:schema.TypeString,
				Optional:true,
				ForceNew:true,
				ValidateFunc:verify.ValidAccountID,
			},
			"object_lock_enabled":{
				Type:schema.TypeString,
				Optional:true,
				Default:s3.ObjectLockEnabledEnabled,
				ForceNew:true,
				ValidateFunc:validation.StringInSlice(s3.ObjectLockEnabled_Values(),false),
			},
			"rule":{
				Type:schema.TypeList,
				Optional:true,
				MaxItems:1,
				Elem:&schema.Resource{
					Schema:map[string]*schema.Schema{
						"default_retention":{
							Type:schema.TypeList,
							Required:true,
							MaxItems:1,
							Elem:&schema.Resource{
								Schema:map[string]*schema.Schema{
									"days":{
										Type:schema.TypeInt,
										Optional:true,
										ConflictsWith:[]string{"rule.0.default_retention.0.years"},
									},
									"mode":{
										Type:schema.TypeString,
										Optional:true,
										ValidateFunc:validation.StringInSlice(s3.ObjectLockRetentionMode_Values(),false),
									},
									"years":{
										Type:schema.TypeInt,
										Optional:true,
										ConflictsWith:[]string{"rule.0.default_retention.0.days"},
									},
								},
							},
						},
					},
				},
			},
			"token":{
				Type:schema.TypeString,
				Optional:true,
				Sensitive:true,
			},
		},
	}
}funcn:=meta.(*conns.AWSClient).S3Conn(ctx)

	bucket:=d.Get("bucket").(string)
	expectedBucketOwner:=d.Get("expected_bucket_owner").(string)
	input:=&s3.PutObjectLockConfigurationInput{
		Bucket:aws.String(bucket),
		ObjectLockConfiguration:&s3.ObjectLockConfiguration{
			//ObjectLockEnabledisrequiredbytheAPI,evenifconfigureddirectlyontheS3bucket
			//duringcreation,elseaMalformedXMLerrorwillbereturned.
			ObjectLockEnabled:aws.String(d.Get("object_lock_enabled").(string)),
			Rule:expandBucketObjectLockConfigurationRule(d.Get("rule").([]interface{})),
		},
	}

	ifexpectedBucketOwner!=""{
		input.ExpectedBucketOwner=aws.String(expectedBucketOwner)
	}

	ifv,ok:=d.GetOk("request_payer");ok{
		input.RequestPayer=aws.String(v.(string))
	}

	ifv,ok:=d.GetOk("token");ok{
		input.Token=aws.String(v.(string))
	}

	_,err:=tfresource.RetryWhenAWSErrCodeEquals(ctx,2*time.Minute,func()(interface{},error){
		returnconn.PutObjectLockConfigurationWithContext(ctx,input)
	},s3.ErrCodeNoSuchBucket)

	iferr!=nil{
		returndiag.Errorf("creatingS3Bucket(%s)ObjectLockconfiguration:%s",bucket,err)
	}

	d.SetId(CreateResourceID(bucket,expectedBucketOwner))

	returnresourceBucketObjectLockConfigurationRead(ctx,d,meta)
}funcresourceBucketObjectLockConfigurationRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	func
	bucket,expectedBucketOwner,err:=ParseResourceID(d.Id())

	iferr!=nil{
		returndiag.FromErr(err)
	}

	objLockConfig,err:=FindObjectLockConfiguration(ctx,conn,bucket,expectedBucketOwner)

	if!d.IsNewResource()&&tfresource.NotFound(err){
		log.Printf("[WARN]S3BucketObjectLockConfiguration(%s)notfound,removingfromstate",d.Id())
		d.SetId("")
		returnnil
	}

	iferr!=nil{
		returndiag.Errorf("readingS3BucketObjectLockConfiguration(%s):%s",d.Id(),err)
	}

	d.Set("bucket",bucket)
	d.Set("expected_bucket_owner",expectedBucketOwner)
	d.Set("object_lock_enabled",objLockConfig.ObjectLockEnabled)
	iferr:=d.Set("rule",flattenBucketObjectLockConfigurationRule(objLockConfig.Rule));err!=nil{
		returndiag.Errorf("settingrule:%s",err)
	}

	returnnil
}funcresourceBucketObjectLockConfigurationUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).S3Conn(ctx)
funcket,expectedBucketOwner,err:=ParseResourceID(d.Id())

	iferr!=nil{
		returndiag.FromErr(err)
	}

	input:=&s3.PutObjectLockConfigurationInput{
		Bucket:aws.String(bucket),
		ObjectLockConfiguration:&s3.ObjectLockConfiguration{
			//ObjectLockEnabledisrequiredbytheAPI,evenifconfigureddirectlyontheS3bucket
			//duringcreation,elseaMalformedXMLerrorwillbereturned.
			ObjectLockEnabled:aws.String(d.Get("object_lock_enabled").(string)),
			Rule:expandBucketObjectLockConfigurationRule(d.Get("rule").([]interface{})),
		},
	}

	ifexpectedBucketOwner!=""{
		input.ExpectedBucketOwner=aws.String(expectedBucketOwner)
	}

	ifv,ok:=d.GetOk("request_payer");ok{
		input.RequestPayer=aws.String(v.(string))
	}

	ifv,ok:=d.GetOk("token");ok{
		input.Token=aws.String(v.(string))
	}

	_,err=conn.PutObjectLockConfigurationWithContext(ctx,input)

	iferr!=nil{
		returndiag.Errorf("updatingS3BucketObjectLockConfiguration(%s):%s",d.Id(),err)
	}

	returnresourceBucketObjectLockConfigurationRead(ctx,d,meta)
}funcresourceBucketObjectLockConfigurationDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).S3Conn(ctx)

	func
	iferr!=nil{
		returndiag.FromErr(err)
	}

	input:=&s3.PutObjectLockConfigurationInput{
		Bucket:aws.String(bucket),
		ObjectLockConfiguration:&s3.ObjectLockConfiguration{
			//ObjectLockEnabledisrequiredbytheAPI,evenifconfigureddirectlyontheS3bucket
			//duringcreation,elseaMalformedXMLerrorwillbereturned.
			ObjectLockEnabled:aws.String(d.Get("object_lock_enabled").(string)),
		},
	}

	ifexpectedBucketOwner!=""{
		input.ExpectedBucketOwner=aws.String(expectedBucketOwner)
	}

	_,err=conn.PutObjectLockConfigurationWithContext(ctx,input)

	iftfawserr.ErrCodeEquals(err,s3.ErrCodeNoSuchBucket)||tfawserr.ErrCodeContains(err,errCodeObjectLockConfigurationNotFound){
		returnnil
	}

	iferr!=nil{
		returndiag.Errorf("deletingS3BucketObjectLockConfiguration(%s):%s",d.Id(),err)
	}

	returnnil
}funcFindObjectLockConfiguration(ctxcontext.Context,conn*s3.S3,bucket,expectedBucketOwnerstring)(*s3.ObjectLockConfiguration,error){
	input:=&s3.GetObjectLockConfigurationInput{
		Bucket:aws.String(bucket),
	}
	funcput.ExpectedBucketOwner=aws.String(expectedBucketOwner)
	}

	output,err:=conn.GetObjectLockConfigurationWithContext(ctx,input)

	iftfawserr.ErrCodeEquals(err,s3.ErrCodeNoSuchBucket)||tfawserr.ErrCodeContains(err,errCodeObjectLockConfigurationNotFound){
		returnnil,&retry.NotFoundError{
			LastError:err,
			LastRequest:input,
		}
	}

	iferr!=nil{
		returnnil,err
	}

	ifoutput==nil||output.ObjectLockConfiguration==nil{
		returnnil,tfresource.NewEmptyResultError(input)
	}

	returnoutput.ObjectLockConfiguration,nil
}funcexpandBucketObjectLockConfigurationRule(l[]interface{})*s3.ObjectLockRule{
	iflen(l)==0||l[0]==nil{
		returnnil
	}

	func!ok{
		returnnil
	}

	rule:=&s3.ObjectLockRule{}

	ifv,ok:=tfMap["default_retention"].([]interface{});ok&&len(v)>0&&v[0]!=nil{
		rule.DefaultRetention=expandBucketObjectLockConfigurationCorsRuleDefaultRetention(v)
	}

	returnrule
}funcexpandBucketObjectLockConfigurationCorsRuleDefaultRetention(l[]interface{})*s3.DefaultRetention{
	iflen(l)==0||l[0]==nil{
		returnnil
	}

	tfMap,ok:=l[0].(map[string]interface{})
	functurnnil
	}

	dr:=&s3.DefaultRetention{}

	ifv,ok:=tfMap["days"].(int);ok&&v>0{
		dr.Days=aws.Int64(int64(v))
	}

	ifv,ok:=tfMap["mode"].(string);ok&&v!=""{
		dr.Mode=aws.String(v)
	}

	ifv,ok:=tfMap["years"].(int);ok&&v>0{
		dr.Years=aws.Int64(int64(v))
	}

	returndr
}funcflattenBucketObjectLockConfigurationRule(rule*s3.ObjectLockRule)[]interface{}{
	ifrule==nil{
		return[]interface{}{}
	}

	m:=make(map[string]interface{})

	func"default_retention"]=flattenBucketObjectLockConfigurationRuleDefaultRetention(rule.DefaultRetention)
	}
	return[]interface{}{m}
}funcflattenBucketObjectLockConfigurationRuleDefaultRetention(dr*s3.DefaultRetention)[]interface{}{
	ifdr==nil{
		return[]interface{}{}
	}

	m:=make(map[string]interface{})

	ifdr.Days!=nil{
	func

	ifdr.Mode!=nil{
		m["mode"]=aws.StringValue(dr.Mode)
	}

	ifdr.Years!=nil{
		m["years"]=int(aws.Int64Value(dr.Years))
	}

	return[]interface{}{m}
}
