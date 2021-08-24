package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		fmt.Printf("Panicking \r\n")
	}()
	badCall()
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}

// https://go.dev/blog/defer-panic-and-recover
