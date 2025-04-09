package main

import (
	"testing"

	t "github.com/natholdallas/gotest/pkg/tools"
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

func TestStruct(testing *testing.T) {
	// Print
	t.PrintStruct(TestA{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	t.PrintStruct(TestB{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	t.PrintStruct(TestC{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	t.PrintStruct(TestD{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})

	// Jsonify
	t.PrintJSON(TestA{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	t.PrintJSON(TestB{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	t.PrintJSON(TestC{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	t.PrintJSON(TestD{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
}
