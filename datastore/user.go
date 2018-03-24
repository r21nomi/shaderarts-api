package datastore

import (
	"time"

	"github.com/rs/xid"
)

type User struct {
	ID        string    `json:"id"`
	Token     string    `gorm:"type:text" json:"token"`
	Name      string    `json:"name"`
	Thumb     string    `json:"thumb"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func CreateUser(id string, token string, name string, thumb string) {
	var user User
	user.ID = id
	user.Token = token
	user.Name = name
	user.Thumb = thumb

	// Create or update
	Db.Where("id = ?", id).Assign(user).FirstOrCreate(&user)
}

func (user User) AddStar(art Art) error {
	var star Star
	star.ID = xid.New().String()

	err := Db.FirstOrCreate(&star, &Star{
		StarID:     art.ID,
		StaredByID: user.ID,
	}).Error

	return err
}

func (user User) RemoveStar(art Art) error {
	err := Db.Where(Star{
		StarID:     art.ID,
		StaredByID: user.ID,
	}).Delete(Star{}).Error

	return err
}

func GetUser(id string) (user User, err error) {
	err = Db.Where("id = ?", id).First(&user).Error

	return user, err
}
