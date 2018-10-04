package main

import (
	"net/http"

	"matcha/db"
	"matcha/router/auth"
	"matcha/router/tags"
	"matcha/router/users"

	"bytes"
	"fmt"
	"os"
	"time"

	"matcha/router/profile"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type MyServer struct {
	r *mux.Router
}

func (s *MyServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		rw.Header().Set("Access-Control-Allow-Origin", origin)
		rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		rw.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}
	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	db.InitDb()
	defer db.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(notFound)
	router = router.PathPrefix("/api/v1").Subrouter()
	auth.AuthSubrouter(router)
	users.UsersSubrouter(router)
	tags.TagsSubrouter(router)
	profile.ProfileSubrouter(router)
	//http.ListenAndServe(":8080", enforceJSONHandler(JSONheader(router)))
	http.ListenAndServe(":3000", handlers.LoggingHandler(os.Stdout, router))
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func JSONheader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Date", time.Now().String())
		next.ServeHTTP(w, r)
	})
}

//TODO session cookie handler with getting user into context

func enforceJSONHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for a request body
		if r.ContentLength == 0 {
			http.Error(w, http.StatusText(400), 400)
			return
		}
		// Check its MIME type
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		fmt.Println(http.DetectContentType(buf.Bytes()))
		if http.DetectContentType(buf.Bytes()) != "application/json" {
			http.Error(w, http.StatusText(415), 415)
			return
		}
		next.ServeHTTP(w, r)
	})
}
