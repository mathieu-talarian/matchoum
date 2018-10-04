package tags

import (
	"github.com/gorilla/mux"

	"matcha/router"
	"matcha/middlewares"
)

var routes = router.Routes{
	router.Route{
		Name:        "Search All",
		Method:      "GET",
		Pattern:     "/all",
		HandlerFunc: All,
	},
	router.Route{
		Name:    "Search",
		Method:  "GET",
		Pattern: "/search",
		// Queries:     []string{"q", ".+"},
		HandlerFunc: SearchTags,
	},
	router.Route{
		Name:        "AddTag",
		Method:      "POST",
		Pattern:     "",
		Queries:     []string{},
		HandlerFunc: AddTag,
	},
}

func TagsSubrouter(r *mux.Router) {
	s := r.PathPrefix("/tags").Subrouter()
	s.Use(middlewares.Authenticate)
	router.Rooters(routes, s)
}
