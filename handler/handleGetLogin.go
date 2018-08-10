package handler

import (
	"fmt"
	"net/http"

	"encoding/json"

	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"
	"github.com/r21nomi/shaderarts-api/domain"
	"github.com/r21nomi/shaderarts-api/entity"
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
	firebaseUser, err := client.GetUser(context.Background(), userID)
	if err != nil {
		http.Error(w, "error getting User: "+err.Error(), 500)
		return
	}

	setUser := domain.SetUser{}
	setUser.Execute(userID, token, firebaseUser.UserInfo.DisplayName, firebaseUser.PhotoURL)

	getUser := domain.GetUserByID{}
	user, err := getUser.Execute(userID)
	if err != nil {
		http.Error(w, "can not get user: "+err.Error(), 500)
		return
	}

	serializer := entity.UserSerializer{user}
	bytes, err := json.Marshal(serializer.Entity())
	if err != nil {
		http.Error(w, "can not parse user: "+err.Error(), 500)
		return
	}

	fmt.Fprint(w, string(bytes))
}
