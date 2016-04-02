package main 

import "net/http"

func serveRoot(res http.ResponseWriter, req *http.Request){
	res.Header.Set("Content-Type", "text/plain")
	res.Write([]byte("Test Server.\n"))
}

func main(){
	http.HandleFunc("/", serveRoot)
	http.ListenAndServe(":10443", "cert.pem", "key.pem", nil)
}