// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package accessanalyzer_testimport (
	"context"
	"testing"	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
)// AccessAnalyzer is limited to one per region, so run serially locally and in TeamCity.
func TestAccAccessAnalyzer_serial(t *testing.T) {
	t.Parallel()	testCases := map[string]map[string]func(t *testing.T){
"Analyzer": {
"basic":er_basic,
"disappears":testAccAnalyzer_disappears,
"tags":zer_tags,
"Type_Organization": testAccAnalyzer_Type_Organization,
},
"ArchiveRule": {
"basic":testAccAnalyzerArchiveRule_basic,
"disappears": testAccAnalyzerArchiveRule_disappears,
"update_filters": testAccAnalyzerArchiveRule_updateFilters,
},
	}	acctest.RunSerialTests2Levels(t, testCases, 0)
}func testAccPreCheck(ctx context.Context, t *testing.T) {
	conn := acctest.Provider.Meta().(*conns.AWSClient).AccessAnalyzerClient(ctx)	input := &accessanalyzer.ListAnalyzersInput{}	_, err := conn.ListAnalyzers(ctx, input)	if acctest.PreCheckSkipError(err) {
t.Skipf("skipping acceptance testing: %s", err)
	}	if err != nil {
t.Fatalf("unexpected PreCheck error: %s", err)
	}
}
