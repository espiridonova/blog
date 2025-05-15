package service

import (
	"github.com/espiridonova/blog/internal/model"
)

func (s *Service) GetList(title string, limit, offset int64) ([]*model.Article, error) {
	dbList, err := s.repo.GetList(title, limit, offset)
	if err != nil {
		return nil, err
	}
	list := make([]*model.Article, 0, len(dbList))
	for _, a := range dbList {
		list = append(list, convertToArticle(a))
	}

	return list, nil
}
