package domain

import (
	"github.com/r21nomi/arto-api/datastore"
)

type AddStar struct{}

func (a *AddStar) Execute(userID string, artID string) error {
	user, err := datastore.GetUser(userID)

	if err != nil {
		return err
	}

	art := datastore.GetArtByID(artID)

	return user.AddStar(art)
}
