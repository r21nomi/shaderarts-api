package datastore

import (
	"time"

	"github.com/rs/xid"
)

type Art struct {
	ID          string
	Title       string
	Type        int
	Thumb       string
	Description string
	Star        int
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      string
	User        User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	Codes       []Code
	Tags        []Tag `gorm:"many2many:art_tags;"`
}

func (art Art) IsStarredBy(userId string) bool {
	var star Star
	Db.Where(Star{
		StarID:     art.ID,
		StaredByID: userId,
	}).First(&star)
	return star.ID != ""
}

func CreateArt(art Art) {
	guid := xid.New()
	art.ID = guid.String()

	for i, _ := range art.Codes {
		guid := xid.New()
		art.Codes[i].ID = guid.String()
		art.Codes[i].ArtID = art.ID
		Db.Create(&art.Codes[i])
	}

	for i, _ := range art.Tags {
		var tag Tag

		if Db.First(&tag, "text = ?", art.Tags[i].Text).RecordNotFound() {
			// New tag.
			guid := xid.New()
			art.Tags[i].ID = guid.String()
			Db.Create(&art.Tags[i])
		} else {
			// Tag already exists.
			art.Tags[i].ID = tag.ID
		}
	}

	// Create
	Db.Create(&art)
}

func GetArtByID(id string) (art Art) {
	art.ID = id

	Db.Find(&art).Related(&art.User).Related(&art.Codes).Related(&art.Tags, "Tags")

	return
}

func GetArts(limit int, offset int) (arts []Art) {
	// Get all Arts
	Db.Order("created_at desc").Limit(limit).Offset(offset).Find(&arts)
	for i, _ := range arts {
		Db.Model(arts[i]).Related(&arts[i].User)
		Db.Model(arts[i]).Related(&arts[i].Codes)
		Db.Model(arts[i]).Related(&arts[i].Tags, "Tags")
	}
	return
}

func getArt(id string) (art Art) {
	Db.First(&art, "id = ?", id).Related(&art.User)
	return
}
