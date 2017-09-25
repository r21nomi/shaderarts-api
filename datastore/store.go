package datastore

import(
	"github.com/jinzhu/gorm"
	"encoding/json"
	_ "github.com/lib/pq"
	"os"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var Db *gorm.DB

// Method name must be uper camel case to be able to be accessed from another class.
func Create(body []byte) {
	// JSON Parse
	var post Post
	json.Unmarshal(body, &post)

	// Create
	Db.Create(&post)
}

func init() {
	var dbUser = os.Getenv("DB_USER")
	var dbName = os.Getenv("DB_NAME")
	var dbPassword = os.Getenv("DB_PASSWORD")
	
	var err error
	Db, err = gorm.Open("postgres", "user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{})
}