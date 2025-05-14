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
	stmt := `INSERT IGNORE INTO urls
			(url, hash, created)
			values
				(?, ?, UTC_TIMESTAMP())`
	_, err := u.DB.Exec(stmt, url, hash)
	if err != nil {
		return err
	}
	return nil
}

func (u *UrlDB) Retrieve(hash string) (string, error) {
	stmt := `SELECT url FROM urls 
		WHERE hash=?`
	row := u.DB.QueryRow(stmt, hash)
	url := ""

	err := row.Scan(&url)
	if err != nil {
		return "", err
	}
	return url, nil
}
