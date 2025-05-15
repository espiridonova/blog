package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/espiridonova/blog/internal/model"
	"github.com/sirupsen/logrus"
)

func (h *Handler) UpdateHandler(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Set("Content-Type", "application/json")
	if article.ID == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.Update(article)
	if err != nil {
		logrus.Errorf("failed update article: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("marshal error: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
