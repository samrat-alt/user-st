package main

import (
	"celebration/user"
	"fmt"
)

func main() {
	var Aswe string
	var store user.UserStore //type

	for {
		user := user.GetUserInput()

		store = append(store, store.AddUser(&user))
		printInput(store)
		fmt.Println("do want to delete something")

		var conformDelete string
		fmt.Scanln(&conformDelete)
		if conformDelete == "yes" {
			fmt.Println("which id do you want to delete")

			var deletebyid string
			fmt.Scanln(&deletebyid)

			store = store.DeletUser(deletebyid)
			printInput(store)

		}
		fmt.Println("want to contiune yes/no")
		fmt.Scanln(&Aswe)
		if Aswe != "yes" {
			break
		}

		// fmt.Println("")

	}

	fmt.Println("bye")

}

func printInput(users []user.User) {
	fmt.Println("printing inputs")

	for index, value := range users {
		fmt.Printf("%d.[%s] %s\n", index+1, value.Id, value.Name)

	}

}
