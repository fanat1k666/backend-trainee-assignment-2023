package handler

import (
	"fmt"
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/entity"
)

func (h *Handler) UserSegment(user entity.SegmentToUser, names []string) error {
	err := h.us.UserSegment(user.UserId, names)
	if err != nil {
		return fmt.Errorf("can't create segment: %w", err)
	}

	return nil
}
