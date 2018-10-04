package tags

import (
	"encoding/json"
	"matcha/models"
	"net/http"
	"strings"
	"fmt"
	"matcha/tools"
	)

func All(w http.ResponseWriter, r *http.Request) {
	res, err := models.NewTags().FindAll()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	ret, err := json.Marshal(map[string]interface{}{"tags": res})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)
	w.Write(ret)
}

//SearchTags func
func SearchTags(w http.ResponseWriter, r *http.Request) {
	query := r.FormValue("q")
	query = strings.TrimSpace(query)
	res, err := models.NewTags().Match(strings.ToLower(query))
	if err != nil {
		panic(err)
	}
	ret, err := json.Marshal(map[string]interface{}{"tags": res})
	if err != nil {
		panic(err)
	}
	w.WriteHeader(200)

	w.Write(ret)
}

func AddTag(w http.ResponseWriter, r *http.Request) {
	var tag *models.Tags
	b := json.NewDecoder(r.Body)
	b.Decode(&tag)
	if tag.Tag == "" {
		tools.ErrorResponse(w, "Tag vide")
		return
	}
	tag.Tag = strings.ToLower(tag.Tag)
	fmt.Println(tag)
	if err := tag.Save(); err != nil {
		tools.ErrorResponse(w, "Ce tag existe deja")
		return
	}
	w.WriteHeader(http.StatusOK)
}
