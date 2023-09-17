//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packagecognitoidp_testimport(
	"context"
	"errors"
	"fmt"
	"testing"	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	sdkacctest"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfcognitoidp"github.com/hashicorp/terraform-provider-aws/internal/service/cognitoidp"
)funcTestAccCognitoIDPUserInGroup_basic(t*testing.T){
	ctx:=acctest.Context(t)
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName:="aws_cognito_user_in_group.test"
	userPoolResourceName:="aws_cognito_user_pool.test"
	userGroupResourceName:="aws_cognito_user_group.test"
	userResourceName:="aws_cognito_user.test"	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){acctest.PreCheck(ctx,t)},
		ErrorCheck:acctest.ErrorCheck(t,cognitoidentityprovider.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckUserInGroupDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccUserInGroupConfig_basic(rName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckUserInGroupExists(ctx,resourceName),
					resource.TestCheckResourceAttrPair(resourceName,"user_pool_id",userPoolResourceName,"id"),
					resource.TestCheckResourceAttrPair(resourceName,"group_name",userGroupResourceName,"name"),
					resource.TestCheckResourceAttrPair(resourceName,"username",userResourceName,"username"),
				),
			},
		},
	})
}funcTestAccCognitoIDPUserInGroup_disappears(t*testing.T){
	ctx:=acctest.Context(t)
	rName:=sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	resourceName:="aws_cognito_user_in_group.test"	resource.ParallelTest(t,resource.TestCase{
		PreCheck:func(){acctest.PreCheck(ctx,t)},
		ErrorCheck:acctest.ErrorCheck(t,cognitoidentityprovider.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		CheckDestroy:testAccCheckUserInGroupDestroy(ctx),
		Steps:[]resource.TestStep{
			{
				Config:testAccUserInGroupConfig_basic(rName),
				Check:resource.ComposeTestCheckFunc(
					testAccCheckUserInGroupExists(ctx,resourceName),
					acctest.CheckResourceDisappears(ctx,acctest.Provider,tfcognitoidp.ResourceUserInGroup(),resourceName),
				),
				ExpectNonEmptyPlan:true,
			},
		},
	})
}functestAccUserInGroupConfig_basic(rNamestring)string{
	returnfmt.Sprintf(`
resource"aws_cognito_user_pool""test"{
name=%[1]q
password_policy{
temporary_password_validity_days=7
minimum_length=6
require_uppercase=false
require_symbols=false
require_numbers=false
}
}resource"aws_cognito_user""test"{
user_pool_id=aws_cognito_user_pool.test.id
username=%[1]q
}resource"aws_cognito_user_group""test"{
user_pool_id=aws_cognito_user_pool.test.id
name=%[1]q
}resource"aws_cognito_user_in_group""test"{
user_pool_id=aws_cognito_user_pool.test.id
group_name=aws_cognito_user_group.test.name
username=aws_cognito_user.test.username
}
`,rName)
}functestAccCheckUserInGroupExists(ctxcontext.Context,resourceNamestring)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
		rs,ok:=s.RootModule().Resources[resourceName]
		if!ok{
			returnfmt.Errorf("resourcenotfound:%s",resourceName)
		}		conn:=acctest.Provider.Meta().(*conns.AWSClient).CognitoIDPConn(ctx)		groupName:=rs.Primary.Attributes["group_name"]
		userPoolId:=rs.Primary.Attributes["user_pool_id"]
		username:=rs.Primary.Attributes["username"]		found,err:=tfcognitoidp.FindCognitoUserInGroup(ctx,conn,groupName,userPoolId,username)		iferr!=nil{
			returnerr
		}		if!found{
			returnerrors.New("useringroupnotfound")
		}		returnnil
	}
}functestAccCheckUserInGroupDestroy(ctxcontext.Context)resource.TestCheckFunc{
	returnfunc(s*terraform.State)error{
		conn:=acctest.Provider.Meta().(*conns.AWSClient).CognitoIDPConn(ctx)		for_,rs:=ranges.RootModule().Resources{
			ifrs.Type!="aws_cognito_user_in_group"{
				continue
			}			groupName:=rs.Primary.Attributes["group_name"]
			userPoolId:=rs.Primary.Attributes["user_pool_id"]
			username:=rs.Primary.Attributes["username"]			found,err:=tfcognitoidp.FindCognitoUserInGroup(ctx,conn,groupName,userPoolId,username)			iftfawserr.ErrCodeEquals(err,cognitoidentityprovider.ErrCodeResourceNotFoundException){
				continue
			}			iferr!=nil{
				returnerr
			}			iffound{
				returnfmt.Errorf("useringroupstillexists(%s)",rs.Primary.ID)
			}
		}		returnnil
	}
}
