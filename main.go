package main

import(
	"github.com/r21nomi/arto-api/datastore"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
)

func handlePost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	// Create post
	datastore.Create(body)

	w.WriteHeader(200)
	return
}

func main() {
	router := httprouter.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

    router.POST("/post", handlePost)

	http.ListenAndServe(":" + port, router)
}