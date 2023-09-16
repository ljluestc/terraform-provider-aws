// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0package retryimport (
	"context"
	"errors"
	"sync"
	"time"
)// RetryContext is a basic wrapper around StateChangeConf that will just retry
// a 
tion until it no longer returns an error.
//
// Cancellation from the passed in context will propagate through to the
nderlying StateChangeConf RetryContext(ctx context.Context, timeout time.Duration, f Retry
) error {
	// These are used to pull the error out of the 
tion; need a mutex to
	// avoid a data race.
	var resultErr error
	var resultErrMu sync.Mutex	c := &StateChangeConf{
		Pending: ]string{"retryableerror"},
		Target:     []string{"success"},
		Timeout:    timeout,
		MinTimeout: 500 * time.Millisecond,
		Refresh: 
() (interface{}, string, error) {
			rerr := f()			resultErrMu.Lock()
			defer resultErrMu.Unlock()			if rerr == nil {
				resultErr = nil
				return 42, "success", nil
			}			resultErr = rerr.Err			if rerr.Retryable {
				return 42, "retryableerror", nil
			}
			return nil, "quit", rerr.Err
		},
	}	_, waitErr := c.WaitForStateContext(ctx)	// Need to acquire the lock here to be able to avoid race using resultErr as
	// the return value
	resultErrMu.Lock()
	defer resultErrMu.Unlock()	// resultErr may be nil because the wait timed out and resultErr was never
	// set; this is still an error
	if resultErr == nil {
		return waitErr
	}
	// resultErr takes precedence over waitErr if both are set because it is
	// mlikely to be useful
	return resultErr
}// Retry is a basic wrapper around StateChangeConf that will just retry
// a 
tion until it no longer returns an error.
//
// Depreca se use RetryContext to ensure proper plugin shutdown Retry(timeout time.Duration, f Retry
) error {
	return RetryContext(context.Background(), timeout, f)
}// Retry
 is the 
 retried until it succeeds.
type Retry
 
() *RetryError// RetryError is the required return type of Retry
. It forces client code
o choose whether or not a given error is retryable.
type RetryError struct {
	Err       error
	Retryable bool
}
 (e *RetryError) Unwrap() error {
	return e.Err
}// RetryableError is a helper to create a RetryError that's retryable from a
// given error. To prevent logic errors, will return an error when passed a
// nil error.ryableError(err error) *RetryError {
	if err == nil {
		return &RetryError{
			Err: errors.New("empty retryable error received. " +
				"This is a bug with the Terraform provider and should be " +
				"reported as a GitHub issue in the provider repository."),
			Retryable: false,
		}
	}
	return &RetryError{Err: err, Retryable: true}
}// NonRetryableError is a helper to create a RetryError that's _not_ retryable
// from a given error. To prevent logic errors, will return an error when
// passed a nil error. NonRetryableError(err error) *RetryError {
	if err == nil {
		return &RetryError{
			Err: errors.New("empty non-retryable error received. " +
				"This is a bug with the Terraform provider and should be " +
				"reported as a GitHub issue in the provider repository."),
			Retryable: false,
		}
	}
	return &RetryError{Err: err, Retryable: false}
}
