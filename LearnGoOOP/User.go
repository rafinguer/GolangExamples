package main

import "fmt"

type User struct {
	name   string
	email  string
	active bool
}

func (u *User) String() string {
	return fmt.Sprintf("User -> name: '%s', email: %s, active: %t", u.name, u.email, u.active)
}
