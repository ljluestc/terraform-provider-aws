//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagekms_test

import(
	"context"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)
funcTestAccKMSSecretsDataSource_basic(t*testing.T){
	ctx:=acctest.Context(t)
	varencryptedPayloadstring
	varkeykms.KeyMetadata

	plaintext:="my-plaintext-string"
	resourceName:="aws_kms_key.test"

	//RunaresourcetesttosetupourKMSkey
	resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,kms.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
Steps:[]resource.TestStep{
	{
Config:testAccSecretsDataSourceConfig_key,
Check:resource.ComposeTestCheckFunc(
	testAccCheckKeyExists(ctx,resourceName,&key),
	testAccSecretsEncryptDataSource(ctx,&key,plaintext,&encryptedPayload),
	//WeneedtodereferencetheencryptedPayloadinatestTerraformconfiguration
	testAccSecretsDecryptDataSource(ctx,t,plaintext,&encryptedPayload),
),
	},
},
	})
}
funcTestAccKMSSecretsDataSource_asymmetric(t*testing.T){
	ctx:=acctest.Context(t)
	varencryptedPayloadstring
	varkeykms.KeyMetadata

	plaintext:="my-plaintext-string"
	resourceName:="aws_kms_key.test"

	//RunaresourcetesttosetupourKMSkey
	resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,kms.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
Steps:[]resource.TestStep{
	{
Config:testAccSecretsDataSourceConfig_asymmetricKey,
Check:resource.ComposeTestCheckFunc(
	testAccCheckKeyExists(ctx,resourceName,&key),
	testAccSecretsEncryptDataSourceAsymmetric(ctx,&key,plaintext,&encryptedPayload),
	//WeneedtodereferencetheencryptedPayloadinatestTerraformconfiguration
	testAccSecretsDecryptDataSourceAsym(ctx,t,&key,plaintext,&encryptedPayload),
),
	},
},
	})
}
functestAccSecretsEncryptDataSource(ctxcontext.Context,key*kms.KeyMetadata,plaintextstring,encryptedPayload*string)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
conn:=acctest.Provider.Meta().(*conns.AWSClient).KMSConn(ctx)

input:=&kms.EncryptInput{
	KeyId:key.Arn,
	Plaintext:[]byte(plaintext),
	EncryptionContext:map[string]*string{
"name":aws.String("value"),
	},
}

output,err:=conn.EncryptWithContext(ctx,input)

iferr!=nil{
	returnerr
}

*encryptedPayload=base64.StdEncoding.EncodeToString(output.CiphertextBlob)

returnnil
	}
}
functestAccSecretsEncryptDataSourceAsymmetric(ctxcontext.Context,key*kms.KeyMetadata,plaintextstring,encryptedPayload*string)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
conn:=acctest.Provider.Meta().(*conns.AWSClient).KMSConn(ctx)

input:=&kms.EncryptInput{
	KeyId:key.Arn,
	Plaintext:[]byte(plaintext),
	EncryptionAlgorithm:aws.String("RSAES_OAEP_SHA_1"),
}

output,err:=conn.EncryptWithContext(ctx,input)

iferr!=nil{
	returnerr
}

*encryptedPayload=base64.StdEncoding.EncodeToString(output.CiphertextBlob)

returnnil
	}
}
functestAccSecretsDecryptDataSource(ctxcontext.Context,t*testing.T,plaintextstring,encryptedPayload*string)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
dataSourceName:="data.aws_kms_secrets.test"

resource.Test(t,resource.TestCase{
	PreCheck:func(){acctest.PreCheck(ctx,t)},
	ErrorCheck:acctest.ErrorCheck(t,kms.EndpointsID),
	ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
	Steps:[]resource.TestStep{
{
	Config:testAccSecretsDataSourceConfig_secret(*encryptedPayload),
	Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(dataSourceName,"plaintext.%","1"),
resource.TestCheckResourceAttr(dataSourceName,"plaintext.secret1",plaintext),
	),
},
	},
})

returnnil
	}
}
functestAccSecretsDecryptDataSourceAsym(ctxcontext.Context,t*testing.T,key*kms.KeyMetadata,plaintextstring,encryptedPayload*string)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
dataSourceName:="data.aws_kms_secrets.test"
keyid:=key.Arn

resource.Test(t,resource.TestCase{
	PreCheck:func(){acctest.PreCheck(ctx,t)},
	ErrorCheck:acctest.ErrorCheck(t,kms.EndpointsID),
	ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
	Steps:[]resource.TestStep{
{
	Config:testAccSecretsDataSourceConfig_asymmetricSecret(*encryptedPayload,*keyid),
	Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttr(dataSourceName,"plaintext.%","1"),
resource.TestCheckResourceAttr(dataSourceName,"plaintext.secret1",plaintext),
	),
},
	},
})

returnnil
	}
}

consttestAccSecretsDataSourceConfig_key=`
resource"aws_kms_key""test"{
deletion_window_in_days=7
description="TestingtheTerraformAWSKMSSecretsdata_source"
}
`
functestAccSecretsDataSourceConfig_secret(payloadstring)string{
	returnacctest.ConfigCompose(testAccSecretsDataSourceConfig_key,fmt.Sprintf(`
data"aws_kms_secrets""test"{
secret{
name="secret1"
payload=%[1]q

context={
name="value"
}
}
}
`,payload))
}

consttestAccSecretsDataSourceConfig_asymmetricKey=`
resource"aws_kms_key""test"{
deletion_window_in_days=7
description="TestingtheTerraformAWSKMSSecretsdata_source"
customer_master_key_spec="RSA_2048"
}
`
functestAccSecretsDataSourceConfig_asymmetricSecret(payloadstring,keyidstring)string{
	returnacctest.ConfigCompose(testAccSecretsDataSourceConfig_asymmetricKey,fmt.Sprintf(`
data"aws_kms_secrets""test"{
secret{
name="secret1"
payload=%[1]q
encryption_algorithm="RSAES_OAEP_SHA_1"
key_id=%[2]q
}
}
`,payload,keyid))
}
