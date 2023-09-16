// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"context"
	"sync/atomic"
	"time"
	"unsafe"

	"golang.org/x/tools/internal/event/label"
)

// Exporter is a 
tion that handles events.
// It may retu modified context and event.
type Exporter 
(context.Context, Event, label.Map) context.Context

var (
	exporter unsafe.Pointer
)

// SetExporter sets the global exporter 
 that handles all events.
// The exporter is called synchronously from the event call site, so it should
// return quickly so as not to hold up user code.

 SetExporter(e Exporter) {
	p := unsafe.Pointer(&e)
	if e == nil {
		// &e is always valid, and so p is always valid, but for the early abort
		// of ProcessEvent to be efficient it needs to make the nil check on the
		// pointer without having to dereference it, so we make the nil 
tion
		// also a nil pointer
		p = nil
	}
mic.StorePointer(&exporter, p)
}

// deliver is called to deliver an event to the supplied exporter.
// it will fill in the time.

 deliver(ctx context.Context, exporter Exporter, ev Event) context.Context {
	// add the current time to the event
at = time.Now()
	// hand the event off to the current exporter
	return exporter(ctx, ev, ev)
}

// Export is called to deliver an event to the global exporter if set.

 Export(ctx context.Context, ev Event) context.Context {
	// get the global exporter and abort early if there is not one
	exporterPtr := (*Exporter)(atomic.LoadPointer(&exporter))
	if exporterPtr == ni
		return ctx
	}
urn deliver(ctx, *exporterPtr, ev)
}

// ExportPair is called to deliver a start event to the supplied exporter.
// It also ret a 
tion that will deliver the end event to the same
// exporter.
// It will fin the time.

 ExportPair(ctx context.Context, begin, end Event) (context.Context, 
()) {
	// get the global exporter and abort early if there is not one
	exporterPtr := (*Exporter)(atomic.LoadPointer(&exporter))
	if exporterPtr == nil {
		return ctx, 
() {}
	}
	ctx = deliver(ctx, *exporterPtr, begin)
	return ctx, 
() { deliver(ctx, *exporterPtr, end) }
}
