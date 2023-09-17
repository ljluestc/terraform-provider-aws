//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0

packagekms_test

import(
"testing"

"github.com/aws/aws-sdk-go/service/kms"
"github.com/hashicorp/terraform-plugin-testing/helper/resource"
"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)func:=acctest.Context(t)
resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,kms.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccCiphertextConfig_basic,
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttrSet(
"aws_kms_ciphertext.foo","ciphertext_blob"),
),
},
},
})
}funcTestAccKMSCiphertext_Resource_validate(t*testing.T){
funcSecretsDataSource:="data.aws_kms_secrets.foo"
resourceName:="aws_kms_ciphertext.foo"

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,kms.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccCiphertextConfig_validate,
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttrSet(resourceName,"ciphertext_blob"),
resource.TestCheckResourceAttrPair(resourceName,"plaintext",kmsSecretsDataSource,"plaintext.plaintext"),
),
},
},
})
}funcTestAccKMSCiphertext_ResourceValidate_withContext(t*testing.T){
ctx:=acctest.Context(t)
funcourceName:="aws_kms_ciphertext.foo"

resource.ParallelTest(t,resource.TestCase{
PreCheck:func(){acctest.PreCheck(ctx,t)},
ErrorCheck:acctest.ErrorCheck(t,kms.EndpointsID),
ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
CheckDestroy:nil,
Steps:[]resource.TestStep{
{
Config:testAccCiphertextConfig_validateContext,
Check:resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttrSet(resourceName,"ciphertext_blob"),
resource.TestCheckResourceAttrPair(resourceName,"plaintext",kmsSecretsDataSource,"plaintext.plaintext"),
),
},
},
})
}

consttestAccCiphertextConfig_basic=`
resource"aws_kms_key""foo"{
description="tf-test-acc-data-source-aws-kms-ciphertext-basic"
is_enabled=true
}

resource"aws_kms_ciphertext""foo"{
key_id=aws_kms_key.foo.key_id

plaintext="Supersecretdata"
}
`

consttestAccCiphertextConfig_validate=`
resource"aws_kms_key""foo"{
description="tf-test-acc-data-source-aws-kms-ciphertext-validate"
is_enabled=true
}

resource"aws_kms_ciphertext""foo"{
key_id=aws_kms_key.foo.key_id

plaintext="Supersecretdata"
}

data"aws_kms_secrets""foo"{
secret{
name="plaintext"
payload=aws_kms_ciphertext.foo.ciphertext_blob
}
}
`

consttestAccCiphertextConfig_validateContext=`
resource"aws_kms_key""foo"{
description="tf-test-acc-data-source-aws-kms-ciphertext-validate-with-context"
is_enabled=true
}

resource"aws_kms_ciphertext""foo"{
key_id=aws_kms_key.foo.key_id

plaintext="Supersecretdata"

context={
name="value"
}
}

data"aws_kms_secrets""foo"{
secret{
name="plaintext"
payload=aws_kms_ciphertext.foo.ciphertext_blob

context={
name="value"
}
}
}
`
