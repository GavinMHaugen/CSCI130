package main 

import (
	"io"
	"net/http"
	"crypto/hmac"
	"net/http"
	"fmt"
	"encoding/json"
	"io"
)

type User struct {
	Name string
	Age  string
}

func getKey(data string) string{
	h := hmac.New(sha256.New, []byte("keyVal"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func foo () string{

	m := User{
		Name: "Bob",
		Age: "33",
	}

	bs, err := json.Marshal(m)
		if err != nil{
			fmt.Println(err)
		}

		return string(bs)
}

func main() {
	

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {


		
		cookie, err := req.Cookie("logged-in")

		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:  "logged-in",
				Value: "0",
				//Secure: true
				HttpOnly: true,
			}
			
		}

		if req.Method == "POST" {

			password := req.FormValue("password")
			if password == "secret"{
				cookie = &http.Cookie{
					Name: "logged-in",
					Value: "1",
					HttpOnly: true,
				}
			}
		}

		//if logout, then logout
		if req.URL.Path == "/logout" {
			cookie = &http.Cookie{
				Name: "logged-in",
				Value: "0",
				MaxAge: -1,
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
			<h1>LOG IN</h1>
			<form method="POST">
				<h3>Name</h3>
				<input type="text" name="name">
				<br>
				<h3>Age</h3>
				<input type="text" name="age">
				<br>
				<h3>Password</h3>
				<input type="text" name="password">
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
		})

		
	http.ListenAndServe(":8020", nil)
}