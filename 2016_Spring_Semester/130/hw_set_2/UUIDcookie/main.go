package main 

import(
	"fmt"
	"log"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

func foo(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("session")
	if err != nil {
		fmt.Fprint(res, "Cookie granted\n")
		id, _ := uuid.NewV4()
		cookie = &http.Cookit{
			Name: "session"
			Value: id.String(),
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
	}
	fmt.Fprint(res, cookie)
}

func main(){
	http.HandleFunc("/", foo)
	http.ListenAndServe(":8080", nil)
}