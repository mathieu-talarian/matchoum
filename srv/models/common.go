package models

type Queries interface {
	All() ([]interface{})
	FindById(id int) (err error)
}
