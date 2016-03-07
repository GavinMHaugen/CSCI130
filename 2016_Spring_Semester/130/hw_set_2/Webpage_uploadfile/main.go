package main 

import(
	"fmt"
	"io"
	"net/http"
	"html/template"
	"io/ioutil"
	"strings"
)

var tpl *template.Template

func init() {
	tpl, _ = template.ParseFiles("index.gothml")
}

func GenPage(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, `<!DOCTYPE html>
					<html>
					<body>
						<form method = "POST"  enctype="multipart/form-data">
							<input type="file" name="inputField"><br>
							<input type ="submit">
						</form>
					</body>
					</html>`)

	if req.Method == "POST" {
		key := "name"
		_, hdr, err := req.FormFile(key)
		if err != nil {
			fmt.Println(err)
		}

		rdr, err2 := hdr.Open()
		if err2 != nil{
			fmt.Println(err2)
		}

		io.Copy(res, rdr)
	}

	tpl.Execute(res, nil)
	hdr, _ := ioutil.ReadFile("file.txt")
	fmt.Fprint(res,strings.Split(string(hdr), "\n"))
}

func main(){
	http.HandleFunc("/", GenPage)
	fmt.Println("running...")
	http.ListenAndServe(":8080", nil)
}