package main

import(
	"github.com/r21nomi/arto-api/datastore"
	"github.com/r21nomi/arto-api/domain"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"net/http"
	"os"
	"log"
	"golang.org/x/net/context"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
	"encoding/json"
)

var app *firebase.App

/**
 * Create Art.
 */
func handlePostArt(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	token := r.Header.Get("X-Token")
	log.Printf("token: %s\n", token)

	getUserID := domain.GetUserID{}
	userID, err := getUserID.Execute(app, token)

	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	// Create art
	datastore.CreateArt(userID, body)

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
	arts := datastore.GetArts()
	bytes, err := json.Marshal(arts)

	if err != nil {
		fmt.Fprint(w, "error")
        return
	}

	log.Printf("arts: %s\n", string(bytes))
	fmt.Fprint(w, string(bytes))
}

func handleGetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	token := r.Header.Get("X-Token")
	log.Printf("token: %s\n", token)

	getUserID := domain.GetUserID{}
	userID, err := getUserID.Execute(app, token)

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	user, err := client.GetUser(context.Background(), userID)
	if err != nil {
		log.Fatalf("error getting User: %v\n", err)
	}

	setUser := domain.SetUser{}
	setUser.Execute(userID, token, user.UserInfo.DisplayName)

	getUser := domain.GetUser{}
	bytes, err := getUser.Execute(userID)
	if err != nil {
		fmt.Fprint(w, "can not get user.")
        return
	}
	log.Printf("user: %s\n", string(bytes))
	
	fmt.Fprint(w, string(bytes))
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