package repository

import (
	"database/sql"
	"kodinggo/internal/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) model.IUserRepository {
	return &UserRepository{db: db}
}

func (u *UserRepository) Create(user model.User) error {
	_, err := u.db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepository) Login(username string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow("SELECT id, username, password FROM users WHERE username=?", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) FindByUsername(username string) (model.User, error) {
	var user model.User
	err := u.db.QueryRow("SELECT id, username, password FROM users WHERE username=?", username).Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}
