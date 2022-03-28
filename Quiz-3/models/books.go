package models

import "time"

type (
	// Books
	Books struct {
		ID           int       `json:"id"`
		Title        string    `json:"title"`
		Description  string    `json:"description"`
		Image_url    string    `json:"image_url"`
		Release_Year int       `json:"release_year"`
		Price        string    `json:"price"`
		Total_Page   int       `json:"total_page"`
		Thickness    string    `json:"thickness"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Category_ID  int
	}
)
