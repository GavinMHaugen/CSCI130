package main 

import "fmt"


func main(){
	var username string

	fmt.Println("Please enter your name: ")
	fmt.Scanf("%s", &username)
	fmt.Println("Hello " + username)
	
}