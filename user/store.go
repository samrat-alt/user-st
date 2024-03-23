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

//create a delete function which will have id as parameter and remove the id from given arry
//update a user function which will have id and user parameter and update taken from para meter
//get by id and retrun
