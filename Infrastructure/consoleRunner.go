package Infrastructure

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

type CancellationTokenSource chan struct{} // bi-directional channel to stop application
type Runnable func(CancellationTokenSource)

// Run will execute your function in a go-routine and block the main thread.
// The CancellationTokenSource channel will run on system interrupt
func Run(fn Runnable) {
	var cts CancellationTokenSource
	cts = make(chan struct{})

	fmt.Println("Starting application at ", time.Now().UTC())
	fmt.Println("Press CTRL+C to stop the application gracefully")

	sigintEventHandler(cts)
	fn(cts) // run main application code

	<-cts                        // block the main thread until cts is pushed to channel
	fmt.Println("Shutting down") // wait for all other goroutines to stop gracefully
	time.Sleep(3 * time.Second)
}

func sigintEventHandler(cts chan struct{}) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	go func() {
		defer close(cts)
		defer close(c)
		defer signal.Stop(c) // deferred to the stack when this go-routine is disposed
		for range c {
			fmt.Println("System interrupt received. Stopping application.")
			return
		}
	}()
}
