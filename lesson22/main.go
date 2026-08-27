package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	user.Created = time.Now()
	err = writeUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"msg":  "User Created",
		"user": user,
	})
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	users, err := readUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("id")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Id must be number", http.StatusBadRequest)
		return
	}
	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	users, err := readUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	for i, _ := range users {
		if users[i].Id == userId {
			users[i].Name = user.Name
			users[i].Email = user.Email
			users[i].Age = user.Age
		}
	}
	err = writeUsers(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"msg":  "User Updated",
		"user": user,
	})

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	
}

func writeUser(user User) error {
	file, err := os.OpenFile("users.json", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("could not open file: %v\n", err)
		return err
	}
	defer file.Close()

	userByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("could not read file: %v\n", err)
		return err
	}
	users := []User{}
	if len(userByte) > 0 {
		err = json.Unmarshal(userByte, &users)
		if err != nil {
			fmt.Printf("could not unmarshal users: %v\n", err)
			return err
		}
	}

	users = append(users, user)
	usersM, err := json.MarshalIndent(users, "", "   ")
	if err != nil {
		fmt.Printf("could not marshal users: %v\n", err)
		return err
	}

	if err = file.Truncate(0); err != nil {
		fmt.Printf("could not truncate users.json: %v\n", err)
		return err
	}
	if _, err = file.Seek(0, 0); err != nil {
		fmt.Printf("could not seek users.json: %v\n", err)
		return err
	}

	_, err = file.Write(usersM)
	if err != nil {
		fmt.Printf("could not write users: %v\n", err)
		return err
	}
	return nil
}

func writeUsers(users []User) error {
	file, err := os.OpenFile("users.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Printf("could not open users json: %v\n", err)
		return err
	}
	defer file.Close()
	usersByte, err := json.MarshalIndent(users, "", "   ")
	if err != nil {
		fmt.Printf("could not marshal users: %v\n", err)
		return err
	}
	if err = file.Truncate(0); err != nil {
		fmt.Printf("could not truncate users.json: %v\n", err)
		return err
	}
	if _, err = file.Seek(0, 0); err != nil {
		fmt.Printf("could not seek users.json: %v\n", err)
		return err
	}
	_, err = file.Write(usersByte)
	if err != nil {
		fmt.Printf("could not write users: %v\n", err)
		return err
	}
	return nil
}

func readUsers() ([]User, error) {
	file, err := os.Open("users.json")
	if err != nil {
		fmt.Printf("could not open users.json: %v\n", err)
		return []User{}, err
	}
	defer file.Close()
	usersByte, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("could not read users.json: %v\n", err)
		return []User{}, err
	}
	if len(usersByte) == 0 {
		return []User{}, nil
	}
	var users []User
	err = json.Unmarshal(usersByte, &users)
	if err != nil {
		fmt.Printf("could not unmarshal users.json: %v\n", err)
		return []User{}, err
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
