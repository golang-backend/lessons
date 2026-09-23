package storage

import (
	"database/sql"
	"mini-project/models"
)

type Storage struct {
	DB *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		DB: db,
	}
}

func (s *Storage) CreateUser(user models.User) (*models.User, error) {
	query := "INSERT INTO users(first_name, last_name, phone) VALUES($1, $2, $3) RETURNING *"
	var response models.User
	err := s.DB.QueryRow(query, user.FirstName, user.LastName, user.Phone).
		Scan(&response.Id,
			&response.FirstName,
			&response.LastName,
			&response.Phone,
			&response.Balance,
			&response.CreatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (s *Storage) CreateTask(task models.Task) (*models.Task, error) {
	query := "INSERT INTO tasks(user_id, title, description, status) VALUES($1, $2, $3, $4) RETURNING *"
	var response models.Task
	task.Status = "new"
	err := s.DB.QueryRow(query, task.UserId, task.Title, task.Description, task.Status).
		Scan(&response.Id,
			&response.UserId,
			&response.Title,
			&response.Description,
			&response.Status,
			&response.CreatedAt,
			&response.UpdatedAt,
		)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
