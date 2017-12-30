package datastore

import(
	"encoding/json"
)

type User struct {
	Id string `json:"id"`
	Token string `json:"token"`
	Name string `json:"name"`
}

func CreateUser(id string, token string, name string) {
	// JSON Parse
	var user User
	user.Id = id
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