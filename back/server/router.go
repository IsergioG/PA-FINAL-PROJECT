package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) router() http.Handler {
	router := mux.NewRouter()
	router.Use(s.logger.RequestLogger)
	router.HandleFunc("/death", s.HandleKills).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/death/{id}", s.HandleKillsWithId).Methods(http.MethodGet)
	router.HandleFunc("/deathUpdate/{id}", s.handleUpdate).Methods(http.MethodPatch)
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("uploads/"))))
	return router
}
