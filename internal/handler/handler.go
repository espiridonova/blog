package handler

import (
	servicepkg "github.com/espiridonova/blog/internal/service"
)

type Handler struct {
	service *servicepkg.Service
}

func NewHandler(service *servicepkg.Service) *Handler {
	return &Handler{service: service}
}
