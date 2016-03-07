package main 

import(
	"net/http"
	"fmt"
)

func PageMaker(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, <!DOCTYPE html>
					<html>
					<head>
						<title></title>
					</head>

					<body>
						<form>
							Input name: <br>
							<input type="text" name="name"><br>
						</form>
					</body>
					</html>)

	fmt.Fprint(res,req.FormValue("name"))
}

func main(){
	http.HandleFunc("/", PageMaker)
	http.ListenAndServe(":8080", nil)
}