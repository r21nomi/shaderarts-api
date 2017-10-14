package datastore

import(
	"encoding/json"
)

type Art struct {
	Id int `json:"id"`
	Title string `json:"title"`
	User_Id string `json:"user_id"`
	Type int `json:"type"`
	Thumb string `json:"thumb"`
	Src string `json:"src"`
	Description string `json:"description"`
	// Tag []string `json:"tag"`
	Star int `json:"star"`
}

// Method name must be uper camel case to be able to be accessed from another class.
func CreateArt(body []byte) {
	// JSON Parse
	var art Art
	json.Unmarshal(body, &art)

	// Create
	Db.Create(&art)
}

func GetArt() ([]byte, error) {
	arts := []Art{}

	// Get all Arts
	Db.Find(&arts)
	
	return json.Marshal(arts)
}