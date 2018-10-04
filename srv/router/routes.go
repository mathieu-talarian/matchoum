package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Queries     []string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func Rooters(routes Routes, router *mux.Router) {
	for _, route := range routes {
		if len(route.Queries) != 0 {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Queries(route.Queries...).
				Name(route.Name).
				Handler(route.HandlerFunc)
		} else {
			router.
				Methods(route.Method).
				Path(route.Pattern).
				Name(route.Name).
				Handler(route.HandlerFunc)
		}
	}
}
