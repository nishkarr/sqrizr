package main

import (
	"net/http"
	"io"
	"log"
	"fmt"
	"strings"

	"github.com/nishkarr/sqrizr/sqrizrlib"
)

var indexHtml = `
<!doctype html>
<html>
	<head>
		<meta charset="utf-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
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

		// this is messy, we need to know if the processing completed successfully so we 
		// can set the request headers. Request headers must be set before writing to ResponseWriter!

		headers := w.Header()
		headers.Set("Content-Type", "image/png")
		headers.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="sqr_%s"`, convertFileName(fileHeader.Filename)))

		_, _, err = sqrizrlib.Sqrize(file, w) // we rely on the sqrizrlib not writing anything to w on failure!
		if err != nil {
			headers.Set("Content-Type", "text/plain")
			headers.Del("Content-Disposition")
			io.WriteString(w, "Error decoding: " + fileHeader.Filename + ": " + err.Error())	
		} else{
			//headers := w.Header()
			//headers.Set("Content-Type", "image/jpg")
			//headers.Set("Content-Disposition", fmt.Sprintf(`attachment; filename="sqr_%s"`, fileHeader.Filename))
		}
	}
}

func convertFileName(file string) string {
	r := strings.NewReplacer(".jpg", ".png", ".gif", ".png")
	return r.Replace(file)
}

func main (){
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}


