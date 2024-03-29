package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/shaderarts-api/domain"
	"github.com/r21nomi/shaderarts-api/entity"
)

/**
 * Get Art Code by ID.
 */
func HandleGetArtCodeDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	id := ps.ByName("id")

	getArtByID := domain.GetArtByID{}
	art := getArtByID.Execute(id)
	serializer := entity.ArtCodeSerializer{art}
	bytes, err := json.Marshal(serializer.Entity())

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprint(w, string(bytes))
}