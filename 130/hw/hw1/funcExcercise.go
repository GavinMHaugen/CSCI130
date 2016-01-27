package main

import "fmt"

func foo(x int) (int, bool) {

	return x / 2, x % 2 == 0
}

func main(){

	half, evencheck :=  foo(4)
	fmt.Println(half, evencheck)

}