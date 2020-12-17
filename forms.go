package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// ContactDetails 联系信息详情
type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func contactDetailsHandler(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("forms.html"))

	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}

	// do something with details

	_ = details

	fmt.Printf("Form Data: %v\n", details)

	tmpl.Execute(w, struct{ Success bool }{true})
}
