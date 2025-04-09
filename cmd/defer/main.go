package main

func defer1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func defer2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func defer3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	println("defer1", defer1())
	println("defer2", defer2())
	println("defer3", defer3())
}
