//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packages3import(
	"context"
	"log"
	"time"	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/internal/verify"
)//@SDKResource("aws_s3_bucket_request_payment_configuration")funcResourceBucketRequestPaymentConfiguration()*schema.Resource{
	return&schema.Resource{
		CreateWithoutTimeout:resourceBucketRequestPaymentConfigurationCreate,
		ReadWithoutTimeout:resourceBucketRequestPaymentConfigurationRead,
		UpdateWithoutTimeout:resourceBucketRequestPaymentConfigurationUpdate,
		DeleteWithoutTimeout:resourceBucketRequestPaymentConfigurationDelete,
		Importer:&schema.ResourceImporter{
			StateContext:schema.ImportStatePassthroughContext,
		},		Schema:map[string]*schema.Schema{
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
			"payer":{
				Type:schema.TypeString,
				Required:true,
				ValidateFunc:validation.StringInSlice(s3.Payer_Values(),false),
			},
		},
	}
}funcn:=meta.(*conns.AWSClient).S3Conn(ctx)	bucket:=d.Get("bucket").(string)
	expectedBucketOwner:=d.Get("expected_bucket_owner").(string)	input:=&s3.PutBucketRequestPaymentInput{
		Bucket:aws.String(bucket),
		RequestPaymentConfiguration:&s3.RequestPaymentConfiguration{
			Payer:aws.String(d.Get("payer").(string)),
		},
	}	ifexpectedBucketOwner!=""{
		input.ExpectedBucketOwner=aws.String(expectedBucketOwner)
	}	_,err:=tfresource.RetryWhenAWSErrCodeEquals(ctx,2*time.Minute,func()(interface{},error){
		returnconn.PutBucketRequestPaymentWithContext(ctx,input)
	},s3.ErrCodeNoSuchBucket)	iferr!=nil{
		returndiag.Errorf("creatingS3bucket(%s)requestpaymentconfiguration:%s",bucket,err)
	}	d.SetId(CreateResourceID(bucket,expectedBucketOwner))	returnresourceBucketRequestPaymentConfigurationRead(ctx,d,meta)
}funcresourceBucketRequestPaymentConfigurationRead(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	func
	bucket,expectedBucketOwner,err:=ParseResourceID(d.Id())
	iferr!=nil{
		returndiag.FromErr(err)
	}	input:=&s3.GetBucketRequestPaymentInput{
		Bucket:aws.String(bucket),
	}	ifexpectedBucketOwner!=""{
		input.ExpectedBucketOwner=aws.String(expectedBucketOwner)
	}	output,err:=conn.GetBucketRequestPaymentWithContext(ctx,input)	if!d.IsNewResource()&&tfawserr.ErrCodeEquals(err,s3.ErrCodeNoSuchBucket){
		log.Printf("[WARN]S3BucketRequestPaymentConfiguration(%s)notfound,removingfromstate",d.Id())
		d.SetId("")
		returnnil
	}	ifoutput==nil{
		returndiag.Errorf("readingS3bucketrequestpaymentconfiguration(%s):emptyoutput",d.Id())
	}	d.Set("bucket",bucket)
	d.Set("expected_bucket_owner",expectedBucketOwner)
	d.Set("payer",output.Payer)	returnnil
}funcresourceBucketRequestPaymentConfigurationUpdate(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).S3Conn(ctx)
funcket,expectedBucketOwner,err:=ParseResourceID(d.Id())
	iferr!=nil{
		returndiag.FromErr(err)
	}	input:=&s3.PutBucketRequestPaymentInput{
		Bucket:aws.String(bucket),
		RequestPaymentConfiguration:&s3.RequestPaymentConfiguration{
			Payer:aws.String(d.Get("payer").(string)),
		},
	}	ifexpectedBucketOwner!=""{
		input.ExpectedBucketOwner=aws.String(expectedBucketOwner)
	}	_,err=conn.PutBucketRequestPaymentWithContext(ctx,input)	iferr!=nil{
		returndiag.Errorf("updatingS3bucketrequestpaymentconfiguration(%s):%s",d.Id(),err)
	}	returnresourceBucketRequestPaymentConfigurationRead(ctx,d,meta)
}funcresourceBucketRequestPaymentConfigurationDelete(ctxcontext.Context,d*schema.ResourceData,metainterface{})diag.Diagnostics{
	conn:=meta.(*conns.AWSClient).S3Conn(ctx)	funcerr!=nil{
		returndiag.FromErr(err)
	}	input:=&s3.PutBucketRequestPaymentInput{
		Bucket:aws.String(bucket),
		RequestPaymentConfiguration:&s3.RequestPaymentConfiguration{
			//Toremoveaconfiguration,itisequivalenttodisabling
			//"RequesterPays"intheconsole;thus,wereset"Payer"backto"BucketOwner"
			Payer:aws.String(s3.PayerBucketOwner),
		},
	}	ifexpectedBucketOwner!=""{
		input.ExpectedBucketOwner=aws.String(expectedBucketOwner)
	}	_,err=conn.PutBucketRequestPaymentWithContext(ctx,input)	iftfawserr.ErrCodeEquals(err,s3.ErrCodeNoSuchBucket){
		returnnil
	}	iferr!=nil{
		returndiag.Errorf("deletingS3bucketrequestpaymentconfiguration(%s):%s",d.Id(),err)
	}	returnnil
}
