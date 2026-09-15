package helper

import (
	"encoding/json"
	"os"
)

func SaveUsers(users []any) error {
	data, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("users.json", data, 0644)
}

func ReadUsers() ([]any, error) {
	data, err := os.ReadFile("users.json")

	if os.IsNotExist(err) {
		return []any{}, nil
	}

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return []any{}, nil
	}

	var users []any
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}
