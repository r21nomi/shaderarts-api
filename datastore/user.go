package datastore

import(
	"encoding/json"
	"time"
)

type User struct {
	ID string `json:"id"`
	Token string `gorm:"type:text" json:"token"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func CreateUser(id string, token string, name string) {
	var user User
	user.ID = id
	user.Token = token
	user.Name = name

	// Create or update
	Db.Where("id = ?", id).Assign(user).FirstOrCreate(&user)
}

func GetUser(id string) ([]byte, error) {
	user := User{}
	
	Db.Where("id = ?", id).First(&user)

	return json.Marshal(user)
}