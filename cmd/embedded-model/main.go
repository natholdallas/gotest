package main

import (
	"testing"

	"github.com/natholdallas/natools4go/spew"
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
	spew.Struct(TestA{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	spew.Struct(TestB{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	spew.Struct(TestC{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	spew.Struct(TestD{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})

	// Jsonify
	spew.JSON(TestA{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	spew.JSON(TestB{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	spew.JSON(TestC{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	spew.JSON(TestD{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
}
