package models

import (
	"matcha/db"
)

type Profile struct {
	id          int
	userId      int
	Gender      int `json:"gender"`
	Orientation int `json:"orientation"`
	TagsId      `json:"tags"`
	Bio         string `json:"bio"`
	Photos      string `json:"photos"` //TODO slice of strings
	tagsValue   TagsValue
	// TODO timestamps
}

func NewProfile() *Profile {
	return new(Profile)
}

func (p *Profile) UserId(uid int) {
	p.userId = uid
}

func (p *Profile) FindByUser(userEmail string) (err error) {
	row := db.PrepareRow("SELECT users.email, profile.* from profile LEFT JOIN users on profile .userid = users.id where email=$1;", userEmail)
	err = row.Scan(p.All()...)
	p.TagsId = p.tagsValue.Value()
	return
}

func (p *Profile) All() ([]interface{}) {
	var u interface{}
	return append([]interface{}{
		&u,
		&p.id,
		&p.userId,
		&p.Gender,
		&p.Orientation,
		&p.Bio,
		&p.tagsValue,
		&p.Photos,
	})
}

func (p *Profile) Save() (err error) {
	row := db.PrepareRow(`INSERT INTO profile (userid, 
gender, 
orientation, 
bio, 
tags, 
photos) 
VAlUES ($1, $2, $3, $4, $5, $6) RETURNING *;`, p.userId, p.Gender, p.Orientation, p.Bio,
		p.TagsId.Value(), p.Photos)
	return row.Scan(p.All()...)
}

func (p *Profile) Update() error {
	row := db.PrepareRow(`UPDATE profile SET (gender, 
orientation, 
bio, 
tags, 
photos) 
VAlUES ($2, $3, $4, $5, $6) WHERE userid = $1 RETURNING *;`, p.userId, p.Gender, p.Orientation, p.Bio,
		p.TagsId.Value(), p.Photos)
	return row.Scan(p.All()...)
}
