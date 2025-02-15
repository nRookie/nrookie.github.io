package main

import (
	"fmt"
	"sync"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("gen send %d \n", n )
			out <- n
		}		
		fmt.Printf("close out gen\n")
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("sq send %d \n", n *n)
			out <- n * n
		}
		fmt.Printf("close out sq\n")
		close(out)
	}()
		return out
}


// func main() {
// 	// Set up the pipeline and consume the output.
// 	for n := range sq(sq(gen(2, 3))) {
// 		fmt.Println(n) // 16 then 81
// 	}
// }


func merge(cs ... <-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 4)

	// Start an output goroutine for each input channel in cs. output
	// copies values from c to out until c is closed, then calls wg.Done.

	output := func(c <-chan int) {
		for n := range c {
			fmt.Printf("merge recv: %d\n", n)
			out <- n 
		}
		wg.Done()
	}

	wg.Add(len(cs))
	for _, c := range cs {
		fmt.Printf("start a new output routine %d \n", c)
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
func main() {
	in := gen(2, 3, 4, 5)
	// Distribute the sq work across two goroutines that both read from in.
	c1 := sq(in)
	// // Consume the merged output from c1 and c2.
	// for n := range merge(c1, c2) {
	// 	fmt.Println(n) // 4 then 9, or 9 then 4.
	// }

	out := merge(c1)
    //fmt.Println(<-out) // 4 or 9

	for i := range out {
		fmt.Printf("range value %d\n", i)
	}
    return

}