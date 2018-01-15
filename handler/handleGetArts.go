package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
	log.Printf("limit: %s\n", queryValues.Get("limit"))

	// Get Arts
	getArts := domain.GetArts{}
	arts := getArts.Execute()
	bytes, err := json.Marshal(arts)

	if err != nil {
		fmt.Fprint(w, "error")
		return
	}

	log.Printf("arts: %s\n", string(bytes))
	fmt.Fprint(w, string(bytes))
}
