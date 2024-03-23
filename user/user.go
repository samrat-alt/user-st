package user

import "fmt"

type User struct {
	Id          string
	Name        string
	Email       string
	PhoneNumber string
	Address     string
}

func getInput(input string) (output string) {
	fmt.Printf("Enter your %s:", input)
	fmt.Scanln(&output)
	return output

}
func GetUserInput() User {
	var user User

	// fmt.Println("Enter your Email:")
	// fmt.Scanln(&user.Email)

	user.Email = getInput("Email")

	// fmt.Println("Enter your PhoneNumber:")
	// fmt.Scanln(&user.PhoneNumber)
	user.PhoneNumber = getInput("PhoneNumber")

	// fmt.Println("Enter your Address:")
	// fmt.Scanln(&user.Address)
	user.Address = getInput("Address")
	return user
}

func DeletUser(id string) (output string) {
	fmt.Printf("enter %s", id)
	fmt.Scanln(&output)
	return output
}

// func DeletUserInput(id string) User {
// 	var user User

// 	user.Email = getInput("Email")
// 	user.PhoneNumber = getInput("PhoneNumber")
// 	user.Address = getInput("Address")
// 	return user

// }

// fmt.Println("Enter your Email:")
// fmt.Scanln(&user.Email)
