package models

import "time"

type User struct {
	Id        int        `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Phone     string     `json:"phone"`
	Balance   float64    `json:"balance"`
	CreatedAt *time.Time `json:"created_at"`
}
