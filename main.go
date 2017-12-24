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
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var app *firebase.App

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

	log.Printf("arts: %s\n", string(bytes))
	fmt.Fprint(w, string(bytes))
}

func handleGetLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	token := r.Header.Get("X-Token")
	log.Printf("token: %s\n", token)

	authToken := verifyIDToken(app, token)

	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	user, err := client.GetUser(context.Background(), authToken.UID)
	if err != nil {
		log.Fatalf("error getting User: %v\n", err)
	}

	domain.SetUser(authToken.UID, token, user.UserInfo.DisplayName)

	bytes, err := domain.GetUser(authToken.UID)
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

func verifyIDToken(app *firebase.App, idToken string) *auth.Token {
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	token, err := client.VerifyIDToken(idToken)
	if err != nil {
		log.Fatalf("error verifying ID token: %v\n", err)
	}

	return token
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