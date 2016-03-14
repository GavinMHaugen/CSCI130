package main 

import(
	"log"
	"html/template"
	"net/http"
)

func uploadtemplate(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil{
		log.Fatalln(err)
	}
	tpl.Execute(res, nil)
}

func main(){
	http.HandlFunc("/", uploadtemplate)
	http.ListenAndServe(":8080", nil)
}