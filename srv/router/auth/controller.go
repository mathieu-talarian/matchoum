package auth

import (
	"net/http"
	"encoding/json"
	"matcha/models"
	"matcha/tools"
	"matcha/mailing"
	"github.com/dgrijalva/jwt-go"
	"os"
	"fmt"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Index func handle /api/v1/auth
func Index(w http.ResponseWriter, r *http.Request) {
	var c struct {
		Credentials `json:"credentials"`
	}
	b := json.NewDecoder(r.Body)
	b.Decode(&c)
	u := models.NewUser()
	if err := u.FindByEmail(c.Email); err != nil {
		tools.ErrorResponse(w, "Pas d'utilisateur avec cet email")
		return
	}
	if err := u.IsValidPassword(c.Password); err != nil {
		tools.ErrorResponse(w, "Invalid Credentials")
		return
	}
	w.WriteHeader(http.StatusOK)
	res, err := u.ToAuthJson()
	if err != nil {
		panic(err)
	}
	ret, err := json.Marshal(map[string]interface{}{"user": res})
	if err != nil {
		tools.ErrorResponse(w, fmt.Sprint(err))
	}
	w.Write(ret)
}

// Confirmation func handle /api/v1/auth/confirmation
func Confirmation(w http.ResponseWriter, r *http.Request) {
	var t struct {
		Token string `json:"token"`
	}
	b := json.NewDecoder(r.Body)
	b.Decode(&t)
	u := models.NewUser()
	if err := u.FindByTokenAndUpdate(t.Token); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	res, err := u.ToAuthJson()
	if err != nil {
		panic(err)
	}
	ret, err := json.Marshal(map[string]interface{}{"user": res})
	if err != nil {
		tools.ErrorResponse(w, fmt.Sprint(err))
	}
	w.Write(ret)
}

// ResetPasswordRequest func handle /api/v1/auth/rest_password_request
func ResetPasswordRequest(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Email string `json:"email"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	u := models.NewUser()
	if err := u.FindByEmail(body.Email); err != nil {
		tools.ErrorResponse(w, "Il n'y a pas d'utilisateur avec cet email")
		return
	}
	go func(u *models.User) {
		if err := mailing.SendResetPasswordRequest(u); err != nil {
			panic(err)
		}
		return
	}(u)
	w.WriteHeader(200)
}

// ValidateToken func handle /api/v1/auth/validate_token
func ValidateToken(w http.ResponseWriter, r *http.Request) {
	// TODO compare with db
	var body struct {
		Token string `json:"token"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	_, err := jwtDecode(body.Token)
	if err != nil {
		tools.ErrorResponse401(w, fmt.Sprint(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(map[string]interface{}{})
	w.Write(b)
}

// ResetPassword func handle /api/v1/auth/reset_password
func ResetPassword(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Data struct {
			Password             string `json:"password"`
			PasswordConfirmation string `json:"passwordConfirmation"`
			Token                string `json:"token"`
		} `json:"data"`
	}
	json.NewDecoder(r.Body).Decode(&body)
	data := body.Data
	if data.Password == "" || data.PasswordConfirmation == "" || data.Token == "" {
		tools.ErrorResponse(w, "Issue with datas")
		return
	}
	if data.PasswordConfirmation != data.Password {
		tools.ErrorResponse(w, "Mot de passe et confirmation ne matchent pas")
		return
	}
	token, err := jwtDecode(data.Token)
	if err != nil {
		tools.ErrorResponse401(w, fmt.Sprint(err))
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok && !token.Valid {
		tools.ErrorResponse(w, "Issue with Token")
	}
	u := new(models.User)
	if err := u.FindById(int(claims["_id"].(float64))); err != nil {
		tools.ErrorResponse401(w, fmt.Sprint(err))
		return
	}
	if err := u.UpdatePassword(data.Password); err != nil {
		tools.ErrorResponse(w, fmt.Sprint(err))
	}
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(map[string]interface{}{})
	w.Write(b)
}

func jwtDecode(token string) (*jwt.Token, error) {
	return jwt.Parse(token,
		func(token *jwt.Token) (interface{}, error) {
			fmt.Println(token)
			return []byte(os.Getenv("SECRET")), nil
		})
}
