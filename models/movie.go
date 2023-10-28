package models

import "time"

type Movie struct {
	ID                  int               `gorm:"primary_key" json:"id"`
	Title               string            `json:"title"`
	Year                int               `json:"year"`
	AgeRatingCategoryID int               `json:"age_rating_category_id"`
	CreatedAt           time.Time         `json:"created_at"`
	UpdatedAt           time.Time         `json:"updated_at"`
	AgeRatingCategory   AgeRatingCategory `json:"-"`
}
