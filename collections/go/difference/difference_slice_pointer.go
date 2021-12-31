package main

// The difference is quite large:

// *[]Users would be a pointer to a slice of Users. Ex:

import (
	"fmt"
)

type Users struct {
	ID   int
	Name string
}

var (
	userList []Users
)

func main() {
	//Make the slice of Users
	userList = []Users{Users{ID: 43215, Name: "Billy"}, Users{ID: 43215, Name: "Billg"}}

	//Then pass the slice as a reference to some function
	myFunc(&userList)

	fmt.Println(userList) // Outputs: [{1337 Bobby}]
}

//Now the function gets a pointer *[]Users that when changed, will affect the global variable "userList"
func myFunc(input *[]Users) {
	//*input = []Users{Users{ID: 1337, Name: "Bobby"}}
	(*input)[1] = Users{ID: 1337, Name: "Bobby"}
	(*input) = append((*input), Users{ID: 13378, Name: "Bobby"})
}
