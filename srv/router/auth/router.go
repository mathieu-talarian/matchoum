package auth

import (
	"matcha/router"

	"github.com/gorilla/mux"
)

var routes = router.Routes{
	router.Route{
		Name:        "Index",
		Method:      "POST",
		Pattern:     "",
		Queries:     []string{},
		HandlerFunc: Index,
	},
	router.Route{
		Name:        "Confirmation",
		Method:      "POST",
		Pattern:     "/confirmation",
		Queries:     []string{},
		HandlerFunc: Confirmation,
	},
	router.Route{
		Name:        "Reset password request",
		Method:      "POST",
		Pattern:     "/reset_password_request",
		Queries:     []string{},
		HandlerFunc: ResetPasswordRequest,
	},
	router.Route{
		Name:        "Validate Token",
		Method:      "POST",
		Pattern:     "/validate_token",
		Queries:     []string{},
		HandlerFunc: ValidateToken,
	},
	router.Route{
		Name:        "Reset Password",
		Method:      "POST",
		Pattern:     "/reset_password",
		Queries:     []string{},
		HandlerFunc: ResetPassword,
	},
}

func AuthSubrouter(r *mux.Router) {
	s := r.PathPrefix("/auth").Subrouter()
	router.Rooters(routes, s)
}
