package repository

import (
	"database/sql"
)

// - создать в отдельном файле структуру для работы с БД (методы create, getByID, getList, delete, update)
type DBArticle struct {
	ID      int
	Title   string
	Content string
	Created int64
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
