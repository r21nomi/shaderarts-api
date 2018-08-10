package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/shaderarts-api/domain"
	"github.com/r21nomi/shaderarts-api/entity"
)

/**
 * Get User by ID.
 */
func HandleGetUserDetail(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	id := ps.ByName("id")

	getUserByID := domain.GetUserByID{}
	user, err := getUserByID.Execute(id)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	serializer := entity.UserSerializer{user}
	bytes, err := json.Marshal(serializer.Entity())

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprint(w, string(bytes))
}
