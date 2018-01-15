package domain

import (
	"github.com/r21nomi/arto-api/datastore"
)

type GetUser struct{}

func (g *GetUser) Execute(id string) ([]byte, error) {
	return datastore.GetUser(id)
}
