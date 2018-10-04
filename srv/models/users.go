package models

import (
	"time"
	"matcha/db"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"fmt"
)

type User struct {
	Id                int       `json:"id"`
	Firstname         string    `json:"firstname"`
	Lastname          string    `json:"lastname"`
	Pseudo            string    `json:"pseudo"`
	Email             string    `json:"email"`
	Password          string    `json:"password"`
	Profile           bool      `json:"profile"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Confirmed         bool      `json:"confirmed"`
	ConfirmationToken string    `json:"confirmation_token"`
}

func (u *User) All() ([]interface{}) {
	return []interface{}{&u.Id,
		&u.Firstname,
		&u.Lastname,
		&u.Pseudo,
		&u.Email,
		&u.Password,
		&u.Profile,
		&u.CreatedAt,
		&u.UpdatedAt,
		&u.Confirmed,
		&u.ConfirmationToken}
}

type TokenClaims struct {
	Id int `json:"_id"`
	jwt.StandardClaims
}

type Auth interface {
	IsValidPassword(password string) error
	ToAuthJson()
}

func NewUser() (u *User) {
	return new(User)
}

func (u *User) FindById(id int) (err error) {
	row := db.QueryRow("SELECT * from users where id = $1", id)
	return row.Scan(u.All()...)
}

func (u *User) FindByEmail(email string) (error) {
	row := db.QueryRow("SELECT * FROM users where email = $1", email)
	return row.Scan(u.All()...)
}

func (u *User) IsValidPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

func (u *User) ToAuthJson() (map[string]interface{}, error) {
	return map[string]interface{}{
		"email":     u.Email,
		"confirmed": u.Confirmed,
		"profile":   u.Profile,
		"token":     u.GenerateJWT(),
	}, nil
}

func (u *User) GenerateJWT() (token string) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":     u.Email,
		"confirmed": u.Confirmed,
		"profile":   u.Profile,
	})
	token, err := t.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		panic(err)
	}
	return
}

func (u *User) Save() error {
	row := db.PrepareRow(`INSERT INTO users 
      (firstname, lastname, pseudo, email, password, created_at, updated_at, confirmed, profile, confirmation_token) 
      VALUES ($1, $2, $3, $4, $5, $6, $6, $7, $7, $8) RETURNING *;`,
		u.Firstname,
		u.Lastname,
		u.Pseudo,
		u.Email,
		u.Password,
		time.Now(),
		false,
		u.ConfirmationToken,
	)
	return row.Scan(u.All()...)
}

func (u *User) FindByTokenAndUpdate(token string) (err error) {
	row := db.PrepareRow("UPDATE users SET confirmation_token = '', confirmed = true WHERE confirmation_token = $1 RETURNING *;", token)
	return row.Scan(u.All()...)
}

func (u *User) SetConfirmationToken() {
	u.ConfirmationToken = u.GenerateJWT()
}

func (u *User) GenerateConfirmationUrl() string {
	return `http://localhost:3000/confirmation/` + u.ConfirmationToken
}

func (u *User) generateResetPasswordToken() string {

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, TokenClaims{
		Id: u.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	})
	tokenStr, err := t.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		panic(err)
	}
	return tokenStr
}

func (u *User) GenerateResetPasswordLink() string {
	return `http://localhost:3000/reset_password/` + u.generateResetPasswordToken()
}

func (u *User) SetPassword(password string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	u.Password = string(b)
	return err
}

func (u *User) UpdatePassword(password string) error {
	if err := u.SetPassword(password); err != nil {
		return err
	}
	fmt.Println(u.Password)
	_, err := db.Prepare("UPDATE users set password = $1 where id = $2;", u.Password, u.Id)
	return err
}
