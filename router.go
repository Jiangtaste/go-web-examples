package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// InitRouters 初始化路由
func InitRouters() *mux.Router {

	r := mux.NewRouter()

	// HTTP Server
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"] // the book title slug
		page := vars["page"]   // the page
		fmt.Fprintf(w, "You've requested the books: %s on page %s\n", title, page)
	})

	// mysql database
	r.HandleFunc("/createUserTable", createUserTable)
	r.HandleFunc("/users/", createUser).Methods("POST")
	r.HandleFunc("/users/", fetchUsers).Methods("GET")
	r.HandleFunc("/users/{id}", fetchUser).Methods("GET")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// templates
	r.HandleFunc("/templates", todoHandler)

	r.HandleFunc("/forms", contactDetailsHandler)

	// basic middleware
	r.HandleFunc("/foo", logging(foo))
	r.HandleFunc("/bar", logging(bar))

	// advanced middleware
	r.HandleFunc("/middleware", Chain(MiddlewareHello, Method("GET"), Logging()))

	// sessions
	r.HandleFunc("/secret", secret)
	r.HandleFunc("/login", login)
	r.HandleFunc("/logout", logout)

	// json
	r.HandleFunc("/decode", decodeHandler)
	r.HandleFunc("/encode", encodeHandler)

	// websockets
	r.HandleFunc("/echo", onMessage)
	r.HandleFunc("/ws", getWSPage)

	return r
}
