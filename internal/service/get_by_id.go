package service

import (
	"database/sql"
	"errors"

	"github.com/espiridonova/blog/internal/model"
)

func (s *Service) GetByID(id int64) (*model.Article, error) {
	dbArticle, err := s.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.NotFoundErr
		}
		return nil, err
	}

	return convertToArticle(dbArticle), nil
}
