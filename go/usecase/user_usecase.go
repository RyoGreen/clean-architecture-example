package usecase

import (
	"clean-architecture/model"
	"clean-architecture/repo"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Signup(user model.User) error
	Login(email, password string) (uint, error)
	SetSID(val string) (*http.Cookie, error)
	CreateSession(session *model.Session) error
}

type userUsecase struct {
	ur repo.IUserRepo
}

func NewUserUsecase(ur repo.IUserRepo) IUserUsecase {
	return &userUsecase{ur}
}

func (uu *userUsecase) Signup(user model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	if err := uu.ur.CreateUser(&user); err != nil {
		return err
	}
	return nil
}

func (uu *userUsecase) Login(email, password string) (uint, error) {
	gotUser, err := uu.ur.GetUserByEmail(email)
	if err != nil {
		return 0, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(gotUser.Password), []byte(password)); err != nil {
		return 0, err
	}
	return gotUser.ID, nil
}

func (uu *userUsecase) SetSID(val string) (*http.Cookie, error) {
	return uu.ur.SetSID(val)
}

func (uu *userUsecase) CreateSession(session *model.Session) error {
	return uu.ur.CreateSession(session)
}
