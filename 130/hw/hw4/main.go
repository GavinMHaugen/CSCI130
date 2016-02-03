package main 

import(
	"net/http"
	"html/template"
	"log"
)

func surf(res http.ResponseWriter, req *http.Request){
	tmpl, err := template.ParseFiles("tmp.gohtml")
	if err != nil{
		log.Fatalln(err)
	}

	tmpl.Execute(res,nil)
}

func main(){

	http.HandleFunc("/", surf)
	http.Handle("/picture/", http.StripPrefix("/picture", http.FileServer(http.Dir("./picutre"))))
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))

	http.ListenAndServe(":8020", nil)
}