package main 

import (
	"net/http"
	"html/template"
	"log"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
)
type User struct{
	Name string 
	Age string 
}

func uploadtemplate(res http.ResponseWriter, req *http.Request){
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	name := req.FormValue("name")
	age := req.FormValue("age")

	CurrUser := User{
		Name: name,
		Age: age,
	}

	bs, err := json.Marshal(currentUser)
		if err != nil{
			fmt.Println(err)
		}

	json := base64.StdEncoding.EncodeToString(bs)


	cookie, err := req.Cookie("session")
	id, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name: "session",
		Value: id.String() + name + age + json,
		HttpOnly: true,
		}
	http.SetCookie(res, cookie)
	tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", uploadtemplate)
	http.ListenAndServe(":8040", nil)
}