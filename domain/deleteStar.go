package domain

import (
	"github.com/r21nomi/shaderarts-api/datastore"
)

type DeleteStar struct{}

func (a *DeleteStar) Execute(userID string, artID string) error {
	user, err := datastore.GetUserByID(userID)

	if err != nil {
		return err
	}

	art := datastore.GetArtByID(artID)

	return user.DeleteStar(art)
}
