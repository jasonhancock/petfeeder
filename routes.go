package main

import (
	"net/http"
	"strings"
)

func (s *Server) buildRoutes() {

	s.router.HandleFunc("/feed", s.feedHandler).Methods("POST")
	s.router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir(s.assetsDir))))

	// catch-all default routing
	s.router.PathPrefix("/").Subrouter().HandleFunc("/{path:.*}", func(w http.ResponseWriter, r *http.Request) {
		name := strings.Replace(r.URL.Path[1:], "/", ".", -1)
		if name == "" {
			name = "index"
		}
		s.renderTemplate(w, r, name)
	})
}
