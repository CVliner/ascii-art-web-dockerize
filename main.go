package main

import (
	"fmt"
	"html/template"
	"log"
	"main/asciiart"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running @ http:localhost:4747")
	log.Println("OK(200)")
	err := http.ListenAndServe(":4747", nil)
	if err != nil {
		log.Fatal("ListenAndServe: Error", http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/ascii.html")
	if err != nil {
		fmt.Println("Internal Server Error")
		Error500(w)
		return
	}
	if r.URL.Path != "/" {
		fmt.Println("Not applicable path")
		Error404(w)
		return
	}
	switch r.Method {
	case "GET":
		template.Execute(w, nil)
	case "POST":
		text, banner := r.FormValue("text"), r.FormValue("banner")
		output, err1 := asciiart.Generate(text, banner)
		if err1 == false && output == "2" {
			fmt.Println("Internal Server Error!")
			Error500(w)
			return
		}
		if err1 == false && output == "1" {
			fmt.Println("Not applicable Ascii")
			Error400(w)
			return
		}
		err := template.Execute(w, output)
		if err != nil {
			fmt.Println(400, 2)
			Error400(w)
			return
		}
	}
}

func Error500(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	t, err := template.ParseFiles("templates/500-error.html")
	if err != nil {
		log.Fatalf("Error happened in parsing file. Error: %s", err)
		return
	}
	t.Execute(w, nil)
}

func Error404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	t, err := template.ParseFiles("templates/404-error.html")
	if err != nil {
		log.Fatalf("Error happened in parsing file. Error: %s", err)
		return
	}
	t.Execute(w, nil)
}

func Error400(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	t, err := template.ParseFiles("templates/400-error.html")
	if err != nil {
		log.Fatalf("Error happened in parsing file. Error: %s", err)
		return
	}
	t.Execute(w, nil)
}
