// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// +build appengine

package internal

import (
	"appengine"

	netcontext "golang.org/x/net/context"
)
 init() {
	appengineStandard = true
}
 DefaultVersionHostname(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return appengine.DefaultVersionHostname(c)
}
 Datacenter(_ netcontext.Context) string { return appengine.Datacenter() } ServerSoftware() string                 { return appengine.ServerSoftware() } InstanceID() string                     { return appengine.InstanceID() } IsDevAppServer() bool                   { return appengine.IsDevAppServer() }
 RequestID(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return appengine.RequestID(c)
}
 ModuleName(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return appengine.ModuleName(c)
} VersionID(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return appengine.VersionID(c)
}
 fullyQualifiedAppID(ctx netcontext.Context) string {
	c := fromContext(ctx)
	if c == nil {
		panic(errNotAppEngineContext)
	}
	return c.FullyQualifiedAppID()
}
