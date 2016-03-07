package main 

import(
	"net/http"
	"io"
	"strconv"
)

func cookie_server(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("user-cookie")
		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:  "user-cookie",
				Value: "0",
			}
		}
		i, _ := strconv.Atoi(cookie.Value)
		i++
		cookie.Value = strconv.Itoa(i)

		http.SetCookie(res, cookie)

		io.WriteString(res, cookie.Value)
}

func throw_the_icon(res http.ResponseWriter, req *http.Request){
	//does nothing 
}

func main(){
	http.HandleFunc("/", cookie_server)
	http.HandleFunc("/favicon.ico", throw_the_icon)
	http.ListenAndServe(":8080", nil)
}