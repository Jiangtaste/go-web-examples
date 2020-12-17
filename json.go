package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JUser 用户模型
type JUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	var user JUser
	json.NewDecoder(r.Body).Decode(&user)

	fmt.Fprintf(w, "%s %s is %d years old!", user.Firstname, user.Lastname, user.Age)
}

func encodeHandler(w http.ResponseWriter, r *http.Request) {
	peter := JUser{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}
	json.NewEncoder(w).Encode(peter)
}
