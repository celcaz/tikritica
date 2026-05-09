package entity

import "time"

type Game struct {
	ID            string    `json:"id"`
	Slug          string    `json:"slug"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CoverURL      string    `json:"coverUrl"`
	ReleaseDate   string    `json:"releaseDate"`
	Developer     string    `json:"developer"`
	Publisher     string    `json:"publisher"`
	Platforms     []string  `json:"platforms"`
	Genres        []string  `json:"genres"`
	AverageRating float64   `json:"averageRating"`
	RatingsCount  int       `json:"ratingsCount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
