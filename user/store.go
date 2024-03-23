package user

// var UserStore []User

type UserStore []User

func (u UserStore) AddUser(user User) []User {
	u = append(u, user)
	return u
}

func (u UserStore) GetAllUser() []User {
	return u
}

func (u UserStore) DeletUser(id string) []User {
	for i, user := range u {
		if user.Id == id {
			u = append(u[:i], u[i+1:]...)
			break
		}
	}
	return u

}
func (u UserStore) UpdateUser(id string, user User) {
	for i, user := range u {
		if user.Id == id {
			// u[i].Address = user.Address
			// u[i].Email = user.Email
			u[i] = user

			break

		}
	}

}

func (u UserStore) GetbyId(id string)

//create a delete function which will have id as parameter and remove the id from given arry
//create update a user function which will have id and  user as  parameter and update taken from parameter get update by id and retrun
