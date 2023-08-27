package handler

import "github.com/fanat1k666/backend-trainee-assignment-2023/internal/repository"

type Handler struct {
	us repository.UserSegment
}

func NewHandler(us repository.UserSegment) *Handler {
	return &Handler{us: us}
}
