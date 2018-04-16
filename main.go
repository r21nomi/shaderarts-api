package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/arto-api/handler"
	"github.com/rs/cors"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
)

var app *firebase.App

func handleIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func handleGetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandleGetLogin(app, w, r, ps)
}

func handleGetUserDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandleGetUserDetail(app, w, r, ps)
}

func handleGetMyArts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandleGetMyArts(app, w, r, ps)
}

func handleGetExplore(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandleGetExplore(app, w, r, ps)
}

func handleGetArts(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandleGetArts(app, w, r, ps)
}

func handlePostArt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandlePostArt(app, w, r, ps)
}

func handleGetArtDetail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandleGetArtDetail(app, w, r, ps)
}

func handlePostStar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandlePostStar(app, w, r, ps)
}

func handleDeleteStar(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	handler.HandleDeleteStar(app, w, r, ps)
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

	router.GET("/", handleIndex)
	router.GET("/v1/login", handleGetLogin)
	router.GET("/v1/user/:id", handleGetUserDetail)
	router.GET("/v1/me/arts", handleGetMyArts)

	router.GET("/v1/explore", handleGetExplore)

	router.GET("/v1/art", handleGetArts)
	router.GET("/v1/art/:id", handleGetArtDetail)
	router.POST("/v1/art", handlePostArt)

	router.POST("/v1/star/:artId", handlePostStar)
	router.DELETE("/v1/star/:artId", handleDeleteStar)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"X-Token"},
	})
	handler := c.Handler(router)

	http.ListenAndServe(":"+port, handler)
}
