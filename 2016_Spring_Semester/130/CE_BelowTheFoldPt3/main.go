package main

import (
	"net/http" 
	"html/template"
	"io/ioutil"
	"io"
	"google.golang.org/appengine"
    "google.golang.org/cloud/storage"
	"golang.org/x/net/context"
	"mime/multipart"
	"crypto/sha1"
	"strings"
	"fmt"
)

func getSha(src multipart.File) string {
	h := sha1.New()
	io.Copy(h, src)
	return fmt.Sprintf("%x", h.Sum(nil))
}

var file1 string
var bucket = "essys-bucket"

type PageData struct{
	Images []string
}


type SessionData struct{
	res http.ResponseWriter
	req *http.Request
	//Current Bucket
	bucket *storage.BucketHandle
	//Current Client
	client *storage.Client
	//Current Context
	ctx context.Context
}

func PopulateImages(d SessionData,directory string)(PageData){
	directory = d.req.FormValue("key")
	var outputString[]string
	//Based on a given directory pull out all files within the bucket on GCS and append them to the Images slice.
	query := &storage.Query{Prefix: directory}
	objs, _ := d.bucket.List(d.ctx, query)
	for _, obj := range objs.Results {
		outputString = append(outputString,obj.Name)
	}
	return PageData{Images:outputString}
}

func store(res http.ResponseWriter, req *http.Request){
	fmt.Fprintf(res,`<!DOCTYPE html>
					<html lang="en">
					<head>
						<meta charset="UTF-8">
						<title></title>
					</head>
					<body>
					<form method="POST" enctype="multipart/form-data">
						<input type="file" name="data">
						<input type="submit">
					</form>
					</body>
					</html>`)

	if req.Method == "POST"{
		file, hdr , _ := req.FormFile("data")
		defer file.Close()
		uploadFile(req,file,hdr,req.FormValue("key"))
	}
}

func uploadFile(req *http.Request,file multipart.File,hdr *multipart.FileHeader, dest string){
	ext := hdr.Filename[strings.LastIndex(hdr.Filename, ".")+1:]
	
	name:= dest + "/" + getSha(file) + "." + ext
	
	file.Seek(0,0)
	ctx := appengine.NewContext(req)
	putFile(ctx,name,file)
}

func putFile(ctx context.Context, name string, rdr io.Reader){
	client, _ := storage.NewClient(ctx)
	defer client.Close()
	
	writer := client.Bucket(bucket).Object(name).NewWriter(ctx)
	
	writer.ACL = []storage.ACLRule{
		{storage.AllUsers, storage.RoleReader},
	}
	io.Copy(writer, rdr)
	writer.Close()
}

func index(res http.ResponseWriter, req *http.Request){
	// Make current context based off the user request.
	ctx := appengine.NewContext(req)
	// Make current client based off the current context.
	client, _ := storage.NewClient(ctx)
	// Construct the session data to be passed around.
	// Get image paths.
	obj := PopulateImages(SessionData{res:res,req:req,ctx:ctx,client:client,bucket:client.Bucket(bucket)},"")
	// Execute the template ranging over the images gathered.
	t,_ := template.New("name").Parse(file1)
	t.Execute(res, obj)
}

func init(){
	// Read template file.
	temp, _ := ioutil.ReadFile("index.html")
	// Save template file globally.
	file1 = string(temp)
	// Handle css files.
	http.Handle("/css/",http.StripPrefix("/css",http.FileServer(http.Dir("css"))))
	// Handle img files.
	http.Handle("/img/",http.StripPrefix("/img",http.FileServer(http.Dir("img"))))
	// Handle main page.
	http.HandleFunc("/", index)
	http.HandleFunc("/store/", store)
}