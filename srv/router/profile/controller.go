package profile

import (
	"net/http"
	"encoding/json"
	"fmt"
	"matcha/models"
	"github.com/gorilla/context"
	"matcha/tools"
)

func GetProfile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("here")
	userEmail, ok := context.Get(r, "user").(string)
	if !ok {
		tools.ErrorResponse(w, fmt.Sprint("Unknown user"))
	}
	fmt.Println(userEmail)
	t := models.NewProfile()
	if err := t.FindByUser(userEmail); err != nil {
		fmt.Println(err)
		ret, _ := json.Marshal(map[string]interface{}{"profile": t})
		w.Write(ret)
		return
	}
	fmt.Printf("%+v\n", t)
	ret, err := json.Marshal(map[string]interface{}{"profile": t})
	if err != nil {
		tools.ErrorResponse(w, fmt.Sprint(err))
		return
	}
	w.Write(ret)
}

func PostProfile(w http.ResponseWriter, r *http.Request) {
	var b struct {
		models.Profile `json:"profile"`
	}

	body := json.NewDecoder(r.Body)
	body.Decode(&b)
	u := models.NewUser()
	userEmail, ok := context.Get(r, "user").(string)
	if !ok {
		tools.ErrorResponse(w, "Issue with context")
	}
	if err := u.FindByEmail(userEmail); err != nil {
		tools.ErrorResponse(w, fmt.Sprint(err))
		return
	}
	b.Profile.UserId(u.Id)
	fmt.Println(b.Profile)
	if err := b.Profile.Save(); err != nil {
		tools.ErrorResponse(w, fmt.Sprint(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}
