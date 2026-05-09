package entity

import "time"

type User struct {
	ID             string    `json:"id"`
	Username       string    `json:"username"`
	DisplayName    string    `json:"displayName"`
	Email          string    `json:"-"`
	PasswordHash   string    `json:"-"`
	AvatarURL      *string   `json:"avatarUrl"`
	Bio            string    `json:"bio"`
	FollowersCount int       `json:"followersCount"`
	FollowingCount int       `json:"followingCount"`
	ReviewsCount   int       `json:"reviewsCount"`
	ListsCount     int       `json:"listsCount"`
	CreatedAt      time.Time `json:"createdAt"`
}
