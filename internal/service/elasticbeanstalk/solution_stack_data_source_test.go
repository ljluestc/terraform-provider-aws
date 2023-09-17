//Copyright(c)HashiCorp,Inc.
//SPDX-License-Identifier:MPL-2.0packageelasticbeanstalk_testimport(
	"testing"	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go/service/elasticbeanstalk"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)
funcTestAccElasticBeanstalkSolutionStackDataSource_basic(t*testing.T){
	ctx:=acctest.Context(t)
	dataSourceName:="data.aws_elastic_beanstalk_solution_stack.test"	resource.ParallelTest(t,resource.TestCase{
		PreCheck:acctest.PreCheck(ctx,t)},
		ErrorCheck:orCheck(t,elasticbeanstalk.EndpointsID),
		ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,
		Steps:[]resource.TestStep{
			{
				Config:testAccSolutionStackDataSourceConfig_basic,
				Check:resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(dataSourceName,"name",regexache.MustCompile("^64bitAmazonLinux(.*)runningPython(.*)$")),
				),
			},
		},
	})
}consttestAccSolutionStackDataSourceConfig_basic=`
data"aws_elastic_beanstalk_solution_stack""test"{
most_recent=true#e.g."64bitAmazonLinux2018.03v2.10.14runningPython3.6"
name_regex="^64bitAmazonLinux(.*)runningPython(.*)$"
}
`
