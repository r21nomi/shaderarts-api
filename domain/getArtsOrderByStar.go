package domain

import (
	"github.com/r21nomi/arto-api/datastore"
)

type GetArtsOrderByStar struct{}

func (g *GetArtsOrderByStar) Execute(limit int, offset int) []datastore.Art {
	return datastore.GetArtsOrderByStar(limit, offset)
}
