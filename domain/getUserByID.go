package domain

import (
	"github.com/r21nomi/shaderarts-api/datastore"
)

type GetUserByID struct{}

func (g *GetUserByID) Execute(id string) (datastore.User, error) {
	return datastore.GetUserByID(id)
}
