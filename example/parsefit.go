/**
* From Go-MultipleFileUpload @
*
* https://github.com/sanatgersappa/Go-MultipleFileUpload/blob/master/app.go
**/
package main

//$ cd c:go/src/github.com/jezard/fit/example
//$ go install
//$ example > ../sample_output.txt

import (
	"fmt"
	"github.com/jezard/fit"
	"html/template"
	"io"
	"net/http"
	"os"
)

var templatePath = "tmpl/"
var crc uint16

//Compile templates on start
var templates = template.Must(template.ParseFiles(templatePath + "uploader.html"))

func main() {
	http.HandleFunc("/upload/", UploadHandler)

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}

//Display the named template
func display(w http.ResponseWriter, tmpl string, data interface{}) {
	templates.ExecuteTemplate(w, tmpl+".html", data)
}
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	//GET displays the upload form.
	case "GET":
		display(w, "uploader", nil)

	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//parse the multipart form in the request
		err := r.ParseMultipartForm(1000000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//get a ref to the parsed multipart form
		m := r.MultipartForm

		//get the *fileheaders
		files := m.File["myfiles"]
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.
			dst, err := os.Create("C:/Users/Administrator/git-projects/jps-frontend/uploads/" + files[i].Filename)
			defer dst.Close()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			fmt.Printf("%v\n", files[i].Header)
			fit.Parse("C:/Users/Administrator/git-projects/jps-frontend/uploads/"+files[i].Filename, true)
		}
		//display success message.
		display(w, "uploader", "Upload successful.")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
