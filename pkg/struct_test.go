package pkg

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Model struct {
	Username string
	Password string
}

type TestA struct {
	Model
	AUsername string
	APassword string
}

type TestB struct {
	Model     Model
	AUsername string
	APassword string
}

type TestC struct {
	Model
	Username string
	Password string
}

type TestD struct {
	Model    `json:"model"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestStruct(t *testing.T) {
	// Print
	fmt.Printf("%#v\n", TestA{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	fmt.Printf("%#v\n", TestB{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	fmt.Printf("%#v\n", TestC{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	fmt.Printf("%#v\n", TestD{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})

	fmt.Println(Marshal(TestA{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"}))
	fmt.Println(Marshal(TestB{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"}))
	fmt.Println(Marshal(TestC{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"}))
	fmt.Println(Marshal(TestD{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"}))
}

func Marshal(v any) string {
	result, _ := json.MarshalIndent(v, "", "\t")
	return string(result)
}
