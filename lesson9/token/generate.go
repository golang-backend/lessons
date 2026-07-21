package token

import "github.com/google/uuid"

func Generate() string {
	id := uuid.NewString()
	return id
}


