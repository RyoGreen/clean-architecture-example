package repo

import (
	"clean-architecture/cookie"
	"clean-architecture/model"
	"database/sql"
	"net/http"
)

const (
	getUserQuery       = "SELECT * FROM users WHERE email = $1;"
	createUserQuery    = "INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4);"
	createSessionQuery = "INSERT INTO sessions (s_id, user_id, expired, created_at) VALUES ($1, $2, $3, $4);"
)

type IUserRepo interface {
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) error
	SetSID(val string) (*http.Cookie, error)
	CreateSession(session *model.Session) error
}

type userRepo struct {
	db *sql.DB
	cm *cookie.CookieManager
}

func NewUserRepo(db *sql.DB, cm *cookie.CookieManager) IUserRepo {
	return &userRepo{db, cm}
}

func (u userRepo) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	if err := u.db.QueryRow(getUserQuery, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u userRepo) CreateUser(user *model.User) error {
	_, err := u.db.Exec(createUserQuery, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	return err
}

func (u userRepo) SetSID(val string) (*http.Cookie, error) {
	return u.cm.SetSID(val)
}

func (u userRepo) CreateSession(session *model.Session) error {
	_, err := u.db.Exec(createSessionQuery, session.SID, session.UserID, session.Expired, session.CreatedAt)
	return err
}
