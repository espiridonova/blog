package service

import (
	"github.com/espiridonova/blog/internal/model"
)

func (s *Service) Update(article *model.Article) error {
	dbArticle := convertToDBArticle(article)

	return s.repo.Update(dbArticle)
}
