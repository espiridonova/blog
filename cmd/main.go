package main

import (
	"database/sql"
	"net/http"

	handlerpkg "github.com/espiridonova/blog/internal/handler"
	"github.com/espiridonova/blog/internal/repository"
	servicepkg "github.com/espiridonova/blog/internal/service"
	"github.com/sirupsen/logrus"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "article.db")
	if err != nil {
		logrus.Fatalf("failed connect to DB: %s", err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	service := servicepkg.NewService(repo)
	handler := handlerpkg.NewHandler(service)

	http.HandleFunc("/article/list", handler.ListHandler)
	http.HandleFunc("/article", handler.GetArticleByIDHandler)
	http.HandleFunc("/create", handler.CreateHandler)
	http.HandleFunc("/update", handler.UpdateHandler)
	http.HandleFunc("/delete", handler.DeleteHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
