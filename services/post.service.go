package services

import (
	"ngobar/berkicau/helper"
	"ngobar/berkicau/repositories"
	"ngobar/berkicau/requests"
)

type PostService interface {
	CreateNewPost(requests.PostRequest) (valid bool, message string)
	AddNewComment(requests.CommentRequest) (valid bool, message string)
	RepostExistingPost(requests.RepostRequest) (valid bool, message string)
}

type postService struct {
	postRepo repositories.PostRepository
}

func NewPostService() PostService {
	return &postService{postRepo: repositories.NewPostRepository(helper.DB)}
}

func (s *postService) CreateNewPost(request requests.PostRequest) (valid bool, message string) {
	return false, ""
}

func (s *postService) AddNewComment(request requests.CommentRequest) (valid bool, message string) {
	return false, ""
}

func (s *postService) RepostExistingPost(request requests.RepostRequest) (valid bool, message string) {
	return false, ""
}