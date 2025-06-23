package handler

import (
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

func (h *Handler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		logrus.Errorf("failed delete article from DB: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
