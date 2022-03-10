package main
import (
	"fmt"
	"sync"
)

/*
This means that main can
 unblock all the senders simply by closing the done channel. 
 This close is effectively a broadcast signal to the senders. 
 We extend each of our pipeline functions to accept done as a parameter and arrange for the close to happen via a defer statement, 
 so that all return paths from main will signal the pipeline stages to exit
*/




func merge(done <-chan struct{}, cs ... <-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs. output
	// copies values from c to out until c is closed or it receives a value
	// from done, then output calls wg.Done.

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
				fmt.Printf("receive %d in merge \n", n)
			case <-done:
				fmt.Println("receive the done channel in merge")
				return
			}
		}
	}

	wg.Add(len(cs))
	fmt.Printf("count is %d\n", len(cs))
	for _, c := range cs {
		fmt.Println("start a new merge output routine")
		go output(c) 
	}

	// Start a go routine to close out once all the output goroutines are done.
	// This must start after the wg.Add call.

	go func() {
		wg.Wait()
		close(out)
	}() 
	fmt.Println("return the merged output channel")
	return out 
}

func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
				fmt.Printf("receive %d in sq\n", n )
			case <-done:
				fmt.Println("receive the done channel in sq")
				return
			}
		}
	}()
	return out
}

func gen(done <-chan struct{}, nums ...int) <-chan int {
    out := make(chan int)
	go func() {
		for _, n := range nums {
			select {
				case out <- n:
					fmt.Printf("generating %d\n", n)
				case <-done:
					fmt.Println("gen: receive done command")
					return
			}
		}
		fmt.Printf("close the gen out channel\n")
		close(out)
	}()
    return out
}


func main() {
	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer func() {
		fmt.Printf("close the main done channel\n")
		//close(done)
	}()

	in := gen(done, 2, 3, 4, 5)

	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(done, in)
	c2 := sq(done, in)
	c3 := sq(done, in)
	// Consume the first value from output.
	out := merge(done, c1, c2,c3)
	fmt.Println(<-out) // 4 or 9
	// done will be closed by the deferred call.

}


// Here are the guidelines for pipeline construction

// stages close their outbound channels when all the send operations are done.
// stages keep receiving values from inbound channels until those channels are close or the senders are unblocked.


// Pipelines unblock senders either by ensuring there's enough buffer for all the values that are sent or by explicitly signalling
// senders when the receiver may abandon the channel.

