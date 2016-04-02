package main 

import (
	"io"
	"net/http"
	"fmt"
	"crypto/hmac"
	"crypto/sha256"
)

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("ourkey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func hmacaaa(res http.ResponseWriter, req *http.Request){
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	} 

	cookie, err := req.Cookie("session-id")
	if err != nil {
		cookie = &http.Cookie{
			Name:     "session-id",
			Value:    "cookie-value",
			HttpOnly: true,
		}
	}

	if req.FormValue("email") != "" {
		email := req.FormValue("email")
		cookie.Value = email + `|` + getCode(email)
	}

	http.SetCookie(res, cookie)
	io.WriteString(res, `<!DOCTYPE html>
	<html>
	  <body>
	    <form method="POST">
	    `+cookie.Value+`
	      <input type="email" name="email">
	      <input type="submit">
	    </form>
	  </body>
	</html>`)
} 

func main() {
	http.HandleFunc("/", hmacaaa)
	http.ListenAndServe(":8060", nil)
}