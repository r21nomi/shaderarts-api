package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/arto-api/domain"
	"github.com/r21nomi/arto-api/entity"
)

/**
 * Get Art by ID.
 */
func HandleGetArtDetail(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	token := r.Header.Get("X-Token")

	getUserID := domain.GetUserID{}
	userID, err := getUserID.Execute(app, token)

	if err != nil {
		http.Error(w, "user id couldn't find: "+err.Error(), 400)
		return
	}

	id := ps.ByName("id")

	getArtByID := domain.GetArtByID{}
	art := getArtByID.Execute(id)
	serializer := entity.ArtSerializer{art}
	bytes, err := json.Marshal(serializer.Entity(userID))

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprint(w, string(bytes))
}
