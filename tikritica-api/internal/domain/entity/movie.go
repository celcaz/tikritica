package entity

import "time"

type Movie struct {
	ID            string    `json:"id"`
	Slug          string    `json:"slug"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CoverURL      string    `json:"coverUrl"`
	ReleaseDate   string    `json:"releaseDate"`
	Director      string    `json:"director"`
	Runtime       int       `json:"runtime"`
	Genres        []string  `json:"genres"`
	AverageRating float64   `json:"averageRating"`
	RatingsCount  int       `json:"ratingsCount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
