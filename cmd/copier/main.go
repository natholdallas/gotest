package main

import (
	"github.com/jinzhu/copier"
	t "github.com/natholdallas/gotest/pkg/tools"
)

type Model struct {
	ID uint
}

type Test struct {
	Model
	Username string
	Password string
}

type TestIn struct {
	ID       uint `copier:"Model"`
	Username string
	Password string
}

func main() {
	test := Test{}
	testIn := TestIn{1, "Username", "Password"}
	copier.Copy(&test, testIn)
	t.PrintJSON(test)
}
