package datastore

import(
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

var Db *gorm.DB

func init() {
	var DBMS = "postgres"
	var DB_USER = os.Getenv("DB_USER")
	var DB_NAME = os.Getenv("DB_NAME")
	var DB_PASSWORD = os.Getenv("DB_PASSWORD")

	var err error
	Db, err = gorm.Open(DBMS, "user=" + DB_USER + " dbname=" + DB_NAME + " password=" + DB_PASSWORD + " sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Art{})
}