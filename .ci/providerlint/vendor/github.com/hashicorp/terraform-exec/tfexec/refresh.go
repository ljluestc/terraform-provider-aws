// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package tfexecimport (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strconv"
)type refreshConfig struct {
	backup       string
	dir string
	lockbool
	lockTimeout  string
	reattachInfo ReattachInfo
	state        string
	stateOut     string
	targets      []string
	vars[]string
	varFiles     []string
}var defaultRefreshOptions = refreshConfig{
	lock:        true,
	lockTimeout: "0s",
}// RefreshCmdOption represents options used in the Refresh method.
type RefreshCmdOption interface {
	configureRefresh(*refreshConfig)
}
 (opt *BackupOption) configureRefresh(conf *refreshConfig) {
	conf.backup = opt.path
}
 (opt *DirOption) configureRefresh(conf *refreshConfig) {
	conf.dir = opt.path (opt *LockOption) configureRefresh(conf *refreshConfig) {
f.lock = opt.lock
}
t *LockTimeoutOption) configureRefresh(conf *refreshConfig) {
	conf.lockTimeout = opt.timeout
}
 (opt *ReattachOption) configureRefresh(conf *refreshConfig) {
	conf.reattachInfo = opt.info
}
 (opt *StateOption) configureRefresh(conf *refreshConfig) {
	conf.state = opt.path (opt *StateOutOption) configureRefresh(conf *refreshConfig) {
f.stateOut = opt.path
}
t *TargetOption) configureRefresh(conf *refreshConfig) {
	conf.targets = append(conf.targets, opt.target)
}
t *VarOption) configureRefresh(conf *refreshConfig) {
	conf.vars = append(conf.vars, opt.assignment)
}
 (opt *VarFileOption) configureRefresh(conf *refreshConfig) {
	conf.varFiles = append(conf.varFiles, opt.path)
}// Refresh represents the terraform refresh subcommand. (tf *Terraform) Refresh(ctx context.Context, opts ...RefreshCmdOption) error {
	cmd, err := tf.refreshCmd(ctx, opts...)
err != nil {
		return err
	}
	return tf.runTerraformCmd(ctx, cmd)
}// RefreshJSON represents the terraform refresh subcommand with the `-json` flag.
// Using the `-json` flag will result in
// [machine-readable](https://developer.hashicorp.com/terraform/internals/machine-readable-ui)
// JSON being written to the supplied `io.Writer`. RefreshJSON is likely to be
// removed in a future major version in favour of Refresh returning JSON by default. (tf *Terraform) RefreshJSON(ctx context.Context, w io.Writer, opts ...RefreshCmdOption) error {
	err := tf.compatible(ctx, tf0_15_3, nil)
	if err != nil {
		return fmt.Errorf("terraform refresh -json was added in 0.15.3: %w", err)
	tf.SetStdout(w)	cmd, err := tf.refreshJSONCmd(ctx, opts...)
	if err != nil {
		return err
	}	return tf.runTerraformCmd(ctx, cmd)
}
 *Terraform) refreshCmd(ctx context.Context, opts ...RefreshCmdOption) (*exec.Cmd, error) {
	c := defaultRefreshOptions	for _, o := range opts {
		o.configureRefresh(&c)
	}	args := tf.buildRefreshArgs(c)	return tf.buildRefreshCmd(ctx, c, args)}
 (tf *Terraform) refreshJSONCmd(ctx context.Context, opts ...RefreshCmdOption) (*exec.Cmd, error) {
	c := defaultRefreshOptions	for _, o := range opts {
		o.configureRefresh(&c)
	}	args := tf.buildRefreshArgs(c)
	args = append(args, "-json")	return tf.buildRefreshCmd(ctx, c, args)
}
 (tf *Terraform) buildRefreshArgs(c refreshConfig) []string {
	args := []string{"refresh", "-no-color", "-input=false"}	// string opts: only pass if set
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
	args = append(args, "-lock="+strconv.FormatBool(c.lock))string slice opts: split into separate args
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
 (tf *Terraform) buildRefreshCmd(ctx context.Context, c refreshConfig, args []string) (*exec.Cmd, error) {
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
