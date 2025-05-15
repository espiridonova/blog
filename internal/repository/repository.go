package repository

import (
	"database/sql"
	"time"
)

type DBArticle struct {
	ID      int64
	Title   string
	Content string
	Created *time.Time
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}
