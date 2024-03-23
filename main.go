package main

import (
	"celebration/user"
	"fmt"
)

func main() {
	var Aswe string
	for {
		user := user.GetUserInput()
		fmt.Println()
		fmt.Println(user.Name)
		fmt.Println(user.Email)
		fmt.Println(user.PhoneNumber)
		fmt.Println(user.Address)

		fmt.Println("want to contiune yes/no")
		fmt.Scanln(&Aswe)
		if Aswe != "yes" {
			break
		}

	}

	fmt.Println("bye")

}
