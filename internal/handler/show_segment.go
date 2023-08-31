package handler

import (
	"fmt"
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/entity"
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/repository"
)

func (h *Handler) ShowSegment(user entity.User) ([]repository.ShowUsersSegment, error) {
	line, err := h.us.ShowSegment(user.UserId)
	if err != nil {
		return nil, fmt.Errorf("can't create segment: %w", err)
	}

	return line, nil
}
