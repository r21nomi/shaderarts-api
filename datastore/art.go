package datastore

import (
	"encoding/json"
	"time"

	"github.com/rs/xid"
)

type Art struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Type        int       `json:"type"`
	Thumb       string    `json:"thumb"`
	Description string    `json:"description"`
	Star        int       `json:"star"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserID      string    `json:"userId"`
	User        User      `gorm:"ForeignKey:UserID;AssociationForeignKey:ID" json:"user"`
	Codes       []Code    `json:"codes"`
}

func CreateArt(userID string, body []byte) {
	var art Art
	guid := xid.New()
	art.ID = guid.String()
	art.UserID = userID
	json.Unmarshal(body, &art)

	for i, _ := range art.Codes {
		guid := xid.New()
		art.Codes[i].ID = guid.String()
		art.Codes[i].ArtID = art.ID
		Db.Create(&art.Codes[i])
	}

	// Create
	Db.Create(&art)
}

func GetArts() (arts []Art) {
	// Get all Arts
	Db.Find(&arts)
	for i, _ := range arts {
		Db.Model(arts[i]).Related(&arts[i].User)
		Db.Model(arts[i]).Related(&arts[i].Codes)
	}
	return
}

func getArt(id string) (art Art) {
	Db.First(&art, "id = ?", id).Related(&art.User)
	return
}
