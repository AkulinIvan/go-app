package main

import (
	"html/template"
	"net/http"
	"fmt"
	"github.com/AkulinIvan/go-app/book/uploadFile"
)

type User struct {
	name      string
	last_name string
}




// func cabinet(w http.ResponseWriter, r *http.Request) {
// 	...
// }

func home_page(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/index.html")
	tmpl.Execute(w, r)
}

func HandleRequest() {
	http.HandleFunc("/", home_page)
	// http.HandleFunc("/about/", about_page)
	// http.HandleFunc("/contacts/", contacts)
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":8080", nil)
}
func main() {
	HandleRequest()
}
