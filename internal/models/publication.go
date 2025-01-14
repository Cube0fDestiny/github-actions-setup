package models

import (
	"time"

	"github.com/google/uuid"
)

type Publication struct {
	ID              uuid.UUID `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Journal         string    `db:"journal" json:"journal"`
	PublicationDate time.Time `db:"publication_date" json:"publication_date"`
	CitationsCount  int       `db:"citations_count" json:"citations_count"`
	ImpactFactor    float64   `db:"impact_factor" json:"impact_factor"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
type CreatePublication struct {
	Title           string    `db:"title" json:"title"`
	Journal         string    `db:"journal" json:"journal"`
	PublicationDate time.Time `db:"publication_date" json:"publication_date"`
	CitationsCount  int       `db:"citations_count" json:"citations_count"`
	ImpactFactor    float64   `db:"impact_factor" json:"impact_factor"`
}
type UpdatePublication struct {
	ID              uuid.UUID `db:"id" json:"id"`
	Title           string    `db:"title" json:"title"`
	Journal         string    `db:"journal" json:"journal"`
	PublicationDate time.Time `db:"publication_date" json:"publication_date"`
	CitationsCount  int       `db:"citations_count" json:"citations_count"`
	ImpactFactor    float64   `db:"impact_factor" json:"impact_factor"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
type DeletePublication struct {
	ID uuid.UUID `db:"id" json:"id"`
}
