// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package tfexecimport (
	"context"
	"fmt"
	"os/exec"	tfjson "github.com/hashicorp/terraform-json"
)// Metadata
s represents the terrm metadata 
tions -json subcommand. (tf *Terraform) Metadata
tions(ctx context.Context) (*tfjson.Metadata
tions, error) {
	:= tf.compatible(ctx, t_0, nil)
	if err != nil {
		return nil, fmt.Errorfrraform metadata 
tions was added in 1.4.0: %w", 
	}	
tionsCmd := tf.metadata
tionsCmd(ctx)	var ret tfjson.Metadata
s
	err = tf.runTerraformCmdJSON(ctx,
tionsCmd, &ret)
	if err != nil {
		return nil, err
	}	return &ret, nil
}
 (tf *Terraform) metadata
tionsCmd(ctx context.Context, args ...string) *exec.Cmd {
	allArgs := []string{"metadata", "
tions", "-json"}
	allArgs = append(allArgs, args...)	return tf.buildTerraformCmd(ctx, nil, allArgs...)
}
