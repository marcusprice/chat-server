package user

type UserData struct {
	UserName  string
	FirstName string
	LastName  string
}

type User struct {
	userMap map[string]UserData
}

func (user User) UserExists(un string) bool {
	if _, ok := user.userMap[un]; ok {
		return true
	}

	return false
}

func (user *User) AddUser(ud UserData) {
	user.userMap[ud.UserName] = ud
}

func (user User) GetListOfUsers(un string) []string {
	var usernames []string
	for k := range user.userMap {
		if k == un {
			continue
		}

		usernames = append(usernames, k)
	}

	return usernames
}

func (user *User) RemoveUser(un string) {
	delete(user.userMap, un)
}

func InitUser() *User {
	um := make(map[string]UserData)
	return &User{
		userMap: um,
	}
}
