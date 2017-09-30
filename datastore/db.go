package datastore

import(
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"os"
)

var Db *gorm.DB

func init() {
	var DBMS = "postgres"
	var RDS_USERNAME = os.Getenv("RDS_USERNAME")
	var RDS_DB_NAME = os.Getenv("RDS_DB_NAME")
	var RDS_PASSWORD = os.Getenv("RDS_PASSWORD")

	var err error
	Db, err = gorm.Open(DBMS, "user=" + RDS_USERNAME + " dbname=" + RDS_DB_NAME + " password=" + RDS_PASSWORD + " sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Art{})
}