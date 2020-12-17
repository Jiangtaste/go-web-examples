package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// hello world
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	// })

	

	// init db
	Open()

	// router
	routers := InitRouters()

	// password hash
	password := "secret"
	hash, err := HashPassword(password)

	if err != nil {
		log.Fatalf("Error to hash password. %v\n", err)
	}

	fmt.Println("Password: ", password)
	fmt.Println("Hash: ", hash)

	match := CheckPasswordHash(password, hash)

	fmt.Println("Match: ", match)

	http.ListenAndServe(":8080", routers)
}
