package usecase

import (
	"clean-architecture/model"
	"clean-architecture/repo"
	"os"

	"golang.org/x/crypto/bcrypt"
)

var secret = os.Getenv("secret")

type IUserUsecase interface {
	Signup(user model.User) error
	Login(user model.User) error
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

func (uu *userUsecase) Login(user model.User) error {
	gotUser, err := uu.ur.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(gotUser.Password), []byte(user.Password)); err != nil {
		return err
	}
	return nil
}
