package repository

import (
	"database/sql"
	"errors"
)

func (r *Repository) Update(a *DBArticle) error {
	if a == nil {
		return errors.New("invalid article")
	}

	_, err := r.db.Exec("UPDATE article SET title = :title, content = :content WHERE id = :ID",
		sql.Named("title", a.Title),
		sql.Named("content", a.Content),
		sql.Named("ID", a.ID))

	return err
}
