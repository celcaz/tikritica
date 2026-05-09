package entity

import "time"

type Series struct {
	ID            string    `json:"id"`
	Slug          string    `json:"slug"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CoverURL      string    `json:"coverUrl"`
	ReleaseDate   string    `json:"releaseDate"`
	Creator       string    `json:"creator"`
	Seasons       int       `json:"seasons"`
	Episodes      int       `json:"episodes"`
	Genres        []string  `json:"genres"`
	AverageRating float64   `json:"averageRating"`
	RatingsCount  int       `json:"ratingsCount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
