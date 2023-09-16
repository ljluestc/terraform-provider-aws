// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mq_test

import (
	"testing"

	tfmq "github.com/hashicorp/terraform-provider-aws/internal/service/mq"
)
func TestCanonicalXML(t *testing.T) {
	t.Parallel()

	cases := []struct {
		Nameng
		Config
		Expected string
		ExpectError bool
	}{
		{
			Name:"Config sample from MSDN",
			Config:testAccForgeConfig_testExampleXMLFromMsdn,
			Expected: testAccForgeConfig_testExampleXMLFromMsdn,
		},
		{
			Name:"Config sample from MSDN, modified",
			Config:testAccForgeConfig_testExampleXMLFromMsdn,
			Expected: testExampleXML_from_msdn_modified,
		},
		{
			Name:fig sample from MSDN, flaw",
			Config:cForgeConfig_testExampleXMLFromMsdn,
			Expected: testExampleXML_from_msdn_flawed,
			ExpectError: true,
		},
		{
			Name: "A note",
			Config: `
<?xml version="1.0"?>
<note>
<to>You</to>
<from>Me</from>
<heading>Reminder</heading>
<body>You're awesome</body>
<rant/>
<rant/>
</note>
`,
			Expected: `
<?xml version="1.0"?>
<note>
	<to>You</to>
	<from>Me</from>
	<heading>
 Reminder
 </heading>
	<body>You're awesome</body>
	<rant/>
	<rant>
</rant>
</note>`,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			config, err := tfmq.CanonicalXML(tc.Config)
			if err != nil {
				t.Fatalf("Error getting canonical xml for given config: %s", err)
			}
			expected, err := tfmq.CanonicalXML(tc.Expected)
			if err != nil {
				t.Fatalf("Error getting canonical xml for expected config: %s", err)
			}

			if config != expected {
				if !tc.ExpectError {
					t.Fatalf("Error matching canonical xmls:\n\tconfig: %s\n\n\texpected: %s\n", config, expected)
				}
			}
		})
	}
}

const testAccForgeConfig_testExampleXMLFromMsdn = `
<?xml version="1.0"?>
<purchaseOrder xmlns="http://tempuri.org/po.xsd" orderDate="1999-10-20">
 <shipTo country="US">
e>Alice Smith</name>
eet>123 Maple Street</street>
y>Mill Valley</city>
te>CA</state>
>90952</zip>
 </shipTo>
 <billTo country="US">
e>Robert Smith</name>
eet>8 Oak Avenue</street>
y>Old Town</city>
te>PA</state>
>95819</zip>
 </billTo>
 <comment>Hurry, my lawn is going wild!</comment>
 <items>
m partNum="872-AA">
>Lawnmower</productName>
/quantity>
.95</USPrice>
firm this is electric</comment>
em>
m partNum="926-AA">
>Baby Monitor</productName>
/quantity>
98</USPrice>
99-05-21</shipDate>
em>
				<item/>
				<item/>
 </items>
</purchaseOrder>
`

const testExampleXML_from_msdn_modified = `
<?xml version="1.0"?>
<purchaseOrder xmlns="http://tempuri.org/po.xsd" orderDate="1999-10-20">
 <shipTo country="US">
e>Alice Smith</name>
eet>123 Maple Street</street>
y>Mill Valley</city>
te>CA</state>
>90952</zip>
 </shipTo>
 <billTo country="US">
e>Robert Smith</name>
eet>8 Oak Avenue</street>
y>Old Town</city>
te>PA</state>
>95819</zip>
 </billTo>
 <comment>Hurry, my lawn is going wild!</comment>
 <items>
m partNum="872-AA">
>Lawnmower</productName>
/quantity>
.95</USPrice>
firm this is electric</comment>
em>
m partNum="926-AA">
>Baby Monitor</productName>
/quantity>
98</USPrice>
99-05-21</shipDate>
em>
				  	 <item></item>
				<item>
</item>
 </items>
</purchaseOrder>
`

const testExampleXML_from_msdn_flawed = `
<?xml version="1.0"?>
<purchaseOrder xmlns="http://tempuri.org/po.xsd" orderDate="1999-10-20">
 <shipTo country="US">
e>Alice Smith</name>
eet>123 Maple Street</street>
y>Mill Valley</city>
te>CA</state>
>90952</zip>
 </shipTo>
 <billTo country="US">
e>Robert Smith</name>
eet>8 Oak Avenue</street>
y>Old Town</city>
te>PA</state>
>95819</zip>
 </billTo>
 <comment>Hurry, my lawn is going wild!</comment>
 <items>
m partNum="872-AA">
>Lawnmower</productName>
/quantity>
.95</USPrice>
firm this is electric</comment>
em>
m partNum="926-AA">
>Baby Monitor</productName>
/quantity>
98</USPrice>
99-05-21</shipDate>
em>
				<item>
				flaw
				</item>
 </items>
</purchaseOrder>
`
