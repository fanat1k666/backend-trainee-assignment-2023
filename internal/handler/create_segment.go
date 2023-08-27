package handler

import (
	"fmt"

	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/entity"
)

func (h *Handler) CreateSegment(segment entity.Segment) error {
	err := h.us.CreateSegment(segment.Name)
	if err != nil {
		return fmt.Errorf("can't create segment: %w", err)
	}

	return nil
}
