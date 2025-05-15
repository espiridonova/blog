package repository

import (
	"database/sql"
)

func (r *Repository) Delete(id int64) error {
	_, err := r.db.Exec("DELETE FROM article WHERE id = :id",
		sql.Named("id", id))

	return err
}
