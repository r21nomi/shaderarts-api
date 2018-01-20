package datastore

import (
	"fmt"
	"image"
	"log"
	"strings"
	"time"

	"github.com/rs/xid"

	"encoding/base64"
	"os"

	"image/png"
)

type Art struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Type        int       `json:"type"`
	Thumb       string    `json:"thumb"`
	Description string    `json:"description"`
	Star        int       `json:"star"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	UserID      string    `json:"userId"`
	User        User      `gorm:"ForeignKey:UserID;AssociationForeignKey:ID" json:"user"`
	Codes       []Code    `json:"codes"`
}

func CreateArt(art Art) {
	guid := xid.New()
	art.ID = guid.String()

	toPngImage(art.Thumb)

	for i, _ := range art.Codes {
		guid := xid.New()
		art.Codes[i].ID = guid.String()
		art.Codes[i].ArtID = art.ID
		Db.Create(&art.Codes[i])
	}

	// Create
	Db.Create(&art)
}

func GetArts() (arts []Art) {
	// Get all Arts
	Db.Find(&arts)
	for i, _ := range arts {
		Db.Model(arts[i]).Related(&arts[i].User)
		Db.Model(arts[i]).Related(&arts[i].Codes)
	}
	return
}

func getArt(id string) (art Art) {
	Db.First(&art, "id = ?", id).Related(&art.User)
	return
}

func toPngImage(base64Thumb string) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(base64Thumb))
	img, formatString, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := img.Bounds()
	fmt.Println(bounds, formatString)

	guid := xid.New()
	filePath := ""
	fileName := guid.String() + ".png"
	f, err := os.OpenFile(filePath+fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = png.Encode(f, img)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Png file", fileName, "created")
}
