package models

import (
	"database/sql"
)

type UrlStr struct {
	Url string `json:"url"`
}

type UrlDB struct {
	DB *sql.DB
}

func (u *UrlDB) Insert(url, hash string) error {
	stmt := `INSERT INTO urls (url, hash)
			values ($1, $2)
			ON CONFLICT (url) DO NOTHING`
	_, err := u.DB.Exec(stmt, url, hash)
	if err != nil {
		return err
	}
	return nil
}

func (u *UrlDB) Retrieve(hash string) (string, error) {
	stmt := `SELECT url FROM urls 
		WHERE hash=$1`
	row := u.DB.QueryRow(stmt, hash)
	url := ""

	err := row.Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
