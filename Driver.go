package jk_error

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	_ "github.com/go-sql-driver/mysql"
)

func NewOpenDb() *sql.DB {
	db, err := sql.Open("mysql", "blog:blog@tcp(101.200.38.157:3306)/bxd_blog")
	if err != nil {
		panic(err)
	}
	return db
}

func MyselfQueryNoRows(ctx context.Context, id int) error {
	db := NewOpenDb()

	var title string
	err := db.QueryRowContext(ctx, "SELECT title FROM bxd_article WHERE id=?", id).Scan(&title)
	switch {
	case err == sql.ErrNoRows:
		return errors.Wrapf(err, "no user with id %d", id)
	case err != nil:
		return errors.Wrap(err, "query error")
	}

	return nil
}
