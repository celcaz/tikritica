package entity

import "time"

type Follow struct {
	ID          string    `json:"id"`
	FollowerID  string    `json:"followerId"`
	FollowingID string    `json:"followingId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Like struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userId"`
	ReviewID  string    `json:"reviewId"`
	CreatedAt time.Time `json:"createdAt"`
}
