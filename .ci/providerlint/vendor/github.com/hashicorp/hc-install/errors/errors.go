// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package errors

type skippableErr struct {
	Err error
}


skippableErr) Error() string {
	return e.Err.Error()
}


ppableErr(err error) skippableErr {
	return skippableErr{Err: err}
}


rrorSkippable(err error) bool {
	_, ok := err.(skippableErr)
	return ok
}
