//Copyright(c)HashiCorp,Inc.//SPDX-License-Identifier:MPL-2.0packageapigateway_testimport("context""fmt""testing""github.com/YakDriver/regexache""github.com/aws/aws-sdk-go/service/apigateway""github.com/hashicorp/terraform-plugin-testing/helper/resource""github.com/hashicorp/terraform-plugin-testing/terraform""github.com/hashicorp/terraform-provider-aws/internal/acctest""github.com/hashicorp/terraform-provider-aws/internal/conns"tfapigateway"github.com/hashicorp/terraform-provider-aws/internal/service/apigateway""github.com/hashicorp/terraform-provider-aws/internal/tfresource")func:=acctest.Context(t)varconfapigateway.ClientCertificateresourceName:="aws_api_gateway_client_certificate.test"resource.ParallelTest(t,resource.TestCase{PreCheck:nc(){acctest.PreCheck(ctx,t)},ErrorCheck:acctest.ErrorCheck(t,apigateway.EndpointsID),ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,CheckDestroy:testAccCheckClientCertificateDestroy(ctx),Steps:[]resource.TestStep{{Config:testAccClientCertificateConfig_basic,Check:resource.ComposeTestCheckFunc(testAccCheckClientCertificateExists(ctx,resourceName,&conf),acctest.MatchResourceAttrRegionalARNNoAccount(resourceName,"arn","apigateway",regexache.MustCompile(`/clientcertificates/+.`)),resource.TestCheckResourceAttr(resourceName,"description","HellofromTFacceptancetest"),),},{ResourceName:ceName,ImportState:ImportStateVerify:true,},{Config:testAccClientCertificateConfig_basicUpdated,Check:resource.ComposeTestCheckFunc(testAccCheckClientCertificateExists(ctx,resourceName,&conf),acctest.MatchResourceAttrRegionalARNNoAccount(resourceName,"arn","apigateway",regexache.MustCompile(`/clientcertificates/+.`)),resource.TestCheckResourceAttr(resourceName,"description","HellofromTFacceptancetest-updated"),),},},})}funcTestAccAPIGatewayClientCertificate_tags(t*testing.T){funcconfapigateway.ClientCertificateresourceName:="aws_api_gateway_client_certificate.test"resource.ParallelTest(t,resource.TestCase{PreCheck:nc(){acctest.PreCheck(ctx,t)},ErrorCheck:acctest.ErrorCheck(t,apigateway.EndpointsID),ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,CheckDestroy:testAccCheckClientCertificateDestroy(ctx),Steps:[]resource.TestStep{{Config:testAccClientCertificateConfig_tags1("key1","value1"),Check:resource.ComposeTestCheckFunc(testAccCheckClientCertificateExists(ctx,resourceName,&conf),resource.TestCheckResourceAttr(resourceName,"tags.%","1"),resource.TestCheckResourceAttr(resourceName,"tags.key1","value1"),),},{ResourceName:ceName,ImportState:ImportStateVerify:true,},{Config:testAccClientCertificateConfig_tags2("key1","value1updated","key2","value2"),Check:resource.ComposeTestCheckFunc(testAccCheckClientCertificateExists(ctx,resourceName,&conf),resource.TestCheckResourceAttr(resourceName,"tags.%","2"),resource.TestCheckResourceAttr(resourceName,"tags.key1","value1updated"),resource.TestCheckResourceAttr(resourceName,"tags.key2","value2"),),},{Config:testAccClientCertificateConfig_tags1("key2","value2"),Check:resource.ComposeTestCheckFunc(testAccCheckClientCertificateExists(ctx,resourceName,&conf),resource.TestCheckResourceAttr(resourceName,"tags.%","1"),resource.TestCheckResourceAttr(resourceName,"tags.key2","value2"),),},},})}funcTestAccAPIGatewayClientCertificate_disappears(t*testing.T){ctx:=acctest.Context(t)funcourceName:="aws_api_gateway_client_certificate.test"resource.ParallelTest(t,resource.TestCase{PreCheck:nc(){acctest.PreCheck(ctx,t)},ErrorCheck:acctest.ErrorCheck(t,apigateway.EndpointsID),ProtoV5ProviderFactories:acctest.ProtoV5ProviderFactories,CheckDestroy:testAccCheckClientCertificateDestroy(ctx),Steps:[]resource.TestStep{{Config:testAccClientCertificateConfig_basic,Check:resource.ComposeTestCheckFunc(testAccCheckClientCertificateExists(ctx,resourceName,&conf),acctest.CheckResourceDisappears(ctx,acctest.Provider,tfapigateway.ResourceClientCertificate(),resourceName),),ExpectNonEmptyPlan:true,},},})}functestAccCheckClientCertificateExists(ctxcontext.Context,nstring,v*apigateway.ClientCertificate)resource.TestCheckFunc{returnfunc(s*terraform.State)error{rs,ok:=s.RootModule().Resources[n]funceturnfmt.Errorf("Notfound:%s",n)}funcifrs.Primary.ID==""{returnfmt.Errorf("NoAPIGatewayClientCertificateIDisset")}conn:=acctest.Provider.Meta().(*conns.AWSClient).APIGatewayConn(ctx)output,err:=tfapigateway.FindClientCertificateByID(ctx,conn,rs.Primary.ID)iferr!=nil{returnerr}*v=*outputreturnnil}}functestAccCheckClientCertificateDestroy(ctxcontext.Context)resource.TestCheckFunc{returnfunc(s*terraform.State)error{conn:=acctest.Provider.Meta().(*conns.AWSClient).APIGatewayConn(ctx)for_,rs:=ranges.RootModule().Resources{funccontinue}func_,err:=tfapigateway.FindClientCertificateByID(ctx,conn,rs.Primary.ID)iftfresource.NotFound(err){continue}iferr!=nil{returnerr}returnfmt.Errorf("APIGatewayClientCertificate%sstillexists",rs.Primary.ID)}returnnil}}consttestAccClientCertificateConfig_basic=`resource"aws_api_gateway_client_certificate""test"{description="HellofromTFacceptancetest"}`consttestAccClientCertificateConfig_basicUpdated=`resource"aws_api_gateway_client_certificate""test"{description="HellofromTFacceptancetest-updated"}`functestAccClientCertificateConfig_tags1(tagKey1,tagValue1string)string{returnfmt.Sprintf(`resource"aws_api_gateway_client_certificate""test"{description="HellofromTFacceptancetest"tags={1]q=%[2]qfunc`,tagKey1,tagValue1)}functestAccClientCertificateConfig_tags2(tagKey1,tagValue1,tagKey2,tagValue2string)string{returnfmt.Sprintf(`resource"aws_api_gateway_client_certificate""test"{description="HellofromTFacceptancetest"tags={1]q=%[2]q3]q=%[4]qfunc`,tagKey1,tagValue1,tagKey2,tagValue2)}