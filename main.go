package main

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

type Todo struct {
	Item string
	Done bool
}
type PageData struct {
	Title string
	Todos []Todo
}
func todo(w http.ResponseWriter, r *http.Request){
	data := PageData{
		Title: "SIMPLE TODOLIST APPLICATION",
		Todos: []Todo{
			{"install go",true},
			{"learn go",false},
			{"i like this video",true},
		},
	}
	tmpl.Execute(w,data)
}

func main(){
	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/index.gohtml"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/",http.StripPrefix("/static/",fs))
	mux.HandleFunc("/todo",todo)

	http.ListenAndServe(":9091",mux)
}
