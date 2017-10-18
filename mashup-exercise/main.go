package main

import (
	"context"
	"sync"

	"git.playchat.net/playchat/mashup-exercise/datasources"
	"git.playchat.net/playchat/mashup-exercise/mashupserver"
)

func main() {
	// Set us up a mechanism to wait forever (or exit on)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// Create 2 new data sources.
	sourceA := datasources.NewSourceA()
	sourceB := datasources.NewSourceB()

	// Create our new mashup service on port 12000
	s := mashupserver.New()

	// Defer shutting everything down (will happen after we exit main())
	defer sourceA.Shutdown(context.Background())
	defer sourceB.Shutdown(context.Background())
	defer s.Shutdown(context.Background())

	// Wait forever.
	wg.Wait()
}
