package datastore

import (
	"time"
)

type Star struct {
	ID         string
	Star       Art
	StarID     string
	StaredBy   User
	StaredByID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
