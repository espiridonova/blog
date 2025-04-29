package repository

import "database/sql"

func (r *Repository) GetByID(id int) (*DBArticle, error) {
	a := DBArticle{}
	row := r.db.QueryRow(
		`SELECT  id, title, content, created FROM article WHERE id = :id`,
		sql.Named("id", id))
	err := row.Scan(&a.ID, &a.Title, &a.Content, &a.Created)

	return &a, err
}
