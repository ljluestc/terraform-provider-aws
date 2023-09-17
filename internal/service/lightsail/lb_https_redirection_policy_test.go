// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package lightsail_testimport (
	"fmt"
	"strings"
	"testing"	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)
func testAccLoadBalancerHTTPSRedirectionPolicy_basic(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix("tf-acc-test")
	enabled := "true"	resource.Test(t, resource.TestCase{
PreCheck: func() {
acctest.PreCheck(ctx, t)
acctest.PreCheckPartitionHasService(t, strings.ToLower(lightsail.ServiceID))
testAccPreCheck(ctx, t)
},
ErrorCheck:orCheck(t, strings.ToLower(lightsail.ServiceID)),
ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
CheckDestroy:oadBalancerDestroy(ctx),
Steps: []resource.TestStep{
{
Config: testAccLoadBalancerHTTPSRedirectionPolicyConfig_basic(rName, enabled),
ExpectError: regexache.MustCompile(`cannot enable https redirection while https is disabled.`),
},
},
	})
}
func testAccLoadBalancerHTTPSRedirectionPolicyConfig_basic(rName string, enabled string) string {
	return fmt.Sprintf(`
resource "aws_lightsail_lb" "test" {
  name
  health_check_path = "/"
  instance_port= "80"
}
resource "aws_lightsail_lb_https_redirection_policy" "test" {
  enabled = %[2]s
  lb_name = aws_lightsail_lb.test.name
}
`, rName, enabled)
}
