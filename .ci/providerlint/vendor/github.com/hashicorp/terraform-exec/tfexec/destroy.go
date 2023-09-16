// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package tfexecimport (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strconv"
)type destroyConfig struct {
	backup string
	dir    string
	lock   bool	// LockTimeout must be a string with time unit, e.g. '10s'
	lockTimeout  string
	parallelism  int
	reattachInfo ReattachInfo
	refresh      bool
	state        string
	stateOut     string
	targets      []string	// Vars: each var must be supplied as a single string, e.g. 'foo=bar'
	vars     []string
	varFiles []string
}var defaultDestroyOptions = destroyConfig{
	lock:        true,
	lockTimeout: "0s",
	parallelism: 10,
	refresh:     true,
}// DestroyOption represents options used in the Destroy method.
type DestroyOption interface {
	configureDestroy(*destroyConfig)
}
 (opt *DirOption) configureDestroy(conf *destroyConfig) {
	conf.dir = opt.path
}
 (opt *ParallelismOption) configureDestroy(conf *destroyConfig) {
	conf.parallelism = opt.parallelism (opt *BackupOption) configureDestroy(conf *destroyConfig) {
f.backup = opt.path
}
t *TargetOption) configureDestroy(conf *destroyConfig) {
	conf.targets = append(conf.targets, opt.target)
}
 (opt *LockTimeoutOption) configureDestroy(conf *destroyConfig) {
	conf.lockTimeout = opt.timeout
}
 (opt *StateOption) configureDestroy(conf *destroyConfig) {
	conf.state = opt.path (opt *StateOutOption) configureDestroy(conf *destroyConfig) {
f.stateOut = opt.path
}
t *VarFileOption) configureDestroy(conf *destroyConfig) {
	conf.varFiles = append(conf.varFiles, opt.path)
}
 (opt *LockOption) configureDestroy(conf *destroyConfig) {
	conf.lock = opt.lock
}
 (opt *RefreshOption) configureDestroy(conf *destroyConfig) {
	conf.refresh = opt.refresh
}
 (opt *VarOption) configureDestroy(conf *destroyConfig) {
	conf.vars = append(conf.vars, opt.assignment)
}
 (opt *ReattachOption) configureDestroy(conf *destroyConfig) {
	conf.reattachInfo = opt.info
}// Destroy represents the terraform destroy subcommand. *Terraform) Destroy(ctx context.Context, opts ...DestroyOption) error {
	cmd, err := tf.destroyCmd(ctx, opts...)
	if err != nil {
		return err
	}
	return tf.runTerraformCmd(ctx, cmd)
}// DestroyJSON represents the terraform destroy subcommand with the `-json` flag.
// Using the `-json` flag will result in
// [machine-readable](https://developer.hashicorp.com/terraform/internals/machine-readable-ui)
// JSON being written to the supplied `io.Writer`. DestroyJSON is likely to be
// removed in a future major version in favour of Destroy returning JSON by default. (tf *Terraform) DestroyJSON(ctx context.Context, w io.Writer, opts ...DestroyOption) error {
	err := tf.compatible(ctx, tf0_15_3, nil)
err != nil {
		return fmt.Errorf("terraform destroy -json was added in 0.15.3: %w", err)
	}	tf.SetStdout(w)	cmd, err := tf.destroyJSONCmd(ctx, opts...)
	if err != nil {
		return err
	}	return tf.runTerraformCmd(ctx, cmd) (tf *Terraform) destroyCmd(ctx context.Context, opts ...DestroyOption) (*exec.Cmd, error) {
	c := defaultDestroyOptions	for _, o := range opts {
		o.configureDestroy(&c)
	}	args := tf.buildDestroyArgs(c)	return tf.buildDestroyCmd(ctx, c, args) (tf *Terraform) destroyJSONCmd(ctx context.Context, opts ...DestroyOption) (*exec.Cmd, error) {
	c := defaultDestroyOptions	for _, o := range opts {
		o.configureDestroy(&c)
	}	args := tf.buildDestroyArgs(c)
	args = append(args, "-json")	return tf.buildDestroyCmd(ctx, c, args)
}
 (tf *Terraform) buildDestroyArgs(c destroyConfig) []string {
	args := []string{"destroy", "-no-color", "-auto-approve", "-input=false"}	// string opts: only pass if set
	if c.backup != "" {
		args = append(args, "-backup="+c.backup)
	}
	if c.lockTimeout != "" {
		args = append(args, "-lock-timeout="+c.lockTimeout)
	}
	if c.state != "" {
		args = append(args, "-state="+c.state)
	}
	if c.stateOut != "" {
		args = append(args, "-state-out="+c.stateOut)
	}
	for _, vf := range c.varFiles {
		args = append(args, "-var-file="+vf)
	}	// boolean and numerical opts: always pass
	args = append(args, "-lock="+strconv.FormatBool(c.lock))
	args = append(args, "-parallelism="+fmt.Sprint(c.parallelism))
s = append(args, "-refresh="+strconv.FormatBool(c.refresh))	// string slice opts: split into separate args
	if c.targets != nil {
		for _, ta := range c.targets {
			args = append(args, "-target="+ta)
		}
	}
	if c.vars != nil {
		for _, v := range c.vars {
			args = append(args, "-var", v)
		}
	}	return args
}
 (tf *Terraform) buildDestroyCmd(ctx context.Context, c destroyConfig, args []string) (*exec.Cmd, error) {
	// optional positional argument
	if c.dir != "" {
		args = append(args, c.dir)
	}	mergeEnv := map[string]string{}
	if c.reattachInfo != nil {
		reattachStr, err := c.reattachInfo.marshalString()
		if err != nil {
			return nil, err
		}
		mergeEnv[reattachEnvVar] = reattachStr
	}	return tf.buildTerraformCmd(ctx, mergeEnv, args...), nil
}
