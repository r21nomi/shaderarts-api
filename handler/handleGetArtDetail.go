package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/arto-api/domain"
	"github.com/r21nomi/arto-api/entity"
)

/**
 * Get Art by ID.
 */
func HandleGetArtDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	getArtByID := domain.GetArtByID{}
	art := getArtByID.Execute(id)
	serializer := entity.ArtSerializer{art}
	bytes, err := json.Marshal(serializer.Entity())

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprint(w, string(bytes))
}
