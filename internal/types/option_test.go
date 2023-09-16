// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package types

import (
	"testing"
)


tIsNone(t *testing.T) {
	t.Parallel()

	none := None[int]()
	some := Some("yes")

	if got, want := none.IsNone(), true; got != want {
t.Errorf("none.IsNone = %v, want = %v", got, want)
	}

	if got, want := some.IsNone(), false; got != want {
t.Errorf("some.IsNone = %v, want = %v", got, want)
	}
}


tIsSome(t *testing.T) {
	t.Parallel()

	none := None[int]()
	some := Some("yes")

	if got, want := none.IsSome(), false; got != want {
t.Errorf("none.IsSome = %v, want = %v", got, want)
	}

	if got, want := some.IsSome(), true; got != want {
t.Errorf("some.IsSome = %v, want = %v", got, want)
	}
}


tUnwrapOr(t *testing.T) {
	t.Parallel()

	none := None[int]()
	some := Some("yes")

	if got, want := none.UnwrapOr(42), 42; got != want {
t.Errorf("none.UnwrapOr = %v, want = %v", got, want)
	}

	if got, want := some.UnwrapOr("no"), "yes"; got != want {
t.Errorf("some.UnwrapOr = %v, want = %v", got, want)
	}
}


tUnwrapOrDefault(t *testing.T) {
	t.Parallel()

	none := None[int]()
	some := Some("yes")

	if got, want := none.UnwrapOrDefault(), 0; got != want {
t.Errorf("none.UnwrapOrDefault = %v, want = %v", got, want)
	}

	if got, want := some.UnwrapOrDefault(), "yes"; got != want {
t.Errorf("some.UnwrapOrDefault = %v, want = %v", got, want)
	}
}


tUnwrapOrElse(t *testing.T) {
	t.Parallel()

	none := None[int]()
	some := Some("yes")

	if got, want := none.UnwrapOrElse(
nt { return 42 }), 42; got != want {
t.Errorf("none.UnwrapOrElse = %v, want = %v", got, want)
	}

	if got, want := some.UnwrapOrElse(
tring { return "no" }), "yes"; got != want {
t.Errorf("some.UnwrapOrElse = %v, want = %v", got, want)
	}
}
