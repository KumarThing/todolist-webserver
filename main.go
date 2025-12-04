package main

import (
	
	"html/template"
	"net/http"
	"strconv"
	"log"
)

var tmpl = template.Must(template.ParseFiles("template/index.html"))

var items []string

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){

		if r.Method == http.MethodPost{
			r.ParseForm()
			newItem := r.FormValue("newitem")

			if newItem != " " {
				items = append(items, newItem)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		data := struct{
			Items []string
		}{
			Items: items,
		}

		tmpl.Execute(w, data)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request){

		if r.Method == http.MethodPost{
			r.ParseForm()
			istr := r.FormValue("index")

			i,err := strconv.Atoi(istr)
			if err == nil && i >= 0 && i < len(items) {
				items = append(items[:i], items[i+1:]... )
			}
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	})


	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}