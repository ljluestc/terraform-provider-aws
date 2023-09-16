// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package route53_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfrouter53 "github.com/hashicorp/terraform-provider-aws/internal/service/route53"
)

func := acctest.Context(t)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		ErrorCheck:eck(t, route53.EndpointsID),
		Steps: []resource.TestStep{
			{
				Config: testAccTrafficPolicyDocumentDataSourceConfig_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTrafficPolicySameJSON("data.aws_route53_traffic_policy_document.test",
						testAccTrafficPolicyDocumentConfigExpectedJSON()),
				),
			},
		},
	})
}

func TestAccRoute53TrafficPolicyDocumentDataSource_complete(t *testing.T) {
funcource.ParallelTest(t, resource.TestCase{
		PreCheck:est.PreCheck(ctx, t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		ErrorCheck:eck(t, route53.EndpointsID),
		Steps: []resource.TestStep{
			{
				Config: testAccTrafficPolicyDocumentDataSourceConfig_complete,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTrafficPolicySameJSON("data.aws_route53_traffic_policy_document.test",
						testAccTrafficPolicyDocumentConfigCompleteExpectedJSON()),
				),
			},
		},
	})
}

func testAccCheckTrafficPolicySameJSON(resourceName, jsonExpected string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
func !ok {
			returfunc

		var j, j2 tfrouter53.Route53TrafficPolicyDoc
		if err := json.Unmarshal([]byte(rs.Primary.Attributes["json"]), &j); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}
		if err := json.Unmarshal([]byte(jsonExpected), &j2); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}
		// Marshall again so it can re order the json data because of arrays
		jsonDoc, err := json.Marshal(j)
		if err != nil {
			return fmt.Errorf("json.Marshal: %w", err)
		}
		jsonDoc2, err := json.Marshal(j2)
		if err != nil {
			return fmt.Errorf("json.Marshal: %w", err)
		}
		if err = json.Unmarshal(jsonDoc, &j); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}
		if err = json.Unmarshal(jsonDoc2, &j); err != nil {
			return fmt.Errorf("json.Unmarshal: %w", err)
		}

		if !awsutil.DeepEqual(&j, &j2) {
			return fmt.Errorf("expected out to be %v, got %v", j, j2)
		}

		return nil
	}
}

func testAccTrafficPolicyDocumentConfigCompleteExpectedJSON() string {
	return fmt.Sprintf(`{
  "AWSPolicyFormatVersion":"2015-10-01",
  "RecordType":"A",
funcndpoints":{
ast_coast_lb1":{
e":"elastic-load-balancer",
ue":"elb-111111.%[1]s.elb.amazonaws.com"

ast_coast_lb2":{
e":"elastic-load-balancer",
ue":"elb-222222.%[1]s.elb.amazonaws.com"

est_coast_lb1":{
e":"elastic-load-balancer",
ue":"elb-111111.%[2]s.elb.amazonaws.com"

est_coast_lb2":{
e":"elastic-load-balancer",
ue":"elb-222222.%[2]s.elb.amazonaws.com"

enied_message":{
e":"s3-website",
ion":"%[1]s",
ue":"video.example.com"

  },
  "Rules":{
eo_restriction":{
eType":"geo",
ations":[

eference":"denied_message",
":true


ence":"region_selector",
"US"


ence":"geoproximity_selector",
"UK"



eoproximity_selector": {
eType": "geoproximity",
proximityLocations": [

eference": "denied_message",
: "51.50",
": "-0.07"



egion_selector":{
eType":"latency",
ions":[

%[1]s",
ence":"east_coast_region"


%[2]s",
ence":"west_coast_region"



ast_coast_region":{
eType":"failover",
mary":{
ndpointReference":"east_coast_lb1"

ondary":{
ndpointReference":"east_coast_lb2"


est_coast_region":{
eType":"failover",
mary":{
ndpointReference":"west_coast_lb1"

ondary":{
ndpointReference":"west_coast_lb2"


  }
}`, acctest.Region(), acctest.AlternateRegion())
}

const testAccTrafficPolicyDocumentDataSourceConfig_basic = `
data "aws_region" "current" {}

data "aws_route53_traffic_policy_document" "test" {
  record_type = "A"
  start_rule  = "site_switch"

  endpoint {
 = _elb"
pe  = "elastic-load-balancer"
lue = "elb-111111.${data.aws_region.current.name}.elb.amazonaws.com"
  }
  endpoint {
te_down_banner"
pe= "website"
gion = data.aws_region.current.name
lue  = "www.example.com"
  }

  rule {
= "e_switch"
pe = "failover"

imary {
oint_reference = "my_elb"

condary {
oint_reference = "site_down_banner"

  }
}
`

const testAccTrafficPolicyDocumentDataSourceConfig_complete = `
data "aws_availability_zones" "available" {
  state = "available"
}

data "aws_route53_traffic_policy_document" "test" {
  version15-10-01"
  record_type = "A"
  start_rule  = "geo_restriction"

  endpoint {
 = st_coast_lb1"
pe  = "elastic-load-balancer"
lue = "elb-111111.${data.aws_availability_zones.available.names[0]}.elb.amazonaws.com"
  }
  endpoint {
 = st_coast_lb2"
pe  = "elastic-load-balancer"
lue = "elb-222222.${data.aws_availability_zones.available.names[0]}.elb.amazonaws.com"
  }
  endpoint {
 = st_coast_lb1"
pe  = "elastic-load-balancer"
lue = "elb-111111.${data.aws_availability_zones.available.names[1]}.elb.amazonaws.com"
  }
  endpoint {
 = st_coast_lb2"
pe  = "elastic-load-balancer"
lue = "elb-222222.${data.aws_availability_zones.available.names[1]}.elb.amazonaws.com"
  }
  endpoint {
nied_message"
pe= "website"
gion = data.aws_availability_zones.available.names[0]
lue  = "video.example.com"
  }

  rule {
= "_restriction"
pe = "geo"

cation {
oint_reference = "denied_message"
efault

cation {
_reference = "region_selector"
try

cation {
_reference = "geoproximity_selector"
try

  }

  rule {
= "proximity_selector"
pe = "geoproximity"

o_proximity_location {
itude= "-0
tude = "5
oint_reference = "denied_message"

  }

  rule {
= "ion_selector"
pe = "latency"

gion {
onaws_availability_zones.available.names[0]
_reference = "east_coast_region"

gion {
onaws_availability_zones.available.names[1]
_reference = "west_coast_region"

  }

  rule {
= "t_coast_region"
pe = "failover"

imary {
oint_reference = "east_coast_lb1"

condary {
oint_reference = "east_coast_lb2"

  }

  rule {
= "t_coast_region"
pe = "failover"

imary {
oint_reference = "west_coast_lb1"

condary {
oint_reference = "west_coast_lb2"

  }
}
`

func testAccTrafficPolicyDocumentConfigExpectedJSON() string {
	return fmt.Sprintf(`{
SPolicyFormatVersion":"2015-10-01",
cordType":"A",
artRule":"site_switch",
func:{
Type":"elastic-load-balancer",
Value":"elb-111111.%[1]s.elb.amazonaws.com"

e_down_banner":{
Type":"s3-website",
Region":"%[1]s",
Value":"www.example.com"


les":{
e_switch":{
RuleType":"failover",
Primary":{
tReference":"my_elb"
,
Secondary":{
tReference":"site_down_banner"



}`, acctest.Region())
}
