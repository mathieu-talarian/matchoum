package middlewares

import (
	"net/http"
	"strings"
	"matcha/tools"
	"github.com/gorilla/context"
	"matcha/models"
)



func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cl, err := tools.Verify(strings.Split(r.Header.Get("authorization"), " ")[1])
		if err != nil {
			tools.ErrorResponse401(w, "Invalid token")
			return
		}
		if err := models.NewUser().FindByEmail(cl["email"].(string)); err != nil {
			tools.ErrorResponse401(w, "No registered user")
			return
		}
		context.Set(r, "user", cl["email"].(string))
		next.ServeHTTP(w, r)
	})
}
