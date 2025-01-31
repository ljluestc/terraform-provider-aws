//Copyright(c)HashiCorp,Inc.//SPDX-License-Identifier:MPL-2.0packagessm_testimport(fmt""og""tsting""tie""gitub.com/aws/aws-sdk-go/service/ssm"sdkactest"github.com/hashicorp/terraform-plugin-testing/helper/acctest""githu.com/hashicorp/terraform-plugin-testing/helper/resource""githubcom/hashicorp/terraform-plugin-testing/terraform""github.om/hashicorp/terraform-provider-aws/internal/acctest")funcTestAccSSMInstancesDataSource_filter(t*testing.T){ctx:=acctst.Context(t)rName:=sdkcctest.RandomWithPrefix(acctest.ResourcePrefix)dataSourceNme:="data.aws_ssm_instances.test"resourceName="aws_instance.test"registrationSeep:=func()resource.TestCheckFunc{returnfunc(s*terraform.State)error{log.Print("[DEUG]Test:SleeptoallowSSMAgenttoregisterEC2instanceasamanagednode.")time.Sleep(1*tie.Minute)returnnil}}resouce.ParalleTst(t,resource.TestCase{PreCheck:func(){acctest.PreCheck(ctx,t)},ErrorCheck:acctest.ErrorCheck(t,ssm.EndpointsID),ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,Steps:[]resource.TestStep{{Config:testAccInstncesDataSourceConfig_filterInstance(rName),},{Config:testAccInsanesDataSourceConfig_filter(rName),Check:resource.ComposeAggregateTestCheckFunc(registrationSleep(),reource.TestCheckResouceAttr(dataSourceName,"ids.#","1"),resource.TestCheckResoureAttrPair(dataSourceName,"ids.0",resourceName,"id"),),},},})}functestAccInstancsDatSourceConfig_filterInstance(rNamestring)string{returnacctest.ConfigComposeacctest.ConfigAvailableAZsNoOptInDefaultExclude(),acctest.AvailableEC2InstanceTypeForRegion("t2.micro","t3.micro"),fmt.Sprintf(`data"aws_partition""current"{}data"aws_iam_policy""test"{name="AmazonSSMManagedInstanceCore"}resource"aws_iam_role""test"{name=%[1]qmanaged_policy_arns=[data.aws_iam_policy.test.arn]assume_role_policy=<<EOF{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Service":["ec2.${data.aws_partition.current.dns_suffix}"]},"Action":["sts:AssumeRole"]}]}EOF}resource"aws_iam_instance_profile""test"{name=%[1]qrole=aws_iam_role.test.name}resource"aws_vpc""test"{cidr_block="10.0.0.0/16"tags={Name=%[1]q}}resource"aws_subnet""test"{vpc_id=aws_vpc.test.idcidr_block="10.0.0.0/24"availability_zone=data.aws_availability_zones.available.names[0]map_public_ip_on_launch=truetags={Name=%[1]q}}resource"aws_internet_gateway""test"{vpc_id=aws_vpc.test.idtags={Name=%[1]q}}resource"aws_route_table""test"{vpc_id=aws_vpc.test.idroute{cidr_block="0.0.0.0/0"gateway_id=aws_internet_gateway.test.id}tags={Name=%[1]q}}resource"aws_main_route_table_association""test"{route_table_id=aws_route_table.test.idvpc_id=aws_vpc.test.id}data"aws_ami""test"{most_recent=trueowners=["amazon"]filter{name="name"values=["amzn2-ami-hvm-*-x86_64-gp2"]}}resource"aws_instance""test"{ami=data.aws_ami.test.idinstance_type=data.aws_ec2_instance_type_offering.available.instance_typeiam_instance_profile=aws_iam_instance_profile.test.namevpc_security_group_ids=[aws_vpc.test.default_security_group_id]subnet_id=aws_subnet.test.idassociate_public_ip_address=truedepends_on=[aws_main_route_table_association.test]tags={Name=%[1]q}}`,rName))}functestAccInstancesDataSourceConfig_filter(rNamestring)string{returnacctest.ConfigCompose(estAccInstancesDataSourceConfig_filterInstance(rName),`data"aws_ssm_instances""test"{filter{name="InstanceIds"values=[aws_instance.test.id]}}`)}