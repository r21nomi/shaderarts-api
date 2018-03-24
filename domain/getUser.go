package domain

import (
	"encoding/json"

	"github.com/r21nomi/arto-api/datastore"
)

type GetUser struct{}

func (g *GetUser) Execute(id string) ([]byte, error) {
	user, err := datastore.GetUser(id)

	if err != nil {
		return nil, err
	}

	return json.Marshal(user)
}
