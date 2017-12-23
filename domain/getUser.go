package domain

import(
	"github.com/r21nomi/arto-api/datastore"
)

func GetUser(id string) ([]byte, error) {
	return datastore.GetUser(id)
}