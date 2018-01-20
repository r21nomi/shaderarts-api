package domain

import (
	"github.com/r21nomi/arto-api/datastore"
)

type SetArt struct{}

func (s *SetArt) Execute(art datastore.Art, userID string, artThumbPath string) {
	art.UserID = userID
	art.Thumb = artThumbPath
	datastore.CreateArt(art)
}
