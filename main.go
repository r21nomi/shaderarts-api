package main

import(
	"github.com/r21nomi/arto-api/datastore"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func handleArt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	// Create art
	datastore.CreateArt(body)

	w.WriteHeader(200)
	return
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router.POST("/v1/art", handleArt)
	
	router.GET("/", Index)

	http.ListenAndServe(":" + port, router)
}