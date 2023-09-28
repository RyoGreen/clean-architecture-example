package usecase

import (
	"clean-architecture/model"
	"clean-architecture/repo"
	"net/http"
)

type ISessionUsecase interface {
	GetSessoionByID(sID string) (*model.User, error)
	GetUserSid(getCookie *http.Cookie) (string, error)
}

type SessionUsecase struct {
	sr repo.ISessionRepo
}

func NewSessionUsecase(sr repo.ISessionRepo) ISessionUsecase {
	return &SessionUsecase{sr}
}

func (su SessionUsecase) GetSessoionByID(sID string) (*model.User, error) {
	return su.sr.GetSessoionByID(sID)
}

func (su SessionUsecase) GetUserSid(getCookie *http.Cookie) (string, error) {
	return su.sr.GetUserSid(getCookie)
}
