package repository

import (
	"database/sql"
	"fmt"
)

func (r *Repository) GetList(title string, limit, offset int64) ([]*DBArticle, error) {
	var args []any
	sqlQuery := "SELECT id, title, content, created FROM article"
	if title != "" {
		sqlQuery += " WHERE title LIKE :title"
		args = append(args, sql.Named("title", fmt.Sprintf("%%%s%%", title)))
	}

	sqlQuery += " ORDER BY created DESC LIMIT :limit OFFSET :offset"
	args = append(args, sql.Named("limit", limit), sql.Named("offset", offset))

	rows, err := r.db.Query(sqlQuery, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res []*DBArticle
	for rows.Next() {
		var art DBArticle
		err = rows.Scan(&art.ID, &art.Title, &art.Content, &art.Created)
		if err != nil {
			return nil, err
		}
		res = append(res, &art)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return res, nil
}
