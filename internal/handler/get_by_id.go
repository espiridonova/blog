package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/espiridonova/blog/internal/model"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	article, err := h.service.GetByID(id)
	if err != nil {
		if errors.Is(err, model.NotFoundErr) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		logrus.Errorf("failed get by id article from DB: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(&article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		logrus.Errorf("error marshal: %s", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
