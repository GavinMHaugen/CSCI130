package main 

import(
	"log"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"net/http"
)

func uploadtemplate(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil{
		log.Fatalln(err)
	}

	cookie, err := req.Cookie("session")

	id, _ := uuid.NewV4()

		logError(err)

		cookie := &http.Cookie{
			Name:	"session",
			Value:	id.String(),
			HttpOnly:	true,
		}
		http.SetCookie(res, cookie)

	tpl.Execute(res, nil)
	
}

func main(){
	http.HandleFunc("/", uploadtemplate)
	http.ListenAndServe(":8080", nil)
}