package server

import (
	"encoding/json"
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/entity"
	"io"
	"net/http"
)

type DropSegmentIn struct {
	Name string `json:"name"`
}

type DropSegmentOut struct {
	Err *string `json:"err,omitempty"`
}

func (s *Server) HandleDropSegment(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	if err != nil {
		errorStr := err.Error()
		ans := DropSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadGateway)
		return
	}

	in := DropSegmentIn{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		errorStr := err.Error()
		ans := DropSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadRequest)
		return
	}

	eSeg := entity.Segment{
		Name: in.Name,
	}

	err = s.h.DropSegment(eSeg)
	if err != nil {
		errorStr := err.Error()
		ans := DropSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusInternalServerError)
		return
	}

	s.SendAnswer(w, DropSegmentOut{}, http.StatusOK)
}
