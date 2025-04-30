package main

import "log"

func main() {
	arr := []int{1, 2, 3, 4}
	a := arr[0]
	b := arr[5] // panic
	log.Println(a, b)
}
