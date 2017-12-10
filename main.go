package main

import(
	"github.com/r21nomi/arto-api/datastore"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"log"
)

/**
 * Create Art.
 */
func handlePostArt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	// Create art
	datastore.CreateArt(body)

	w.WriteHeader(200)
	return
}

/**
 * Get Arts.
 */
func handleGetArt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	//　Query param
	queryValues := r.URL.Query()
	log.Printf("limit: %s\n", queryValues.Get("limit"))

	// Get Arts
	bytes, err := datastore.GetArt()
	if err != nil {
		fmt.Fprint(w, "error")
        return
    }

	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprint(w, string(bytes))
}

func handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func main() {
	router := httprouter.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router.POST("/v1/art", handlePostArt)
	router.GET("/v1/art", handleGetArt)
	router.GET("/", handleIndex)

	http.ListenAndServe(":" + port, router)
}