package handler

import (
	"fmt"

	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/entity"
)

func (h *Handler) DropSegment(segment entity.Segment) error {
	err := h.us.DropSegment(segment.Name)
	if err != nil {
		return fmt.Errorf("can't drop segment: %w", err)
	}

	return nil
}
