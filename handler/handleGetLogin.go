package handler

import (
	"fmt"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/arto-api/domain"
	"golang.org/x/net/context"
)

func HandleGetLogin(app *firebase.App, w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)

	token := r.Header.Get("X-Token")

	getUserID := domain.GetUserID{}
	userID, err := getUserID.Execute(app, token)

	client, err := app.Auth(context.Background())
	if err != nil {
		http.Error(w, "error getting Auth client: "+err.Error(), 500)
		return
	}
	user, err := client.GetUser(context.Background(), userID)
	if err != nil {
		http.Error(w, "error getting User: "+err.Error(), 500)
		return
	}

	setUser := domain.SetUser{}
	setUser.Execute(userID, token, user.UserInfo.DisplayName, user.PhotoURL)

	getUser := domain.GetUser{}
	bytes, err := getUser.Execute(userID)
	if err != nil {
		http.Error(w, "can not get user: "+err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(bytes))
}
