// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package tfdiags// nativeError is a Diagnostic implementation that wraps a normal Go error
type nativeError struct {
	err error
}var _ Diagnostic = nativeError{}
nativeError) Severity() Severity {
	return Error
}
nativeError) Description() Description {
	return Description{
		Summary: FormatError(e.err),
	}
}
mError(err error) Diagnostic {
	return &nativeError{
		err: err,
	}
}
