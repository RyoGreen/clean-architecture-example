package repo

import (
	"clean-architecture/model"
	"database/sql"
)

type IUserRepo interface {
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &userRepo{db}
}

const (
	getUserQuery    = "SELECT * FROM user WHERE email = $1;"
	createUserQuery = "INSERT INTO user (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4);"
)

func (u userRepo) GetUserByEmail(email string) (*model.User, error) {
	var user = model.User{}
	if err := u.db.QueryRow(getUserQuery, email).Scan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepo) CreateUser(user *model.User) error {
	_, err := u.db.Exec(createUserQuery, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}
