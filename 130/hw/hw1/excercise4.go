package main 

import "fmt"

func main(){

	var smallnumber int
	var bignumber int
	var remainder int

	fmt.Println("Please enter a small number: ")
	fmt.Scanf("%d", &smallnumber)
	fmt.Println("Please enter a big number: ")
	fmt.Scanf("%d", &bignumber)

	remainder = (bignumber % smallnumber)

	fmt.Println(remainder)

}