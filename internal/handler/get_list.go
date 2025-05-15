package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	maxLimit     = 100
	defaultLimit = 10
)

func (h *Handler) ListHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	limitParam := r.URL.Query().Get("limit")
	limit, _ := strconv.ParseInt(limitParam, 10, 64)
	if limit == 0 || limit > maxLimit {
		limit = defaultLimit
	}

	offsetParam := r.URL.Query().Get("offset")
	offset, _ := strconv.ParseInt(offsetParam, 10, 64)

	list, err := h.service.GetList(title, limit, offset)
	if err != nil {
		logrus.Errorf("failed get article list: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp, err := json.Marshal(&list)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
