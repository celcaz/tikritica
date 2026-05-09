package entity

import "time"

type Review struct {
	ID               string    `json:"id"`
	UserID           string    `json:"userId"`
	MediaID          string    `json:"mediaId"`
	MediaType        string    `json:"mediaType"`
	Rating           int       `json:"rating"`
	Content          string    `json:"content"`
	ContainsSpoilers bool      `json:"containsSpoilers"`
	LikesCount       int       `json:"likesCount"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
