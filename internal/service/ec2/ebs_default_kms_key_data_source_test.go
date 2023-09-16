// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
)


func := acctest.Context(t)
	resource.ParallelTest(t, resource.TestCase{
PreCheck:
func() { acctest.PreCheck(ctx, t) },
funcoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
Steps: []resource.TestStep{
	{
Config: testAccEBSDefaultKMSKeyDataSourceConfig_basic,
Check: resource.ComposeTestCheck
func(
	testAccCheckEBSDefaultKMSKey(ctx, "data.aws_ebs_default_kms_key.current"),
func
},
	})
}

const testAccEBSDefaultKMSKeyDataSourceConfig_basic = `
data "aws_ebs_default_kms_key" "current" {}
`
