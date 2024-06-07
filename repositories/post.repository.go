package repositories

import (
	"database/sql"
	"ngobar/berkicau/models"
)

type PostRepository interface {
	GetDataByPostId(string) (models.Post, error)
}

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) GetDataByPostId(postId string) (models.Post, error) {
	var post models.Post
	err := r.db.QueryRow("SELECT * FROM postss WHERE id = ?", postId).Scan(&post.ID, &post.Content, &post.OriginalPostId)
	if err != nil {
		return models.Post{}, err
	}

	return post, nil
}