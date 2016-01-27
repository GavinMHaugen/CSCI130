package main 

import "fmt"

func findmax(numbers ...int) int {
	var largestnum int

	for _, num := range numbers {
		if num > largestnum{
			largestnum = num
		}
	}

	return largestnum
}

func main(){
	nums := findmax(1, 2, 3, 4, 5, 6, 93453, 7, 8)
	fmt.Println(nums)
}