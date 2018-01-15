package domain

import (
	"github.com/r21nomi/arto-api/datastore"
)

type GetArts struct{}

func (g *GetArts) Execute() []datastore.Art {
	return datastore.GetArts()
}
