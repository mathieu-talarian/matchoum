package users

import (
	"matcha/router"

	"github.com/gorilla/mux"
)

var routes = router.Routes{
	router.Route{
		Name:        "Index",
		Method:      "POST",
		Pattern:     "",
		HandlerFunc: Index,
	},
}

func UsersSubrouter(r *mux.Router) {
	s := r.PathPrefix("/users").Subrouter()
	router.Rooters(routes, s)
}
