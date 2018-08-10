package domain

import (
	"github.com/r21nomi/shaderarts-api/datastore"
)

type SetUser struct{}

func (s *SetUser) Execute(id string, token string, name string, thumb string) {
	datastore.CreateUser(id, token, name, thumb)
}
