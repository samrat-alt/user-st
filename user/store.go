package user

var userstore []User

func AddUser(user User) []User {
	userstore = append(userstore, user)
	return userstore
}

func GetAllUser() []User {
	return userstore
}
