package user

import "fmt"

type User struct {
	Name        string
	Email       string
	PhoneNumber string
	Address     string
}

func GetUserInput() User {
	var user User

	fmt.Println("Enter your Name:")
	fmt.Scanln(&user.Name)

	fmt.Println("Enter your Email:")
	fmt.Scanln(&user.Email)

	fmt.Println("Enter your PhoneNumber:")
	fmt.Scanln(&user.PhoneNumber)

	fmt.Println("Enter your Address:")
	fmt.Scanln(&user.Address)

	return user
}
