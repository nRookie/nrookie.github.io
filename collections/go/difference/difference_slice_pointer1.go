package main

// The difference is quite large:

// *[]Users would be a pointer to a slice of Users. Ex:

// On the contrary, []*Users would be a slice of pointers to Users. Ex:

import (
	"fmt"
)

type Users struct {
	ID   int
	Name string
}

var (
	user1 Users
	user2 Users
)

func main() {
	//Make a couple Users:
	user1 = Users{ID: 43215, Name: "Billy"}
	user2 = Users{ID: 84632, Name: "Bobby"}

	//Then make a list of pointers to those Users:
	var userList []*Users = []*Users{&user1, &user2}

	//Now you can change an individual Users in that list.
	//This changes the variable user2:
	*userList[1] = Users{ID: 1337, Name: "Larry"}

	fmt.Println(user1) // Outputs: {43215 Billy}
	fmt.Println(user2) // Outputs: {1337 Larry}
}
