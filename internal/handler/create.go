package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/espiridonova/blog/internal/model"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	r.Body.Close()

	article := &model.Article{}
	err = json.Unmarshal(body, &article)
	if err != nil {
		logrus.Errorf("error unmarshal: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id, err := h.service.Create(article)
	if err != nil {
		logrus.Errorf("failed create article: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(&model.Article{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("error marshal: %s", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
