package requests

type CommentRequest struct {
	PostId             string `json:"postId"`
	Content             string `json:"content"`
}
