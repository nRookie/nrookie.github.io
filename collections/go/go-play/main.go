package main

import (
	store "GO_PLAY/B"
	"fmt"
)

func main() {
	a := &store.A{
		B: getB(),
	}
	a.AA = 20
	fmt.Print(a)
	fmt.Println()
	newA := &store.A{
		AA: 10,
		B:  getB(),
	}
	fmt.Print(newA)
}

func getB() *store.B {
	return &store.B{
		BB: 10,
	}
}
