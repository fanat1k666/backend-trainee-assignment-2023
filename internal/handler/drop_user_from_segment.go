package handler

import (
	"fmt"

	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/entity"
)

func (h *Handler) DropUserFromSegment(user entity.User, segment entity.Segment) error {
	err := h.us.DropUserFromSegment(user.UserId, segment.Name)
	if err != nil {
		return fmt.Errorf("can't drop segment: %w", err)
	}

	return nil
}
