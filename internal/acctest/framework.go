// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package acctest

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	fwresource "github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/fwdiag"
)

// Terraform Plugin Framework variants of standard acceptance test helpers.


eteFrameworkResource(ctx context.Context, factory 
text.Context) (fwresource.ResourceWithConfigure, error), is *terraform.InstanceState, meta interface{}) error {
resource, err := factory(ctx)

if err != nil {
return err
}

resource.Configure(ctx, fwresource.ConfigureRequest{ProviderData: meta}, &fwresource.ConfigureResponse{})

schemaResp := fwresource.SchemaResponse{}
resource.Schema(ctx, fwresource.SchemaRequest{}, &schemaResp)

// Construct a simple Framework State that contains just top-level attributes.
state := tfsdk.State{
Raw:    tftypes.NewValue(schemaResp.Schema.Type().TerraformType(ctx), nil),
Schema: schemaResp.Schema,
}

for name, v := range is.Attributes {
if name == "%" || strings.Contains(name, ".") {
continue
}

if err := fwdiag.DiagnosticsError(state.SetAttribute(ctx, path.Root(name), v)); err != nil {
log.Printf("[WARN] %s(%s): %s", name, v, err)
}
}

response := fwresource.DeleteResponse{}
resource.Delete(ctx, fwresource.DeleteRequest{State: state}, &response)

if response.Diagnostics.HasError() {
return fwdiag.DiagnosticsError(response.Diagnostics)
}

return nil
}


ckFrameworkResourceDisappears(ctx context.Context, provo *schema.Provider, factory 
text.Context) (fwresource.ResourceWithConfigure, error), n string) resource.TestCheck

return 
terraform.State) error {
rs, ok := s.RootModule().Resources[n]
if !ok {
return fmt.Errorf("resource not found: %s", n)
}

if rs.Primary.ID == "" {
return fmt.Errorf("resource ID missing: %s", n)
}

return deleteFrameworkResource(ctx, factory, rs.Primary, provo.Meta())
}
}
