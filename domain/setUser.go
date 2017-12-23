package domain

import(
	"github.com/r21nomi/arto-api/datastore"
)

func SetUser(id string, token string, name string) {
	datastore.CreateUser(id, token, name)
}