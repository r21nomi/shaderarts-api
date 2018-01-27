package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/arto-api/domain"
)

/**
 * Get Arts.
 */
func HandleGetArts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	//　Query param
	queryValues := r.URL.Query()
	l := queryValues.Get("limit")
	o := queryValues.Get("offset")

	var limit, offset = 8, 0
	var err error

	if l != "" {
		limit, err = strconv.Atoi(l)
	}
	if o != "" {
		offset, err = strconv.Atoi(o)
	}

	if err != nil {
		log.Println("strconv error: %s\n", err.Error())
	}

	// Get Arts
	getArts := domain.GetArts{}
	arts := getArts.Execute(limit, offset)
	bytes, err := json.Marshal(arts)

	if err != nil {
		fmt.Fprint(w, "error")
		return
	}

	fmt.Fprint(w, string(bytes))
}
