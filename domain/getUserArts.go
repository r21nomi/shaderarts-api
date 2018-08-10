package domain

import (
	"github.com/r21nomi/shaderarts-api/datastore"
)

type GetUserArts struct{}

func (g *GetUserArts) Execute(userID string, limit int, offset int) []datastore.Art {
	return datastore.GetArtsByUserId(userID, limit, offset)
}
