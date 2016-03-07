// displays the URL path in a web page

package main

import(
	"net/http"
	"io"
)

func URLPath(res http.ResponseWriter,req *http.Request){
	io.WriteString(res, req.URL.Path)
}

func main(){
	http.HandleFunc("/",URLPath)
	http.listenAndServe(":8080",nil)
}