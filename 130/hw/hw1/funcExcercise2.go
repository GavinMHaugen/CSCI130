package main 

import "fmt"

func main() {

	foo := func(x int) (int, bool) {
		return x / 2, x % 2 == 0
	}

	fmt.Println(foo(4))
}