package server

import (
	"encoding/json"
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/entity"
	"io"
	"net/http"
)

type UserSegmentIn struct {
	Segments []string `json:"segments"`
	UserId   int      `json:"user_id"`
}

type UserSegmentOut struct {
	Err *string `json:"err,omitempty"`
}

func (s *Server) HandleAddSegmentToUser(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	if err != nil {
		errorStr := err.Error()
		ans := UserSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadGateway)
		return
	}

	in := UserSegmentIn{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		errorStr := err.Error()
		ans := UserSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadRequest)
		return
	}

	eUs := entity.SegmentToUser{
		UserId: in.UserId,
	}

	err = s.h.UserSegment(eUs, in.Segments)
	if err != nil {
		errorStr := err.Error()
		ans := UserSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusInternalServerError)
		return
	}

	s.SendAnswer(w, UserSegmentOut{}, http.StatusOK)
}
