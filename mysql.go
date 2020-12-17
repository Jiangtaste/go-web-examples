package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// DB 数据库连接池实例
var DB *sql.DB

// Open 链接数据库
func Open() {
	db, err := sql.Open("mysql", "dev:Fqk.V;@kU7yF7Afw+YBu@(localhost:3306)/go_web?parseTime=true")

	if err != nil {
		log.Fatalf("Error to connect db %v\n", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error to ping db. %v\n", err)
	}

	DB = db
}

// CreateUserTable 创建用户表
func CreateUserTable(db *sql.DB) error {
	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`

	// Executes the SQL query in our database. Check err to ensure there was no error.
	if _, err := db.Exec(query); err != nil {
		// log.Fatalf("Error to create user table %v\n", err)
		return err
	}
	return nil
}

// InsertUser 插入用户，返回uid
func InsertUser(db *sql.DB) (int64, error) {

	username := "johndoe"
	password := "secret"
	createdAt := time.Now()

	result, err := db.Exec(`INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`, username, password, createdAt)
	if err != nil {
		// log.Fatalf("Error to insert user %v\n", err)
		return -1, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		// log.Fatalf("Error to get userID %v\n", err)
		return -1, err
	}

	fmt.Printf("uid: %d\n", userID)

	return userID, nil
}

// FetchUser 获取用户
func FetchUser(uid int64, db *sql.DB) (*User, error) {

	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := `SELECT id, username, password, created_at FROM users WHERE id = ?`
	err := db.QueryRow(query, uid).Scan(&id, &username, &password, &createdAt)

	if err != nil {
		// log.Fatalf("Error to get user %v\n", err)
		return nil, err
	}

	fmt.Printf("uid: %d, username: %s, password: %s, createdAt: %v\n", id, username, password, createdAt)

	return &User{ID: id, Username: username, Password: password, CreatedAt: createdAt}, nil
}

// User 用户模型
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

// GetUsers 获取所有用户
func GetUsers(db *sql.DB) ([]User, error) {

	rows, err := db.Query(`SELECT id, username, password, created_at FROM users`)

	if err != nil {
		// log.Fatalf("Error to get users %v\n", err)
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var u User

		err := rows.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt)
		if err != nil {
			// log.Fatalf("Error to scan user %v\n", err)
			return nil, err
		}

		users = append(users, u)
	}

	err = rows.Err()

	if err != nil {
		// log.Fatalf("Error to get users %v\n", err)
		return nil, err
	}

	fmt.Printf("%#v\n", users)

	return users, nil
}

// DeleteUser 删除用户
func DeleteUser(id int64, db *sql.DB) error {

	_, err := db.Exec(`DELETE FROM users WHERE id = ?`, id)

	return err
}

// handlers

func createUserTable(w http.ResponseWriter, r *http.Request) {

	if err := CreateUserTable(DB); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "User Table Created.")

}

func createUser(w http.ResponseWriter, r *http.Request) {

	userID, err := InsertUser(DB)

	if err != nil {
		http.Error(w, "Failed to create user.", http.StatusBadRequest)
	}

	fmt.Fprintf(w, "User ID: %d\n", userID)
}

func fetchUser(w http.ResponseWriter, r *http.Request) {

	uid := parseUID(w, r)

	user, err := FetchUser(uid, DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(user)
}

func fetchUsers(w http.ResponseWriter, r *http.Request) {

	users, err := GetUsers(DB)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	json.NewEncoder(w).Encode(users)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {

	uid := parseUID(w, r)

	if err := DeleteUser(uid, DB); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Fprintf(w, "delete user success. uid: %d\n", uid)
}

func parseUID(w http.ResponseWriter, r *http.Request) int64 {

	vars := mux.Vars(r)
	id := vars["id"] // the user id slug

	uid, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	return uid
}
