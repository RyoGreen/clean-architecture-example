package repo

import (
	"clean-architecture/model"
	"database/sql"
)

const (
	getPostsQuery   = "SELECT * FROM post;"
	createPostQuery = "INSERT INTO post (content, created_at, updated_at, user_id) VALUES($1, $2, $3, $4);"
)

type IPostRepo interface {
	ListPosts() ([]*model.Post, error)
	CreatePost(post *model.Post) error
}
type postRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) IPostRepo {
	return &postRepo{db}
}

func (p postRepo) ListPosts() ([]*model.Post, error) {
	var posts = make([]*model.Post, 0)
	rows, err := p.db.Query(getPostsQuery)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var post model.Post
		if err := rows.Scan(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}
	return posts, nil
}

func (p postRepo) CreatePost(post *model.Post) error {
	_, err := p.db.Exec(createPostQuery, post.Content, post.CreatedAt, post.UpdatedAt, post.UserID)
	return err
}
