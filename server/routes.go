package server

import (
	"net/http"
)

func (s *server) routes() {
	r := s.r

	// file servers
	r.PathPrefix("/public/").
		Handler(http.FileServer(http.FS(s.fs)))

	r.HandleFunc("/", s.handleView("home"))
}
