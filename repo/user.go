package repo

import (
	"clean-architecture/model"
	"database/sql"
)

const (
	getUserQuery       = "SELECT * FROM users WHERE email = $1;"
	createUserQuery    = "INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4);"
	getSessionQuery    = "SELECT users.* FROM sessions INNER JOIN users ON sessions.user_id = users.id WHERE sessions.id = $1;"
	insertSessionQuery = "INSERT INTO sessions (s_id, user_id, expired, created_at) VALUES ($1, $2, $3, $4);"
)

type IUserRepo interface {
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
	GetSessoionByID(sID string) (*model.User, error)
	CreateSession(session *model.Session) error
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) IUserRepo {
	return &userRepo{db}
}

func (u userRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.QueryRow(getUserQuery, email).Scan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepo) CreateUser(user *model.User) error {
	_, err := u.db.Exec(createUserQuery, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func (u userRepo) GetSessoionByID(sID string) (*model.User, error) {
	var user model.User
	if err := u.db.QueryRow(getUserQuery, sID).Scan(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepo) CreateSession(session *model.Session) error {
	_, err := u.db.Exec(insertSessionQuery, session.SID, session.UserID, session.Expired, session.CreatedAt)
	return err
}
