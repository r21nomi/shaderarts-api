package domain

import (
	"github.com/r21nomi/arto-api/datastore"
)

type DeleteStar struct{}

func (a *DeleteStar) Execute(userID string, artID string) error {
	user, err := datastore.GetUser(userID)

	if err != nil {
		return err
	}

	art := datastore.GetArtByID(artID)

	return user.DeleteStar(art)
}
