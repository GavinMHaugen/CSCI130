package main 

import(
	"net/http"
	"html/template"
	"log"
	"github.com/nu7hatch/gouuid"
)

func uploadtemplate(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	name := req.FormValue("name")
	age := req.FormValue("age")

	cookie, err := req.Cookie("session")
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name: "session",
		Value: id.String() + name + age,
		HttpOnly: true,
		}
	http.SetCookie(res, cookie)
	tpl.Execute(res, nil)
}

func main(){
	http.HandlFunc("/", uploadtemplate)
	http.ListenAndServe(":8040", nil)
}