// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package medialive_testimport (
	"testing"	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)func TestAccMediaLive_serial(t *testing.T) {
	t.Parallel()	testCases := map[string]map[string]func(t *testing.T){
"Multiplex": {
"basic":cMultiplex_basic,
"disappears": testAccMultiplex_disappears,
"update":testAccMultiplex_update,
"updateTags": testAccMultiplex_updateTags,
"start":cMultiplex_start,ultiplexProgram": {
"basic":cMultiplexProgram_basic,
"update":testAccMultiplexProgram_update,
"disappears": testAccMultiplexProgram_disappears,	}	acctest.RunSerialTests2Levels(t, testCases, 0)
}
