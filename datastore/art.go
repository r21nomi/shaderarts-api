package datastore

import(
	"encoding/json"
	"github.com/rs/xid"
	"time"
)

type Art struct {
	ID string `json:"id"`
	Title string `json:"title"`
	UserID string `json:"user_id"`
	User User `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
	Type int `json:"type"`
	Thumb string `json:"thumb"`
	Src string `json:"src"`
	Description string `json:"description"`
	Star int `json:"star"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Method name must be uper camel case to be able to be accessed from another class.
func CreateArt(userID string, body []byte) {
	var art Art
	guid := xid.New()
	art.ID = guid.String()
	art.UserID = userID
	json.Unmarshal(body, &art)

	// Create
	Db.Create(&art)
}

func GetArts() (arts []Art) {
    // Get all Arts
    Db.Find(&arts)
    for i, _ := range arts {
        Db.Model(arts[i]).Related(&arts[i].User)
    }
    return
}

func getArt(id string) (art Art) {
	Db.First(&art, "id = ?", id).Related(&art.User)
	return
  }