package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/espiridonova/blog/internal/handler"
	"github.com/espiridonova/blog/internal/repository"
	"io"
	"net/http"
	"strconv"
	"time"

	_ "modernc.org/sqlite"
)

/*
1. GET /article - апи метод для получения списка статей (готов)

2. GET /article/{id} - апи метод для получения одной статьи по id (готов)

3. POST /article - апи метод для создания новой статьи, в body запроса передавать json вида: (готов)
	{
		"title": "Название новой статьи",
		"content": "Текст новой статьи"
	}

4. PUT /article/{id} - апи метод для изменения существующей статьи по id
	в body запроса передавать json вида:(готов)
	{
		"title": "Название новой статьи",
		"content": "Текст новой статьи"
	}

5. DELETE /article/{id} - апи метод для удаления статьи по id(готов)

6. Перенести хранение статей в базу данных (используем sqlite так же как в уроках практикума).
Что нужно сделать:
	- написать SQL запрос на создание таблицы "article" в БД (поля такие же как в структуре Article)(готово)
	- создать подключение к бд (одно на весь сервис)(готово)
	- создать в отдельном файле структуру для работы с БД (методы create, getByID, getList, delete, update)
	- добавить использование методов сервиса БД в хэндлеры api:
		listHandler
		getArticleByIDHandler
		createHandler
		putArticleByIDHandler
		deleteHandler
	- проверить что все работает через постман и что после перезапуска добавленные статьи не теряются
*/

type Article struct {
	Id      int       `json:"id"`
	Title   string    `json:"title"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

var list = map[int]Article{
	1: {
		Id:      1,
		Title:   "Симба серый прохвост",
		Content: "какой то текст",
		Created: time.Now(),
	},
	2: {
		Id:      2,
		Title:   "Артур очень сильно любит одну девушку",
		Content: "её зовут Лена",
		Created: time.Now(),
	},
	3: {
		Id:      3,
		Title:   "Артур уезжает от Лены",
		Content: "Артура отправляют на неделю в командировку, поэтому он оставляет Лену на 1 недел, во всем виноват тупой озон",
		Created: time.Now(),
	},
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(&list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body.Close()

	article := Article{}

	err = json.Unmarshal(body, &article)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	newID := len(list) + 1
	article.Id = newID
	article.Created = time.Now()

	list[newID] = article

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(string(rune(newID))))
}

func putArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body.Close()

	article := Article{}
	err = json.Unmarshal(body, &article)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	art, ok := list[article.Id]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
	}

	art.Title = article.Title
	art.Content = article.Content
	list[art.Id] = art

	w.WriteHeader(http.StatusOK)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	delete(list, id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func main() {
	db, err := sql.Open("sqlite", "article.db")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	handler := handler.NewHandler(repo)

	http.HandleFunc("/article/list", listHandler)
	http.HandleFunc("/article", handler.GetArticleByIDHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/update", putArticleByIDHandler)
	http.HandleFunc("/delete", deleteHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
