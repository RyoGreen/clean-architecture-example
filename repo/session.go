package repo

import (
	"clean-architecture/cookie"
	"clean-architecture/model"
	"database/sql"
	"net/http"
)

const (
	getSessionQuery = "SELECT users.* FROM sessions INNER JOIN users ON sessions.user_id = users.id WHERE sessions.s_id = $1;"
)

type ISessionRepo interface {
	GetSessoionByID(sID string) (*model.User, error)
	GetUserSid(getCookie *http.Cookie) (string, error)
}

type sessionRepo struct {
	db *sql.DB
	cm *cookie.CookieManager
}

func NewSessionRepo(db *sql.DB, cm *cookie.CookieManager) ISessionRepo {
	return &sessionRepo{db, cm}
}

func (u sessionRepo) GetSessoionByID(sID string) (*model.User, error) {
	var user model.User
	if err := u.db.QueryRow(getSessionQuery, sID).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u sessionRepo) GetUserSid(getCookie *http.Cookie) (string, error) {
	return u.cm.GetUserSID(getCookie)
}
