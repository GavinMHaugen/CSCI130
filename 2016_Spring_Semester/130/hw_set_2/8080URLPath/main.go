package main 

import (
	"net/http"
	"fmt"
	"strings"
)

func PathN(res http.ResponseWriter,req *http.Request){
	name := strings.Split(req.URL.Path,"/")
	fmt.Fprint(res, name[len(name)-1])
}

func main(){
	http.HandleFunc("/", PathN)
	http.listenAndServe(":8080", nil)
}