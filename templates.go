package main

import (
	"html/template"
	"net/http"
)

// Todo 任务模型
type Todo struct {
	Title string
	Done  bool
}

// TodoPageData 页面数据模型
type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

// templates
func todoHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("layout.html"))

	data := TodoPageData{
		PageTitle: "My TODO list",
		Todos: []Todo{
			{Title: "Task 1", Done: false},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: true},
		},
	}

	tmpl.Execute(w, data)
}
