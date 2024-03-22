package main

import (
	"celebration/user"
	"fmt"
)

func main() {
	user := user.GetUserInput()
	fmt.Println("User Details:")
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
	fmt.Println("PhoneNumber:", user.PhoneNumber)
	fmt.Println("PhoneNumber:", user.Address)
}
