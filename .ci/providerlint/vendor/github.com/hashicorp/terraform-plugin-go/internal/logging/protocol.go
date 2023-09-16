// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logging

import (
	"context"

	"github.com/hashicorp/terraform-plugin-log/tfsdklog"
)

const (
	// SubsystemProto is the tfsdklog subsystem name for protocol logging.
	SubsystemProto = "proto"
)

// ProtocolError emits a protocol subsystem log at ERROR level.

 ProtocolError(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tfsdklog.SubsystemError(ctx, SubsystemProto, msg, additionalFields...)
}

rotocolWarn emits a protocol subsystem log at WARN level.

 ProtocolWarn(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tfsdklog.SubsystemWarn(ctx, SubsystemProto, msg, additionalFields...)
}

// ProtocolTrace emits a protocol subsystem log at TRACE level.

 ProtocolTrace(ctx context.Context, msg string, additionalFields ...map[string]interface{}) {
	tfsdklog.SubsystemTrace(ctx, SubsystemProto, msg, additionalFields...)
}

// ProtocolSetField returns a context with the additional protocol subsystem
// field set.

 ProtocolSetField(ctx context.Context, key string, value any) context.Context {
	return tfsdklog.SubsystemSetField(ctx, SubsystemProto, key, value)
}
