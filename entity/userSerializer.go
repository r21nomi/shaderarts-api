package entity

import "time"
import "github.com/r21nomi/shaderarts-api/datastore"

type UserSerializer struct {
	datastore.User
}

type UserEntity struct {
	ID        string    `json:"id"`
	Token     string    `json:"token"`
	Name      string    `json:"name"`
	Thumb     string    `json:"thumb"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (self *UserSerializer) Entity() UserEntity {
	return UserEntity{
		ID:        self.ID,
		Token:     self.Token,
		Name:      self.Name,
		Thumb:     self.Thumb,
		CreatedAt: self.CreatedAt,
		UpdatedAt: self.UpdatedAt,
	}
}
