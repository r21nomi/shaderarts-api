package domain

import (
	"github.com/r21nomi/shaderarts-api/datastore"
)

type GetArtByID struct{}

func (g *GetArtByID) Execute(id string) datastore.Art {
	return datastore.GetArtByID(id)
}
