package main

import (
	"github.com/jinzhu/copier"
	"github.com/natholdallas/natools4go/ptr"
)

type Model struct {
	ID        uint
	CreatedAt string
}

type ModelOut struct {
	ID        uint
	CreatedAt string
}

type Config struct {
	DarkMode bool
}

type ConfigOut struct {
	DarkMode bool
}

type Test struct {
	Model
	Username string
	Password string
	Config   Config
}

type TestOut struct {
	ModelOut
	Username string
	Password string
	Config   ConfigOut
}

type In1 struct {
	ID        uint   `copier:"Model"`
	CreatedAt string `copier:"Model"`
	Username  string
	Password  string
}

type In2 struct {
	ID        uint
	CreatedAt string
	Username  string
	Password  string
}

func main() {
	// Copy test
	println("\nCopy test\n")
	test := Test{}
	copier.Copy(&test, In1{1, "Now", "Username", "Password"})
	ptr.JSON(test)
	test = Test{}
	copier.Copy(&test, In2{1, "Now", "Username", "Password"})
	ptr.JSON(test)
	test = Test{Username: "1234"}
	copier.Copy(&test, In2{1, "Now", "", "Password"})
	ptr.JSON(test)

	// DeepCopy and Copy compare
	println("\nDeepCopy and Copy compare\n")
	test3 := TestOut{}
	test = Test{Model: Model{1, "Now"}, Username: "Username", Password: "Password", Config: Config{true}}
	copier.Copy(&test3, test)
	ptr.JSON(test3)

	// Copy pointer value source
	println("\nCopy pointer value source\n")
	test = Test{}
	test4in := Test{Model: Model{1, "Now"}, Username: "Username", Password: "Password", Config: Config{true}}
	copier.Copy(&test, &test4in)
	ptr.JSON(test)
	test.Username = ""
	ptr.JSON(test)
	ptr.JSON(test4in)

	// Copy test reverse
	println("\nCopy test reverse\n")
	test = Test{Model: Model{ID: 1, CreatedAt: "12345"}}
	model := Model{}
	copier.Copy(&model, test)
	ptr.JSON(test)
	ptr.JSON(model)
}
