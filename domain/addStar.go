package domain

import (
	"github.com/r21nomi/shaderarts-api/datastore"
)

type AddStar struct{}

func (a *AddStar) Execute(userID string, artID string) error {
	user, err := datastore.GetUserByID(userID)

	if err != nil {
		return err
	}

	art := datastore.GetArtByID(artID)

	return user.AddStar(art)
}
