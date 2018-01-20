package handler

import (
	"encoding/json"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"

	"github.com/r21nomi/arto-api/datastore"
	"github.com/r21nomi/arto-api/domain"
)

/**
 * Create Art.
 */
func HandlePostArt(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var art datastore.Art
	err := decoder.Decode(&art)

	if err != nil {
		log.Fatalf("error decoding: %v\n", err)
	}

	token := r.Header.Get("X-Token")

	getUserID := domain.GetUserID{}
	userID, err := getUserID.Execute(app, token)

	if err != nil {
		log.Fatalf("error getting user id: %v\n", err)
	}

	// Create art
	setArt := domain.SetArt{}
	setArt.Execute(userID, art)

	w.WriteHeader(200)
	return
}
