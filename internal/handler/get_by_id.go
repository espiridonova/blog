package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) GetArticleByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	dbArticle, err := h.repo.GetByID(id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	article := convertToResp(dbArticle)
	resp, err := json.Marshal(&article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
