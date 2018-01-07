package datastore

import(
	"encoding/json"
	"github.com/rs/xid"
	"time"
)

type Art struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Type int `json:"type"`
	Thumb string `json:"thumb"`
	Description string `json:"description"`
	Star int `json:"star"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID string `json:"userId"`
	User User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID" json:"user"`
	Programs []Program `json:"programs"`
}

func CreateArt(userID string, body []byte) {
	var art Art
	guid := xid.New()
	art.ID = guid.String()
	art.UserID = userID
	json.Unmarshal(body, &art)

	for i, _ := range art.Programs {
		guid := xid.New()
		art.Programs[i].ID = guid.String()
		art.Programs[i].ArtID = art.ID
		Db.Create(&art.Programs[i])
	}

	// Create
	Db.Create(&art)
}

func GetArts() (arts []Art) {
    // Get all Arts
    Db.Find(&arts)
    for i, _ := range arts {
        Db.Model(arts[i]).Related(&arts[i].User)
        Db.Model(arts[i]).Related(&arts[i].Programs)
    }
    return
}

func getArt(id string) (art Art) {
	Db.First(&art, "id = ?", id).Related(&art.User)
	return
}