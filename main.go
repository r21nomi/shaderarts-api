package main

import(
	"github.com/r21nomi/arto-api/handler"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	"os"
	"log"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

var app *firebase.App

func handlePostArt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandlePostArt(app, w, r, ps)
}

func handleGetArt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandleGetArts(w, r, ps)
}

func handleGetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandleGetLogin(app, w, r, ps)
}

func handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    fmt.Fprint(w, "Welcome!\n")
}

func initializeAppWithServiceAccount() *firebase.App {
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)

	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

func main() {
	app = initializeAppWithServiceAccount()

	router := httprouter.New()
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	router.POST("/v1/art", handlePostArt)
	router.GET("/v1/art", handleGetArt)
	router.GET("/v1/login", handleGetLogin)
	router.GET("/", handleIndex)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"X-Token"},
	})
	handler := c.Handler(router)

	http.ListenAndServe(":" + port, handler)
}