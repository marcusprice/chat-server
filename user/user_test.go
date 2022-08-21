package user

import (
	"reflect"
	"testing"
)

func TestInitUser(t *testing.T) {
	u := InitUser()
	if reflect.TypeOf(u) != reflect.TypeOf(&User{}) {
		t.Error("InitUser doesn't return a User type")
	}
}

func TestUserExists(t *testing.T) {
	u := InitUser()

	u.AddUser(UserData{
		UserName:  "username",
		FirstName: "firstName",
		LastName:  "lastName",
	})

	if !u.UserExists("username") {
		t.Error("UserExists returns false when user in fact exists")
	}

	if u.UserExists("freshUsername") {
		t.Error("UserExists returns true when user does not actually exist")
	}
}

func TestAddUser(t *testing.T) {
	u := InitUser()
	u.AddUser(UserData{
		UserName:  "username",
		FirstName: "firstName",
		LastName:  "lastName",
	})

	if len(u.userMap) != 1 {
		t.Error("AddUser does not result in user store in user.userMap")
	}

	if _, ok := u.userMap["username"]; !ok {
		t.Error("AddUser does not result in searchable user in user.userMap")
	}
}

func TestGetListOfUsers(t *testing.T) {
	u := InitUser()

	u.AddUser(UserData{
		UserName:  "username",
		FirstName: "firstName",
		LastName:  "lastName",
	})

	u.AddUser(UserData{
		UserName:  "username2",
		FirstName: "firstName",
		LastName:  "lastName",
	})

	ul := u.GetListOfUsers("username")

	if len(ul) == 2 || ul[0] == "username" {
		t.Error("GetListOfUsers is returning the current user in the user list")
	}

	if ul[0] != "username2" {
		t.Error("GetListOfUsers does not return other users")
	}
}

func TestRemoveUser(t *testing.T) {
	u := InitUser()

	u.AddUser(UserData{
		UserName:  "username",
		FirstName: "firstName",
		LastName:  "lastName",
	})

	u.RemoveUser("username")

	if len(u.userMap) == 1 {
		t.Error("User wasn't removed from userMap")
	}
}
