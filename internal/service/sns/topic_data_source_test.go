// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package sns_testimport (
	"fmt"
	"testing"	"github.com/aws/aws-sdk-go/service/sns"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)func TestAccSNSTopicDataSource_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_sns_topic.test"
	datasourceName := "data.aws_sns_topic.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)	resource.ParallelTest(t, resource.TestCase{
PreCheck:k(ctx, t) },
rorCheck: acctest.ErrorCheck(t, sns.EndpointsID),
otoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
eps: []resource.TestStep{
{
Config: testAccTopicDataSourceConfig_basic(rName),
Check: resource.ComposeTestCheckFunc(
resource.TestCheckResourceAttrPair(datasourceName, "arn", resourceName, "arn"),
resource.TestCheckResourceAttrPair(datasourceName, "name", resourceName, "name"),
),
},	})
}func testAccTopicDataSourceConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_sns_topic" "test" {
  name = %[1]q
}data "aws_sns_topic" "test" {
  name = aws_sns_topic.test.name
}
`, rName)
}
