package datastore

import(
	"encoding/json"
)

type User struct {
	Id int
	Name string
}

func CreateUser(body []byte) {
	// JSON Parse
	var user User
	json.Unmarshal(body, &user)

	// Create
	Db.Create(&user)
}