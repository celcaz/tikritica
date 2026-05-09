package entity

import "time"

type Book struct {
	ID            string    `json:"id"`
	Slug          string    `json:"slug"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CoverURL      string    `json:"coverUrl"`
	ReleaseDate   string    `json:"releaseDate"`
	Author        string    `json:"author"`
	Pages         int       `json:"pages"`
	ISBN          string    `json:"isbn"`
	Genres        []string  `json:"genres"`
	AverageRating float64   `json:"averageRating"`
	RatingsCount  int       `json:"ratingsCount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
