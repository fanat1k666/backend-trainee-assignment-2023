package server

import (
	"net/http"
)

func (s *Server) HandleSegment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.HandleCreateSegment(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
