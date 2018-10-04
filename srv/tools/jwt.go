package tools

import (
	"github.com/dgrijalva/jwt-go"
	"os"
)

func Verify(token string) (ret jwt.MapClaims, err error) {
	var tk *jwt.Token
	tk, err = jwt.Parse(token,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET")), nil
		})
	if err != nil || !tk.Valid {
		return
	}
	return tk.Claims.(jwt.MapClaims), nil
}
