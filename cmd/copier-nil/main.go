package main

import (
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/natholdallas/natools4go/spew"
)

type A struct {
	Name string
}

type B struct {
	Name *string
}

func main() {
	// Copy test
	println("\nCopy test\n")
	a := A{}
	err := copier.Copy(&a, B{nil})
	if err != nil {
		fmt.Println(err)
	}
	spew.JSON(a)
}
