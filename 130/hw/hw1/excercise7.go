package main 

import "fmt"

func main(){

	var sum int

	for i := 0; i < 1000; i++ {
		if i % 3 == 0 {
			sum += i
			continue
		} else if i % 5 == 0{
			sum += i
			continue
		}

	}

	fmt.Println(sum)


}