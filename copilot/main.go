package main

import (
	"fmt"
)

type User struct {
	email    string
	username string
	age      int
}

func (u User) Email() string {
	return u.email
}

func (u *User) updateEmail(email string) {
	u.email = email
}

func Email(user *User) string {
	return user.email
}

func main() {
	user := User{
		email: "agg@foo.com",
	}
	user.updateEmail("foo@bar.com")
	fmt.Println(user.Email())
}

func updateUserAge(user *User, age int) {
	user.age = age
}
