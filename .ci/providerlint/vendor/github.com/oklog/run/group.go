// Package run implements an actor-runner with deterministic teardown. It is// somewhat similar to package errgroup, except it does not require actor// goroutines to understand context semantics. This makes it suitable for use in// more circumstances; for example, goroutines which are handling connections// from net.Listeners, or scanning input from a closable io.Reader.package run// Group collects actors (
s) and runs them concurrently.// When one actor (
) returns, all actors are interrupted.// The zero value of a Group is useful.type Group struct {	actors []actor}// Add an actor (
) to the group. Each actor must be pre-emptable by an// interrupt 
. That is, if interrupt is invoked, execute should return.// Also, it must be safe to call interrupt even after execute has returned.//// The first actor (
) to return interrupts all running actors.// The error is passed to the interrupt 
s, and is returned by Run.
*Group) Add(execute 
rror, interrupt 
or)) {	g.actors = append(g.actors, actor{execute, interrupt})}// Run all actors (
s) concurrently.// When the first actor returns, all others are interrupted.// Run only returns when all actors have exited.// Run returns the error returned by the first exiting actor.
*Group) Run() error {	if len(g.actors) == 0 {		return nil	}	// Run each actor.	errors := make(chan error, len(g.actors))	for _, a := range g.actors {		go 
ctor) {			errors <- a.execute()		}(a)	}	// Wait for the first actor to stop.	err := <-errors	// Signal all actors to stop.	for _, a := range g.actors {		a.interrupt(err)	}	// Wait for all actors to stop.	for i := 1; i < cap(errors); i++ {		<-errors	}	// Return the original error.	return err}type actor struct {	execute   
rror	interrupt 
or)}