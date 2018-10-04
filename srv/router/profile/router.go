package profile

import (
	"github.com/gorilla/mux"
	"matcha/router"
	"matcha/middlewares"
)

var routes = router.Routes{
	router.Route{
		Name:        "get Info",
		Method:      "GET",
		Pattern:     "",
		HandlerFunc: GetProfile,
	},
	router.Route{
		Name:    "post Info",
		Method:  "POST",
		Pattern: "",
		// Queries:     []string{"q", ".+"},
		HandlerFunc: PostProfile,
	},
}

func ProfileSubrouter(r *mux.Router) {
	s := r.PathPrefix("/profile").Subrouter()
	s.Use(middlewares.Authenticate)
	router.Rooters(routes, s)
}
