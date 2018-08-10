package handler

import (
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"

	"github.com/r21nomi/shaderarts-api/domain"
)

/**
 * Delete Star.
 */
func HandleDeleteStar(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	token := r.Header.Get("X-Token")
	getUserID := domain.GetUserID{}
	userID, err := getUserID.Execute(app, token)

	if err != nil {
		http.Error(w, "user id couldn't find: "+err.Error(), 400)
		return
	}

	artID := ps.ByName("artId")
	deleteStar := domain.DeleteStar{}
	err = deleteStar.Execute(userID, artID)

	if err != nil {
		http.Error(w, "Couldn't add star: "+err.Error(), 500)
		return
	}

	w.WriteHeader(200)
	return
}
