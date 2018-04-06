package datastore

import (
	"time"

	"github.com/rs/xid"
)

type User struct {
	ID        string
	Token     string `gorm:"type:text"`
	Name      string
	Thumb     string
	CreatedAt time.Time
	UpdatedAt time.Time
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

	// Create or update
	err := Db.Attrs(Star{ // Set ID only if the record is inserted.
		ID: xid.New().String(),
	}).Assign(Star{ // Set StarID regardless of the record is inserted or updated (to update an updated_at column).
		StarID: art.ID,
	}).FirstOrCreate(&star, &Star{
		StarID:     art.ID,
		StaredByID: user.ID,
	}).Error

	return err
}

func (user User) DeleteStar(art Art) error {
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
