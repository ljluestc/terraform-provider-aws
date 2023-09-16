// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package tfdiagstype simpleWarning stringvar _ Diagnostic = simpleWarning("")// SimpleWarning constructs a simple (summary-only) warning diagnostic.pleWarning(msg string) Diagnostic {
	return simpleWarning(msg)
}
simpleWarning) Severity() Severity {
	return Warning
}
simpleWarning) Description() Description {
	return Description{
		Summary: string(e),
	}
}
