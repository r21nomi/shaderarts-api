package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/shaderarts-api/domain"
	"github.com/r21nomi/shaderarts-api/entity"
)

/**
 * Get User Arts.
 */
func HandleGetUserArts(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	userID := ps.ByName("id")

	//　Query param
	queryValues := r.URL.Query()
	l := queryValues.Get("limit")
	o := queryValues.Get("offset")

	var limit, offset = 8, 0
	var err error = nil

	if l != "" {
		limit, err = strconv.Atoi(l)
	}
	if o != "" {
		offset, err = strconv.Atoi(o)
	}

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Get Arts
	getUserArts := domain.GetUserArts{}
	arts := getUserArts.Execute(userID, limit, offset)
	serializer := entity.ArtsSerializer{arts}
	bytes, err := json.Marshal(serializer.Entities(userID))

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprint(w, string(bytes))
}
