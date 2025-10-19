package main

/*
 * Gaurav Sablok
 * codeprog@icloud.com
 */

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Message string
}

func main() {
	tmpl := template.Must(template.ParseFiles("templates/web.tmpl"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{}
		if r.Method == http.MethodPost {
			r.ParseForm()
			userInput := r.FormValue("userInput")
			if userInput != "" {
				data.Message = userInput
			}
		}

		err := tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			return
		}
	})
	http.ListenAndServe(":8080", nil)
}
