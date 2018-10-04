package models

import (
	"matcha/db"
	"database/sql"
	"fmt"
	"sort"
	"database/sql/driver"
	"strconv"
	"strings"
	"log"
)

type Tags struct {
	Id  int    `json:"id"`
	Tag string `json:"tag"`
}

type TagsId []int64

func (t TagsId) Value() (driver.Value) {
	var strs []string
	for _, i := range t {
		strs = append(strs, strconv.FormatInt(i, 10))
	}
	return "{" + strings.Join(strs, ",") + "}"
}

type TagsValue string

func (t TagsValue) Value() (tid TagsId) {
	tid = TagsId{}
	var y []string
	s := strings.TrimPrefix(string(t), "{")
	s = strings.TrimSuffix(string(s), "}")
	y = strings.Split(s, ",")
	for _, v := range y {
		if v != "" {
			i, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			tid = append(tid, int64(i))
		}
	}
	return
}

type TagsArr []*Tags

func (t TagsArr) Len() int           { return len(t) }
func (t TagsArr) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t TagsArr) Less(i, j int) bool { return t[i].Tag < t[j].Tag }

func (t *Tags) Save() error {
	row := db.PrepareRow(`INSERT INTO tags (tag) VALUES ($1) RETURNING *;`, t.Tag)
	return row.Scan(t.All()...)
}

func (t *Tags) All() ([]interface{}) {
	return []interface{}{
		&t.Id,
		&t.Tag,
	}
}

func (t *Tags) FindAll() (ret TagsArr, err error) {
	ret = make([]*Tags, 0)
	rows, err := db.Query("SELECT * FROM TAGS")
	if err != nil {
		return
	}
	for rows.Next() {
		nt := NewTags()
		if err = rows.Scan(nt.All()...); err != nil {
			return
		}
		ret = append(ret, nt)
	}
	sort.Sort(interface{}(ret).(TagsArr))
	return
}

func NewTags() (*Tags) {
	return new(Tags)
}

func (t *Tags) FindById(id int) (err error) {
	return db.QueryRow("SELECT * FROM tags where id = $1", id).Scan(t.All()...)
}

func (t *Tags) Match(query string) (ret []*Tags, err error) {
	var rows *sql.Rows
	rows, err = db.Prepare("SELECT * FROM tags where tag LIKE '%" + query + "%'")
	if err != nil {
		return
	}
	for rows.Next() {
		nt := NewTags()
		if err = rows.Scan(nt.All()...); err != nil {
			return
		}
		fmt.Println(nt)
		ret = append(ret, nt)
	}
	return
}
