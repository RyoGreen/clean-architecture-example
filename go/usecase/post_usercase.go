package usecase

import (
	"clean-architecture/model"
	"clean-architecture/repo"
)

type IPostUsecase interface {
	ListsPosts() ([]*model.Post, error)
	Create(post *model.Post) error
}

type postUsecase struct {
	ir repo.IPostRepo
}

func NewPostUsecase(pr repo.IPostRepo) IPostUsecase {
	return &postUsecase{pr}
}

func (pu postUsecase) ListsPosts() ([]*model.Post, error) {
	posts, err := pu.ir.ListPosts()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (pu *postUsecase) Create(post *model.Post) error {
	if err := pu.ir.CreatePost(post); err != nil {
		return err
	}
	return nil
}
