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

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("thisisthekey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func upload(res http.ResponseWriter, req *http.Request){
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
		Value: id.String() + name + age + json + getCode(id.String()),
		HttpOnly: true,
		}

	http.SetCookie(res, cookie)

	xs := string.Split(cookie.Value, "|")
	newcook := xs[0]
	newcodecook := xs[1]
	if (getCode(newcook) == newcodecook){
		fmt.Fprintf(res, "This cookie is still the same.")
	}else{
		fmt.Fprintf(res,"The cookie has been changed.")
	}
	
		tpl.Execute(res, nil)
}

func main() {
	http.HandleFunc("/", upload)
	http.ListenAndServe(":8040", nil)
}