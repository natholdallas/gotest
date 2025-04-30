package main

type Test struct {
	Username string
	Password string
}

func main() {
	var test *Test = nil
	if test != nil {
		println(test.Username)
		println(test.Password)
	}
}
