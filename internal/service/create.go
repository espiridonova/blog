package service

import (
	"fmt"

	"github.com/espiridonova/blog/internal/model"
)

func (s *Service) Create(article *model.Article) (int64, error) {
	dbArticle := convertToDBArticle(article)
	id, err := s.repo.Create(dbArticle)
	if err != nil {
		return 0, fmt.Errorf("create article failed: %w", err)
	}

	return id, nil
}
