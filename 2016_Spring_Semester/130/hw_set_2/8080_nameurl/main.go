package main 

import(
	"fmt"
	"net/http"
)

func URLPathInput(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res, req.FormValue("n"))
}

func main(){
	http.HandleFunc("/", URLPathInput)
	http.ListenAndServe(":8080", nil)
}