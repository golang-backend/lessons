package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	route := http.NewServeMux()

	route.HandleFunc("/health", healthHandler)

	// users
	route.HandleFunc("/users", userHandler)

	fmt.Printf("Server is running\n")
	if err := http.ListenAndServe(":8080", route); err != nil {
		log.Fatalf("Server could not run on 8080 port: %v\n", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Server is running",
	})
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPost:
		createUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	}
}

type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Phone   string `json:"phone"`
	Balance int    `json:"balance"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
	user := User{
		Name:    "Umidjon",
		Age:     17,
		Phone:   "+998911234567",
		Balance: 1000,
	}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// request dan kelgan body ni o'qib user struct ga yozish
	json.NewDecoder(r.Body).Decode(&user)

	// response qaytarish json formatida
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"msg":  "Users yaratildi",
		"name": user.Name,
	})
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	// request dan kelgan body ni o'qib user struct ga yozish
	json.NewDecoder(r.Body).Decode(&user)

	// response qaytarish json formatida
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"msg":  "User ma'lumoti yangilandi",
		"user": user,
	})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	// response qaytarish json formatida
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"msg": "User ma'lumoti o'chirildi!!!",
	})
}
