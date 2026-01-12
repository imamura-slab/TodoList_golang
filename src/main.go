package main

import (
	//"fmt"
	"log"
	"text/template"
	"net/http"
)

// トップページハンドラ (/)
func topPageHandler(w http.ResponseWriter, r *http.Request) {
	filename := "templates/index.html"
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}


// registerページハンドラ (/register)
func registerPageHandler(w http.ResponseWriter, r *http.Request) {
	filename := "templates/register.html"
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
	}
	
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal(err)
	}
}


func main() {
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./templates/css"))))

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(topPageHandler))
	mux.Handle("/register", http.HandlerFunc(registerPageHandler))
	
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", mux)
}

