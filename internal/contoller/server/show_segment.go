package server

import (
	"encoding/json"
	"github.com/fanat1k666/backend-trainee-assignment-2023/internal/entity"
	"io"
	"net/http"
)

type ShowSegmentIn struct {
	UserId int `json:"user_id"`
}

type ShowSegmentOut struct {
	UserId   int      `json:"userId"`
	Segments []string `json:"segments"`
	Err      *string  `json:"err,omitempty"`
}

func (s *Server) HandleShowSegment(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	defer r.Body.Close()
	if err != nil {
		errorStr := err.Error()
		ans := ShowSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadGateway)
		return
	}

	in := ShowSegmentIn{}
	err = json.Unmarshal(body, &in)
	if err != nil {
		errorStr := err.Error()
		ans := ShowSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusBadRequest)
		return
	}

	eUs := entity.User{
		UserId: in.UserId,
	}

	rows, err := s.h.ShowSegment(eUs)
	if err != nil {
		errorStr := err.Error()
		ans := ShowSegmentOut{Err: &errorStr}
		s.SendAnswer(w, ans, http.StatusInternalServerError)
		return
	}
	var end []string
	var u int
	if len(rows) != 0 {
		u = rows[0].UserId
		for i := range rows {
			end = append(end, rows[i].Name)
		}
	}
	s.SendAnswer(w, ShowSegmentOut{UserId: u, Segments: end}, http.StatusOK)
}
