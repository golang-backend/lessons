package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	PostgresUser     = "postgres"
	PostgresPassword = "postgres"
	PostgresDatabase = "postgres"
	PostgresHost     = "localhost"
	PostgresPort     = "5432"
)

type User struct {
	Id        int
	Name      string
	Phone     string
	CreatedAt *time.Time
}

func main() {
	psqlConnString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost, PostgresPort, PostgresUser, PostgresPassword, PostgresDatabase)
	// host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable
	conn, err := sql.Open("postgres", psqlConnString)
	if err != nil {
		fmt.Printf("failed to connect postgres: %v\n", err)
		return
	}

	if err := conn.Ping(); err != nil {
		fmt.Printf("failed to ping postgres: %v\n", err)
		return
	}

	fmt.Println("Postgres connected successfully!!!")

	// Migrate
	query, err := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS products(
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		price DECIMAL(10, 2) DEFAULT 0.00,
		created_at TIMESTAMP DEFAULT NOW()
	) 
	`)
	_, err = query.Exec()
	if err != nil {
		fmt.Printf("could not create products table: %v\n", err)
		return
	}
	fmt.Println("products table created successfully!!!")

	// Query a user by ID
	id := 1
	raw := conn.QueryRow("SELECT id, name, phone, created_at FROM users WHERE id = $1;", id)
	user := User{}
	err = raw.Scan(&user.Id, &user.Name, &user.Phone, &user.CreatedAt)
	if err != nil {
		fmt.Printf("failed to scan user: %v\n", err)
		return
	}
	fmt.Println("execute select 1 row")
	// userByte, err := json.MarshalIndent(user, "", "  ")
	// if err != nil {
	// 	fmt.Printf("failed to marshal user: %v\n", err)
	// 	return
	// }
	// fmt.Printf("User JSON: %s\n", string(userByte))

	users := []User{}

	raws, err := conn.Query("SELECT id, name, phone, created_at FROM users")
	if err != nil {
		fmt.Printf("failed to run query: %v\n", err)
		return
	}
	if raws.Err() != nil {
		fmt.Printf("failed to run query: %v\n", err)
		return
	}
	for raws.Next() {
		tmp := User{}
		err = raws.Scan(&tmp.Id, &tmp.Name, &tmp.Phone, &tmp.CreatedAt)
		if err != nil {
			fmt.Printf("could not scan raws: %v\n", err)
			return
		}
		users = append(users, tmp)
	}
	raws.Close()
	fmt.Println("execute select more rows")

	// usersByte, err := json.MarshalIndent(users, "", "   ")
	// if err != nil {
	// 	fmt.Printf("failed to marshal users: %v\n", err)
	// 	return
	// }
	// fmt.Println(string(usersByte))

	_, err = conn.Exec("UPDATE users SET phone='+998913450909' WHERE id = $1", id)
	if err != nil {
		fmt.Printf("could not update users phone: %v\n", err)
		return
	}
	fmt.Println("execute update 1 row")

	_, err = conn.Exec("DELETE FROM users WHERE id = $1", 61)
	if err != nil {
		fmt.Printf("could not delete user: %v\n", err)
		return
	}
	fmt.Println("execute delete 1 row")
	newUser := User{}
	err = conn.QueryRow("INSERT INTO users(id, name, phone) VALUES($1, $2, $3) RETURNING id, name, phone, created_at",
		62, "Karimjon", "+998982348989",
	).Scan(&newUser.Id, &newUser.Name, &newUser.Phone, &newUser.CreatedAt)
	if err != nil {
		fmt.Printf("could not insert user: %v\n", err)
		return
	}

	newByte, err := json.MarshalIndent(newUser, "", "   ")
	if err != nil {
		fmt.Printf("could marshal new user: %v\n", err)
		return
	}
	fmt.Println(string(newByte))
}
