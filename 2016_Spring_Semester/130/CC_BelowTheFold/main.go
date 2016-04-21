package main

import (
	"net/http" 
	"html/template"
	"io/ioutil"
	"google.golang.org/appengine"
    "google.golang.org/cloud/storage"
	"golang.org/x/net/context"
)

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
	var outputString[]string
	//Based on a given directory pull out all files within the bucket on GCS and append them to the Images slice.
	query := &storage.Query{Delimiter: directory}
	objs, _ := d.bucket.List(d.ctx, query)
	for _, obj := range objs.Results {
		outputString = append(outputString,obj.Name)
	}
	return PageData{Images:outputString}
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
}