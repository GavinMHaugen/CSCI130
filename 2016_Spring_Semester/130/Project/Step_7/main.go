package main 

import (
	"io"
	"net/http"
)

func login_out(res http.ResponseWriter, req *http.Request){
	cookie, err := req.Cookie("in")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name: "in",
			Value: "0",
			HttpOnly: true,
		}
	}
	if req.Method == "POST" {
		password := req.FormValue("password")
		if password == "secret" {
			cookie = &http.Cookie{
				Name:"in",
				Value: "1",
				HttpOnly: true,
			}
		}
	}

	if req.URL.Path == "/logout" {
		cookie = &http.Cookie{
			Name: "in",
			Value: "0",
			HttpOnly: true,
		}
		http.SetCookie(res, cookie)
		http.Redirect(res, req, "/", 303)
		return
	}
	http.SetCookie(res, cookie)
	var html string

	if cookie.Value == "0" {
			html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<form method="POST">
				Enter name: <input type="text" name="name">
				<br>
				Enter age:  <input type="text" name="age">
				<br>
				Enter password: <input type="text" name="password">
				<br>
				<input type="submit">
			</form>
			</body>
			</html>`
		}

	if cookie.Value == "1" {
			html = `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title></title>
			</head>
			<body>
			<h1><a href="/logout">LOG OUT</a></h1>
			</body>
			</html>`
		}

		io.WriteString(res, html)

}

func main() {
	http.HandleFunc("/", login_out)
	http.ListenAndServe(":8010", nil)
}