package service

import (
	"github.com/espiridonova/blog/internal/model"
	"github.com/espiridonova/blog/internal/repository"
)

func convertToArticle(article *repository.DBArticle) *model.Article {
	return &model.Article{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Created: article.Created,
	}
}

func convertToDBArticle(article *model.Article) *repository.DBArticle {
	return &repository.DBArticle{
		ID:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Created: article.Created,
	}
}
