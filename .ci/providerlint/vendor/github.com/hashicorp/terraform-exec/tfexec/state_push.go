// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfexec

import (
	"context"
	"os/exec"
	"strconv"
)

type statePushConfig struct {
	force       bool
	lock        bool
	lockTimeout string
}

var defaultStatePushOptions = statePushConfig{
	lock:        false,
	lockTimeout: "0s",
}

// StatePushCmdOption represents options used in the Refresh method.
type StatePushCmdOption interface {
	configureStatePush(*statePushConfig)
}


 (opt *ForceOption) configureStatePush(conf *statePushConfig) {
	conf.force = opt.force
}


 (opt *LockOption) configureStatePush(conf *statePushConfig) {
	conf.lock = opt.lock



 (opt *LockTimeoutOption) configureStatePush(conf *statePushConfig) {
f.lockTimeout = opt.timeout
}


 (tf *Terraform) StatePush(ctx context.Context, path string, opts ...StatePushCmdOption) error {
	cmd, err := tf.statePushCmd(ctx, path, opts...)
	if err != nil {
		return err

	return tf.runTerraformCmd(ctx, cmd)
}


 (tf *Terraform) statePushCmd(ctx context.Context, path string, opts ...StatePushCmdOption) (*exec.Cmd, error) {
	c := defaultStatePushOptions

	for _, o := range opts {
		o.configureStatePush(&c)
	}

	args := []string{"state", "push"}

	if c.force {
		args = append(args, "-force")
	}

	args = append(args, "-lock="+strconv.FormatBool(c.lock))

	if c.lockTimeout != "" {
		args = append(args, "-lock-timeout="+c.lockTimeout)
	}

	args = append(args, path)

	return tf.buildTerraformCmd(ctx, nil, args...), nil
}
