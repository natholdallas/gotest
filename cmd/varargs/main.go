package main

import "fmt"

func main() {
	PrintName("hello")
	PrintName("hello", "world")
}

func PrintName(name string, prefix ...string) {
	fmt.Println(name, Varargs(prefix))
}

func Varargs[T any](value []T) T {
	var res T
	if len(value) > 0 {
		res = value[0]
	}
	return res
}
