package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	PostgresHost = "localhost"
	PostgresPort = "5432"
	PostgresUser = "postgres"
	PostgresPass = "postgres"
	PostgresDb   = "postgres"
)

type User struct {
	Id        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Phone     string     `json:"phone"`
	CreatedAt *time.Time `json:"created_at"`
}

func main() {
	// open connection with postgres
	db, err := connPostgres()
	if err != nil {
		fmt.Println(err)
		return
	}

	// migrate tables
	err = migrate(db)
	if err != nil {
		fmt.Println("failed to migrate table: ", err)
		return
	}

	storage := NewPostgres(db)

	route := gin.Default()

	route.POST("/users", storage.userCreate)
	route.GET("/users/:user_id", storage.getUser)
	route.PUT("/users/:user_id" )
	route.DELETE("/users/:user_id" )

	if err := route.Run(":8081"); err != nil {
		fmt.Println("could not run on port: 8080, err:", err)
		return
	}
}

func (h *Conn) userCreate(c *gin.Context) {
	var body User
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user User
	err := h.db.QueryRow(`
	INSERT INTO users(
		first_name, last_name, phone
	) VALUES($1, $2, $3) 
	 RETURNING id, first_name, last_name, phone, created_at`,
		body.FirstName,
		body.LastName,
		body.Phone).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.CreatedAt,
	)
	if err != nil {
		fmt.Println("could not create user: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *Conn) getUser(c *gin.Context) {
	userIdStr := c.Param("user_id")

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user User
	err = h.db.QueryRow("SELECT id, first_name, last_name, phone, created_at FROM users WHERE id = $1",
		userId).Scan(
		&user.Id,
		&user.FirstName,
		&user.LastName,
		&user.Phone,
		&user.CreatedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

type Conn struct {
	db *sql.DB
}

func NewPostgres(conn *sql.DB) *Conn {
	return &Conn{
		db: conn,
	}
}

func connPostgres() (*sql.DB, error) {
	psqlConnStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		PostgresHost,
		PostgresPort,
		PostgresUser,
		PostgresPass,
		PostgresDb,
	)

	conn, err := sql.Open("postgres", psqlConnStr)
	if err != nil {
		fmt.Printf("failed to open postgres: %v\n", err)
		return nil, err
	}
	if err := conn.Ping(); err != nil {
		fmt.Printf("failed to check postgres ping: %v\n", err)
		return nil, err
	}

	return conn, nil
}

func migrate(conn *sql.DB) error {
	// prepare create table query
	pre, err := conn.Prepare(`
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255),
		phone VARCHAR(20),
		created_at TIMESTAMP DEFAULT NOW()
	);
	`)
	if err != nil {
		return err
	}
	// execute create table query
	_, err = pre.Exec()
	if err != nil {
		return err
	}
	return nil
}
