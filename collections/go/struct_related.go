package main

import "fmt"

type A struct {
	B
}

type B struct {
	hello string `default:"FTP"`
}

func main() {
	b := B{}
	fmt.Println(b)
}
