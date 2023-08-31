package server

import (
	"encoding/json"
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/entity"
	"io"
	"net/http"
)

type DropUserFromSegmentIn struct {
	UserId  int    `json:"user_id"`
	Segment string `json:"segment"`
}

type DropUserFromSegmentOut struct {
	Err *string `json:"err,omitempty"`
}

func (s *Server) HandleDropUserFromSegment(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	if err != nil {
		errorStr := err.Error()
		ans := DropUserFromSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadGateway)
		return
	}

	in := DropUserFromSegmentIn{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		errorStr := err.Error()
		ans := DropUserFromSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadRequest)
		return
	}

	eUs := entity.User{
		UserId: in.UserId,
	}
	eSeg := entity.Segment{
		Name: in.Segment,
	}

	err = s.h.DropUserFromSegment(eUs, eSeg)
	if err != nil {
		errorStr := err.Error()
		ans := DropUserFromSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusInternalServerError)
		return
	}

	s.SendAnswer(w, DropUserFromSegmentOut{}, http.StatusOK)
}
