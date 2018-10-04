package tools

import (
	"net/http"
	"encoding/json"
)

func GlobalError(error interface{}) ([]byte, error) {
	var c = map[string]interface{}{
		"errors": map[string]interface{}{
			"global": error,
		},
	}
	return json.Marshal(c)
}

func e(w http.ResponseWriter, error string, status int) {
	w.WriteHeader(status)
	res, err := GlobalError(error)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func ErrorResponse(w http.ResponseWriter, error string) {
	e(w, error, http.StatusBadRequest)
}

func ErrorResponse401(w http.ResponseWriter, error string) {
	w.WriteHeader(http.StatusUnauthorized)
	res, err := GlobalError(error)
	if err != nil {
		panic(err)
	}
	w.Write(res)
}

func ErrorResponseMail(w http.ResponseWriter, error string) {
	w.WriteHeader(http.StatusBadRequest)
	res, err := GlobalError(map[string]interface{}{
		"email": "Email deja utilisé",
		"stack": error,
	})
	if err != nil {
		panic(err)
	}
	w.Write(res)
}
