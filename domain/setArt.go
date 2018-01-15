package domain

import (
	"github.com/r21nomi/arto-api/datastore"
)

type SetArt struct{}

func (s *SetArt) Execute(userID string, body []byte) {
	datastore.CreateArt(userID, body)
}
