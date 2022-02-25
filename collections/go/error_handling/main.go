package main

import (
	"errors"
	"fmt"
	"os"
)

type PathError struct {
	Path string
}

func (e *PathError) Error() string {
	return fmt.Sprintf("error in path: %v", e.Path)
}

func throwError() error {
	return &PathError{Path: "/test"}
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, errors.New("cannot divide through zero")
	}

	return a / b, nil
}

func accessSlice(slice []int, index int) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("internal error: %v", p)
		}
	}()

	fmt.Printf("item %d, value %d \n", index, slice[index])
	defer fmt.Printf("defer %d \n", index)
	accessSlice(slice, index+1)
}

func openFile(filename string) error {
	if _, err := os.Open(filename); err != nil {
		return fmt.Errorf("error opening %s: %w", filename, err)
	}

	return nil
}

func main() {
	// Casting error
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *os.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}

	err := openFile("non-existing")

	if err != nil {
		fmt.Printf("error running program: %s \n", err.Error())

		// Unwrap error
		unwrappedErr := errors.Unwrap(err)
		fmt.Printf("unwrapped error: %v \n", unwrappedErr)
	}

	accessSlice([]int{1, 2, 5, 6, 7, 8}, 0)
	err = throwError()

	if err != nil {
		fmt.Println(err)
	}

	switch e := err.(type) {
	case *PathError:
		fmt.Println("do something with the path")
	default:
		fmt.Println(e)
	}

	num, err := divide(100, 0)

	if err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Println("Number: ", num)
	}

}

//https://gabrieltanner.org/blog/golang-error-handling-definitive-guide

/*
As() function
Similar to Is(), the As(err error, target interface{}) bool checks if any error in the chain of wrapped errors matches the target. The difference is that this function checks whether the error has a specific type, unlike the Is(), which examines if it is a particular error object. Because As considers the whole chain of errors, it should be preferable to the type assertion if e, ok := err.(*BadInputError); ok.

target argument of the As(err error, target interface{}) bool function should be a pointer to the error type, which in this case is *BadInputError
*/
