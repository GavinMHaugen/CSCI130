package main 

import (
	"log"
	"os"
	"text/template"
)

type person struct{
	Name string
}

type condition struct{

	person 
	YOUQ bool
}

func main(){

	Person1 := condition{
		person: person{
			Name: "Henry",
		},
		YOUQ: false,
	}

	if Person1.Name == "gavin"{
		Person1.YOUQ = true
	}

	tmp, err := template.ParseFiles("template.gohtml")
	if err != nil{
		log.Fatalln(err)
	}

	err = tmp.Execute(os.Stdout, Person1)
	if err != nil{
		log.Fatalln(err)
	}

}