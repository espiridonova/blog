package repository

import "database/sql"

func (r *Repository) Create(a *DBArticle) (int64, error) {
	res, err := r.db.Exec(`INSERT INTO article (title, content) VALUES (:title, :content)`,
		sql.Named("title", a.Title),
		sql.Named("content", a.Content))
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}
