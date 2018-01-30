package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/arto-api/domain"
)

/**
 * Get Art by ID.
 */
func HandleGetArtDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	getArtByID := domain.GetArtByID{}
	art := getArtByID.Execute(id)
	bytes, err := json.Marshal(art)

	if err != nil {
		fmt.Fprint(w, "error")
		return
	}

	fmt.Fprint(w, string(bytes))
}
