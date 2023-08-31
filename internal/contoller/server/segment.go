package server

import (
	"net/http"
)

func (s *Server) HandleSegment(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.HandleCreateSegment(w, r)
	case http.MethodDelete:
		s.HandleDropSegment(w, r)
	case http.MethodGet:
		s.HandleShowSegment(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) HandleSegmentToUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		s.HandleAddSegmentToUser(w, r)
	case http.MethodDelete:
		s.HandleDropUserFromSegment(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
