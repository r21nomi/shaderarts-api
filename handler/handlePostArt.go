package handler

import(
	"github.com/r21nomi/arto-api/datastore"
	"github.com/r21nomi/arto-api/domain"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	firebase "firebase.google.com/go"
)

/**
 * Create Art.
 */
 func HandlePostArt(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	token := r.Header.Get("X-Token")
	log.Printf("token: %s\n", token)

	getUserID := domain.GetUserID{}
	userID, err := getUserID.Execute(app, token)

	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	// Create art
	datastore.CreateArt(userID, body)

	w.WriteHeader(200)
	return
}