package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// API - Application Programming Interface
// REST API - Reprentation State Transfer |  Graphql - API architecture
// /users /users/:id -> Rest Architecture
// Resource vs Endpoint
//   User        /users
//   Product     /products
// Rest Full 6-prinsple
//

type User struct {
	Id      int       `json:"id"`
	Name    string    `json:"name"`
	Age     int       `json:"age"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
}

func main() {
	route := http.NewServeMux()

	route.HandleFunc("POST /users", createUser)
	route.HandleFunc("GET /users", getUsers)
	route.HandleFunc("PATCH /users/{id}", updateUser)
	route.HandleFunc("DELETE /users/{id}", deleteUser)

	fmt.Println("Server is running")
	if err := http.ListenAndServe(":8080", enableCORS(route)); err != nil {
		fmt.Printf("Could not run server: %v\n", err)
		return
	}
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	users, err := readUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.Id = len(users) + 1
	user.Created = time.Now()

	users = append(users, user)

	if err := saveUsers(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	Middleware(r)
	users, err := readUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func Middleware(r *http.Request) {
	fmt.Println("Method: ", r.Method)
	fmt.Println("Path: ", r.URL)
	fmt.Println("Time: ", time.Now().Format(time.DateTime))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	var req User
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	users, err := readUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	found := false

	for i := range users {
		if users[i].Id == id {
			if req.Name != "" {
				users[i].Name = req.Name
			}
			if req.Email != "" {
				users[i].Email = req.Email
			}
			if req.Age != 0 {
				users[i].Age = req.Age
			}
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := saveUsers(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User updated",
	})
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	users, err := readUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result []User
	found := false

	for _, user := range users {
		if user.Id == id {
			found = true
			continue
		}
		result = append(result, user)
	}

	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	if err := saveUsers(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func saveUsers(users []User) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("users.json", data, 0644)
}

func readUsers() ([]User, error) {
	data, err := os.ReadFile("users.json")

	if os.IsNotExist(err) {
		return []User{}, nil
	}

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []User{}, nil
	}

	var users []User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Brauzer yuborgan tekshiruv (preflight) so'roviga OK qaytaramiz
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
