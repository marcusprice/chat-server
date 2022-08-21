package user

type UserData struct {
	UserName  string
	FirstName string
	LastName  string
}

type User struct {
	userList map[string]UserData
}

func (user User) UserExists(un string) bool {
	if _, ok := user.userList[un]; ok {
		return true
	}

	return false
}

func (user *User) AddUser(ud UserData) {
	user.userList[ud.UserName] = ud
}

func (user User) GetListOfUsers(un string) []string {
	var usernames []string
	for k := range user.userList {
		if k == un {
			continue
		}

		usernames = append(usernames, k)
	}

	return usernames
}

func (user *User) RemoveUser(un string) {
	delete(user.userList, un)
}

func InitUser() *User {
	ul := make(map[string]UserData)
	return &User{
		userList: ul,
	}
}
