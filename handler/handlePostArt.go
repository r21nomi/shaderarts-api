package handler

import (
	"encoding/json"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"

	"github.com/r21nomi/shaderarts-api/datastore"
	"github.com/r21nomi/shaderarts-api/domain"
)

/**
 * Create Art.
 */
func HandlePostArt(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var art datastore.Art
	err := decoder.Decode(&art)

	if err != nil {
		http.Error(w, "error decoding: "+err.Error(), 400)
		return
	}

	token := r.Header.Get("X-Token")

	getUserID := domain.GetUserID{}
	userID, err := getUserID.Execute(app, token)

	if err != nil {
		http.Error(w, "error getting user id: "+err.Error(), 500)
		return
	}

	uploadImage := domain.UploadImage{}
	artThumbPath, err := uploadImage.Execute(art.Thumb)

	if err != nil {
		http.Error(w, "error upload image: "+err.Error(), 500)
		return
	}

	log.Println("artThumbPath: " + artThumbPath)

	// Create art
	setArt := domain.SetArt{}
	setArt.Execute(art, userID, artThumbPath)

	w.WriteHeader(200)
	return
}
