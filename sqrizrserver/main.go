package main

import (
	"net/http"
	"io"
	"log"
	//"sqrizr/sqrizrlib"
	//"mime/multipart"
)

var indexHtml = `
<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<title>SQRIZR (square-izer)</title>
		<style>
			body {
				padding: 0;
				margin: 0;
				font-size: 14px;
			}

			header{
				font-size: 150%;
			}
		</style>
	</head>
	<body>
		<header>Sqrizr - turn your images into squares!</header>
		<form method="post" enctype="multipart/form-data">
			<label for="image">Choose an image:</label> <input type="file" name="image">
			<input type="submit" name="submit">
		</form>
	</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET" {
		io.WriteString(w, indexHtml)
	} else {
		file, fileHeader, err := r.FormFile("image")
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer file.Close()
		io.WriteString(w, "Uploaded: " + fileHeader.Filename)
	}
}

func main (){
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}


