package repositories

import (
	"database/sql"
	"repo/internal/domain/entities"
	"repo/internal/domain/repositories"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) repositories.UserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Create(user *entities.User) error {
	_, err := r.db.Exec("INSERT INTO users (name) VALUES (?)", user.Name)
	return err
}

func (r *MySQLUserRepository) GetByID(id int) (*entities.User, error) {
	var user entities.User
	err := r.db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) Update(user *entities.User) error {
	_, err := r.db.Exec("UPDATE users SET name = ? WHERE id = ?", user.Name, user.ID)
	return err
}
