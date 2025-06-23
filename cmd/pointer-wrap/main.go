package main

import (
	"log"
	"time"
)

type User struct {
	Username  *string
	Password  *string
	FirstName *string
	LastName  *string
}

type UserNoPtr struct {
	Username  string
	Password  string
	FirstName string
	LastName  string
}

func main() {
	// s := UserNoPtr{
	// 	Username:  "arasf",
	// 	Password:  "asfb",
	// 	FirstName: "asjfba",
	// 	LastName:  "jaskbf",
	// }

	now := time.Now()
	// user1 := User{
	// 	Username: &s.Username,
	// }
	log.Println("origin: ", time.Since(now))

	now = time.Now()
	log.Println("wrap: ", time.Since(now))
}

func wrap[T any](s T) *T {
	return &s
}
