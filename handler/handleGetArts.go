package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/arto-api/domain"
	"github.com/r21nomi/arto-api/entity"
)

/**
 * Get Arts.
 */
func HandleGetArts(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	token := r.Header.Get("X-Token")
	getUserID := domain.GetUserID{}
	userID, err := getUserID.Execute(app, token)

	if err != nil {
		http.Error(w, "user id couldn't find: " + err.Error(), 400)
		return
	}

	//　Query param
	queryValues := r.URL.Query()
	l := queryValues.Get("limit")
	o := queryValues.Get("offset")

	var limit, offset = 8, 0

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
	getArts := domain.GetArts{}
	arts := getArts.Execute(limit, offset)
	serializer := entity.ArtsSerializer{arts}
	bytes, err := json.Marshal(serializer.Entities(userID))

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprint(w, string(bytes))
}
