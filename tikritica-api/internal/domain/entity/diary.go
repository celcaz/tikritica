package entity

import "time"

type DiaryEntry struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	MediaID   string    `json:"mediaId"`
	MediaType string    `json:"mediaType"`
	WatchedAt time.Time `json:"watchedAt"`
	Rating    *int      `json:"rating"`
	ReviewID  *string   `json:"reviewId"`
	CreatedAt time.Time `json:"createdAt"`
}
