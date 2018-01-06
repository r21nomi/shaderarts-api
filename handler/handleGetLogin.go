package handler

import(
	"github.com/r21nomi/arto-api/domain"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	firebase "firebase.google.com/go"
	"golang.org/x/net/context"
	"fmt"
)

func HandleGetLogin(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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