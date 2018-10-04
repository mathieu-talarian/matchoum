package db

import "fmt"

// Queries struct
type Queries map[string]string

var q = Queries{
	"users": `CREATE TABLE IF NOT EXISTS users (
    id         SERIAL PRIMARY KEY,
    firstname  VARCHAR(255) not null,
    lastname   VARCHAR(255) not null,
    pseudo     VARCHAR(255) not  null,
    email       VARCHAR(255) unique,
    password     VARCHAR(500) not null,
    profile      BOOLEAN,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    confirmed BOOLEAN,
    confirmation_token VARCHAR(500)
  );`,
	"profile": `CREATE TABLE IF NOT EXISTS profile (
    id SERIAL PRIMARY KEY,
    userid integer unique,
    gender VARCHAR(1) not null,
    orientation VARCHAR(50) not null,
    bio VARCHAR(500),
    tags integer[],
    photos VARCHAR(500),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
  );`,
	"tags": `CREATE TABLE IF NOT EXISTS tags (
    id SERIAL PRIMARY KEY,
    tag VARCHAR(200) unique
  );`,
}

func installDbs() {
	for k, v := range q {
		if _, err := db.Exec(v); err != nil {
			panic(err)
		}
		fmt.Println("Successfully installed", k)
	}
}
