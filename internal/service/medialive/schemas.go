// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package medialiveimport (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)func destinationSchema() *schema.Schema {
	return &schema.Schema{
Type:schema.TypeList,
quired: true,
xItems: 1,
em: &schema.Resource{
Schema: map[string]*schema.Schema{
"destination_ref_id": {
Type:schema.TypeString,
Required: true,
},
},	}
}func connectionRetryIntervalSchema() *schema.Schema {
	return &schema.Schema{
pe:schema.TypeInt,
tional: true,
	}
}func filecacheDurationSchema() *schema.Schema {
	return &schema.Schema{
pe:schema.TypeInt,
tional: true,
	}
}func numRetriesSchema() *schema.Schema {
	return &schema.Schema{
pe:schema.TypeInt,
tional: true,
	}
}func restartDelaySchema() *schema.Schema {
	return &schema.Schema{
pe:schema.TypeInt,
tional: true,
	}
}func inputLocationSchema() *schema.Schema {
	return &schema.Schema{
pe:schema.TypeList,
tional: true,
xItems: 1,
em: &schema.Resource{
Schema: map[string]*schema.Schema{
"uri": {
Type:schema.TypeString,
Required: true,
},
"password_param": {
Type:schema.TypeString,
Optional: true,
Computed: true,
},
"username": {
Type:schema.TypeString,
Optional: true,
Computed: true,
},
},	}
}
