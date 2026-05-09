package entity

import "time"

type List struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsPublic    bool      `json:"isPublic"`
	ItemsCount  int       `json:"itemsCount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
