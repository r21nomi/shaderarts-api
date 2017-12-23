package datastore

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"log"
)

var Db *gorm.DB

func init() {
	var DBMS = "mysql"
	var USER = os.Getenv("RDS_USER")
	var PASS = os.Getenv("RDS_PASS")
	var PROTOCOL = "tcp(" + os.Getenv("RDS_PROTOCOL") + ")"
	var DBNAME = os.Getenv("RDS_DBNAME")
	var CONNECT = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME

	log.Printf("RDS_USERNAME: %s\n\n", USER)
	log.Printf("RDS_DB_NAME: %s\n\n", DBNAME)

	var err error
	Db, err = gorm.Open(DBMS, CONNECT)
	if err != nil {
			panic(err)
	}
	Db.AutoMigrate(&User{}, &Art{})
}