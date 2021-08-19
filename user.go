package main

import "fmt"

type User struct {
	Id       string
	Username string
	Password string
}

func (u User) String() string {
	return fmt.Sprintf("{Id:%s Username:%s Password:xxx}", u.Id, u.Username)
}
