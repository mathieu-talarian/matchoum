package users

import (
	"net/http"
	"matcha/models"
	"encoding/json"
	"matcha/tools"
	"fmt"
	"matcha/mailing"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var body struct {
		*models.User `json:"user"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	u := body.User
	u.SetPassword(u.Password)
	u.SetConfirmationToken()
	if err := u.Save(); err != nil {
		tools.ErrorResponseMail(w, fmt.Sprint(err))
		return
	}
	go func(u *models.User) {
		if err := mailing.SendConfirmationEmail(u); err != nil {
			panic(err)
		}
	}(u)
	w.WriteHeader(http.StatusOK)
	res, err := u.ToAuthJson()
	if err != nil {
		panic(err)
	}
	ret, err := json.Marshal(map[string]interface{}{"user": res})
	if err != nil {
		tools.ErrorResponse(w, fmt.Sprint(err))
	}
	w.Write(ret)
}
