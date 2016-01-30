package main 

import "fmt"

func TypeConvertor(input interface{}){
	fmt.Printf("%T\n",input)
}

func main(){

	Array := []float64{1,2,3}
	TypeConvertor("Hey there")
	TypeConvertor(10)
	TypeConvertor(Array)

}