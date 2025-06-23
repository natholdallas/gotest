package main

import (
	"testing"

	"github.com/natholdallas/natools4go/ptr"
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
	ptr.Struct(TestA{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	ptr.Struct(TestB{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	ptr.Struct(TestC{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	ptr.Struct(TestD{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})

	// Jsonify
	ptr.JSON(TestA{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	ptr.JSON(TestB{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	ptr.JSON(TestC{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
	ptr.JSON(TestD{Model{Username: "MUsername", Password: "MPassword"}, "Username", "Password"})
}
