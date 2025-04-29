package handler

import (
	"github.com/espiridonova/blog/internal/repository"

	"github.com/espiridonova/blog/internal/model"
)

type Handler struct {
	repo *repository.Repository
}

func NewHandler(repo *repository.Repository) *Handler {
	return &Handler{repo}
}

func convertToResp(article *repository.DBArticle) *model.Article {
	return &model.Article{
		Id:      article.ID,
		Title:   article.Title,
		Content: article.Content,
		Created: article.Created,
	}
}
