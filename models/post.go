package models

import "time"

type Post struct {
	ID        		uint      `json:"id"`
	Content     	string    `json:"content"`
	OriginalPostId  string    `json:"originalPostId"`
	CreatedAt 		time.Time `json:"createdAt"`
	UpdatedAt 		time.Time `json:"updatedAt"`
}