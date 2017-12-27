package domain

import(
	"github.com/r21nomi/arto-api/datastore"
)

type SetUser struct{}

func (s *SetUser) Execute(id string, token string, name string) {
	datastore.CreateUser(id, token, name)
}