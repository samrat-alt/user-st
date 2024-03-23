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
	user.Id = getInput("Id")
	user.Name = getInput("name")

	user.Email = getInput("Email")

	user.PhoneNumber = getInput("PhoneNumber")

	user.Address = getInput("Address")
	return user
}
