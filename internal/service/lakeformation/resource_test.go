//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagelakeformation_test

import(
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lakeformation"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tflakeformation"github.com/hashicorp/terraform-provider-aws/internal/service/lakeformation"
)

funcTestAccLakeFormationResource_basic(t*testing.T){
	ctx:=acctest.Context(t)
	bucketName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	roleName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceAddr:="aws_lakeformation_resource.test"
	bucketAddr:="aws_s3_bucket.test"
	roleAddr:="aws_iam_role.test"

	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){acctest.PreCheck(ctx,t);acctest.PreCheckPartitionHasService(t,lakeformation.EndpointsID)},
		ErrorCheck:acctest.ErrorCheck(t,lakeformation.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckResourceDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccResourceConfig_basic(bucketName,roleName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(ctx,resourceAddr),
					resource.TestCheckResourceAttrPair(resourceAddr,"role_arn",roleAddr,"arn"),
					resource.TestCheckResourceAttrPair(resourceAddr,"arn",bucketAddr,"arn"),
				),
			},
		},
	})
}

funcTestAccLakeFormationResource_disappears(t*testing.T){
	ctx:=acctest.Context(t)
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName:="aws_lakeformation_resource.test"

	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){acctest.PreCheck(ctx,t);acctest.PreCheckPartitionHasService(t,lakeformation.EndpointsID)},
		ErrorCheck:acctest.ErrorCheck(t,lakeformation.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckResourceDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccResourceConfig_basic(rName,rName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(ctx,resourceName),
					acctest.CheckResourceDisappears(ctx,acctest.Provider,tflakeformation.ResourceResource(),resourceName),
				),
				ExpectNonEmptyPlan:true,
			},
		},
	})
}

funcTestAccLakeFormationResource_serviceLinkedRole(t*testing.T){
	ctx:=acctest.Context(t)
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceAddr:="aws_lakeformation_resource.test"
	bucketAddr:="aws_s3_bucket.test"

	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){
			acctest.PreCheck(ctx,t)
			acctest.PreCheckPartitionHasService(t,lakeformation.EndpointsID)
			acctest.PreCheckIAMServiceLinkedRole(ctx,t,"/aws-service-role/lakeformation.amazonaws.com")
		},
		ErrorCheck:acctest.ErrorCheck(t,lakeformation.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckResourceDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccResourceConfig_serviceLinkedRole(rName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(ctx,resourceAddr),
					resource.TestCheckResourceAttrPair(resourceAddr,"arn",bucketAddr,"arn"),
					acctest.CheckResourceAttrGlobalARN(resourceAddr,"role_arn","iam","role/aws-service-role/lakeformation.amazonaws.com/AWSServiceRoleForLakeFormationDataAccess"),
				),
			},
		},
	})
}

funcTestAccLakeFormationResource_updateRoleToRole(t*testing.T){
	ctx:=acctest.Context(t)
	bucketName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	roleName1:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	roleName2:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceAddr:="aws_lakeformation_resource.test"
	bucketAddr:="aws_s3_bucket.test"
	roleAddr:="aws_iam_role.test"

	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){acctest.PreCheck(ctx,t);acctest.PreCheckPartitionHasService(t,lakeformation.EndpointsID)},
		ErrorCheck:acctest.ErrorCheck(t,lakeformation.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckResourceDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccResourceConfig_basic(bucketName,roleName1),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(ctx,resourceAddr),
					resource.TestCheckResourceAttrPair(resourceAddr,"role_arn",roleAddr,"arn"),
					resource.TestCheckResourceAttrPair(resourceAddr,"arn",bucketAddr,"arn"),
				),
			},
			{
				Config:testAccResourceConfig_basic(bucketName,roleName2),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(ctx,resourceAddr),
					resource.TestCheckResourceAttrPair(resourceAddr,"role_arn",roleAddr,"arn"),
					resource.TestCheckResourceAttrPair(resourceAddr,"arn",bucketAddr,"arn"),
				),
			},
		},
	})
}

funcTestAccLakeFormationResource_updateSLRToRole(t*testing.T){
	ctx:=acctest.Context(t)
	bucketName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	roleName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceAddr:="aws_lakeformation_resource.test"
	bucketAddr:="aws_s3_bucket.test"
	roleAddr:="aws_iam_role.test"

	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){
			acctest.PreCheck(ctx,t)
			acctest.PreCheckPartitionHasService(t,lakeformation.EndpointsID)
			acctest.PreCheckIAMServiceLinkedRole(ctx,t,"/aws-service-role/lakeformation.amazonaws.com")
		},
		ErrorCheck:acctest.ErrorCheck(t,lakeformation.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckResourceDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccResourceConfig_serviceLinkedRole(bucketName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(ctx,resourceAddr),
					resource.TestCheckResourceAttrPair(resourceAddr,"arn",bucketAddr,"arn"),
					acctest.CheckResourceAttrGlobalARN(resourceAddr,"role_arn","iam","role/aws-service-role/lakeformation.amazonaws.com/AWSServiceRoleForLakeFormationDataAccess"),
				),
			},
			{
				Config:testAccResourceConfig_basic(bucketName,roleName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckResourceExists(ctx,resourceAddr),
					resource.TestCheckResourceAttrPair(resourceAddr,"role_arn",roleAddr,"arn"),
					resource.TestCheckResourceAttrPair(resourceAddr,"arn",bucketAddr,"arn"),
				),
			},
		},
	})
}

//AWSdoesnotsupportchangingfromanIAMroletoanSLR.Noerroristhrown
//buttheregistrationisnotchanged(theIAMrolecontinuesintheregistration).
//
//funcTestAccLakeFormationResource_updateRoleToSLR(t*testing.T){

functestAccCheckResourceDestroy(ctxcontext.Context)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
		conn:=acctest.Provider.Meta().(*conns.AWSClient).LakeFormationConn(ctx)

		for_,rs:=ranges.RootModule().Resources{
			ifrs.Type!="aws_lakeformation_resource"{
				continue
			}

			resourceArn:=rs.Primary.Attributes["arn"]

			input:=&lakeformation.DescribeResourceInput{
				ResourceArn:aws.String(resourceArn),
			}

			_,err:=conn.DescribeResourceWithContext(ctx,input)
			iferr==nil{
				returnfmt.Errorf("resourcestillregistered:%s",resourceArn)
			}
			if!isResourceNotFoundErr(err){
				returnerr
			}
		}

		returnnil
	}
}

functestAccCheckResourceExists(ctxcontext.Context,resourceNamestring)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
		rs,ok:=s.RootModule().Resources[resourceName]
		if!ok{
			returnfmt.Errorf("resourcenotfound:%s",resourceName)
		}

		conn:=acctest.Provider.Meta().(*conns.AWSClient).LakeFormationConn(ctx)

		input:=&lakeformation.DescribeResourceInput{
			ResourceArn:aws.String(rs.Primary.ID),
		}

		_,err:=conn.DescribeResourceWithContext(ctx,input)

		iferr!=nil{
			returnfmt.Errorf("errorgettingLakeFormationresource(%s):%w",rs.Primary.ID,err)
		}

		returnnil
	}
}

funcisResourceNotFoundErr(errerror)bool{
	returntfawserr.ErrMessageContains(
		err,
		"EntityNotFoundException",
		"Entitynotfound")
}

functestAccResourceConfig_basic(bucket,rolestring)string{
	returnfmt.Sprintf(`
resource"aws_s3_bucket""test"{
bucket=%[1]q
}

resource"aws_iam_role""test"{
name=%[2]q
path="/test/"

assume_role_policy=<<EOF
{
"Version":"2012-10-17",
"Statement":[
{
"Action":"sts:AssumeRole",
"Principal":{
"Service":"s3.amazonaws.com"
},
"Effect":"Allow",
"Sid":""
}
]
}
EOF
}

data"aws_partition""current"{}

resource"aws_iam_role_policy""test"{
name=%[2]q
role=aws_iam_role.test.id

policy=<<EOF
{
"Version":"2012-10-17",
"Statement":[
{
"Effect":"Allow",
"Action":[
"s3:GetBucketLocation",
"s3:ListAllMyBuckets",
"s3:GetObjectVersion",
"s3:GetBucketAcl",
"s3:GetObject",
"s3:GetObjectACL",
"s3:PutObject",
"s3:PutObjectAcl"
],
"Resource":[
"arn:${data.aws_partition.current.partition}:s3:::${aws_s3_bucket.test.id}/*",
"arn:${data.aws_partition.current.partition}:s3:::${aws_s3_bucket.test.id}"
]
}
]
}
EOF
}

resource"aws_lakeformation_resource""test"{
arn=aws_s3_bucket.test.arn
role_arn=aws_iam_role.test.arn
}
`,bucket,role)
}

functestAccResourceConfig_serviceLinkedRole(rNamestring)string{
	returnfmt.Sprintf(`
resource"aws_s3_bucket""test"{
bucket=%[1]q
}

resource"aws_lakeformation_resource""test"{
arn=aws_s3_bucket.test.arn
}
`,rName)
}
